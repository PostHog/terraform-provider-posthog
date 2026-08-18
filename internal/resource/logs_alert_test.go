package resource

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// emptyStringSet builds an explicitly empty (non-null) set. Note that calling the
// variadic stringSet with no arguments yields a *null* set instead, since a
// zero-argument variadic call passes a nil slice.
func emptyStringSet(t *testing.T) types.Set {
	t.Helper()
	set, diags := types.SetValueFrom(context.Background(), types.StringType, []string{})
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
	return set
}

// window builds a blocked window from plain strings. Both fields are types.String, so
// unkeyed literals would silently invert if the struct's field order ever changed.
func window(start, end string) BlockedWindowTFModel {
	return BlockedWindowTFModel{Start: types.StringValue(start), End: types.StringValue(end)}
}

// TestNotNotifyingWarnings pins the states in which the provider warns that an alert exists
// but notifies nobody. Acceptance tests cannot reach these: a test alert never fires, and
// nothing can mark it broken on demand.
func TestNotNotifyingWarnings(t *testing.T) {
	tests := map[string]struct {
		state       *string
		wantWarning string
	}{
		"broken":     {state: util.StringPtr(stateBroken), wantWarning: "Log alert is not notifying"},
		"snoozed":    {state: util.StringPtr(stateSnoozed), wantWarning: "Log alert is snoozed"},
		"not firing": {state: util.StringPtr("not_firing")},
		"firing":     {state: util.StringPtr("firing")},
		"absent":     {state: nil},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			diags := notNotifyingWarnings(httpclient.LogsAlert{
				ID:    "01a00000-0000-0000-0000-00000000beef",
				Name:  util.StringPtr("Checkout errors"),
				State: test.state,
			})

			if test.wantWarning == "" {
				assert.Empty(t, diags.Warnings(), "unexpected warnings: %v", diags)
				return
			}
			require.Len(t, diags.Warnings(), 1)
			assert.Equal(t, test.wantWarning, diags.Warnings()[0].Summary())
			// The alert has to be identifiable: on import there is no config for
			// Terraform to attach the diagnostic to.
			assert.Contains(t, diags.Warnings()[0].Detail(), "Checkout errors")
			assert.Contains(t, diags.Warnings()[0].Detail(), "01a00000-0000-0000-0000-00000000beef")
		})
	}
}

// blockedWindowObjectType reuses the production attribute map so a new field on
// BlockedWindowTFModel cannot leave the tests exercising a stale shape.
func blockedWindowObjectType() types.ObjectType {
	return types.ObjectType{AttrTypes: blockedWindowAttrTypes}
}

func blockedWindowSet(t *testing.T, windows ...BlockedWindowTFModel) types.Set {
	t.Helper()
	set, diags := types.SetValueFrom(context.Background(), blockedWindowObjectType(), windows)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
	return set
}

// emptyBlockedWindowSet builds an explicitly empty (non-null) set. Calling the variadic
// blockedWindowSet with no arguments would yield a null set instead, since a
// zero-argument variadic call passes a nil slice.
func emptyBlockedWindowSet(t *testing.T) types.Set {
	t.Helper()
	set, diags := types.SetValueFrom(context.Background(), blockedWindowObjectType(), []BlockedWindowTFModel{})
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
	return set
}

// PostHog echoes back only the filter keys that are set, so the filters the user
// never configured must land as null rather than as empty collections.
func TestLogsAlertMapResponseToModel_AbsentFilterKeysAreNull(t *testing.T) {
	model := LogsAlertTFModel{}

	resp := httpclient.LogsAlert{
		ID:      "019dbe94-cec8-781b-9470-4a970cd69ebf",
		Name:    util.StringPtr("Contour Errors"),
		Enabled: util.BoolPtr(true),
		State:   util.StringPtr("not_firing"),
		Filters: &httpclient.LogsAlertFilters{
			ServiceNames: []string{"contour"},
		},
		ThresholdCount:    util.Int64Ptr(10),
		ThresholdOperator: util.StringPtr("above"),
		WindowMinutes:     util.Int64Ptr(10),
	}

	diags := LogsAlertOps{}.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	assert.Equal(t, types.StringValue("019dbe94-cec8-781b-9470-4a970cd69ebf"), model.ID)
	assert.Equal(t, stringSet(t, "contour"), model.ServiceNames)
	assert.True(t, model.SeverityLevels.IsNull(), "severity_levels was never configured, so it must be null")
	assert.True(t, model.FilterGroupJSON.IsNull(), "filter_group_json was never configured, so it must be null")
	assert.True(t, model.BlockedWindows.IsNull(), "no schedule_restriction means quiet hours are off")
	assert.Equal(t, types.Int64Value(10), model.ThresholdCount)
	assert.Equal(t, types.StringValue("not_firing"), model.State)
}

// A user who writes `severity_levels = []` must keep an empty set in state; flipping
// it to null would show as drift on every plan.
func TestLogsAlertMapResponseToModel_PreservesConfiguredEmptySet(t *testing.T) {
	model := LogsAlertTFModel{
		SeverityLevels: emptyStringSet(t),
	}

	diags := LogsAlertOps{}.MapResponseToModel(context.Background(), httpclient.LogsAlert{
		ID:      "019dbe94-cec8-781b-9470-4a970cd69ebf",
		Filters: &httpclient.LogsAlertFilters{ServiceNames: []string{"api"}},
	}, &model)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	assert.False(t, model.SeverityLevels.IsNull())
	assert.Empty(t, model.SeverityLevels.Elements())
}

// PostHog re-serializes filterGroup through its own model and adds keys the user never
// wrote — a leaf sent as {key,type,operator,value} reads back with "label":null (verified
// against a live alert). jsontypes.Normalized ignores key order and whitespace but NOT
// added keys, so without projection the state written after apply differs from the config
// and Terraform aborts with "Provider produced inconsistent result after apply".
//
// The response here is a realistic API payload, deliberately NOT an echo of the request:
// echoing the request back cannot detect server-side normalization by construction.
func TestLogsAlertMapResponseToModel_StripsServerAddedFilterGroupFields(t *testing.T) {
	userFilterGroup := `{"type":"AND","values":[{"type":"AND","values":[{"type":"log_attribute","key":"status_code","operator":"exact","value":["500"]}]}]}`

	model := LogsAlertTFModel{
		FilterGroupJSON: jsontypes.NewNormalizedValue(userFilterGroup),
	}

	diags := LogsAlertOps{}.MapResponseToModel(context.Background(), httpclient.LogsAlert{
		ID: "019dbe94-cec8-781b-9470-4a970cd69ebf",
		Filters: &httpclient.LogsAlertFilters{
			FilterGroup: map[string]any{
				"type": "AND",
				"values": []any{
					map[string]any{
						"type": "AND",
						"values": []any{
							map[string]any{
								"type":     "log_attribute",
								"key":      "status_code",
								"operator": "exact",
								"value":    []any{"500"},
								"label":    nil,
							},
						},
					},
				},
			},
		},
	}, &model)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	assert.JSONEq(t, userFilterGroup, model.FilterGroupJSON.ValueString(),
		"server-added keys must be projected away so state matches config")
}

// On import there is no prior config, so the whole API filter group is adopted verbatim.
func TestLogsAlertMapResponseToModel_ImportAdoptsFullFilterGroup(t *testing.T) {
	model := LogsAlertTFModel{}

	diags := LogsAlertOps{}.MapResponseToModel(context.Background(), httpclient.LogsAlert{
		ID: "019dbe94-cec8-781b-9470-4a970cd69ebf",
		Filters: &httpclient.LogsAlertFilters{
			FilterGroup: map[string]any{"type": "AND", "values": []any{}},
		},
	}, &model)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	assert.JSONEq(t, `{"type":"AND","values":[]}`, model.FilterGroupJSON.ValueString())
}

// name is Optional+Computed because PostHog defaults it to "Untitled alert". A config
// that omits name must be able to adopt that default without failing the apply.
func TestLogsAlertMapResponseToModel_AdoptsServerDefaultedName(t *testing.T) {
	model := LogsAlertTFModel{}

	diags := LogsAlertOps{}.MapResponseToModel(context.Background(), httpclient.LogsAlert{
		ID:      "019dbe94-cec8-781b-9470-4a970cd69ebf",
		Name:    util.StringPtr("Untitled alert"),
		Filters: &httpclient.LogsAlertFilters{ServiceNames: []string{"api"}},
	}, &model)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	assert.Equal(t, types.StringValue("Untitled alert"), model.Name)
}

// An explicitly empty blocked_windows set must stay empty rather than flipping to null,
// mirroring how the string sets are handled.
func TestLogsAlertMapResponseToModel_PreservesConfiguredEmptyBlockedWindows(t *testing.T) {
	model := LogsAlertTFModel{
		BlockedWindows: emptyBlockedWindowSet(t),
	}

	diags := LogsAlertOps{}.MapResponseToModel(context.Background(), httpclient.LogsAlert{
		ID:                  "019dbe94-cec8-781b-9470-4a970cd69ebf",
		ScheduleRestriction: nil,
	}, &model)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	assert.False(t, model.BlockedWindows.IsNull(), "a configured empty set must not read back as null")
	assert.Empty(t, model.BlockedWindows.Elements())
}

func TestLogsAlertBuildCreateRequest_RejectsInvalidFilterGroupJSON(t *testing.T) {
	model := LogsAlertTFModel{
		FilterGroupJSON: jsontypes.NewNormalizedValue(`{"type":`),
	}

	_, diags := LogsAlertOps{}.BuildCreateRequest(context.Background(), model)
	assert.True(t, diags.HasError(), "expected a diagnostic for unparseable filter_group_json")
}

// filters is a whole-object replace, so it must always be sent — otherwise a draft
// alert can never have its filters cleared.
func TestLogsAlertBuildCreateRequest_AlwaysSendsFilters(t *testing.T) {
	req, diags := LogsAlertOps{}.BuildCreateRequest(context.Background(), LogsAlertTFModel{})
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	require.NotNil(t, req.Filters)
	assert.Nil(t, req.Filters.SeverityLevels)
	assert.Nil(t, req.Filters.ServiceNames)
	assert.Nil(t, req.Filters.FilterGroup)
	assert.Nil(t, req.ScheduleRestriction, "no blocked_windows must send null to clear quiet hours")
}

func TestLogsAlertBuildCreateRequest_MapsBlockedWindows(t *testing.T) {
	model := LogsAlertTFModel{
		SeverityLevels: stringSet(t, "error"),
		BlockedWindows: blockedWindowSet(t, BlockedWindowTFModel{
			Start: types.StringValue("22:00"),
			End:   types.StringValue("23:30"),
		}),
	}

	req, diags := LogsAlertOps{}.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	require.NotNil(t, req.ScheduleRestriction)
	require.Len(t, req.ScheduleRestriction.BlockedWindows, 1)
	assert.Equal(t, "22:00", req.ScheduleRestriction.BlockedWindows[0].Start)
	assert.Equal(t, "23:30", req.ScheduleRestriction.BlockedWindows[0].End)
	assert.Equal(t, []string{"error"}, req.Filters.SeverityLevels)
}

// Plan-time validation has to reason about the value PostHog will actually end up with.
// Neither plan nor config alone identifies it:
//   - create, attribute omitted    → plan unknown, config null  → server default applies
//   - create, unresolved reference → plan unknown, config unknown → nothing can be concluded
//   - update, attribute omitted    → plan carries the prior state value (Optional+Computed)
//
// Getting this wrong is what produced both a false "never fires" rejection on update and a
// skipped check on create, so each case is pinned here.
func TestValidateLogsAlertPlan(t *testing.T) {
	withFilter := func(m LogsAlertTFModel) LogsAlertTFModel {
		m.SeverityLevels = stringSet(t, "error")
		return m
	}

	tests := []struct {
		name      string
		plan      LogsAlertTFModel
		config    LogsAlertTFModel
		expectErr string
	}{
		{
			name:   "minimal valid config",
			plan:   withFilter(LogsAlertTFModel{}),
			config: withFilter(LogsAlertTFModel{}),
		},
		{
			name: "datapoints exceeds periods",
			plan: withFilter(LogsAlertTFModel{
				EvaluationPeriods: types.Int64Value(2),
				DatapointsToAlarm: types.Int64Value(5),
			}),
			config: withFilter(LogsAlertTFModel{
				EvaluationPeriods: types.Int64Value(2),
				DatapointsToAlarm: types.Int64Value(5),
			}),
			expectErr: "Alert can never fire",
		},
		{
			// Create with evaluation_periods omitted: the server default of 1 applies, so
			// datapoints_to_alarm = 5 can never be reached.
			name: "create: datapoints exceeds omitted periods default",
			plan: withFilter(LogsAlertTFModel{
				EvaluationPeriods: types.Int64Unknown(),
				DatapointsToAlarm: types.Int64Value(5),
			}),
			config: withFilter(LogsAlertTFModel{
				DatapointsToAlarm: types.Int64Value(5),
			}),
			expectErr: "Alert can never fire",
		},
		{
			// Update with the evaluation_periods line removed: the attribute keeps its last
			// applied value (10), which the plan carries, so this is valid and must not be
			// rejected against the server default of 1.
			name: "update: omitted periods keeps its prior value",
			plan: withFilter(LogsAlertTFModel{
				EvaluationPeriods: types.Int64Value(10),
				DatapointsToAlarm: types.Int64Value(5),
			}),
			config: withFilter(LogsAlertTFModel{
				DatapointsToAlarm: types.Int64Value(5),
			}),
		},
		{
			// Mirror case: the retained value is the one that breaks, and must be caught.
			name: "update: retained periods below datapoints still fails",
			plan: withFilter(LogsAlertTFModel{
				EvaluationPeriods: types.Int64Value(2),
				DatapointsToAlarm: types.Int64Value(5),
			}),
			config: withFilter(LogsAlertTFModel{
				EvaluationPeriods: types.Int64Value(2),
			}),
			expectErr: "Alert can never fire",
		},
		{
			// Cannot conclude anything: the reference may resolve to any value.
			name: "unresolved periods skips the check",
			plan: withFilter(LogsAlertTFModel{
				EvaluationPeriods: types.Int64Unknown(),
				DatapointsToAlarm: types.Int64Value(5),
			}),
			config: withFilter(LogsAlertTFModel{
				EvaluationPeriods: types.Int64Unknown(),
				DatapointsToAlarm: types.Int64Value(5),
			}),
		},
		{
			name: "below zero can never fire",
			plan: withFilter(LogsAlertTFModel{
				ThresholdOperator: types.StringValue("below"),
				ThresholdCount:    types.Int64Value(0),
			}),
			config: withFilter(LogsAlertTFModel{
				ThresholdOperator: types.StringValue("below"),
				ThresholdCount:    types.Int64Value(0),
			}),
			expectErr: "Alert can never fire",
		},
		{
			// Update: threshold_count = 0 retained from state, operator newly set to below.
			// Resolving the omitted count to the server default of 100 would miss this.
			name: "update: retained zero count with below operator fails",
			plan: withFilter(LogsAlertTFModel{
				ThresholdOperator: types.StringValue("below"),
				ThresholdCount:    types.Int64Value(0),
			}),
			config: withFilter(LogsAlertTFModel{
				ThresholdOperator: types.StringValue("below"),
			}),
			expectErr: "Alert can never fire",
		},
		{
			name: "above zero is valid",
			plan: withFilter(LogsAlertTFModel{
				ThresholdOperator: types.StringValue("above"),
				ThresholdCount:    types.Int64Value(0),
			}),
			config: withFilter(LogsAlertTFModel{
				ThresholdOperator: types.StringValue("above"),
				ThresholdCount:    types.Int64Value(0),
			}),
		},
		{
			name:      "enabled alert with no filters",
			plan:      LogsAlertTFModel{},
			config:    LogsAlertTFModel{},
			expectErr: "no filters",
		},
		{
			name:   "disabled alert may omit filters",
			plan:   LogsAlertTFModel{Enabled: types.BoolValue(false)},
			config: LogsAlertTFModel{Enabled: types.BoolValue(false)},
		},
		{
			// Update: the enabled = false line removed. The alert stays disabled, so the
			// filter rule must not kick in against the server default of true.
			name:   "update: omitted enabled keeps its prior false value",
			plan:   LogsAlertTFModel{Enabled: types.BoolValue(false)},
			config: LogsAlertTFModel{},
		},
		{
			// Regression: an unknown set has zero elements but is not empty. Treating it as
			// empty blocked valid configs whose filters come from another resource.
			name:   "unknown severity levels does not trip the filter check",
			plan:   LogsAlertTFModel{SeverityLevels: types.SetUnknown(types.StringType)},
			config: LogsAlertTFModel{SeverityLevels: types.SetUnknown(types.StringType)},
		},
		{
			name:   "unknown service names does not trip the filter check",
			plan:   LogsAlertTFModel{ServiceNames: types.SetUnknown(types.StringType)},
			config: LogsAlertTFModel{ServiceNames: types.SetUnknown(types.StringType)},
		},
		{
			name:   "unknown filter group does not trip the filter check",
			plan:   LogsAlertTFModel{FilterGroupJSON: jsontypes.NewNormalizedUnknown()},
			config: LogsAlertTFModel{FilterGroupJSON: jsontypes.NewNormalizedUnknown()},
		},
		{
			name:      "explicitly empty sets still trip the filter check",
			plan:      LogsAlertTFModel{SeverityLevels: emptyStringSet(t), ServiceNames: emptyStringSet(t)},
			config:    LogsAlertTFModel{SeverityLevels: emptyStringSet(t), ServiceNames: emptyStringSet(t)},
			expectErr: "no filters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := validateLogsAlertPlan(context.Background(), tt.plan, tt.config)
			if tt.expectErr == "" {
				assert.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
				return
			}
			require.True(t, diags.HasError(), "expected a diagnostic")
			assert.Contains(t, diags.Errors()[0].Summary(), tt.expectErr)
		})
	}
}

func TestValidateBlockedWindows(t *testing.T) {
	tests := []struct {
		name      string
		windows   []BlockedWindowTFModel
		expectErr string
	}{
		{
			name:    "single valid window",
			windows: []BlockedWindowTFModel{window("22:00", "23:00")},
		},
		{
			name:    "window crossing midnight",
			windows: []BlockedWindowTFModel{window("22:00", "06:00")},
		},
		{
			name: "two non-overlapping windows",
			windows: []BlockedWindowTFModel{
				window("01:00", "02:00"),
				window("03:00", "04:00"),
			},
		},
		{
			name:      "window shorter than 30 minutes",
			windows:   []BlockedWindowTFModel{window("22:00", "22:15")},
			expectErr: "too short",
		},
		{
			// Must not report "spans 1440 minutes" — the midnight-wrap correction would
			// otherwise turn a zero-length window into a full day.
			name:      "zero-length window",
			windows:   []BlockedWindowTFModel{window("22:00", "22:00")},
			expectErr: "covers no time",
		},
		{
			name:    "exactly 30 minutes is allowed",
			windows: []BlockedWindowTFModel{window("01:00", "01:30")},
		},
		{
			name:      "29 minutes is rejected",
			windows:   []BlockedWindowTFModel{window("01:00", "01:29")},
			expectErr: "too short",
		},
		{
			name:    "30-minute window spanning midnight",
			windows: []BlockedWindowTFModel{window("23:45", "00:15")},
		},
		// PostHog merges on `next.start <= prev.end`, so windows that merely touch are
		// stored as one. Verified against validate_and_normalize_schedule_restriction:
		// 01:00-02:00 plus 02:00-03:00 is saved as a single 01:00-03:00 window.
		{
			name: "adjacent windows are merged, so they conflict",
			windows: []BlockedWindowTFModel{
				window("01:00", "02:00"),
				window("02:00", "03:00"),
			},
			expectErr: "overlap",
		},
		{
			name: "two windows with a gap between them",
			windows: []BlockedWindowTFModel{
				window("01:00", "05:00"),
				window("12:00", "13:00"),
			},
		},
		// A window crossing midnight is re-encoded as one wrapping window only when it is
		// the whole configuration. Alongside another window PostHog stores it as two, so
		// 22:00-07:00 plus 12:00-13:00 reads back as three windows.
		{
			name: "wrapping window alone is fine",
			windows: []BlockedWindowTFModel{
				window("22:00", "07:00"),
			},
		},
		{
			name: "wrapping window alongside another window",
			windows: []BlockedWindowTFModel{
				window("22:00", "07:00"),
				window("12:00", "13:00"),
			},
			expectErr: "must be the only window",
		},
		// Blocking both sides of midnight with two separate windows is the same shape:
		// PostHog recombines them into a single 22:00-06:00 window.
		{
			name: "window starting at midnight plus one ending at midnight",
			windows: []BlockedWindowTFModel{
				window("00:00", "06:00"),
				window("22:00", "00:00"),
			},
			expectErr: "meeting at midnight",
		},
		// PostHog only rejoins a midnight pair while it is the whole timeline. A third
		// window anywhere in the day leaves all of them stored exactly as written.
		{
			name: "midnight pair with a third window",
			windows: []BlockedWindowTFModel{
				window("22:00", "00:00"),
				window("00:00", "07:00"),
				window("12:00", "13:00"),
			},
		},
		{
			name: "midnight pair with two other windows",
			windows: []BlockedWindowTFModel{
				window("00:00", "06:00"),
				window("08:00", "09:00"),
				window("12:00", "13:00"),
				window("19:00", "00:00"),
			},
		},
		// Ending at 23:59 stops short of midnight, so nothing is recombined.
		{
			name: "window starting at midnight plus one ending at 23:59",
			windows: []BlockedWindowTFModel{
				window("00:00", "06:00"),
				window("22:00", "23:59"),
			},
		},
		// A window ending exactly at midnight does not wrap into the next morning, so it
		// coexists with a daytime window.
		{
			name: "window ending at midnight plus a daytime window",
			windows: []BlockedWindowTFModel{
				window("19:00", "00:00"),
				window("12:00", "13:00"),
			},
		},
		{
			name: "two wrapping windows that overlap",
			windows: []BlockedWindowTFModel{
				window("22:00", "02:00"),
				window("23:00", "03:00"),
			},
			expectErr: "overlap",
		},
		{
			name: "overlapping windows",
			windows: []BlockedWindowTFModel{
				window("01:00", "03:00"),
				window("02:00", "04:00"),
			},
			expectErr: "overlap",
		},
		{
			name: "overlap only after midnight wrap",
			windows: []BlockedWindowTFModel{
				window("22:00", "06:00"),
				window("05:00", "07:00"),
			},
			expectErr: "overlap",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := validateBlockedWindows(context.Background(), blockedWindowSet(t, tt.windows...))
			if tt.expectErr == "" {
				assert.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
				return
			}
			require.True(t, diags.HasError(), "expected a diagnostic")
			assert.Contains(t, strings.ToLower(diags.Errors()[0].Summary()), tt.expectErr)
		})
	}
}

// A null set means quiet hours are off, which is always valid.
func TestValidateBlockedWindows_NullIsValid(t *testing.T) {
	diags := validateBlockedWindows(context.Background(), types.SetNull(blockedWindowObjectType()))
	assert.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
}

func TestNonEmptyJSONObjectValidator(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		expectErr bool
	}{
		{name: "valid object", value: `{"type":"AND","values":[]}`},
		{name: "empty object", value: `{}`, expectErr: true},
		{name: "json null", value: `null`, expectErr: true},
		{name: "array", value: `[1,2]`, expectErr: true},
		{name: "scalar", value: `"hello"`, expectErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &validator.StringResponse{}
			nonEmptyJSONObjectValidator{}.ValidateString(context.Background(), validator.StringRequest{
				Path:        path.Root("filter_group_json"),
				ConfigValue: types.StringValue(tt.value),
			}, resp)
			assert.Equal(t, tt.expectErr, resp.Diagnostics.HasError(), "diags: %v", resp.Diagnostics)
		})
	}
}

func TestParseHHMM(t *testing.T) {
	for _, tt := range []struct {
		in   string
		want int
		ok   bool
	}{
		{"00:00", 0, true},
		{"06:30", 390, true},
		{"23:59", 1439, true},
		{"nope", 0, false},
	} {
		got, ok := parseHHMM(tt.in)
		assert.Equal(t, tt.ok, ok, "input %q", tt.in)
		if tt.ok {
			assert.Equal(t, tt.want, got, "input %q", tt.in)
		}
	}
}

func TestLogsAlertMapResponseToModel_MapsBlockedWindows(t *testing.T) {
	model := LogsAlertTFModel{}

	diags := LogsAlertOps{}.MapResponseToModel(context.Background(), httpclient.LogsAlert{
		ID: "019dbe94-cec8-781b-9470-4a970cd69ebf",
		ScheduleRestriction: &httpclient.LogsAlertSchedule{
			BlockedWindows: []httpclient.LogsAlertBlockedWindow{{Start: "22:00", End: "23:30"}},
		},
	}, &model)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	expected := blockedWindowSet(t, BlockedWindowTFModel{
		Start: types.StringValue("22:00"),
		End:   types.StringValue("23:30"),
	})
	assert.Equal(t, expected, model.BlockedWindows)
}

// snooze_until is managed only when the configuration declares it. A snooze an operator
// sets in the PostHog UI must not be pulled into state, because the next apply would then
// revert it.
func TestLogsAlertMapResponseToModel_SnoozeUntil(t *testing.T) {
	t.Run("adopted when declared", func(t *testing.T) {
		model := LogsAlertTFModel{SnoozeUntil: types.StringValue("2026-01-01T00:00:00Z")}

		diags := LogsAlertOps{}.MapResponseToModel(context.Background(), httpclient.LogsAlert{
			ID:          "019dbe94-cec8-781b-9470-4a970cd69ebf",
			SnoozeUntil: util.StringPtr("2026-02-01T09:00:00Z"),
		}, &model)
		require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

		assert.Equal(t, types.StringValue("2026-02-01T09:00:00Z"), model.SnoozeUntil)
	})

	t.Run("left null when the config never declared it", func(t *testing.T) {
		model := LogsAlertTFModel{}

		diags := LogsAlertOps{}.MapResponseToModel(context.Background(), httpclient.LogsAlert{
			ID:          "019dbe94-cec8-781b-9470-4a970cd69ebf",
			SnoozeUntil: util.StringPtr("2026-02-01T09:00:00Z"),
		}, &model)
		require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

		assert.True(t, model.SnoozeUntil.IsNull(),
			"an out-of-band snooze must not be adopted, or the next apply would revert it")
	})
}

// An undeclared snooze is omitted from the request entirely, so the server keeps its own.
func TestLogsAlertBuildCreateRequest_SnoozeUntil(t *testing.T) {
	req, diags := LogsAlertOps{}.BuildCreateRequest(context.Background(), LogsAlertTFModel{
		SeverityLevels: stringSet(t, "error"),
	})
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
	assert.Nil(t, req.SnoozeUntil)

	body, err := json.Marshal(req)
	require.NoError(t, err)
	assert.NotContains(t, string(body), "snooze_until")
}

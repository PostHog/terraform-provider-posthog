package resource

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
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

func blockedWindowObjectType() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"start": types.StringType,
		"end":   types.StringType,
	}}
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

// Plan-time validation runs against the config, where an omitted Optional+Computed
// attribute is null (server default applies) and an unresolved reference is unknown
// (nothing can be concluded yet). Conflating the two either blocks valid configs or
// silently skips the check on create, so both cases are covered explicitly.
func TestValidateLogsAlertConfig(t *testing.T) {
	withFilter := func(m LogsAlertTFModel) LogsAlertTFModel {
		m.SeverityLevels = stringSet(t, "error")
		return m
	}

	tests := []struct {
		name      string
		config    LogsAlertTFModel
		expectErr string
	}{
		{
			name:   "minimal valid config",
			config: withFilter(LogsAlertTFModel{}),
		},
		{
			name: "datapoints exceeds periods",
			config: withFilter(LogsAlertTFModel{
				EvaluationPeriods: types.Int64Value(2),
				DatapointsToAlarm: types.Int64Value(5),
			}),
			expectErr: "Alert can never fire",
		},
		{
			// The create path: evaluation_periods omitted, so the server default of 1
			// applies and datapoints_to_alarm = 5 can never be reached.
			name: "datapoints exceeds omitted periods default",
			config: withFilter(LogsAlertTFModel{
				DatapointsToAlarm: types.Int64Value(5),
			}),
			expectErr: "Alert can never fire",
		},
		{
			// Cannot conclude anything: the reference may resolve to any value.
			name: "unknown periods skips the check",
			config: withFilter(LogsAlertTFModel{
				EvaluationPeriods: types.Int64Unknown(),
				DatapointsToAlarm: types.Int64Value(5),
			}),
		},
		{
			name: "below zero can never fire",
			config: withFilter(LogsAlertTFModel{
				ThresholdOperator: types.StringValue("below"),
				ThresholdCount:    types.Int64Value(0),
			}),
			expectErr: "Alert can never fire",
		},
		{
			name: "above zero is valid",
			config: withFilter(LogsAlertTFModel{
				ThresholdOperator: types.StringValue("above"),
				ThresholdCount:    types.Int64Value(0),
			}),
		},
		{
			name:      "enabled alert with no filters",
			config:    LogsAlertTFModel{},
			expectErr: "no filters",
		},
		{
			name:   "draft alert may omit filters",
			config: LogsAlertTFModel{Enabled: types.BoolValue(false)},
		},
		{
			// Regression: an unknown set has zero elements but is not empty. Treating it
			// as empty blocked valid configs whose filters come from another resource.
			name:   "unknown severity levels does not trip the filter check",
			config: LogsAlertTFModel{SeverityLevels: types.SetUnknown(types.StringType)},
		},
		{
			name:   "unknown service names does not trip the filter check",
			config: LogsAlertTFModel{ServiceNames: types.SetUnknown(types.StringType)},
		},
		{
			name:   "unknown filter group does not trip the filter check",
			config: LogsAlertTFModel{FilterGroupJSON: jsontypes.NewNormalizedUnknown()},
		},
		{
			name:      "explicitly empty sets still trip the filter check",
			config:    LogsAlertTFModel{SeverityLevels: emptyStringSet(t), ServiceNames: emptyStringSet(t)},
			expectErr: "no filters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := validateLogsAlertConfig(context.Background(), tt.config)
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
			windows: []BlockedWindowTFModel{{types.StringValue("22:00"), types.StringValue("23:00")}},
		},
		{
			name:    "window crossing midnight",
			windows: []BlockedWindowTFModel{{types.StringValue("22:00"), types.StringValue("06:00")}},
		},
		{
			name: "two non-overlapping windows",
			windows: []BlockedWindowTFModel{
				{types.StringValue("01:00"), types.StringValue("02:00")},
				{types.StringValue("03:00"), types.StringValue("04:00")},
			},
		},
		{
			name:      "window shorter than 30 minutes",
			windows:   []BlockedWindowTFModel{{types.StringValue("22:00"), types.StringValue("22:15")}},
			expectErr: "too short",
		},
		{
			// Must not report "spans 1440 minutes" — the midnight-wrap correction would
			// otherwise turn a zero-length window into a full day.
			name:      "zero-length window",
			windows:   []BlockedWindowTFModel{{types.StringValue("22:00"), types.StringValue("22:00")}},
			expectErr: "covers no time",
		},
		{
			name:    "exactly 30 minutes is allowed",
			windows: []BlockedWindowTFModel{{types.StringValue("01:00"), types.StringValue("01:30")}},
		},
		{
			name:      "29 minutes is rejected",
			windows:   []BlockedWindowTFModel{{types.StringValue("01:00"), types.StringValue("01:29")}},
			expectErr: "too short",
		},
		{
			name:    "30-minute window spanning midnight",
			windows: []BlockedWindowTFModel{{types.StringValue("23:45"), types.StringValue("00:15")}},
		},
		{
			name: "adjacent windows do not overlap",
			windows: []BlockedWindowTFModel{
				{types.StringValue("01:00"), types.StringValue("02:00")},
				{types.StringValue("02:00"), types.StringValue("03:00")},
			},
		},
		{
			name: "two wrapping windows that overlap",
			windows: []BlockedWindowTFModel{
				{types.StringValue("22:00"), types.StringValue("02:00")},
				{types.StringValue("23:00"), types.StringValue("03:00")},
			},
			expectErr: "overlap",
		},
		{
			name: "overlapping windows",
			windows: []BlockedWindowTFModel{
				{types.StringValue("01:00"), types.StringValue("03:00")},
				{types.StringValue("02:00"), types.StringValue("04:00")},
			},
			expectErr: "overlap",
		},
		{
			name: "overlap only after midnight wrap",
			windows: []BlockedWindowTFModel{
				{types.StringValue("22:00"), types.StringValue("06:00")},
				{types.StringValue("05:00"), types.StringValue("07:00")},
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

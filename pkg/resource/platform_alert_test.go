package resource

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testPlatformAlertID = "0190a000-0000-0000-0000-00000000a1e7"
	testLogsAlertID     = "0190b000-0000-0000-0000-00000000106s"
)

func keyList(keys ...string) types.List {
	values := make([]attr.Value, len(keys))
	for i, key := range keys {
		values[i] = types.StringValue(key)
	}
	return types.ListValueMust(types.StringType, values)
}

// platformAlertPlan is a plan as Terraform builds it from a config that sets only what the
// arguments name: every schema default is already applied.
func platformAlertPlan(t *testing.T, filters *PlatformAlertLogsFiltersTFModel) PlatformAlertTFModel {
	t.Helper()
	return PlatformAlertTFModel{
		SourceKind:        types.StringValue(sourceKindLogs),
		Name:              types.StringValue("Checkout errors"),
		Enabled:           types.BoolValue(true),
		EvaluationPeriods: types.Int64Value(1),
		DatapointsToAlarm: types.Int64Value(1),
		CooldownMinutes:   types.Int64Value(0),
		BlockedWindows:    types.SetNull(blockedWindowObjectType()),
		Grouping: &PlatformAlertGroupingTFModel{
			Mode:         types.StringValue(groupingModeSingle),
			Keys:         keyList(),
			MaxInstances: types.Int64Value(defaultMaxInstances),
			OnOverflow:   types.StringValue(onOverflowDegradeVisible),
		},
		Logs: &PlatformAlertLogsTFModel{
			Filters: filters,
			Condition: &PlatformAlertLogsConditionTFModel{
				ThresholdCount:    types.Int64Value(100),
				ThresholdOperator: types.StringValue("above"),
				WindowMinutes:     types.Int64Value(5),
			},
		},
	}
}

func severityFilter(t *testing.T, levels ...string) *PlatformAlertLogsFiltersTFModel {
	t.Helper()
	return &PlatformAlertLogsFiltersTFModel{
		SeverityLevels:  stringSet(t, levels...),
		ServiceNames:    types.SetNull(types.StringType),
		FilterGroupJSON: jsontypes.NewNormalizedNull(),
	}
}

// A draft alert with null for each filter key is refused by PostHog as having filters that narrow
// nothing, so a draft must send an empty filters object.
func TestPlatformAlertBuildCreateRequest_DraftSendsEmptyFilters(t *testing.T) {
	plan := platformAlertPlan(t, nil)
	plan.Enabled = types.BoolValue(false)

	req, diags := PlatformAlertOps{}.BuildCreateRequest(context.Background(), plan)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	body, err := json.Marshal(req)
	require.NoError(t, err)
	var decoded map[string]any
	require.NoError(t, json.Unmarshal(body, &decoded))

	sourceConfig := decoded["source_config"].(map[string]any)
	assert.Equal(t, map[string]any{}, sourceConfig["filters"])
	assert.Equal(t, map[string]any{"threshold_count": 100.0, "threshold_operator": "above", "window_minutes": 5.0}, sourceConfig["condition"])
	// PUT replaces every field, so none may be left out.
	for _, field := range []string{"name", "enabled", "grouping", "schedule_restriction", "evaluation_periods", "datapoints_to_alarm", "cooldown_minutes"} {
		assert.Contains(t, decoded, field)
	}
	assert.Equal(t, []any{}, decoded["grouping"].(map[string]any)["keys"])
}

// A moved alert must read back into exactly the model the rewritten config plans, or the first
// plan after the move shows a change.
func TestPlatformAlertMove_ReadsBackAsTheRewrittenConfigPlans(t *testing.T) {
	source := LogsAlertTFModel{
		Name:              types.StringValue("Checkout errors"),
		Enabled:           types.BoolValue(true),
		SeverityLevels:    stringSet(t, "error"),
		ServiceNames:      types.SetNull(types.StringType),
		FilterGroupJSON:   jsontypes.NewNormalizedNull(),
		ThresholdCount:    types.Int64Value(100),
		ThresholdOperator: types.StringValue("above"),
		WindowMinutes:     types.Int64Value(5),
		EvaluationPeriods: types.Int64Value(1),
		DatapointsToAlarm: types.Int64Value(1),
		CooldownMinutes:   types.Int64Value(0),
		BlockedWindows:    types.SetNull(blockedWindowObjectType()),
		SnoozeUntil:       types.StringNull(),
		State:             types.StringValue("not_firing"),
	}
	source.ID = types.StringValue(testLogsAlertID)
	source.ProjectID = types.StringValue("2")

	moved := platformAlertFromLogsAlert(source)
	assert.Equal(t, legacyIDPrefix+testLogsAlertID, moved.ID.ValueString())
	assert.Equal(t, testLogsAlertID, moved.LegacyConfigurationID.ValueString())

	resp := httpclient.PlatformAlert{
		ID:                    testPlatformAlertID,
		Name:                  util.StringPtr("Checkout errors"),
		Enabled:               util.BoolPtr(true),
		SourceKind:            util.StringPtr(sourceKindLogs),
		SourceConfig:          json.RawMessage(`{"filters":{"severityLevels":["error"]},"condition":{"threshold_count":100,"threshold_operator":"above","window_minutes":5}}`),
		EvaluationPeriods:     util.Int64Ptr(1),
		DatapointsToAlarm:     util.Int64Ptr(1),
		CooldownMinutes:       util.Int64Ptr(0),
		Grouping:              &httpclient.PlatformAlertGrouping{Mode: groupingModeSingle, Keys: []string{}, MaxInstances: 100, OnOverflow: onOverflowDegradeVisible},
		LegacyConfigurationID: util.StringPtr(testLogsAlertID),
		Owner:                 util.StringPtr("platform"),
	}
	diags := PlatformAlertOps{}.MapResponseToModel(context.Background(), resp, &moved)
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

	want := platformAlertPlan(t, severityFilter(t, "error"))
	want.ID = types.StringValue(testPlatformAlertID)
	want.ProjectID = types.StringValue("2")
	want.LegacyConfigurationID = types.StringValue(testLogsAlertID)
	assert.Equal(t, want, moved)
}

func TestLogsSourceConfigToModel(t *testing.T) {
	tests := map[string]struct {
		raw         string
		current     *PlatformAlertLogsTFModel
		wantFilters *PlatformAlertLogsFiltersTFModel
	}{
		"draft with no filters stays without a filters object": {
			raw:     `{"filters":{},"condition":{"threshold_count":1,"threshold_operator":"above","window_minutes":5}}`,
			current: &PlatformAlertLogsTFModel{},
		},
		"import adopts stored filters": {
			raw:         `{"filters":{"severityLevels":["warn"]},"condition":{"threshold_count":1,"threshold_operator":"above","window_minutes":5}}`,
			wantFilters: severityFilter(t, "warn"),
		},
		"declared empty set stays empty": {
			raw: `{"filters":{"serviceNames":["api"]},"condition":{"threshold_count":1,"threshold_operator":"above","window_minutes":5}}`,
			current: &PlatformAlertLogsTFModel{Filters: &PlatformAlertLogsFiltersTFModel{
				SeverityLevels: emptyStringSet(t), ServiceNames: stringSet(t, "api"), FilterGroupJSON: jsontypes.NewNormalizedNull(),
			}},
			wantFilters: &PlatformAlertLogsFiltersTFModel{
				SeverityLevels: emptyStringSet(t), ServiceNames: stringSet(t, "api"), FilterGroupJSON: jsontypes.NewNormalizedNull(),
			},
		},
		"server-added filter group fields are not tracked": {
			raw: `{"filters":{"filterGroup":{"type":"AND","values":[{"key":"status","type":"log_attribute","operator":"exact","value":["500"],"label":null}]}},"condition":{"threshold_count":1,"threshold_operator":"above","window_minutes":5}}`,
			current: &PlatformAlertLogsTFModel{Filters: &PlatformAlertLogsFiltersTFModel{
				SeverityLevels:  types.SetNull(types.StringType),
				ServiceNames:    types.SetNull(types.StringType),
				FilterGroupJSON: jsontypes.NewNormalizedValue(`{"type":"AND","values":[{"key":"status","type":"log_attribute","operator":"exact","value":["500"]}]}`),
			}},
			wantFilters: &PlatformAlertLogsFiltersTFModel{
				SeverityLevels:  types.SetNull(types.StringType),
				ServiceNames:    types.SetNull(types.StringType),
				FilterGroupJSON: jsontypes.NewNormalizedValue(`{"type":"AND","values":[{"key":"status","operator":"exact","type":"log_attribute","value":["500"]}]}`),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			logs, diags := logsSourceConfigToModel(context.Background(), json.RawMessage(test.raw), test.current)
			require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
			if test.wantFilters == nil {
				assert.Nil(t, logs.Filters)
				return
			}
			require.NotNil(t, logs.Filters)
			assert.True(t, test.wantFilters.SeverityLevels.Equal(logs.Filters.SeverityLevels), "severity_levels: %v", logs.Filters.SeverityLevels)
			assert.True(t, test.wantFilters.ServiceNames.Equal(logs.Filters.ServiceNames), "service_names: %v", logs.Filters.ServiceNames)
			if test.wantFilters.FilterGroupJSON.IsNull() {
				assert.True(t, logs.Filters.FilterGroupJSON.IsNull(), "filter_group_json: %s", logs.Filters.FilterGroupJSON.ValueString())
				return
			}
			equal, d := test.wantFilters.FilterGroupJSON.StringSemanticEquals(context.Background(), logs.Filters.FilterGroupJSON)
			require.False(t, d.HasError())
			assert.True(t, equal, "filter_group_json: %s", logs.Filters.FilterGroupJSON.ValueString())
		})
	}
}

// After a move the state holds the logs alert's id. Resolving it must never answer 404, because the
// generic read would drop the alert from state and plan a second one beside the logs alert.
func TestPlatformAlertRead_ResolvesAMovedAlert(t *testing.T) {
	tests := map[string]struct {
		results []httpclient.PlatformAlert
		wantErr string
	}{
		"platform-owned alert resolves to its platform id": {
			results: []httpclient.PlatformAlert{{ID: testPlatformAlertID, LegacyConfigurationID: util.StringPtr(testLogsAlertID), Owner: util.StringPtr("platform")}},
		},
		"a server that ignores the filter never adopts an unrelated alert": {
			results: []httpclient.PlatformAlert{{ID: testPlatformAlertID, LegacyConfigurationID: util.StringPtr("0190c000-0000-0000-0000-0000000000ff"), Owner: util.StringPtr("platform")}},
			wantErr: "has no platform alert",
		},
		"no platform copy fails the plan": {
			wantErr: "has no platform alert",
		},
		"source-owned alert fails the plan": {
			results: []httpclient.PlatformAlert{{ID: testPlatformAlertID, LegacyConfigurationID: util.StringPtr(testLogsAlertID), Owner: util.StringPtr("source")}},
			wantErr: "still evaluated by the logs product",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				require.Equal(t, "/api/projects/2/platform_alerts/", r.URL.Path)
				require.Equal(t, testLogsAlertID, r.URL.Query().Get("legacy_configuration_id"))
				w.Header().Set("Content-Type", "application/json")
				require.NoError(t, json.NewEncoder(w).Encode(map[string]any{"results": test.results}))
			}))
			defer server.Close()
			client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")

			model := PlatformAlertTFModel{}
			model.ID = types.StringValue(legacyIDPrefix + testLogsAlertID)
			model.ProjectID = types.StringValue("2")

			alert, status, err := PlatformAlertOps{}.Read(context.Background(), client, model)

			assert.NotEqual(t, http.StatusNotFound, int(status))
			if test.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, testPlatformAlertID, alert.ID)
		})
	}
}

func TestSnoozedUntil(t *testing.T) {
	now := time.Date(2026, 10, 21, 12, 0, 0, 0, time.UTC)
	tests := map[string]struct {
		snoozeUntil *string
		want        bool
	}{
		"not snoozed":       {snoozeUntil: nil},
		"snooze running":    {snoozeUntil: util.StringPtr("2026-10-21T13:00:00Z"), want: true},
		"snooze expired":    {snoozeUntil: util.StringPtr("2026-10-20T13:00:00.123456Z")},
		"unparseable value": {snoozeUntil: util.StringPtr("tomorrow"), want: true},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.want, snoozedUntil(test.snoozeUntil, now))
		})
	}
}

func TestValidatePlatformAlertPlan(t *testing.T) {
	tests := map[string]struct {
		mutate    func(*PlatformAlertTFModel)
		wantError string
	}{
		"valid": {},
		"logs kind without a logs block": {
			mutate:    func(m *PlatformAlertTFModel) { m.Logs = nil },
			wantError: "Missing logs settings",
		},
		"enabled without filters": {
			mutate:    func(m *PlatformAlertTFModel) { m.Logs.Filters = nil },
			wantError: "Log alert has no filters",
		},
		"disabled draft without filters": {
			mutate: func(m *PlatformAlertTFModel) {
				m.Logs.Filters = nil
				m.Enabled = types.BoolValue(false)
			},
		},
		"more datapoints than periods": {
			mutate:    func(m *PlatformAlertTFModel) { m.DatapointsToAlarm = types.Int64Value(3) },
			wantError: "Alert can never fire",
		},
		"below zero": {
			mutate: func(m *PlatformAlertTFModel) {
				m.Logs.Condition.ThresholdOperator = types.StringValue("below")
				m.Logs.Condition.ThresholdCount = types.Int64Value(0)
			},
			wantError: "Alert can never fire",
		},
		"grouped by service": {
			mutate: func(m *PlatformAlertTFModel) {
				m.Grouping.Mode = types.StringValue(groupingModeByLabels)
				m.Grouping.Keys = keyList("service_name")
			},
		},
		"single with keys": {
			mutate:    func(m *PlatformAlertTFModel) { m.Grouping.Keys = keyList("service_name") },
			wantError: "Grouping keys need by_result_labels",
		},
		"by labels without keys": {
			mutate:    func(m *PlatformAlertTFModel) { m.Grouping.Mode = types.StringValue(groupingModeByLabels) },
			wantError: "Grouping needs keys",
		},
		"unknown logs grouping key": {
			mutate: func(m *PlatformAlertTFModel) {
				m.Grouping.Mode = types.StringValue(groupingModeByLabels)
				m.Grouping.Keys = keyList("host")
			},
			wantError: "Unsupported grouping key",
		},
		"grouped below": {
			mutate: func(m *PlatformAlertTFModel) {
				m.Grouping.Mode = types.StringValue(groupingModeByLabels)
				m.Grouping.Keys = keyList("service_name")
				m.Logs.Condition.ThresholdOperator = types.StringValue("below")
			},
			wantError: "Grouped alert must fire above",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			plan := platformAlertPlan(t, severityFilter(t, "error"))
			if test.mutate != nil {
				test.mutate(&plan)
			}
			diags := validatePlatformAlertPlan(context.Background(), plan)
			if test.wantError == "" {
				assert.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
				return
			}
			require.True(t, diags.HasError())
			assert.Equal(t, test.wantError, diags.Errors()[0].Summary())
		})
	}
}

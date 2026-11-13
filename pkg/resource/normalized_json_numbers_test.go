package resource

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/posthog/terraform-provider/pkg/httpclient"
	"github.com/posthog/terraform-provider/pkg/resource/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The tests in this file cover every JSON attribute that PostHog reads back
// from its API. Configs often write whole numbers as floats: Pulumi Python's
// json.dumps writes 2.0, for example. The API client decodes responses into
// float64 and re-encodes them as 2. If state compared number text, the apply
// would fail with "Provider produced inconsistent result after apply".
//
// Each config is compact with sorted keys, which is how the provider encodes
// state, so the only text difference left is the number format.

// decodeAPIResponse decodes an API response body the way the API client does.
func decodeAPIResponse[T any](t *testing.T, body string) T {
	t.Helper()
	var v T
	require.NoError(t, json.Unmarshal([]byte(body), &v))
	return v
}

// assertStateMatchesConfig fails unless the state mapped from the API response
// holds the integer form of the config's number and still equals the config.
func assertStateMatchesConfig(t *testing.T, state core.NormalizedJSON, configJSON string) {
	t.Helper()
	require.False(t, state.IsNull(), "state must be populated from the API response")
	require.NotEqual(t, configJSON, state.ValueString(),
		"state must hold the API's integer form, otherwise this test compares the config with itself")

	equal, diags := core.NewNormalizedJSONValue(configJSON).StringSemanticEquals(context.Background(), state)
	require.False(t, diags.HasError(), diags.Errors())
	assert.True(t, equal, "config %s must equal state %s", configJSON, state.ValueString())
}

func TestCohortFilters_WholeNumberFloatMatchesAPIInteger(t *testing.T) {
	configJSON := `{"properties":{"type":"AND","values":[{"key":"$session_duration","operator":"gt","type":"person","value":60.0}]}}`
	model := CohortTFModel{Filters: core.NewNormalizedJSONValue(configJSON)}
	resp := decodeAPIResponse[httpclient.Cohort](t,
		`{"id":1,"filters":{"properties":{"type":"AND","values":[{"key":"$session_duration","operator":"gt","type":"person","value":60}]}}}`)

	diags := CohortOps{}.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), diags.Errors())

	assertStateMatchesConfig(t, model.Filters, configJSON)
}

func TestFeatureFlagFilters_WholeNumberFloatMatchesAPIInteger(t *testing.T) {
	configJSON := `{"groups":[{"properties":[],"rollout_percentage":50.0}]}`
	model := FeatureFlagTFModel{Filters: core.NewNormalizedJSONValue(configJSON)}
	resp := decodeAPIResponse[httpclient.FeatureFlag](t,
		`{"id":1,"key":"my_flag","filters":{"groups":[{"properties":[],"rollout_percentage":50}]}}`)

	diags := FeatureFlagOps{}.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), diags.Errors())

	assertStateMatchesConfig(t, model.Filters, configJSON)
}

func TestExperimentJSON_WholeNumberFloatMatchesAPIInteger(t *testing.T) {
	configJSON := `[{"conversion_window":14.0,"kind":"ExperimentMetric","metric_type":"funnel"}]`
	exposureConfigJSON := `{"filterTestAccounts":true,"multiple_variant_handling":"exclude","rollout_window":7.0}`
	model := ExperimentTFModel{
		Metrics:          core.NewNormalizedJSONValue(configJSON),
		MetricsSecondary: core.NewNormalizedJSONValue(configJSON),
		ExposureCriteria: core.NewNormalizedJSONValue(exposureConfigJSON),
	}
	resp := decodeAPIResponse[httpclient.Experiment](t, `{
		"id": 1,
		"metrics": [{"conversion_window": 14, "kind": "ExperimentMetric", "metric_type": "funnel"}],
		"metrics_secondary": [{"conversion_window": 14, "kind": "ExperimentMetric", "metric_type": "funnel"}],
		"exposure_criteria": {"filterTestAccounts": true, "multiple_variant_handling": "exclude", "rollout_window": 7}
	}`)

	diags := ExperimentOps{}.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), diags.Errors())

	t.Run("metrics", func(t *testing.T) { assertStateMatchesConfig(t, model.Metrics, configJSON) })
	t.Run("metrics_secondary", func(t *testing.T) { assertStateMatchesConfig(t, model.MetricsSecondary, configJSON) })
	t.Run("exposure_criteria", func(t *testing.T) { assertStateMatchesConfig(t, model.ExposureCriteria, exposureConfigJSON) })
}

func TestSubscriptionAIPromptConfig_WholeNumberFloatMatchesAPIInteger(t *testing.T) {
	configJSON := `{"window":{"end_days_ago":null,"mode":"last_n_days","start_days_ago":7.0}}`
	model := SubscriptionResourceTFModel{AIPromptConfig: core.NewNormalizedJSONValue(configJSON)}
	resp := decodeAPIResponse[httpclient.Subscription](t,
		`{"id":1,"start_date":"2026-08-17T07:00:00Z","ai_prompt_config":{"window":{"end_days_ago":null,"mode":"last_n_days","start_days_ago":7}}}`)

	diags := SubscriptionOps{}.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), diags.Errors())

	assertStateMatchesConfig(t, model.AIPromptConfig, configJSON)
}

// sensitive_inputs_json is not covered: PostHog never returns secret values, so
// the provider keeps the configured value in state instead of reading it back.
func TestHogFunctionJSON_WholeNumberFloatMatchesAPIInteger(t *testing.T) {
	inputsConfig := `{"timeout":{"value":30.0}}`
	inputsSchemaConfig := `[{"default":30.0,"key":"timeout","type":"number"}]`
	filtersConfig := `{"properties":[{"key":"$browser_version","operator":"gt","type":"event","value":100.0}]}`
	maskingConfig := `{"hash":"{person.id}","threshold":5.0,"ttl":3600.0}`
	mappingsConfig := `[{"inputs":{"timeout":{"value":10.0}},"name":"Purchase"}]`
	model := HogFunctionResourceTFModel{
		InputsJSON:          core.NewNormalizedJSONValue(inputsConfig),
		SensitiveInputsJSON: core.NewNormalizedJSONNull(),
		InputsSchemaJSON:    core.NewNormalizedJSONValue(inputsSchemaConfig),
		FiltersJSON:         core.NewNormalizedJSONValue(filtersConfig),
		MaskingJSON:         core.NewNormalizedJSONValue(maskingConfig),
		MappingsJSON:        core.NewNormalizedJSONValue(mappingsConfig),
	}
	resp := decodeAPIResponse[httpclient.HogFunction](t, `{
		"id": "0190d0c5-0000-0000-0000-000000000001",
		"inputs": {"timeout": {"value": 30}},
		"inputs_schema": [{"default": 30, "key": "timeout", "type": "number"}],
		"filters": {"properties": [{"key": "$browser_version", "operator": "gt", "type": "event", "value": 100}]},
		"masking": {"hash": "{person.id}", "threshold": 5, "ttl": 3600},
		"mappings": [{"inputs": {"timeout": {"value": 10}}, "name": "Purchase"}]
	}`)

	diags := HogFunctionOps{}.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), diags.Errors())

	t.Run("inputs_json", func(t *testing.T) { assertStateMatchesConfig(t, model.InputsJSON, inputsConfig) })
	t.Run("inputs_schema_json", func(t *testing.T) { assertStateMatchesConfig(t, model.InputsSchemaJSON, inputsSchemaConfig) })
	t.Run("filters_json", func(t *testing.T) { assertStateMatchesConfig(t, model.FiltersJSON, filtersConfig) })
	t.Run("masking_json", func(t *testing.T) { assertStateMatchesConfig(t, model.MaskingJSON, maskingConfig) })
	t.Run("mappings_json", func(t *testing.T) { assertStateMatchesConfig(t, model.MappingsJSON, mappingsConfig) })
}

func TestEarlyAccessFeaturePayload_WholeNumberFloatMatchesAPIInteger(t *testing.T) {
	configJSON := `{"max_items":10.0}`
	model := EarlyAccessFeatureTFModel{Payload: core.NewNormalizedJSONValue(configJSON)}
	resp := decodeAPIResponse[httpclient.EarlyAccessFeature](t, `{"id":"0190d0c5-0000-0000-0000-000000000002","payload":{"max_items":10}}`)

	diags := EarlyAccessFeatureOps{}.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), diags.Errors())

	assertStateMatchesConfig(t, model.Payload, configJSON)
}

func TestDashboardLayoutLayoutsJSON_WholeNumberFloatMatchesAPIInteger(t *testing.T) {
	configJSON := `{"sm":{"h":5.0,"w":6.0,"x":0.0,"y":0.0}}`
	tile := decodeAPIResponse[httpclient.DashboardTile](t,
		`{"id":1,"text":{"body":"Notes"},"layouts":{"sm":{"h":5,"w":6,"x":0,"y":0}}}`)

	assertStateMatchesConfig(t, apiTileToTFModel(tile).LayoutsJSON, configJSON)
}

func TestProjectSettingsTestAccountFilters_WholeNumberFloatMatchesAPIInteger(t *testing.T) {
	configJSON := `[{"key":"$session_duration","operator":"gt","type":"event","value":60.0}]`
	model := ProjectSettingsModel{TestAccountFilters: core.NewNormalizedJSONValue(configJSON)}
	resp := decodeAPIResponse[httpclient.Environment](t,
		`{"id":1,"test_account_filters":[{"key":"$session_duration","operator":"gt","type":"event","value":60}]}`)

	diags := ProjectSettingsOps{}.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), diags.Errors())

	assertStateMatchesConfig(t, model.TestAccountFilters, configJSON)
}

func TestLogsAlertFilterGroupJSON_WholeNumberFloatMatchesAPIInteger(t *testing.T) {
	configJSON := `{"type":"AND","values":[{"key":"status_code","operator":"gt","type":"log_attribute","value":499.0}]}`
	model := LogsAlertTFModel{FilterGroupJSON: core.NewNormalizedJSONValue(configJSON)}
	resp := decodeAPIResponse[httpclient.LogsAlert](t,
		`{"id":"019dbe94-cec8-781b-9470-4a970cd69ebf","filters":{"filterGroup":{"type":"AND","values":[{"key":"status_code","label":null,"operator":"gt","type":"log_attribute","value":499}]}}}`)

	diags := LogsAlertOps{}.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), diags.Errors())

	assertStateMatchesConfig(t, model.FilterGroupJSON, configJSON)
}

package resource

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFeatureFlagMapResponseToModel_NormalizesServerOnlyFilterFields(t *testing.T) {
	ops := FeatureFlagOps{}
	model := FeatureFlagTFModel{
		Filters: jsontypes.NewNormalizedValue(`{"groups":[{"rollout_percentage":0}]}`),
	}

	resp := httpclient.FeatureFlag{
		ID:  123,
		Key: "classic_questionnaires_enabled",
		Filters: map[string]interface{}{
			"aggregation_group_type_index": nil,
			"groups": []interface{}{
				map[string]interface{}{
					"aggregation_group_type_index": nil,
					"rollout_percentage":           float64(0),
				},
			},
		},
	}

	diags := ops.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), diags.Errors())
	assert.Equal(t, `{"groups":[{"rollout_percentage":0}]}`, model.Filters.ValueString())
	assert.Equal(t, int64(0), model.RolloutPercentage.ValueInt64())
}

// TestFeatureFlagMapResponseToModel_NoPerpetualDiffOnKeyOrder reproduces issue #111:
// PostHog's API returns filters JSON with alphabetically-sorted object keys, while the
// user's config uses a natural (non-alphabetical) key order. A byte-for-byte comparison
// of config vs. state therefore reports an endless diff. Using jsontypes.Normalized makes
// the framework compare the two semantically, so equivalent JSON no longer drifts.
func TestFeatureFlagMapResponseToModel_NoPerpetualDiffOnKeyOrder(t *testing.T) {
	ops := FeatureFlagOps{}

	// User config: property object keys in a natural (non-alphabetical) order.
	configJSON := `{"groups":[{"properties":[{"key":"email","type":"person","operator":"icontains","value":"@posthog.com"}]}]}`
	model := FeatureFlagTFModel{
		Filters: jsontypes.NewNormalizedValue(configJSON),
	}

	// API returns the same data but with alphabetically-sorted keys.
	resp := httpclient.FeatureFlag{
		ID:  1,
		Key: "my_flag",
		Filters: map[string]interface{}{
			"groups": []interface{}{
				map[string]interface{}{
					"properties": []interface{}{
						map[string]interface{}{
							"key":      "email",
							"operator": "icontains",
							"type":     "person",
							"value":    "@posthog.com",
						},
					},
				},
			},
		},
	}

	diags := ops.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), diags.Errors())

	// The stored state must be semantically equal to the user's config; otherwise
	// Terraform reports a perpetual diff even though nothing changed.
	eq, d := model.Filters.StringSemanticEquals(context.Background(), jsontypes.NewNormalizedValue(configJSON))
	require.False(t, d.HasError(), d.Errors())
	assert.True(t, eq, "filters state must be semantically equal to config to avoid a perpetual diff")
}

// Empty/unparseable prior state is treated as nothing configured (the initial-create
// path), so empty API defaults drop while non-empty values are kept.
func TestNormalizeFeatureFlagFiltersForState_UnparsedStateDropsEmptyDefaults(t *testing.T) {
	apiData := map[string]interface{}{
		"aggregation_group_type_index": nil,
		"payloads":                     map[string]interface{}{},
		"groups": []interface{}{
			map[string]interface{}{
				"aggregation_group_type_index": nil,
				"variant":                      nil,
				"properties": []interface{}{
					map[string]interface{}{
						"key":      "email",
						"operator": "exact",
						"type":     "person",
						"value":    []interface{}{"brad@example.com"},
					},
				},
				"rollout_percentage": float64(100),
			},
		},
	}

	for _, stateJSON := range []string{"", "not valid json"} {
		got, err := normalizeFeatureFlagFiltersForState(apiData, stateJSON, defaultIgnoredFilterKeys)
		require.NoError(t, err)
		assert.JSONEq(t, `{
			"groups": [{
				"properties": [{
					"key": "email",
					"operator": "exact",
					"type": "person",
					"value": ["brad@example.com"]
				}],
				"rollout_percentage": 100
			}]
		}`, got, "stateJSON=%q", stateJSON)
	}
}

// ffFilters with super_groups + holdout_groups (cross-product wiring, populated by Early
// Access Features / Experiments) alongside user-authored groups, multivariate and payloads.
func ffFiltersWithWiring() map[string]interface{} {
	return map[string]interface{}{
		"groups": []interface{}{
			map[string]interface{}{"rollout_percentage": float64(100)},
		},
		"multivariate": map[string]interface{}{
			"variants": []interface{}{map[string]interface{}{"key": "control", "rollout_percentage": float64(100)}},
		},
		"payloads":       map[string]interface{}{"control": `{"x":1}`},
		"super_groups":   []interface{}{map[string]interface{}{"rollout_percentage": float64(100)}},
		"holdout_groups": []interface{}{map[string]interface{}{"rollout_percentage": float64(10)}},
		"holdout":        map[string]interface{}{"id": float64(7)},
	}
}

// The default ignore set drops the cross-product wiring keys (super_groups, holdout_groups)
// that the user did not configure, but KEEPS multivariate and payloads — those are part of
// the flag's own definition and stay in the normal drift-detected flow.
func TestNormalizeFeatureFlagFiltersForState_DefaultIgnoresWiringKeepsMultivariateAndPayloads(t *testing.T) {
	got, err := normalizeFeatureFlagFiltersForState(ffFiltersWithWiring(), `{"groups":[{"rollout_percentage":100}]}`, defaultIgnoredFilterKeys)
	require.NoError(t, err)
	assert.JSONEq(t, `{
		"groups": [{"rollout_percentage": 100}],
		"multivariate": {"variants": [{"key": "control", "rollout_percentage": 100}]},
		"payloads": {"control": "{\"x\":1}"}
	}`, got)
}

// When no ignore key matches, every non-empty field survives — whether the set is empty
// (`[]` = track everything) or holds only near-misses/typos (no prefix or fuzzy matching).
func TestNormalizeFeatureFlagFiltersForState_KeepsAllWhenNoIgnoreKeyMatches(t *testing.T) {
	const wantAll = `{
		"groups": [{"rollout_percentage": 100}],
		"multivariate": {"variants": [{"key": "control", "rollout_percentage": 100}]},
		"payloads": {"control": "{\"x\":1}"},
		"super_groups": [{"rollout_percentage": 100}],
		"holdout_groups": [{"rollout_percentage": 10}],
		"holdout": {"id": 7}
	}`
	cases := []struct {
		name        string
		ignoredKeys []string
	}{
		{"explicit empty set tracks everything", []string{}},
		{"near-miss and typo keys are a no-op", []string{"super_group", "payload", "not_a_key"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := normalizeFeatureFlagFiltersForState(ffFiltersWithWiring(), `{"groups":[{"rollout_percentage":100}]}`, tc.ignoredKeys)
			require.NoError(t, err)
			assert.JSONEq(t, wantAll, got)
		})
	}
}

// Configured beats ignored: a key the user declares in filters is tracked even when it is
// in the ignore set, while the other default-ignored wiring keys are still dropped.
func TestNormalizeFeatureFlagFiltersForState_ConfiguredKeyBeatsIgnore(t *testing.T) {
	// super_groups is in the default ignore set, but the user declared it in prior state.
	got, err := normalizeFeatureFlagFiltersForState(ffFiltersWithWiring(),
		`{"groups":[{"rollout_percentage":100}],"super_groups":[{"rollout_percentage":100}]}`, defaultIgnoredFilterKeys)
	require.NoError(t, err)
	// super_groups kept (configured wins); holdout_groups + holdout dropped (ignored, not
	// configured); multivariate + payloads kept (never default-ignored).
	assert.JSONEq(t, `{
		"groups": [{"rollout_percentage": 100}],
		"super_groups": [{"rollout_percentage": 100}],
		"multivariate": {"variants": [{"key": "control", "rollout_percentage": 100}]},
		"payloads": {"control": "{\"x\":1}"}
	}`, got)
}

func TestResolveIgnoredFilterKeys(t *testing.T) {
	ctx := context.Background()

	// Unset → default set.
	keys, diags := resolveIgnoredFilterKeys(ctx, types.SetNull(types.StringType))
	require.False(t, diags.HasError(), diags.Errors())
	assert.Equal(t, defaultIgnoredFilterKeys, keys)

	// Explicit empty → empty (track everything), NOT the default.
	empty, _ := types.SetValue(types.StringType, []attr.Value{})
	keys, diags = resolveIgnoredFilterKeys(ctx, empty)
	require.False(t, diags.HasError(), diags.Errors())
	assert.Empty(t, keys)

	// Explicit set → that set.
	custom, _ := types.SetValue(types.StringType, []attr.Value{types.StringValue("payloads")})
	keys, diags = resolveIgnoredFilterKeys(ctx, custom)
	require.False(t, diags.HasError(), diags.Errors())
	assert.Equal(t, []string{"payloads"}, keys)
}

// An explicit ignore_filter_fields set is threaded through MapResponseToModel and replaces
// the default: here it ignores only super_groups, so holdout_groups (default-ignored) is
// now tracked. Exercises the full ops-layer wiring, not just the pure helpers.
func TestFeatureFlagMapResponseToModel_ExplicitIgnoreFilterFieldsReplacesDefault(t *testing.T) {
	ops := FeatureFlagOps{}
	ignore, d := types.SetValue(types.StringType, []attr.Value{types.StringValue("super_groups")})
	require.False(t, d.HasError(), d.Errors())
	model := FeatureFlagTFModel{
		Filters:            jsontypes.NewNormalizedValue(`{"groups":[{"rollout_percentage":100}]}`),
		IgnoreFilterFields: ignore,
	}

	resp := httpclient.FeatureFlag{ID: 1, Key: "my_flag", Filters: ffFiltersWithWiring()}
	diags := ops.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), diags.Errors())

	// super_groups dropped (in the explicit set); everything else — including holdout_groups
	// and holdout, which the default would have dropped — is kept because the set replaces it.
	assert.JSONEq(t, `{
		"groups": [{"rollout_percentage": 100}],
		"multivariate": {"variants": [{"key": "control", "rollout_percentage": 100}]},
		"payloads": {"control": "{\"x\":1}"},
		"holdout_groups": [{"rollout_percentage": 10}],
		"holdout": {"id": 7}
	}`, model.Filters.ValueString())
}

func TestFeatureFlagMapResponseToModel_PreservesRemoteFilterDrift(t *testing.T) {
	ops := FeatureFlagOps{}
	model := FeatureFlagTFModel{
		Filters: jsontypes.NewNormalizedValue(`{"groups":[{"rollout_percentage":100},{"properties":[{"key":"email","operator":"exact","type":"person","value":["admin@example.com"]}],"rollout_percentage":100}]}`),
	}

	resp := httpclient.FeatureFlag{
		ID:  123,
		Key: "create-additional-business",
		Filters: map[string]interface{}{
			"aggregation_group_type_index": nil,
			"payloads":                     map[string]interface{}{},
			"groups": []interface{}{
				map[string]interface{}{
					"aggregation_group_type_index": nil,
					"properties": []interface{}{
						map[string]interface{}{
							"key":      "email",
							"operator": "exact",
							"type":     "person",
							"value":    []interface{}{"brad@example.com"},
						},
					},
					"rollout_percentage": float64(100),
				},
				map[string]interface{}{
					"aggregation_group_type_index": nil,
					"variant":                      nil,
					"properties": []interface{}{
						map[string]interface{}{
							"key":      "email",
							"operator": "exact",
							"type":     "person",
							"value":    []interface{}{"admin@example.com"},
						},
					},
					"rollout_percentage": float64(100),
				},
			},
		},
	}

	diags := ops.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), diags.Errors())
	assert.JSONEq(t, `{
		"groups": [
			{
				"properties": [{
					"key": "email",
					"operator": "exact",
					"type": "person",
					"value": ["brad@example.com"]
				}],
				"rollout_percentage": 100
			},
			{
				"properties": [{
					"key": "email",
					"operator": "exact",
					"type": "person",
					"value": ["admin@example.com"]
				}],
				"rollout_percentage": 100
			}
		]
	}`, model.Filters.ValueString())
}

func TestFeatureFlagBuildCreateRequestSetsEnsureExperienceContinuity(t *testing.T) {
	ops := FeatureFlagOps{}
	model := FeatureFlagTFModel{
		Key:                        types.StringValue("my_flag"),
		EnsureExperienceContinuity: types.BoolValue(false),
	}

	req, diags := ops.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError(), diags.Errors())
	require.NotNil(t, req.EnsureExperienceContinuity)
	assert.False(t, *req.EnsureExperienceContinuity)
}

func TestFeatureFlagBuildCreateRequestOmitsEnsureExperienceContinuityWhenNull(t *testing.T) {
	ops := FeatureFlagOps{}
	model := FeatureFlagTFModel{
		Key: types.StringValue("my_flag"),
	}

	req, diags := ops.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError(), diags.Errors())
	assert.Nil(t, req.EnsureExperienceContinuity)
}

func TestFeatureFlagBuildCreateRequestSuppressesUsageDashboardByDefault(t *testing.T) {
	ops := FeatureFlagOps{}
	model := FeatureFlagTFModel{
		Key: types.StringValue("my_flag"),
	}

	req, diags := ops.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError(), diags.Errors())
	require.NotNil(t, req.ShouldCreateUsageDashboard)
	assert.False(t, *req.ShouldCreateUsageDashboard)
}

func TestFeatureFlagBuildCreateRequestAllowsUsageDashboardOptIn(t *testing.T) {
	ops := FeatureFlagOps{}
	model := FeatureFlagTFModel{
		Key:                  types.StringValue("my_flag"),
		CreateUsageDashboard: types.BoolValue(true),
	}

	req, diags := ops.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError(), diags.Errors())
	require.NotNil(t, req.ShouldCreateUsageDashboard)
	assert.True(t, *req.ShouldCreateUsageDashboard)
}

func TestFeatureFlagBuildUpdateRequestNeverSendsUsageDashboardField(t *testing.T) {
	ops := FeatureFlagOps{}
	plan := FeatureFlagTFModel{
		Key:                  types.StringValue("my_flag"),
		CreateUsageDashboard: types.BoolValue(true),
	}

	req, diags := ops.BuildUpdateRequest(context.Background(), plan, FeatureFlagTFModel{})
	require.False(t, diags.HasError(), diags.Errors())
	assert.Nil(t, req.ShouldCreateUsageDashboard)
}

func TestFeatureFlagBuildUpdateRequestSetsEnsureExperienceContinuity(t *testing.T) {
	ops := FeatureFlagOps{}
	plan := FeatureFlagTFModel{
		Key:                        types.StringValue("my_flag"),
		EnsureExperienceContinuity: types.BoolValue(true),
	}

	req, diags := ops.BuildUpdateRequest(context.Background(), plan, FeatureFlagTFModel{})
	require.False(t, diags.HasError(), diags.Errors())
	require.NotNil(t, req.EnsureExperienceContinuity)
	assert.True(t, *req.EnsureExperienceContinuity)
}

func TestFeatureFlagMapResponseToModel_EnsureExperienceContinuity(t *testing.T) {
	ops := FeatureFlagOps{}

	enabled := true
	respEnabled := httpclient.FeatureFlag{
		ID:                         1,
		Key:                        "my_flag",
		EnsureExperienceContinuity: &enabled,
	}
	var modelEnabled FeatureFlagTFModel
	diags := ops.MapResponseToModel(context.Background(), respEnabled, &modelEnabled)
	require.False(t, diags.HasError(), diags.Errors())
	assert.True(t, modelEnabled.EnsureExperienceContinuity.ValueBool())

	// A nil pointer (field absent from the API response) maps to a null value.
	respMissing := httpclient.FeatureFlag{ID: 2, Key: "my_flag"}
	var modelMissing FeatureFlagTFModel
	diags = ops.MapResponseToModel(context.Background(), respMissing, &modelMissing)
	require.False(t, diags.HasError(), diags.Errors())
	assert.True(t, modelMissing.EnsureExperienceContinuity.IsNull())
}

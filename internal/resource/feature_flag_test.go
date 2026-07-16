package resource

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
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

// TestNormalizeFeatureFlagFiltersForState_UnparsedStateDropsEmptyDefaults pins the
// documented fallback: when the prior state is empty or unparseable it is treated as
// "nothing configured", so every empty API default (null, {}, []) is dropped while
// non-empty values are kept. This is the initial-create path (state is null).
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
		got, err := normalizeFeatureFlagFiltersForState(apiData, stateJSON)
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

package resource

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFeatureFlagMapResponseToModel_NormalizesServerOnlyFilterFields(t *testing.T) {
	ops := FeatureFlagOps{}
	model := FeatureFlagTFModel{
		Filters: types.StringValue(`{"groups":[{"rollout_percentage":0}]}`),
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

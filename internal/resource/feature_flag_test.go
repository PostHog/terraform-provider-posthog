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

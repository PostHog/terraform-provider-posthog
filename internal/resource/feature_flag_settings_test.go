package resource

import (
	"context"
	"encoding/json"
	"maps"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	tfresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFeatureFlagSettingsRequests(t *testing.T) {
	for _, tc := range []struct {
		name  string
		model FeatureFlagTFModel
		want  map[string]interface{}
	}{
		{"omitted", FeatureFlagTFModel{}, nil},
		{"unknown", FeatureFlagTFModel{
			Archived: types.BoolUnknown(), BucketingIdentifier: types.StringUnknown(),
			EvaluationContexts: types.SetUnknown(types.StringType), IsRemoteConfiguration: types.BoolUnknown(),
		}, nil},
		{"explicit_false_and_empty", FeatureFlagTFModel{
			Archived: types.BoolValue(false), BucketingIdentifier: types.StringValue("distinct_id"),
			EvaluationContexts: types.SetValueMust(types.StringType, nil), IsRemoteConfiguration: types.BoolValue(false),
		}, map[string]interface{}{"archived": false, "bucketing_identifier": "distinct_id", "evaluation_contexts": []interface{}{}, "is_remote_configuration": false}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			create, diags := (FeatureFlagOps{}).BuildCreateRequest(context.Background(), tc.model)
			require.False(t, diags.HasError(), diags)
			update, diags := (FeatureFlagOps{}).BuildUpdateRequest(context.Background(), tc.model, tc.model)
			require.False(t, diags.HasError(), diags)
			for _, request := range []httpclient.FeatureFlagRequest{create, update} {
				raw, err := json.Marshal(request)
				require.NoError(t, err)
				var body map[string]interface{}
				require.NoError(t, json.Unmarshal(raw, &body))
				for _, field := range []string{"archived", "bucketing_identifier", "evaluation_contexts", "is_remote_configuration"} {
					if tc.want == nil {
						assert.NotContains(t, body, field)
					} else {
						assert.Equal(t, tc.want[field], body[field], field)
					}
				}
			}
		})
	}
}

func TestFeatureFlagRemoteConfigurationCreate(t *testing.T) {
	for _, tc := range []struct {
		name, filters string
		rollout       types.Int64
		wantError     bool
	}{
		{"payload_requires_group", `{"payloads":{"true":"{}"}}`, types.Int64Null(), true},
		{"payload_with_unknown_rollout", `{"payloads":{"true":"{}"}}`, types.Int64Unknown(), true},
		{"payload_with_rollout_override", `{"payloads":{"true":"{}"}}`, types.Int64Value(100), true},
		{"zero_group", `{"groups":[{"rollout_percentage":0}]}`, types.Int64Null(), true},
		{"overridden_zero_group", `{"groups":[{"rollout_percentage":0}]}`, types.Int64Value(100), true},
		{"conflicting_rollout", `{"groups":[{"rollout_percentage":50}]}`, types.Int64Value(100), true},
		{"matching_rollout", `{"groups":[{"rollout_percentage":100}]}`, types.Int64Value(100), false},
		{"second_zero_group", `{"groups":[{"rollout_percentage":100},{"rollout_percentage":0}]}`, types.Int64Value(100), true},
		{"explicit_group_with_unknown_convenience_field", `{"groups":[{"rollout_percentage":100}]}`, types.Int64Unknown(), false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			model := FeatureFlagTFModel{IsRemoteConfiguration: types.BoolValue(true),
				Filters: jsontypes.NewNormalizedValue(tc.filters), RolloutPercentage: tc.rollout}
			_, diags := (FeatureFlagOps{}).BuildCreateRequest(context.Background(), model)
			assert.Equal(t, tc.wantError, diags.HasError(), diags)
		})
	}
}

func TestFeatureFlagGrandfatheredDevicePersistence(t *testing.T) {
	ctx := context.Background()
	ops := FeatureFlagOps{}
	schema := ops.Schema()
	objectType := schema.Type().TerraformType(ctx).(tftypes.Object)
	for _, existing := range []bool{false, true} {
		t.Run(map[bool]string{false: "new_combination", true: "existing_combination"}[existing], func(t *testing.T) {
			prior := make(map[string]tftypes.Value)
			for name, typ := range objectType.AttributeTypes {
				prior[name] = tftypes.NewValue(typ, nil)
			}
			prior["key"] = tftypes.NewValue(tftypes.String, "grandfathered")
			prior["bucketing_identifier"] = tftypes.NewValue(tftypes.String, "device_id")
			prior["ensure_experience_continuity"] = tftypes.NewValue(tftypes.Bool, existing)
			planned := maps.Clone(prior)
			planned["name"] = tftypes.NewValue(tftypes.String, "unrelated update")
			planned["ensure_experience_continuity"] = tftypes.NewValue(tftypes.Bool, true)
			req := tfresource.ModifyPlanRequest{
				Config: tfsdk.Config{Schema: schema, Raw: tftypes.NewValue(objectType, planned)},
				State:  tfsdk.State{Schema: schema, Raw: tftypes.NewValue(objectType, prior)},
				Plan:   tfsdk.Plan{Schema: schema, Raw: tftypes.NewValue(objectType, planned)},
			}
			resp := tfresource.ModifyPlanResponse{Plan: req.Plan}
			ops.ModifyResourcePlan(ctx, req, &resp)
			assert.Equal(t, !existing, resp.Diagnostics.HasError(), resp.Diagnostics)
		})
	}
}

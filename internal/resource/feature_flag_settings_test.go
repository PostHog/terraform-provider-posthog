package resource

import (
	"context"
	"encoding/json"
	"maps"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	tfresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
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
			Archived: types.BoolUnknown(), BucketingIdentifier: types.StringUnknown(), EvaluationRuntime: types.StringUnknown(),
			EvaluationContexts: types.SetUnknown(types.StringType), IsRemoteConfiguration: types.BoolUnknown(),
		}, nil},
		{"explicit_false_and_empty", FeatureFlagTFModel{
			Archived: types.BoolValue(false), BucketingIdentifier: types.StringValue("distinct_id"), EvaluationRuntime: types.StringValue("server"),
			EvaluationContexts: types.SetValueMust(types.StringType, nil), IsRemoteConfiguration: types.BoolValue(false),
		}, map[string]interface{}{
			"archived": false, "bucketing_identifier": "distinct_id", "evaluation_runtime": "server",
			"evaluation_contexts": []interface{}{}, "is_remote_configuration": false,
		}},
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
				for _, field := range []string{"archived", "bucketing_identifier", "evaluation_runtime", "evaluation_contexts", "is_remote_configuration"} {
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

// modifyFeatureFlagPlan runs ModifyResourcePlan on an existing flag. Attributes missing from
// prior, planned or config are null.
func modifyFeatureFlagPlan(t *testing.T, prior, planned, config map[string]tftypes.Value) tfresource.ModifyPlanResponse {
	t.Helper()
	ctx := context.Background()
	ops := FeatureFlagOps{}
	schema := ops.Schema()
	objectType := schema.Type().TerraformType(ctx).(tftypes.Object)
	object := func(values map[string]tftypes.Value) tftypes.Value {
		all := make(map[string]tftypes.Value)
		for name, typ := range objectType.AttributeTypes {
			all[name] = tftypes.NewValue(typ, nil)
		}
		maps.Copy(all, values)
		return tftypes.NewValue(objectType, all)
	}
	req := tfresource.ModifyPlanRequest{
		Config: tfsdk.Config{Schema: schema, Raw: object(config)},
		State:  tfsdk.State{Schema: schema, Raw: object(prior)},
		Plan:   tfsdk.Plan{Schema: schema, Raw: object(planned)},
	}
	resp := tfresource.ModifyPlanResponse{Plan: req.Plan}
	ops.ModifyResourcePlan(ctx, req, &resp)
	return resp
}

func TestFeatureFlagModifyPlanOmittedSettings(t *testing.T) {
	prior := map[string]tftypes.Value{
		"key":                tftypes.NewValue(tftypes.String, "flag"),
		"evaluation_runtime": tftypes.NewValue(tftypes.String, "server"),
		"archived":           tftypes.NewValue(tftypes.Bool, false),
	}
	renamed := maps.Clone(prior)
	renamed["name"] = tftypes.NewValue(tftypes.String, "rename")
	config := map[string]tftypes.Value{
		"key":      tftypes.NewValue(tftypes.String, "flag"),
		"name":     tftypes.NewValue(tftypes.String, "rename"),
		"archived": tftypes.NewValue(tftypes.Bool, false),
	}

	t.Run("an update plans omitted settings unknown and keeps configured ones", func(t *testing.T) {
		resp := modifyFeatureFlagPlan(t, prior, renamed, config)
		require.False(t, resp.Diagnostics.HasError(), resp.Diagnostics)
		var plan FeatureFlagTFModel
		require.False(t, resp.Plan.Get(context.Background(), &plan).HasError())
		assert.True(t, plan.EvaluationRuntime.IsUnknown())
		assert.True(t, plan.EvaluationContexts.IsUnknown())
		assert.True(t, plan.Tags.IsUnknown())
		assert.Equal(t, types.BoolValue(false), plan.Archived)
	})

	t.Run("a plan with no changes keeps the state's values", func(t *testing.T) {
		resp := modifyFeatureFlagPlan(t, prior, prior, map[string]tftypes.Value{"key": prior["key"]})
		require.False(t, resp.Diagnostics.HasError(), resp.Diagnostics)
		var plan FeatureFlagTFModel
		require.False(t, resp.Plan.Get(context.Background(), &plan).HasError())
		assert.Equal(t, types.StringValue("server"), plan.EvaluationRuntime)
	})
}

func TestCheckEvaluationContextsApplied(t *testing.T) {
	contexts := func(names ...string) *[]string { return &names }
	for _, tc := range []struct {
		name      string
		requested *[]string
		applied   []string
		wantError bool
	}{
		{"not requested", nil, []string{"web"}, false},
		{"applied in another order", contexts("web", "api"), []string{"api", "web"}, false},
		{"cleared", contexts(), nil, false},
		{"ignored when setting", contexts("web"), nil, true},
		{"ignored when clearing", contexts(), []string{"web"}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := checkEvaluationContextsApplied(
				httpclient.FeatureFlagRequest{EvaluationContexts: tc.requested},
				httpclient.FeatureFlag{EvaluationContexts: tc.applied},
			)
			if tc.wantError {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "PostHog did not apply evaluation_contexts")
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestFeatureFlagCreateDeletesFlagWhenContextsIgnored(t *testing.T) {
	for _, tc := range []struct {
		name        string
		deleteFails bool
		wantError   []string
	}{
		{"deleted", false, []string{"PostHog did not apply evaluation_contexts", "deleted again"}},
		{"delete fails", true, []string{"PostHog did not apply evaluation_contexts", "serving without evaluation contexts", "Deleting it failed", "1/7"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var deleted bool
			// PostHog with evaluation contexts disabled: it creates the flag but stores no contexts.
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if r.Method == http.MethodPatch {
					if tc.deleteFails {
						w.WriteHeader(http.StatusInternalServerError)
						_, _ = w.Write([]byte(`{"detail":"server error"}`))
						return
					}
					var body map[string]interface{}
					_ = json.NewDecoder(r.Body).Decode(&body)
					deleted = body["deleted"] == true
				}
				_, _ = w.Write([]byte(`{"id":7,"key":"flag","evaluation_contexts":[]}`))
			}))
			defer server.Close()
			client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")
			contexts := []string{"web"}

			model := FeatureFlagTFModel{BaseProjectID: core.BaseProjectID{ProjectID: types.StringValue("1")}}
			_, err := FeatureFlagOps{}.Create(context.Background(), client, model, httpclient.FeatureFlagRequest{Key: "flag", EvaluationContexts: &contexts})

			require.Error(t, err)
			for _, want := range tc.wantError {
				assert.Contains(t, err.Error(), want)
			}
			assert.Equal(t, !tc.deleteFails, deleted)
		})
	}
}

func TestFeatureFlagMapResponseToModel_MissingContextsAreEmpty(t *testing.T) {
	var model FeatureFlagTFModel
	diags := FeatureFlagOps{}.MapResponseToModel(context.Background(), httpclient.FeatureFlag{ID: 1, Key: "flag"}, &model)
	require.False(t, diags.HasError(), diags)
	assert.Equal(t, types.SetValueMust(types.StringType, []attr.Value{}), model.EvaluationContexts)
}

func TestFeatureFlagUpdateSendsChangedContextsFirst(t *testing.T) {
	for _, tc := range []struct {
		name        string
		stored      []string // the flag's contexts in state and in PostHog
		requested   *[]string
		gateOn      bool // whether PostHog applies evaluation contexts
		disable     bool // the update sets active = false instead of true
		wantPatches int
		wantError   bool
	}{
		{name: "contexts ignored: nothing else is sent", stored: []string{}, requested: &[]string{"web"}, gateOn: false, wantPatches: 1, wantError: true},
		{name: "contexts applied: then the full update", stored: []string{}, requested: &[]string{"web"}, gateOn: true, wantPatches: 2},
		{name: "contexts unchanged: one update", stored: []string{"web"}, requested: &[]string{"web"}, wantPatches: 1},
		{name: "contexts not configured: one update", stored: []string{"web"}, requested: nil, wantPatches: 1},
		{name: "disabling while clearing contexts: disabled in the first request", stored: []string{"web"}, requested: &[]string{}, gateOn: true, disable: true, wantPatches: 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			stored := tc.stored
			var patches []map[string]interface{}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var body map[string]interface{}
				_ = json.NewDecoder(r.Body).Decode(&body)
				patches = append(patches, body)
				if contexts, ok := body["evaluation_contexts"].([]interface{}); ok && tc.gateOn {
					stored = nil
					for _, c := range contexts {
						stored = append(stored, c.(string))
					}
				}
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(httpclient.FeatureFlag{ID: 7, Key: "flag", EvaluationContexts: stored})
			}))
			defer server.Close()
			client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")

			contexts, diags := types.SetValueFrom(context.Background(), types.StringType, tc.stored)
			require.False(t, diags.HasError(), diags)
			state := FeatureFlagTFModel{
				BaseProjectID:      core.BaseProjectID{ProjectID: types.StringValue("1")},
				Key:                types.StringValue("flag"),
				EvaluationContexts: contexts,
			}
			state.ID = types.Int64Value(7)
			active := !tc.disable
			_, _, err := FeatureFlagOps{}.Update(context.Background(), client, state, httpclient.FeatureFlagRequest{
				Key: "flag", Active: &active, EvaluationContexts: tc.requested,
			})

			assert.Equal(t, tc.wantError, err != nil, err)
			require.Len(t, patches, tc.wantPatches)
			if tc.wantPatches == 2 || tc.wantError {
				if tc.disable {
					assert.Equal(t, false, patches[0]["active"], "a deactivation goes with the contexts")
				} else {
					assert.NotContains(t, patches[0], "active", "an activation waits for the contexts")
				}
			}
			if !tc.wantError {
				assert.Equal(t, active, patches[len(patches)-1]["active"])
			}
		})
	}
}

func TestFeatureFlagMapResponseToModel_Tags(t *testing.T) {
	for _, tc := range []struct {
		name  string
		prior types.Set
		api   []string
		want  types.Set
	}{
		{name: "tags from PostHog", prior: types.SetNull(types.StringType), api: []string{"ui"}, want: types.SetValueMust(types.StringType, []attr.Value{types.StringValue("ui")})},
		{name: "cleared with tags = []", prior: types.SetValueMust(types.StringType, []attr.Value{}), api: nil, want: types.SetValueMust(types.StringType, []attr.Value{})},
		{name: "never set", prior: types.SetNull(types.StringType), api: nil, want: types.SetNull(types.StringType)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			model := FeatureFlagTFModel{Tags: tc.prior}
			diags := FeatureFlagOps{}.MapResponseToModel(context.Background(), httpclient.FeatureFlag{ID: 1, Key: "flag", Tags: tc.api}, &model)
			require.False(t, diags.HasError(), diags)
			assert.Equal(t, tc.want, model.Tags)
		})
	}
}

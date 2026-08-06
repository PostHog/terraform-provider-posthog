package resource

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func stoppedStatus(state string, stopped *ExperimentStoppedModel) *ExperimentStatusModel {
	return &ExperimentStatusModel{State: types.StringValue(state), Stopped: stopped}
}

// --- BuildCreateRequest -----------------------------------------------------

func TestExperimentBuildCreateRequest_VariantsAndMetrics(t *testing.T) {
	ops := ExperimentOps{}
	model := ExperimentTFModel{
		Name:           types.StringValue("Pricing test"),
		Description:    types.StringValue("Does the redesign convert?"),
		FeatureFlagKey: types.StringValue("pricing-test"),
		Variant: []ExperimentVariantModel{
			{Key: types.StringValue("control"), Name: types.StringValue("Original"), RolloutPercentage: types.Int64Value(50)},
			{Key: types.StringValue("test"), Name: types.StringValue("Redesign"), RolloutPercentage: types.Int64Value(50)},
		},
		Metrics:            jsontypes.NewNormalizedValue(`[{"kind":"ExperimentFunnelsQuery","name":"Signup"}]`),
		HoldoutID:          types.Int64Value(42),
		AllowUnknownEvents: types.BoolValue(true),
		Status:             stoppedStatus(stateDraft, nil),
	}

	req, diags := ops.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError(), diags.Errors())

	assert.Equal(t, "Pricing test", *req.body.Name)
	assert.Equal(t, "Does the redesign convert?", *req.body.Description)
	assert.Equal(t, "pricing-test", *req.body.FeatureFlagKey)
	require.NotNil(t, req.body.HoldoutID)
	assert.Equal(t, int64(42), *req.body.HoldoutID)
	require.NotNil(t, req.body.AllowUnknownEvents)
	assert.True(t, *req.body.AllowUnknownEvents)

	// metrics pass through verbatim as raw JSON
	assert.JSONEq(t, `[{"kind":"ExperimentFunnelsQuery","name":"Signup"}]`, string(req.body.Metrics))

	// variants serialize into feature_flag.filters.multivariate.variants
	var ff map[string]interface{}
	require.NoError(t, json.Unmarshal(req.body.FeatureFlag, &ff))
	variants := ff["filters"].(map[string]interface{})["multivariate"].(map[string]interface{})["variants"].([]interface{})
	require.Len(t, variants, 2)
	first := variants[0].(map[string]interface{})
	assert.Equal(t, "control", first["key"])
	assert.Equal(t, "Original", first["name"])
	assert.Equal(t, float64(50), first["rollout_percentage"])
}

func TestExperimentBuildCreateRequest_NoVariants_OmitsFeatureFlag(t *testing.T) {
	ops := ExperimentOps{}
	model := ExperimentTFModel{
		Name:           types.StringValue("x"),
		FeatureFlagKey: types.StringValue("x-flag"),
		Metrics:        jsontypes.NewNormalizedNull(),
		Status:         stoppedStatus(stateDraft, nil),
	}
	req, diags := ops.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError())
	assert.Nil(t, req.body.FeatureFlag)
	assert.Nil(t, req.body.Metrics)
}

// --- BuildUpdateRequest: status transition dispatch table -------------------

func TestExperimentBuildUpdateRequest_TransitionTable(t *testing.T) {
	ops := ExperimentOps{}

	base := func(status *ExperimentStatusModel) ExperimentTFModel {
		return ExperimentTFModel{
			Name:           types.StringValue("exp"),
			FeatureFlagKey: types.StringValue("exp-flag"),
			Status:         status,
		}
	}

	tests := []struct {
		name        string
		fromState   *ExperimentStatusModel
		toState     *ExperimentStatusModel
		wantActions []string
		wantErr     bool
		wantShipKey string
		wantRelease bool
		wantConcl   string
	}{
		{name: "draft->running", fromState: stoppedStatus(stateDraft, nil), toState: stoppedStatus(stateRunning, nil), wantActions: []string{"launch"}},
		{name: "draft->paused", fromState: stoppedStatus(stateDraft, nil), toState: stoppedStatus(statePaused, nil), wantActions: []string{"launch", "pause"}},
		{name: "draft->stopped(end)", fromState: stoppedStatus(stateDraft, nil), toState: stoppedStatus(stateStopped, nil), wantActions: []string{"launch", "end"}},
		{
			name:      "draft->stopped(ship)",
			fromState: stoppedStatus(stateDraft, nil),
			toState: stoppedStatus(stateStopped, &ExperimentStoppedModel{
				ShipVariant:       types.StringValue("test"),
				ReleaseToEveryone: types.BoolValue(true),
				Conclusion:        types.StringValue("won"),
			}),
			wantActions: []string{"launch", "ship"}, wantShipKey: "test", wantRelease: true, wantConcl: "won",
		},
		{name: "running->paused", fromState: stoppedStatus(stateRunning, nil), toState: stoppedStatus(statePaused, nil), wantActions: []string{"pause"}},
		{name: "paused->running", fromState: stoppedStatus(statePaused, nil), toState: stoppedStatus(stateRunning, nil), wantActions: []string{"resume"}},
		{
			name:      "running->stopped(end w/ conclusion)",
			fromState: stoppedStatus(stateRunning, nil),
			toState: stoppedStatus(stateStopped, &ExperimentStoppedModel{
				Conclusion: types.StringValue("inconclusive"),
			}),
			wantActions: []string{"end"}, wantConcl: "inconclusive",
		},
		{
			name:      "paused->stopped(ship)",
			fromState: stoppedStatus(statePaused, nil),
			toState: stoppedStatus(stateStopped, &ExperimentStoppedModel{
				ShipVariant: types.StringValue("winner"),
			}),
			wantActions: []string{"ship"}, wantShipKey: "winner",
		},
		{name: "running->running(noop)", fromState: stoppedStatus(stateRunning, nil), toState: stoppedStatus(stateRunning, nil), wantActions: nil},
		{name: "stopped->stopped(no ship change)", fromState: stoppedStatus(stateStopped, nil), toState: stoppedStatus(stateStopped, nil), wantActions: nil},
		// illegal backward transitions
		{name: "running->draft(illegal)", fromState: stoppedStatus(stateRunning, nil), toState: stoppedStatus(stateDraft, nil), wantErr: true},
		{name: "paused->draft(illegal)", fromState: stoppedStatus(statePaused, nil), toState: stoppedStatus(stateDraft, nil), wantErr: true},
		{name: "stopped->running(illegal)", fromState: stoppedStatus(stateStopped, nil), toState: stoppedStatus(stateRunning, nil), wantErr: true},
		{name: "stopped->paused(illegal)", fromState: stoppedStatus(stateStopped, nil), toState: stoppedStatus(statePaused, nil), wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, diags := ops.BuildUpdateRequest(context.Background(), base(tc.toState), base(tc.fromState))
			if tc.wantErr {
				assert.True(t, diags.HasError(), "expected an error diagnostic")
				return
			}
			require.False(t, diags.HasError(), diags.Errors())
			assert.Equal(t, tc.wantActions, req.transition.actions)
			if tc.wantShipKey != "" {
				assert.Equal(t, tc.wantShipKey, req.transition.shipVariant)
				assert.Equal(t, tc.wantRelease, req.transition.releaseToEveryone)
			}
			if tc.wantConcl != "" {
				require.NotNil(t, req.transition.conclusion)
				assert.Equal(t, tc.wantConcl, *req.transition.conclusion)
			}
		})
	}
}

// ship-later: already stopped, ship_variant newly set relative to prior state -> ship.
func TestExperimentBuildUpdateRequest_ShipLater(t *testing.T) {
	ops := ExperimentOps{}

	state := ExperimentTFModel{
		Name:           types.StringValue("exp"),
		FeatureFlagKey: types.StringValue("exp-flag"),
		Status:         stoppedStatus(stateStopped, nil), // no ship_variant recorded yet
	}
	plan := ExperimentTFModel{
		Name:           types.StringValue("exp"),
		FeatureFlagKey: types.StringValue("exp-flag"),
		Status: stoppedStatus(stateStopped, &ExperimentStoppedModel{
			ShipVariant: types.StringValue("test"),
		}),
	}

	req, diags := ops.BuildUpdateRequest(context.Background(), plan, state)
	require.False(t, diags.HasError(), diags.Errors())
	assert.Equal(t, []string{"ship"}, req.transition.actions)
	assert.Equal(t, "test", req.transition.shipVariant)

	// ship idempotency: same ship_variant already recorded -> no action.
	statePrior := ExperimentTFModel{
		Name:           types.StringValue("exp"),
		FeatureFlagKey: types.StringValue("exp-flag"),
		Status: stoppedStatus(stateStopped, &ExperimentStoppedModel{
			ShipVariant: types.StringValue("test"),
		}),
	}
	req2, diags2 := ops.BuildUpdateRequest(context.Background(), plan, statePrior)
	require.False(t, diags2.HasError(), diags2.Errors())
	assert.Empty(t, req2.transition.actions)
}

// variant edits are only sent when changed, and carry update_feature_flag_params opt-in.
func TestExperimentBuildUpdateRequest_VariantEdits(t *testing.T) {
	ops := ExperimentOps{}
	variants := func(pct int64) []ExperimentVariantModel {
		return []ExperimentVariantModel{
			{Key: types.StringValue("control"), RolloutPercentage: types.Int64Value(100 - pct)},
			{Key: types.StringValue("test"), RolloutPercentage: types.Int64Value(pct)},
		}
	}

	unchanged := ExperimentTFModel{Name: types.StringValue("e"), FeatureFlagKey: types.StringValue("f"), Variant: variants(50), Status: stoppedStatus(stateRunning, nil)}
	reqSame, _ := ops.BuildUpdateRequest(context.Background(), unchanged, unchanged)
	assert.Nil(t, reqSame.body.FeatureFlag, "unchanged variants should not resend flag config")

	planChanged := ExperimentTFModel{Name: types.StringValue("e"), FeatureFlagKey: types.StringValue("f"), Variant: variants(60), UpdateFeatureFlagParams: types.BoolValue(true), Status: stoppedStatus(stateRunning, nil)}
	reqChanged, _ := ops.BuildUpdateRequest(context.Background(), planChanged, unchanged)
	require.NotNil(t, reqChanged.body.FeatureFlag)
	require.NotNil(t, reqChanged.body.UpdateFeatureFlagParams)
	assert.True(t, *reqChanged.body.UpdateFeatureFlagParams)
}

// --- MapResponseToModel -----------------------------------------------------

// Normalization: reordered + extra API keys in metrics must not produce a diff.
func TestExperimentMapResponseToModel_MetricsNoPerpetualDiff(t *testing.T) {
	ops := ExperimentOps{}
	configJSON := `[{"kind":"ExperimentFunnelsQuery","name":"Signup"}]`
	model := ExperimentTFModel{Metrics: jsontypes.NewNormalizedValue(configJSON)}

	// API returns reordered keys plus a server-computed fingerprint field.
	resp := httpclient.Experiment{
		ID:      7,
		Name:    "exp",
		Metrics: json.RawMessage(`[{"name":"Signup","fingerprint":"abc123","kind":"ExperimentFunnelsQuery"}]`),
		Status:  stateRunning,
	}

	diags := ops.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), diags.Errors())

	eq, d := model.Metrics.StringSemanticEquals(context.Background(), jsontypes.NewNormalizedValue(configJSON))
	require.False(t, d.HasError(), d.Errors())
	assert.True(t, eq, "metrics state should semantically equal config; got %s", model.Metrics.ValueString())
}

// Config-only: ship_variant / release_to_everyone / conclusion_comment must not be
// clobbered from the API; conclusion is refreshed from the API.
func TestExperimentMapResponseToModel_ConfigOnlyShipFields(t *testing.T) {
	ops := ExperimentOps{}
	model := ExperimentTFModel{
		Status: stoppedStatus(stateStopped, &ExperimentStoppedModel{
			ShipVariant:       types.StringValue("test"),
			ReleaseToEveryone: types.BoolValue(true),
			Conclusion:        types.StringValue("won"),
			ConclusionComment: types.StringValue("configured comment"),
		}),
	}

	conclusion := "won"
	respComment := "server comment that must be ignored"
	resp := httpclient.Experiment{
		ID:                9,
		Name:              "exp",
		Status:            stateStopped,
		Conclusion:        &conclusion,
		ConclusionComment: &respComment,
	}

	diags := ops.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), diags.Errors())

	assert.Equal(t, "test", model.Status.Stopped.ShipVariant.ValueString(), "ship_variant is config-only")
	assert.True(t, model.Status.Stopped.ReleaseToEveryone.ValueBool(), "release_to_everyone is config-only")
	assert.Equal(t, "configured comment", model.Status.Stopped.ConclusionComment.ValueString(), "conclusion_comment is config-only")
	assert.Equal(t, "won", model.Status.Stopped.Conclusion.ValueString(), "conclusion is readable")
	assert.Equal(t, stateStopped, model.Status.State.ValueString())
}

// Variants are read back from parameters.feature_flag_variants.
func TestExperimentMapResponseToModel_VariantsFromParameters(t *testing.T) {
	ops := ExperimentOps{}
	model := ExperimentTFModel{}

	resp := httpclient.Experiment{
		ID:   3,
		Name: "exp",
		Parameters: json.RawMessage(`{"feature_flag_variants":[` +
			`{"key":"control","name":"Original","rollout_percentage":50},` +
			`{"key":"test","name":"Redesign","rollout_percentage":50}]}`),
		Status: stateRunning,
	}

	diags := ops.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), diags.Errors())

	require.Len(t, model.Variant, 2)
	assert.Equal(t, "control", model.Variant[0].Key.ValueString())
	assert.Equal(t, "Original", model.Variant[0].Name.ValueString())
	assert.Equal(t, int64(50), model.Variant[0].RolloutPercentage.ValueInt64())
	assert.Equal(t, "test", model.Variant[1].Key.ValueString())
}

// Import starts from an empty model (nil status block): status must be populated without panic.
func TestExperimentMapResponseToModel_ImportPopulatesStatus(t *testing.T) {
	ops := ExperimentOps{}
	model := ExperimentTFModel{} // empty, as after the import parser

	resp := httpclient.Experiment{ID: 11, Name: "exp", FeatureFlagKey: "exp-flag", Status: statePaused}
	diags := ops.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError(), diags.Errors())

	require.NotNil(t, model.Status)
	assert.Equal(t, statePaused, model.Status.State.ValueString())
	assert.Nil(t, model.Status.Stopped)
	assert.Equal(t, "exp-flag", model.FeatureFlagKey.ValueString())
}

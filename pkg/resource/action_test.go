package resource

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/pkg/httpclient"
	"github.com/posthog/terraform-provider/pkg/resource/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestActionStepsJSON_ComparedByContent(t *testing.T) {
	assert.Equal(t, core.NormalizedJSONType{}, ActionOps{}.Schema().Attributes["steps_json"].GetType())
}

// PostHog returns no steps both for an action that never had any and for one whose
// steps were cleared, so state depends on what steps_json was before the response.
func TestActionMapResponseToModel_NoSteps(t *testing.T) {
	tests := map[string]struct {
		before, after core.NormalizedJSON
	}{
		"import":                {core.NewNormalizedJSONNull(), core.NewNormalizedJSONNull()},
		"omitted on create":     {core.NewNormalizedJSONUnknown(), core.NewNormalizedJSONValue("[]")},
		"configured as empty":   {core.NewNormalizedJSONValue("[]"), core.NewNormalizedJSONValue("[]")},
		"steps gone in PostHog": {core.NewNormalizedJSONValue(`[{"event":"a"}]`), core.NewNormalizedJSONValue("[]")},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			model := ActionTFModel{StepsJSON: tt.before}
			diags := ActionOps{}.MapResponseToModel(context.Background(), httpclient.Action{ID: 1}, &model)
			require.False(t, diags.HasError(), diags.Errors())
			assert.Equal(t, tt.after, model.StepsJSON)
		})
	}
}

func TestActionBuildUpdateRequest_Steps(t *testing.T) {
	state := ActionTFModel{Name: types.StringValue("a"), StepsJSON: core.NewNormalizedJSONValue(`[{"event":"a"}]`)}
	tests := map[string]struct {
		plan core.NormalizedJSON
		want []map[string]interface{}
	}{
		"formatted JSON": {core.NewNormalizedJSONValue("[\n  { \"event\": \"b\" }\n]\n"), []map[string]interface{}{{"event": "b"}}},
		"empty array":    {core.NewNormalizedJSONValue("[ ]"), []map[string]interface{}{}},
		"null":           {core.NewNormalizedJSONNull(), []map[string]interface{}{}},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			plan := state
			plan.StepsJSON = tt.plan
			req, diags := ActionOps{}.BuildUpdateRequest(context.Background(), plan, state)
			require.False(t, diags.HasError(), diags.Errors())
			require.NotNil(t, req.Steps)
			assert.Equal(t, tt.want, *req.Steps)
		})
	}
}

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

func TestEarlyAccessFeatureBuildCreateRequest_LinksConfiguredFlag(t *testing.T) {
	model := EarlyAccessFeatureTFModel{
		Name:             types.StringValue("New editor"),
		Stage:            types.StringValue("beta"),
		Description:      types.StringValue("A faster editor"),
		DocumentationURL: types.StringValue("https://example.com/docs"),
		Payload:          jsontypes.NewNormalizedValue(`{"theme":"dark"}`),
		FeatureFlagID:    types.Int64Value(42),
	}

	req, diags := EarlyAccessFeatureOps{}.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError(), diags.Errors())

	body, err := json.Marshal(req)
	require.NoError(t, err)
	assert.JSONEq(t, `{
		"name": "New editor",
		"stage": "beta",
		"description": "A faster editor",
		"documentation_url": "https://example.com/docs",
		"payload": {"theme": "dark"},
		"feature_flag_id": 42
	}`, string(body))
}

// feature_flag_id is Optional+Computed, so it is unknown in the plan when omitted. It must
// be left out of the request so PostHog creates the flag itself.
func TestEarlyAccessFeatureBuildCreateRequest_OmitsUnsetFields(t *testing.T) {
	model := EarlyAccessFeatureTFModel{
		Name:             types.StringValue("New editor"),
		Stage:            types.StringValue("draft"),
		Description:      types.StringNull(),
		DocumentationURL: types.StringNull(),
		Payload:          jsontypes.NewNormalizedNull(),
		FeatureFlagID:    types.Int64Unknown(),
	}

	req, diags := EarlyAccessFeatureOps{}.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError(), diags.Errors())

	body, err := json.Marshal(req)
	require.NoError(t, err)
	assert.JSONEq(t, `{"name": "New editor", "stage": "draft"}`, string(body))
}

func TestEarlyAccessFeatureBuildUpdateRequest_ClearsRemovedFieldsAndNeverRelinks(t *testing.T) {
	state := EarlyAccessFeatureTFModel{
		Name:             types.StringValue("New editor"),
		Stage:            types.StringValue("beta"),
		Description:      types.StringValue("A faster editor"),
		DocumentationURL: types.StringValue("https://example.com/docs"),
		Payload:          jsontypes.NewNormalizedValue(`{"theme":"dark"}`),
		FeatureFlagID:    types.Int64Value(42),
	}
	plan := state
	plan.Stage = types.StringValue("general-availability")
	plan.Description = types.StringNull()
	plan.DocumentationURL = types.StringNull()
	plan.Payload = jsontypes.NewNormalizedNull()

	req, diags := EarlyAccessFeatureOps{}.BuildUpdateRequest(context.Background(), plan, state)
	require.False(t, diags.HasError(), diags.Errors())

	body, err := json.Marshal(req)
	require.NoError(t, err)
	assert.JSONEq(t, `{
		"name": "New editor",
		"stage": "general-availability",
		"description": "",
		"documentation_url": "",
		"payload": {}
	}`, string(body))
}

func TestEarlyAccessFeatureMapResponseToModel(t *testing.T) {
	description := "A faster editor"
	empty := ""

	tests := map[string]struct {
		resp  httpclient.EarlyAccessFeature
		model EarlyAccessFeatureTFModel
		check func(t *testing.T, m EarlyAccessFeatureTFModel)
	}{
		"maps the linked flag and set fields": {
			resp: httpclient.EarlyAccessFeature{
				ID:          "eaf-1",
				FeatureFlag: &httpclient.EarlyAccessFeatureFeatureFlag{ID: 42, Key: "new-editor"},
				Name:        "New editor",
				Description: &description,
				Stage:       "beta",
				Payload:     json.RawMessage(`{"theme": "dark"}`),
				CreatedAt:   "2026-10-01T10:00:00Z",
			},
			check: func(t *testing.T, m EarlyAccessFeatureTFModel) {
				assert.Equal(t, "eaf-1", m.ID.ValueString())
				assert.Equal(t, int64(42), m.FeatureFlagID.ValueInt64())
				assert.Equal(t, "new-editor", m.FeatureFlagKey.ValueString())
				assert.Equal(t, "A faster editor", m.Description.ValueString())
				// Stored compact, as jsonencode() writes it, not with the API's spacing.
				assert.Equal(t, `{"theme":"dark"}`, m.Payload.ValueString())
				assert.Equal(t, "2026-10-01T10:00:00Z", m.CreatedAt.ValueString())
			},
		},
		"API empty defaults read as unset": {
			resp: httpclient.EarlyAccessFeature{
				ID: "eaf-1", Name: "New editor", Stage: "draft",
				Description: &empty, DocumentationURL: &empty, Payload: json.RawMessage(`{}`),
			},
			model: EarlyAccessFeatureTFModel{Payload: jsontypes.NewNormalizedNull()},
			check: func(t *testing.T, m EarlyAccessFeatureTFModel) {
				assert.True(t, m.Description.IsNull())
				assert.True(t, m.DocumentationURL.IsNull())
				assert.True(t, m.Payload.IsNull())
				assert.True(t, m.FeatureFlagID.IsNull())
				assert.True(t, m.FeatureFlagKey.IsNull())
			},
		},
		"an explicitly configured empty payload is kept": {
			resp:  httpclient.EarlyAccessFeature{ID: "eaf-1", Name: "New editor", Stage: "draft", Payload: json.RawMessage(`{}`)},
			model: EarlyAccessFeatureTFModel{Payload: jsontypes.NewNormalizedValue(`{}`)},
			check: func(t *testing.T, m EarlyAccessFeatureTFModel) {
				assert.Equal(t, `{}`, m.Payload.ValueString())
			},
		},
		"a payload changed outside Terraform surfaces as drift": {
			resp:  httpclient.EarlyAccessFeature{ID: "eaf-1", Name: "New editor", Stage: "draft", Payload: json.RawMessage(`{"theme":"dark","extra":1}`)},
			model: EarlyAccessFeatureTFModel{Payload: jsontypes.NewNormalizedValue(`{"theme":"dark"}`)},
			check: func(t *testing.T, m EarlyAccessFeatureTFModel) {
				assert.JSONEq(t, `{"theme":"dark","extra":1}`, m.Payload.ValueString())
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			model := tt.model
			diags := EarlyAccessFeatureOps{}.MapResponseToModel(context.Background(), tt.resp, &model)
			require.False(t, diags.HasError(), diags.Errors())
			tt.check(t, model)
		})
	}
}

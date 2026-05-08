package resource

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
	"github.com/posthog/terraform-provider/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testEDSProjectID  = "proj-1"
	testEDSSourceID   = "01900000-0000-7000-8000-000000000000"
	testEDSAPIKey     = "test-key"
	testEDSClient     = "test"
	testEDSStripeType = "Stripe"
	testEDSPrefix     = "stripe_"
)

func TestExternalDataSourceResourceNameAndSchema(t *testing.T) {
	ops := ExternalDataSourceOps{}
	s := ops.Schema()

	assert.Equal(t, "external_data_source", ops.ResourceName())
	assert.Contains(t, s.MarkdownDescription, "immutable")

	// source_type, prefix, and schemas must all be RequiresReplace — the API
	// rejects or silently ignores changes to any of them.
	for _, name := range []string{"source_type", "prefix"} {
		attr, ok := s.Attributes[name].(schema.StringAttribute)
		require.Truef(t, ok, "%s must be a string attribute", name)
		require.Lenf(t, attr.PlanModifiers, 1, "%s should carry exactly one plan modifier (RequiresReplace)", name)
		assert.Containsf(t, planModifierDescription(t, attr.PlanModifiers[0]), "destroy and recreate", "%s plan modifier should be RequiresReplace", name)
	}

	schemasAttr, ok := s.Attributes["schemas"].(schema.ListAttribute)
	require.True(t, ok, "schemas must be a list attribute")
	require.Len(t, schemasAttr.PlanModifiers, 1)

	jobInputsAttr, ok := s.Attributes["job_inputs_json"].(schema.StringAttribute)
	require.True(t, ok)
	assert.True(t, jobInputsAttr.Required)
	assert.True(t, jobInputsAttr.Sensitive)

	syncAttr, ok := s.Attributes["sync_frequency"].(schema.StringAttribute)
	require.True(t, ok)
	assert.True(t, syncAttr.Computed)
	assert.False(t, syncAttr.Optional, "sync_frequency must not be configurable — the create endpoint does not accept it")
}

func TestExternalDataSourceBuildCreateRequest(t *testing.T) {
	ops := ExternalDataSourceOps{}
	model := ExternalDataSourceTFModel{
		SourceType:    types.StringValue(testEDSStripeType),
		Prefix:        types.StringValue(testEDSPrefix),
		JobInputsJSON: types.StringValue(`{"stripe_account_id":"acct_123","stripe_secret_key":"sk_test"}`),
		Schemas:       mustListOf(t, "charges", "customers"),
	}

	req, diags := ops.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError(), diags.Errors())

	// Bug-fix regression: prefix must be at the top level of the request, not
	// inside payload. The PostHog create serializer expects it as a sibling of
	// source_type and payload.
	assert.Equal(t, testEDSStripeType, req["source_type"])
	assert.Equal(t, testEDSPrefix, req["prefix"], "prefix must be at the top level, not inside payload")

	payload, ok := req["payload"].(map[string]any)
	require.True(t, ok, "payload should be a map[string]any")
	assert.Equal(t, "acct_123", payload["stripe_account_id"])
	assert.Equal(t, "sk_test", payload["stripe_secret_key"])
	_, hasPrefixInPayload := payload["prefix"]
	assert.False(t, hasPrefixInPayload, "prefix must not appear inside the payload dict")

	schemas, ok := payload["schemas"].([]map[string]any)
	require.True(t, ok)
	require.Len(t, schemas, 2)
	assert.Equal(t, "charges", schemas[0]["name"])
	assert.Equal(t, true, schemas[0]["should_sync"])
	assert.Equal(t, "customers", schemas[1]["name"])
}

func TestExternalDataSourceBuildCreateRequest_NoPrefix(t *testing.T) {
	ops := ExternalDataSourceOps{}
	model := ExternalDataSourceTFModel{
		SourceType:    types.StringValue("Postgres"),
		Prefix:        types.StringNull(),
		JobInputsJSON: types.StringValue(`{"host":"localhost"}`),
		Schemas:       mustListOf(t, "users"),
	}

	req, diags := ops.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError())

	_, hasPrefix := req["prefix"]
	assert.False(t, hasPrefix, "prefix must be omitted when null")
}

func TestExternalDataSourceBuildCreateRequest_InvalidJSON(t *testing.T) {
	ops := ExternalDataSourceOps{}
	model := ExternalDataSourceTFModel{
		SourceType:    types.StringValue("Postgres"),
		JobInputsJSON: types.StringValue(`not-json`),
		Schemas:       mustListOf(t, "users"),
	}

	_, diags := ops.BuildCreateRequest(context.Background(), model)
	require.True(t, diags.HasError())
	assert.Contains(t, diags.Errors()[0].Summary(), "Invalid job_inputs_json")
}

func TestExternalDataSourceBuildUpdateRequest(t *testing.T) {
	ops := ExternalDataSourceOps{}
	plan := ExternalDataSourceTFModel{
		JobInputsJSON: types.StringValue(`{"host":"new-host","port":5433}`),
	}

	req, diags := ops.BuildUpdateRequest(context.Background(), plan, ExternalDataSourceTFModel{})
	require.False(t, diags.HasError())

	// Bug-fix regression: PATCH body must use `job_inputs` (matches the update
	// serializer), not `payload` (matches the create serializer). Sending under
	// the wrong key silently no-ops the update.
	jobInputs, ok := req["job_inputs"].(map[string]any)
	require.True(t, ok, "PATCH body must contain a job_inputs map")
	assert.Equal(t, "new-host", jobInputs["host"])

	_, hasPayload := req["payload"]
	assert.False(t, hasPayload, "PATCH body must not use payload — that key is silently dropped by the update serializer")
}

func TestExternalDataSourceBuildUpdateRequest_EmptyInputs(t *testing.T) {
	ops := ExternalDataSourceOps{}
	plan := ExternalDataSourceTFModel{
		JobInputsJSON: types.StringValue(`{}`),
	}

	req, diags := ops.BuildUpdateRequest(context.Background(), plan, ExternalDataSourceTFModel{})
	require.False(t, diags.HasError())
	assert.Empty(t, req, "empty job inputs should produce an empty PATCH body")
}

func TestExternalDataSourceMapResponseToModel_PreservesPlanJobInputs(t *testing.T) {
	ops := ExternalDataSourceOps{}
	planValue := `{"host":"db.local","password":"hunter2"}`
	model := ExternalDataSourceTFModel{
		JobInputsJSON: types.StringValue(planValue),
	}

	resp := httpclient.ExternalDataSource{
		ID:         testEDSSourceID,
		SourceType: "Postgres",
		JobInputs:  map[string]any{"host": "db.local", "password": "[REDACTED]"},
	}

	diags := ops.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError())
	// The redacted server value must not overwrite the user's planned secret —
	// otherwise sensitive-attribute consistency checks would fail on apply.
	assert.Equal(t, planValue, model.JobInputsJSON.ValueString())
}

func TestExternalDataSourceMapResponseToModel_ImportFromResponse(t *testing.T) {
	ops := ExternalDataSourceOps{}
	model := ExternalDataSourceTFModel{
		JobInputsJSON: types.StringNull(),
	}

	resp := httpclient.ExternalDataSource{
		ID:         testEDSSourceID,
		SourceType: "Postgres",
		JobInputs:  map[string]any{"host": "db.local"},
	}

	diags := ops.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError())
	assert.Contains(t, model.JobInputsJSON.ValueString(), `"host":"db.local"`)
}

func TestExternalDataSourceMapResponseToModel_SyncFrequencyFromFirstSchema(t *testing.T) {
	ops := ExternalDataSourceOps{}
	model := ExternalDataSourceTFModel{}

	resp := httpclient.ExternalDataSource{
		ID:         testEDSSourceID,
		SourceType: "Stripe",
		Schemas: []httpclient.ExternalDataSourceSchema{
			{Name: "charges", SyncFrequency: util.StringPtr("6hour")},
			{Name: "customers", SyncFrequency: util.StringPtr("day")},
		},
	}

	diags := ops.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError())
	assert.Equal(t, "6hour", model.SyncFrequency.ValueString())
}

func TestExternalDataSourceCRUD_VerifiesRequestShape(t *testing.T) {
	var capturedCreateBody, capturedUpdateBody map[string]any

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			require.NoError(t, json.NewDecoder(r.Body).Decode(&capturedCreateBody))
			w.Header().Set("Content-Type", "application/json")
			require.NoError(t, json.NewEncoder(w).Encode(httpclient.ExternalDataSource{
				ID:         testEDSSourceID,
				SourceType: testEDSStripeType,
				Prefix:     util.StringPtr(testEDSPrefix),
			}))
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			require.NoError(t, json.NewEncoder(w).Encode(httpclient.ExternalDataSource{
				ID:         testEDSSourceID,
				SourceType: testEDSStripeType,
			}))
		case http.MethodPatch:
			require.NoError(t, json.NewDecoder(r.Body).Decode(&capturedUpdateBody))
			w.Header().Set("Content-Type", "application/json")
			require.NoError(t, json.NewEncoder(w).Encode(httpclient.ExternalDataSource{
				ID:         testEDSSourceID,
				SourceType: testEDSStripeType,
			}))
		case http.MethodDelete:
			w.WriteHeader(http.StatusOK)
		default:
			t.Fatalf("unexpected method: %s", r.Method)
		}
	}))
	defer server.Close()

	client := httpclient.NewClient(server.Client(), server.URL, testEDSAPIKey, testEDSClient)
	ops := ExternalDataSourceOps{}
	model := ExternalDataSourceTFModel{
		BaseStringIdentifiable: core.BaseStringIdentifiable{ID: types.StringValue(testEDSSourceID)},
		BaseProjectID:          core.BaseProjectID{ProjectID: types.StringValue(testEDSProjectID)},
		SourceType:             types.StringValue(testEDSStripeType),
		Prefix:                 types.StringValue(testEDSPrefix),
		JobInputsJSON:          types.StringValue(`{"stripe_account_id":"acct_123"}`),
		Schemas:                mustListOf(t, "charges"),
	}

	createReq, diags := ops.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError())
	_, err := ops.Create(context.Background(), client, model, createReq)
	require.NoError(t, err)

	// Confirm the wire shape for create.
	assert.Equal(t, testEDSPrefix, capturedCreateBody["prefix"])
	createPayload, ok := capturedCreateBody["payload"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "acct_123", createPayload["stripe_account_id"])

	_, status, err := ops.Read(context.Background(), client, model)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))

	updateModel := model
	updateModel.JobInputsJSON = types.StringValue(`{"stripe_account_id":"acct_999"}`)
	updateReq, diags := ops.BuildUpdateRequest(context.Background(), updateModel, model)
	require.False(t, diags.HasError())
	_, status, err = ops.Update(context.Background(), client, updateModel, updateReq)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))

	// Confirm the wire shape for update.
	updateInputs, ok := capturedUpdateBody["job_inputs"].(map[string]any)
	require.True(t, ok, "PATCH body must include job_inputs")
	assert.Equal(t, "acct_999", updateInputs["stripe_account_id"])

	status, err = ops.Delete(context.Background(), client, model)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
}

func mustListOf(t *testing.T, values ...string) types.List {
	t.Helper()
	listVal, diags := types.ListValueFrom(context.Background(), types.StringType, values)
	require.False(t, diags.HasError())
	return listVal
}

// planModifierDescription is a small helper to introspect a plan modifier's
// description without needing a typed assertion against the framework's
// internal structs. RequiresReplace modifiers describe themselves as containing
// the word "replace".
func planModifierDescription(t *testing.T, m planmodifier.String) string {
	t.Helper()
	return m.Description(context.Background())
}

package resource

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/posthog/terraform-provider/internal/httpclient"
	"github.com/posthog/terraform-provider/internal/resource/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testESResourceProjectID = "project-123"
	testESResourceID        = "es-123"
	testESResourceDefID     = "def-123"
	testESResourceGroupID   = "pg-123"
	testESResourceEvent     = "checkout_completed"
)

func TestEventSchemaMetadataAndSchema(t *testing.T) {
	require.NotNil(t, NewEventSchema())

	ops := EventSchemaOps{}
	assert.Equal(t, "Event Schema", ops.ResourceName())

	s := ops.Schema()
	eventAttr, ok := s.Attributes["event"].(schema.StringAttribute)
	require.True(t, ok)
	assert.True(t, eventAttr.Required)

	groupAttr, ok := s.Attributes["property_group_id"].(schema.StringAttribute)
	require.True(t, ok)
	assert.True(t, groupAttr.Required)

	defAttr, ok := s.Attributes["event_definition_id"].(schema.StringAttribute)
	require.True(t, ok)
	assert.True(t, defAttr.Computed)
	assert.Empty(t, defAttr.PlanModifiers, "must re-resolve when event changes; no UseStateForUnknown")
}

func TestEventSchemaBuildCreateRequest(t *testing.T) {
	ops := EventSchemaOps{}
	model := EventSchemaTFModel{
		Event:           types.StringValue(testESResourceEvent),
		PropertyGroupID: types.StringValue(testESResourceGroupID),
	}

	req, diags := ops.BuildCreateRequest(context.Background(), model)
	require.False(t, diags.HasError())
	assert.Equal(t, testESResourceEvent, req.EventName)
	assert.Equal(t, testESResourceGroupID, req.PropertyGroupID)
	assert.Empty(t, req.EventDefinition, "resolution happens in Create/Update, not in request building")
}

func TestEventSchemaMapResponseToModel(t *testing.T) {
	ops := EventSchemaOps{}
	resp := httpclient.EventSchema{
		ID:              testESResourceID,
		EventDefinition: testESResourceDefID,
		PropertyGroup:   &httpclient.SchemaPropertyGroup{ID: testESResourceGroupID, Name: "Checkout"},
		EventName:       testESResourceEvent,
	}

	var model EventSchemaTFModel
	diags := ops.MapResponseToModel(context.Background(), resp, &model)
	require.False(t, diags.HasError())
	assert.Equal(t, testESResourceID, model.ID.ValueString())
	assert.Equal(t, testESResourceDefID, model.EventDefinitionID.ValueString())
	assert.Equal(t, testESResourceGroupID, model.PropertyGroupID.ValueString())
	assert.Equal(t, testESResourceEvent, model.Event.ValueString())
}

// eventSchemaTestServer covers create/list/patch/delete plus the event
// definition search and retrieve endpoints used for name resolution.
func eventSchemaTestServer(t *testing.T) *httptest.Server {
	listPath := "/api/projects/project-123/event_schemas/"
	defListPath := "/api/projects/project-123/event_definitions/"

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch {
		case r.Method == http.MethodGet && r.URL.Path == defListPath:
			assert.Equal(t, testESResourceEvent, r.URL.Query().Get("search"))
			_ = json.NewEncoder(w).Encode(httpclient.PaginatedResponse[httpclient.EventDefinition]{
				Results: []httpclient.EventDefinition{{ID: testESResourceDefID, Name: testESResourceEvent}},
			})
		case r.Method == http.MethodGet && r.URL.Path == defListPath+testESResourceDefID+"/":
			_ = json.NewEncoder(w).Encode(httpclient.EventDefinition{ID: testESResourceDefID, Name: testESResourceEvent})
		case r.Method == http.MethodPost && r.URL.Path == listPath:
			var req httpclient.EventSchemaRequest
			require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
			assert.Equal(t, testESResourceDefID, req.EventDefinition, "Create must send the resolved UUID")
			_ = json.NewEncoder(w).Encode(httpclient.EventSchema{ID: testESResourceID, EventDefinition: testESResourceDefID})
		case r.Method == http.MethodGet && r.URL.Path == listPath:
			_ = json.NewEncoder(w).Encode(httpclient.PaginatedResponse[httpclient.EventSchema]{
				Results: []httpclient.EventSchema{{ID: testESResourceID, EventDefinition: testESResourceDefID}},
			})
		case r.Method == http.MethodPatch && r.URL.Path == listPath+testESResourceID+"/":
			_ = json.NewEncoder(w).Encode(httpclient.EventSchema{ID: testESResourceID, EventDefinition: testESResourceDefID})
		case r.Method == http.MethodDelete && r.URL.Path == listPath+testESResourceID+"/":
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
}

func TestEventSchemaOpsResolveAndCRUD(t *testing.T) {
	server := eventSchemaTestServer(t)
	defer server.Close()

	client := httpclient.NewClient(server.Client(), server.URL, "test-key", "test")
	ops := EventSchemaOps{}
	model := EventSchemaTFModel{
		BaseStringIdentifiable: core.BaseStringIdentifiable{ID: types.StringValue(testESResourceID)},
		BaseProjectID:          core.BaseProjectID{ProjectID: types.StringValue(testESResourceProjectID)},
		Event:                  types.StringValue(testESResourceEvent),
		EventDefinitionID:      types.StringValue(testESResourceDefID),
	}
	req := httpclient.EventSchemaRequest{EventName: testESResourceEvent, PropertyGroupID: testESResourceGroupID}

	created, err := ops.Create(context.Background(), client, model, req)
	require.NoError(t, err)
	assert.Equal(t, testESResourceID, created.ID)
	assert.Equal(t, testESResourceEvent, created.EventName, "Create must backfill the resolved event name")

	read, status, err := ops.Read(context.Background(), client, model)
	require.NoError(t, err)
	assert.Equal(t, httpclient.HTTPStatusCode(http.StatusOK), status)
	assert.Equal(t, testESResourceEvent, read.EventName, "Read must recover the event name for state")

	updated, status, err := ops.Update(context.Background(), client, model, req)
	require.NoError(t, err)
	assert.Equal(t, httpclient.HTTPStatusCode(http.StatusOK), status)
	assert.Equal(t, testESResourceEvent, updated.EventName)

	status, err = ops.Delete(context.Background(), client, model)
	require.NoError(t, err)
	assert.Equal(t, httpclient.HTTPStatusCode(http.StatusNoContent), status)
}

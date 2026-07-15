package httpclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testESProjectID   = "project-123"
	testESID          = "es-1"
	testESDefID       = "def-1"
	testESListPath    = "/api/projects/project-123/event_schemas/"
	testESDefListPath = "/api/projects/project-123/event_definitions/"
)

func newEventSchemaTestServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch {
		case r.Method == http.MethodPost && r.URL.Path == testESListPath:
			var req EventSchemaRequest
			require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
			assert.Equal(t, testESDefID, req.EventDefinition)
			_ = json.NewEncoder(w).Encode(EventSchema{ID: testESID, EventDefinition: testESDefID})
		case r.Method == http.MethodGet && r.URL.Path == testESListPath:
			assert.Equal(t, testESDefID, r.URL.Query().Get("event_definition"))
			_ = json.NewEncoder(w).Encode(PaginatedResponse[EventSchema]{
				Results: []EventSchema{{ID: "other"}, {ID: testESID, EventDefinition: testESDefID}},
			})
		case r.Method == http.MethodPatch && r.URL.Path == testESListPath+testESID+"/":
			_ = json.NewEncoder(w).Encode(EventSchema{ID: testESID, EventDefinition: testESDefID})
		case r.Method == http.MethodDelete && r.URL.Path == testESListPath+testESID+"/":
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodGet && r.URL.Path == testESDefListPath:
			// search returns fuzzy matches; exact-name filtering is the client's job
			_ = json.NewEncoder(w).Encode(PaginatedResponse[EventDefinition]{
				Results: []EventDefinition{
					{ID: "def-2", Name: "checkout_completed_v2"},
					{ID: testESDefID, Name: "checkout_completed"},
				},
			})
		case r.Method == http.MethodGet && r.URL.Path == testESDefListPath+testESDefID+"/":
			_ = json.NewEncoder(w).Encode(EventDefinition{ID: testESDefID, Name: "checkout_completed"})
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
}

func TestEventSchemaCRUD(t *testing.T) {
	server := newEventSchemaTestServer(t)
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", "test")
	ctx := context.Background()

	created, err := client.CreateEventSchema(ctx, testESProjectID, EventSchemaRequest{
		EventDefinition: testESDefID,
		PropertyGroupID: "pg-1",
	})
	require.NoError(t, err)
	assert.Equal(t, testESID, created.ID)

	got, status, err := client.GetEventSchema(ctx, testESProjectID, testESID, testESDefID)
	require.NoError(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusOK), status)
	assert.Equal(t, testESID, got.ID)

	_, status, err = client.GetEventSchema(ctx, testESProjectID, "missing-id", testESDefID)
	require.Error(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusNotFound), status)

	_, status, err = client.UpdateEventSchema(ctx, testESProjectID, testESID, EventSchemaRequest{
		EventDefinition: testESDefID, PropertyGroupID: "pg-2",
	})
	require.NoError(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusOK), status)

	status, err = client.DeleteEventSchema(ctx, testESProjectID, testESID)
	require.NoError(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusNoContent), status)
}

func TestFindEventDefinitionByName(t *testing.T) {
	server := newEventSchemaTestServer(t)
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", "test")

	def, err := client.FindEventDefinitionByName(context.Background(), testESProjectID, "checkout_completed")
	require.NoError(t, err)
	assert.Equal(t, testESDefID, def.ID, "must exact-match, not take the first fuzzy result")

	_, err = client.FindEventDefinitionByName(context.Background(), testESProjectID, "never_ingested")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "never_ingested")
	assert.Contains(t, err.Error(), "ingested at least once")
}

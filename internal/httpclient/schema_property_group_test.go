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
	testSPGProjectID      = "project-123"
	testSPGID             = "0196-aaaa"
	testSPGCollectionPath = "/api/projects/project-123/schema_property_groups/"
	testSPGResourcePath   = "/api/projects/project-123/schema_property_groups/0196-aaaa/"
)

func TestSchemaPropertyGroupCRUD(t *testing.T) {
	var lastCreateBody SchemaPropertyGroupRequest

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch {
		case r.Method == http.MethodPost && r.URL.Path == testSPGCollectionPath:
			require.NoError(t, json.NewDecoder(r.Body).Decode(&lastCreateBody))
			_ = json.NewEncoder(w).Encode(SchemaPropertyGroup{ID: testSPGID, Name: "Checkout"})
		case r.Method == http.MethodGet && r.URL.Path == testSPGResourcePath:
			_ = json.NewEncoder(w).Encode(SchemaPropertyGroup{ID: testSPGID, Name: "Checkout"})
		case r.Method == http.MethodPatch && r.URL.Path == testSPGResourcePath:
			_ = json.NewEncoder(w).Encode(SchemaPropertyGroup{ID: testSPGID, Name: "Checkout v2"})
		case r.Method == http.MethodDelete && r.URL.Path == testSPGResourcePath:
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", "test")
	ctx := context.Background()

	name := "Checkout"
	required := true
	props := []SchemaPropertyGroupProperty{{Name: "cart_value", PropertyType: "Numeric", IsRequired: &required}}
	created, err := client.CreateSchemaPropertyGroup(ctx, testSPGProjectID, SchemaPropertyGroupRequest{Name: &name, Properties: &props})
	require.NoError(t, err)
	assert.Equal(t, testSPGID, created.ID)
	require.NotNil(t, lastCreateBody.Properties)
	assert.Equal(t, "cart_value", (*lastCreateBody.Properties)[0].Name)
	assert.Equal(t, "Numeric", (*lastCreateBody.Properties)[0].PropertyType)

	got, status, err := client.GetSchemaPropertyGroup(ctx, testSPGProjectID, testSPGID)
	require.NoError(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusOK), status)
	assert.Equal(t, testSPGID, got.ID)

	updated, status, err := client.UpdateSchemaPropertyGroup(ctx, testSPGProjectID, testSPGID, SchemaPropertyGroupRequest{Name: &name})
	require.NoError(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusOK), status)
	assert.Equal(t, "Checkout v2", updated.Name)

	status, err = client.DeleteSchemaPropertyGroup(ctx, testSPGProjectID, testSPGID)
	require.NoError(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusNoContent), status)
}

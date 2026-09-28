package httpclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/posthog/terraform-provider/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testWarehouseTableID      = "01900000-0000-7000-8000-000000000001"
	testWarehouseTableProject = "proj-1"
)

func warehouseTableCollectionPath() string {
	return "/api/projects/" + testWarehouseTableProject + "/warehouse_tables/"
}

func warehouseTableItemPath() string {
	return warehouseTableCollectionPath() + testWarehouseTableID + "/"
}

func TestCreateWarehouseTable(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, warehouseTableCollectionPath(), r.URL.Path)
		// Create must build the HogQL database, which PostHog needs to check that the name is unique.
		assert.Empty(t, r.URL.Query().Get("include_columns"))

		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "s3_orders", body["name"])
		credential, ok := body["credential"].(map[string]any)
		require.True(t, ok, "POST body must include a credential object")
		assert.Equal(t, "AKIA", credential["access_key"])

		writeJSONResponse(t, w, WarehouseTable{
			ID:         testWarehouseTableID,
			Name:       "s3_orders",
			Format:     "Parquet",
			URLPattern: "https://bucket.s3.amazonaws.com/orders/*.parquet",
		})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	resp, err := client.CreateWarehouseTable(context.Background(), testWarehouseTableProject, map[string]any{
		"name":        "s3_orders",
		"format":      "Parquet",
		"url_pattern": "https://bucket.s3.amazonaws.com/orders/*.parquet",
		"credential":  map[string]any{"access_key": "AKIA", "access_secret": "secret"},
	})

	require.NoError(t, err)
	assert.Equal(t, testWarehouseTableID, resp.ID)
	assert.Equal(t, "Parquet", resp.Format)
}

func TestGetWarehouseTable(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, warehouseTableItemPath(), r.URL.Path)
		assert.Equal(t, "false", r.URL.Query().Get("include_columns"))
		writeJSONResponse(t, w, WarehouseTable{
			ID:        testWarehouseTableID,
			Name:      "s3_orders",
			HogQLName: util.StringPtr("s3_orders"),
			Options:   map[string]any{"csv_allow_double_quotes": true},
		})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	resp, status, err := client.GetWarehouseTable(context.Background(), testWarehouseTableProject, testWarehouseTableID)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	require.NotNil(t, resp.HogQLName)
	assert.Equal(t, "s3_orders", *resp.HogQLName)
	assert.Equal(t, true, resp.Options["csv_allow_double_quotes"])
}

func TestUpdateWarehouseTable(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method)
		assert.Equal(t, warehouseTableItemPath(), r.URL.Path)
		assert.Empty(t, r.URL.Query().Get("include_columns"))

		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "s3_orders_v2", body["name"])

		writeJSONResponse(t, w, WarehouseTable{ID: testWarehouseTableID, Name: "s3_orders_v2"})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	resp, status, err := client.UpdateWarehouseTable(context.Background(), testWarehouseTableProject, testWarehouseTableID, map[string]any{
		"name": "s3_orders_v2",
	})

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	assert.Equal(t, "s3_orders_v2", resp.Name)
}

func TestDeleteWarehouseTable(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, warehouseTableItemPath(), r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	status, err := client.DeleteWarehouseTable(context.Background(), testWarehouseTableProject, testWarehouseTableID)

	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, int(status))
}

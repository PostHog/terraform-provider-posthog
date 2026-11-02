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
	testWarehouseTableID      = "01900000-0000-7000-8000-00000000abcd"
	testWarehouseTableProject = "proj-1"
)

func warehouseTableCollectionPath() string {
	return "/api/projects/" + testWarehouseTableProject + "/warehouse_tables/"
}

func warehouseTableItemPath() string {
	return warehouseTableCollectionPath() + testWarehouseTableID + "/"
}

func testWarehouseTableRequest() WarehouseTableRequest {
	return WarehouseTableRequest{
		Name:       "orders",
		Format:     "Parquet",
		URLPattern: "https://bucket.s3.us-east-2.amazonaws.com/orders/*.parquet",
		Credential: WarehouseTableCredential{AccessKey: "AKIAEXAMPLE", AccessSecret: "shh"},
	}
}

func TestCreateWarehouseTable_SendsTheKeysNestedUnderCredential(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, warehouseTableCollectionPath(), r.URL.Path)

		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, map[string]any{
			"name":        "orders",
			"format":      "Parquet",
			"url_pattern": "https://bucket.s3.us-east-2.amazonaws.com/orders/*.parquet",
			"credential":  map[string]any{"access_key": "AKIAEXAMPLE", "access_secret": "shh"},
		}, body)

		w.Header().Set(jsonContentTypeHeader, jsonContentTypeValue)
		_, _ = w.Write([]byte(`{
			"id": "` + testWarehouseTableID + `",
			"name": "orders",
			"format": "Parquet",
			"url_pattern": "https://bucket.s3.us-east-2.amazonaws.com/orders/*.parquet",
			"credential": {"id": "cred-1", "created_at": "2026-09-28T17:00:00Z"},
			"created_at": "2026-09-28T17:00:00Z"
		}`))
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	resp, err := client.CreateWarehouseTable(context.Background(), testWarehouseTableProject, testWarehouseTableRequest())

	require.NoError(t, err)
	assert.Equal(t, WarehouseTable{
		ID:         testWarehouseTableID,
		Name:       "orders",
		Format:     "Parquet",
		URLPattern: "https://bucket.s3.us-east-2.amazonaws.com/orders/*.parquet",
		CreatedAt:  "2026-09-28T17:00:00Z",
	}, resp)
}

func TestGetWarehouseTable(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, warehouseTableItemPath(), r.URL.Path)
		writeJSONResponse(t, w, WarehouseTable{ID: testWarehouseTableID, Name: "orders"})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	resp, status, err := client.GetWarehouseTable(context.Background(), testWarehouseTableProject, testWarehouseTableID)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	assert.Equal(t, "orders", resp.Name)
}

func TestUpdateWarehouseTable(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method)
		assert.Equal(t, warehouseTableItemPath(), r.URL.Path)
		writeJSONResponse(t, w, WarehouseTable{ID: testWarehouseTableID, Name: "orders"})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	resp, status, err := client.UpdateWarehouseTable(context.Background(), testWarehouseTableProject, testWarehouseTableID, testWarehouseTableRequest())

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	assert.Equal(t, "orders", resp.Name)
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

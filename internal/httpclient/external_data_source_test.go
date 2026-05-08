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
	testEDSID         = "01900000-0000-7000-8000-000000000000"
	testEDSProject    = "proj-1"
	externalSourceAPI = "/api/projects/"
	externalSourceSub = "/external_data_sources/"
)

func externalDataSourceCollectionPath() string {
	return externalSourceAPI + testEDSProject + externalSourceSub
}

func externalDataSourceItemPath() string {
	return externalDataSourceCollectionPath() + testEDSID + "/"
}

func TestCreateExternalDataSource(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, externalDataSourceCollectionPath(), r.URL.Path)

		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "Stripe", body["source_type"])
		assert.Equal(t, "stripe_", body["prefix"])

		writeJSONResponse(t, w, ExternalDataSource{
			ID:         testEDSID,
			SourceType: "Stripe",
			Prefix:     util.StringPtr("stripe_"),
		})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	resp, err := client.CreateExternalDataSource(context.Background(), testEDSProject, map[string]any{
		"source_type": "Stripe",
		"prefix":      "stripe_",
		"payload": map[string]any{
			"stripe_account_id": "acct_123",
			"schemas":           []map[string]any{{"name": "charges", "should_sync": true}},
		},
	})

	require.NoError(t, err)
	assert.Equal(t, testEDSID, resp.ID)
	require.NotNil(t, resp.Prefix)
	assert.Equal(t, "stripe_", *resp.Prefix)
}

func TestGetExternalDataSource(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, externalDataSourceItemPath(), r.URL.Path)
		writeJSONResponse(t, w, ExternalDataSource{
			ID:         testEDSID,
			SourceType: "Stripe",
			Schemas: []ExternalDataSourceSchema{
				{Name: "charges", ShouldSync: true, SyncFrequency: util.StringPtr("6hour")},
			},
		})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	resp, status, err := client.GetExternalDataSource(context.Background(), testEDSProject, testEDSID)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	require.Len(t, resp.Schemas, 1)
	assert.Equal(t, "charges", resp.Schemas[0].Name)
	require.NotNil(t, resp.Schemas[0].SyncFrequency)
	assert.Equal(t, "6hour", *resp.Schemas[0].SyncFrequency)
}

func TestUpdateExternalDataSource(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method)
		assert.Equal(t, externalDataSourceItemPath(), r.URL.Path)

		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		// PATCH must use job_inputs at top level — that's what the PostHog
		// update serializer reads. Sending under `payload` would silently no-op.
		_, ok := body["job_inputs"].(map[string]any)
		assert.True(t, ok, "PATCH body must include job_inputs")

		writeJSONResponse(t, w, ExternalDataSource{ID: testEDSID, SourceType: "Stripe"})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	resp, status, err := client.UpdateExternalDataSource(context.Background(), testEDSProject, testEDSID, map[string]any{
		"job_inputs": map[string]any{"stripe_account_id": "acct_999"},
	})

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	assert.Equal(t, testEDSID, resp.ID)
}

func TestDeleteExternalDataSource(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, externalDataSourceItemPath(), r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	status, err := client.DeleteExternalDataSource(context.Background(), testEDSProject, testEDSID)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
}

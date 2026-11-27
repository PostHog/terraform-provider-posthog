package httpclient

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/posthog/terraform-provider/pkg/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testEndpointProjectID = "123"
	testEndpointName      = "daily-signups"
	testEndpointPath      = "/api/projects/123/endpoints/daily-signups/"
)

func TestEndpointRequests(t *testing.T) {
	emptyOverrides := map[string]string{}
	emptyTags := []string{}

	tests := map[string]struct {
		method   string
		path     string
		wantBody string
		call     func(c PosthogClient) error
	}{
		"create posts to the collection": {
			method:   http.MethodPost,
			path:     "/api/projects/123/endpoints/",
			wantBody: `{"name":"daily-signups","query":{"kind":"HogQLQuery","query":"SELECT 1"}}`,
			call: func(c PosthogClient) error {
				_, err := c.CreateEndpoint(context.Background(), testEndpointProjectID, EndpointRequest{
					Name:  util.StringPtr(testEndpointName),
					Query: map[string]interface{}{"kind": "HogQLQuery", "query": "SELECT 1"},
				})
				return err
			},
		},
		"get reads the endpoint by name": {
			method: http.MethodGet,
			path:   testEndpointPath,
			call: func(c PosthogClient) error {
				_, _, err := c.GetEndpoint(context.Background(), testEndpointProjectID, testEndpointName)
				return err
			},
		},
		"update patches the endpoint and sends empty collections to clear them": {
			method:   http.MethodPatch,
			path:     testEndpointPath,
			wantBody: `{"bucket_overrides":{},"tags":[]}`,
			call: func(c PosthogClient) error {
				_, _, err := c.UpdateEndpoint(context.Background(), testEndpointProjectID, testEndpointName, EndpointRequest{
					BucketOverrides: &emptyOverrides,
					Tags:            &emptyTags,
				})
				return err
			},
		},
		"delete removes the endpoint by name": {
			method: http.MethodDelete,
			path:   testEndpointPath,
			call: func(c PosthogClient) error {
				_, err := c.DeleteEndpoint(context.Background(), testEndpointProjectID, testEndpointName)
				return err
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, tc.method, r.Method)
				assert.Equal(t, tc.path, r.URL.Path)
				if tc.wantBody != "" {
					body, err := io.ReadAll(r.Body)
					require.NoError(t, err)
					assert.JSONEq(t, tc.wantBody, string(body))
				}
				if r.Method == http.MethodDelete {
					w.WriteHeader(http.StatusNoContent)
					return
				}
				writeJSONResponse(t, w, Endpoint{ID: "0196f8c0-0000-0000-0000-000000000001", Name: testEndpointName})
			}))
			defer server.Close()

			require.NoError(t, tc.call(newTestPosthogClient(server)))
		})
	}
}

func TestGetEndpoint_DecodesMaterializationAndVersion(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(jsonContentTypeHeader, jsonContentTypeValue)
		_, _ = w.Write([]byte(`{
			"id": "0196f8c0-0000-0000-0000-000000000001",
			"name": "daily-signups",
			"description": "",
			"query": {"kind": "HogQLQuery", "query": "SELECT 1", "variables": null},
			"is_active": true,
			"data_freshness_seconds": 3600,
			"is_materialized": false,
			"materialization": {"enabled": true, "ready": false, "status": "Running", "can_materialize": true},
			"bucket_overrides": {"timestamp": "hour"},
			"current_version": 3,
			"url": "https://us.posthog.com/api/projects/123/endpoints/daily-signups/run",
			"tags": ["growth"],
			"optional_breakdown_properties": [],
			"created_at": "2026-11-27T10:00:00Z"
		}`))
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	endpoint, status, err := client.GetEndpoint(context.Background(), testEndpointProjectID, testEndpointName)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	assert.Equal(t, testEndpointName, endpoint.Name)
	require.NotNil(t, endpoint.Materialization)
	require.NotNil(t, endpoint.Materialization.Enabled)
	assert.True(t, *endpoint.Materialization.Enabled)
	require.NotNil(t, endpoint.CurrentVersion)
	assert.Equal(t, int64(3), *endpoint.CurrentVersion)
	assert.Equal(t, map[string]string{"timestamp": "hour"}, endpoint.BucketOverrides)
	assert.Equal(t, []string{"growth"}, endpoint.Tags)
}

func TestEndpointRequest_OmitsUnsetFields(t *testing.T) {
	body, err := json.Marshal(EndpointRequest{})
	require.NoError(t, err)
	assert.JSONEq(t, `{}`, string(body))
}

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
	testEarlyAccessFeatureID      = "01900000-0000-7000-8000-0000000000ea"
	testEarlyAccessFeatureProject = "proj-1"
)

func earlyAccessFeatureCollectionPath() string {
	return "/api/projects/" + testEarlyAccessFeatureProject + "/early_access_feature/"
}

func earlyAccessFeatureItemPath() string {
	return earlyAccessFeatureCollectionPath() + testEarlyAccessFeatureID + "/"
}

const testEarlyAccessFeatureResponse = `{
	"id": "` + testEarlyAccessFeatureID + `",
	"feature_flag": {"id": 42, "key": "new-editor", "filters": {"feature_enrollment": true}},
	"name": "New editor",
	"description": "",
	"stage": "beta",
	"documentation_url": "",
	"payload": {},
	"created_at": "2026-10-01T10:00:00Z",
	"created_by": null,
	"assignee": null,
	"user_access_level": "editor"
}`

func TestCreateEarlyAccessFeature_SendsFlagIDAndParsesNestedFlag(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, earlyAccessFeatureCollectionPath(), r.URL.Path)

		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, map[string]any{
			"name":            "New editor",
			"stage":           "beta",
			"feature_flag_id": float64(42),
		}, body)

		w.Header().Set(jsonContentTypeHeader, jsonContentTypeValue)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(testEarlyAccessFeatureResponse))
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", "test")
	flagID := int64(42)
	got, err := client.CreateEarlyAccessFeature(context.Background(), testEarlyAccessFeatureProject, EarlyAccessFeatureRequest{
		Name:          "New editor",
		Stage:         "beta",
		FeatureFlagID: &flagID,
	})
	require.NoError(t, err)
	assert.Equal(t, testEarlyAccessFeatureID, got.ID)
	require.NotNil(t, got.FeatureFlag)
	assert.Equal(t, int64(42), got.FeatureFlag.ID)
	assert.Equal(t, "new-editor", got.FeatureFlag.Key)
}

// Without a flag ID PostHog creates one from the name, so the key must be omitted, not null.
func TestCreateEarlyAccessFeature_OmitsUnsetOptionalFields(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, map[string]any{"name": "New editor", "stage": "draft"}, body)

		w.Header().Set(jsonContentTypeHeader, jsonContentTypeValue)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(testEarlyAccessFeatureResponse))
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", "test")
	_, err := client.CreateEarlyAccessFeature(context.Background(), testEarlyAccessFeatureProject, EarlyAccessFeatureRequest{
		Name:  "New editor",
		Stage: "draft",
	})
	require.NoError(t, err)
}

func TestEarlyAccessFeature_ReadUpdateDeletePaths(t *testing.T) {
	var methods []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, earlyAccessFeatureItemPath(), r.URL.Path)
		methods = append(methods, r.Method)
		if r.Method == http.MethodDelete {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.Header().Set(jsonContentTypeHeader, jsonContentTypeValue)
		_, _ = w.Write([]byte(testEarlyAccessFeatureResponse))
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", "test")
	ctx := context.Background()

	_, _, err := client.GetEarlyAccessFeature(ctx, testEarlyAccessFeatureProject, testEarlyAccessFeatureID)
	require.NoError(t, err)
	_, _, err = client.UpdateEarlyAccessFeature(ctx, testEarlyAccessFeatureProject, testEarlyAccessFeatureID, EarlyAccessFeatureRequest{Name: "x", Stage: "alpha"})
	require.NoError(t, err)
	_, err = client.DeleteEarlyAccessFeature(ctx, testEarlyAccessFeatureProject, testEarlyAccessFeatureID)
	require.NoError(t, err)

	assert.Equal(t, []string{http.MethodGet, http.MethodPatch, http.MethodDelete}, methods)
}

package httpclient

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testSurveyProjectID   = "project-123"
	testSurveyID          = "survey-123"
	testSurveyNameCreated = "Created"
	testSurveyNameFetched = "Fetched"
	testSurveyNameUpdated = "Updated"
	testSurveyPath        = "/api/projects/project-123/surveys/survey-123/"
)

func TestSurveyCRUDRequests(t *testing.T) {
	var createBody map[string]interface{}
	var updateBody map[string]interface{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/api/projects/"+testSurveyProjectID+"/surveys/":
			require.NoError(t, json.NewDecoder(r.Body).Decode(&createBody))
			_ = json.NewEncoder(w).Encode(Survey{ID: testSurveyID, Name: stringPtr(testSurveyNameCreated)})
		case r.Method == http.MethodGet && r.URL.Path == testSurveyPath:
			_ = json.NewEncoder(w).Encode(Survey{ID: testSurveyID, Name: stringPtr(testSurveyNameFetched)})
		case r.Method == http.MethodPut && r.URL.Path == testSurveyPath:
			require.NoError(t, json.NewDecoder(r.Body).Decode(&updateBody))
			_ = json.NewEncoder(w).Encode(Survey{ID: testSurveyID, Name: stringPtr(testSurveyNameUpdated)})
		case r.Method == http.MethodDelete && r.URL.Path == testSurveyPath:
			body, err := io.ReadAll(r.Body)
			require.NoError(t, err)
			assert.Empty(t, body)
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", "test")
	ctx := context.Background()

	created, err := client.CreateSurvey(ctx, testSurveyProjectID, SurveyRequest{
		Name:        testSurveyNameCreated,
		Type:        "popover",
		Description: nil,
	})
	require.NoError(t, err)
	assert.Equal(t, testSurveyID, created.ID)

	fetched, status, err := client.GetSurvey(ctx, testSurveyProjectID, testSurveyID)
	require.NoError(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusOK), status)
	assert.Equal(t, testSurveyID, fetched.ID)

	updated, status, err := client.UpdateSurvey(ctx, testSurveyProjectID, testSurveyID, SurveyRequest{
		Name:        testSurveyNameUpdated,
		Type:        "popover",
		Description: stringPtr("desc"),
	})
	require.NoError(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusOK), status)
	assert.Equal(t, testSurveyID, updated.ID)

	status, err = client.DeleteSurvey(ctx, testSurveyProjectID, testSurveyID)
	require.NoError(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusNoContent), status)

	assert.Equal(t, testSurveyNameCreated, createBody["name"])
	assert.Equal(t, "popover", createBody["type"])
	_, hasTargetingFlagID := createBody["targeting_flag_id"]
	assert.False(t, hasTargetingFlagID)

	assert.Equal(t, testSurveyNameUpdated, updateBody["name"])
	assert.Equal(t, "desc", updateBody["description"])
}

func stringPtr(value string) *string {
	return &value
}

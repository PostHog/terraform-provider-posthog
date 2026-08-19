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
	testLogsAlertProjectID     = "123"
	testLogsAlertID            = "019dbe94-cec8-781b-9470-4a970cd69ebf"
	testLogsAlertDestinationsP = "/api/projects/123/logs/alerts/019dbe94-cec8-781b-9470-4a970cd69ebf/destinations/"
)

func writeDestinationPage(t *testing.T, w http.ResponseWriter, next any, destinations ...LogsAlertDestination) {
	t.Helper()
	writeJSONResponse(t, w, map[string]any{
		"count":    len(destinations),
		"next":     next,
		"previous": nil,
		"results":  destinations,
	})
}

func absoluteNextPageURL(server *httptest.Server, query string) string {
	return server.URL + testLogsAlertDestinationsP + "?" + query
}

func TestListLogsAlertDestinations_PopulatesOnlyTheFieldsOfEachDestinationType(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testLogsAlertDestinationsP, r.URL.Path)

		writeDestinationPage(t, w, nil,
			LogsAlertDestination{
				HogFunctionIDs:   []string{"hf-2", "hf-1"},
				Type:             "slack",
				SlackWorkspaceID: util.Int64Ptr(1),
				SlackChannelID:   util.StringPtr("C0123456789"),
			},
			LogsAlertDestination{
				HogFunctionIDs:     []string{"hf-3"},
				Type:               "webhook",
				RedactedWebhookURL: util.StringPtr("https://example.com/…"),
			},
		)
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	destinations, status, err := client.ListLogsAlertDestinations(context.Background(), testLogsAlertProjectID, testLogsAlertID)

	require.NoError(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusOK), status)
	require.Len(t, destinations, 2)
	assert.Equal(t, []string{"hf-2", "hf-1"}, destinations[0].HogFunctionIDs)
	assert.Equal(t, int64(1), *destinations[0].SlackWorkspaceID)
	assert.Nil(t, destinations[0].RedactedWebhookURL)
	assert.Equal(t, "https://example.com/…", *destinations[1].RedactedWebhookURL)
	assert.Nil(t, destinations[1].SlackChannelID)
}

func TestListLogsAlertDestinations_ReturnsDestinationsFromEveryPage(t *testing.T) {
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, testLogsAlertDestinationsP, r.URL.Path)

		if r.URL.Query().Get("offset") == "" {
			writeDestinationPage(t, w, absoluteNextPageURL(server, "limit=1&offset=1"),
				LogsAlertDestination{
					HogFunctionIDs:     []string{"hf-1"},
					Type:               "webhook",
					RedactedWebhookURL: util.StringPtr("https://first.example.com/…"),
				})
			return
		}

		assert.Equal(t, "1", r.URL.Query().Get("offset"))
		writeDestinationPage(t, w, nil,
			LogsAlertDestination{
				HogFunctionIDs:     []string{"hf-2"},
				Type:               "teams",
				RedactedWebhookURL: util.StringPtr("https://second.example.com/…"),
			})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	destinations, status, err := client.ListLogsAlertDestinations(context.Background(), testLogsAlertProjectID, testLogsAlertID)

	require.NoError(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusOK), status)
	require.Len(t, destinations, 2)
	assert.Equal(t, []string{"hf-1"}, destinations[0].HogFunctionIDs)
	assert.Equal(t, []string{"hf-2"}, destinations[1].HogFunctionIDs)
}

func TestCreateLogsAlertDestination_ReturnsOnlyTheNewHogFunctionIDs(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testLogsAlertDestinationsP, r.URL.Path)

		writeJSONResponse(t, w, LogsAlertDestination{HogFunctionIDs: []string{"hf-1", "hf-2"}})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	destination, err := client.CreateLogsAlertDestination(context.Background(), testLogsAlertProjectID, testLogsAlertID, LogsAlertDestinationRequest{
		Type:             "slack",
		SlackWorkspaceID: util.Int64Ptr(1),
		SlackChannelID:   util.StringPtr("C0123456789"),
		SlackChannelName: util.StringPtr("#alerts"),
	})

	require.NoError(t, err)
	assert.Equal(t, []string{"hf-1", "hf-2"}, destination.HogFunctionIDs)
	assert.Empty(t, destination.Type)
}

func TestCreateLogsAlertDestination_OmitsTheAttributesOfOtherDestinationTypes(t *testing.T) {
	var body map[string]any

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		writeJSONResponse(t, w, LogsAlertDestination{HogFunctionIDs: []string{"hf-1"}})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	_, err := client.CreateLogsAlertDestination(context.Background(), testLogsAlertProjectID, testLogsAlertID, LogsAlertDestinationRequest{
		Type:             "slack",
		SlackWorkspaceID: util.Int64Ptr(1),
		SlackChannelID:   util.StringPtr("C0123456789"),
		SlackChannelName: util.StringPtr("#alerts"),
	})

	require.NoError(t, err)
	assert.Equal(t, "slack", body["type"])
	assert.EqualValues(t, 1, body["slack_workspace_id"])
	assert.Equal(t, "C0123456789", body["slack_channel_id"])
	assert.Equal(t, "#alerts", body["slack_channel_name"])
	assert.NotContains(t, body, "webhook_url")
}

func TestDeleteLogsAlertDestination_SendsTheWholeGroupInOneCall(t *testing.T) {
	var calls int

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testLogsAlertDestinationsP+"delete", r.URL.Path)

		var body logsAlertDestinationDeleteRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, []string{"hf-1", "hf-2"}, body.HogFunctionIDs)

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	status, err := client.DeleteLogsAlertDestination(context.Background(), testLogsAlertProjectID, testLogsAlertID, []string{"hf-1", "hf-2"})

	require.NoError(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusNoContent), status)
	assert.Equal(t, 1, calls)
}

func TestListLogsAlertDestinations_ReturnsTheNotFoundStatusToTheCaller(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	_, status, err := client.ListLogsAlertDestinations(context.Background(), testLogsAlertProjectID, testLogsAlertID)

	require.Error(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusNotFound), status)
}

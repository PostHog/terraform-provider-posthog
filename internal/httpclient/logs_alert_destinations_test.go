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

func TestListLogsAlertDestinations(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, testLogsAlertDestinationsP, r.URL.Path)

		// The paginated envelope the rest of the API uses, carrying webhook_url redacted to
		// scheme and host the way the read returns it.
		writeJSONResponse(t, w, map[string]any{
			"count":    2,
			"next":     nil,
			"previous": nil,
			"results": []LogsAlertDestination{
				{
					HogFunctionIDs:   []string{"hf-2", "hf-1"},
					Type:             "slack",
					SlackWorkspaceID: util.Int64Ptr(1),
					SlackChannelID:   util.StringPtr("C0123456789"),
				},
				{
					HogFunctionIDs: []string{"hf-3"},
					Type:           "webhook",
					WebhookURL:     util.StringPtr("https://example.com/…"),
				},
			},
		})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	destinations, status, err := client.ListLogsAlertDestinations(context.Background(), testLogsAlertProjectID, testLogsAlertID)

	require.NoError(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusOK), status)
	require.Len(t, destinations, 2)
	assert.Equal(t, []string{"hf-2", "hf-1"}, destinations[0].HogFunctionIDs)
	assert.Equal(t, int64(1), *destinations[0].SlackWorkspaceID)
	assert.Nil(t, destinations[0].WebhookURL)
	assert.Equal(t, "https://example.com/…", *destinations[1].WebhookURL)
	assert.Nil(t, destinations[1].SlackChannelID)
}

// A destination on a later page has to come back too. The caller finds its own destination by
// scanning the list, so a first-page-only read would report it deleted and Terraform would
// drop a live destination from state.
func TestListLogsAlertDestinations_FollowsPages(t *testing.T) {
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, testLogsAlertDestinationsP, r.URL.Path)

		if r.URL.Query().Get("offset") == "" {
			writeJSONResponse(t, w, map[string]any{
				"count": 2,
				// Absolute, as PostHog builds it. Only the path and query are reused.
				"next":     server.URL + testLogsAlertDestinationsP + "?limit=1&offset=1",
				"previous": nil,
				"results": []LogsAlertDestination{
					{HogFunctionIDs: []string{"hf-1"}, Type: "webhook", WebhookURL: util.StringPtr("https://first.example.com/…")},
				},
			})
			return
		}

		assert.Equal(t, "1", r.URL.Query().Get("offset"))
		writeJSONResponse(t, w, map[string]any{
			"count":    2,
			"next":     nil,
			"previous": server.URL + testLogsAlertDestinationsP + "?limit=1",
			"results": []LogsAlertDestination{
				{HogFunctionIDs: []string{"hf-2"}, Type: "teams", WebhookURL: util.StringPtr("https://second.example.com/…")},
			},
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

func TestCreateLogsAlertDestination(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testLogsAlertDestinationsP, r.URL.Path)

		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "slack", body["type"])
		assert.EqualValues(t, 1, body["slack_workspace_id"])
		assert.Equal(t, "C0123456789", body["slack_channel_id"])
		assert.Equal(t, "#alerts", body["slack_channel_name"])
		// Fields the type does not use must be omitted rather than sent as null.
		assert.NotContains(t, body, "webhook_url")

		// The create response carries the new hog function ids and nothing else.
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

func TestDeleteLogsAlertDestination(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, testLogsAlertDestinationsP+"delete", r.URL.Path)

		var body logsAlertDestinationDeleteRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		// The whole group goes in one call: PostHog deletes the hog functions atomically.
		assert.Equal(t, []string{"hf-1", "hf-2"}, body.HogFunctionIDs)

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	status, err := client.DeleteLogsAlertDestination(context.Background(), testLogsAlertProjectID, testLogsAlertID, []string{"hf-1", "hf-2"})

	require.NoError(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusNoContent), status)
}

// A 404 has to reach the caller with its status intact. The resource decides what it means:
// a deleted alert, or a PostHog too old to serve the endpoint.
func TestListLogsAlertDestinations_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	_, status, err := client.ListLogsAlertDestinations(context.Background(), testLogsAlertProjectID, testLogsAlertID)

	require.Error(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusNotFound), status)
}

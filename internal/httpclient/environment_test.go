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
	testEnvironmentID   = "123"
	apiEnvironmentsPath = "/api/environments/123/"
)

func TestGetEnvironment(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, apiEnvironmentsPath, r.URL.Path)

		writeJSONResponse(t, w, Environment{
			ID:                         123,
			HeatmapsOptIn:              util.BoolPtr(true),
			AutocaptureExceptionsOptIn: util.BoolPtr(false),
			SessionRecordingOptIn:      util.BoolPtr(true),
			SurveysOptIn:               util.BoolPtr(false),
			CookielessServerHashMode:   util.Int64Ptr(2),
			AutocaptureWebVitalsOptIn:  util.BoolPtr(true),
			AnonymizeIps:               util.BoolPtr(true),
		})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	env, status, err := client.GetEnvironment(context.Background(), testEnvironmentID)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	assert.Equal(t, int64(123), env.ID)
	require.NotNil(t, env.HeatmapsOptIn)
	assert.True(t, *env.HeatmapsOptIn)
	require.NotNil(t, env.AutocaptureExceptionsOptIn)
	assert.False(t, *env.AutocaptureExceptionsOptIn)
	require.NotNil(t, env.SessionRecordingOptIn)
	assert.True(t, *env.SessionRecordingOptIn)
	require.NotNil(t, env.SurveysOptIn)
	assert.False(t, *env.SurveysOptIn)
	require.NotNil(t, env.CookielessServerHashMode)
	assert.Equal(t, int64(2), *env.CookielessServerHashMode)
	require.NotNil(t, env.AutocaptureWebVitalsOptIn)
	assert.True(t, *env.AutocaptureWebVitalsOptIn)
	require.NotNil(t, env.AnonymizeIps)
	assert.True(t, *env.AnonymizeIps)
}

func TestGetEnvironmentDeserializesNetworkPayloadCaptureFromRawJSON(t *testing.T) {
	// The other tests round-trip Environment through this package's own json tags,
	// which would agree with each other even if the tags were wrong. This handler
	// writes a hand-written body instead (mirroring a real /api/environments/{id}/
	// response) to pin the inbound contract: camelCase keys inside the
	// session_recording_network_payload_capture_config blob, snake_case around it.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, apiEnvironmentsPath, r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{
			"id": 123,
			"capture_performance_opt_in": true,
			"session_recording_network_payload_capture_config": {"recordHeaders": true, "recordBody": false}
		}`))
		require.NoError(t, err)
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	env, status, err := client.GetEnvironment(context.Background(), testEnvironmentID)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	require.NotNil(t, env.CapturePerformanceOptIn)
	assert.True(t, *env.CapturePerformanceOptIn)
	require.NotNil(t, env.SessionRecordingNetworkPayloadCaptureConfig)
	require.NotNil(t, env.SessionRecordingNetworkPayloadCaptureConfig.RecordHeaders)
	assert.True(t, *env.SessionRecordingNetworkPayloadCaptureConfig.RecordHeaders)
	require.NotNil(t, env.SessionRecordingNetworkPayloadCaptureConfig.RecordBody)
	assert.False(t, *env.SessionRecordingNetworkPayloadCaptureConfig.RecordBody, "explicit false must deserialize as a set value")
}

// TestGetEnvironmentDeserializesTestAccountFilters pins the inbound contract for
// test_account_filters: an array of filter objects, including the server-injected
// cohort_name that a cohort-referencing filter carries.
func TestGetEnvironmentDeserializesTestAccountFilters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{
			"id": 123,
			"test_account_filters": [{"key":"id","type":"cohort","value":2,"operator":"in","cohort_name":"Internal users"}],
			"test_account_filters_default_checked": true
		}`))
		require.NoError(t, err)
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	env, status, err := client.GetEnvironment(context.Background(), testEnvironmentID)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	require.NotNil(t, env.TestAccountFilters)
	require.Len(t, *env.TestAccountFilters, 1)
	filter, ok := (*env.TestAccountFilters)[0].(map[string]interface{})
	require.True(t, ok)
	assert.Equal(t, "cohort", filter["type"])
	assert.Equal(t, float64(2), filter["value"])
	assert.Equal(t, "Internal users", filter["cohort_name"])
	require.NotNil(t, env.TestAccountFiltersDefaultChecked)
	assert.True(t, *env.TestAccountFiltersDefaultChecked)
}

func TestUpdateEnvironmentSerializesTestAccountFilters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method)

		var raw map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&raw))
		filters, ok := raw["test_account_filters"].([]any)
		require.True(t, ok, "test_account_filters must be present as an array")
		require.Len(t, filters, 1)
		assert.Equal(t, true, raw["test_account_filters_default_checked"])

		writeJSONResponse(t, w, Environment{ID: 123})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	_, status, err := client.UpdateEnvironment(context.Background(), testEnvironmentID, EnvironmentSettingsRequest{
		TestAccountFilters: &[]interface{}{
			map[string]interface{}{"key": "id", "type": "cohort", "value": 2, "operator": "in"},
		},
		TestAccountFiltersDefaultChecked: util.BoolPtr(true),
	})

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
}

func TestUpdateEnvironmentSerializesOnlySetFields(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method)
		assert.Equal(t, apiEnvironmentsPath, r.URL.Path)

		// Decode into a raw map to prove that only set fields are present on the wire.
		var raw map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&raw))
		assert.Equal(t, true, raw["heatmaps_opt_in"])
		_, hasSurveys := raw["surveys_opt_in"]
		assert.False(t, hasSurveys, "unset fields must be omitted from the PATCH body")
		_, hasHash := raw["cookieless_server_hash_mode"]
		assert.False(t, hasHash, "unset fields must be omitted from the PATCH body")

		writeJSONResponse(t, w, Environment{
			ID:            123,
			HeatmapsOptIn: util.BoolPtr(true),
		})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	env, status, err := client.UpdateEnvironment(context.Background(), testEnvironmentID, EnvironmentSettingsRequest{
		HeatmapsOptIn: util.BoolPtr(true),
	})

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
	require.NotNil(t, env.HeatmapsOptIn)
	assert.True(t, *env.HeatmapsOptIn)
}

func TestUpdateEnvironmentSerializesZeroValues(t *testing.T) {
	// Zero-values (false / 0) must be sent on the wire, not dropped: omitempty on a
	// pointer only fires on nil. This guards against a future switch to value types.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method)

		var raw map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&raw))

		heatmaps, hasHeatmaps := raw["heatmaps_opt_in"]
		require.True(t, hasHeatmaps, "explicit false must be present in the PATCH body")
		assert.Equal(t, false, heatmaps)

		hash, hasHash := raw["cookieless_server_hash_mode"]
		require.True(t, hasHash, "explicit 0 must be present in the PATCH body")
		assert.Equal(t, float64(0), hash) // JSON numbers decode to float64

		writeJSONResponse(t, w, Environment{ID: 123})
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	_, status, err := client.UpdateEnvironment(context.Background(), testEnvironmentID, EnvironmentSettingsRequest{
		HeatmapsOptIn:            util.BoolPtr(false),
		CookielessServerHashMode: util.Int64Ptr(0),
	})

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, int(status))
}

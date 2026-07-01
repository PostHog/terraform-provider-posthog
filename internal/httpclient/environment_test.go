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

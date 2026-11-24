package httpclient

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetEventsRetention(t *testing.T) {
	for _, tc := range []struct {
		name     string
		response string
		want     *int64
	}{
		{"project period", `{"retention_months":6,"project_retention_months":12}`, new(int64(12))},
		{"organization period", `{"retention_months":6,"project_retention_months":null}`, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, "/api/projects/123/events_retention/", r.URL.Path)
				w.Header().Set(jsonContentTypeHeader, jsonContentTypeValue)
				_, err := w.Write([]byte(tc.response))
				assert.NoError(t, err)
			}))
			defer server.Close()

			client := newTestPosthogClient(server)
			result, _, err := client.GetEventsRetention(context.Background(), "123")
			require.NoError(t, err)
			assert.Equal(t, tc.want, result.ProjectRetentionMonths)
		})
	}
}

func TestSetEventsRetention(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/projects/123/events_retention/", r.URL.Path)
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.JSONEq(t, `{"retention_months":12}`, string(body))
		w.Header().Set(jsonContentTypeHeader, jsonContentTypeValue)
		_, err = w.Write([]byte(`{"retention_months":12,"project_retention_months":12}`))
		assert.NoError(t, err)
	}))
	defer server.Close()

	client := newTestPosthogClient(server)
	result, _, err := client.SetEventsRetention(context.Background(), "123", EventsRetentionRequest{RetentionMonths: 12})
	require.NoError(t, err)
	assert.Equal(t, new(int64(12)), result.ProjectRetentionMonths)
}

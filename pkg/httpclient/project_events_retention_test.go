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

func TestProjectEventsRetention(t *testing.T) {
	for _, tc := range []struct {
		name, method, response, body string
		status                       int
		want                         *int64
	}{
		{"get set", http.MethodGet, `{"retention_months":6,"project_retention_months":12}`, "", 200, new(int64(12))},
		{"set", http.MethodPost, `{"retention_months":12,"project_retention_months":12}`, `{"retention_months":12}`, 200, new(int64(12))},
		{"clear", http.MethodPost, `{"retention_months":6,"project_retention_months":null}`, `{"retention_months":null}`, 200, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, tc.method, r.Method)
				assert.Equal(t, "/api/projects/123/events_retention/", r.URL.Path)
				if tc.body != "" {
					body, err := io.ReadAll(r.Body)
					if err != nil {
						t.Errorf("read request: %v", err)
						http.Error(w, "invalid request", http.StatusBadRequest)
						return
					}
					assert.JSONEq(t, tc.body, string(body))
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tc.status)
				_, err := w.Write([]byte(tc.response))
				assert.NoError(t, err)
			}))
			defer server.Close()
			client := newTestPosthogClient(server)
			var result ProjectEventsRetention
			var err error
			if tc.method == http.MethodGet {
				result, _, err = client.GetProjectEventsRetention(context.Background(), "123")
			} else {
				result, _, err = client.SetProjectEventsRetention(context.Background(), "123", ProjectEventsRetentionRequest{RetentionMonths: tc.want})
			}
			require.NoError(t, err)
			assert.Equal(t, tc.want, result.ProjectRetentionMonths)
		})
	}
}

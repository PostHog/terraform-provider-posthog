package httpclient

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/posthog/terraform-provider/internal/httpclient/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultClient_LeavesTimeoutToTheRetryTransport(t *testing.T) {
	client := NewDefaultClient("https://us.posthog.com", "key", "test")

	// An http.Client.Timeout would bound the retry loop and its backoff sleeps,
	// so a request that failed by timing out could never be retried.
	assert.Zero(t, client.httpClient.Timeout)

	transport, ok := client.httpClient.Transport.(*middleware.RetryTransport)
	require.True(t, ok, "expected the retry transport to be installed")
	assert.Equal(t, middleware.DefaultAttemptTimeout, transport.Config.AttemptTimeout)
	assert.Equal(t, middleware.DefaultMaxRetries, transport.Config.MaxRetries)

	_, ok = transport.Base.(*http.Transport)
	assert.True(t, ok, "expected the retry transport to wrap the pooled http.Transport")
}

// failingServer counts requests and always answers 500.
func failingServer(t *testing.T, attempts *atomic.Int32) *httptest.Server {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts.Add(1)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	t.Cleanup(server.Close)
	return server
}

func TestClientOptions_ReplaceTheRetryTransport(t *testing.T) {
	// NewDefaultClient has already installed a retry transport. Stacking a second
	// one on top would leave the inner loop retrying, and would make the outer
	// per-attempt deadline cover the inner loop end to end.
	tests := map[string]struct {
		option           ClientOption
		expectedAttempts int32
	}{
		"no retry":    {WithNoRetry(), 1},
		"one retry":   {WithRetryConfig(middleware.RetryConfig{MaxRetries: 1}), 2},
		"three tries": {WithRetryConfig(middleware.RetryConfig{MaxRetries: 2}), 3},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var attempts atomic.Int32
			server := failingServer(t, &attempts)

			client := NewDefaultClient(server.URL, "key", "test", tt.option)
			_, status, err := client.doRequest(context.Background(), http.MethodGet, "/api/whatever", nil)

			require.Error(t, err)
			assert.Equal(t, HTTPStatusCode(http.StatusInternalServerError), status)
			assert.Equal(t, tt.expectedAttempts, attempts.Load())
		})
	}
}

func TestClientOptions_DoNotNestAttemptDeadlines(t *testing.T) {
	var attempts atomic.Int32
	server := failingServer(t, &attempts)

	client := NewDefaultClient(server.URL, "key", "test", WithNoRetry())

	transport, ok := client.httpClient.Transport.(*middleware.RetryTransport)
	require.True(t, ok)
	_, nested := transport.Base.(*middleware.RetryTransport)
	assert.False(t, nested, "the option should replace the retry transport, not wrap it")
}

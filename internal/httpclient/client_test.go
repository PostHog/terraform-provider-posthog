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

func TestClientOptions_DoNotStackRetryTransports(t *testing.T) {
	client := NewDefaultClient("https://example.invalid", "key", "test", WithNoRetry())

	transport, ok := client.httpClient.Transport.(*middleware.RetryTransport)
	require.True(t, ok)
	_, nested := transport.Base.(*middleware.RetryTransport)
	assert.False(t, nested, "the option should replace the retry transport, not wrap it")
}

func TestDoRequest_RateLimitErrorSaysHowLongItWaitedAndWhy(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "1800")
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"detail":"Request was throttled. Expected available in 1800 seconds."}`))
	}))
	defer server.Close()

	client := NewDefaultClient(server.URL, testAPIKey, testClientVersion)
	_, status, err := client.doRequest(context.Background(), http.MethodGet, "/api/projects/1/feature_flags/1/", nil)

	require.Error(t, err)
	assert.Equal(t, HTTPStatusCode(http.StatusTooManyRequests), status)
	assert.ErrorContains(t, err, "rate limited by PostHog; stopped because", "an immediate stop reports no wait")
	assert.ErrorContains(t, err, "Retry-After of 30m0s")
	assert.ErrorContains(t, err, "-parallelism")
	assert.ErrorContains(t, err, "Expected available in 1800 seconds", "PostHog's own message is kept")
}

func TestRateLimitError(t *testing.T) {
	stopped := func(waited, reason string) *http.Response {
		resp := &http.Response{StatusCode: http.StatusTooManyRequests, Header: http.Header{}}
		if reason != "" {
			resp.Header.Set(middleware.RateLimitWaitedHeader, waited)
			resp.Header.Set(middleware.RateLimitStopReasonHeader, reason)
		}
		return resp
	}
	body := []byte(`{"detail":"Request was throttled. Expected available in 30 seconds."}`)

	tests := map[string]struct {
		resp    *http.Response
		want    []string
		wantNot []string
	}{
		"a 429 the transport did not stop on keeps the plain error": {
			resp:    stopped("", ""),
			want:    []string{"API error (status 429): {\"detail\""},
			wantNot: []string{"rate limited by PostHog"},
		},
		"an immediate stop reports no wait": {
			resp:    stopped("0s", "retries are disabled"),
			want:    []string{"rate limited by PostHog; stopped because retries are disabled", "-parallelism", "Expected available in 30 seconds"},
			wantNot: []string{"after waiting"},
		},
		"a stop after waiting says how long": {
			resp: stopped("4m12s", "waiting again would pass the 5m0s rate-limit budget"),
			want: []string{"rate limited by PostHog after waiting 4m12s; stopped because waiting again would pass the 5m0s rate-limit budget"},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := rateLimitError(tt.resp, body)

			for _, w := range tt.want {
				assert.ErrorContains(t, err, w)
			}
			for _, w := range tt.wantNot {
				assert.NotContains(t, err.Error(), w)
			}
		})
	}
}

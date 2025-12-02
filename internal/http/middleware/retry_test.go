package middleware

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// testRetryConfig returns a fast retry config for testing.
func testRetryConfig() RetryConfig {
	return RetryConfig{
		MaxRetries:           3,
		InitialBackoff:       10 * time.Millisecond,
		MaxBackoff:           100 * time.Millisecond,
		BackoffFactor:        2.0,
		JitterFactor:         0, // Disable jitter for predictable tests
		RetryableStatusCodes: []int{429, 500, 502, 503, 504},
	}
}

func TestRetryTransport_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	transport := NewRetryTransport(http.DefaultTransport, NoRetryConfig())
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := client.Do(req)

	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestRetryTransport_RetryableStatusCodes(t *testing.T) {
	tests := map[string]struct {
		failStatus              int
		failUntilAttempt        int32
		expectedAttempts        int32
		expectedStatus          int
		shouldEventuallySucceed bool
	}{
		"retry on 500 and succeed": {
			failStatus:              http.StatusInternalServerError,
			failUntilAttempt:        3,
			expectedAttempts:        3,
			expectedStatus:          http.StatusOK,
			shouldEventuallySucceed: true,
		},
		"retry on 502 and succeed": {
			failStatus:              http.StatusBadGateway,
			failUntilAttempt:        2,
			expectedAttempts:        2,
			expectedStatus:          http.StatusOK,
			shouldEventuallySucceed: true,
		},
		"retry on 503 and succeed": {
			failStatus:              http.StatusServiceUnavailable,
			failUntilAttempt:        2,
			expectedAttempts:        2,
			expectedStatus:          http.StatusOK,
			shouldEventuallySucceed: true,
		},
		"retry on 504 and succeed": {
			failStatus:              http.StatusGatewayTimeout,
			failUntilAttempt:        2,
			expectedAttempts:        2,
			expectedStatus:          http.StatusOK,
			shouldEventuallySucceed: true,
		},
		"exhaust retries on persistent 500": {
			failStatus:              http.StatusInternalServerError,
			failUntilAttempt:        100, // Never succeed
			expectedAttempts:        4,   // 1 initial + 3 retries
			expectedStatus:          http.StatusInternalServerError,
			shouldEventuallySucceed: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var attempts atomic.Int32

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				attempt := attempts.Add(1)
				if attempt < tt.failUntilAttempt {
					w.WriteHeader(tt.failStatus)
					w.Write([]byte(`{"error": "server error"}`))
					return
				}
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"status": "ok"}`))
			}))
			defer server.Close()

			transport := NewRetryTransport(http.DefaultTransport, testRetryConfig())
			client := &http.Client{Transport: transport}

			req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
			resp, err := client.Do(req)

			require.NoError(t, err)
			defer resp.Body.Close()
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
			assert.Equal(t, tt.expectedAttempts, attempts.Load())
		})
	}
}

func TestRetryTransport_NonRetryableStatusCodes(t *testing.T) {
	nonRetryableCodes := []int{
		http.StatusBadRequest,          // 400
		http.StatusUnauthorized,        // 401
		http.StatusForbidden,           // 403
		http.StatusNotFound,            // 404
		http.StatusMethodNotAllowed,    // 405
		http.StatusUnprocessableEntity, // 422
	}

	for _, statusCode := range nonRetryableCodes {
		t.Run(http.StatusText(statusCode), func(t *testing.T) {
			var attempts atomic.Int32

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				attempts.Add(1)
				w.WriteHeader(statusCode)
				w.Write([]byte(`{"error": "client error"}`))
			}))
			defer server.Close()

			transport := NewRetryTransport(http.DefaultTransport, testRetryConfig())
			client := &http.Client{Transport: transport}

			req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
			resp, err := client.Do(req)

			require.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, statusCode, resp.StatusCode)
			assert.Equal(t, int32(1), attempts.Load(), "should not retry on %d", statusCode)
		})
	}
}

func TestRetryTransport_RetryOn429WithRetryAfter(t *testing.T) {
	var attempts atomic.Int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempt := attempts.Add(1)
		if attempt < 2 {
			w.Header().Set("Retry-After", "1") // 1 second
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"error": "rate limited"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	// Increment MaxBackOff to confirm that we're respecting the "Retry-After" header
	config := testRetryConfig()
	config.MaxBackoff = 2 * time.Second

	transport := NewRetryTransport(http.DefaultTransport, config)
	client := &http.Client{Transport: transport}

	start := time.Now()
	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := client.Do(req)
	elapsed := time.Since(start)

	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, int32(2), attempts.Load())
	assert.InEpsilon(t, 1*time.Second, elapsed, 0.1, "should have waited for retry after limit")
}

func TestRetryTransport_RetryOn429WithRetryAfter_RespectingLowerMaxBackoff(t *testing.T) {
	var attempts atomic.Int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempt := attempts.Add(1)
		if attempt < 2 {
			w.Header().Set("Retry-After", "1") // 1 second
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"error": "rate limited"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	transport := NewRetryTransport(http.DefaultTransport, testRetryConfig())
	client := &http.Client{Transport: transport}

	start := time.Now()
	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := client.Do(req)
	elapsed := time.Since(start)

	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, int32(2), attempts.Load())
	assert.InEpsilon(t, testRetryConfig().MaxBackoff, elapsed, 0.1, "should have respected the max backoff from our test config")
}

func TestRetryTransport_ContextCancellation(t *testing.T) {
	var attempts atomic.Int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts.Add(1)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "internal server error"}`))
	}))
	defer server.Close()

	ctx, cancel := context.WithCancel(context.Background())
	// Cancel after a short delay (should interrupt retries)
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	transport := NewRetryTransport(http.DefaultTransport, testRetryConfig())
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, server.URL, nil)
	_, err := client.Do(req)

	assert.Error(t, err, "expected error due to context cancellation")
	assert.Less(t, attempts.Load(), int32(4), "should have been cancelled before exhausting all retries")
}

func TestRetryTransport_WithBody(t *testing.T) {
	var receivedBodies []string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		receivedBodies = append(receivedBodies, string(body))

		if len(receivedBodies) < 2 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	transport := NewRetryTransport(http.DefaultTransport, testRetryConfig())
	client := &http.Client{Transport: transport}

	bodyContent := `{"key": "value"}`
	req, _ := http.NewRequest(http.MethodPost, server.URL, strings.NewReader(bodyContent))
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(strings.NewReader(bodyContent)), nil
	}

	resp, err := client.Do(req)

	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Len(t, receivedBodies, 2, "expected 2 attempts")

	for i, body := range receivedBodies {
		assert.Equal(t, bodyContent, body, "attempt %d should have received the body", i+1)
	}
}

func TestRetryTransport_OnRetryCallback(t *testing.T) {
	var attempts atomic.Int32
	var callbackInvocations []int

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempt := attempts.Add(1)
		if attempt < 3 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := testRetryConfig()
	config.OnRetry = func(attempt int, resp *http.Response, err error) {
		callbackInvocations = append(callbackInvocations, attempt)
	}

	transport := NewRetryTransport(http.DefaultTransport, config)
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := client.Do(req)

	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Len(t, callbackInvocations, 2, "callback should be invoked twice (after attempt 1 and 2)")
	assert.Equal(t, []int{1, 2}, callbackInvocations)
}

func TestCalculateBackoff(t *testing.T) {
	config := RetryConfig{
		InitialBackoff: 100 * time.Millisecond,
		MaxBackoff:     10 * time.Second,
		BackoffFactor:  2.0,
		JitterFactor:   0, // Disable jitter for predictable tests
	}
	tests := []struct {
		attempt  int
		expected time.Duration
	}{
		{0, 100 * time.Millisecond},  // 100ms * 2^0 = 100ms
		{1, 200 * time.Millisecond},  // 100ms * 2^1 = 200ms
		{2, 400 * time.Millisecond},  // 100ms * 2^2 = 400ms
		{3, 800 * time.Millisecond},  // 100ms * 2^3 = 800ms
		{4, 1600 * time.Millisecond}, // 100ms * 2^4 = 1600ms
		{10, 10 * time.Second},       // Would be 102.4s, but capped at 10s
	}
	transport := NewRetryTransport(http.DefaultTransport, config)

	for _, tt := range tests {
		t.Run(http.StatusText(tt.attempt), func(t *testing.T) {
			backoff := transport.calculateBackoff(tt.attempt)
			assert.Equal(t, tt.expected, backoff)
		})
	}
}

func TestCalculateBackoff_WithJitter(t *testing.T) {
	config := RetryConfig{
		InitialBackoff: 100 * time.Millisecond,
		MaxBackoff:     10 * time.Second,
		BackoffFactor:  2.0,
		JitterFactor:   0.2, // ±20% jitter
	}

	baseBackoff := 100 * time.Millisecond
	minExpected := time.Duration(float64(baseBackoff) * 0.8)
	maxExpected := time.Duration(float64(baseBackoff) * 1.2)
	transport := NewRetryTransport(http.DefaultTransport, config)

	uniqueValues := make(map[time.Duration]bool)
	for i := 0; i < 100; i++ {
		backoff := transport.calculateBackoff(0)
		uniqueValues[backoff] = true

		assert.GreaterOrEqual(t, backoff, minExpected, "backoff should be >= min")
		assert.LessOrEqual(t, backoff, maxExpected, "backoff should be <= max")
	}

	assert.GreaterOrEqual(t, len(uniqueValues), 10, "jitter should produce varied backoffs")
}

func TestParseRetryAfter(t *testing.T) {
	tests := []struct {
		name     string
		header   string
		expected time.Duration
		ok       bool
	}{
		{"empty", "", 0, false},
		{"zero", "0", 0, false},
		{"negative", "-1", 0, false},
		{"one second", "1", 1 * time.Second, true},
		{"five seconds", "5", 5 * time.Second, true},
		{"sixty seconds", "60", 60 * time.Second, true},
		{"invalid", "invalid", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duration, ok := parseRetryAfterHeader(tt.header)
			assert.Equal(t, tt.ok, ok)
			assert.Equal(t, tt.expected, duration)
		})
	}
}

func TestIsRetryableError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{"context.Canceled", context.Canceled, false},
		{"context.DeadlineExceeded", context.DeadlineExceeded, true},
		{"nil error", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, isRetryableError(tt.err))
		})
	}
}

func TestSleep(t *testing.T) {
	t.Run("cancelled context returns immediately", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		err := sleep(ctx, 10*time.Second)
		assert.ErrorIs(t, err, context.Canceled)
	})

	t.Run("completes after duration", func(t *testing.T) {
		start := time.Now()
		err := sleep(context.Background(), 50*time.Millisecond)
		elapsed := time.Since(start)

		assert.NoError(t, err)
		assert.InEpsilon(t, 50*time.Millisecond, elapsed, 0.1)
	})
}

package middleware

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net"
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
		AttemptTimeout:       5 * time.Second,
		RateLimitMaxWait:     2 * time.Second,
		RateLimitTotalWait:   5 * time.Second,
		RetryableStatusCodes: []int{429, 500, 502, 503, 504},
	}
}

func TestRetryTransport_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	transport := NewRetryTransport(http.DefaultTransport, NoRetryConfig())
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := client.Do(req)

	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

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
					_, _ = w.Write([]byte(`{"error": "server error"}`))
					return
				}
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(`{"status": "ok"}`))
			}))
			defer server.Close()

			transport := NewRetryTransport(http.DefaultTransport, testRetryConfig())
			client := &http.Client{Transport: transport}

			req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
			resp, err := client.Do(req)

			require.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()
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
				_, _ = w.Write([]byte(`{"error": "client error"}`))
			}))
			defer server.Close()

			transport := NewRetryTransport(http.DefaultTransport, testRetryConfig())
			client := &http.Client{Transport: transport}

			req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
			resp, err := client.Do(req)

			require.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, statusCode, resp.StatusCode)
			assert.Equal(t, int32(1), attempts.Load(), "should not retry on %d", statusCode)
		})
	}
}

func TestRetryTransport_RetryOn429WaitsTheFullRetryAfterBeyondMaxBackoff(t *testing.T) {
	var attempts atomic.Int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if attempts.Add(1) < 2 {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// MaxBackoff is 100ms in the test config. It bounds the waits between 5xx
	// retries, not the wait PostHog asks for on a 429.
	transport := NewRetryTransport(http.DefaultTransport, testRetryConfig())
	client := &http.Client{Transport: transport}

	start := time.Now()
	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := client.Do(req)
	elapsed := time.Since(start)

	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, int32(2), attempts.Load())
	assert.GreaterOrEqual(t, elapsed, 1*time.Second, "should have waited the whole 1s Retry-After")
}

func TestRetryTransport_RateLimitRetriesDoNotUseUpMaxRetries(t *testing.T) {
	var attempts atomic.Int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if attempts.Add(1) <= 2 {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := testRetryConfig()
	config.MaxRetries = 1
	transport := NewRetryTransport(http.DefaultTransport, config)
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := client.Do(req)

	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, int32(3), attempts.Load(), "two 429s then a success, despite MaxRetries = 1")
}

func TestRetryTransport_RetryAfterBeyondRateLimitMaxWaitIsReturnedAtOnce(t *testing.T) {
	var attempts atomic.Int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts.Add(1)
		// PostHog's hourly limit asks for a wait of minutes.
		w.Header().Set("Retry-After", "1800")
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"detail": "Request was throttled. Expected available in 1800 seconds."}`))
	}))
	defer server.Close()

	transport := NewRetryTransport(http.DefaultTransport, testRetryConfig())
	client := &http.Client{Transport: transport}

	start := time.Now()
	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := client.Do(req)
	elapsed := time.Since(start)

	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusTooManyRequests, resp.StatusCode, "the 429 reaches the caller, whose error carries PostHog's message")
	assert.Equal(t, "0s", resp.Header.Get(RateLimitWaitedHeader))
	assert.Contains(t, resp.Header.Get(RateLimitStopReasonHeader), "Retry-After of 30m0s is longer than the 2s")
	assert.Equal(t, int32(1), attempts.Load())
	assert.Less(t, elapsed, time.Second, "should not have waited at all")
	body, _ := io.ReadAll(resp.Body)
	assert.Contains(t, string(body), "Expected available in 1800 seconds")
}

func TestRetryTransport_RateLimitWaitStopsAtTheTotalBudget(t *testing.T) {
	var attempts atomic.Int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts.Add(1)
		w.Header().Set("Retry-After", "1")
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer server.Close()

	config := testRetryConfig()
	config.RateLimitTotalWait = 1500 * time.Millisecond
	transport := NewRetryTransport(http.DefaultTransport, config)
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := client.Do(req)

	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
	// One 1s wait fits the 1.5s budget; a second would not, so it is not started.
	assert.Equal(t, int32(2), attempts.Load())
	assert.Equal(t, "1s", resp.Header.Get(RateLimitWaitedHeader))
	assert.Contains(t, resp.Header.Get(RateLimitStopReasonHeader), "rate-limit budget")
}

func TestRetryTransport_RateLimitBudgetCountsSlowThrottledAttempts(t *testing.T) {
	var attempts atomic.Int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts.Add(1)
		// Each 429 takes a large share of the 1s wait it asks for.
		time.Sleep(600 * time.Millisecond)
		w.Header().Set("Retry-After", "1")
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer server.Close()

	config := testRetryConfig()
	config.RateLimitTotalWait = 3 * time.Second
	transport := NewRetryTransport(http.DefaultTransport, config)
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := client.Do(req)

	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
	// Counting only the 1s sleeps would allow four attempts in a 3s budget;
	// counting the 600ms attempts too allows two.
	assert.LessOrEqual(t, attempts.Load(), int32(3))
	assert.GreaterOrEqual(t, attempts.Load(), int32(2), "a rate-limited request is still retried")
}

func TestRetryTransport_ZeroRetryAfterUsesTheRateLimitBudget(t *testing.T) {
	var attempts atomic.Int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if attempts.Add(1) <= 2 {
			// What PostHog sends when a slot frees up within the second.
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := testRetryConfig()
	config.MaxRetries = 1
	transport := NewRetryTransport(http.DefaultTransport, config)
	client := &http.Client{Transport: transport}

	start := time.Now()
	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := client.Do(req)

	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "two 429s then a success, despite MaxRetries = 1")
	assert.Equal(t, int32(3), attempts.Load())
	assert.GreaterOrEqual(t, time.Since(start), 2*time.Second, "each zero Retry-After still waits a second")
}

func TestRetryTransport_A429WithoutRetryAfterIsExplainedWhenRetriesRunOut(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer server.Close()

	transport := NewRetryTransport(http.DefaultTransport, testRetryConfig())
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := client.Do(req)

	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	waited, reason, ok := RateLimitStop(resp)
	require.True(t, ok, "the caller gets the rate-limit explanation, not a bare 429")
	assert.Contains(t, reason, "no usable Retry-After and 3 retries ran out")
	assert.Zero(t, waited, "the 10-80ms backoffs round down to no wait")
}

func TestRetryTransport_NoRetryConfigDoesNotRetryRateLimits(t *testing.T) {
	var attempts atomic.Int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts.Add(1)
		w.Header().Set("Retry-After", "1")
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer server.Close()

	transport := NewRetryTransport(http.DefaultTransport, NoRetryConfig())
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := client.Do(req)

	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
	assert.Equal(t, int32(1), attempts.Load())
}

func TestRetryTransport_ContextCancellation(t *testing.T) {
	var attempts atomic.Int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts.Add(1)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "internal server error"}`))
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
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	transport := NewRetryTransport(http.DefaultTransport, testRetryConfig())
	client := &http.Client{Transport: transport}

	// PUT rather than POST: only idempotent methods are replayed on a 500.
	bodyContent := `{"key": "value"}`
	req, _ := http.NewRequest(http.MethodPut, server.URL, strings.NewReader(bodyContent))
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(strings.NewReader(bodyContent)), nil
	}

	resp, err := client.Do(req)

	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

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
	defer func() { _ = resp.Body.Close() }()

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
		{"zero, a sub-second wait rounded down", "0", 0, true},
		{"negative", "-1", 0, false},
		{"one second", "1", 1 * time.Second, true},
		{"five seconds", "5", 5 * time.Second, true},
		{"sixty seconds", "60", 60 * time.Second, true},
		{"invalid", "invalid", 0, false},
		{"longer than a day is clamped", "90000", 24 * time.Hour, true},
		{"too large for a Duration is clamped, not wrapped", "9223372037", 24 * time.Hour, true},
		{"a value that wraps to exactly zero is clamped", "36028797018963968", 24 * time.Hour, true},
		{"too large for an int64 is clamped", "99999999999999999999", 24 * time.Hour, true},
		{"too negative for an int64", "-99999999999999999999", 0, false},
		{"padded with spaces", " 5 ", 0, false},
		{"fractional", "1.5", 0, false},
		{"an HTTP date is not used", "Wed, 21 Oct 2026 07:28:00 GMT", 0, false},
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
		// A 10% band was flaky: a timer can only fire late, never early.
		assert.GreaterOrEqual(t, elapsed, 50*time.Millisecond)
		assert.Less(t, elapsed, 5*time.Second)
	})
}

// awaitOrCancel stalls a handler for d, or until the client gives up. Sleeping
// the full duration would make httptest's Close wait it out after the attempt
// deadline has already closed the connection.
func awaitOrCancel(r *http.Request, d time.Duration) {
	select {
	case <-r.Context().Done():
	case <-time.After(d):
	}
}

func TestRetryTransport_RetriesAfterAttemptTimeout(t *testing.T) {
	var attempts atomic.Int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if attempts.Add(1) < 3 {
			awaitOrCancel(r, 600*time.Millisecond)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	config := testRetryConfig()
	config.AttemptTimeout = 150 * time.Millisecond

	transport := NewRetryTransport(http.DefaultTransport, config)
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := client.Do(req)

	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, int32(3), attempts.Load(), "the first two attempts should time out and still be retried")
}

func TestRetryTransport_AppliesAttemptTimeoutWithoutRetries(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		awaitOrCancel(r, 600*time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := NoRetryConfig()
	config.AttemptTimeout = 100 * time.Millisecond

	transport := NewRetryTransport(http.DefaultTransport, config)
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	_, err := client.Do(req)

	require.Error(t, err)
	assert.ErrorIs(t, err, context.DeadlineExceeded, "the attempt deadline is the only timeout left on this path")
}

func TestRetryTransport_ResponseBodyOutlivesRoundTrip(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.(http.Flusher).Flush()
		// Delay the body so it is read after RoundTrip has already returned.
		time.Sleep(100 * time.Millisecond)
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	transport := NewRetryTransport(http.DefaultTransport, testRetryConfig())

	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := transport.RoundTrip(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "the body is buffered during the attempt, so it reads back fine")
	assert.Equal(t, `{"status": "ok"}`, string(body))
}

// A body that stalls has to fail the attempt, not the caller's later read.
// RoundTrip used to return as soon as the headers arrived, so the timeout landed
// outside the retry loop, where it could be neither retried nor explained.
func TestRetryTransport_RetriesWhenTheBodyStalls(t *testing.T) {
	var attempts atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.(http.Flusher).Flush()
		if attempts.Add(1) < 3 {
			// Headers arrive, the body never does.
			awaitOrCancel(r, 3*time.Second)
			return
		}
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	config := testRetryConfig()
	config.AttemptTimeout = 150 * time.Millisecond
	client := &http.Client{Transport: NewRetryTransport(http.DefaultTransport, config)}

	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := client.Do(req)

	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, `{"status": "ok"}`, string(body))
	assert.Equal(t, int32(3), attempts.Load(), "a stalled body should fail its attempt and be retried")
}

// The same stall on a create must not be replayed, and the error has to say so:
// the write may have landed even though the body never arrived.
func TestRetryTransport_StalledBodyOnCreateSaysItMayHaveApplied(t *testing.T) {
	var attempts atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts.Add(1)
		w.WriteHeader(http.StatusOK)
		w.(http.Flusher).Flush()
		awaitOrCancel(r, 3*time.Second)
	}))
	defer server.Close()

	config := testRetryConfig()
	config.AttemptTimeout = 150 * time.Millisecond
	client := &http.Client{Transport: NewRetryTransport(http.DefaultTransport, config)}

	req, _ := http.NewRequest(http.MethodPost, server.URL, nil)
	_, err := client.Do(req)

	require.Error(t, err)
	assert.ErrorContains(t, err, "not retried, POST may already have been applied")
	assert.ErrorContains(t, err, "attempt timed out after 150ms")
	assert.Equal(t, int32(1), attempts.Load())
}
func TestRetryTransport_RetriesOnlyIdempotentMethods(t *testing.T) {
	tests := map[string]struct {
		method           string
		status           int
		expectedAttempts int32
	}{
		"GET retries on 500":          {http.MethodGet, http.StatusInternalServerError, 4},
		"HEAD retries on 500":         {http.MethodHead, http.StatusInternalServerError, 4},
		"PUT retries on 502":          {http.MethodPut, http.StatusBadGateway, 4},
		"DELETE retries on 503":       {http.MethodDelete, http.StatusServiceUnavailable, 4},
		"POST does not retry on 500":  {http.MethodPost, http.StatusInternalServerError, 1},
		"POST does not retry on 502":  {http.MethodPost, http.StatusBadGateway, 1},
		"POST does not retry on 504":  {http.MethodPost, http.StatusGatewayTimeout, 1},
		"PATCH does not retry on 503": {http.MethodPatch, http.StatusServiceUnavailable, 1},
		// A 429 was never processed, so it is replayed for any method. With no
		// Retry-After it gets the ordinary MaxRetries bound.
		"POST retries on 429":  {http.MethodPost, http.StatusTooManyRequests, 4},
		"PATCH retries on 429": {http.MethodPatch, http.StatusTooManyRequests, 4},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var attempts atomic.Int32

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				attempts.Add(1)
				w.WriteHeader(tt.status)
			}))
			defer server.Close()

			transport := NewRetryTransport(http.DefaultTransport, testRetryConfig())
			client := &http.Client{Transport: transport}

			req, _ := http.NewRequest(tt.method, server.URL, nil)
			resp, err := client.Do(req)

			require.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.status, resp.StatusCode)
			assert.Equal(t, tt.expectedAttempts, attempts.Load())
		})
	}
}

// errorTransport fails every request, standing in for a connection that drops
// without the caller learning whether the server saw the request.
type errorTransport struct {
	attempts atomic.Int32
	err      error
}

func (e *errorTransport) RoundTrip(*http.Request) (*http.Response, error) {
	e.attempts.Add(1)
	return nil, e.err
}

func TestRetryTransport_RetriesTransportErrorsOnlyForIdempotentMethods(t *testing.T) {
	tests := map[string]struct {
		method           string
		expectedAttempts int32
	}{
		"GET retries":          {http.MethodGet, 4},
		"DELETE retries":       {http.MethodDelete, 4},
		"POST does not retry":  {http.MethodPost, 1},
		"PATCH does not retry": {http.MethodPatch, 1},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			base := &errorTransport{err: errors.New("connection reset by peer")}
			transport := NewRetryTransport(base, testRetryConfig())

			req, _ := http.NewRequest(tt.method, "http://example.invalid", nil)
			_, err := transport.RoundTrip(req)

			require.Error(t, err)
			assert.Equal(t, tt.expectedAttempts, base.attempts.Load())
		})
	}
}

func TestNewRetryTransport_FillsUnsetTunables(t *testing.T) {
	// A hand-built config must not end up with no deadline, or with a zero
	// backoff that would retry with no spacing between attempts.
	for name, config := range map[string]RetryConfig{
		"no retries": NoRetryConfig(),
		"zero value": {MaxRetries: 2},
	} {
		t.Run(name, func(t *testing.T) {
			transport := NewRetryTransport(http.DefaultTransport, config)
			got := transport.Config

			assert.Equal(t, DefaultAttemptTimeout, got.AttemptTimeout)
			assert.Equal(t, DefaultInitialBackoff, got.InitialBackoff)
			assert.Equal(t, DefaultMaxBackoff, got.MaxBackoff)
			assert.Equal(t, DefaultBackoffFactor, got.BackoffFactor)
			assert.Equal(t, DefaultRateLimitMaxWait, got.RateLimitMaxWait)
			assert.Equal(t, DefaultRateLimitTotalWait, got.RateLimitTotalWait)
			assert.Equal(t, defaultRetryableStatusCodes(), got.RetryableStatusCodes)
			assert.Positive(t, transport.calculateBackoff(0), "retries must be spaced out")
		})
	}
}

func TestNewRetryTransport_KeepsMeaningfulZeroes(t *testing.T) {
	// Zero is a real choice for both of these, so filling them in would silently
	// override the caller.
	got := NewRetryTransport(http.DefaultTransport, RetryConfig{MaxRetries: 0, JitterFactor: 0}).Config

	assert.Zero(t, got.MaxRetries)
	assert.Zero(t, got.JitterFactor)
}

func TestRetryDecision(t *testing.T) {
	transport := NewRetryTransport(http.DefaultTransport, testRetryConfig())

	tests := map[string]struct {
		method     string
		resp       *http.Response
		err        error
		ctx        context.Context
		header     http.Header
		wantRetry  bool
		wantReason string
	}{
		"get on a retryable status": {
			method: http.MethodGet, resp: &http.Response{StatusCode: 500},
			wantRetry: true, wantReason: "retryable failure",
		},
		"post on a retryable status": {
			method: http.MethodPost, resp: &http.Response{StatusCode: 502},
			wantReason: "POST may already have been applied",
		},
		"post on a rate limit with Retry-After": {
			method: http.MethodPost, resp: &http.Response{StatusCode: 429, Header: http.Header{"Retry-After": []string{"5"}}},
			wantRetry: true, wantReason: "rate limited",
		},
		"a zero Retry-After is still a rate limit": {
			method: http.MethodGet, resp: &http.Response{StatusCode: 429, Header: http.Header{"Retry-After": []string{"0"}}},
			wantRetry: true, wantReason: "rate limited",
		},
		"post on a rate limit without Retry-After": {
			method: http.MethodPost, resp: &http.Response{StatusCode: 429},
			wantRetry: true, wantReason: "rate limited without Retry-After",
		},
		"status is not retryable": {
			method: http.MethodGet, resp: &http.Response{StatusCode: 404},
			wantReason: "failure is not retryable",
		},
		"post on a transport error": {
			method: http.MethodPost, err: errors.New("connection reset by peer"),
			wantReason: "POST may already have been applied",
		},
		"post carrying an idempotency key": {
			method: http.MethodPost, resp: &http.Response{StatusCode: 502},
			header:    http.Header{"Idempotency-Key": []string{"abc123"}},
			wantRetry: true, wantReason: "retryable failure",
		},
		"post carrying the x-prefixed key": {
			method: http.MethodPost, resp: &http.Response{StatusCode: 500},
			header:    http.Header{"X-Idempotency-Key": []string{"abc123"}},
			wantRetry: true, wantReason: "retryable failure",
		},
		"an empty key still counts, as it does in net/http": {
			method: http.MethodPost, resp: &http.Response{StatusCode: 500},
			header:    http.Header{"Idempotency-Key": []string{""}},
			wantRetry: true, wantReason: "retryable failure",
		},
		"post that never left the machine": {
			method: http.MethodPost, err: &net.OpError{Op: "dial", Err: errors.New("connection refused")},
			wantRetry: true, wantReason: "retryable failure",
		},
		"caller cancelled": {
			method: http.MethodPost, err: context.Canceled, ctx: cancelledContext(),
			wantReason: "caller's context ended",
		},
		"caller's own deadline expired": {
			method: http.MethodGet, err: context.DeadlineExceeded, ctx: cancelledContext(),
			wantReason: "caller's context ended",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := tt.ctx
			if ctx == nil {
				ctx = context.Background()
			}
			req, err := http.NewRequestWithContext(ctx, tt.method, "http://example.invalid/x", nil)
			require.NoError(t, err)
			for name, values := range tt.header {
				req.Header[name] = values
			}

			got := transport.decide(req, tt.resp, tt.err)

			assert.Equal(t, tt.wantRetry, got.retry)
			assert.Equal(t, tt.wantReason, got.reason)
		})
	}
}

func cancelledContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

// The message an operator actually reads is the one worth pinning.
func TestRetryTransport_ErrorNamesAttemptAndDeadline(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		awaitOrCancel(r, 3*time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := testRetryConfig()
	config.AttemptTimeout = 80 * time.Millisecond
	client := &http.Client{Transport: NewRetryTransport(http.DefaultTransport, config)}

	tests := map[string]struct {
		method      string
		wantAttempt string
	}{
		"idempotent method retries to exhaustion": {http.MethodGet, "(attempt: 4)"},
		"create is not replayed after a timeout":  {http.MethodPost, "(attempt: 1)"},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.method, server.URL, nil)
			_, err := client.Do(req)

			require.Error(t, err)
			assert.ErrorContains(t, err, tt.wantAttempt)
			assert.ErrorContains(t, err, "attempt timed out after 80ms")
			if tt.method == http.MethodPost {
				assert.ErrorContains(t, err, "not retried, POST may already have been applied")
			}
			assert.ErrorIs(t, err, context.DeadlineExceeded, "callers still match on the cause")
		})
	}
}

// A caller deadline expiring mid-attempt must not be reported as a failed
// backoff. That message is the one this whole change set out to remove.
func TestRetryTransport_CallerDeadlineDoesNotLookLikeAFailedBackoff(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		awaitOrCancel(r, 3*time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := testRetryConfig()
	config.AttemptTimeout = 5 * time.Second
	config.InitialBackoff = 200 * time.Millisecond
	client := &http.Client{Transport: NewRetryTransport(http.DefaultTransport, config)}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Millisecond)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, server.URL, nil)
	_, err := client.Do(req)

	require.Error(t, err)
	assert.NotContains(t, err.Error(), "sleeping", "the caller's deadline ended it, not a backoff")
}

func TestRetryTransport_RetriesUnsentRequestsForAnyMethod(t *testing.T) {
	// A dial that never connected sent nothing, so replaying it cannot duplicate
	// a resource the way a response-side failure could.
	for _, method := range []string{http.MethodGet, http.MethodPost, http.MethodPatch} {
		t.Run(method, func(t *testing.T) {
			base := &errorTransport{err: &net.OpError{Op: "dial", Err: errors.New("connection refused")}}
			transport := NewRetryTransport(base, testRetryConfig())

			req, _ := http.NewRequest(method, "http://example.invalid", nil)
			_, err := transport.RoundTrip(req)

			require.Error(t, err)
			assert.Equal(t, int32(4), base.attempts.Load())
		})
	}
}

func TestRetryTransport_Stopped(t *testing.T) {
	tests := map[string]struct {
		maxRetries  int
		decision    retryDecision
		exhausted   bool
		wantMessage string
		wantReason  string
	}{
		"ran out of attempts": {
			maxRetries: 3, decision: retryDecision{retry: true}, exhausted: true,
			wantMessage: "giving up on request", wantReason: "retries exhausted",
		},
		"retries turned off": {
			maxRetries: 0, decision: retryDecision{retry: true}, exhausted: true,
			wantMessage: "not retrying request", wantReason: "retries are disabled",
		},
		"never retryable in the first place": {
			maxRetries: 3, decision: retryDecision{reason: "failure is not retryable"},
			wantMessage: "not retrying request", wantReason: "failure is not retryable",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			config := testRetryConfig()
			config.MaxRetries = tt.maxRetries
			transport := NewRetryTransport(http.DefaultTransport, config)

			message, reason := transport.stopped(tt.decision, tt.exhausted)

			assert.Equal(t, tt.wantMessage, message)
			assert.Equal(t, tt.wantReason, reason)
		})
	}
}

func TestReplaceRetryTransport_DoesNotStack(t *testing.T) {
	base := http.DefaultTransport
	once := ReplaceRetryTransport(base, DefaultRetryConfig())
	twice := ReplaceRetryTransport(once, NoRetryConfig())

	assert.Same(t, base, twice.Base, "replacing must unwrap, not nest")
	assert.Zero(t, twice.Config.MaxRetries)
}

func TestRetryTransport_RateLimitWaitAddsJitterButNeverWaitsLess(t *testing.T) {
	config := testRetryConfig()
	config.JitterFactor = 0.2
	transport := NewRetryTransport(http.DefaultTransport, config)

	seen := map[time.Duration]bool{}
	for range 50 {
		wait, stopReason := transport.rateLimitBackoff(time.Second, 0)
		require.Empty(t, stopReason)
		assert.GreaterOrEqual(t, wait, time.Second, "never sooner than PostHog asked")
		assert.LessOrEqual(t, wait, 1200*time.Millisecond)
		seen[wait] = true
	}
	assert.Greater(t, len(seen), 1, "parallel requests given the same Retry-After should not wake together")
}

func TestRetryTransport_CancellingDuringARateLimitWaitStopsAtOnce(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "2")
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer server.Close()

	transport := NewRetryTransport(http.DefaultTransport, testRetryConfig())
	client := &http.Client{Transport: transport}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, server.URL, nil)

	start := time.Now()
	_, err := client.Do(req)

	require.Error(t, err)
	assert.ErrorContains(t, err, "sleeping")
	assert.Less(t, time.Since(start), time.Second, "Ctrl-C must not wait out the Retry-After")
}

func TestRetryTransport_OnRetryFiresForRateLimitWaits(t *testing.T) {
	var attempts, callbacks atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if attempts.Add(1) <= 2 {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := testRetryConfig()
	config.OnRetry = func(_ int, resp *http.Response, _ error) {
		assert.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
		callbacks.Add(1)
	}
	transport := NewRetryTransport(http.DefaultTransport, config)
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := client.Do(req)

	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()
	assert.Equal(t, int32(2), callbacks.Load())
}

func TestRetryTransport_RateLimitBackoff(t *testing.T) {
	// RateLimitMaxWait is 2s and RateLimitTotalWait 5s in the test config, and
	// jitter is off.
	tests := map[string]struct {
		retryAfter     time.Duration
		rateLimitedFor time.Duration
		maxRetries     int
		want           time.Duration
		wantStop       string
	}{
		"Retry-After is honoured past MaxBackoff":         {retryAfter: time.Second, want: time.Second},
		"a zero Retry-After still waits a second":         {retryAfter: 0, want: time.Second},
		"Retry-After up to RateLimitMaxWait is honoured":  {retryAfter: 2 * time.Second, want: 2 * time.Second},
		"Retry-After past RateLimitMaxWait is not waited": {retryAfter: 3 * time.Second, wantStop: "Retry-After of 3s is longer than the 2s"},
		"a clamped huge Retry-After is not waited":        {retryAfter: maxRetryAfter, wantStop: "is longer than"},
		"a wait past the budget is not started":           {retryAfter: time.Second, rateLimitedFor: 4500 * time.Millisecond, wantStop: "5s rate-limit budget"},
		"no retries at all with MaxRetries 0":             {retryAfter: time.Second, maxRetries: -1, wantStop: "retries are disabled"},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			config := testRetryConfig()
			if tt.maxRetries < 0 {
				config.MaxRetries = 0
			}
			transport := NewRetryTransport(http.DefaultTransport, config)

			got, stopReason := transport.rateLimitBackoff(tt.retryAfter, tt.rateLimitedFor)

			if tt.wantStop != "" {
				assert.Contains(t, stopReason, tt.wantStop)
				return
			}
			assert.Empty(t, stopReason)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRateLimitStop(t *testing.T) {
	stopped := func(waited, reason string) *http.Response {
		h := http.Header{}
		if waited != "" {
			h.Set(RateLimitWaitedHeader, waited)
		}
		if reason != "" {
			h.Set(RateLimitStopReasonHeader, reason)
		}
		return &http.Response{StatusCode: 429, Header: h}
	}

	waited, reason, ok := RateLimitStop(stopped("42s", "budget spent"))
	assert.True(t, ok)
	assert.Equal(t, 42*time.Second, waited)
	assert.Equal(t, "budget spent", reason)

	_, _, ok = RateLimitStop(stopped("", ""))
	assert.False(t, ok, "a 429 the transport did not stop on")
	_, _, ok = RateLimitStop(stopped("soon", "budget spent"))
	assert.False(t, ok, "an unparseable wait")
}

func TestIsIdempotentMethod(t *testing.T) {
	idempotent := []string{
		http.MethodGet, http.MethodHead, http.MethodPut,
		http.MethodDelete, http.MethodOptions, http.MethodTrace,
	}
	for _, method := range idempotent {
		assert.True(t, isIdempotentMethod(method), "%s should be idempotent", method)
	}

	for _, method := range []string{http.MethodPost, http.MethodPatch, http.MethodConnect} {
		assert.False(t, isIdempotentMethod(method), "%s should not be idempotent", method)
	}
}

// net/http replays a request itself when a pooled connection turns out to be
// dead before any bytes were written, including a POST, but only when GetBody
// survives to rebuild the body. Cloning per attempt must not drop it.
func TestRetryTransport_AttemptKeepsGetBodyForTheTransport(t *testing.T) {
	body := []byte(`{"name":"thing"}`)
	var seen *http.Request
	base := roundTripFunc(func(r *http.Request) (*http.Response, error) {
		seen = r
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader("{}"))}, nil
	})

	req, err := http.NewRequest(http.MethodPost, "http://example.invalid", bytes.NewReader(body))
	require.NoError(t, err)
	req.GetBody = func() (io.ReadCloser, error) { return io.NopCloser(bytes.NewReader(body)), nil }

	resp, err := NewRetryTransport(base, testRetryConfig()).RoundTrip(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	require.NotNil(t, seen)
	require.NotNil(t, seen.GetBody, "the transport below needs GetBody to replay safely")

	rebuilt, err := seen.GetBody()
	require.NoError(t, err)
	got, err := io.ReadAll(rebuilt)
	require.NoError(t, err)
	assert.Equal(t, body, got)
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

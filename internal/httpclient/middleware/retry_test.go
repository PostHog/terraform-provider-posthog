package middleware

import (
	"context"
	"errors"
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
		AttemptTimeout:       5 * time.Second,
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

func TestRetryTransport_RetryOn429WithRetryAfter(t *testing.T) {
	var attempts atomic.Int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempt := attempts.Add(1)
		if attempt < 2 {
			w.Header().Set("Retry-After", "1") // 1 second
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write([]byte(`{"error": "rate limited"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "ok"}`))
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
	defer func() { _ = resp.Body.Close() }()

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
			_, _ = w.Write([]byte(`{"error": "rate limited"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "ok"}`))
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

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, int32(2), attempts.Load())
	// A 10% band around MaxBackoff was flaky: the two round trips are counted in
	// elapsed too. Assert the property instead, that the 1s Retry-After was
	// clamped down to MaxBackoff.
	assert.GreaterOrEqual(t, elapsed, testRetryConfig().MaxBackoff, "should have waited at least the max backoff")
	// 500ms sits between the clamped backoff and the 1s Retry-After it replaced.
	assert.Less(t, elapsed, 500*time.Millisecond, "should have clamped the 1s Retry-After to the max backoff")
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
	require.NoError(t, err, "the attempt context must stay alive until the body is closed")
	assert.Equal(t, `{"status": "ok"}`, string(body))
}

func TestRetryTransport_AttemptTimeoutCoversBodyRead(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.(http.Flusher).Flush()
		awaitOrCancel(r, 600*time.Millisecond)
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	config := testRetryConfig()
	config.AttemptTimeout = 100 * time.Millisecond

	transport := NewRetryTransport(http.DefaultTransport, config)

	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := transport.RoundTrip(req)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	_, err = io.ReadAll(resp.Body)
	require.Error(t, err, "a stalled body read must still hit the attempt deadline")
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
		"POST retries on 429":         {http.MethodPost, http.StatusTooManyRequests, 4},
		"PATCH retries on 429":        {http.MethodPatch, http.StatusTooManyRequests, 4},
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

func TestGiveUpReason(t *testing.T) {
	transport := NewRetryTransport(http.DefaultTransport, testRetryConfig())

	tests := map[string]struct {
		retry  bool
		method string
		resp   *http.Response
		err    error
		want   string
	}{
		"ran out of attempts": {
			retry: true, method: http.MethodGet,
			resp: &http.Response{StatusCode: 500}, want: "retries exhausted",
		},
		"post on a retryable status": {
			method: http.MethodPost,
			resp:   &http.Response{StatusCode: 502}, want: "POST is not safe to replay",
		},
		"post on a transport error": {
			method: http.MethodPost,
			err:    errors.New("connection reset by peer"), want: "POST is not safe to replay",
		},
		"status is not retryable at all": {
			method: http.MethodGet,
			resp:   &http.Response{StatusCode: 404}, want: "failure is not retryable",
		},
		"caller cancelled": {
			method: http.MethodGet,
			err:    context.Canceled, want: "failure is not retryable",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, transport.giveUpReason(tt.retry, tt.method, tt.resp, tt.err))
		})
	}
}

func TestRetryTransport_LogsDecisionsWithoutALogger(t *testing.T) {
	// tflog is a no-op on a context with no provider logger, which is what the
	// retry tests and any direct RoundTrip caller pass.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	transport := NewRetryTransport(http.DefaultTransport, testRetryConfig())

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL, nil)
	resp, err := transport.RoundTrip(req)

	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
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

// Package middleware provides HTTP middleware components.
package middleware

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand/v2"
	"net/http"
	"strconv"
	"time"
)

const (
	DefaultMaxRetries     = 3
	DefaultInitialBackoff = 200 * time.Millisecond
	DefaultMaxBackoff     = 5 * time.Second
	DefaultBackoffFactor  = 2.0
	DefaultJitterFactor   = 0.2 // 20% jitter
	DefaultAttemptTimeout = 30 * time.Second
)

type RetryConfig struct {
	// MaxRetries is the maximum number of retry attempts (0 means no retries).
	MaxRetries int
	// InitialBackoff is the initial delay before the first retry.
	InitialBackoff time.Duration
	// MaxBackoff is the maximum delay between retries.
	MaxBackoff time.Duration
	// BackoffFactor is the multiplier applied to backoff after each retry.
	BackoffFactor float64
	// JitterFactor adds randomness to backoff (0.0-1.0, where 0.2 = ±20%).
	JitterFactor float64
	// AttemptTimeout bounds a single attempt, from dialing through reading the
	// response body. Each attempt gets its own budget, so a request that failed
	// by timing out still has time left to be retried. 0 disables it.
	AttemptTimeout time.Duration
	// RetryableStatusCodes are HTTP status codes that should trigger a retry.
	// If nil, defaults to 429, 500, 502, 503, 504.
	RetryableStatusCodes []int
	// OnRetry is an optional callback invoked before each retry attempt.
	// It receives the attempt number (1-indexed), the response (may be nil), and the error.
	OnRetry func(attempt int, resp *http.Response, err error)
}

func DefaultRetryConfig() RetryConfig {
	return RetryConfig{
		MaxRetries:           DefaultMaxRetries,
		InitialBackoff:       DefaultInitialBackoff,
		MaxBackoff:           DefaultMaxBackoff,
		BackoffFactor:        DefaultBackoffFactor,
		JitterFactor:         DefaultJitterFactor,
		AttemptTimeout:       DefaultAttemptTimeout,
		RetryableStatusCodes: []int{429, 500, 502, 503, 504},
	}
}

// NoRetryConfig disables retries but keeps the per-attempt deadline, which is
// the only timeout on requests once http.Client.Timeout is unset.
func NoRetryConfig() RetryConfig {
	return RetryConfig{MaxRetries: 0, AttemptTimeout: DefaultAttemptTimeout}
}

type RetryTransport struct {
	Base   http.RoundTripper
	Config RetryConfig
}

func NewRetryTransport(base http.RoundTripper, config RetryConfig) *RetryTransport {
	if base == nil {
		base = http.DefaultTransport
	}
	return &RetryTransport{
		Base:   base,
		Config: config,
	}
}

func (t *RetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// If retries are disabled, still go through attempt so the request keeps its
	// deadline.
	if t.Config.MaxRetries == 0 {
		return t.attempt(req)
	}

	for attempt := 0; ; attempt++ {
		if err := req.Context().Err(); err != nil {
			return nil, fmt.Errorf("(attempt: %d) context error: %w", attempt+1, err)
		}

		resp, err := t.attempt(req)

		// Success case
		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return resp, nil
		}

		// If not retryable or last attempt, return
		if !t.shouldRetry(req.Method, resp, err) || attempt >= t.Config.MaxRetries {
			if err != nil {
				return nil, err
			}
			return resp, nil
		}

		backoff := t.calculateBackoff(attempt)
		if resp != nil {
			// Close response body before retry to avoid resource leak
			if resp.Body != nil {
				_ = resp.Body.Close()
			}
			if resp.StatusCode == http.StatusTooManyRequests {
				if retryAfter, ok := parseRetryAfterHeader(resp.Header.Get("Retry-After")); ok {
					if retryAfter > backoff {
						backoff = retryAfter
					}
					if backoff > t.Config.MaxBackoff {
						backoff = t.Config.MaxBackoff
					}
				}
			}
		}
		if t.Config.OnRetry != nil {
			t.Config.OnRetry(attempt+1, resp, err)
		}
		// Sleep under the caller's context, not the attempt's. An attempt that
		// timed out leaves its own context expired, which would skip every backoff.
		if err := sleep(req.Context(), backoff); err != nil {
			return nil, fmt.Errorf("(attempt: %d) sleeping: %w", attempt+1, err)
		}
	}
}

// attempt sends the request once under its own deadline. The deadline covers
// reading the response body, so on success the cancel func is handed to the
// body's Close and the caller must close the body to release it.
func (t *RetryTransport) attempt(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	cancel := context.CancelFunc(func() {})
	if t.Config.AttemptTimeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, t.Config.AttemptTimeout)
	}

	// Clone the request for retry (body needs special handling)
	reqCopy := req.Clone(ctx)

	// If there's a body, we need to handle GetBody for retries (due to it being a read-once construct)
	if req.Body != nil && req.GetBody != nil {
		body, err := req.GetBody()
		if err != nil {
			cancel()
			return nil, err
		}
		reqCopy.Body = body
	}

	resp, err := t.Base.RoundTrip(reqCopy)
	if err != nil {
		cancel()
		return nil, err
	}
	if resp.Body == nil {
		cancel()
		return resp, nil
	}

	resp.Body = &cancelOnCloseBody{ReadCloser: resp.Body, cancel: cancel}
	return resp, nil
}

// cancelOnCloseBody releases an attempt's context once the caller is done with
// the response body. The body is read after RoundTrip returns, so the attempt
// deadline has to outlive the call that created it.
type cancelOnCloseBody struct {
	io.ReadCloser
	cancel context.CancelFunc
}

func (b *cancelOnCloseBody) Close() error {
	err := b.ReadCloser.Close()
	b.cancel()
	return err
}

// shouldRetry reports whether the attempt can be replayed. A retryable failure
// is not enough on its own: for POST and PATCH the request may already have
// been applied, and replaying it would create a duplicate resource.
func (t *RetryTransport) shouldRetry(method string, resp *http.Response, err error) bool {
	if err != nil {
		// Nothing in the error says whether it happened before or after the
		// request reached the server, so only replay methods that are safe to
		// apply twice.
		return isIdempotentMethod(method) && isRetryableError(err)
	}

	if !t.isRetryableStatus(resp.StatusCode) {
		return false
	}

	// A rate-limited request was rejected before it was processed, so replaying
	// it is safe for any method. The other retryable statuses can come from a
	// proxy in front of an app that already committed the write.
	return resp.StatusCode == http.StatusTooManyRequests || isIdempotentMethod(method)
}

// isIdempotentMethod reports whether RFC 9110 guarantees that sending the
// request twice has the same effect as sending it once. POST and PATCH carry no
// such guarantee.
func isIdempotentMethod(method string) bool {
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodPut, http.MethodDelete, http.MethodOptions, http.MethodTrace:
		return true
	default:
		return false
	}
}

func (t *RetryTransport) isRetryableStatus(statusCode int) bool {
	for _, code := range t.Config.RetryableStatusCodes {
		if code == statusCode {
			return true
		}
	}
	return false
}

func (t *RetryTransport) calculateBackoff(attempt int) time.Duration {
	backoff := float64(t.Config.InitialBackoff) * math.Pow(t.Config.BackoffFactor, float64(attempt))
	backoff = math.Min(backoff, float64(t.Config.MaxBackoff))

	if t.Config.JitterFactor > 0 {
		jitter := backoff * t.Config.JitterFactor
		backoff = backoff - jitter + (rand.Float64() * 2 * jitter)
	}

	return time.Duration(backoff)
}

func isRetryableError(err error) bool {
	if err == nil {
		return false
	}

	// Retry on context deadline exceeded (timeout) but not on explicit cancellation
	if errors.Is(err, context.DeadlineExceeded) {
		return true
	}
	if errors.Is(err, context.Canceled) {
		return false
	}

	return true
}

func sleep(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func parseRetryAfterHeader(header string) (time.Duration, bool) {
	if header == "" {
		return 0, false
	}

	if seconds, err := strconv.Atoi(header); err == nil && seconds > 0 {
		return time.Duration(seconds) * time.Second, true
	}

	return 0, false
}

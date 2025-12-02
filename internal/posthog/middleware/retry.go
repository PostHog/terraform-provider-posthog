// Package middleware provides HTTP middleware components.
package middleware

import (
	"context"
	"errors"
	"fmt"
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
		RetryableStatusCodes: []int{429, 500, 502, 503, 504},
	}
}

func NoRetryConfig() RetryConfig {
	return RetryConfig{MaxRetries: 0}
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
	// If retries are disabled, just do the request
	if t.Config.MaxRetries == 0 {
		return t.Base.RoundTrip(req)
	}

	var lastResp *http.Response
	var lastErr error

	maxAttempts := t.Config.MaxRetries + 1 // +1 for the initial attempt
	for attempt := 0; attempt < maxAttempts; attempt++ {
		if err := req.Context().Err(); err != nil {
			return nil, fmt.Errorf("(attempt: %d) context error: %w", attempt+1, err)
		}

		// Clone the request for retry (body needs special handling)
		reqCopy := req.Clone(req.Context())

		// If there's a body, we need to handle GetBody for retries (due to it being a read-once construct)
		if req.Body != nil && req.GetBody != nil {
			body, err := req.GetBody()
			if err != nil {
				return nil, err
			}
			reqCopy.Body = body
		}

		resp, err := t.Base.RoundTrip(reqCopy)

		// Success case
		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return resp, nil
		}

		// Store last response/error for potential return
		lastResp = resp
		lastErr = err

		shouldRetry := false
		if err != nil {
			shouldRetry = isRetryableError(err)
		} else if t.isRetryableStatus(resp.StatusCode) {
			shouldRetry = true
		}

		// If not retryable or last attempt, return
		if !shouldRetry || attempt >= t.Config.MaxRetries-1 {
			if err != nil {
				return nil, err
			}
			return resp, nil
		}

		backoff := t.calculateBackoff(attempt)
		if resp != nil {
			// Close response body before retry to avoid resource leak
			if resp.Body != nil {
				resp.Body.Close()
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
		if err := sleep(req.Context(), backoff); err != nil {
			return nil, fmt.Errorf("(attempt: %d) sleeping: %w", attempt+1, err)
		}
	}

	// Should not reach here, but handle gracefully
	if lastErr != nil {
		return nil, lastErr
	}
	return lastResp, nil
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

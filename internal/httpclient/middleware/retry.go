// Package middleware provides HTTP middleware components.
package middleware

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand/v2"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
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
	// by timing out still has time left to be retried. Unset takes
	// DefaultAttemptTimeout. A whole request can therefore take
	// (MaxRetries+1) * AttemptTimeout plus the backoff between attempts.
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
		RetryableStatusCodes: defaultRetryableStatusCodes(),
	}
}

// defaultRetryableStatusCodes are rate limiting plus the gateway and server
// errors that usually clear on their own.
func defaultRetryableStatusCodes() []int {
	return []int{429, 500, 502, 503, 504}
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
	// Fill in the unset tunables. A zero AttemptTimeout would leave a request
	// with no bound at all now that the http.Client carries no Timeout, and a
	// zero backoff would retry with no spacing between attempts. JitterFactor
	// and MaxRetries are left alone because zero is a meaningful value for both.
	if config.AttemptTimeout <= 0 {
		config.AttemptTimeout = DefaultAttemptTimeout
	}
	if config.InitialBackoff <= 0 {
		config.InitialBackoff = DefaultInitialBackoff
	}
	if config.MaxBackoff <= 0 {
		config.MaxBackoff = DefaultMaxBackoff
	}
	if config.BackoffFactor <= 0 {
		config.BackoffFactor = DefaultBackoffFactor
	}
	if config.RetryableStatusCodes == nil {
		config.RetryableStatusCodes = defaultRetryableStatusCodes()
	}
	return &RetryTransport{
		Base:   base,
		Config: config,
	}
}

// ReplaceRetryTransport installs config on base, swapping out an already
// installed retry transport rather than stacking a second one on it. Stacked,
// the outer transport's per-attempt deadline would have to cover the inner
// transport's whole retry loop.
func ReplaceRetryTransport(base http.RoundTripper, config RetryConfig) *RetryTransport {
	if retry, ok := base.(*RetryTransport); ok {
		base = retry.Base
	}
	return NewRetryTransport(base, config)
}

func (t *RetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for attempt := 0; ; attempt++ {
		if err := req.Context().Err(); err != nil {
			return nil, fmt.Errorf("(attempt: %d) context error: %w", attempt+1, err)
		}

		resp, err := t.attempt(req)

		// Success case
		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return resp, nil
		}

		decision := t.decide(req, resp, err)
		exhausted := decision.retry && attempt >= t.Config.MaxRetries
		if !decision.retry || exhausted {
			message, reason := t.stopped(decision, exhausted)
			fields := attemptFields(req, attempt+1, resp, err)
			fields["reason"] = reason
			tflog.Debug(req.Context(), message, fields)
			if err != nil {
				return nil, fmt.Errorf("(attempt: %d) %w", attempt+1, err)
			}
			return resp, nil
		}

		backoff := t.nextBackoff(attempt, resp)
		if resp != nil && resp.Body != nil {
			// Close the body before retrying so the connection is not leaked.
			_ = resp.Body.Close()
		}
		if t.Config.OnRetry != nil {
			t.Config.OnRetry(attempt+1, resp, err)
		}
		fields := attemptFields(req, attempt+1, resp, err)
		fields["backoff"] = backoff.String()
		tflog.Debug(req.Context(), "retrying request", fields)
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
	ctx, cancel := context.WithTimeout(req.Context(), t.Config.AttemptTimeout)

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
	if err != nil && ctx.Err() != nil && req.Context().Err() == nil {
		// Say whose deadline fired. On its own the wrapped error reads as
		// "context deadline exceeded", which looks like the caller cancelled.
		err = fmt.Errorf("attempt timed out after %s: %w", t.Config.AttemptTimeout, err)
	}
	if err != nil || resp.Body == nil {
		cancel()
		return resp, err
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

// retryDecision is the outcome of decide: whether to send the attempt again,
// and why.
type retryDecision struct {
	retry bool
	// reason is written to the debug log, so an operator can tell "ran out of
	// attempts" apart from "never tried again, because replaying this method
	// could have created the resource twice".
	reason string
}

// decide reports whether the failed attempt can be sent again, and why.
func (t *RetryTransport) decide(req *http.Request, resp *http.Response, err error) retryDecision {
	// The caller's own deadline or cancellation ended this, so nothing about the
	// failure or the method is what stopped the retry.
	if err != nil && req.Context().Err() != nil {
		return retryDecision{reason: "caller's context ended"}
	}

	retryable := isRetryableError(err)
	if err == nil {
		retryable = t.isRetryableStatus(resp.StatusCode)
	}
	if !retryable {
		return retryDecision{reason: "failure is not retryable"}
	}

	// A rate-limited request was rejected before it was processed, so replaying
	// it is safe for any method. Every other retryable failure can come from a
	// proxy in front of an app that already committed the write, and nothing in
	// the response or the error says which happened.
	if err == nil && resp.StatusCode == http.StatusTooManyRequests {
		return retryDecision{retry: true, reason: "rate limited"}
	}
	if !safeToReplay(req, err) {
		return retryDecision{reason: req.Method + " is not safe to replay"}
	}
	return retryDecision{retry: true, reason: "retryable failure"}
}

// safeToReplay reports whether sending the request a second time cannot apply
// it twice.
//
// net/http answers the same question in Request.isReplayable, but it is
// unexported and reachable only from inside the transport, which is why the
// rule is restated here. Two deliberate differences from it: PUT and DELETE are
// included, because RFC 9110 defines them as idempotent and this retry is
// explicit and logged rather than the transport's invisible one; and the
// "nothing was written to the wire" case is left to the transport, which
// already replays any method when it can prove that (see failedBeforeSending).
func safeToReplay(req *http.Request, err error) bool {
	if isIdempotentMethod(req.Method) || failedBeforeSending(err) {
		return true
	}
	// Honour the same two headers net/http does, for the same reason: the caller
	// is saying the server will collapse a duplicate.
	return len(req.Header.Values("Idempotency-Key")) > 0 ||
		len(req.Header.Values("X-Idempotency-Key")) > 0
}

// failedBeforeSending reports whether the error proves the request never
// reached the server. A dial that never connected wrote no bytes, so even a
// create is safe to send again. Every other error leaves it unknowable.
//
// This is the case net/http's own transport refuses: shouldRetryRequest bails
// on a connection that was not reused, so a failed dial is never replayed below
// this layer. The reverse also holds, so the two do not overlap: when a pooled
// connection turns out to be dead before any bytes are written, the transport
// replays the request itself, including a POST, as long as GetBody is set.
func failedBeforeSending(err error) bool {
	var opErr *net.OpError
	return errors.As(err, &opErr) && opErr.Op == "dial"
}

// stopped describes why the loop ended, for the debug log. A failure that was
// never retryable is the normal path for a 404 during drift detection, so it
// does not get the same wording as running out of attempts.
func (t *RetryTransport) stopped(decision retryDecision, exhausted bool) (message, reason string) {
	switch {
	case exhausted && t.Config.MaxRetries == 0:
		return "not retrying request", "retries are disabled"
	case exhausted:
		return "giving up on request", "retries exhausted"
	default:
		return "not retrying request", decision.reason
	}
}

// nextBackoff is how long to wait before trying again. A Retry-After on a
// rate-limited response outranks the computed backoff, still bounded by
// MaxBackoff so one hostile header cannot stall an apply.
func (t *RetryTransport) nextBackoff(attempt int, resp *http.Response) time.Duration {
	backoff := t.calculateBackoff(attempt)
	if resp == nil || resp.StatusCode != http.StatusTooManyRequests {
		return backoff
	}

	retryAfter, ok := parseRetryAfterHeader(resp.Header.Get("Retry-After"))
	if !ok {
		return backoff
	}
	if retryAfter > backoff {
		backoff = retryAfter
	}
	return min(backoff, t.Config.MaxBackoff)
}

// attemptFields describes an attempt's outcome for the debug log.
func attemptFields(req *http.Request, attempt int, resp *http.Response, err error) map[string]any {
	fields := map[string]any{"method": req.Method, "path": req.URL.Path, "attempt": attempt}
	if resp != nil {
		fields["status"] = resp.StatusCode
	}
	if err != nil {
		fields["error"] = err.Error()
	}
	return fields
}

// isIdempotentMethod reports whether RFC 9110 guarantees that sending the
// request twice has the same effect as sending it once. POST and PATCH carry no
// such guarantee. This is a wider set than net/http's isReplayable, which stops
// at GET, HEAD, OPTIONS and TRACE.
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

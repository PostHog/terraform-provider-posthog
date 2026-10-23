// Package middleware provides HTTP middleware components.
package middleware

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand/v2"
	"net"
	"net/http"
	"strconv"
	"strings"
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
	// PostHog's Retry-After is the time until the API key's sliding window
	// frees a request: at most a minute under its per-minute limits. A longer
	// ask comes from an hourly limit, which no apply should sit through.
	DefaultRateLimitMaxWait = 60 * time.Second
	// A plan that refreshes more resources than the per-minute limit allows
	// waits out several windows, but not without end.
	DefaultRateLimitTotalWait = 5 * time.Minute
	// maxRateLimitJitter caps the jitter added to a Retry-After wait.
	maxRateLimitJitter = 2 * time.Second
	// minRateLimitWait is the shortest wait on a 429. PostHog writes
	// Retry-After in whole seconds rounded down, so a slot freeing up within
	// the second arrives as "0".
	minRateLimitWait = time.Second
)

// RateLimitWaitedHeader and RateLimitStopReasonHeader are set on a 429 the
// transport stops retrying: how long it slept on 429s, and why it stopped.
// RateLimitStop reads them back. They ride on the response rather than a typed
// error because a RoundTripper that returns an error loses the response, and
// with it PostHog's message.
const (
	RateLimitWaitedHeader     = "X-Posthog-Provider-Rate-Limit-Waited"
	RateLimitStopReasonHeader = "X-Posthog-Provider-Rate-Limit-Stop-Reason"
)

// RateLimitStop reports how long the transport slept on 429s before giving up
// on resp, and why it gave up. ok is false when resp is not a 429 the transport
// stopped retrying.
func RateLimitStop(resp *http.Response) (waited time.Duration, reason string, ok bool) {
	reason = resp.Header.Get(RateLimitStopReasonHeader)
	waited, err := time.ParseDuration(resp.Header.Get(RateLimitWaitedHeader))
	if reason == "" || err != nil {
		return 0, "", false
	}
	return waited, reason, true
}

type RetryConfig struct {
	// MaxRetries bounds retries of server errors, timeouts, and 429s that carry
	// no usable Retry-After. 0 turns off every retry, including waits on 429s
	// that do carry one, which are otherwise bounded by RateLimitTotalWait.
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
	// (MaxRetries+1) * AttemptTimeout plus the backoff between attempts, plus
	// up to RateLimitTotalWait when PostHog rate-limits it.
	AttemptTimeout time.Duration
	// RateLimitMaxWait is the longest Retry-After waited out on a 429. One up
	// to it is waited in full, even past MaxBackoff, plus up to 2 seconds of
	// jitter; a longer one returns the 429 to the caller straight away. Unset
	// takes DefaultRateLimitMaxWait.
	RateLimitMaxWait time.Duration
	// RateLimitTotalWait bounds the time one request spends rate-limited, from
	// sending the first attempt that got a 429 with a Retry-After: the waits and
	// the throttled attempts both count. Unset takes DefaultRateLimitTotalWait.
	RateLimitTotalWait time.Duration
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
		RateLimitMaxWait:     DefaultRateLimitMaxWait,
		RateLimitTotalWait:   DefaultRateLimitTotalWait,
		RetryableStatusCodes: defaultRetryableStatusCodes(),
	}
}

// defaultRetryableStatusCodes are rate limiting plus the gateway and server
// errors that usually clear on their own.
func defaultRetryableStatusCodes() []int {
	return []int{429, 500, 502, 503, 504}
}

// NoRetryConfig turns off every retry, rate-limited ones included.
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
	if config.RateLimitMaxWait <= 0 {
		config.RateLimitMaxWait = DefaultRateLimitMaxWait
	}
	if config.RateLimitTotalWait <= 0 {
		config.RateLimitTotalWait = DefaultRateLimitTotalWait
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
	// retries counts the failures MaxRetries bounds. 429s that carry a
	// Retry-After are bounded by RateLimitTotalWait instead, measured from when
	// the first of them was sent, because PostHog rejected them unprocessed.
	var retries int
	var rateLimitedSince time.Time
	var rateLimitSlept time.Duration
	for attempt := 0; ; attempt++ {
		if err := req.Context().Err(); err != nil {
			return nil, fmt.Errorf("(attempt: %d) context error: %w", attempt+1, err)
		}

		sent := time.Now()
		resp, err := t.attempt(req)

		// Success case
		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return resp, nil
		}

		decision := t.decide(req, resp, err)
		if !decision.retry {
			return t.stop(req, attempt+1, resp, err, decision, false)
		}

		var backoff time.Duration
		if decision.rateLimited {
			if rateLimitedSince.IsZero() {
				rateLimitedSince = sent
			}
			var stopReason string
			backoff, stopReason = t.rateLimitBackoff(decision.retryAfter, time.Since(rateLimitedSince))
			if stopReason != "" {
				return t.stopRateLimited(req, attempt+1, resp, stopReason, rateLimitSlept)
			}
		} else {
			if retries >= t.Config.MaxRetries {
				if err == nil && resp.StatusCode == http.StatusTooManyRequests {
					return t.stopRateLimited(req, attempt+1, resp, t.noRetryAfterStopReason(), rateLimitSlept)
				}
				return t.stop(req, attempt+1, resp, err, decision, true)
			}
			backoff = t.calculateBackoff(retries)
		}

		// No body to close before looping: attempt already drained and closed it.
		t.logRetry(req, attempt+1, resp, err, decision, backoff)
		// Sleep under the caller's context, not the attempt's. An attempt that
		// timed out leaves its own context expired, which would skip every backoff.
		if err := sleep(req.Context(), backoff); err != nil {
			return nil, fmt.Errorf("(attempt: %d) sleeping: %w", attempt+1, err)
		}
		if err == nil && resp.StatusCode == http.StatusTooManyRequests {
			rateLimitSlept += backoff
		}
		if !decision.rateLimited {
			retries++
		}
	}
}

// stop ends the loop with the last attempt's outcome, for every failure but a
// 429 that carried a Retry-After.
func (t *RetryTransport) stop(req *http.Request, attempt int, resp *http.Response, err error, decision retryDecision, exhausted bool) (*http.Response, error) {
	message, reason := t.stopped(decision, exhausted)
	fields := attemptFields(req, attempt, resp, err)
	fields["reason"] = reason
	tflog.Debug(req.Context(), message, fields)
	if err != nil {
		if exhausted {
			return nil, fmt.Errorf("(attempt: %d) %w", attempt, err)
		}
		// Name why it was not sent again. Without this the operator sees
		// only the timeout and cannot tell that re-running may duplicate
		// a resource this run already created.
		return nil, fmt.Errorf("(attempt: %d) not retried, %s: %w", attempt, reason, err)
	}
	return resp, nil
}

// stopRateLimited returns a 429 the transport will not wait on again, with the
// time slept and the reason written onto it for RateLimitStop.
func (t *RetryTransport) stopRateLimited(req *http.Request, attempt int, resp *http.Response, reason string, slept time.Duration) (*http.Response, error) {
	waited := slept.Round(time.Second).String()
	fields := attemptFields(req, attempt, resp, nil)
	fields["reason"] = reason
	fields["rate_limit_waited"] = waited
	tflog.Warn(req.Context(), "giving up on rate-limited request", fields)
	resp.Header.Set(RateLimitWaitedHeader, waited)
	resp.Header.Set(RateLimitStopReasonHeader, reason)
	return resp, nil
}

func (t *RetryTransport) logRetry(req *http.Request, attempt int, resp *http.Response, err error, decision retryDecision, backoff time.Duration) {
	if t.Config.OnRetry != nil {
		t.Config.OnRetry(attempt, resp, err)
	}
	fields := attemptFields(req, attempt, resp, err)
	fields["backoff"] = backoff.String()
	if decision.rateLimited {
		// A wait of up to a minute, repeated across parallel requests, looks
		// like a hung plan unless it is visible without TF_LOG=DEBUG.
		fields["retry_after"] = decision.retryAfter.String()
		tflog.Warn(req.Context(), "PostHog rate limit reached for this API key; waiting before retrying", fields)
		return
	}
	tflog.Debug(req.Context(), "retrying request", fields)
}

// noRetryAfterStopReason explains giving up on a 429 that carried no usable
// Retry-After, which gets the ordinary MaxRetries bound.
func (t *RetryTransport) noRetryAfterStopReason() string {
	if t.Config.MaxRetries == 0 {
		return "retries are disabled"
	}
	return fmt.Sprintf("the response had no usable Retry-After and %d retries ran out", t.Config.MaxRetries)
}

// rateLimitBackoff is how long to wait before sending a request again after a
// 429 that asked for retryAfter, given the time since the request was first
// rate-limited. A non-empty stopReason means not to send it again.
func (t *RetryTransport) rateLimitBackoff(retryAfter, rateLimitedFor time.Duration) (backoff time.Duration, stopReason string) {
	// NoRetryConfig has to mean no retries at all.
	if t.Config.MaxRetries == 0 {
		return 0, "retries are disabled"
	}
	if retryAfter > t.Config.RateLimitMaxWait {
		return 0, fmt.Sprintf("PostHog's Retry-After of %s is longer than the %s the provider waits for one 429", retryAfter, t.Config.RateLimitMaxWait)
	}
	// Parallel requests rejected together get the same Retry-After, and the
	// window frees one slot at a time. Spread them out, never sooner than asked.
	wait := max(retryAfter, minRateLimitWait)
	jitter := min(time.Duration(float64(wait)*t.Config.JitterFactor), maxRateLimitJitter)
	backoff = wait + time.Duration(rand.Float64()*float64(jitter))
	if rateLimitedFor+backoff > t.Config.RateLimitTotalWait {
		return 0, fmt.Sprintf("waiting again would pass the %s rate-limit budget", t.Config.RateLimitTotalWait)
	}
	return backoff, ""
}

// attempt sends the request once under its own deadline, response body
// included, and returns the body already read into memory.
func (t *RetryTransport) attempt(req *http.Request) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(req.Context(), t.Config.AttemptTimeout)
	defer cancel()

	// Clone the request for retry (body needs special handling)
	reqCopy := req.Clone(ctx)

	// If there's a body, we need to handle GetBody for retries (due to it being a read-once construct)
	if req.Body != nil && req.GetBody != nil {
		body, err := req.GetBody()
		if err != nil {
			return nil, err
		}
		reqCopy.Body = body
	}

	resp, err := t.Base.RoundTrip(reqCopy)
	if err != nil {
		return nil, t.nameTheDeadline(ctx, req, err)
	}

	// Read the body here rather than leaving it to the caller. RoundTrip returns
	// as soon as the headers arrive, so a body that stalls would fail outside the
	// retry loop, where it can be neither retried nor explained. Every response
	// this client handles is read whole anyway, so nothing is buffered that the
	// caller was not about to buffer itself.
	if err := bufferBody(resp); err != nil {
		return nil, t.nameTheDeadline(ctx, req, err)
	}
	return resp, nil
}

// nameTheDeadline marks an error caused by this attempt's own timeout. Left
// alone it reads as "context deadline exceeded", which looks to the caller like
// their own cancellation.
func (t *RetryTransport) nameTheDeadline(ctx context.Context, req *http.Request, err error) error {
	if ctx.Err() != nil && req.Context().Err() == nil {
		return fmt.Errorf("attempt timed out after %s: %w", t.Config.AttemptTimeout, err)
	}
	return err
}

// bufferBody replaces the response body with its own contents, so that reading
// it happens under the attempt's deadline and the connection goes back to the
// pool straight away.
func bufferBody(resp *http.Response) error {
	if resp.Body == nil {
		return nil
	}

	streamed := resp.Body
	defer func() { _ = streamed.Close() }()

	body, err := io.ReadAll(streamed)
	if err != nil {
		return err
	}
	resp.Body = io.NopCloser(bytes.NewReader(body))
	return nil
}

// retryDecision is the outcome of decide: whether to send the attempt again,
// and why.
type retryDecision struct {
	retry bool
	// mayHaveApplied marks the case where the request was not sent again because
	// a replay could apply it twice. The write may still have landed, and the
	// caller has to know that before running again.
	mayHaveApplied bool
	// rateLimited marks a 429 with a usable Retry-After, kept in retryAfter.
	// Its retries wait as long as PostHog asks and are bounded by the
	// rate-limit budget instead of MaxRetries.
	rateLimited bool
	retryAfter  time.Duration
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
		if retryAfter, ok := parseRetryAfterHeader(resp.Header.Get("Retry-After")); ok {
			return retryDecision{retry: true, rateLimited: true, retryAfter: retryAfter, reason: "rate limited"}
		}
		// Without a Retry-After, possibly from a proxy in front of PostHog,
		// there is no ask to honour, so it gets the ordinary bounded retries.
		return retryDecision{retry: true, reason: "rate limited without Retry-After"}
	}
	if !safeToReplay(req, err) {
		return retryDecision{
			reason:         req.Method + " may already have been applied",
			mayHaveApplied: true,
		}
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

// maxRetryAfter clamps a parsed Retry-After. It is far past any wait the
// transport accepts, so a clamped value is still refused, and it keeps the
// conversion to a Duration from overflowing: a huge header would otherwise wrap
// to a negative wait that passes every cap and is never added to a budget.
const maxRetryAfter = 24 * time.Hour

func parseRetryAfterHeader(header string) (time.Duration, bool) {
	if header == "" {
		return 0, false
	}

	seconds, err := strconv.ParseInt(header, 10, 64)
	if errors.Is(err, strconv.ErrRange) && !strings.HasPrefix(header, "-") {
		// Too long to hold, so far too long to wait.
		return maxRetryAfter, true
	}
	if err != nil || seconds < 0 {
		return 0, false
	}
	if seconds > int64(maxRetryAfter/time.Second) {
		return maxRetryAfter, true
	}
	return time.Duration(seconds) * time.Second, true
}

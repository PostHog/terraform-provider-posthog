package httpclient

import (
	"net/http"

	"github.com/posthog/terraform-provider/internal/httpclient/middleware"
)

// ClientOption is a functional option for configuring the client.
type ClientOption func(*http.Client)

// WithRetryConfig sets a custom retry configuration.
func WithRetryConfig(config middleware.RetryConfig) ClientOption {
	return func(c *http.Client) {
		c.Transport = middleware.NewRetryTransport(unwrapRetryTransport(c.Transport), config)
	}
}

// WithNoRetry disables retry logic. The request still gets the retry transport's
// per-attempt deadline, which is the only bound left now that the client itself
// carries no Timeout.
func WithNoRetry() ClientOption {
	return WithRetryConfig(middleware.NoRetryConfig())
}

// unwrapRetryTransport returns what a retry transport wraps, so the options
// above replace an installed one instead of stacking a second retry loop on it.
// A stacked pair nests deadlines: the outer transport's per-attempt budget would
// have to cover the inner one's whole loop.
func unwrapRetryTransport(rt http.RoundTripper) http.RoundTripper {
	if retry, ok := rt.(*middleware.RetryTransport); ok {
		return retry.Base
	}
	return rt
}

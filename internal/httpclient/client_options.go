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
		c.Transport = middleware.ReplaceRetryTransport(c.Transport, config)
	}
}

// WithNoRetry disables retry logic. The request still gets the retry transport's
// per-attempt deadline, which is the only bound left now that the client itself
// carries no Timeout.
func WithNoRetry() ClientOption {
	return WithRetryConfig(middleware.NoRetryConfig())
}

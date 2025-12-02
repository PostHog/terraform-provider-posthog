package http

import (
	"net/http"

	"github.com/posthog/terraform-provider/internal/http/middleware"
)

// ClientOption is a functional option for configuring the client.
type ClientOption func(*http.Client)

// WithRetryConfig sets a custom retry configuration.
func WithRetryConfig(config middleware.RetryConfig) ClientOption {
	return func(c *http.Client) {
		c.Transport = middleware.NewRetryTransport(c.Transport, config)
	}
}

// WithNoRetry disables retry logic.
func WithNoRetry() ClientOption {
	return func(c *http.Client) {
		c.Transport = middleware.NewRetryTransport(c.Transport, middleware.NoRetryConfig())
	}
}

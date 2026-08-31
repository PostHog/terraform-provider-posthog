package httpclient

import (
	"net/http"
	"testing"

	"github.com/posthog/terraform-provider/internal/httpclient/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultClient_LeavesTimeoutToTheRetryTransport(t *testing.T) {
	client := NewDefaultClient("https://us.posthog.com", "key", "test")

	// An http.Client.Timeout would bound the retry loop and its backoff sleeps,
	// so a request that failed by timing out could never be retried.
	assert.Zero(t, client.httpClient.Timeout)

	transport, ok := client.httpClient.Transport.(*middleware.RetryTransport)
	require.True(t, ok, "expected the retry transport to be installed")
	assert.Equal(t, middleware.DefaultAttemptTimeout, transport.Config.AttemptTimeout)
	assert.Equal(t, middleware.DefaultMaxRetries, transport.Config.MaxRetries)

	_, ok = transport.Base.(*http.Transport)
	assert.True(t, ok, "expected the retry transport to wrap the pooled http.Transport")
}

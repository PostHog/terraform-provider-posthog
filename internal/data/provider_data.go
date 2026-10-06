package data

import (
	"github.com/posthog/terraform-provider/internal/httpclient"
)

// ProviderData passes the configured client and scope defaults to resources.
type ProviderData struct {
	Client                httpclient.PosthogClient
	DefaultProjectID      string
	DefaultOrganizationID string
}

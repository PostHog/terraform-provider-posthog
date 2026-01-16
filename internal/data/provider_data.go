package data

import (
	"github.com/posthog/terraform-provider/internal/httpclient"
)

// ProviderData passes configured client to resources.
type ProviderData struct {
	Client           httpclient.PosthogClient
	DefaultProjectID string
}

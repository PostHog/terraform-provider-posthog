package data

import (
	"github.com/posthog/terraform-provider/internal/http"
)

// ProviderData passes configured client to resources.
type ProviderData struct {
	Client http.PosthogClient
}

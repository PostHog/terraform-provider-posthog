package data

import (
	posthogapi "github.com/posthog/terraform-provider/internal/posthog/swagger"
)

// ProviderData passes configured client to resources.
type ProviderData struct {
	SwaggerClient *posthogapi.APIClient
	ProjectID     string
}

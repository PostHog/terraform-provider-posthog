package data

import (
	"github.com/posthog/terraform-provider/internal/posthog"
	posthogapi "github.com/posthog/terraform-provider/internal/posthog/swagger"
)

// ProviderData passes configured client to resources.
type ProviderData struct {
	Client        posthog.Client
	SwaggerClient *posthogapi.APIClient
	ProjectID     string
}

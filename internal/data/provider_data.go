package data

import "github.com/posthog/terraform-provider/internal/posthog"

// ProviderData passes configured client to resources.
type ProviderData struct {
	Client    posthog.Client
	ProjectID string
}

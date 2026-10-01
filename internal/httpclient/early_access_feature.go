package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
)

// EarlyAccessFeature is a feature users can opt into from the PostHog in-app
// "Feature previews" list. Each one is backed by a feature flag.
type EarlyAccessFeature struct {
	ID               string                         `json:"id"`
	FeatureFlag      *EarlyAccessFeatureFeatureFlag `json:"feature_flag"`
	Name             string                         `json:"name"`
	Description      *string                        `json:"description"`
	Stage            string                         `json:"stage"`
	DocumentationURL *string                        `json:"documentation_url"`
	Payload          json.RawMessage                `json:"payload"`
	CreatedAt        string                         `json:"created_at"`
}

// EarlyAccessFeatureFeatureFlag is the linked flag as nested in the response.
// The API returns more of the flag; only its identity is needed here.
type EarlyAccessFeatureFeatureFlag struct {
	ID  int64  `json:"id"`
	Key string `json:"key"`
}

type EarlyAccessFeatureRequest struct {
	Name             string          `json:"name"`
	Stage            string          `json:"stage"`
	Description      *string         `json:"description,omitempty"`
	DocumentationURL *string         `json:"documentation_url,omitempty"`
	Payload          json.RawMessage `json:"payload,omitempty"`
	// Create-only: PostHog ignores it on update. When omitted on create,
	// PostHog creates a flag keyed by the slugified name.
	FeatureFlagID *int64 `json:"feature_flag_id,omitempty"`
}

func (c *PosthogClient) CreateEarlyAccessFeature(ctx context.Context, projectID string, input EarlyAccessFeatureRequest) (EarlyAccessFeature, error) {
	path := fmt.Sprintf("/api/projects/%s/early_access_feature/", projectID)
	result, _, err := doPost[EarlyAccessFeature](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetEarlyAccessFeature(ctx context.Context, projectID, id string) (EarlyAccessFeature, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/early_access_feature/%s/", projectID, id)
	return doGet[EarlyAccessFeature](c, ctx, path)
}

func (c *PosthogClient) UpdateEarlyAccessFeature(ctx context.Context, projectID, id string, input EarlyAccessFeatureRequest) (EarlyAccessFeature, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/early_access_feature/%s/", projectID, id)
	return doPatch[EarlyAccessFeature](c, ctx, path, input)
}

// DeleteEarlyAccessFeature hard-deletes the feature. PostHog turns off opt-in on
// the linked flag but keeps the flag itself.
func (c *PosthogClient) DeleteEarlyAccessFeature(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/early_access_feature/%s/", projectID, id)
	return doDelete(c, ctx, path)
}

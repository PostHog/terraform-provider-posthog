package httpclient

import (
	"context"
	"fmt"
)

type FeatureFlag struct {
	ID                int64                  `json:"id"`
	Key               string                 `json:"key"`
	Name              *string                `json:"name,omitempty"`
	Active            *bool                  `json:"active,omitempty"`
	Filters           map[string]interface{} `json:"filters,omitempty"`
	RolloutPercentage *int32                 `json:"rollout_percentage,omitempty"`
	Tags              []string               `json:"tags,omitempty"`
	Deleted           *bool                  `json:"deleted,omitempty"`
}

type FeatureFlagRequest struct {
	Key     string                 `json:"key"`
	Name    *string                `json:"name,omitempty"`
	Active  *bool                  `json:"active,omitempty"`
	Filters map[string]interface{} `json:"filters,omitempty"`
	Tags    []string               `json:"tags,omitempty"`
	Deleted *bool                  `json:"deleted,omitempty"`
}

func (c *PosthogClient) CreateFeatureFlag(ctx context.Context, input FeatureFlagRequest) (FeatureFlag, error) {
	path := fmt.Sprintf("/api/projects/%s/feature_flags/", c.projectID)
	result, _, err := doPost[FeatureFlag](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetFeatureFlag(ctx context.Context, id string) (FeatureFlag, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/feature_flags/%s/", c.projectID, id)
	return doGet[FeatureFlag](c, ctx, path)
}

func (c *PosthogClient) UpdateFeatureFlag(ctx context.Context, id string, input FeatureFlagRequest) (FeatureFlag, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/feature_flags/%s/", c.projectID, id)
	return doPatch[FeatureFlag](c, ctx, path, input)
}

func (c *PosthogClient) DeleteFeatureFlag(ctx context.Context, id string) (HTTPStatusCode, error) {
	existing, _, err := c.GetFeatureFlag(ctx, id)
	if err != nil {
		return 0, fmt.Errorf("failed to get feature flag for deletion: %w", err)
	}

	deleted := true
	input := FeatureFlagRequest{
		Key:     existing.Key,
		Deleted: &deleted,
	}
	_, statusCode, err := c.UpdateFeatureFlag(ctx, id, input)
	return statusCode, err
}

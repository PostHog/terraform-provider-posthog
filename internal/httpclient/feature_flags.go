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

func (c *PosthogClient) GetFeatureFlag(ctx context.Context, id int64) (FeatureFlag, error) {
	path := fmt.Sprintf("/api/projects/%s/feature_flags/%d/", c.projectID, id)
	result, _, err := doGet[FeatureFlag](c, ctx, path)
	return result, err
}

func (c *PosthogClient) UpdateFeatureFlag(ctx context.Context, id int64, input FeatureFlagRequest) (FeatureFlag, error) {
	path := fmt.Sprintf("/api/projects/%s/feature_flags/%d/", c.projectID, id)
	result, _, err := doPatch[FeatureFlag](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) DeleteFeatureFlag(ctx context.Context, id int64) error {
	existing, err := c.GetFeatureFlag(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get feature flag for deletion: %w", err)
	}

	deleted := true
	input := FeatureFlagRequest{
		Key:     existing.Key,
		Deleted: &deleted,
	}
	_, err = c.UpdateFeatureFlag(ctx, id, input)
	return err
}

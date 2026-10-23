package httpclient

import (
	"context"
	"fmt"
)

type FeatureFlag struct {
	ID                         int64                  `json:"id"`
	Key                        string                 `json:"key"`
	Name                       *string                `json:"name,omitempty"`
	Active                     *bool                  `json:"active,omitempty"`
	Filters                    map[string]interface{} `json:"filters,omitempty"`
	RolloutPercentage          *int32                 `json:"rollout_percentage,omitempty"`
	Tags                       []string               `json:"tags,omitempty"`
	Deleted                    *bool                  `json:"deleted,omitempty"`
	Archived                   *bool                  `json:"archived,omitempty"`
	EnsureExperienceContinuity *bool                  `json:"ensure_experience_continuity,omitempty"`
}

type FeatureFlagRequest struct {
	Key                        string                 `json:"key"`
	Name                       *string                `json:"name,omitempty"`
	Active                     *bool                  `json:"active,omitempty"`
	Filters                    map[string]interface{} `json:"filters,omitempty"`
	Tags                       []string               `json:"tags,omitempty"`
	Deleted                    *bool                  `json:"deleted,omitempty"`
	EnsureExperienceContinuity *bool                  `json:"ensure_experience_continuity,omitempty"`
	// Legacy write-only create option. Current PostHog ignores it, but sending
	// false still suppresses automatic dashboards on older PostHog versions.
	// The resource rejects true before calling the API.
	ShouldCreateUsageDashboard *bool `json:"_should_create_usage_dashboard,omitempty"`
}

func (c *PosthogClient) CreateFeatureFlag(ctx context.Context, projectID string, input FeatureFlagRequest) (FeatureFlag, error) {
	path := fmt.Sprintf("/api/projects/%s/feature_flags/", projectID)
	result, _, err := doPost[FeatureFlag](c, ctx, path, input)
	return result, err
}

// ListFeatureFlags returns every flag in the project that has not been deleted,
// archived ones included. The list endpoint hides archived flags unless the
// request passes `archived`, and that parameter selects one state or the other,
// so it takes one pass per state. A PostHog version that predates archiving
// ignores the parameter and returns every flag in both passes, so flags are
// de-duplicated by ID.
//
// Pages are offsets into a list ordered newest first, so a flag deleted or
// archived while the pages are read can shift another past a page boundary and
// out of the result. A flag missing that way surfaces when creating its key
// again fails on apply, not as a wrong plan that succeeds.
func (c *PosthogClient) ListFeatureFlags(ctx context.Context, projectID string) ([]FeatureFlag, error) {
	var all []FeatureFlag
	seen := make(map[int64]bool)
	for _, archived := range []string{"false", "true"} {
		path := fmt.Sprintf("/api/projects/%s/feature_flags/?archived=%s", projectID, archived)
		flags, err := listAll[FeatureFlag](c, ctx, path)
		if err != nil {
			return nil, fmt.Errorf("listing flags with archived=%s: %w", archived, err)
		}
		for _, f := range flags {
			if !seen[f.ID] {
				seen[f.ID] = true
				all = append(all, f)
			}
		}
	}
	return all, nil
}

func (c *PosthogClient) GetFeatureFlag(ctx context.Context, projectID, id string) (FeatureFlag, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/feature_flags/%s/", projectID, id)
	return doGet[FeatureFlag](c, ctx, path)
}

func (c *PosthogClient) UpdateFeatureFlag(ctx context.Context, projectID, id string, input FeatureFlagRequest) (FeatureFlag, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/feature_flags/%s/", projectID, id)
	return doPatch[FeatureFlag](c, ctx, path, input)
}

func (c *PosthogClient) DeleteFeatureFlag(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	existing, _, err := c.GetFeatureFlag(ctx, projectID, id)
	if err != nil {
		return 0, fmt.Errorf("failed to get feature flag for deletion: %w", err)
	}

	deleted := true
	input := FeatureFlagRequest{
		Key:     existing.Key,
		Deleted: &deleted,
	}
	_, statusCode, err := c.UpdateFeatureFlag(ctx, projectID, id, input)
	return statusCode, err
}

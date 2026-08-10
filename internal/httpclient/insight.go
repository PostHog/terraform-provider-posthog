package httpclient

import (
	"context"
	"fmt"
)

type Insight struct {
	ID             int64                  `json:"id"`
	ShortID        *string                `json:"short_id,omitempty"`
	Name           *string                `json:"name,omitempty"`
	DerivedName    *string                `json:"derived_name,omitempty"`
	Description    *string                `json:"description,omitempty"`
	Query          map[string]interface{} `json:"query,omitempty"`
	Tags           []string               `json:"tags,omitempty"`
	Dashboards     []int32                `json:"dashboards,omitempty"`
	CreateInFolder *string                `json:"_create_in_folder,omitempty"`
	Favorited      *bool                  `json:"favorited,omitempty"`
	Deleted        *bool                  `json:"deleted,omitempty"`
}

// InsightQueryServerFields are query-internal fields PostHog injects on the
// insight response that are not part of the user-authored query. Both the
// posthog_insight resource and data source strip them (via util.StripFields)
// so the normalized query_json stays stable across refreshes and import.
var InsightQueryServerFields = map[string]struct{}{
	"version":   {},
	"result":    {},
	"hogql":     {},
	"is_cached": {},
}

type InsightRequest struct {
	Name           *string                `json:"name,omitempty"`
	DerivedName    *string                `json:"derived_name,omitempty"`
	Description    *string                `json:"description,omitempty"`
	Query          map[string]interface{} `json:"query,omitempty"`
	Tags           []string               `json:"tags,omitempty"`
	Dashboards     []int32                `json:"dashboards,omitempty"`
	CreateInFolder *string                `json:"_create_in_folder,omitempty"`
	Deleted        *bool                  `json:"deleted,omitempty"`
}

func (c *PosthogClient) CreateInsight(ctx context.Context, projectID string, input InsightRequest) (Insight, error) {
	path := fmt.Sprintf("/api/projects/%s/insights/", projectID)
	result, _, err := doPost[Insight](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetInsight(ctx context.Context, projectID, id string) (Insight, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/insights/%s/", projectID, id)
	return doGet[Insight](c, ctx, path)
}

func (c *PosthogClient) UpdateInsight(ctx context.Context, projectID, id string, input InsightRequest) (Insight, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/insights/%s/", projectID, id)
	return doPatch[Insight](c, ctx, path, input)
}

func (c *PosthogClient) DeleteInsight(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	deleted := true
	input := InsightRequest{Deleted: &deleted}
	_, statusCode, err := c.UpdateInsight(ctx, projectID, id, input)
	return statusCode, err
}

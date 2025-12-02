package httpclient

import (
	"context"
	"fmt"
)

type Insight struct {
	ID             int64                  `json:"id"`
	Name           *string                `json:"name,omitempty"`
	DerivedName    *string                `json:"derived_name,omitempty"`
	Description    *string                `json:"description,omitempty"`
	Query          map[string]interface{} `json:"query,omitempty"`
	Tags           []string               `json:"tags,omitempty"`
	Dashboards     []int32                `json:"dashboards,omitempty"`
	CreateInFolder *string                `json:"_create_in_folder,omitempty"`
	Deleted        *bool                  `json:"deleted,omitempty"`
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

func (c *PosthogClient) CreateInsight(ctx context.Context, input InsightRequest) (Insight, error) {
	path := fmt.Sprintf("/api/projects/%s/insights/", c.projectID)
	result, _, err := doPost[Insight](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetInsight(ctx context.Context, id int64) (Insight, error) {
	path := fmt.Sprintf("/api/projects/%s/insights/%d/", c.projectID, id)
	result, _, err := doGet[Insight](c, ctx, path)
	return result, err
}

func (c *PosthogClient) UpdateInsight(ctx context.Context, id int64, input InsightRequest) (Insight, error) {
	path := fmt.Sprintf("/api/projects/%s/insights/%d/", c.projectID, id)
	result, _, err := doPatch[Insight](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) DeleteInsight(ctx context.Context, id int64) error {
	deleted := true
	input := InsightRequest{Deleted: &deleted}
	_, err := c.UpdateInsight(ctx, id, input)
	return err
}

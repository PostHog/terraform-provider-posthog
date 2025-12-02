package httpclient

import (
	"context"
	"fmt"
)

type Dashboard struct {
	ID          int64    `json:"id"`
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Pinned      *bool    `json:"pinned,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Deleted     *bool    `json:"deleted,omitempty"`
}

type DashboardRequest struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Pinned      *bool    `json:"pinned,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Deleted     *bool    `json:"deleted,omitempty"`
}

func (c *PosthogClient) CreateDashboard(ctx context.Context, input DashboardRequest) (Dashboard, error) {
	path := fmt.Sprintf("/api/projects/%s/dashboards/", c.projectID)
	result, _, err := doPost[Dashboard](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetDashboard(ctx context.Context, id int64) (Dashboard, error) {
	path := fmt.Sprintf("/api/projects/%s/dashboards/%d/", c.projectID, id)
	result, _, err := doGet[Dashboard](c, ctx, path)
	return result, err
}

func (c *PosthogClient) UpdateDashboard(ctx context.Context, id int64, input DashboardRequest) (Dashboard, error) {
	path := fmt.Sprintf("/api/projects/%s/dashboards/%d/", c.projectID, id)
	result, _, err := doPatch[Dashboard](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) DeleteDashboard(ctx context.Context, id int64) error {
	deleted := true
	input := DashboardRequest{Deleted: &deleted}
	_, err := c.UpdateDashboard(ctx, id, input)
	return err
}

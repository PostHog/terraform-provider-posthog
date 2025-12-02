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

func (c *PosthogClient) GetDashboard(ctx context.Context, id int64) (Dashboard, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/dashboards/%d/", c.projectID, id)
	return doGet[Dashboard](c, ctx, path)
}

func (c *PosthogClient) UpdateDashboard(ctx context.Context, id int64, input DashboardRequest) (Dashboard, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/dashboards/%d/", c.projectID, id)
	return doPatch[Dashboard](c, ctx, path, input)
}

func (c *PosthogClient) DeleteDashboard(ctx context.Context, id int64) (HTTPStatusCode, error) {
	deleted := true
	input := DashboardRequest{Deleted: &deleted}
	_, statusCode, err := c.UpdateDashboard(ctx, id, input)
	return statusCode, err
}

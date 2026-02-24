package httpclient

import (
	"context"
	"fmt"
)

type DashboardTileInsight struct {
	ID int64 `json:"id"`
}

type DashboardTileText struct {
	ID   int64  `json:"id"`
	Body string `json:"body"`
}

type DashboardTile struct {
	ID      int64                  `json:"id"`
	Insight *DashboardTileInsight  `json:"insight,omitempty"`
	Text    *DashboardTileText     `json:"text,omitempty"`
	Layouts map[string]interface{} `json:"layouts"`
	Color   *string                `json:"color,omitempty"`
	Deleted *bool                  `json:"deleted,omitempty"`
}

type DashboardLayoutResponse struct {
	ID    int64           `json:"id"`
	Tiles []DashboardTile `json:"tiles"`
}

type DashboardTileTextPatch struct {
	Body string `json:"body"`
}

type DashboardTilePatchItem struct {
	ID      int64                   `json:"id,omitempty"`
	Deleted *bool                   `json:"deleted,omitempty"`
	Color   *string                 `json:"color,omitempty"`
	Layouts *map[string]interface{} `json:"layouts,omitempty"`
	Text    *DashboardTileTextPatch `json:"text,omitempty"`
}

type DashboardLayoutPatchRequest struct {
	Tiles []DashboardTilePatchItem `json:"tiles"`
}

// GetDashboardLayout fetches the tile layout for a specific dashboard.
func (c *PosthogClient) GetDashboardLayout(ctx context.Context, projectID, dashboardID string) (DashboardLayoutResponse, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/dashboards/%s/", projectID, dashboardID)
	return doGet[DashboardLayoutResponse](c, ctx, path)
}

// UpdateDashboardLayout applies a partial update to the tile layout of a specific dashboard.
// Only the tiles listed in the request are modified; unmanaged tiles are left untouched.
func (c *PosthogClient) UpdateDashboardLayout(ctx context.Context, projectID, dashboardID string, input DashboardLayoutPatchRequest) (DashboardLayoutResponse, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/dashboards/%s/", projectID, dashboardID)
	return doPatch[DashboardLayoutResponse](c, ctx, path, input)
}

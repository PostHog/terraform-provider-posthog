package httpclient

import (
	"context"
	"fmt"
)

// WarehouseTable represents a PostHog data warehouse table that reads files
// directly from an S3-compatible bucket (S3, GCS, R2, etc.).
type WarehouseTable struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	HogQLName  *string        `json:"hogql_name,omitempty"`
	Format     string         `json:"format"`
	URLPattern string         `json:"url_pattern"`
	Options    map[string]any `json:"options,omitempty"`
	CreatedAt  *string        `json:"created_at,omitempty"`
}

func (c *PosthogClient) CreateWarehouseTable(ctx context.Context, projectID string, input any) (WarehouseTable, error) {
	path := fmt.Sprintf("/api/projects/%s/warehouse_tables/", projectID)
	result, _, err := doPost[WarehouseTable](c, ctx, path, input)
	return result, err
}

// GetWarehouseTable skips column serialization, which builds the full HogQL
// database on every read. Create and update cannot skip it: PostHog needs that
// database to check that the table name is unique.
func (c *PosthogClient) GetWarehouseTable(ctx context.Context, projectID, id string) (WarehouseTable, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/warehouse_tables/%s/?include_columns=false", projectID, id)
	return doGet[WarehouseTable](c, ctx, path)
}

func (c *PosthogClient) UpdateWarehouseTable(ctx context.Context, projectID, id string, input any) (WarehouseTable, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/warehouse_tables/%s/", projectID, id)
	return doPatch[WarehouseTable](c, ctx, path, input)
}

func (c *PosthogClient) DeleteWarehouseTable(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/warehouse_tables/%s/", projectID, id)
	return doDelete(c, ctx, path)
}

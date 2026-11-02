package httpclient

import (
	"context"
	"fmt"
)

// WarehouseTable is a self-managed data warehouse table: PostHog queries the
// files matching URLPattern in place, in a bucket the customer controls.
type WarehouseTable struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Format     string `json:"format"`
	URLPattern string `json:"url_pattern"`
	CreatedAt  string `json:"created_at"`
}

type WarehouseTableRequest struct {
	Name       string                   `json:"name"`
	Format     string                   `json:"format"`
	URLPattern string                   `json:"url_pattern"`
	Credential WarehouseTableCredential `json:"credential"`
}

// WarehouseTableCredential is write-only: PostHog stores it encrypted and never
// returns the key or the secret.
type WarehouseTableCredential struct {
	AccessKey    string `json:"access_key"`
	AccessSecret string `json:"access_secret"`
}

func (c *PosthogClient) CreateWarehouseTable(ctx context.Context, projectID string, input WarehouseTableRequest) (WarehouseTable, error) {
	path := fmt.Sprintf("/api/projects/%s/warehouse_tables/", projectID)
	result, _, err := doPost[WarehouseTable](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetWarehouseTable(ctx context.Context, projectID, id string) (WarehouseTable, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/warehouse_tables/%s/", projectID, id)
	return doGet[WarehouseTable](c, ctx, path)
}

func (c *PosthogClient) UpdateWarehouseTable(ctx context.Context, projectID, id string, input WarehouseTableRequest) (WarehouseTable, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/warehouse_tables/%s/", projectID, id)
	return doPatch[WarehouseTable](c, ctx, path, input)
}

func (c *PosthogClient) DeleteWarehouseTable(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/warehouse_tables/%s/", projectID, id)
	return doDelete(c, ctx, path)
}

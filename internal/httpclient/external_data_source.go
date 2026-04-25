package httpclient

import (
	"context"
	"fmt"
)

// ExternalDataSource represents a PostHog data warehouse source (Stripe, Postgres,
// Snowflake, BigQuery, Hubspot, etc.) that syncs external data into PostHog.
type ExternalDataSource struct {
	ID         string                       `json:"id"`
	SourceType string                       `json:"source_type"`
	Prefix     *string                      `json:"prefix,omitempty"`
	JobInputs  map[string]interface{}       `json:"job_inputs,omitempty"`
	Schemas    []ExternalDataSourceSchema   `json:"schemas,omitempty"`
	Status     *string                      `json:"status,omitempty"`
	LastRunAt  *string                      `json:"last_run_at,omitempty"`
	CreatedAt  *string                      `json:"created_at,omitempty"`
}

type ExternalDataSourceSchema struct {
	Name          string  `json:"name"`
	ShouldSync    bool    `json:"should_sync"`
	SyncFrequency *string `json:"sync_frequency,omitempty"`
	Status        *string `json:"status,omitempty"`
}

func (c *PosthogClient) CreateExternalDataSource(ctx context.Context, projectID string, input interface{}) (ExternalDataSource, error) {
	path := fmt.Sprintf("/api/projects/%s/external_data_sources/", projectID)
	result, _, err := doPost[ExternalDataSource](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetExternalDataSource(ctx context.Context, projectID, id string) (ExternalDataSource, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/external_data_sources/%s/", projectID, id)
	return doGet[ExternalDataSource](c, ctx, path)
}

func (c *PosthogClient) UpdateExternalDataSource(ctx context.Context, projectID, id string, input interface{}) (ExternalDataSource, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/external_data_sources/%s/", projectID, id)
	return doPatch[ExternalDataSource](c, ctx, path, input)
}

func (c *PosthogClient) DeleteExternalDataSource(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/external_data_sources/%s/", projectID, id)
	return doDelete(c, ctx, path)
}

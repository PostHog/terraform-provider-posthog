package httpclient

import (
	"context"
	"fmt"
	"net/url"
)

// Endpoint is a named, versioned API route that runs a HogQL or insight query.
// PostHog addresses an endpoint by its name, not its UUID, so every path below
// takes the name.
type Endpoint struct {
	ID                   string                 `json:"id"`
	Name                 string                 `json:"name"`
	Description          *string                `json:"description,omitempty"`
	Query                map[string]interface{} `json:"query,omitempty"`
	IsActive             *bool                  `json:"is_active,omitempty"`
	DataFreshnessSeconds *int64                 `json:"data_freshness_seconds,omitempty"`
	// The top-level is_materialized is true only once a materialization run has
	// finished. Materialization.Enabled says whether materialization is turned on,
	// which is what a configuration controls.
	Materialization *EndpointMaterialization `json:"materialization,omitempty"`
	BucketOverrides map[string]string        `json:"bucket_overrides,omitempty"`
	CurrentVersion  *int64                   `json:"current_version,omitempty"`
	// URL is the absolute URL that runs the endpoint.
	URL  *string  `json:"url,omitempty"`
	Tags []string `json:"tags,omitempty"`
	// OptionalBreakdownProperties lists the breakdowns a caller may leave out of a run.
	OptionalBreakdownProperties []string `json:"optional_breakdown_properties,omitempty"`
	CreatedAt                   *string  `json:"created_at,omitempty"`
	// MaterializationError is set when an update saved a new query version but
	// could not materialize it. The update itself still returns 200.
	MaterializationError *string `json:"materialization_error,omitempty"`
}

type EndpointMaterialization struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type EndpointRequest struct {
	// Name is only read on create. PostHog ignores it on update.
	Name *string `json:"name,omitempty"`
	// A query that differs from the current version's creates a new version.
	Query                map[string]interface{} `json:"query,omitempty"`
	Description          *string                `json:"description,omitempty"`
	DataFreshnessSeconds *int64                 `json:"data_freshness_seconds,omitempty"`
	IsActive             *bool                  `json:"is_active,omitempty"`
	IsMaterialized       *bool                  `json:"is_materialized,omitempty"`
	// The pointer lets an empty map reach PostHog, which is how overrides are cleared.
	BucketOverrides             *map[string]string `json:"bucket_overrides,omitempty"`
	Tags                        *[]string          `json:"tags,omitempty"`
	OptionalBreakdownProperties *[]string          `json:"optional_breakdown_properties,omitempty"`
}

func endpointPath(projectID, name string) string {
	return fmt.Sprintf("/api/projects/%s/endpoints/%s/", projectID, url.PathEscape(name))
}

func (c *PosthogClient) CreateEndpoint(ctx context.Context, projectID string, input EndpointRequest) (Endpoint, error) {
	path := fmt.Sprintf("/api/projects/%s/endpoints/", projectID)
	result, _, err := doPost[Endpoint](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetEndpoint(ctx context.Context, projectID, name string) (Endpoint, HTTPStatusCode, error) {
	return doGet[Endpoint](c, ctx, endpointPath(projectID, name))
}

func (c *PosthogClient) UpdateEndpoint(ctx context.Context, projectID, name string, input EndpointRequest) (Endpoint, HTTPStatusCode, error) {
	return doPatch[Endpoint](c, ctx, endpointPath(projectID, name), input)
}

// DeleteEndpoint soft-deletes the endpoint and turns off its materializations.
// The name is free for a new endpoint straight after.
func (c *PosthogClient) DeleteEndpoint(ctx context.Context, projectID, name string) (HTTPStatusCode, error) {
	return doDelete(c, ctx, endpointPath(projectID, name))
}

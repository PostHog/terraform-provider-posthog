package httpclient

import (
	"context"
	"fmt"
)

const (
	cohortCollectionPathFormat = "/api/projects/%s/cohorts/"
	cohortResourcePathFormat   = "/api/projects/%s/cohorts/%s/"
)

type Cohort struct {
	ID          int64                  `json:"id"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Filters     map[string]interface{} `json:"filters,omitempty"`
	IsStatic    *bool                  `json:"is_static,omitempty"`
	Deleted     *bool                  `json:"deleted,omitempty"`
	CreatedAt   *string                `json:"created_at,omitempty"`
}

type CohortRequest struct {
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Filters     map[string]interface{} `json:"filters,omitempty"`
	IsStatic    *bool                  `json:"is_static,omitempty"`
	Deleted     *bool                  `json:"deleted,omitempty"`
}

func (c *PosthogClient) CreateCohort(ctx context.Context, projectID string, input CohortRequest) (Cohort, error) {
	path := fmt.Sprintf(cohortCollectionPathFormat, projectID)
	result, _, err := doPost[Cohort](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetCohort(ctx context.Context, projectID, id string) (Cohort, HTTPStatusCode, error) {
	path := fmt.Sprintf(cohortResourcePathFormat, projectID, id)
	return doGet[Cohort](c, ctx, path)
}

func (c *PosthogClient) UpdateCohort(ctx context.Context, projectID, id string, input CohortRequest) (Cohort, HTTPStatusCode, error) {
	path := fmt.Sprintf(cohortResourcePathFormat, projectID, id)
	return doPatch[Cohort](c, ctx, path, input)
}

// DeleteCohort soft-deletes the cohort. PostHog rejects a hard DELETE on this
// endpoint with 405 ("Hard delete of this model is not allowed"), so clear it
// by patching deleted=true, the same way actions and feature flags are removed.
func (c *PosthogClient) DeleteCohort(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	deleted := true
	_, statusCode, err := c.UpdateCohort(ctx, projectID, id, CohortRequest{Deleted: &deleted})
	return statusCode, err
}

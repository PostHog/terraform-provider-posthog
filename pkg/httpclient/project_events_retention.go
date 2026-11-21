package httpclient

import (
	"context"
	"fmt"
)

type ProjectEventsRetention struct {
	ProjectRetentionMonths *int64 `json:"project_retention_months"`
}

type ProjectEventsRetentionRequest struct {
	RetentionMonths *int64 `json:"retention_months"`
}

func (c *PosthogClient) GetProjectEventsRetention(ctx context.Context, projectID string) (ProjectEventsRetention, HTTPStatusCode, error) {
	return doGet[ProjectEventsRetention](c, ctx, fmt.Sprintf("/api/projects/%s/events_retention/", projectID))
}

func (c *PosthogClient) SetProjectEventsRetention(ctx context.Context, projectID string, input ProjectEventsRetentionRequest) (ProjectEventsRetention, HTTPStatusCode, error) {
	return doPost[ProjectEventsRetention](c, ctx, fmt.Sprintf("/api/projects/%s/events_retention/", projectID), input)
}

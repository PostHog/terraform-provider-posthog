package httpclient

import (
	"context"
	"fmt"
)

// EventsRetention is the part of a project's events retention response that the
// provider reads. PostHog serves it from its own endpoint, not from the
// environment, and only organization admins can change it.
type EventsRetention struct {
	// ProjectRetentionMonths is the period set for this project. It is null while
	// the project uses its organization's period.
	ProjectRetentionMonths *int64 `json:"project_retention_months"`
}

// EventsRetentionRequest sets a project's events retention period. The API also
// accepts null to return the project to its organization's period, which can
// delete events; the provider never sends it, so the field is not a pointer.
type EventsRetentionRequest struct {
	RetentionMonths int64 `json:"retention_months"`
}

func (c *PosthogClient) GetEventsRetention(ctx context.Context, projectID string) (EventsRetention, HTTPStatusCode, error) {
	return doGet[EventsRetention](c, ctx, fmt.Sprintf("/api/projects/%s/events_retention/", projectID))
}

func (c *PosthogClient) SetEventsRetention(ctx context.Context, projectID string, input EventsRetentionRequest) (EventsRetention, HTTPStatusCode, error) {
	return doPost[EventsRetention](c, ctx, fmt.Sprintf("/api/projects/%s/events_retention/", projectID), input)
}

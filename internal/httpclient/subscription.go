package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
)

// Subscription is the API representation of a PostHog subscription (recurring
// dashboard/insight digest delivered to email or Slack on an rrule schedule).
type Subscription struct {
	ID                      int64    `json:"id"`
	TargetType              string   `json:"target_type"`
	TargetValue             string   `json:"target_value"`
	IntegrationID           *int64   `json:"integration_id,omitempty"`
	Dashboard               *int64   `json:"dashboard,omitempty"`
	Insight                 *int64   `json:"insight,omitempty"`
	DashboardExportInsights []int64  `json:"dashboard_export_insights,omitempty"`
	Frequency               string   `json:"frequency"`
	Interval                int64    `json:"interval"`
	StartDate               string   `json:"start_date"`
	ByWeekday               []string `json:"byweekday,omitempty"`
	BySetPos                *int64   `json:"bysetpos,omitempty"`
	Enabled                 *bool    `json:"enabled,omitempty"`
	Title                   *string  `json:"title,omitempty"`
	Deleted                 *bool    `json:"deleted,omitempty"`

	// AI-summary ("prompt") subscription fields. Prompt selects the ai_prompt resource_type
	// (mutually exclusive with dashboard/insight). AIPromptConfig is a free-form JSON blob
	// (e.g. an analysis window). SummaryEnabled/SummaryPromptGuide drive the optional
	// AI-generated summary that can accompany any subscription.
	Prompt             *string         `json:"prompt,omitempty"`
	AIPromptConfig     json.RawMessage `json:"ai_prompt_config,omitempty"`
	SummaryEnabled     *bool           `json:"summary_enabled,omitempty"`
	SummaryPromptGuide *string         `json:"summary_prompt_guide,omitempty"`

	// Read-only / server-computed fields.
	ResourceType     *string `json:"resource_type,omitempty"`
	Summary          *string `json:"summary,omitempty"`
	NextDeliveryDate *string `json:"next_delivery_date,omitempty"`
	CreatedAt        *string `json:"created_at,omitempty"`
}

// SubscriptionRequest is the write payload for create (POST) and update (PATCH).
// resource_type is server-inferred from which subject is set and must never be sent.
//
// integration_id, dashboard_export_insights, byweekday and bysetpos deliberately omit
// ,omitempty: an optional field removed from config must reach the wire as an explicit
// null (int64s) or [] (slices) so the API clears the previously-stored value instead of
// keeping it. BuildCreateRequest defaults the two slices to a non-nil [] because the API
// rejects a null slice on create; [] both creates-as-unset and clears-on-update.
type SubscriptionRequest struct {
	TargetType              string   `json:"target_type,omitempty"`
	TargetValue             string   `json:"target_value,omitempty"`
	IntegrationID           *int64   `json:"integration_id"`
	Dashboard               *int64   `json:"dashboard,omitempty"`
	Insight                 *int64   `json:"insight,omitempty"`
	DashboardExportInsights []int64  `json:"dashboard_export_insights"`
	Frequency               string   `json:"frequency,omitempty"`
	Interval                int64    `json:"interval,omitempty"`
	StartDate               string   `json:"start_date,omitempty"`
	ByWeekday               []string `json:"byweekday"`
	BySetPos                *int64   `json:"bysetpos"`
	Enabled                 *bool    `json:"enabled,omitempty"`
	Title                   *string  `json:"title,omitempty"`

	// AI-summary ("prompt") subscription write fields. AIPromptConfig is sent verbatim; an
	// explicit "{}" clears a previously-set config (omitempty would drop it otherwise).
	Prompt             *string         `json:"prompt,omitempty"`
	AIPromptConfig     json.RawMessage `json:"ai_prompt_config,omitempty"`
	SummaryEnabled     *bool           `json:"summary_enabled,omitempty"`
	SummaryPromptGuide *string         `json:"summary_prompt_guide,omitempty"`
}

func (c *PosthogClient) CreateSubscription(ctx context.Context, projectID string, input SubscriptionRequest) (Subscription, error) {
	path := fmt.Sprintf("/api/projects/%s/subscriptions/", projectID)
	result, _, err := doPost[Subscription](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetSubscription(ctx context.Context, projectID, id string) (Subscription, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/subscriptions/%s/", projectID, id)
	return doGet[Subscription](c, ctx, path)
}

func (c *PosthogClient) UpdateSubscription(ctx context.Context, projectID, id string, input SubscriptionRequest) (Subscription, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/subscriptions/%s/", projectID, id)
	return doPatch[Subscription](c, ctx, path, input)
}

// DeleteSubscription soft-deletes the subscription. Hard DELETE is forbidden by the
// API (405, the viewset is ForbidDestroyModel); deletion is a PATCH setting deleted=true,
// mirroring DeleteDashboard/DeleteInsight.
func (c *PosthogClient) DeleteSubscription(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	// Send a minimal {"deleted": true} payload rather than a SubscriptionRequest: that struct
	// intentionally omits ,omitempty on integration_id/byweekday/bysetpos/dashboard_export_insights,
	// so a bare SubscriptionRequest would emit nulls the API rejects (e.g. a Slack sub requires
	// integration_id, a dashboard sub requires a non-empty insight selection).
	path := fmt.Sprintf("/api/projects/%s/subscriptions/%s/", projectID, id)
	_, statusCode, err := doPatch[Subscription](c, ctx, path, map[string]bool{"deleted": true})
	return statusCode, err
}

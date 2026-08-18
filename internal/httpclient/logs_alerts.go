package httpclient

import (
	"context"
	"fmt"
)

type LogsAlert struct {
	ID                  string             `json:"id"`
	Name                *string            `json:"name,omitempty"`
	Enabled             *bool              `json:"enabled,omitempty"`
	Filters             *LogsAlertFilters  `json:"filters,omitempty"`
	ThresholdCount      *int64             `json:"threshold_count,omitempty"`
	ThresholdOperator   *string            `json:"threshold_operator,omitempty"`
	WindowMinutes       *int64             `json:"window_minutes,omitempty"`
	EvaluationPeriods   *int64             `json:"evaluation_periods,omitempty"`
	DatapointsToAlarm   *int64             `json:"datapoints_to_alarm,omitempty"`
	CooldownMinutes     *int64             `json:"cooldown_minutes,omitempty"`
	ScheduleRestriction *LogsAlertSchedule `json:"schedule_restriction,omitempty"`
	SnoozeUntil         *string            `json:"snooze_until,omitempty"`
	State               *string            `json:"state,omitempty"`
	CreatedAt           *string            `json:"created_at,omitempty"`
	UpdatedAt           *string            `json:"updated_at,omitempty"`
}

type LogsAlertRequest struct {
	Name    *string `json:"name,omitempty"`
	Enabled *bool   `json:"enabled,omitempty"`
	// Filters and ScheduleRestriction are whole-object replacements, so they are
	// sent on every request. A nil ScheduleRestriction marshals to null, which is
	// how quiet hours are cleared.
	Filters             *LogsAlertFilters  `json:"filters"`
	ThresholdCount      *int64             `json:"threshold_count,omitempty"`
	ThresholdOperator   *string            `json:"threshold_operator,omitempty"`
	WindowMinutes       *int64             `json:"window_minutes,omitempty"`
	EvaluationPeriods   *int64             `json:"evaluation_periods,omitempty"`
	DatapointsToAlarm   *int64             `json:"datapoints_to_alarm,omitempty"`
	CooldownMinutes     *int64             `json:"cooldown_minutes,omitempty"`
	ScheduleRestriction *LogsAlertSchedule `json:"schedule_restriction"`
	// Sent only when configured. Omitting it leaves whatever snooze an operator set in
	// the PostHog UI untouched, which is the point: Terraform manages this only if asked.
	SnoozeUntil *string `json:"snooze_until,omitempty"`
}

// LogsAlertFilters omits no fields: the API accepts null for each key and treats the
// whole object as a replacement, so sending an explicit null is what clears a filter.
// With omitempty, "cleared" and "unchanged" would be identical on the wire.
type LogsAlertFilters struct {
	SeverityLevels []string       `json:"severityLevels"`
	ServiceNames   []string       `json:"serviceNames"`
	FilterGroup    map[string]any `json:"filterGroup"`
}

type LogsAlertSchedule struct {
	BlockedWindows []LogsAlertBlockedWindow `json:"blocked_windows"`
}

type LogsAlertBlockedWindow struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func (c *PosthogClient) CreateLogsAlert(ctx context.Context, projectID string, input LogsAlertRequest) (LogsAlert, error) {
	path := fmt.Sprintf("/api/projects/%s/logs/alerts/", projectID)
	result, _, err := doPost[LogsAlert](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetLogsAlert(ctx context.Context, projectID, id string) (LogsAlert, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/logs/alerts/%s/", projectID, id)
	return doGet[LogsAlert](c, ctx, path)
}

func (c *PosthogClient) UpdateLogsAlert(ctx context.Context, projectID, id string, input LogsAlertRequest) (LogsAlert, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/logs/alerts/%s/", projectID, id)
	return doPatch[LogsAlert](c, ctx, path, input)
}

func (c *PosthogClient) DeleteLogsAlert(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/logs/alerts/%s/", projectID, id)
	return doDelete(c, ctx, path)
}

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

type LogsAlertDestination struct {
	HogFunctionIDs   []string `json:"hog_function_ids"`
	Type             string   `json:"type,omitempty"`
	SlackWorkspaceID *int64   `json:"slack_workspace_id,omitempty"`
	SlackChannelID   *string  `json:"slack_channel_id,omitempty"`
	WebhookURL       *string  `json:"webhook_url,omitempty"`
}

type LogsAlertDestinationRequest struct {
	Type             string  `json:"type"`
	SlackWorkspaceID *int64  `json:"slack_workspace_id,omitempty"`
	SlackChannelID   *string `json:"slack_channel_id,omitempty"`
	SlackChannelName *string `json:"slack_channel_name,omitempty"`
	WebhookURL       *string `json:"webhook_url,omitempty"`
}

type logsAlertDestinationDeleteRequest struct {
	HogFunctionIDs []string `json:"hog_function_ids"`
}

func logsAlertDestinationsPath(projectID, alertID string) string {
	return fmt.Sprintf("/api/projects/%s/logs/alerts/%s/destinations/", projectID, alertID)
}

func (c *PosthogClient) ListLogsAlertDestinations(ctx context.Context, projectID, alertID string) ([]LogsAlertDestination, HTTPStatusCode, error) {
	hogFunctions, status, err := c.ListLogsAlertHogFunctions(ctx, projectID, alertID)
	if err != nil {
		return nil, status, err
	}
	return groupLogsAlertDestinations(hogFunctions), status, nil
}

func groupLogsAlertDestinations(hogFunctions []HogFunction) []LogsAlertDestination {
	groups := make(map[string]int)
	destinations := make([]LogsAlertDestination, 0)

	for _, hogFunction := range hogFunctions {
		destination, key := logsAlertDestinationFromHogFunction(hogFunction)
		if index, exists := groups[key]; exists {
			destinations[index].HogFunctionIDs = append(destinations[index].HogFunctionIDs, hogFunction.ID)
			continue
		}

		groups[key] = len(destinations)
		destinations = append(destinations, destination)
	}

	return destinations
}

func logsAlertDestinationFromHogFunction(hogFunction HogFunction) (LogsAlertDestination, string) {
	templateID := ""
	if hogFunction.TemplateID != nil {
		templateID = *hogFunction.TemplateID
	} else if hogFunction.Template != nil {
		templateID = hogFunction.Template.ID
	}

	destination := LogsAlertDestination{HogFunctionIDs: []string{hogFunction.ID}}
	switch templateID {
	case "template-slack":
		destination.Type = "slack"
		destination.SlackWorkspaceID = inputInt64(hogFunction.Inputs, "slack_workspace")
		destination.SlackChannelID = inputString(hogFunction.Inputs, "channel")
		return destination, fmt.Sprintf("%s:%d:%s", templateID, inputInt64Value(hogFunction.Inputs, "slack_workspace"), inputStringValue(hogFunction.Inputs, "channel"))
	case "template-microsoft-teams":
		destination.Type = "teams"
		destination.WebhookURL = inputString(hogFunction.Inputs, "webhookUrl")
		return destination, templateID + ":" + inputStringValue(hogFunction.Inputs, "webhookUrl")
	case "template-webhook":
		destination.Type = "webhook"
		destination.WebhookURL = inputString(hogFunction.Inputs, "url")
		return destination, templateID + ":" + inputStringValue(hogFunction.Inputs, "url")
	default:
		return destination, hogFunction.ID
	}
}

func inputValue(inputs map[string]interface{}, key string) any {
	input, ok := inputs[key].(map[string]interface{})
	if !ok {
		return nil
	}
	return input["value"]
}

func inputString(inputs map[string]interface{}, key string) *string {
	value := inputStringValue(inputs, key)
	if value == "" {
		return nil
	}
	return &value
}

func inputStringValue(inputs map[string]interface{}, key string) string {
	value, _ := inputValue(inputs, key).(string)
	return value
}

func inputInt64(inputs map[string]interface{}, key string) *int64 {
	value := inputInt64Value(inputs, key)
	if value == 0 {
		return nil
	}
	return &value
}

func inputInt64Value(inputs map[string]interface{}, key string) int64 {
	value, _ := inputValue(inputs, key).(float64)
	return int64(value)
}

func (c *PosthogClient) CreateLogsAlertDestination(ctx context.Context, projectID, alertID string, input LogsAlertDestinationRequest) (LogsAlertDestination, error) {
	result, _, err := doPost[LogsAlertDestination](c, ctx, logsAlertDestinationsPath(projectID, alertID), input)
	return result, err
}

func (c *PosthogClient) DeleteLogsAlertDestination(ctx context.Context, projectID, alertID string, hogFunctionIDs []string) (HTTPStatusCode, error) {
	path := logsAlertDestinationsPath(projectID, alertID) + "delete"
	return doPostNoContent(c, ctx, path, logsAlertDestinationDeleteRequest{HogFunctionIDs: hogFunctionIDs})
}

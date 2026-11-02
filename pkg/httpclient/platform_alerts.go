package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// PlatformAlert is a configuration on PostHog's shared alerts platform. Its settings are the same
// for every source kind, except source_config, whose shape depends on source_kind.
type PlatformAlert struct {
	ID                    string                 `json:"id"`
	Name                  *string                `json:"name,omitempty"`
	Enabled               *bool                  `json:"enabled,omitempty"`
	SourceKind            *string                `json:"source_kind,omitempty"`
	SourceConfig          json.RawMessage        `json:"source_config,omitempty"`
	EvaluationPeriods     *int64                 `json:"evaluation_periods,omitempty"`
	DatapointsToAlarm     *int64                 `json:"datapoints_to_alarm,omitempty"`
	CooldownMinutes       *int64                 `json:"cooldown_minutes,omitempty"`
	ScheduleRestriction   *LogsAlertSchedule     `json:"schedule_restriction,omitempty"`
	Grouping              *PlatformAlertGrouping `json:"grouping,omitempty"`
	LegacyConfigurationID *string                `json:"legacy_configuration_id,omitempty"`
	Owner                 *string                `json:"owner,omitempty"`
	CheckStatus           *string                `json:"check_status,omitempty"`
	SnoozeUntil           *string                `json:"snooze_until,omitempty"`
}

// PlatformAlertRequest is the body of both a create and a full update. PostHog's PUT requires
// every field, so none is omitted.
type PlatformAlertRequest struct {
	SourceKind          string                `json:"source_kind"`
	Name                string                `json:"name"`
	Enabled             bool                  `json:"enabled"`
	SourceConfig        any                   `json:"source_config"`
	Grouping            PlatformAlertGrouping `json:"grouping"`
	ScheduleRestriction *LogsAlertSchedule    `json:"schedule_restriction"`
	EvaluationPeriods   int64                 `json:"evaluation_periods"`
	DatapointsToAlarm   int64                 `json:"datapoints_to_alarm"`
	CooldownMinutes     int64                 `json:"cooldown_minutes"`
}

type PlatformAlertGrouping struct {
	Mode         string   `json:"mode"`
	Keys         []string `json:"keys"`
	MaxInstances int64    `json:"max_instances"`
	OnOverflow   string   `json:"on_overflow"`
}

// PlatformAlertLogsSourceConfig is source_config for source_kind = logs.
type PlatformAlertLogsSourceConfig struct {
	Filters   PlatformAlertLogsFilters   `json:"filters"`
	Condition PlatformAlertLogsCondition `json:"condition"`
}

// PlatformAlertLogsFilters omits empty keys, unlike LogsAlertFilters. The platform validates a
// present filters object as a filter, so a draft that sent null for every key would be refused as
// having filters that narrow nothing. source_config is replaced whole, so an omitted key is cleared.
type PlatformAlertLogsFilters struct {
	SeverityLevels []string       `json:"severityLevels,omitempty"`
	ServiceNames   []string       `json:"serviceNames,omitempty"`
	FilterGroup    map[string]any `json:"filterGroup,omitempty"`
}

type PlatformAlertLogsCondition struct {
	ThresholdCount    int64  `json:"threshold_count"`
	ThresholdOperator string `json:"threshold_operator"`
	WindowMinutes     int64  `json:"window_minutes"`
}

type platformAlertList struct {
	Results []PlatformAlert `json:"results"`
}

func platformAlertsPath(projectID string) string {
	return fmt.Sprintf("/api/projects/%s/platform_alerts/", projectID)
}

func (c *PosthogClient) CreatePlatformAlert(ctx context.Context, projectID string, input PlatformAlertRequest) (PlatformAlert, error) {
	result, _, err := doPost[PlatformAlert](c, ctx, platformAlertsPath(projectID), input)
	return result, err
}

func (c *PosthogClient) GetPlatformAlert(ctx context.Context, projectID, id string) (PlatformAlert, HTTPStatusCode, error) {
	return doGet[PlatformAlert](c, ctx, platformAlertsPath(projectID)+id+"/")
}

func (c *PosthogClient) UpdatePlatformAlert(ctx context.Context, projectID, id string, input PlatformAlertRequest) (PlatformAlert, HTTPStatusCode, error) {
	return doPut[PlatformAlert](c, ctx, platformAlertsPath(projectID)+id+"/", input)
}

func (c *PosthogClient) DeletePlatformAlert(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	return doDelete(c, ctx, platformAlertsPath(projectID)+id+"/")
}

// FindPlatformAlertByLegacyID returns the platform alert copied from, or created with, a source
// product's own alert, such as a posthog_logs_alert id. found is false when there is none.
func (c *PosthogClient) FindPlatformAlertByLegacyID(ctx context.Context, projectID, legacyID string) (alert PlatformAlert, found bool, err error) {
	path := platformAlertsPath(projectID) + "?legacy_configuration_id=" + url.QueryEscape(legacyID)
	list, _, err := doGet[platformAlertList](c, ctx, path)
	if err != nil {
		return PlatformAlert{}, false, err
	}
	// A server older than the filter ignores it and lists every alert, so match the field too.
	for _, alert := range list.Results {
		if alert.LegacyConfigurationID != nil && *alert.LegacyConfigurationID == legacyID {
			return alert, true, nil
		}
	}
	return PlatformAlert{}, false, nil
}

package httpclient

import (
	"context"
	"fmt"
)

type Alert struct {
	ID                  string             `json:"id"`
	Name                *string            `json:"name,omitempty"`
	Insight             AlertInsight       `json:"insight"`
	Enabled             *bool              `json:"enabled,omitempty"`
	SubscribedUsers     []AlertUser        `json:"subscribed_users,omitempty"`
	Threshold           *AlertThreshold    `json:"threshold,omitempty"`
	Condition           *AlertCondition    `json:"condition,omitempty"`
	Config              *TrendsAlertConfig `json:"config,omitempty"`
	CalculationInterval *string            `json:"calculation_interval,omitempty"`
	SnoozedUntil        *string            `json:"snoozed_until,omitempty"`
	SkipWeekend         *bool              `json:"skip_weekend,omitempty"`
	State               *string            `json:"state,omitempty"`
	CreatedAt           *string            `json:"created_at,omitempty"`
	CreatedBy           map[string]any     `json:"created_by,omitempty"`
	LastNotifiedAt      *string            `json:"last_notified_at,omitempty"`
	LastCheckedAt       *string            `json:"last_checked_at,omitempty"`
	NextCheckAt         *string            `json:"next_check_at,omitempty"`
	Checks              []AlertCheck       `json:"checks,omitempty"`
}

type AlertRequest struct {
	Name                *string            `json:"name,omitempty"`
	Insight             int64              `json:"insight"`
	Enabled             *bool              `json:"enabled,omitempty"`
	SubscribedUsers     []int64            `json:"subscribed_users"`
	Threshold           *AlertThreshold    `json:"threshold,omitempty"`
	Condition           *AlertCondition    `json:"condition,omitempty"`
	Config              *TrendsAlertConfig `json:"config,omitempty"`
	CalculationInterval *string            `json:"calculation_interval,omitempty"`
	SnoozedUntil        *string            `json:"snoozed_until,omitempty"`
	SkipWeekend         *bool              `json:"skip_weekend,omitempty"`
}

type AlertInsight struct {
	ID int64 `json:"id"`
}

type AlertUser struct {
	ID int64 `json:"id"`
}

type AlertThreshold struct {
	Configuration ThresholdConfiguration `json:"configuration"`
}

type ThresholdConfiguration struct {
	Type   string           `json:"type"` // "absolute" or "percentage"
	Bounds *ThresholdBounds `json:"bounds,omitempty"`
}

type ThresholdBounds struct {
	Lower *float64 `json:"lower,omitempty"`
	Upper *float64 `json:"upper,omitempty"`
}

type AlertCondition struct {
	Type string `json:"type"` // "absolute_value", "relative_increase", "relative_decrease"
}

type TrendsAlertConfig struct {
	Type                 string `json:"type"` // "TrendsAlertConfig"
	SeriesIndex          *int   `json:"series_index,omitempty"`
	CheckOngoingInterval *bool  `json:"check_ongoing_interval,omitempty"`
}

type AlertCheck struct {
	ID              string   `json:"id"`
	CreatedAt       string   `json:"created_at"`
	CalculatedValue *float64 `json:"calculated_value,omitempty"`
	State           string   `json:"state"`
	TargetsNotified bool     `json:"targets_notified"`
}

func (c *PosthogClient) CreateAlert(ctx context.Context, projectID string, input AlertRequest) (Alert, error) {
	path := fmt.Sprintf("/api/environments/%s/alerts/", projectID)
	result, _, err := doPost[Alert](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetAlert(ctx context.Context, projectID, id string) (Alert, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/environments/%s/alerts/%s/", projectID, id)
	return doGet[Alert](c, ctx, path)
}

func (c *PosthogClient) UpdateAlert(ctx context.Context, projectID, id string, input AlertRequest) (Alert, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/environments/%s/alerts/%s/", projectID, id)
	return doPatch[Alert](c, ctx, path, input)
}

func (c *PosthogClient) DeleteAlert(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/environments/%s/alerts/%s/", projectID, id)
	return doDelete(c, ctx, path)
}

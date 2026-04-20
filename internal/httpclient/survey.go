package httpclient

import (
	"context"
	"fmt"
)

const (
	surveyCollectionPathFormat = "/api/projects/%s/surveys/"
	surveyResourcePathFormat   = "/api/projects/%s/surveys/%s/"
)

type Survey struct {
	ID                           string                 `json:"id"`
	Name                         *string                `json:"name,omitempty"`
	Description                  *string                `json:"description,omitempty"`
	Type                         *string                `json:"type,omitempty"`
	Schedule                     *string                `json:"schedule,omitempty"`
	LinkedFlag                   map[string]interface{} `json:"linked_flag,omitempty"`
	TargetingFlag                map[string]interface{} `json:"targeting_flag,omitempty"`
	InternalTargetingFlag        map[string]interface{} `json:"internal_targeting_flag,omitempty"`
	Questions                    []interface{}          `json:"questions,omitempty"`
	Conditions                   interface{}            `json:"conditions,omitempty"`
	Appearance                   interface{}            `json:"appearance,omitempty"`
	CreatedAt                    *string                `json:"created_at,omitempty"`
	CreatedBy                    map[string]interface{} `json:"created_by,omitempty"`
	StartDate                    *string                `json:"start_date,omitempty"`
	EndDate                      *string                `json:"end_date,omitempty"`
	Archived                     *bool                  `json:"archived,omitempty"`
	ResponsesLimit               *int64                 `json:"responses_limit,omitempty"`
	IterationCount               *int64                 `json:"iteration_count,omitempty"`
	IterationFrequencyDays       *int64                 `json:"iteration_frequency_days,omitempty"`
	IterationStartDates          []interface{}          `json:"iteration_start_dates,omitempty"`
	CurrentIteration             *int64                 `json:"current_iteration,omitempty"`
	CurrentIterationStartDate    *string                `json:"current_iteration_start_date,omitempty"`
	ResponseSamplingStartDate    *string                `json:"response_sampling_start_date,omitempty"`
	ResponseSamplingIntervalType *string                `json:"response_sampling_interval_type,omitempty"`
	ResponseSamplingInterval     *int64                 `json:"response_sampling_interval,omitempty"`
	ResponseSamplingLimit        *int64                 `json:"response_sampling_limit,omitempty"`
	ResponseSamplingDailyLimits  interface{}            `json:"response_sampling_daily_limits,omitempty"`
	EnablePartialResponses       *bool                  `json:"enable_partial_responses,omitempty"`
	EnableIframeEmbedding        *bool                  `json:"enable_iframe_embedding,omitempty"`
	Translations                 interface{}            `json:"translations,omitempty"`
	FormContent                  interface{}            `json:"form_content,omitempty"`
}

type SurveyRequest struct {
	Name                         string        `json:"name"`
	Description                  *string       `json:"description,omitempty"`
	Type                         string        `json:"type"`
	Schedule                     *string       `json:"schedule,omitempty"`
	LinkedFlagID                 *int64        `json:"linked_flag_id,omitempty"`
	LinkedInsightID              *int64        `json:"linked_insight_id,omitempty"`
	TargetingFlagID              *int64        `json:"targeting_flag_id,omitempty"`
	TargetingFlagFilters         interface{}   `json:"targeting_flag_filters,omitempty"`
	RemoveTargetingFlag          *bool         `json:"remove_targeting_flag,omitempty"`
	Questions                    []interface{} `json:"questions"`
	Conditions                   interface{}   `json:"conditions,omitempty"`
	Appearance                   interface{}   `json:"appearance,omitempty"`
	StartDate                    *string       `json:"start_date,omitempty"`
	EndDate                      *string       `json:"end_date,omitempty"`
	Archived                     *bool         `json:"archived,omitempty"`
	ResponsesLimit               *int64        `json:"responses_limit,omitempty"`
	IterationCount               *int64        `json:"iteration_count,omitempty"`
	IterationFrequencyDays       *int64        `json:"iteration_frequency_days,omitempty"`
	ResponseSamplingStartDate    *string       `json:"response_sampling_start_date,omitempty"`
	ResponseSamplingIntervalType *string       `json:"response_sampling_interval_type,omitempty"`
	ResponseSamplingInterval     *int64        `json:"response_sampling_interval,omitempty"`
	ResponseSamplingLimit        *int64        `json:"response_sampling_limit,omitempty"`
	ResponseSamplingDailyLimits  interface{}   `json:"response_sampling_daily_limits,omitempty"`
	EnablePartialResponses       *bool         `json:"enable_partial_responses,omitempty"`
	EnableIframeEmbedding        *bool         `json:"enable_iframe_embedding,omitempty"`
	Translations                 interface{}   `json:"translations,omitempty"`
	CreateInFolder               *string       `json:"_create_in_folder,omitempty"`
	FormContent                  interface{}   `json:"form_content,omitempty"`
}

func (c *PosthogClient) CreateSurvey(ctx context.Context, projectID string, input SurveyRequest) (Survey, error) {
	path := fmt.Sprintf(surveyCollectionPathFormat, projectID)
	result, _, err := doPost[Survey](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetSurvey(ctx context.Context, projectID, id string) (Survey, HTTPStatusCode, error) {
	path := fmt.Sprintf(surveyResourcePathFormat, projectID, id)
	return doGet[Survey](c, ctx, path)
}

func (c *PosthogClient) UpdateSurvey(ctx context.Context, projectID, id string, input SurveyRequest) (Survey, HTTPStatusCode, error) {
	path := fmt.Sprintf(surveyResourcePathFormat, projectID, id)
	return doPut[Survey](c, ctx, path, input)
}

func (c *PosthogClient) DeleteSurvey(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	path := fmt.Sprintf(surveyResourcePathFormat, projectID, id)
	return doDelete(c, ctx, path)
}

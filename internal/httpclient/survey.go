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

// SurveyNullableIntegerJSONFields lists the JSON keys whose Go fields drop
// `omitempty` per the SurveyRequest comment below. Both the resource and
// httpclient test packages reference this slice so the list lives in one place.
var SurveyNullableIntegerJSONFields = []string{
	"linked_flag_id",
	"linked_insight_id",
	"responses_limit",
	"iteration_count",
	"iteration_frequency_days",
	"response_sampling_interval",
	"response_sampling_limit",
}

// SurveyRequest is the wire format for POST/PUT to the surveys endpoint.
//
// Fields marked nullable in the upstream OpenAPI schema (the keys in
// SurveyNullableIntegerJSONFields above) deliberately omit `,omitempty` so that
// a nil pointer serialises as JSON `null`. PostHog interprets `null` as "clear
// this column", which is the semantic users expect when they remove the
// attribute from their Terraform config. Fields that are not nullable upstream
// (targeting_flag_id, archived) keep `,omitempty` so an absent value is omitted
// rather than rejected.
type SurveyRequest struct {
	Name                         string        `json:"name"`
	Description                  *string       `json:"description,omitempty"`
	Type                         string        `json:"type"`
	Schedule                     *string       `json:"schedule,omitempty"`
	LinkedFlagID                 *int64        `json:"linked_flag_id"`
	LinkedInsightID              *int64        `json:"linked_insight_id"`
	TargetingFlagID              *int64        `json:"targeting_flag_id,omitempty"`
	TargetingFlagFilters         interface{}   `json:"targeting_flag_filters,omitempty"`
	RemoveTargetingFlag          *bool         `json:"remove_targeting_flag,omitempty"`
	Questions                    []interface{} `json:"questions"`
	Conditions                   interface{}   `json:"conditions,omitempty"`
	Appearance                   interface{}   `json:"appearance,omitempty"`
	StartDate                    *string       `json:"start_date,omitempty"`
	EndDate                      *string       `json:"end_date,omitempty"`
	Archived                     *bool         `json:"archived,omitempty"`
	ResponsesLimit               *int64        `json:"responses_limit"`
	IterationCount               *int64        `json:"iteration_count"`
	IterationFrequencyDays       *int64        `json:"iteration_frequency_days"`
	ResponseSamplingStartDate    *string       `json:"response_sampling_start_date,omitempty"`
	ResponseSamplingIntervalType *string       `json:"response_sampling_interval_type,omitempty"`
	ResponseSamplingInterval     *int64        `json:"response_sampling_interval"`
	ResponseSamplingLimit        *int64        `json:"response_sampling_limit"`
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
	// PostHog's SurveyViewSet.get_serializer_class() routes only POST and PATCH
	// to SurveySerializerCreateUpdateOnly; PUT falls through to SurveySerializer,
	// whose linked_flag_id / linked_insight_id are declared with a dotted source
	// (`source="linked_flag.id"`) and therefore trip DRF's "writable dotted-source
	// fields" AssertionError on update. Use PATCH so the request hits the
	// write-safe serializer path.
	return doPatch[Survey](c, ctx, path, input)
}

func (c *PosthogClient) DeleteSurvey(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	path := fmt.Sprintf(surveyResourcePathFormat, projectID, id)
	return doDelete(c, ctx, path)
}

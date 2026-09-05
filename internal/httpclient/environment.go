package httpclient

import (
	"context"
	"fmt"
)

// Environment represents a PostHog environment (also referred to as a project/team).
// Project-level settings such as heatmaps, autocapture and session recording opt-ins
// live on this object and are managed via the /api/environments/{id}/ endpoint.
//
// Pointer types are used so that an unset field (nil) can be distinguished from a
// field that is explicitly set to its zero value (e.g. false or 0).
type Environment struct {
	ID                                          int64                        `json:"id"`
	HeatmapsOptIn                               *bool                        `json:"heatmaps_opt_in,omitempty"`
	AutocaptureExceptionsOptIn                  *bool                        `json:"autocapture_exceptions_opt_in,omitempty"`
	SessionRecordingOptIn                       *bool                        `json:"session_recording_opt_in,omitempty"`
	SurveysOptIn                                *bool                        `json:"surveys_opt_in,omitempty"`
	CookielessServerHashMode                    *int64                       `json:"cookieless_server_hash_mode,omitempty"`
	AutocaptureWebVitalsOptIn                   *bool                        `json:"autocapture_web_vitals_opt_in,omitempty"`
	CapturePerformanceOptIn                     *bool                        `json:"capture_performance_opt_in,omitempty"`
	AnonymizeIps                                *bool                        `json:"anonymize_ips,omitempty"`
	SessionRecordingNetworkPayloadCaptureConfig *NetworkPayloadCaptureConfig `json:"session_recording_network_payload_capture_config,omitempty"`
	AppURLs                                     *[]string                    `json:"app_urls,omitempty"`
	RecordingDomains                            *[]string                    `json:"recording_domains,omitempty"`
	// TestAccountFilters is the "internal and test users" filter list, a JSON array of
	// filter objects; server-injected cohort_name is stripped in MapResponseToModel.
	TestAccountFilters *[]interface{} `json:"test_account_filters,omitempty"`
	// TestAccountFiltersDefaultChecked is the "apply these filters by default" toggle.
	// PostHog may return null when it has never been set.
	TestAccountFiltersDefaultChecked *bool `json:"test_account_filters_default_checked,omitempty"`
}

// NetworkPayloadCaptureConfig mirrors the nullable JSON blob PostHog stores in
// session_recording_network_payload_capture_config. Note the camelCase keys:
// this is a JSON value consumed by the posthog-js SDK, not a REST field.
// Inner fields are pointers so an explicit false survives serialization and a
// missing key can be distinguished from false.
type NetworkPayloadCaptureConfig struct {
	RecordHeaders *bool `json:"recordHeaders,omitempty"`
	RecordBody    *bool `json:"recordBody,omitempty"`
}

// EnvironmentSettingsRequest carries the writable subset of environment settings.
// Fields left nil are omitted from the PATCH body so that only configured settings
// are sent to the API. It deliberately mirrors Environment minus ID rather than
// reusing Environment as the request body: ID has no omitempty, so reusing it would
// serialize "id":0 into every PATCH.
type EnvironmentSettingsRequest struct {
	HeatmapsOptIn                               *bool                        `json:"heatmaps_opt_in,omitempty"`
	AutocaptureExceptionsOptIn                  *bool                        `json:"autocapture_exceptions_opt_in,omitempty"`
	SessionRecordingOptIn                       *bool                        `json:"session_recording_opt_in,omitempty"`
	SurveysOptIn                                *bool                        `json:"surveys_opt_in,omitempty"`
	CookielessServerHashMode                    *int64                       `json:"cookieless_server_hash_mode,omitempty"`
	AutocaptureWebVitalsOptIn                   *bool                        `json:"autocapture_web_vitals_opt_in,omitempty"`
	CapturePerformanceOptIn                     *bool                        `json:"capture_performance_opt_in,omitempty"`
	AnonymizeIps                                *bool                        `json:"anonymize_ips,omitempty"`
	SessionRecordingNetworkPayloadCaptureConfig *NetworkPayloadCaptureConfig `json:"session_recording_network_payload_capture_config,omitempty"`
	AppURLs                                     *[]string                    `json:"app_urls,omitempty"`
	RecordingDomains                            *[]string                    `json:"recording_domains,omitempty"`
	TestAccountFilters                          *[]interface{}               `json:"test_account_filters,omitempty"`
	TestAccountFiltersDefaultChecked            *bool                        `json:"test_account_filters_default_checked,omitempty"`
}

func (c *PosthogClient) GetEnvironment(ctx context.Context, projectID string) (Environment, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/environments/%s/", projectID)
	return doGet[Environment](c, ctx, path)
}

func (c *PosthogClient) UpdateEnvironment(ctx context.Context, projectID string, input EnvironmentSettingsRequest) (Environment, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/environments/%s/", projectID)
	return doPatch[Environment](c, ctx, path, input)
}

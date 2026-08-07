package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
)

// Experiment is the API representation of a PostHog experiment. JSON-blob fields
// (parameters, feature_flag, metrics, metrics_secondary, exposure_criteria) are kept as
// raw JSON so the resource layer can normalize them for state (semantic compare), the same
// way feature_flag filters are handled. Scalar/lifecycle fields are typed.
type Experiment struct {
	ID                int64           `json:"id"`
	Name              string          `json:"name"`
	Description       *string         `json:"description,omitempty"`
	FeatureFlagKey    string          `json:"feature_flag_key,omitempty"`
	FeatureFlag       json.RawMessage `json:"feature_flag,omitempty"`
	Parameters        json.RawMessage `json:"parameters,omitempty"`
	Metrics           json.RawMessage `json:"metrics,omitempty"`
	MetricsSecondary  json.RawMessage `json:"metrics_secondary,omitempty"`
	ExposureCriteria  json.RawMessage `json:"exposure_criteria,omitempty"`
	HoldoutID         *int64          `json:"holdout_id,omitempty"`
	StartDate         *string         `json:"start_date,omitempty"`
	EndDate           *string         `json:"end_date,omitempty"`
	Archived          *bool           `json:"archived,omitempty"`
	Deleted           *bool           `json:"deleted,omitempty"`
	Conclusion        *string         `json:"conclusion,omitempty"`
	ConclusionComment *string         `json:"conclusion_comment,omitempty"`
	// Status is the server-derived lifecycle state: draft | running | paused |
	// exposure_frozen | stopped. Read-only.
	Status string `json:"status,omitempty"`
}

// ExperimentRequest is the create/update body. All fields are pointers/omitempty so updates
// send only what changed (PATCH is a partial update) and soft-delete can send just
// {deleted: true} without blanking other fields.
type ExperimentRequest struct {
	Name               *string         `json:"name,omitempty"`
	Description        *string         `json:"description,omitempty"`
	FeatureFlagKey     *string         `json:"feature_flag_key,omitempty"`
	Metrics            json.RawMessage `json:"metrics,omitempty"`
	MetricsSecondary   json.RawMessage `json:"metrics_secondary,omitempty"`
	ExposureCriteria   json.RawMessage `json:"exposure_criteria,omitempty"`
	HoldoutID          *int64          `json:"holdout_id,omitempty"`
	Archived           *bool           `json:"archived,omitempty"`
	Deleted            *bool           `json:"deleted,omitempty"`
	AllowUnknownEvents *bool           `json:"allow_unknown_events,omitempty"`
}

// ExperimentEndRequest is the body for the end sub-action.
type ExperimentEndRequest struct {
	Conclusion        *string `json:"conclusion,omitempty"`
	ConclusionComment *string `json:"conclusion_comment,omitempty"`
}

// ExperimentShipVariantRequest is the body for the ship_variant sub-action.
type ExperimentShipVariantRequest struct {
	VariantKey        string  `json:"variant_key"`
	ReleaseToEveryone bool    `json:"release_to_everyone"`
	Conclusion        *string `json:"conclusion,omitempty"`
	ConclusionComment *string `json:"conclusion_comment,omitempty"`
}

func experimentsPath(projectID string) string {
	return fmt.Sprintf("/api/projects/%s/experiments/", projectID)
}

func experimentPath(projectID, id string) string {
	return fmt.Sprintf("/api/projects/%s/experiments/%s/", projectID, id)
}

func experimentActionPath(projectID, id, action string) string {
	return fmt.Sprintf("/api/projects/%s/experiments/%s/%s/", projectID, id, action)
}

func (c *PosthogClient) CreateExperiment(ctx context.Context, projectID string, input ExperimentRequest) (Experiment, error) {
	result, _, err := doPost[Experiment](c, ctx, experimentsPath(projectID), input)
	return result, err
}

func (c *PosthogClient) GetExperiment(ctx context.Context, projectID, id string) (Experiment, HTTPStatusCode, error) {
	return doGet[Experiment](c, ctx, experimentPath(projectID, id))
}

func (c *PosthogClient) UpdateExperiment(ctx context.Context, projectID, id string, input ExperimentRequest) (Experiment, HTTPStatusCode, error) {
	return doPatch[Experiment](c, ctx, experimentPath(projectID, id), input)
}

// DeleteExperiment soft-deletes the experiment. Hard DELETE is forbidden by the API (405);
// deletion is a PATCH setting deleted=true, mirroring DeleteFeatureFlag. The backing feature
// flag is left in place.
func (c *PosthogClient) DeleteExperiment(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	deleted := true
	_, statusCode, err := c.UpdateExperiment(ctx, projectID, id, ExperimentRequest{Deleted: &deleted})
	return statusCode, err
}

// LaunchExperiment transitions a draft to running (POST .../launch/).
func (c *PosthogClient) LaunchExperiment(ctx context.Context, projectID, id string) (Experiment, HTTPStatusCode, error) {
	return doPost[Experiment](c, ctx, experimentActionPath(projectID, id, "launch"), struct{}{})
}

// PauseExperiment deactivates the flag of a running experiment (POST .../pause/).
func (c *PosthogClient) PauseExperiment(ctx context.Context, projectID, id string) (Experiment, HTTPStatusCode, error) {
	return doPost[Experiment](c, ctx, experimentActionPath(projectID, id, "pause"), struct{}{})
}

// ResumeExperiment reactivates the flag of a paused experiment (POST .../resume/).
func (c *PosthogClient) ResumeExperiment(ctx context.Context, projectID, id string) (Experiment, HTTPStatusCode, error) {
	return doPost[Experiment](c, ctx, experimentActionPath(projectID, id, "resume"), struct{}{})
}

// EndExperiment stops a running experiment without shipping a variant (POST .../end/).
func (c *PosthogClient) EndExperiment(ctx context.Context, projectID, id string, input ExperimentEndRequest) (Experiment, HTTPStatusCode, error) {
	return doPost[Experiment](c, ctx, experimentActionPath(projectID, id, "end"), input)
}

// ShipVariant ends the experiment (if running) and rewrites the flag so the chosen variant
// gets 100% (POST .../ship_variant/).
func (c *PosthogClient) ShipVariant(ctx context.Context, projectID, id string, input ExperimentShipVariantRequest) (Experiment, HTTPStatusCode, error) {
	return doPost[Experiment](c, ctx, experimentActionPath(projectID, id, "ship_variant"), input)
}

package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
)

// InsightVariable is a SQL variable. Queries reference it as
// {variables.code_name}, which renders an input on any dashboard holding an
// insight that uses it.
//
// DefaultValue and Values are arbitrary JSON: a scalar for most variable
// types, a string array for List variables. json.RawMessage keeps the payload
// intact so the resource can hand the user's own JSON back to state unchanged.
type InsightVariable struct {
	ID           string           `json:"id"`
	Name         *string          `json:"name,omitempty"`
	Type         *string          `json:"type,omitempty"`
	CodeName     *string          `json:"code_name,omitempty"`
	DefaultValue *json.RawMessage `json:"default_value,omitempty"`
	Values       *json.RawMessage `json:"values,omitempty"`
	CreatedAt    *string          `json:"created_at,omitempty"`
}

type InsightVariableRequest struct {
	Name         *string          `json:"name,omitempty"`
	Type         *string          `json:"type,omitempty"`
	DefaultValue *json.RawMessage `json:"default_value,omitempty"`
	Values       *json.RawMessage `json:"values,omitempty"`
}

func (c *PosthogClient) CreateInsightVariable(ctx context.Context, projectID string, input InsightVariableRequest) (InsightVariable, error) {
	path := fmt.Sprintf("/api/projects/%s/insight_variables/", projectID)
	result, _, err := doPost[InsightVariable](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetInsightVariable(ctx context.Context, projectID, id string) (InsightVariable, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/insight_variables/%s/", projectID, id)
	return doGet[InsightVariable](c, ctx, path)
}

func (c *PosthogClient) UpdateInsightVariable(ctx context.Context, projectID, id string, input InsightVariableRequest) (InsightVariable, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/insight_variables/%s/", projectID, id)
	return doPatch[InsightVariable](c, ctx, path, input)
}

func (c *PosthogClient) DeleteInsightVariable(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/insight_variables/%s/", projectID, id)
	return doDelete(c, ctx, path)
}

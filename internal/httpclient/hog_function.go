package httpclient

import (
	"context"
	"fmt"
)

type HogFunction struct {
	ID             string                   `json:"id"`
	Type           *string                  `json:"type,omitempty"`
	Name           *string                  `json:"name,omitempty"`
	Description    *string                  `json:"description,omitempty"`
	Enabled        *bool                    `json:"enabled,omitempty"`
	Hog            *string                  `json:"hog,omitempty"`
	InputsSchema   []map[string]interface{} `json:"inputs_schema,omitempty"`
	Inputs         map[string]interface{}   `json:"inputs,omitempty"`
	Filters        map[string]interface{}   `json:"filters,omitempty"`
	Masking        map[string]interface{}   `json:"masking,omitempty"`
	Mappings       []map[string]interface{} `json:"mappings,omitempty"`
	IconURL        *string                  `json:"icon_url,omitempty"`
	Template       *HogFunctionTemplate     `json:"template,omitempty"`
	ExecutionOrder *int                     `json:"execution_order,omitempty"`
	CreatedAt      *string                  `json:"created_at,omitempty"`
	CreatedBy      map[string]interface{}   `json:"created_by,omitempty"`
	UpdatedAt      *string                  `json:"updated_at,omitempty"`
	Status         *HogFunctionStatus       `json:"status,omitempty"`
}

type HogFunctionRequest struct {
	Type           *string                  `json:"type,omitempty"`
	Name           *string                  `json:"name,omitempty"`
	Description    *string                  `json:"description,omitempty"`
	Enabled        *bool                    `json:"enabled,omitempty"`
	Hog            *string                  `json:"hog,omitempty"`
	InputsSchema   []map[string]interface{} `json:"inputs_schema,omitempty"`
	Inputs         map[string]interface{}   `json:"inputs,omitempty"`
	Filters        map[string]interface{}   `json:"filters,omitempty"`
	Masking        map[string]interface{}   `json:"masking,omitempty"`
	Mappings       []map[string]interface{} `json:"mappings,omitempty"`
	IconURL        *string                  `json:"icon_url,omitempty"`
	TemplateID     *string                  `json:"template_id,omitempty"`
	ExecutionOrder *int                     `json:"execution_order,omitempty"`
	Deleted        *bool                    `json:"deleted,omitempty"`
}

type HogFunctionTemplate struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Type        string  `json:"type"`
	Status      *string `json:"status,omitempty"`
	IconURL     *string `json:"icon_url,omitempty"`
}

type HogFunctionStatus struct {
	State  int `json:"state"`
	Tokens int `json:"tokens"`
}

func (c *PosthogClient) CreateHogFunction(ctx context.Context, projectID string, input HogFunctionRequest) (HogFunction, error) {
	path := fmt.Sprintf("/api/environments/%s/hog_functions/", projectID)
	result, _, err := doPost[HogFunction](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetHogFunction(ctx context.Context, projectID, id string) (HogFunction, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/environments/%s/hog_functions/%s/", projectID, id)
	return doGet[HogFunction](c, ctx, path)
}

func (c *PosthogClient) UpdateHogFunction(ctx context.Context, projectID, id string, input HogFunctionRequest) (HogFunction, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/environments/%s/hog_functions/%s/", projectID, id)
	return doPatch[HogFunction](c, ctx, path, input)
}

// DeleteHogFunction performs a soft delete by setting deleted=true via PATCH.
// The HogFunction API does not support hard deletes
func (c *PosthogClient) DeleteHogFunction(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/environments/%s/hog_functions/%s/", projectID, id)
	deleted := true
	req := HogFunctionRequest{Deleted: &deleted}
	_, status, err := doPatch[HogFunction](c, ctx, path, req)
	return status, err
}

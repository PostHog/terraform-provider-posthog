package httpclient

import (
	"context"
	"fmt"
)

// SchemaPropertyGroupProperty is one typed property inside a schema property group.
// ID is server-assigned; requests omit it, which makes the API treat updates as
// replace-all (existing properties whose ids are absent get deleted, incoming
// entries without ids get created — transactionally).
type SchemaPropertyGroupProperty struct {
	ID                *string `json:"id,omitempty"`
	Name              string  `json:"name"`
	PropertyType      string  `json:"property_type"`
	IsRequired        *bool   `json:"is_required,omitempty"`
	IsOptionalInTypes *bool   `json:"is_optional_in_types,omitempty"`
	Description       *string `json:"description,omitempty"`
}

type SchemaPropertyGroup struct {
	ID          string                        `json:"id"`
	Name        string                        `json:"name"`
	Description *string                       `json:"description,omitempty"`
	Properties  []SchemaPropertyGroupProperty `json:"properties,omitempty"`
	CreatedAt   *string                       `json:"created_at,omitempty"`
	UpdatedAt   *string                       `json:"updated_at,omitempty"`
}

// SchemaPropertyGroupRequest uses a pointer-to-slice for Properties so callers
// can distinguish "leave untouched" (nil) from "delete all" (empty slice).
type SchemaPropertyGroupRequest struct {
	Name        *string                        `json:"name,omitempty"`
	Description *string                        `json:"description,omitempty"`
	Properties  *[]SchemaPropertyGroupProperty `json:"properties,omitempty"`
}

func (c *PosthogClient) CreateSchemaPropertyGroup(ctx context.Context, projectID string, input SchemaPropertyGroupRequest) (SchemaPropertyGroup, error) {
	path := fmt.Sprintf("/api/projects/%s/schema_property_groups/", projectID)
	result, _, err := doPost[SchemaPropertyGroup](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetSchemaPropertyGroup(ctx context.Context, projectID, id string) (SchemaPropertyGroup, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/schema_property_groups/%s/", projectID, id)
	return doGet[SchemaPropertyGroup](c, ctx, path)
}

func (c *PosthogClient) UpdateSchemaPropertyGroup(ctx context.Context, projectID, id string, input SchemaPropertyGroupRequest) (SchemaPropertyGroup, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/schema_property_groups/%s/", projectID, id)
	return doPatch[SchemaPropertyGroup](c, ctx, path, input)
}

func (c *PosthogClient) DeleteSchemaPropertyGroup(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/schema_property_groups/%s/", projectID, id)
	return doDelete(c, ctx, path)
}

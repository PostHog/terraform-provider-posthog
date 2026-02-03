package httpclient

import (
	"context"
	"fmt"
)

type AccessControl struct {
	AccessLevel        *string `json:"access_level"`
	Resource           string  `json:"resource"`
	ResourceID         *string `json:"resource_id,omitempty"`
	OrganizationMember *string `json:"organization_member,omitempty"`
	Role               *string `json:"role,omitempty"`
	CreatedBy          *string `json:"created_by,omitempty"`
	CreatedAt          *string `json:"created_at,omitempty"`
	UpdatedAt          *string `json:"updated_at,omitempty"`
}

type AccessControlRequest struct {
	AccessLevel        *string `json:"access_level"`
	Resource           string  `json:"resource"`
	ResourceID         *string `json:"resource_id,omitempty"`
	OrganizationMember *string `json:"organization_member,omitempty"`
	Role               *string `json:"role,omitempty"`
}

type AccessControlListResponse struct {
	AccessControls        []AccessControl `json:"access_controls"`
	AvailableAccessLevels []string        `json:"available_access_levels"`
	DefaultAccessLevel    string          `json:"default_access_level"`
}

// UpsertAccessControl creates or updates an access control.
// If resourceID is provided, it sets permissions for that specific resource instance.
// Otherwise, it sets permissions at the resource-type level (all resources of that type).
func (c *PosthogClient) UpsertAccessControl(ctx context.Context, projectID string, input AccessControlRequest) (AccessControl, error) {
	path := c.accessControlPath(projectID, input.Resource, input.ResourceID)
	result, _, err := doPut[AccessControl](c, ctx, path, input)
	return result, err
}

// GetAccessControls returns access controls for a project.
// If resourceID is provided, returns controls for that specific resource.
// Otherwise, returns resource-type level controls.
func (c *PosthogClient) GetAccessControls(ctx context.Context, projectID, resourceType string, resourceID *string) (AccessControlListResponse, HTTPStatusCode, error) {
	path := c.accessControlPath(projectID, resourceType, resourceID)
	return doGet[AccessControlListResponse](c, ctx, path)
}

func (c *PosthogClient) DeleteAccessControl(ctx context.Context, projectID string, input AccessControlRequest) (HTTPStatusCode, error) {
	path := c.accessControlPath(projectID, input.Resource, input.ResourceID)
	input.AccessLevel = nil
	return doPutNoContent(c, ctx, path, input)
}

// accessControlPath returns the appropriate API path based on whether resourceID is set.
func (c *PosthogClient) accessControlPath(projectID, resourceType string, resourceID *string) string {
	if resourceID != nil && *resourceID != "" {
		// Instance-level: /api/projects/{project}/{resource_type}/{resource_id}/access_controls/
		return fmt.Sprintf("/api/projects/%s/%s/%s/access_controls/", projectID, resourceType, *resourceID)
	}
	// Resource-Type-level: /api/projects/{project}/resource_access_controls/
	return fmt.Sprintf("/api/projects/%s/resource_access_controls/", projectID)
}

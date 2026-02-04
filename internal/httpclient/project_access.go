package httpclient

import (
	"context"
	"fmt"

	"github.com/posthog/terraform-provider/internal/util"
)

type ProjectAccessControl struct {
	AccessLevel        *string `json:"access_level"`
	Resource           string  `json:"resource"`
	ResourceID         *string `json:"resource_id,omitempty"`
	OrganizationMember *string `json:"organization_member,omitempty"`
	Role               *string `json:"role,omitempty"`
	CreatedBy          *string `json:"created_by,omitempty"`
	CreatedAt          *string `json:"created_at,omitempty"`
	UpdatedAt          *string `json:"updated_at,omitempty"`
}

// BuildProjectMemberCompositeID constructs a unique identifier for project member access.
// Format: project_id/role/<role_id> or project_id/member/<member_id>
func BuildProjectMemberCompositeID(projectID, role, organizationMember string) string {
	if role != "" {
		return fmt.Sprintf("%s/role/%s", projectID, role)
	}
	if organizationMember != "" {
		return fmt.Sprintf("%s/member/%s", projectID, organizationMember)
	}
	return ""
}

// BuildCompositeID constructs a unique identifier for this project access control.
func (ac ProjectAccessControl) BuildCompositeID(projectID string) string {
	return BuildProjectMemberCompositeID(projectID, util.PtrToString(ac.Role), util.PtrToString(ac.OrganizationMember))
}

type ProjectAccessControlRequest struct {
	AccessLevel        *string `json:"access_level"`
	OrganizationMember *string `json:"organization_member,omitempty"`
	Role               *string `json:"role,omitempty"`
}

type ProjectAccessControlListResponse struct {
	AccessControls          []ProjectAccessControl `json:"access_controls"`
	AvailableAccessLevels   []string               `json:"available_access_levels"`
	DefaultAccessLevel      string                 `json:"default_access_level"`
	UserAccessLevel         string                 `json:"user_access_level"`
	UserCanEditAccessLevels bool                   `json:"user_can_edit_access_levels"`
}

func (c *PosthogClient) GetProjectAccessControls(ctx context.Context, projectID string) (ProjectAccessControlListResponse, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/access_controls/", projectID)
	return doGet[ProjectAccessControlListResponse](c, ctx, path)
}

// UpsertProjectAccessControl creates or updates a project-level access control.
// This can set:
// - Default access level (when role and organization_member are nil)
// - Role access (when role is set)
// - Member access (when organization_member is set)
func (c *PosthogClient) UpsertProjectAccessControl(ctx context.Context, projectID string, input ProjectAccessControlRequest) (ProjectAccessControl, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/access_controls/", projectID)
	return doPut[ProjectAccessControl](c, ctx, path, input)
}

// DeleteProjectAccessControl removes a project-level access control by setting access_level to nil.
func (c *PosthogClient) DeleteProjectAccessControl(ctx context.Context, projectID string, input ProjectAccessControlRequest) (HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/access_controls/", projectID)
	input.AccessLevel = nil
	return doPutNoContent(c, ctx, path, input)
}

package httpclient

import (
	"context"
	"fmt"
)

// ManagerTerraform is the value PostHog stores on rules this provider owns.
const ManagerTerraform = "terraform"

// AccessControlClaimRequest addresses one rule the same way its unique constraint does.
// A nil ManagedBy releases the rule instead of claiming it.
type AccessControlClaimRequest struct {
	Resource           string  `json:"resource"`
	ResourceID         *string `json:"resource_id,omitempty"`
	Role               *string `json:"role,omitempty"`
	OrganizationMember *string `json:"organization_member,omitempty"`
	ManagedBy          *string `json:"managed_by"`
}

type AccessControlClaimResponse struct {
	ManagedBy *string `json:"managed_by"`
	ManagedAt *string `json:"managed_at"`
}

// ClaimAccessControl tells PostHog that this provider owns the rule.
//
// Terraform only writes a rule when that rule changed, so a rule that is already correct is
// never written and would never be recorded as ours. Refresh is the one point that touches every
// managed rule on every plan, so ownership is asserted from Read instead of from the writes.
func (c *PosthogClient) ClaimAccessControl(ctx context.Context, projectID string, input AccessControlClaimRequest) (AccessControlClaimResponse, HTTPStatusCode, error) {
	managedBy := ManagerTerraform
	input.ManagedBy = &managedBy
	return c.setAccessControlManager(ctx, projectID, input)
}

// ReleaseAccessControl gives the rule back, so PostHog stops refusing changes to it.
func (c *PosthogClient) ReleaseAccessControl(ctx context.Context, projectID string, input AccessControlClaimRequest) (AccessControlClaimResponse, HTTPStatusCode, error) {
	input.ManagedBy = nil
	return c.setAccessControlManager(ctx, projectID, input)
}

func (c *PosthogClient) setAccessControlManager(ctx context.Context, projectID string, input AccessControlClaimRequest) (AccessControlClaimResponse, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/access_control_management/", projectID)
	return doPut[AccessControlClaimResponse](c, ctx, path, input)
}

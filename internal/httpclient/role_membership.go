package httpclient

import (
	"context"
	"fmt"
)

type RoleMembershipUser struct {
	UUID  string  `json:"uuid"`
	Email *string `json:"email,omitempty"`
}

type RoleMembership struct {
	ID       string              `json:"id"`
	RoleID   string              `json:"role_id"`
	User     *RoleMembershipUser `json:"user,omitempty"`
	JoinedAt *string             `json:"joined_at,omitempty"`
}

type RoleMembershipRequest struct {
	UserUUID string `json:"user_uuid"`
}

func (c *PosthogClient) CreateRoleMembership(ctx context.Context, organizationID, roleID string, input RoleMembershipRequest) (RoleMembership, error) {
	organizationID, err := c.ResolveOrganizationID(ctx, organizationID)
	if err != nil {
		return RoleMembership{}, err
	}
	path := fmt.Sprintf("/api/organizations/%s/roles/%s/role_memberships/", organizationID, roleID)
	result, _, err := doPost[RoleMembership](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetRoleMembership(ctx context.Context, organizationID, roleID, membershipID string) (RoleMembership, HTTPStatusCode, error) {
	organizationID, err := c.ResolveOrganizationID(ctx, organizationID)
	if err != nil {
		return RoleMembership{}, 0, err
	}
	path := fmt.Sprintf("/api/organizations/%s/roles/%s/role_memberships/%s/", organizationID, roleID, membershipID)
	return doGet[RoleMembership](c, ctx, path)
}

func (c *PosthogClient) DeleteRoleMembership(ctx context.Context, organizationID, roleID, membershipID string) (HTTPStatusCode, error) {
	organizationID, err := c.ResolveOrganizationID(ctx, organizationID)
	if err != nil {
		return 0, err
	}
	path := fmt.Sprintf("/api/organizations/%s/roles/%s/role_memberships/%s/", organizationID, roleID, membershipID)
	return doDelete(c, ctx, path)
}

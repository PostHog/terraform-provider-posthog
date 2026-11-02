package httpclient

import (
	"context"
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

func (c *PosthogClient) CreateRoleMembership(ctx context.Context, orgIDOrSlug, roleID string, input RoleMembershipRequest) (RoleMembership, error) {
	path, err := c.orgPath(ctx, orgIDOrSlug, "/api/organizations/%s/roles/%s/role_memberships/", roleID)
	if err != nil {
		return RoleMembership{}, err
	}
	result, _, err := doPost[RoleMembership](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetRoleMembership(ctx context.Context, orgIDOrSlug, roleID, membershipID string) (RoleMembership, HTTPStatusCode, error) {
	path, err := c.orgPath(ctx, orgIDOrSlug, "/api/organizations/%s/roles/%s/role_memberships/%s/", roleID, membershipID)
	if err != nil {
		return RoleMembership{}, 0, err
	}
	return doGet[RoleMembership](c, ctx, path)
}

func (c *PosthogClient) DeleteRoleMembership(ctx context.Context, orgIDOrSlug, roleID, membershipID string) (HTTPStatusCode, error) {
	path, err := c.orgPath(ctx, orgIDOrSlug, "/api/organizations/%s/roles/%s/role_memberships/%s/", roleID, membershipID)
	if err != nil {
		return 0, err
	}
	return doDelete(c, ctx, path)
}

package httpclient

import (
	"context"
)

type Role struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	CreatedAt *string `json:"created_at,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
}

type RoleRequest struct {
	Name string `json:"name"`
}

func (c *PosthogClient) ListRoles(ctx context.Context, orgIDOrSlug string) ([]Role, error) {
	path, err := c.orgPath(ctx, orgIDOrSlug, "/api/organizations/%s/roles/")
	if err != nil {
		return nil, err
	}
	return listAll[Role](c, ctx, path)
}

func (c *PosthogClient) CreateRole(ctx context.Context, orgIDOrSlug string, input RoleRequest) (Role, error) {
	path, err := c.orgPath(ctx, orgIDOrSlug, "/api/organizations/%s/roles/")
	if err != nil {
		return Role{}, err
	}
	result, _, err := doPost[Role](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetRole(ctx context.Context, orgIDOrSlug, roleID string) (Role, HTTPStatusCode, error) {
	path, err := c.orgPath(ctx, orgIDOrSlug, "/api/organizations/%s/roles/%s/", roleID)
	if err != nil {
		return Role{}, 0, err
	}
	return doGet[Role](c, ctx, path)
}

func (c *PosthogClient) UpdateRole(ctx context.Context, orgIDOrSlug, roleID string, input RoleRequest) (Role, HTTPStatusCode, error) {
	path, err := c.orgPath(ctx, orgIDOrSlug, "/api/organizations/%s/roles/%s/", roleID)
	if err != nil {
		return Role{}, 0, err
	}
	return doPatch[Role](c, ctx, path, input)
}

func (c *PosthogClient) DeleteRole(ctx context.Context, orgIDOrSlug, roleID string) (HTTPStatusCode, error) {
	path, err := c.orgPath(ctx, orgIDOrSlug, "/api/organizations/%s/roles/%s/", roleID)
	if err != nil {
		return 0, err
	}
	return doDelete(c, ctx, path)
}

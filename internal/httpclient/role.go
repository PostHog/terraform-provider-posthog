package httpclient

import (
	"context"
	"fmt"
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

func (c *PosthogClient) CreateRole(ctx context.Context, organizationID string, input RoleRequest) (Role, error) {
	path := fmt.Sprintf("/api/organizations/%s/roles/", organizationID)
	result, _, err := doPost[Role](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetRole(ctx context.Context, organizationID, roleID string) (Role, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/organizations/%s/roles/%s/", organizationID, roleID)
	return doGet[Role](c, ctx, path)
}

func (c *PosthogClient) UpdateRole(ctx context.Context, organizationID, roleID string, input RoleRequest) (Role, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/organizations/%s/roles/%s/", organizationID, roleID)
	return doPatch[Role](c, ctx, path, input)
}

func (c *PosthogClient) DeleteRole(ctx context.Context, organizationID, roleID string) (HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/organizations/%s/roles/%s/", organizationID, roleID)
	return doDelete(c, ctx, path)
}

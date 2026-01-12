package httpclient

import (
	"context"
	"fmt"
)

type Project struct {
	ID             int64   `json:"id"`
	Name           *string `json:"name,omitempty"`
	OrganizationID *string `json:"organization,omitempty"`
	APIToken       *string `json:"api_token,omitempty"`
	Timezone       *string `json:"timezone,omitempty"`
}

type ProjectRequest struct {
	Name     *string `json:"name,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}

func (c *PosthogClient) CreateProject(ctx context.Context, organizationID string, input ProjectRequest) (Project, error) {
	path := fmt.Sprintf("/api/organizations/%s/projects/", organizationID)
	result, _, err := doPost[Project](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetProject(ctx context.Context, organizationID, projectID string) (Project, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/organizations/%s/projects/%s/", organizationID, projectID)
	return doGet[Project](c, ctx, path)
}

func (c *PosthogClient) UpdateProject(ctx context.Context, organizationID, projectID string, input ProjectRequest) (Project, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/organizations/%s/projects/%s/", organizationID, projectID)
	return doPatch[Project](c, ctx, path, input)
}

func (c *PosthogClient) DeleteProject(ctx context.Context, organizationID, projectID string) (HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/organizations/%s/projects/%s/", organizationID, projectID)
	return doDelete(c, ctx, path)
}

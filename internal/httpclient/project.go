package httpclient

import (
	"context"
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

func (c *PosthogClient) CreateProject(ctx context.Context, orgIDOrSlug string, input ProjectRequest) (Project, error) {
	path, err := c.orgPath(ctx, orgIDOrSlug, "/api/organizations/%s/projects/")
	if err != nil {
		return Project{}, err
	}
	result, _, err := doPost[Project](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetProject(ctx context.Context, orgIDOrSlug, projectID string) (Project, HTTPStatusCode, error) {
	path, err := c.orgPath(ctx, orgIDOrSlug, "/api/organizations/%s/projects/%s/", projectID)
	if err != nil {
		return Project{}, 0, err
	}
	return doGet[Project](c, ctx, path)
}

func (c *PosthogClient) UpdateProject(ctx context.Context, orgIDOrSlug, projectID string, input ProjectRequest) (Project, HTTPStatusCode, error) {
	path, err := c.orgPath(ctx, orgIDOrSlug, "/api/organizations/%s/projects/%s/", projectID)
	if err != nil {
		return Project{}, 0, err
	}
	return doPatch[Project](c, ctx, path, input)
}

func (c *PosthogClient) DeleteProject(ctx context.Context, orgIDOrSlug, projectID string) (HTTPStatusCode, error) {
	path, err := c.orgPath(ctx, orgIDOrSlug, "/api/organizations/%s/projects/%s/", projectID)
	if err != nil {
		return 0, err
	}
	return doDelete(c, ctx, path)
}

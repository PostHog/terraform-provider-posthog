package httpclient

import (
	"context"
	"fmt"
)

type Action struct {
	ID                 int64                    `json:"id"`
	Name               string                   `json:"name"`
	Description        *string                  `json:"description,omitempty"`
	Tags               []string                 `json:"tags,omitempty"`
	PostToSlack        *bool                    `json:"post_to_slack,omitempty"`
	SlackMessageFormat *string                  `json:"slack_message_format,omitempty"`
	Steps              []map[string]interface{} `json:"steps,omitempty"`
	Deleted            *bool                    `json:"deleted,omitempty"`
}

type ActionRequest struct {
	Name               string                   `json:"name"`
	Description        *string                  `json:"description,omitempty"`
	Tags               []string                 `json:"tags,omitempty"`
	PostToSlack        *bool                    `json:"post_to_slack,omitempty"`
	SlackMessageFormat *string                  `json:"slack_message_format,omitempty"`
	Steps              []map[string]interface{} `json:"steps,omitempty"`
	Deleted            *bool                    `json:"deleted,omitempty"`
}

func (c *PosthogClient) CreateAction(ctx context.Context, projectID string, input ActionRequest) (Action, error) {
	path := fmt.Sprintf("/api/projects/%s/actions/", projectID)
	result, _, err := doPost[Action](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetAction(ctx context.Context, projectID, id string) (Action, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/actions/%s/", projectID, id)
	return doGet[Action](c, ctx, path)
}

func (c *PosthogClient) UpdateAction(ctx context.Context, projectID, id string, input ActionRequest) (Action, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/projects/%s/actions/%s/", projectID, id)
	return doPatch[Action](c, ctx, path, input)
}

func (c *PosthogClient) DeleteAction(ctx context.Context, projectID, id string) (HTTPStatusCode, error) {
	existing, _, err := c.GetAction(ctx, projectID, id)
	if err != nil {
		return 0, fmt.Errorf("failed to get action for deletion: %w", err)
	}

	deleted := true
	input := ActionRequest{
		Name:    existing.Name,
		Deleted: &deleted,
	}
	_, statusCode, err := c.UpdateAction(ctx, projectID, id, input)
	return statusCode, err
}

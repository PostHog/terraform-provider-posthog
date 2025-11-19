package posthog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Insight struct {
	ID             int64           `json:"id"`
	Name           string          `json:"name"`
	DerivedName    string          `json:"derived_name"`
	Description    string          `json:"description"`
	Query          json.RawMessage `json:"query"`
	Tags           []string        `json:"tags"`
	CreateInFolder string          `json:"_create_in_folder"`
}

type InsightRequest struct {
	Name           string          `json:"name,omitempty"`
	DerivedName    string          `json:"derived_name,omitempty"`
	Description    string          `json:"description,omitempty"`
	Query          json.RawMessage `json:"query"`
	Tags           []string        `json:"tags,omitempty"`
	Saved          bool            `json:"saved,default=1"`
	CreateInFolder string          `json:"_create_in_folder,omitempty"`
}

func (c *DefaultClient) CreateInsight(ctx context.Context, input InsightRequest) (Insight, error) {
	endpoint, err := c.insightCollectionURL()
	if err != nil {
		return Insight{}, err
	}

	payload, err := json.Marshal(input)
	if err != nil {
		return Insight{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(payload))
	if err != nil {
		return Insight{}, fmt.Errorf("failed to build request: %w", err)
	}

	body, err := c.doRequestAndReadBody(req)
	if err != nil {
		return Insight{}, fmt.Errorf("failed to create insight: %w", err)
	}

	var insight Insight
	if len(body) > 0 {
		if err := json.Unmarshal(body, &insight); err != nil {
			return Insight{}, fmt.Errorf("failed to decode create insight response: %w", err)
		}
	}

	return insight, nil
}

func (c *DefaultClient) GetInsight(ctx context.Context, id int64) (Insight, error) {
	endpoint, err := c.insightDetailURL(id)
	if err != nil {
		return Insight{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, http.NoBody)
	if err != nil {
		return Insight{}, fmt.Errorf("failed to build request: %w", err)
	}

	body, err := c.doRequestAndReadBody(req)
	if err != nil {
		return Insight{}, fmt.Errorf("failed to retrieve insight: %w", err)
	}

	var insight Insight
	if len(body) > 0 {
		if err := json.Unmarshal(body, &insight); err != nil {
			return Insight{}, fmt.Errorf("failed to decode insight response: %w", err)
		}
	}

	return insight, nil
}

func (c *DefaultClient) UpdateInsight(ctx context.Context, id int64, input InsightRequest) (Insight, error) {
	endpoint, err := c.insightDetailURL(id)
	if err != nil {
		return Insight{}, err
	}

	payload, err := json.Marshal(input)
	if err != nil {
		return Insight{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, endpoint, bytes.NewReader(payload))
	if err != nil {
		return Insight{}, fmt.Errorf("failed to build request: %w", err)
	}

	body, err := c.doRequestAndReadBody(req)
	if err != nil {
		return Insight{}, fmt.Errorf("failed to update insight: %w", err)
	}

	var insight Insight
	if len(body) > 0 {
		if err := json.Unmarshal(body, &insight); err != nil {
			return Insight{}, fmt.Errorf("failed to decode update insight response: %w", err)
		}
	}

	return insight, nil
}

func (c *DefaultClient) DeleteInsight(ctx context.Context, id int64) error {
	endpoint, err := c.insightDetailURL(id)
	if err != nil {
		return err
	}

	payload := map[string]any{"deleted": true}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal delete payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, endpoint, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}

	if _, err := c.doRequestAndReadBody(req); err != nil {
		return fmt.Errorf("failed to delete insight: %w", err)
	}

	return nil
}

func (c *DefaultClient) insightCollectionURL() (string, error) {
	return fmt.Sprintf("%s/api/projects/%s/insights/", c.host, c.projectId), nil
}

func (c *DefaultClient) insightDetailURL(id int64) (string, error) {
	collection, err := c.insightCollectionURL()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%d/", collection, id), nil
}

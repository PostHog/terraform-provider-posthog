package posthog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

const (
	headerAuthorization = "Authorization"
	headerContentType   = "Content-Type"
)

type Client interface {
	CreateFeatureFlag(ctx context.Context, input FeatureFlagRequest) (FeatureFlag, error)
	GetFeatureFlag(ctx context.Context, id int64) (FeatureFlag, error)
	UpdateFeatureFlag(ctx context.Context, id int64, input FeatureFlagRequest) (FeatureFlag, error)
	DeleteFeatureFlag(ctx context.Context, id int64) error
}

type DefaultClient struct {
	host      string
	apiKey    string
	projectId string

	httpClient *http.Client
	logger     *slog.Logger
}

func NewDefaultClient(logger *slog.Logger, host, apiKey, projectId string) Client {
	return NewClient(&http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        10,
			IdleConnTimeout:     30 * time.Second,
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}, logger, host, apiKey, projectId)
}

// NewClient creates a new HTTP client
// baseUrl should NOT end with a trailing slash
func NewClient(client *http.Client, logger *slog.Logger, host, apiKey, projectId string) Client {
	return &DefaultClient{
		host:       host,
		apiKey:     apiKey,
		projectId:  projectId,
		httpClient: client,
		logger:     logger,
	}
}

func (c *DefaultClient) setCommonHeaders(req *http.Request) *http.Request {
	req.Header.Set(headerAuthorization, fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set(headerContentType, "application/json")

	return req
}

// doRequestAndReadBody sends a request and reads the body of the response, it also closes
// the respective reader so that callees do not have to worry about that.
// If we receive a non 200 status code, an error is returned.
func (c *DefaultClient) doRequestAndReadBody(req *http.Request) ([]byte, error) {
	logger := c.logger.With(slog.Any("uri", req.URL.String()))

	resp, err := c.httpClient.Do(c.setCommonHeaders(req))
	if err != nil {
		return []byte{}, fmt.Errorf("failed to send request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			logger.Error("failed to close response body", slog.Any("error", err))
		}
	}()
	logger.Debug("received response", slog.Any("status", resp.StatusCode))

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return []byte{}, fmt.Errorf("failed to get success response, received status code: %d, body: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Info("received an invalid response", slog.Any("body", string(body)))
		return []byte{}, fmt.Errorf("failed to read response body: %w", err)
	}
	logger.Debug("received a valid response", slog.Any("body", string(body)))

	return body, nil
}

// FeatureFlag represents a PostHog feature flag
type FeatureFlag struct {
	ID                int64                  `json:"id"`
	Key               string                 `json:"key"`
	Name              *string                `json:"name,omitempty"`
	Active            *bool                  `json:"active,omitempty"`
	Filters           map[string]interface{} `json:"filters,omitempty"`
	RolloutPercentage *int32                 `json:"rollout_percentage,omitempty"`
	Tags              []string               `json:"tags,omitempty"`
	Deleted           *bool                  `json:"deleted,omitempty"`
}

// FeatureFlagRequest represents the request body for creating/updating feature flags
type FeatureFlagRequest struct {
	Key     string                 `json:"key"`
	Name    *string                `json:"name,omitempty"`
	Active  *bool                  `json:"active,omitempty"`
	Filters map[string]interface{} `json:"filters,omitempty"`
	Tags    []string               `json:"tags,omitempty"`
	Deleted *bool                  `json:"deleted,omitempty"`
}

// CreateFeatureFlag creates a new feature flag
func (c *DefaultClient) CreateFeatureFlag(ctx context.Context, input FeatureFlagRequest) (FeatureFlag, error) {
	url := fmt.Sprintf("%s/api/projects/%s/feature_flags/", c.host, c.projectId)

	body, err := json.Marshal(input)
	if err != nil {
		return FeatureFlag{}, fmt.Errorf("failed to marshal feature flag request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return FeatureFlag{}, fmt.Errorf("failed to create request: %w", err)
	}

	respBody, err := c.doRequestAndReadBody(req)
	if err != nil {
		return FeatureFlag{}, err
	}

	var flag FeatureFlag
	if err := json.Unmarshal(respBody, &flag); err != nil {
		return FeatureFlag{}, fmt.Errorf("failed to unmarshal feature flag response: %w", err)
	}

	return flag, nil
}

// GetFeatureFlag retrieves a feature flag by ID
func (c *DefaultClient) GetFeatureFlag(ctx context.Context, id int64) (FeatureFlag, error) {
	url := fmt.Sprintf("%s/api/projects/%s/feature_flags/%d/", c.host, c.projectId, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return FeatureFlag{}, fmt.Errorf("failed to create request: %w", err)
	}

	respBody, err := c.doRequestAndReadBody(req)
	if err != nil {
		return FeatureFlag{}, err
	}

	var flag FeatureFlag
	if err := json.Unmarshal(respBody, &flag); err != nil {
		return FeatureFlag{}, fmt.Errorf("failed to unmarshal feature flag response: %w", err)
	}

	return flag, nil
}

// UpdateFeatureFlag updates an existing feature flag
func (c *DefaultClient) UpdateFeatureFlag(ctx context.Context, id int64, input FeatureFlagRequest) (FeatureFlag, error) {
	url := fmt.Sprintf("%s/api/projects/%s/feature_flags/%d/", c.host, c.projectId, id)

	body, err := json.Marshal(input)
	if err != nil {
		return FeatureFlag{}, fmt.Errorf("failed to marshal feature flag request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, bytes.NewReader(body))
	if err != nil {
		return FeatureFlag{}, fmt.Errorf("failed to create request: %w", err)
	}

	respBody, err := c.doRequestAndReadBody(req)
	if err != nil {
		return FeatureFlag{}, err
	}

	var flag FeatureFlag
	if err := json.Unmarshal(respBody, &flag); err != nil {
		return FeatureFlag{}, fmt.Errorf("failed to unmarshal feature flag response: %w", err)
	}

	return flag, nil
}

// DeleteFeatureFlag soft-deletes a feature flag
func (c *DefaultClient) DeleteFeatureFlag(ctx context.Context, id int64) error {
	url := fmt.Sprintf("%s/api/projects/%s/feature_flags/%d/", c.host, c.projectId, id)

	deleted := true
	input := FeatureFlagRequest{Deleted: &deleted}
	body, err := json.Marshal(input)
	if err != nil {
		return fmt.Errorf("failed to marshal delete request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	_, err = c.doRequestAndReadBody(req)
	return err
}

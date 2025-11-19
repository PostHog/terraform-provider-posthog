package posthog

import (
	"bytes"
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
	CreateDashboard(params CreateDashboardParams) (Dashboard, error)
	GetDashboard(id string) (Dashboard, error)
	UpdateDashboard(id string, params UpdateDashboardParams) (Dashboard, error)
	DeleteDashboard(id string) error
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

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("failed to get success response, received status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Info("received an invalid response", slog.Any("body", string(body)))
		return []byte{}, fmt.Errorf("failed to read response body: %w", err)
	}
	logger.Debug("received a valid response", slog.Any("body", string(body)))

	return body, nil
}

// Dashboard types
type Dashboard struct {
	ID          int32    `json:"id"`
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Pinned      *bool    `json:"pinned,omitempty"`
	Deleted     *bool    `json:"deleted,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

type CreateDashboardParams struct {
	Name        string   `json:"name"`
	Description *string  `json:"description,omitempty"`
	Pinned      *bool    `json:"pinned,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

type UpdateDashboardParams struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Pinned      *bool    `json:"pinned,omitempty"`
	Deleted     *bool    `json:"deleted,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

func (c *DefaultClient) CreateDashboard(params CreateDashboardParams) (Dashboard, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return Dashboard{}, fmt.Errorf("failed to marshal params: %w", err)
	}

	url := fmt.Sprintf("%s/api/projects/%s/dashboards/", c.host, c.projectId)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return Dashboard{}, fmt.Errorf("failed to build request: %w", err)
	}

	body, err := c.doRequestAndReadBody(req)
	if err != nil {
		return Dashboard{}, fmt.Errorf("failed to handle the HTTP request: %w", err)
	}

	var response Dashboard
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Dashboard{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response, nil
}

func (c *DefaultClient) GetDashboard(id string) (Dashboard, error) {
	url := fmt.Sprintf("%s/api/projects/%s/dashboards/%s/", c.host, c.projectId, id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Dashboard{}, fmt.Errorf("failed to build request: %w", err)
	}

	body, err := c.doRequestAndReadBody(req)
	if err != nil {
		return Dashboard{}, fmt.Errorf("failed to handle the HTTP request: %w", err)
	}

	var response Dashboard
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Dashboard{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response, nil
}

func (c *DefaultClient) UpdateDashboard(id string, params UpdateDashboardParams) (Dashboard, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return Dashboard{}, fmt.Errorf("failed to marshal params: %w", err)
	}

	url := fmt.Sprintf("%s/api/projects/%s/dashboards/%s/", c.host, c.projectId, id)
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		return Dashboard{}, fmt.Errorf("failed to build request: %w", err)
	}

	body, err := c.doRequestAndReadBody(req)
	if err != nil {
		return Dashboard{}, fmt.Errorf("failed to handle the HTTP request: %w", err)
	}

	var response Dashboard
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Dashboard{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response, nil
}

func (c *DefaultClient) DeleteDashboard(id string) error {
	params := UpdateDashboardParams{
		Deleted: ptrBool(true),
	}

	data, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("failed to marshal params: %w", err)
	}

	url := fmt.Sprintf("%s/api/projects/%s/dashboards/%s/", c.host, c.projectId, id)
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}

	_, err = c.doRequestAndReadBody(req)
	if err != nil {
		return fmt.Errorf("failed to handle the HTTP request: %w", err)
	}

	return nil
}

func ptrString(s string) *string {
	return &s
}

func ptrBool(b bool) *bool {
	return &b
}

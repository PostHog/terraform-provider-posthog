package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/posthog/terraform-provider/internal/http/middleware"
)

type HTTPStatusCode int

type PosthogClient struct {
	host       string
	apiKey     string
	projectID  string
	httpClient *http.Client
}

func NewDefaultClient(host, apiKey, projectID string, opts ...ClientOption) PosthogClient {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}

	httpClient.Transport = middleware.NewRetryTransport(httpClient.Transport, middleware.DefaultRetryConfig())

	for _, opt := range opts {
		opt(httpClient)
	}

	return PosthogClient{
		host:       host,
		apiKey:     apiKey,
		projectID:  projectID,
		httpClient: httpClient,
	}
}

func NewClient(client *http.Client, host, apiKey, projectID string, opts ...ClientOption) PosthogClient {
	for _, opt := range opts {
		opt(client)
	}

	return PosthogClient{
		host:       host,
		apiKey:     apiKey,
		projectID:  projectID,
		httpClient: client,
	}
}

func (c *PosthogClient) doRequest(ctx context.Context, method, path string, body any) ([]byte, HTTPStatusCode, error) {
	url := fmt.Sprintf("%s%s", c.host, path)

	var reqBody io.Reader
	var bodyBytes []byte
	if body != nil {
		var err error
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to marshal request: %w", err)
		}
		reqBody = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create request: %w", err)
	}

	// Set GetBody for retry support
	if bodyBytes != nil {
		req.GetBody = func() (io.ReadCloser, error) {
			return io.NopCloser(bytes.NewReader(bodyBytes)), nil
		}
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Posthog/terraform-provider-http v0.0.1")

	tflog.Debug(ctx, "sending http request", map[string]any{"method": method, "url": url})

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, HTTPStatusCode(resp.StatusCode), fmt.Errorf("failed to read response: %w", err)
	}

	tflog.Debug(ctx, "received response", map[string]any{"status": resp.StatusCode})

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return respBody, HTTPStatusCode(resp.StatusCode), fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	return respBody, HTTPStatusCode(resp.StatusCode), nil
}

func doPost[T any](c *PosthogClient, ctx context.Context, path string, body any) (T, HTTPStatusCode, error) {
	var result T
	respBody, status, err := c.doRequest(ctx, http.MethodPost, path, body)
	if err != nil {
		return result, status, fmt.Errorf("failed to send POST request: %w", err)
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return result, status, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return result, status, nil
}

func doGet[T any](c *PosthogClient, ctx context.Context, path string) (T, HTTPStatusCode, error) {
	var result T
	respBody, status, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return result, status, fmt.Errorf("failed to send GET request: %w", err)
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return result, status, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return result, status, nil
}

func doPatch[T any](c *PosthogClient, ctx context.Context, path string, body any) (T, HTTPStatusCode, error) {
	var result T
	respBody, status, err := c.doRequest(ctx, http.MethodPatch, path, body)
	if err != nil {
		return result, status, fmt.Errorf("failed to send PATCH request: %w", err)
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return result, status, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return result, status, nil
}

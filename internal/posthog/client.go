package posthog

import (
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

// Package http provides a high-level client for interacting with Glean's REST API.
// It handles authentication, request formatting, and response parsing while providing
// both standard and streaming request capabilities.
package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/scalvert/glean-cli/internal/config"
)

// Request represents a Glean API request with authentication and headers.
type Request struct {
	Body    interface{}
	Headers map[string]string
	Method  string
	Path    string
	Stream  bool
}

// HTTPClient matches the standard library's http.Client interface for testing.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client provides high-level access to Glean's HTTP API.
type Client interface {
	// SendRequest executes a single request and returns its response body.
	SendRequest(req *Request) ([]byte, error)
	// SendStreamingRequest executes a request that returns a stream of data.
	SendStreamingRequest(req *Request) (io.ReadCloser, error)
	// GetFullURL returns the complete API URL for a given path.
	GetFullURL(path string) string
}

// client wraps http.Client with Glean-specific functionality
type client struct {
	http    HTTPClient
	cfg     *config.Config
	baseURL string
}

// For dependency injection in tests
var NewClientFunc = defaultNewClient

// NewClient creates a new authenticated Glean API client.
func NewClient(cfg *config.Config) (Client, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}
	return NewClientFunc(cfg)
}

// defaultNewClient is the default implementation of NewClient
func defaultNewClient(cfg *config.Config) (Client, error) {
	if cfg.GleanHost == "" {
		return nil, fmt.Errorf("Glean host not configured. Run 'glean config --host <host>' to set it")
	}

	if cfg.GleanToken == "" {
		return nil, fmt.Errorf("Glean token not configured. Run 'glean config --token <token>' to set it")
	}

	baseURL := fmt.Sprintf("https://%s", cfg.GleanHost)

	return &client{
		http:    &http.Client{},
		cfg:     cfg,
		baseURL: baseURL,
	}, nil
}

// GetFullURL ensures the path includes the required /rest/api/v1/ prefix.
func (c *client) GetFullURL(path string) string {
	if !strings.HasPrefix(path, "/rest/api/v1/") {
		path = fmt.Sprintf("/rest/api/v1/%s", strings.TrimPrefix(path, "/"))
	}
	return fmt.Sprintf("%s%s", strings.TrimRight(c.baseURL, "/"), path)
}

// SendRequest executes a single request and returns its response body.
func (c *client) SendRequest(req *Request) ([]byte, error) {
	url := c.GetFullURL(req.Path)

	var bodyReader io.Reader
	if req.Body != nil {
		bodyBytes, err := json.Marshal(req.Body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	httpReq, err := http.NewRequest(req.Method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if c.cfg.GleanToken != "" {
		httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.cfg.GleanToken))
	}
	if c.cfg.GleanEmail != "" {
		httpReq.Header.Set("X-Scio-Actas", c.cfg.GleanEmail)
	}
	httpReq.Header.Set("X-Glean-Auth-Type", "string")

	// Add custom headers
	for k, v := range req.Headers {
		httpReq.Header.Set(k, v)
	}

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		var errorResp struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}
		if err := json.Unmarshal(body, &errorResp); err == nil {
			if errorResp.Message != "" {
				return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, errorResp.Message)
			}
			if errorResp.Error != "" {
				return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, errorResp.Error)
			}
		}
		return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// SendStreamingRequest executes the HTTP request and returns a streaming response
func (c *client) SendStreamingRequest(req *Request) (io.ReadCloser, error) {
	url := c.GetFullURL(req.Path)

	var bodyReader io.Reader
	if req.Body != nil {
		bodyBytes, err := json.Marshal(req.Body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	httpReq, err := http.NewRequest(req.Method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if c.cfg.GleanToken != "" {
		httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.cfg.GleanToken))
	}
	if c.cfg.GleanEmail != "" {
		httpReq.Header.Set("X-Scio-Actas", c.cfg.GleanEmail)
	}
	httpReq.Header.Set("X-Glean-Auth-Type", "string")

	// Add custom headers
	for k, v := range req.Headers {
		httpReq.Header.Set(k, v)
	}

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading error response: %w", err)
		}

		var errorResp struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}
		if err := json.Unmarshal(body, &errorResp); err == nil {
			if errorResp.Message != "" {
				return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, errorResp.Message)
			}
			if errorResp.Error != "" {
				return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, errorResp.Error)
			}
		}
		return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, string(body))
	}

	return resp.Body, nil
}

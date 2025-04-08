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
	"net/http/httputil"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/scalvert/glean-cli/internal/config"
)

// DebugLevel controls the verbosity of client logging
// 0: disabled
// 1: basic request/response info
// 2: detailed including headers
// 3: full request/response including bodies
func getDebugLevel() int {
	if level, exists := os.LookupEnv("GLEAN_HTTP_DEBUG"); exists {
		if lvl, err := strconv.Atoi(level); err == nil {
			return lvl
		}
	}

	return 0
}

// logDebug logs a message if the debug level is >= the specified level
func logDebug(level int, format string, args ...any) {
	if getDebugLevel() >= level {
		debugColor := color.New(color.FgCyan).SprintFunc()
		boldStyle := color.New(color.Bold).SprintFunc()

		fmt.Print("\n")

		fmt.Printf("%s %s\n", debugColor("[DEBUG]"), boldStyle(format))

		for _, arg := range args {
			argStr := fmt.Sprintf("%v", arg)

			lines := strings.Split(argStr, "\n")
			for _, line := range lines {
				fmt.Printf("\t%s\n", line)
			}
		}

		fmt.Print("\n")
	}
}

// dumpRequest returns the http request as a string for debugging
func dumpRequest(req *http.Request) string {
	debugLevel := getDebugLevel()
	if debugLevel < 2 {
		return fmt.Sprintf("%s %s", req.Method, req.URL.String())
	}

	dump, err := httputil.DumpRequestOut(req, debugLevel >= 3)
	if err != nil {
		return fmt.Sprintf("Error dumping request: %s", err)
	}
	return string(dump)
}

// dumpResponse returns the http response as a string for debugging
func dumpResponse(resp *http.Response, includeBody bool) string {
	debugLevel := getDebugLevel()
	if debugLevel < 2 {
		return fmt.Sprintf("Status: %s", resp.Status)
	}

	dump, err := httputil.DumpResponse(resp, includeBody && debugLevel >= 3)
	if err != nil {
		return fmt.Sprintf("Error dumping response: %s", err)
	}
	return string(dump)
}

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

func buildBaseURL(cfg *config.Config) string {
	host := cfg.GleanHost
	port := cfg.GleanPort

	if port != "" {
		return fmt.Sprintf("https://%s:%s", host, port)
	}

	return fmt.Sprintf("https://%s", host)
}

// defaultNewClient is the default implementation of NewClient
func defaultNewClient(cfg *config.Config) (Client, error) {
	if cfg.GleanHost == "" {
		return nil, fmt.Errorf("Glean host not configured. Run 'glean config --host <host>' to set it")
	}

	if cfg.GleanToken == "" {
		return nil, fmt.Errorf("Glean token not configured. Run 'glean config --token <token>' to set it")
	}

	baseURL := buildBaseURL(cfg)

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
	logDebug(1, "Sending request to %s %s", req.Method, url)

	var bodyReader io.Reader
	if req.Body != nil {
		bodyBytes, err := json.Marshal(req.Body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
		logDebug(3, "Request body:", string(bodyBytes))
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

	logDebug(2, "Sending HTTP request:", dumpRequest(httpReq))

	resp, err := c.http.Do(httpReq)
	if err != nil {
		logDebug(1, "Request error:", err)
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	logDebug(2, "Received HTTP response:", dumpResponse(resp, true))

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %s", err)
	}

	logDebug(3, "Response body:", string(body))

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
	logDebug(1, "Sending streaming request to %s %s", req.Method, url)

	var bodyReader io.Reader
	if req.Body != nil {
		bodyBytes, err := json.Marshal(req.Body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
		logDebug(3, "Request body:", string(bodyBytes))
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

	logDebug(2, "Sending HTTP streaming request:", dumpRequest(httpReq))

	resp, err := c.http.Do(httpReq)
	if err != nil {
		logDebug(1, "Streaming request error:", err)
		return nil, fmt.Errorf("error making request: %w", err)
	}

	logDebug(2, "Received HTTP streaming response:", dumpResponse(resp, false))

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading error response: %w", err)
		}

		logDebug(3, "Error response body:", string(body))

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

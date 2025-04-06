package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/scalvert/glean-cli/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockHTTPClient implements HTTPClient for testing
type mockHTTPClient struct {
	doFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if m.doFunc != nil {
		return m.doFunc(req)
	}
	return nil, fmt.Errorf("doFunc not implemented")
}

func TestBuildBaseURL(t *testing.T) {
	tests := []struct {
		name     string
		config   *config.Config
		expected string
	}{
		{
			name: "host only",
			config: &config.Config{
				GleanHost: "test-be.glean.com",
			},
			expected: "https://test-be.glean.com",
		},
		{
			name: "host with port",
			config: &config.Config{
				GleanHost: "foo.bar.com",
				GleanPort: "8080",
			},
			expected: "https://foo.bar.com:8080",
		},
		{
			name: "host with empty port",
			config: &config.Config{
				GleanHost: "test-be.glean.com",
				GleanPort: "",
			},
			expected: "https://test-be.glean.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildBaseURL(tt.config)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNewClient(t *testing.T) {
	// Save original NewClientFunc and restore after test
	originalNewClientFunc := NewClientFunc
	defer func() { NewClientFunc = originalNewClientFunc }()

	t.Run("uses custom NewClientFunc when provided", func(t *testing.T) {
		called := false
		NewClientFunc = func(cfg *config.Config) (Client, error) {
			called = true
			return &client{}, nil
		}

		cfg := &config.Config{
			GleanHost:  "test-be.glean.com",
			GleanToken: "test-token",
		}

		_, err := NewClient(cfg)
		require.NoError(t, err)
		assert.True(t, called, "custom NewClientFunc should have been called")
	})

	t.Run("creates client with valid config", func(t *testing.T) {
		NewClientFunc = defaultNewClient
		cfg := &config.Config{
			GleanHost:  "test-be.glean.com",
			GleanToken: "test-token",
			GleanEmail: "test@example.com",
		}

		client, err := NewClient(cfg)
		require.NoError(t, err)
		assert.NotNil(t, client)
	})

	t.Run("fails with nil config", func(t *testing.T) {
		NewClientFunc = defaultNewClient
		_, err := NewClient(nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "config cannot be nil")
	})

	t.Run("fails with missing host", func(t *testing.T) {
		NewClientFunc = defaultNewClient
		cfg := &config.Config{
			GleanToken: "test-token",
			GleanEmail: "test@example.com",
		}

		_, err := NewClient(cfg)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Glean host not configured")
	})

	t.Run("fails with missing token", func(t *testing.T) {
		NewClientFunc = defaultNewClient
		cfg := &config.Config{
			GleanHost:  "test-be.glean.com",
			GleanEmail: "test@example.com",
		}

		_, err := NewClient(cfg)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Glean token not configured")
	})
}

func TestSendRequest(t *testing.T) {
	cfg := &config.Config{
		GleanToken: "test-token",
		GleanEmail: "test@example.com",
	}

	t.Run("successful GET request with foo.bar.com:7960 ", func(t *testing.T) {
		mock := &mockHTTPClient{
			doFunc: func(req *http.Request) (*http.Response, error) {
				// Verify request
				assert.Equal(t, "GET", req.Method)
				assert.Equal(t, "Bearer test-token", req.Header.Get("Authorization"))
				assert.Equal(t, "test@example.com", req.Header.Get("X-Scio-Actas"))
				assert.Equal(t, "string", req.Header.Get("X-Glean-Auth-Type"))
				assert.Equal(t, "application/json", req.Header.Get("Content-Type"))

				// Verify URL construction
				assert.Equal(t, "https://foo.bar.com:7960/rest/api/v1/test", req.URL.String())

				// Return mock response
				responseBody := map[string]string{"status": "ok"}
				jsonBody, _ := json.Marshal(responseBody)
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewReader(jsonBody)),
				}, nil
			},
		}

		cfgWithPort := &config.Config{
			GleanHost:  "foo.bar.com",
			GleanPort:  "7960",
			GleanToken: "test-token",
			GleanEmail: "test@example.com",
		}

		client := &client{
			http:    mock,
			baseURL: buildBaseURL(cfgWithPort),
			cfg:     cfgWithPort,
		}

		req := &Request{
			Method: "GET",
			Path:   "test",
		}

		resp, err := client.SendRequest(req)
		require.NoError(t, err)

		var result map[string]string
		err = json.Unmarshal(resp, &result)
		require.NoError(t, err)
		assert.Equal(t, "ok", result["status"])
	})

	t.Run("successful GET request with custom headers", func(t *testing.T) {
		mock := &mockHTTPClient{
			doFunc: func(req *http.Request) (*http.Response, error) {
				// Verify request
				assert.Equal(t, "GET", req.Method)
				assert.Equal(t, "Bearer test-token", req.Header.Get("Authorization"))
				assert.Equal(t, "test@example.com", req.Header.Get("X-Scio-Actas"))
				assert.Equal(t, "string", req.Header.Get("X-Glean-Auth-Type"))
				assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
				assert.Equal(t, "custom-value", req.Header.Get("X-Custom-Header"))

				// Verify URL construction
				assert.Equal(t, "https://test-be.glean.com/rest/api/v1/test", req.URL.String())

				// Return mock response
				responseBody := map[string]string{"status": "ok"}
				jsonBody, _ := json.Marshal(responseBody)
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewReader(jsonBody)),
				}, nil
			},
		}

		client := &client{
			http:    mock,
			baseURL: "https://test-be.glean.com",
			cfg:     cfg,
		}

		req := &Request{
			Method: "GET",
			Path:   "test",
			Headers: map[string]string{
				"X-Custom-Header": "custom-value",
			},
		}

		resp, err := client.SendRequest(req)
		require.NoError(t, err)

		var result map[string]string
		err = json.Unmarshal(resp, &result)
		require.NoError(t, err)
		assert.Equal(t, "ok", result["status"])
	})

	t.Run("error marshaling request body", func(t *testing.T) {
		client := &client{
			http:    &mockHTTPClient{},
			baseURL: "https://test-be.glean.com",
			cfg:     cfg,
		}

		req := &Request{
			Method: "POST",
			Path:   "test",
			Body:   make(chan int), // channels cannot be marshaled to JSON
		}

		_, err := client.SendRequest(req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error marshaling request body")
	})

	t.Run("error response with message field", func(t *testing.T) {
		mock := &mockHTTPClient{
			doFunc: func(req *http.Request) (*http.Response, error) {
				responseBody := map[string]string{"message": "invalid request"}
				jsonBody, _ := json.Marshal(responseBody)
				return &http.Response{
					StatusCode: http.StatusBadRequest,
					Body:       io.NopCloser(bytes.NewReader(jsonBody)),
				}, nil
			},
		}

		client := &client{
			http:    mock,
			baseURL: "https://test-be.glean.com",
			cfg:     cfg,
		}

		req := &Request{
			Method: "GET",
			Path:   "error",
		}

		_, err := client.SendRequest(req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid request")
	})

	t.Run("error response with error field", func(t *testing.T) {
		mock := &mockHTTPClient{
			doFunc: func(req *http.Request) (*http.Response, error) {
				responseBody := map[string]string{"error": "bad request"}
				jsonBody, _ := json.Marshal(responseBody)
				return &http.Response{
					StatusCode: http.StatusBadRequest,
					Body:       io.NopCloser(bytes.NewReader(jsonBody)),
				}, nil
			},
		}

		client := &client{
			http:    mock,
			baseURL: "https://test-be.glean.com",
			cfg:     cfg,
		}

		req := &Request{
			Method: "GET",
			Path:   "error",
		}

		_, err := client.SendRequest(req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "bad request")
	})

	t.Run("error response with non-JSON body", func(t *testing.T) {
		mock := &mockHTTPClient{
			doFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusInternalServerError,
					Body:       io.NopCloser(bytes.NewReader([]byte("Internal Server Error"))),
				}, nil
			},
		}

		client := &client{
			http:    mock,
			baseURL: "https://test-be.glean.com",
			cfg:     cfg,
		}

		req := &Request{
			Method: "GET",
			Path:   "error",
		}

		_, err := client.SendRequest(req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Internal Server Error")
	})

	t.Run("successful GET request", func(t *testing.T) {
		mock := &mockHTTPClient{
			doFunc: func(req *http.Request) (*http.Response, error) {
				// Verify request
				assert.Equal(t, "GET", req.Method)
				assert.Equal(t, "Bearer test-token", req.Header.Get("Authorization"))
				assert.Equal(t, "test@example.com", req.Header.Get("X-Scio-Actas"))
				assert.Equal(t, "string", req.Header.Get("X-Glean-Auth-Type"))
				assert.Equal(t, "application/json", req.Header.Get("Content-Type"))

				// Return mock response
				responseBody := map[string]string{"status": "ok"}
				jsonBody, _ := json.Marshal(responseBody)
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewReader(jsonBody)),
				}, nil
			},
		}

		client := &client{
			http:    mock,
			baseURL: "https://test-be.glean.com",
			cfg:     cfg,
		}

		req := &Request{
			Method: "GET",
			Path:   "test",
		}

		resp, err := client.SendRequest(req)
		require.NoError(t, err)

		var result map[string]string
		err = json.Unmarshal(resp, &result)
		require.NoError(t, err)
		assert.Equal(t, "ok", result["status"])
	})

	t.Run("successful POST request", func(t *testing.T) {
		mock := &mockHTTPClient{
			doFunc: func(req *http.Request) (*http.Response, error) {
				// Verify request
				assert.Equal(t, "POST", req.Method)
				assert.Equal(t, "Bearer test-token", req.Header.Get("Authorization"))

				// Verify request body
				var body map[string]interface{}
				json.NewDecoder(req.Body).Decode(&body)
				assert.Equal(t, "test", body["key"])

				// Return mock response
				responseBody := map[string]string{"status": "created"}
				jsonBody, _ := json.Marshal(responseBody)
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewReader(jsonBody)),
				}, nil
			},
		}

		client := &client{
			http:    mock,
			baseURL: "https://test-be.glean.com",
			cfg:     cfg,
		}

		req := &Request{
			Method: "POST",
			Path:   "post",
			Body:   map[string]interface{}{"key": "test"},
		}

		resp, err := client.SendRequest(req)
		require.NoError(t, err)

		var result map[string]string
		err = json.Unmarshal(resp, &result)
		require.NoError(t, err)
		assert.Equal(t, "created", result["status"])
	})

	t.Run("network error", func(t *testing.T) {
		mock := &mockHTTPClient{
			doFunc: func(req *http.Request) (*http.Response, error) {
				return nil, fmt.Errorf("network error")
			},
		}

		client := &client{
			http:    mock,
			baseURL: "https://test-be.glean.com",
			cfg:     cfg,
		}

		req := &Request{
			Method: "GET",
			Path:   "test",
		}

		_, err := client.SendRequest(req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "network error")
	})
}

func TestGetFullURL(t *testing.T) {
	client := &client{
		baseURL: "https://test-be.glean.com",
	}

	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "path with leading slash",
			path:     "/test",
			expected: "https://test-be.glean.com/rest/api/v1/test",
		},
		{
			name:     "path without leading slash",
			path:     "test",
			expected: "https://test-be.glean.com/rest/api/v1/test",
		},
		{
			name:     "empty path",
			path:     "",
			expected: "https://test-be.glean.com/rest/api/v1/",
		},
		{
			name:     "path with full API prefix",
			path:     "/rest/api/v1/test",
			expected: "https://test-be.glean.com/rest/api/v1/test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := client.GetFullURL(tt.path)
			assert.Equal(t, tt.expected, result)
		})
	}
}

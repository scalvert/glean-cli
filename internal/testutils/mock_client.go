// Package testutils provides testing utilities for the Glean CLI,
// including mock HTTP transports for injecting test responses into the SDK.
package testutils

import (
	"bytes"
	"io"
	"net/http"

	glean "github.com/gleanwork/api-client-go"
	"github.com/scalvert/glean-cli/internal/client"
	"github.com/scalvert/glean-cli/internal/config"
)

// MockTransport implements http.RoundTripper (the Do method expected by glean.HTTPClient).
// It returns a predefined response body for every request, making it easy to test
// command output without making real network calls.
type MockTransport struct {
	// Body is returned for every request
	Body []byte
	// Err is returned instead of a response when non-nil
	Err error
	// StatusCode defaults to 200 when zero
	StatusCode int
	// ContentType defaults to "application/json" when empty
	ContentType string
	// Requests records all requests received for inspection
	Requests []*http.Request
}

func (m *MockTransport) Do(req *http.Request) (*http.Response, error) {
	m.Requests = append(m.Requests, req)
	if m.Err != nil {
		return nil, m.Err
	}
	statusCode := m.StatusCode
	if statusCode == 0 {
		statusCode = 200
	}
	// Mirror the Accept header the SDK sends so the SDK can parse its own response.
	// CreateStream sets Accept: text/plain; Create sets Accept: application/json.
	contentType := m.ContentType
	if contentType == "" {
		if accept := req.Header.Get("Accept"); accept != "" {
			contentType = accept
		} else {
			contentType = "application/json"
		}
	}
	return &http.Response{
		StatusCode: statusCode,
		Header:     http.Header{"Content-Type": []string{contentType}},
		Body:       io.NopCloser(bytes.NewReader(m.Body)),
	}, nil
}

// SetupMockClient injects a MockTransport into the SDK client factory and
// returns the mock plus a cleanup function that restores the original factory.
func SetupMockClient(body []byte, err error) (*MockTransport, func()) {
	mock := &MockTransport{Body: body, Err: err}
	origFunc := client.NewFunc
	client.NewFunc = func(cfg *config.Config) (*glean.Glean, error) {
		return glean.New(
			glean.WithInstance("test-company"),
			glean.WithSecurity("test-token"),
			glean.WithClient(mock),
		), nil
	}
	return mock, func() {
		client.NewFunc = origFunc
	}
}

// Package testutils provides testing utilities for the Glean CLI,
// including mock implementations of interfaces and test setup helpers.
package testutils

import (
	"bytes"
	"io"

	"github.com/scalvert/glean-cli/internal/config"
	"github.com/scalvert/glean-cli/internal/http"
)

// MockClient implements http.Client for testing with predefined responses.
// It supports both single responses and sequences of responses for testing
// multiple requests in order.
type MockClient struct {
	// Err is returned instead of Response if non-nil
	Err error
	// Response is returned for single-response scenarios
	Response []byte
	// Responses is used for multi-response scenarios, returning each response in sequence
	Responses [][]byte
	// CallCount tracks the number of requests made
	CallCount int
}

// SendRequest returns the next response in the sequence or the single Response.
func (m *MockClient) SendRequest(req *http.Request) ([]byte, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	if len(m.Responses) > 0 {
		response := m.Responses[m.CallCount]
		m.CallCount++
		if m.CallCount >= len(m.Responses) {
			m.CallCount = len(m.Responses) - 1
		}
		return response, nil
	}
	return m.Response, nil
}

// SendStreamingRequest simulates a streaming response by ensuring each response
// ends with a newline, making it compatible with line-by-line readers.
func (m *MockClient) SendStreamingRequest(req *http.Request) (io.ReadCloser, error) {
	if m.Err != nil {
		return nil, m.Err
	}

	var response []byte
	if len(m.Responses) > 0 {
		response = m.Responses[m.CallCount]
		m.CallCount++
		if m.CallCount >= len(m.Responses) {
			m.CallCount = len(m.Responses) - 1
		}
	} else {
		response = m.Response
	}

	if !bytes.HasSuffix(response, []byte("\n")) {
		response = append(response, '\n')
	}

	return io.NopCloser(bytes.NewReader(response)), nil
}

// GetFullURL returns a test URL with the given path.
func (m *MockClient) GetFullURL(path string) string {
	return "https://test-company-be.glean.com" + path
}

// SetupMockClient creates a mock client for testing and returns a cleanup function.
// The cleanup function should be deferred to restore the original client factory.
func SetupMockClient(response []byte, err error) (*MockClient, func()) {
	mock := &MockClient{
		Response: response,
		Err:      err,
	}
	origFunc := http.NewClientFunc
	http.NewClientFunc = func(cfg *config.Config) (http.Client, error) {
		return mock, nil
	}
	return mock, func() {
		http.NewClientFunc = origFunc
	}
}

package testutils

import (
	"bytes"
	"io"

	"github.com/scalvert/glean-cli/pkg/config"
	"github.com/scalvert/glean-cli/pkg/http"
)

// MockClient is a test double that can return predefined responses
type MockClient struct {
	Err       error
	Response  []byte
	Responses [][]byte
	CallCount int
}

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

	// Ensure each line ends with a newline for proper streaming simulation
	if !bytes.HasSuffix(response, []byte("\n")) {
		response = append(response, '\n')
	}

	// Convert response to a ReadCloser that returns each line separately
	// This simulates the streaming behavior where each line is a complete JSON object
	return io.NopCloser(bytes.NewReader(response)), nil
}

func (m *MockClient) GetFullURL(path string) string {
	return "https://test-company-be.glean.com" + path
}

// SetupMockClient creates a new mock client and returns a cleanup function
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

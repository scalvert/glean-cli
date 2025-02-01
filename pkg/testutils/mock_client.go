package testutils

import (
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

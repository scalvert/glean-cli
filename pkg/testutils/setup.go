package testutils

import (
	"testing"
)

// SetupTestWithResponse sets up both config and mock client for testing
func SetupTestWithResponse(t *testing.T, response []byte) (*MockClient, func()) {
	cleanupConfig := SetupTestConfig(t)
	mock, cleanupMock := SetupMockClient(response, nil)

	cleanup := func() {
		cleanupMock()
		cleanupConfig()
	}

	return mock, cleanup
}

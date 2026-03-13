package testutils

import (
	"testing"
)

// SetupTestWithResponse sets up both config and mock transport for testing.
func SetupTestWithResponse(t *testing.T, response []byte) (*MockTransport, func()) {
	t.Helper()
	cleanupConfig := SetupTestConfig(t)
	mock, cleanupMock := SetupMockClient(response, nil)

	cleanup := func() {
		cleanupMock()
		cleanupConfig()
	}

	return mock, cleanup
}

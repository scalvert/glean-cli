package auth

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindCallbackPort_Automatic(t *testing.T) {
	port, err := findCallbackPort(0)
	require.NoError(t, err)
	assert.Greater(t, port, 0)
	assert.LessOrEqual(t, port, 65535)
}

func TestFindCallbackPort_Preferred(t *testing.T) {
	port, err := findCallbackPort(0)
	require.NoError(t, err)

	got, err := findCallbackPort(port)
	require.NoError(t, err)
	assert.Equal(t, port, got)
}

func TestFindCallbackPort_InUse(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port
	_, err = findCallbackPort(port)
	require.Error(t, err)
}

func TestFindCallbackPort_Invalid(t *testing.T) {
	for _, port := range []int{-1, 65536} {
		_, err := findCallbackPort(port)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "between 1 and 65535")
	}
}

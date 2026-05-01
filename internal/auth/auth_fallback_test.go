package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gleanwork/glean-cli/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDcrOrStaticClient_NoClientAvailable(t *testing.T) {
	t.Setenv("GLEAN_SERVER_URL", "")
	config.ConfigPath = t.TempDir() + "/config.json"

	_, _, err := dcrOrStaticClient(context.Background(), "test-host", "", "http://127.0.0.1:9999/callback")
	require.Error(t, err)
	assert.True(t, errors.Is(err, errNoOAuthClient), "expected errNoOAuthClient, got: %v", err)
}

func TestDcrOrStaticClient_DCRFails_NoStaticClient(t *testing.T) {
	t.Setenv("GLEAN_SERVER_URL", "")
	config.ConfigPath = t.TempDir() + "/config.json"

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	}))
	defer srv.Close()

	_, _, err := dcrOrStaticClient(context.Background(), "test-host", srv.URL, "http://127.0.0.1:9999/callback")
	require.Error(t, err)
	assert.True(t, errors.Is(err, errNoOAuthClient),
		"DCR rejection (403) with no static client means no OAuth client is available")
	assert.Contains(t, err.Error(), "dynamic client registration failed")
}

func TestDcrOrStaticClient_DCRFails_StaticClientFallback(t *testing.T) {
	dir := t.TempDir()
	config.ConfigPath = dir + "/config.json"
	t.Setenv("GLEAN_SERVER_URL", "test-host")

	cfgData, _ := json.Marshal(map[string]string{
		"server_url":          "test-host",
		"oauth_client_id":     "static-id",
		"oauth_client_secret": "static-secret",
	})
	require.NoError(t, os.WriteFile(config.ConfigPath, cfgData, 0o600))

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	}))
	defer srv.Close()

	clientID, clientSecret, err := dcrOrStaticClient(context.Background(), "test-host", srv.URL, "http://127.0.0.1:9999/callback")
	require.NoError(t, err)
	assert.Equal(t, "static-id", clientID)
	assert.Equal(t, "static-secret", clientSecret)
}

func TestDcrOrStaticClient_DCRSucceeds(t *testing.T) {
	dir := t.TempDir()
	config.ConfigPath = dir + "/config.json"
	setStoragePath(t, dir)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"client_id":     "dcr-id",
			"client_secret": "dcr-secret",
		})
	}))
	defer srv.Close()

	clientID, clientSecret, err := dcrOrStaticClient(context.Background(), "test-host", srv.URL, "http://127.0.0.1:9999/callback")
	require.NoError(t, err)
	assert.Equal(t, "dcr-id", clientID)
	assert.Equal(t, "dcr-secret", clientSecret)
}

func TestErrNoOAuthClient_NotMatchedByOtherErrors(t *testing.T) {
	other := errors.New("finding callback port: address already in use")
	assert.False(t, errors.Is(other, errNoOAuthClient))
}

// setStoragePath points token/client storage at a temp directory.
func setStoragePath(t *testing.T, dir string) {
	t.Helper()
	t.Setenv("GLEAN_AUTH_DIR", dir)
}

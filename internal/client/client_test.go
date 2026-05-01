package client

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/gleanwork/glean-cli/internal/config"
	"github.com/gleanwork/glean-cli/internal/httputil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockRoundTripper struct {
	fn func(*http.Request) (*http.Response, error)
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.fn(req)
}

func TestTransport_OAuthSetsHeader(t *testing.T) {
	httputil.SetVersion("test")

	var captured *http.Request
	base := &mockRoundTripper{fn: func(req *http.Request) (*http.Response, error) {
		captured = req
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(""))}, nil
	}}

	transport := httputil.NewTransport(base, httputil.WithHeader("X-Glean-Auth-Type", authTypeOAuth))
	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)
	_, _ = transport.RoundTrip(req)

	assert.Equal(t, authTypeOAuth, captured.Header.Get("X-Glean-Auth-Type"))
	assert.Equal(t, "glean-cli/test", captured.Header.Get("User-Agent"))
}

func TestTransport_APITokenOmitsHeader(t *testing.T) {
	httputil.SetVersion("test")

	var captured *http.Request
	base := &mockRoundTripper{fn: func(req *http.Request) (*http.Response, error) {
		captured = req
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(""))}, nil
	}}

	transport := httputil.NewTransport(base)
	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)
	_, _ = transport.RoundTrip(req)

	assert.Empty(t, captured.Header.Get("X-Glean-Auth-Type"))
	assert.Equal(t, "glean-cli/test", captured.Header.Get("User-Agent"))
}

func TestResolveToken_APIToken(t *testing.T) {
	cfg := &config.Config{GleanServerURL: "https://test-be.glean.com", GleanToken: "api-token-123"}
	token, authType := ResolveToken(cfg)
	assert.Equal(t, "api-token-123", token)
	assert.Empty(t, authType)
}

func TestResolveToken_NoToken(t *testing.T) {
	// auth.LoadOAuthToken returns "" for an unrecognized host, so token stays empty.
	cfg := &config.Config{GleanServerURL: "https://nonexistent-be.glean.com", GleanToken: ""}
	token, authType := ResolveToken(cfg)
	assert.Empty(t, token)
	assert.Empty(t, authType)
}

func TestValidateToken_NoToken(t *testing.T) {
	cfg := &config.Config{GleanServerURL: "https://test-be.glean.com", GleanToken: ""}
	err := ValidateToken(context.Background(), cfg)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no token available")
}

func TestValidateToken_Unreachable(t *testing.T) {
	cfg := &config.Config{GleanServerURL: "http://localhost:1", GleanToken: "some-token"}
	err := ValidateToken(context.Background(), cfg)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validating token")
}

func TestNew_EmptyServerURL(t *testing.T) {
	cfg := &config.Config{GleanServerURL: "", GleanToken: "some-token"}
	_, err := New(cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "server URL not configured")
}

func TestNew_EmptyToken(t *testing.T) {
	// auth.LoadOAuthToken will return "" for a fake URL, so token stays empty
	cfg := &config.Config{GleanServerURL: "https://test-be.glean.com", GleanToken: ""}
	_, err := New(cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not authenticated")
}

func TestNew_Success(t *testing.T) {
	cfg := &config.Config{GleanServerURL: "https://test-be.glean.com", GleanToken: "valid-token"}
	client, err := New(cfg)
	require.NoError(t, err)
	assert.NotNil(t, client)
}

// TestNew_FullHostnames verifies that custom-shape Glean server URLs (vanity
// domains, non-"-be" suffixes, obfuscated subdomains) all produce a working
// SDK client. Covers the bugs originally reported in #102.
func TestNew_FullHostnames(t *testing.T) {
	urls := []string{
		"https://acmecorp-be.glean.com",
		"https://acmecorp-pl.glean.com",
		"https://sub.domain.acmecorp-be.glean.com",
		"https://search.acmecorp.com",
		"https://a7c3d91b-be.glean.com",
	}
	for _, u := range urls {
		t.Run(u, func(t *testing.T) {
			cfg := &config.Config{GleanServerURL: u, GleanToken: "valid-token"}
			client, err := New(cfg)
			require.NoError(t, err)
			assert.NotNil(t, client)
		})
	}
}

package httputil

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockRoundTripper struct {
	fn func(*http.Request) (*http.Response, error)
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.fn(req)
}

func TestUATransport_SetsUserAgent(t *testing.T) {
	SetVersion("1.2.3")

	var captured *http.Request
	base := &mockRoundTripper{fn: func(req *http.Request) (*http.Response, error) {
		captured = req
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(""))}, nil
	}}

	transport := NewTransport(base)
	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)

	_, err = transport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, "glean-cli/1.2.3", captured.Header.Get("User-Agent"))
}

func TestUATransport_DefaultVersion(t *testing.T) {
	SetVersion("dev")

	var captured *http.Request
	base := &mockRoundTripper{fn: func(req *http.Request) (*http.Response, error) {
		captured = req
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(""))}, nil
	}}

	transport := NewTransport(base)
	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)

	_, err = transport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, "glean-cli/dev", captured.Header.Get("User-Agent"))
}

func TestUATransport_DoesNotMutateOriginalRequest(t *testing.T) {
	SetVersion("1.0.0")

	base := &mockRoundTripper{fn: func(req *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(""))}, nil
	}}

	transport := NewTransport(base)
	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)

	_, err = transport.RoundTrip(req)
	require.NoError(t, err)

	assert.Empty(t, req.Header.Get("User-Agent"), "original request should not be mutated")
}

func TestWithHeader_SetsExtraHeader(t *testing.T) {
	SetVersion("1.0.0")

	var captured *http.Request
	base := &mockRoundTripper{fn: func(req *http.Request) (*http.Response, error) {
		captured = req
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(""))}, nil
	}}

	transport := NewTransport(base, WithHeader("X-Custom", "value"))
	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)

	_, err = transport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, "value", captured.Header.Get("X-Custom"))
	assert.Equal(t, "glean-cli/1.0.0", captured.Header.Get("User-Agent"))
}

func TestWithHeader_EmptyValueIsIgnored(t *testing.T) {
	SetVersion("1.0.0")

	var captured *http.Request
	base := &mockRoundTripper{fn: func(req *http.Request) (*http.Response, error) {
		captured = req
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(""))}, nil
	}}

	transport := NewTransport(base, WithHeader("X-Custom", ""))
	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)

	_, err = transport.RoundTrip(req)
	require.NoError(t, err)

	assert.Empty(t, captured.Header.Get("X-Custom"))
}

func TestNewHTTPClient_SetsTransport(t *testing.T) {
	SetVersion("2.0.0")

	client := NewHTTPClient(0)
	require.NotNil(t, client)

	assert.NotNil(t, client.Transport)
}

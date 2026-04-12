package httputil

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gleanwork/glean-cli/internal/debug"
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

func TestTruncate_UnderLimit(t *testing.T) {
	s := "short string"
	assert.Equal(t, s, truncate(s, 100))
}

func TestTruncate_ExactLimit(t *testing.T) {
	s := "exact"
	assert.Equal(t, s, truncate(s, 5))
}

func TestTruncate_OverLimit(t *testing.T) {
	s := "this is a long string that should be truncated"
	result := truncate(s, 10)
	assert.Equal(t, "this is a ... (truncated)", result)
}

func TestBodyLogging_RequestBodyPreserved(t *testing.T) {
	debug.Enable()
	SetVersion("1.0.0")

	requestBody := `{"query":"test"}`
	var capturedBody string
	base := &mockRoundTripper{fn: func(req *http.Request) (*http.Response, error) {
		data, _ := io.ReadAll(req.Body)
		capturedBody = string(data)
		return &http.Response{
			StatusCode: 200,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader(`{"result":"ok"}`)),
		}, nil
	}}

	transport := NewTransport(base)
	req, err := http.NewRequest("POST", "https://example.com/api", strings.NewReader(requestBody))
	require.NoError(t, err)

	_, err = transport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, requestBody, capturedBody, "request body should be preserved after logging")
}

func TestBodyLogging_ResponseBodyPreserved(t *testing.T) {
	debug.Enable()
	SetVersion("1.0.0")

	responseBody := `{"result":"ok"}`
	base := &mockRoundTripper{fn: func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader(responseBody)),
		}, nil
	}}

	transport := NewTransport(base)
	req, err := http.NewRequest("GET", "https://example.com/api", nil)
	require.NoError(t, err)

	resp, err := transport.RoundTrip(req)
	require.NoError(t, err)

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, responseBody, string(data), "response body should be preserved after logging")
}

type trackingReadCloser struct {
	io.ReadCloser
	read bool
}

func (t *trackingReadCloser) Read(p []byte) (int, error) {
	t.read = true
	return t.ReadCloser.Read(p)
}

func TestBodyLogging_SSEResponseNotBuffered(t *testing.T) {
	debug.Enable()
	SetVersion("1.0.0")

	sseBody := &trackingReadCloser{ReadCloser: io.NopCloser(strings.NewReader("data: hello\n\n"))}
	base := &mockRoundTripper{fn: func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Header:     http.Header{"Content-Type": []string{"text/event-stream"}},
			Body:       sseBody,
		}, nil
	}}

	transport := NewTransport(base)
	req, err := http.NewRequest("GET", "https://example.com/stream", nil)
	require.NoError(t, err)

	_, err = transport.RoundTrip(req)
	require.NoError(t, err)

	assert.False(t, sseBody.read, "SSE response body should not be read for logging")
}

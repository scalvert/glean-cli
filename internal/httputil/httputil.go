package httputil

import (
	"net/http"
	"time"

	"github.com/gleanwork/glean-cli/internal/debug"
)

// cliVersion is set at startup via SetVersion. Defaults to "dev" for local builds.
var cliVersion = "dev"

// SetVersion records the build-time version for use in the User-Agent header.
func SetVersion(v string) { cliVersion = v }

// Version returns the current CLI version string.
func Version() string { return cliVersion }

// TransportOption configures a cliTransport.
type TransportOption func(*cliTransport)

// WithHeader adds a static header to every outgoing request.
// If value is empty the header is not set.
func WithHeader(key, value string) TransportOption {
	return func(t *cliTransport) {
		if value != "" {
			t.extraHeaders[key] = value
		}
	}
}

// cliTransport wraps an http.RoundTripper, injects the CLI User-Agent header,
// and applies any additional static headers on every outgoing request.
type cliTransport struct {
	base         http.RoundTripper
	extraHeaders map[string]string
}

var (
	reqLog = debug.New("http:request")
	resLog = debug.New("http:response")
)

func (t *cliTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	req.Header.Set("User-Agent", "glean-cli/"+cliVersion)
	for k, v := range t.extraHeaders {
		req.Header.Set(k, v)
	}

	reqLog.Log("%s %s", req.Method, req.URL.String())

	start := time.Now()
	resp, err := t.base.RoundTrip(req)
	if err != nil {
		resLog.Log("%s %s error: %v (%s)", req.Method, req.URL.String(), err, time.Since(start).Round(time.Millisecond))
		return nil, err
	}

	resLog.Log("%d %s (%s)", resp.StatusCode, http.StatusText(resp.StatusCode), time.Since(start).Round(time.Millisecond))
	return resp, nil
}

// NewTransport returns an http.RoundTripper that injects the CLI User-Agent
// header (and any extra headers from opts) before delegating to base.
// If base is nil, http.DefaultTransport is used.
func NewTransport(base http.RoundTripper, opts ...TransportOption) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	t := &cliTransport{base: base, extraHeaders: make(map[string]string)}
	for _, o := range opts {
		o(t)
	}
	return t
}

// NewHTTPClient returns an *http.Client with the given timeout whose transport
// injects the CLI User-Agent header on every request.
func NewHTTPClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout:   timeout,
		Transport: NewTransport(http.DefaultTransport),
	}
}

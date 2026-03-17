// Package client provides a thin wrapper around the official Glean Go SDK,
// initializing the SDK client from glean-cli's config (env vars, keyring,
// or ~/.glean/config.json).
package client

import (
	"fmt"
	"net/http"
	"strings"

	glean "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/glean-cli/internal/auth"
	"github.com/gleanwork/glean-cli/internal/config"
)

// cliVersion is set at startup via SetVersion. Defaults to "dev" for local builds.
var cliVersion = "dev"

// SetVersion records the build-time version for use in the User-Agent header.
func SetVersion(v string) { cliVersion = v }

// Version returns the current CLI version string.
func Version() string { return cliVersion }

// userAgentTransport wraps an http.RoundTripper and appends the CLI identifier
// to the User-Agent header on every request.
type userAgentTransport struct {
	base http.RoundTripper
}

func (t *userAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	existing := req.Header.Get("User-Agent")
	if existing != "" {
		req.Header.Set("User-Agent", existing+" glean-cli/"+cliVersion)
	} else {
		req.Header.Set("User-Agent", "glean-cli/"+cliVersion)
	}
	return t.base.RoundTrip(req)
}

// New creates an authenticated Glean SDK client from the loaded configuration.
//
// Authentication priority:
//  1. GLEAN_API_TOKEN environment variable (via config.LoadConfig)
//  2. System keyring / ~/.glean/config.json (via config.LoadConfig)
//  3. OAuth token from local storage (via auth.LoadOAuthToken)
//
// The GleanHost value is accepted in two forms:
//   - Full hostname: "linkedin-be.glean.com" → instance = "linkedin"
//   - Short name:   "linkedin"              → passed as-is to WithInstance
func New(cfg *config.Config) (*glean.Glean, error) {
	if cfg.GleanHost == "" {
		return nil, fmt.Errorf("Glean host not configured. Run 'glean auth login' or set GLEAN_HOST")
	}

	token := cfg.GleanToken
	if token == "" {
		token = auth.LoadOAuthToken(cfg.GleanHost)
	}
	if token == "" {
		return nil, fmt.Errorf("not authenticated — run 'glean auth login' or set GLEAN_API_TOKEN")
	}

	instance := extractInstance(cfg.GleanHost)

	opts := []glean.SDKOption{
		glean.WithInstance(instance),
		glean.WithSecurity(token),
		glean.WithClient(&http.Client{
			Transport: &userAgentTransport{base: http.DefaultTransport},
		}),
	}

	return glean.New(opts...), nil
}

// NewFunc is the factory used by NewFromConfig. Override in tests to inject
// a mock HTTP transport: set NewFunc to return glean.New(glean.WithClient(mock)).
var NewFunc = New

// NewFromConfig loads config then creates the SDK client via NewFunc.
// Convenience wrapper for command handlers.
func NewFromConfig() (*glean.Glean, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	return NewFunc(cfg)
}

// extractInstance derives the Glean instance name from a host value.
// "linkedin-be.glean.com" → "linkedin"
// "linkedin"              → "linkedin"
func extractInstance(host string) string {
	if strings.HasSuffix(host, "-be.glean.com") {
		return strings.TrimSuffix(host, "-be.glean.com")
	}
	if strings.Contains(host, ".") {
		// Custom hostname — use as-is; SDK will accept a full URL via WithServerURL
		// but WithInstance only sets the variable. Return the part before the first dot
		// as a best-effort fallback.
		return strings.SplitN(host, ".", 2)[0]
	}
	return host
}

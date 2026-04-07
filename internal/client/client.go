// Package client provides a thin wrapper around the official Glean Go SDK,
// initializing the SDK client from glean-cli's config (env vars, keyring,
// or ~/.glean/config.json).
package client

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	glean "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/glean-cli/internal/auth"
	"github.com/gleanwork/glean-cli/internal/config"
	"github.com/gleanwork/glean-cli/internal/debug"
	"github.com/gleanwork/glean-cli/internal/httputil"
)

var resolveLog = debug.New("client:resolve")

// authTypeOAuth is the X-Glean-Auth-Type header value required for External IdP OAuth tokens.
const authTypeOAuth = "OAUTH"

// ResolveToken returns the bearer token and the X-Glean-Auth-Type value for the
// request. API tokens (cfg.GleanToken) return an empty authType; OAuth tokens
// sourced from local storage return authTypeOAuth.
func ResolveToken(cfg *config.Config) (token, authType string) {
	if cfg.GleanToken != "" {
		resolveLog.Log("using API token from env/config")
		return cfg.GleanToken, ""
	}
	tok := auth.LoadOAuthToken(cfg.GleanHost)
	if tok != "" {
		resolveLog.Log("using OAuth token for %s", cfg.GleanHost)
		return tok, authTypeOAuth
	}
	resolveLog.Log("no credentials found")
	return "", ""
}

// ValidateToken makes a lightweight POST /rest/api/v1/search request to
// verify that the resolved token is accepted by the Glean backend. Returns nil
// if the token is valid, or an error describing the failure.
func ValidateToken(ctx context.Context, cfg *config.Config) error {
	token, authType := ResolveToken(cfg)
	if token == "" {
		return fmt.Errorf("no token available")
	}

	url := "https://" + cfg.GleanHost + "/rest/api/v1/search"
	body := strings.NewReader(`{"query":"","pageSize":1}`)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	if authType != "" {
		req.Header.Set("X-Glean-Auth-Type", authType)
	}

	resp, err := httputil.NewHTTPClient(10 * time.Second).Do(req)
	if err != nil {
		return fmt.Errorf("validating token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		return fmt.Errorf("token rejected by server (HTTP %d)", resp.StatusCode)
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("unexpected status validating token (HTTP %d)", resp.StatusCode)
	}
	return nil
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
		return nil, fmt.Errorf("glean host not configured. Run 'glean auth login' or set GLEAN_HOST")
	}

	token, authType := ResolveToken(cfg)
	if token == "" {
		return nil, fmt.Errorf("not authenticated — run 'glean auth login' or set GLEAN_API_TOKEN")
	}

	instance := extractInstance(cfg.GleanHost)
	resolveLog.Log("instance=%s authType=%s", instance, authType)

	opts := []glean.SDKOption{
		glean.WithInstance(instance),
		glean.WithSecurity(token),
		glean.WithClient(&http.Client{
			Transport: httputil.NewTransport(http.DefaultTransport,
				httputil.WithHeader("X-Glean-Auth-Type", authType),
			),
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

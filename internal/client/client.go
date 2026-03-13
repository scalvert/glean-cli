// Package client provides a thin wrapper around the official Glean Go SDK,
// initialising the SDK client from glean-cli's config (env vars, keyring,
// or ~/.glean/config.json).
package client

import (
	"fmt"
	"net/http"
	"strings"

	glean "github.com/gleanwork/api-client-go"
	"github.com/scalvert/glean-cli/internal/config"
)

// New creates an authenticated Glean SDK client from the loaded configuration.
//
// Authentication priority (handled by config.LoadConfig):
//  1. GLEAN_API_TOKEN environment variable
//  2. System keyring
//  3. ~/.glean/config.json
//
// The GleanHost value is accepted in two forms:
//   - Full hostname: "linkedin-be.glean.com" → instance = "linkedin"
//   - Short name:   "linkedin"              → passed as-is to WithInstance
func New(cfg *config.Config) (*glean.Glean, error) {
	if cfg.GleanHost == "" {
		return nil, fmt.Errorf("Glean host not configured. Run 'glean config --host <host>' or set GLEAN_HOST")
	}
	if cfg.GleanToken == "" {
		return nil, fmt.Errorf("Glean token not configured. Run 'glean config --token <token>' or set GLEAN_API_TOKEN")
	}

	instance := extractInstance(cfg.GleanHost)

	opts := []glean.SDKOption{
		glean.WithInstance(instance),
		glean.WithSecurity(cfg.GleanToken),
	}

	if cfg.GleanEmail != "" {
		opts = append(opts, glean.WithClient(actasClient(cfg.GleanEmail)))
	}

	return glean.New(opts...), nil
}

// NewFromConfig loads config then creates the SDK client.
// Convenience wrapper for command handlers.
func NewFromConfig() (*glean.Glean, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	return New(cfg)
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

// actasHTTPClient wraps the default HTTP client to inject the X-Scio-Actas header
// required for delegated/impersonation requests.
type actasHTTPClient struct {
	inner http.Client
	email string
}

func (c *actasHTTPClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-Scio-Actas", c.email)
	return c.inner.Do(req)
}

func actasClient(email string) *actasHTTPClient {
	return &actasHTTPClient{email: email}
}

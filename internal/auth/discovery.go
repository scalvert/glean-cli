package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

var discoveryHTTPClient = &http.Client{Timeout: 10 * time.Second}

// ErrOAuthNotSupported indicates the Glean instance does not have OAuth configured.
// This is distinct from transient failures (network errors, rate limits, etc).
type ErrOAuthNotSupported struct {
	URL string
}

func (e *ErrOAuthNotSupported) Error() string {
	return fmt.Sprintf("OAuth is not configured at %s", e.URL)
}

type protectedResourceMetadata struct {
	Resource             string   `json:"resource"`
	AuthorizationServers []string `json:"authorization_servers"`
}

// fetchProtectedResource fetches RFC 9728 protected resource metadata.
// baseURL is the Glean backend root (e.g. "https://myco-be.glean.com").
func fetchProtectedResource(ctx context.Context, baseURL string) (*protectedResourceMetadata, error) {
	u := strings.TrimRight(baseURL, "/") + "/.well-known/oauth-protected-resource"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("building protected-resource request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := discoveryHTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetching protected resource metadata: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return nil, &ErrOAuthNotSupported{URL: u}
	default:
		return nil, fmt.Errorf("protected resource metadata returned HTTP %d", resp.StatusCode)
	}

	var meta protectedResourceMetadata
	if err := json.NewDecoder(resp.Body).Decode(&meta); err != nil {
		return nil, fmt.Errorf("parsing protected resource metadata: %w", err)
	}
	if len(meta.AuthorizationServers) == 0 {
		return nil, fmt.Errorf("protected resource metadata has no authorization_servers")
	}
	return &meta, nil
}

// registerClient performs RFC 7591 Dynamic Client Registration.
func registerClient(ctx context.Context, registrationEndpoint, redirectURI string) (*StoredClient, error) {
	body := map[string]any{
		"client_name":                "glean-cli",
		"redirect_uris":              []string{redirectURI},
		"grant_types":                []string{"authorization_code", "refresh_token"},
		"response_types":             []string{"code"},
		"token_endpoint_auth_method": "none",
	}
	data, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("marshaling DCR request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, registrationEndpoint, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("building DCR request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := discoveryHTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("DCR request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("DCR returned HTTP %d", resp.StatusCode)
	}

	var result struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret,omitempty"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("parsing DCR response: %w", err)
	}
	if result.ClientID == "" {
		return nil, fmt.Errorf("DCR response missing client_id")
	}
	return &StoredClient{ClientID: result.ClientID, ClientSecret: result.ClientSecret}, nil
}

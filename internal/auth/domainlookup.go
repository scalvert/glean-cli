package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gleanwork/glean-cli/internal/httputil"
)

const gleanConfigSearchURL = "https://app.glean.com/config/search"

// LookupBackendURL resolves a work email to a Glean backend base URL
// using Glean's domain discovery API.
func LookupBackendURL(ctx context.Context, email string) (string, error) {
	return lookupBackendURL(ctx, email, gleanConfigSearchURL)
}

func lookupBackendURL(ctx context.Context, email, endpoint string) (string, error) {
	domain := extractDomain(email)
	if domain == "" {
		return "", fmt.Errorf("invalid email address: %q", email)
	}
	hostLog.Log("domain lookup: domain=%s endpoint=%s", domain, endpoint)

	body := map[string]any{
		"email":       email,
		"emailDomain": domain,
		"isGleanApp":  true,
	}
	data, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("building domain lookup request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := httputil.NewHTTPClient(10 * time.Second).Do(req)
	if err != nil {
		return "", fmt.Errorf("domain lookup request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("domain lookup returned HTTP %d", resp.StatusCode)
	}

	var result struct {
		SearchConfig struct {
			QueryURL string `json:"queryURL"`
		} `json:"search_config"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("parsing domain lookup response: %w", err)
	}
	if result.SearchConfig.QueryURL == "" {
		return "", fmt.Errorf("no Glean instance found for domain %q", domain)
	}

	backendURL := strings.TrimRight(result.SearchConfig.QueryURL, "/")
	hostLog.Log("domain lookup resolved: %s", backendURL)
	return backendURL, nil
}

func extractDomain(email string) string {
	parts := strings.SplitN(email, "@", 2)
	if len(parts) != 2 || parts[1] == "" {
		return ""
	}
	return parts[1]
}

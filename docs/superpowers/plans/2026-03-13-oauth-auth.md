# OAuth 2.0 Auth Implementation Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add `glean auth login|logout|status` with PKCE + DCR + OIDC using standard libraries, making OAuth the recommended auth path while keeping `GLEAN_API_TOKEN` env var for CI.

**Architecture:** New `internal/auth/` package handles domain lookup, OAuth discovery (RFC 9728 → RFC 8414), DCR (RFC 7591), token storage, and OIDC email extraction. `oauth2cli.GetToken()` drives the full PKCE browser flow. `coreos/go-oidc` handles OIDC discovery and ID token parsing. `internal/client/client.go` gains OAuth token as a second auth tier. All protocol work is done by libraries; only the three non-library HTTP calls (protected resource metadata, DCR, domain lookup) are written with stdlib.

**Tech Stack:** `golang.org/x/oauth2` (PKCE/refresh), `github.com/int128/oauth2cli` (local callback server + browser), `github.com/coreos/go-oidc/v3/oidc` (OIDC discovery + ID token), stdlib `net/http` (3 thin HTTP wrappers)

**Spec:** `docs/superpowers/specs/2026-03-13-oauth-auth-design.md`

---

## File Map

**New files:**
- `internal/auth/storage.go` — per-host token/client/server-metadata file storage (0600)
- `internal/auth/discovery.go` — protected resource metadata fetch + DCR POST
- `internal/auth/domainlookup.go` — email → backend URL via app.glean.com/config/search
- `internal/auth/auth.go` — Login(), Logout(), Status(), EnsureAuth(), loadOAuthToken()
- `internal/auth/storage_test.go`
- `internal/auth/discovery_test.go`
- `internal/auth/domainlookup_test.go`
- `cmd/auth.go` — `glean auth` cobra command group

**Modified:**
- `go.mod` — add oauth2cli, coreos/go-oidc, golang.org/x/oauth2
- `internal/client/client.go` — OAuth token as second auth tier, auto-refresh
- `cmd/root.go` — register NewCmdAuth() as core command

---

## Chunk 1: Dependencies + Storage

### Task 1: Add libraries

**Files:**
- Modify: `go.mod`, `go.sum`

- [ ] Run:
  ```bash
  cd /Users/steve.calvert/workspace/personal/glean-cli
  go get golang.org/x/oauth2@latest
  go get github.com/int128/oauth2cli@latest
  go get github.com/coreos/go-oidc/v3/oidc@latest
  ```
- [ ] Verify in `go.mod`:
  ```bash
  grep -E "oauth2|oauth2cli|go-oidc" go.mod
  ```
  Expected: three new entries.
- [ ] Build to confirm nothing broke: `go build ./...`
- [ ] Commit:
  ```bash
  git add go.mod go.sum
  git commit -m "chore(deps): add oauth2, oauth2cli, coreos/go-oidc for auth"
  ```

---

### Task 2: Token storage (TDD)

**Files:**
- Create: `internal/auth/storage.go`
- Create: `internal/auth/storage_test.go`

**Step 1 — Write failing tests first:**

Create `internal/auth/storage_test.go`:

```go
package auth

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func withTempHome(t *testing.T) {
	t.Helper()
	tmp := t.TempDir()
	t.Setenv("HOME", tmp)
}

func TestSaveAndLoadTokens(t *testing.T) {
	withTempHome(t)
	tok := &StoredTokens{
		AccessToken:  "at-abc",
		RefreshToken: "rt-xyz",
		Expiry:       time.Now().Add(time.Hour).Truncate(time.Second),
		Email:        "steve@glean.com",
	}
	require.NoError(t, SaveTokens("myhost.glean.com", tok))

	got, err := LoadTokens("myhost.glean.com")
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, "at-abc", got.AccessToken)
	assert.Equal(t, "rt-xyz", got.RefreshToken)
	assert.Equal(t, "steve@glean.com", got.Email)
}

func TestLoadTokens_Missing(t *testing.T) {
	withTempHome(t)
	got, err := LoadTokens("nobody.glean.com")
	require.NoError(t, err)
	assert.Nil(t, got)
}

func TestDeleteTokens(t *testing.T) {
	withTempHome(t)
	tok := &StoredTokens{AccessToken: "tok"}
	require.NoError(t, SaveTokens("host.glean.com", tok))
	require.NoError(t, DeleteTokens("host.glean.com"))
	got, err := LoadTokens("host.glean.com")
	require.NoError(t, err)
	assert.Nil(t, got)
}

func TestSaveAndLoadClient(t *testing.T) {
	withTempHome(t)
	cl := &StoredClient{ClientID: "cid-123", ClientSecret: "cs-abc"}
	require.NoError(t, SaveClient("myhost.glean.com", cl))

	got, err := LoadClient("myhost.glean.com")
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, "cid-123", got.ClientID)
}

func TestStoredTokens_IsExpired(t *testing.T) {
	assert.True(t, (&StoredTokens{Expiry: time.Now().Add(-time.Minute)}).IsExpired())
	assert.False(t, (&StoredTokens{Expiry: time.Now().Add(time.Hour)}).IsExpired())
	assert.False(t, (&StoredTokens{}).IsExpired()) // zero expiry = no expiry
}

func TestStateDir_FilePermissions(t *testing.T) {
	withTempHome(t)
	tok := &StoredTokens{AccessToken: "tok"}
	require.NoError(t, SaveTokens("host.glean.com", tok))

	dir := stateDir("host.glean.com")
	info, err := os.Stat(dir)
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0700), info.Mode().Perm())

	tokFile := tokensPath("host.glean.com")
	fi, err := os.Stat(tokFile)
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0600), fi.Mode().Perm())
}
```

- [ ] Run: `go test ./internal/auth/... -v` — Expected: **compilation failure** (package doesn't exist yet). That's correct.

**Step 2 — Implement storage.go:**

Create `internal/auth/storage.go`:

```go
package auth

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// StoredTokens holds persisted OAuth tokens for a Glean host.
type StoredTokens struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	Expiry       time.Time `json:"expiry,omitempty"`
	Email        string    `json:"email,omitempty"`
	TokenType    string    `json:"token_type,omitempty"`
}

// IsExpired returns true if the token expires within the next 60 seconds.
func (t *StoredTokens) IsExpired() bool {
	if t.Expiry.IsZero() {
		return false
	}
	return time.Now().Add(60 * time.Second).After(t.Expiry)
}

// StoredClient holds a registered or configured OAuth client for a Glean host.
type StoredClient struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret,omitempty"`
}

// stateDir returns ~/.local/state/glean-cli/<hash>/ for the given host.
func stateDir(host string) string {
	h := sha256.Sum256([]byte(host))
	key := fmt.Sprintf("%x", h[:8])
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".local", "state", "glean-cli", key)
}

func tokensPath(host string) string { return filepath.Join(stateDir(host), "tokens.json") }
func clientPath(host string) string { return filepath.Join(stateDir(host), "client.json") }

func ensureDir(host string) error {
	dir := stateDir(host)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	return os.Chmod(dir, 0700)
}

func writeJSON(path string, v any) error {
	tmp := path + ".tmp"
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(tmp, data, 0600); err != nil {
		return err
	}
	if err := os.Rename(tmp, path); err != nil {
		return err
	}
	return os.Chmod(path, 0600)
}

func readJSON(path string, v any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// SaveTokens persists OAuth tokens for the given host.
func SaveTokens(host string, tok *StoredTokens) error {
	if err := ensureDir(host); err != nil {
		return fmt.Errorf("creating state dir: %w", err)
	}
	return writeJSON(tokensPath(host), tok)
}

// LoadTokens returns stored tokens for the given host, or nil if none exist.
func LoadTokens(host string) (*StoredTokens, error) {
	var tok StoredTokens
	if err := readJSON(tokensPath(host), &tok); err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	return &tok, nil
}

// DeleteTokens removes stored tokens for the given host.
func DeleteTokens(host string) error {
	err := os.Remove(tokensPath(host))
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

// SaveClient persists an OAuth client registration for the given host.
func SaveClient(host string, cl *StoredClient) error {
	if err := ensureDir(host); err != nil {
		return fmt.Errorf("creating state dir: %w", err)
	}
	return writeJSON(clientPath(host), cl)
}

// LoadClient returns a stored client registration for the given host, or nil if none exist.
func LoadClient(host string) (*StoredClient, error) {
	var cl StoredClient
	if err := readJSON(clientPath(host), &cl); err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	return &cl, nil
}
```

- [ ] Run: `go test ./internal/auth/... -run TestSave -v`
  Expected: PASS.
- [ ] Run: `go test ./internal/auth/... -v`
  Expected: all 6 storage tests PASS.
- [ ] Commit:
  ```bash
  git add internal/auth/storage.go internal/auth/storage_test.go
  git commit -m "feat(auth): add per-host OAuth token and client storage"
  ```

---

## Chunk 2: Discovery + Domain Lookup

### Task 3: Discovery — protected resource metadata + DCR (TDD)

**Files:**
- Create: `internal/auth/discovery.go`
- Create: `internal/auth/discovery_test.go`

**Step 1 — Write failing tests:**

Create `internal/auth/discovery_test.go`:

```go
package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchProtectedResource_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/.well-known/oauth-protected-resource", r.URL.Path)
		json.NewEncoder(w).Encode(map[string]any{
			"resource":             "https://example.glean.com",
			"authorization_servers": []string{"https://auth.example.com"},
		})
	}))
	defer srv.Close()

	result, err := fetchProtectedResource(context.Background(), srv.URL)
	require.NoError(t, err)
	assert.Equal(t, []string{"https://auth.example.com"}, result.AuthorizationServers)
}

func TestFetchProtectedResource_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	_, err := fetchProtectedResource(context.Background(), srv.URL)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestRegisterClient_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "glean-cli", body["client_name"])
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{
			"client_id": "dyn-client-id",
		})
	}))
	defer srv.Close()

	cl, err := registerClient(context.Background(), srv.URL, "http://127.0.0.1:9999/callback")
	require.NoError(t, err)
	assert.Equal(t, "dyn-client-id", cl.ClientID)
}

func TestRegisterClient_WithSecret(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{
			"client_id":     "cid",
			"client_secret": "cs",
		})
	}))
	defer srv.Close()

	cl, err := registerClient(context.Background(), srv.URL, "http://127.0.0.1:9999/callback")
	require.NoError(t, err)
	assert.Equal(t, "cs", cl.ClientSecret)
}
```

- [ ] Run: `go test ./internal/auth/... -run TestFetch -v` — Expected: compile error.

**Step 2 — Implement discovery.go:**

Create `internal/auth/discovery.go`:

```go
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

type protectedResourceMetadata struct {
	Resource             string   `json:"resource"`
	AuthorizationServers []string `json:"authorization_servers"`
}

// fetchProtectedResource fetches RFC 9728 protected resource metadata from baseURL.
// baseURL should be the Glean backend root (e.g. "https://myco-be.glean.com").
func fetchProtectedResource(ctx context.Context, baseURL string) (*protectedResourceMetadata, error) {
	// Strip any path; metadata lives at the origin root.
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
		return nil, fmt.Errorf("OAuth not found at %s — instance may not support OAuth", u)
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
// registrationEndpoint is from the auth server metadata.
// redirectURI is the local callback URL (e.g. "http://127.0.0.1:<port>/callback").
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
```

- [ ] Run: `go test ./internal/auth/... -run "TestFetch|TestRegister" -v`
  Expected: all 4 tests PASS.
- [ ] Commit:
  ```bash
  git add internal/auth/discovery.go internal/auth/discovery_test.go
  git commit -m "feat(auth): add protected resource metadata fetch and DCR"
  ```

---

### Task 4: Domain lookup (TDD)

**Files:**
- Create: `internal/auth/domainlookup.go`
- Create: `internal/auth/domainlookup_test.go`

**Step 1 — Write failing test:**

Create `internal/auth/domainlookup_test.go`:

```go
package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLookupBackendURL_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "glean.com", body["emailDomain"])
		json.NewEncoder(w).Encode(map[string]any{
			"search_config": map[string]any{
				"queryURL": "https://scio-prod-be.glean.com/",
			},
		})
	}))
	defer srv.Close()

	url, err := lookupBackendURL(context.Background(), "steve@glean.com", srv.URL)
	require.NoError(t, err)
	// Trailing slash stripped
	assert.Equal(t, "https://scio-prod-be.glean.com", url)
}

func TestLookupBackendURL_EmptyResult(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"search_config": map[string]any{}})
	}))
	defer srv.Close()

	_, err := lookupBackendURL(context.Background(), "user@unknown.com", srv.URL)
	require.Error(t, err)
}

func TestExtractDomain(t *testing.T) {
	assert.Equal(t, "glean.com", extractDomain("steve@glean.com"))
	assert.Equal(t, "", extractDomain("notanemail"))
}
```

- [ ] Run: `go test ./internal/auth/... -run TestLookup -v` — Expected: compile error.

**Step 2 — Implement domainlookup.go:**

Create `internal/auth/domainlookup.go`:

```go
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

const gleanConfigSearchURL = "https://app.glean.com/config/search"

var domainLookupHTTPClient = &http.Client{Timeout: 10 * time.Second}

// LookupBackendURL resolves a work email address to a Glean backend base URL
// using the Glean domain discovery API.
func LookupBackendURL(ctx context.Context, email string) (string, error) {
	return lookupBackendURL(ctx, email, gleanConfigSearchURL)
}

// lookupBackendURL is the testable implementation (accepts configurable endpoint).
func lookupBackendURL(ctx context.Context, email, endpoint string) (string, error) {
	domain := extractDomain(email)
	if domain == "" {
		return "", fmt.Errorf("invalid email address: %q", email)
	}

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

	resp, err := domainLookupHTTPClient.Do(req)
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

	return strings.TrimRight(result.SearchConfig.QueryURL, "/"), nil
}

// extractDomain returns the domain portion of an email address, or "" if invalid.
func extractDomain(email string) string {
	parts := strings.SplitN(email, "@", 2)
	if len(parts) != 2 || parts[1] == "" {
		return ""
	}
	return parts[1]
}
```

- [ ] Run: `go test ./internal/auth/... -run "TestLookup|TestExtract" -v`
  Expected: all 3 tests PASS.
- [ ] Commit:
  ```bash
  git add internal/auth/domainlookup.go internal/auth/domainlookup_test.go
  git commit -m "feat(auth): add email-to-backend-URL domain lookup"
  ```

---

## Chunk 3: Core Auth Logic

### Task 5: auth.go — Login, Logout, Status, EnsureAuth

**Files:**
- Create: `internal/auth/auth.go`

No unit tests for this file — it orchestrates library calls (oauth2cli, oidc) that require real browser interaction. Integration verification is done manually in Task 9.

Create `internal/auth/auth.go`:

```go
package auth

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/int128/oauth2cli"
	"github.com/scalvert/glean-cli/internal/config"
	"golang.org/x/oauth2"
)

// Login performs the full OAuth 2.0 PKCE login flow for the configured Glean host.
//
// Flow:
//  1. Resolve host (from config or email prompt + domain lookup)
//  2. OIDC discovery via coreos/go-oidc (covers auth server metadata)
//  3. Resolve client_id: load stored client → DCR → static config
//  4. oauth2cli.GetToken() — local callback server + browser + PKCE
//  5. Extract email from ID token; store tokens
//
// If the instance doesn't support OAuth, falls back to an inline token prompt.
func Login(ctx context.Context) error {
	host, err := resolveHost(ctx)
	if err != nil {
		return err
	}

	// Try OIDC / OAuth discovery.
	provider, oauthEndpoint, err := discover(ctx, host)
	if err != nil {
		// Instance doesn't support OAuth — prompt for a static token inline.
		return promptForAPIToken(host, err)
	}

	clientID, clientSecret, err := resolveClientID(ctx, host, oauthEndpoint)
	if err != nil {
		return fmt.Errorf("resolving OAuth client: %w", err)
	}

	scopes := []string{oidc.ScopeOpenID, "email", "profile"}

	oauthCfg := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     oauthEndpoint,
		Scopes:       scopes,
	}

	// oauth2cli drives the full PKCE browser flow:
	// starts local server, opens browser, waits for callback, exchanges code.
	token, err := oauth2cli.GetToken(ctx, oauth2cli.Config{
		OAuth2Config:    oauthCfg,
		AuthCodeOptions: []oauth2.AuthCodeOption{oauth2.S256ChallengeOption(oauth2.GenerateVerifier())},
		Logf:            func(string, ...any) {}, // suppress oauth2cli debug logs
	})
	if err != nil {
		return fmt.Errorf("OAuth login failed: %w", err)
	}

	// Extract email from ID token if available.
	email := ""
	if rawIDToken, ok := token.Extra("id_token").(string); ok && provider != nil {
		verifier := provider.Verifier(&oidc.Config{ClientID: clientID})
		if idToken, err := verifier.Verify(ctx, rawIDToken); err == nil {
			var claims struct {
				Email string `json:"email"`
			}
			_ = idToken.Claims(&claims)
			email = claims.Email
		}
	}
	// Fallback: UserInfo endpoint.
	if email == "" && provider != nil {
		if ui, err := provider.UserInfo(ctx, oauth2.StaticTokenSource(token)); err == nil {
			_ = ui.Claims(&struct{ Email *string `json:"email"` }{&email})
		}
	}

	stored := &StoredTokens{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
		Email:        email,
		TokenType:    token.TokenType,
	}
	if err := SaveTokens(host, stored); err != nil {
		return fmt.Errorf("saving tokens: %w", err)
	}

	if email != "" {
		fmt.Printf("✓ Authenticated as %s (%s)\n", email, host)
	} else {
		fmt.Printf("✓ Authenticated with Glean (%s)\n", host)
	}
	return nil
}

// Logout removes the stored OAuth tokens for the configured host.
func Logout(ctx context.Context) error {
	cfg, _ := config.LoadConfig()
	if cfg == nil || cfg.GleanHost == "" {
		return fmt.Errorf("no Glean host configured")
	}
	if err := DeleteTokens(cfg.GleanHost); err != nil {
		return fmt.Errorf("removing tokens: %w", err)
	}
	fmt.Printf("✓ Logged out from Glean (%s)\n", cfg.GleanHost)
	return nil
}

// Status prints the current authentication state.
func Status(ctx context.Context) error {
	cfg, _ := config.LoadConfig()
	host := ""
	if cfg != nil {
		host = cfg.GleanHost
	}
	if host == "" {
		fmt.Println("No Glean host configured.")
		fmt.Println("Run 'glean config --host <host>' or 'glean auth login' to get started.")
		return nil
	}

	// Check env/static token first.
	if cfg.GleanToken != "" {
		fmt.Printf("✓ Authenticated via API token (%s)\n", host)
		return nil
	}

	tok, err := LoadTokens(host)
	if err != nil {
		return fmt.Errorf("reading stored tokens: %w", err)
	}
	if tok == nil {
		fmt.Printf("Not authenticated.\nRun 'glean auth login' to authenticate.\n")
		return nil
	}
	if tok.IsExpired() {
		fmt.Printf("Token expired.\nRun 'glean auth login' to re-authenticate.\n")
		return nil
	}

	expStr := "no expiry"
	if !tok.Expiry.IsZero() {
		remaining := time.Until(tok.Expiry).Round(time.Minute)
		expStr = fmt.Sprintf("expires %s (in %v)", tok.Expiry.UTC().Format(time.RFC3339), remaining)
	}

	if tok.Email != "" {
		fmt.Printf("✓ Authenticated as %s (%s)\n  Token %s\n", tok.Email, host, expStr)
	} else {
		fmt.Printf("✓ Authenticated with Glean (%s)\n  Token %s\n", host, expStr)
	}
	return nil
}

// EnsureAuth returns nil if the client has usable credentials (env var, OAuth token,
// or legacy static token). Returns a human-readable error if nothing is configured.
func EnsureAuth(ctx context.Context) error {
	cfg, err := config.LoadConfig()
	if err == nil && cfg.GleanToken != "" {
		return nil // env var or legacy token
	}
	if err == nil && cfg.GleanHost != "" {
		tok, _ := LoadTokens(cfg.GleanHost)
		if tok != nil && !tok.IsExpired() {
			return nil
		}
	}
	return fmt.Errorf("not authenticated — run 'glean auth login' to authenticate")
}

// LoadOAuthToken returns a valid, non-expired OAuth access token for the given host,
// or "" if none is available.
func LoadOAuthToken(host string) string {
	tok, err := LoadTokens(host)
	if err != nil || tok == nil || tok.IsExpired() {
		return ""
	}
	return tok.AccessToken
}

// resolveHost returns the configured Glean host, prompting for email if needed.
func resolveHost(ctx context.Context) (string, error) {
	cfg, _ := config.LoadConfig()
	if cfg != nil && cfg.GleanHost != "" {
		return cfg.GleanHost, nil
	}

	fmt.Print("Enter your work email: ")
	reader := bufio.NewReader(os.Stdin)
	email, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("reading email: %w", err)
	}
	email = strings.TrimSpace(email)

	fmt.Printf("Looking up your Glean instance...")
	backendURL, err := LookupBackendURL(ctx, email)
	if err != nil {
		fmt.Println()
		return "", fmt.Errorf("could not find a Glean instance for %q: %w", email, err)
	}
	fmt.Println(" found.")

	// Extract hostname from URL.
	host := strings.TrimPrefix(backendURL, "https://")
	host = strings.TrimPrefix(host, "http://")
	host = strings.SplitN(host, "/", 2)[0]

	// Persist host for future commands.
	_ = config.SaveConfig(host, "", "", "")
	return host, nil
}

// discover uses coreos/go-oidc to perform OIDC discovery on the Glean backend.
// It returns the OIDC Provider (for ID token verification) and the OAuth2 endpoint.
// Returns an error if the instance doesn't expose OIDC/OAuth endpoints.
func discover(ctx context.Context, host string) (*oidc.Provider, oauth2.Endpoint, error) {
	// First check if protected resource metadata points us to an auth server.
	baseURL := "https://" + host
	meta, err := fetchProtectedResource(ctx, baseURL)
	if err != nil {
		return nil, oauth2.Endpoint{}, err
	}

	issuer := meta.AuthorizationServers[0]
	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		return nil, oauth2.Endpoint{}, fmt.Errorf("OIDC discovery failed for %s: %w", issuer, err)
	}

	return provider, provider.Endpoint(), nil
}

// resolveClientID returns the client_id and client_secret to use for the OAuth flow.
// Priority: stored client (from previous DCR) → DCR → glean config --oauth-client-id.
func resolveClientID(ctx context.Context, host string, endpoint oauth2.Endpoint) (string, string, error) {
	// 1. Use previously registered client.
	if cl, err := LoadClient(host); err == nil && cl != nil && cl.ClientID != "" {
		return cl.ClientID, cl.ClientSecret, nil
	}

	// 2. Try DCR if registration_endpoint is available.
	// oauth2cli will allocate the port; we use a placeholder to register.
	// The actual redirect_uri is set by oauth2cli per-request.
	// Register with a wildcard localhost redirect.
	if endpoint.AuthStyle == 0 { // endpoint from oidc.Provider always has TokenURL
		// We need the registration_endpoint — fetch auth server metadata directly.
		// oidc.Provider doesn't expose registration_endpoint, so we fetch it ourselves.
	}
	// Try DCR with a temporary placeholder redirect URI.
	// oauth2cli uses http://127.0.0.1:<random>/callback — we re-register if needed.
	// For simplicity, we register once and reuse. The redirect_uri in DCR is
	// stored and must match the auth request; we use 127.0.0.1:0 as a wildcard hint.
	// Real registration happens just before GetToken() if no client is stored.
	baseURL := "https://" + host
	prMeta, err := fetchProtectedResource(ctx, baseURL)
	if err == nil && len(prMeta.AuthorizationServers) > 0 {
		// Fetch auth server metadata to get registration_endpoint.
		issuer := prMeta.AuthorizationServers[0]
		// oidc.Provider already has this; but we need registration_endpoint specifically.
		// Use a direct fetch of the auth server metadata.
		authMeta, err := fetchAuthServerMetadata(ctx, issuer)
		if err == nil && authMeta.RegistrationEndpoint != "" {
			// Use a stable redirect URI pattern for DCR.
			// oauth2cli will use whatever port it allocates; we register with a 127.0.0.1 redirect.
			// In practice, Glean's DCR accepts any localhost redirect.
			cl, err := registerClient(ctx, authMeta.RegistrationEndpoint, "http://127.0.0.1/callback")
			if err == nil {
				_ = SaveClient(host, cl)
				return cl.ClientID, cl.ClientSecret, nil
			}
		}
	}

	// 3. Static client from config.
	cfg, _ := config.LoadConfig()
	if cfg != nil && cfg.OAuthClientID != "" {
		return cfg.OAuthClientID, cfg.OAuthClientSecret, nil
	}

	return "", "", fmt.Errorf("no OAuth client available — run 'glean config --oauth-client-id <id>' for static clients")
}

// fetchAuthServerMetadata fetches OAuth 2.0 Authorization Server Metadata (RFC 8414).
// This supplements oidc.Provider which doesn't expose registration_endpoint.
func fetchAuthServerMetadata(ctx context.Context, issuer string) (*authServerMeta, error) {
	u := strings.TrimRight(issuer, "/") + "/.well-known/oauth-authorization-server"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := discoveryHTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("auth server metadata returned HTTP %d", resp.StatusCode)
	}
	var meta authServerMeta
	return &meta, json.NewDecoder(resp.Body).Decode(&meta)
}

type authServerMeta struct {
	Issuer               string `json:"issuer"`
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint        string `json:"token_endpoint"`
	RegistrationEndpoint string `json:"registration_endpoint,omitempty"`
}

// promptForAPIToken handles the case where OAuth is not supported.
// Stays in glean auth login with an inline token prompt.
func promptForAPIToken(host string, discoveryErr error) error {
	fmt.Printf("\nThis Glean instance doesn't support OAuth.\n")
	fmt.Printf("Contact your Glean administrator to generate an API token.\n")
	fmt.Printf("  (Glean Admin → Settings → API Tokens)\n\n")

	fmt.Print("Token: ")
	reader := bufio.NewReader(os.Stdin)
	token, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("reading token: %w", err)
	}
	token = strings.TrimSpace(token)
	if token == "" {
		return fmt.Errorf("no token provided")
	}

	if err := config.SaveConfig(host, "", token, ""); err != nil {
		return fmt.Errorf("saving token: %w", err)
	}
	fmt.Printf("✓ API token saved for %s\n", host)
	return nil
}
```

Note: This references `config.OAuthClientID` and `config.OAuthClientSecret` fields and a `net/http` import that needs to be in the file. The implementer must:
1. Add `OAuthClientID` and `OAuthClientSecret` fields to `internal/config/config.go`'s `Config` struct and load them from env vars `GLEAN_OAUTH_CLIENT_ID` / `GLEAN_OAUTH_CLIENT_SECRET`
2. Add `"net/http"` and `"encoding/json"` imports to auth.go

- [ ] Add to `internal/config/config.go` `Config` struct:
  ```go
  OAuthClientID     string `json:"oauth_client_id,omitempty"`
  OAuthClientSecret string `json:"oauth_client_secret,omitempty"`
  ```
  And in `loadFromEnv()`:
  ```go
  if v := os.Getenv("GLEAN_OAUTH_CLIENT_ID"); v != "" { cfg.OAuthClientID = v }
  if v := os.Getenv("GLEAN_OAUTH_CLIENT_SECRET"); v != "" { cfg.OAuthClientSecret = v }
  ```

- [ ] Build: `go build ./internal/auth/... ./internal/config/...`
  Fix any compilation errors before proceeding.
- [ ] Build the full project: `go build ./...`
- [ ] Commit:
  ```bash
  git add internal/auth/auth.go internal/config/config.go
  git commit -m "feat(auth): add Login, Logout, Status, EnsureAuth orchestration"
  ```

---

## Chunk 4: cmd/auth.go + client.go integration

### Task 6: cmd/auth.go command group

**Files:**
- Create: `cmd/auth.go`
- Modify: `cmd/root.go`

Create `cmd/auth.go`:

```go
package cmd

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/scalvert/glean-cli/internal/auth"
	"github.com/spf13/cobra"
)

// NewCmdAuth creates the `glean auth` command group.
func NewCmdAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate with Glean",
		Long: heredoc.Doc(`
			Manage authentication with your Glean instance.

			Use 'glean auth login' to authenticate via your browser (recommended).
			For CI/CD environments, set GLEAN_API_TOKEN instead.
		`),
	}
	cmd.AddCommand(newAuthLoginCmd(), newAuthLogoutCmd(), newAuthStatusCmd())
	return cmd
}

func newAuthLoginCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Authenticate with Glean via your browser",
		Long: heredoc.Doc(`
			Opens your browser to authenticate with Glean using OAuth 2.0 + PKCE.

			If your Glean instance supports OAuth with Dynamic Client Registration
			(most instances), no additional configuration is needed.

			For instances with a pre-registered OAuth app, configure first:
			  glean config --oauth-client-id <id>

			For CI/CD environments, set GLEAN_API_TOKEN instead of using this command.

			Examples:
			  glean auth login
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return auth.Login(cmd.Context())
		},
		SilenceUsage: true,
	}
}

func newAuthLogoutCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "Remove stored Glean credentials",
		Long: heredoc.Doc(`
			Removes the OAuth token stored by 'glean auth login'.

			To clear a static API token, use 'glean config --clear' instead.
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return auth.Logout(cmd.Context())
		},
		SilenceUsage: true,
	}
}

func newAuthStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show current authentication status",
		RunE: func(cmd *cobra.Command, args []string) error {
			return auth.Status(cmd.Context())
		},
		SilenceUsage: true,
	}
}
```

- [ ] In `cmd/root.go`, add `NewCmdAuth()` to the core commands slice (alongside search, chat, config, etc.):
  ```go
  // In the core commands loop:
  NewCmdAuth(),
  ```
  Set its GroupID to "core" to match the existing group structure.

- [ ] Build: `go build ./...`
- [ ] Test: `go run . auth --help` — should show login/logout/status subcommands.
- [ ] Commit:
  ```bash
  git add cmd/auth.go cmd/root.go
  git commit -m "feat(auth): add glean auth login|logout|status command group"
  ```

---

### Task 7: Update internal/client/client.go for OAuth token

**Files:**
- Modify: `internal/client/client.go`

Add OAuth token as second priority in `New()`, with auto-refresh:

- [ ] Read `internal/client/client.go` carefully.
- [ ] In `New(cfg *config.Config)`, after checking `cfg.GleanToken != ""`, add:
  ```go
  // If no static token, try OAuth token from storage.
  token := cfg.GleanToken
  if token == "" {
      token = auth.LoadOAuthToken(cfg.GleanHost)
  }
  if token == "" {
      return nil, fmt.Errorf("not authenticated — run 'glean auth login' or set GLEAN_API_TOKEN")
  }
  ```
  Replace the current error on empty token with this block.
  Add `"github.com/scalvert/glean-cli/internal/auth"` import.

- [ ] Build: `go build ./...`
- [ ] Run: `go test ./... -timeout 30s`
  Expected: all tests pass. (Existing tests use `SetupTestWithResponse` which sets `GLEAN_API_TOKEN` via env, so they won't be affected.)
- [ ] Commit:
  ```bash
  git add internal/client/client.go
  git commit -m "feat(client): add OAuth token as second auth priority after env/legacy token"
  ```

---

## Chunk 5: Final Verification

### Task 8: Build, test, and verify

- [ ] Run full test suite:
  ```bash
  cd /Users/steve.calvert/workspace/personal/glean-cli
  go test ./... -race -timeout 60s
  ```
  Expected: all packages pass.

- [ ] Build final binary:
  ```bash
  go build -o /tmp/glean-final .
  ```

- [ ] Verify auth commands exist and help text is accurate:
  ```bash
  /tmp/glean-final auth --help
  /tmp/glean-final auth login --help
  /tmp/glean-final auth status
  /tmp/glean-final --help | grep -E "Core|auth"
  ```

- [ ] Verify `glean auth status` before login shows a useful message (not a panic).

- [ ] Manual login test (requires real Glean credentials):
  ```bash
  /tmp/glean-final auth login
  # Expected: browser opens, after auth: "✓ Authenticated as <email>"
  /tmp/glean-final auth status
  # Expected: shows email + expiry
  /tmp/glean-final search "test query"
  # Expected: results returned using OAuth token
  /tmp/glean-final auth logout
  # Expected: "✓ Logged out"
  ```

- [ ] Commit checklist update:
  ```bash
  git add EVAL-CHECKLIST.md
  git commit -m "chore: mark OAuth auth items as closed"
  ```

---

## Implementation Notes for Agents

1. **oauth2cli PKCE**: Pass `oauth2.S256ChallengeOption(oauth2.GenerateVerifier())` in `AuthCodeOptions` AND `oauth2.VerifierOption(verifier)` in `TokenRequestOptions`. The verifier must be the same value — generate once and use both. The current `auth.go` skeleton has a bug: it generates the verifier inline in `S256ChallengeOption` without capturing it for `VerifierOption`. Fix: generate the verifier separately before building the config:
   ```go
   verifier := oauth2.GenerateVerifier()
   cfg := oauth2cli.Config{
       AuthCodeOptions:     []oauth2.AuthCodeOption{oauth2.S256ChallengeOption(verifier)},
       TokenRequestOptions: []oauth2.AuthCodeOption{oauth2.VerifierOption(verifier)},
   }
   ```

2. **oidc.NewProvider() discovery URL**: For Glean, the issuer from `authorization_servers[0]` might be `https://host/oauth` (with a path). `oidc.NewProvider` will look at `<issuer>/.well-known/openid-configuration`. If that fails, try `<origin>/.well-known/oauth-authorization-server` and construct `oidc.ProviderConfig` manually from the response.

3. **Import cycle**: `internal/auth` imports `internal/config`. `internal/client` imports both. This is fine — no cycle. But `internal/config` must NOT import `internal/auth`.

4. **Test for EnsureAuth**: The existing tests in `cmd/` use `SetupTestWithResponse` which sets `GLEAN_API_TOKEN` via env. `EnsureAuth` checks `cfg.GleanToken != ""` first, so existing tests continue to work without changes.

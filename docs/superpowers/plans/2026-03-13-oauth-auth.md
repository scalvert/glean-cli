# OAuth 2.0 Auth Implementation Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add `glean auth login|logout|status` with PKCE Authorization Code Flow and Device Flow fallback, making OAuth the recommended auth path while keeping `GLEAN_API_TOKEN` env var for CI.

**Architecture:** New `internal/auth/` package handles discovery, PKCE, local callback server, and token storage. `internal/client/client.go` gains a third auth priority tier (OAuth tokens). `cmd/auth.go` exposes the user-facing commands. `golang.org/x/oauth2` does all crypto/protocol work — no hand-rolled PKCE.

**Tech Stack:** `golang.org/x/oauth2` (PKCE + device flow), `github.com/pkg/browser` (cross-platform browser open), `net/http` stdlib (callback server), Go keyring (existing, for token storage)

**Spec:** `docs/superpowers/specs/2026-03-13-auth-and-bug-fixes-design.md`

**Prerequisite:** Run `docs/superpowers/plans/2026-03-13-chk-bug-fixes.md` first (config loading fix in CHK-005/006 must be in place).

---

## Chunk 1: Dependencies + Package Skeleton

### Task 1: Add required libraries

**Files:**
- Modify: `go.mod`, `go.sum` (via `go get`)

- [ ] Add dependencies:
  ```bash
  go get golang.org/x/oauth2@latest
  go get github.com/pkg/browser@latest
  ```
- [ ] Verify they appear in `go.mod`:
  ```bash
  grep "oauth2\|pkg/browser" go.mod
  ```
- [ ] Commit:
  ```bash
  git add go.mod go.sum
  git commit -m "chore(deps): add golang.org/x/oauth2 and github.com/pkg/browser for OAuth PKCE"
  ```

---

### Task 2: Create internal/auth package skeleton

**Files:**
- Create: `internal/auth/auth.go`
- Create: `internal/auth/discovery.go`
- Create: `internal/auth/storage.go`
- Create: `internal/auth/callback.go`

- [ ] Create `internal/auth/discovery.go`:

```go
package auth

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
)

// authServerMetadata holds the subset of RFC 8414 we need.
type authServerMetadata struct {
    AuthorizationEndpoint         string `json:"authorization_endpoint"`
    TokenEndpoint                 string `json:"token_endpoint"`
    DeviceAuthorizationEndpoint   string `json:"device_authorization_endpoint"`
}

// DiscoverMetadata fetches OAuth 2.0 Authorization Server Metadata (RFC 8414)
// from the Glean host. host may be a short name ("linkedin") or full hostname
// ("linkedin-be.glean.com"); both are expanded to the full form.
func DiscoverMetadata(ctx context.Context, host string) (*authServerMetadata, error) {
    baseURL := hostToBaseURL(host)
    url := baseURL + "/.well-known/oauth-authorization-server"

    req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
    if err != nil {
        return nil, fmt.Errorf("building discovery request: %w", err)
    }

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("fetching OAuth metadata from %s: %w", url, err)
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusNotFound {
        return nil, fmt.Errorf("this Glean instance does not support OAuth; use 'glean config --token' instead")
    }
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("OAuth discovery returned HTTP %d from %s", resp.StatusCode, url)
    }

    var meta authServerMetadata
    if err := json.NewDecoder(resp.Body).Decode(&meta); err != nil {
        return nil, fmt.Errorf("parsing OAuth metadata: %w", err)
    }

    if meta.AuthorizationEndpoint == "" || meta.TokenEndpoint == "" {
        return nil, fmt.Errorf("incomplete OAuth metadata from %s (missing authorization_endpoint or token_endpoint)", url)
    }

    return &meta, nil
}

// hostToBaseURL converts a Glean host to its HTTPS base URL.
// Short names (e.g. "linkedin") are expanded to "linkedin-be.glean.com".
// Full hostnames are used as-is.
func hostToBaseURL(host string) string {
    if !strings.Contains(host, ".") {
        host = host + "-be.glean.com"
    }
    return "https://" + host
}
```

- [ ] Create `internal/auth/storage.go`:

```go
package auth

import (
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
    "time"

    "github.com/zalando/go-keyring"
)

const keyringService = "glean-cli-oauth"

// OAuthToken holds a stored OAuth token set.
type OAuthToken struct {
    AccessToken  string    `json:"access_token"`
    RefreshToken string    `json:"refresh_token,omitempty"`
    Expiry       time.Time `json:"expiry,omitempty"`
    TokenType    string    `json:"token_type"`
    Host         string    `json:"host"`
}

// IsExpired returns true if the token has expired (with a 60s buffer).
func (t *OAuthToken) IsExpired() bool {
    if t.Expiry.IsZero() {
        return false
    }
    return time.Now().Add(60 * time.Second).After(t.Expiry)
}

// SaveOAuthToken stores the token in the OS keyring (primary) with
// fallback to ~/.config/glean-cli/oauth-tokens.json (0600).
func SaveOAuthToken(host string, tok *OAuthToken) error {
    tok.Host = host
    data, err := json.Marshal(tok)
    if err != nil {
        return fmt.Errorf("marshaling token: %w", err)
    }

    // Try keyring first.
    if err := keyring.Set(keyringService, host, string(data)); err == nil {
        return nil
    }

    // Fallback: write to file.
    return saveTokenToFile(host, data)
}

// LoadOAuthToken retrieves a stored OAuth token for the given host.
// Returns nil (no error) if no token is stored.
func LoadOAuthToken(host string) (*OAuthToken, error) {
    // Try keyring.
    if val, err := keyring.Get(keyringService, host); err == nil {
        var tok OAuthToken
        if err := json.Unmarshal([]byte(val), &tok); err == nil {
            return &tok, nil
        }
    }

    // Try file fallback.
    return loadTokenFromFile(host)
}

// DeleteOAuthToken removes the stored OAuth token for the given host.
func DeleteOAuthToken(host string) error {
    _ = keyring.Delete(keyringService, host)
    return deleteTokenFile(host)
}

func tokenFilePath(host string) (string, error) {
    home, err := os.UserHomeDir()
    if err != nil {
        return "", err
    }
    dir := filepath.Join(home, ".config", "glean-cli")
    return filepath.Join(dir, "oauth-"+host+".json"), nil
}

func saveTokenToFile(host string, data []byte) error {
    path, err := tokenFilePath(host)
    if err != nil {
        return err
    }
    if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
        return err
    }
    return os.WriteFile(path, data, 0600)
}

func loadTokenFromFile(host string) (*OAuthToken, error) {
    path, err := tokenFilePath(host)
    if err != nil {
        return nil, err
    }
    data, err := os.ReadFile(path)
    if os.IsNotExist(err) {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    var tok OAuthToken
    if err := json.Unmarshal(data, &tok); err != nil {
        return nil, err
    }
    return &tok, nil
}

func deleteTokenFile(host string) error {
    path, err := tokenFilePath(host)
    if err != nil {
        return err
    }
    err = os.Remove(path)
    if os.IsNotExist(err) {
        return nil
    }
    return err
}
```

- [ ] Create `internal/auth/callback.go`:

```go
package auth

import (
    "context"
    "fmt"
    "net"
    "net/http"
)

// CallbackResult holds the authorization code (or error) from the OAuth callback.
type CallbackResult struct {
    Code  string
    State string
    Err   error
}

// StartCallbackServer starts a local HTTP server on a random port and returns
// the port number and a channel that receives exactly one CallbackResult when
// the browser redirects to /callback.
//
// The server shuts itself down after receiving one request.
func StartCallbackServer(ctx context.Context) (port int, results <-chan CallbackResult, err error) {
    listener, err := net.Listen("tcp", "localhost:0")
    if err != nil {
        return 0, nil, fmt.Errorf("starting callback server: %w", err)
    }
    port = listener.Addr().(*net.TCPAddr).Port

    ch := make(chan CallbackResult, 1)
    srv := &http.Server{}

    mux := http.NewServeMux()
    mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
        q := r.URL.Query()
        if errVal := q.Get("error"); errVal != "" {
            ch <- CallbackResult{Err: fmt.Errorf("OAuth error: %s — %s", errVal, q.Get("error_description"))}
            fmt.Fprintf(w, "<html><body><h2>Authentication failed: %s</h2><p>You may close this tab.</p></body></html>", errVal)
        } else {
            ch <- CallbackResult{Code: q.Get("code"), State: q.Get("state")}
            fmt.Fprint(w, "<html><body><h2>✓ Authenticated with Glean!</h2><p>You may close this tab and return to your terminal.</p></body></html>")
        }
        // Shut down after responding.
        go srv.Shutdown(ctx)
    })
    srv.Handler = mux

    go srv.Serve(listener)
    return port, ch, nil
}
```

- [ ] Create `internal/auth/auth.go`:

```go
package auth

import (
    "context"
    "fmt"
    "time"

    "github.com/pkg/browser"
    "github.com/scalvert/glean-cli/internal/config"
    "golang.org/x/oauth2"
)

// gleanOAuthClientID is the public client ID for the Glean CLI.
// This is discovered from the server metadata, but we use this as
// the registered client_id for the CLI application.
const gleanOAuthClientID = "glean-cli"

// Login performs the OAuth 2.0 PKCE Authorization Code Flow.
// If noBrowser is true, falls back to Device Authorization Flow.
// On success, the token is stored and the user email is returned.
func Login(ctx context.Context, host string, noBrowser bool) error {
    cfg, err := config.LoadConfig()
    if err != nil || host == "" {
        host = cfg.GleanHost
    }
    if host == "" {
        return fmt.Errorf("Glean host not configured; run 'glean config --host <host>' first")
    }

    meta, err := DiscoverMetadata(ctx, host)
    if err != nil {
        return err
    }

    oauthCfg := &oauth2.Config{
        ClientID: gleanOAuthClientID,
        Endpoint: oauth2.Endpoint{
            AuthURL:       meta.AuthorizationEndpoint,
            TokenURL:      meta.TokenEndpoint,
            DeviceAuthURL: meta.DeviceAuthorizationEndpoint,
        },
        Scopes: []string{"openid", "profile", "email"},
    }

    var token *oauth2.Token
    if noBrowser || meta.DeviceAuthorizationEndpoint != "" && meta.AuthorizationEndpoint == "" {
        token, err = deviceFlow(ctx, oauthCfg)
    } else {
        token, err = pkceFlow(ctx, oauthCfg)
    }
    if err != nil {
        return err
    }

    stored := &OAuthToken{
        AccessToken:  token.AccessToken,
        RefreshToken: token.RefreshToken,
        Expiry:       token.Expiry,
        TokenType:    token.TokenType,
        Host:         host,
    }
    if err := SaveOAuthToken(host, stored); err != nil {
        return fmt.Errorf("saving token: %w", err)
    }

    fmt.Printf("✓ Authenticated with Glean (%s)\n", host)
    return nil
}

// pkceFlow runs the Authorization Code + PKCE flow.
func pkceFlow(ctx context.Context, cfg *oauth2.Config) (*oauth2.Token, error) {
    verifier := oauth2.GenerateVerifier()
    state := generateState()

    port, callbackCh, err := StartCallbackServer(ctx)
    if err != nil {
        return nil, err
    }
    cfg.RedirectURL = fmt.Sprintf("http://localhost:%d/callback", port)

    authURL := cfg.AuthCodeURL(state, oauth2.S256ChallengeOption(verifier))

    fmt.Printf("Opening your browser to authenticate with Glean...\n")
    fmt.Printf("If your browser doesn't open, visit:\n  %s\n\n", authURL)
    _ = browser.OpenURL(authURL)

    result := <-callbackCh
    if result.Err != nil {
        return nil, result.Err
    }
    if result.State != state {
        return nil, fmt.Errorf("OAuth state mismatch (possible CSRF)")
    }

    return cfg.Exchange(ctx, result.Code, oauth2.VerifierOption(verifier))
}

// deviceFlow runs the Device Authorization Flow (RFC 8628) for headless environments.
func deviceFlow(ctx context.Context, cfg *oauth2.Config) (*oauth2.Token, error) {
    deviceAuth, err := cfg.DeviceAuth(ctx, oauth2.AccessTypeOffline)
    if err != nil {
        return nil, fmt.Errorf("starting device authorization: %w", err)
    }

    fmt.Printf("Visit this URL in your browser:\n  %s\n\n", deviceAuth.VerificationURIComplete)
    fmt.Printf("And enter the code: %s\n\n", deviceAuth.UserCode)
    fmt.Printf("Waiting for authorization...\n")

    return cfg.DeviceAccessToken(ctx, deviceAuth, oauth2.AccessTypeOffline)
}

// Logout removes the stored OAuth token for the host.
func Logout(host string) error {
    if err := DeleteOAuthToken(host); err != nil {
        return fmt.Errorf("removing token: %w", err)
    }
    fmt.Printf("✓ Logged out from Glean (%s)\n", host)
    return nil
}

// Status prints the current authentication state for the host.
func Status(host string) error {
    tok, err := LoadOAuthToken(host)
    if err != nil {
        return fmt.Errorf("checking auth status: %w", err)
    }
    if tok == nil {
        fmt.Printf("Not logged in to %s\n", host)
        fmt.Printf("Run 'glean auth login' to authenticate\n")
        return nil
    }
    if tok.IsExpired() {
        fmt.Printf("Token for %s is expired\n", host)
        fmt.Printf("Run 'glean auth login' to re-authenticate\n")
        return nil
    }
    expStr := "no expiry"
    if !tok.Expiry.IsZero() {
        expStr = "expires " + tok.Expiry.Format(time.RFC3339)
    }
    fmt.Printf("✓ Authenticated with Glean (%s) — %s\n", host, expStr)
    return nil
}

// EnsureAuth checks that the client has usable credentials.
// Returns nil if GLEAN_API_TOKEN env var is set, or if an OAuth token exists.
// Returns an error with "run 'glean auth login'" guidance if nothing is configured.
func EnsureAuth(ctx context.Context) error {
    cfg, err := config.LoadConfig()
    if err == nil && cfg.GleanToken != "" {
        return nil // env var or legacy token configured
    }
    if err == nil && cfg.GleanHost != "" {
        tok, err := LoadOAuthToken(cfg.GleanHost)
        if err == nil && tok != nil && !tok.IsExpired() {
            return nil
        }
    }
    return fmt.Errorf("not authenticated — run 'glean auth login' or set GLEAN_API_TOKEN")
}

// generateState creates a random OAuth state parameter to prevent CSRF.
func generateState() string {
    // Use oauth2's verifier generation for the same randomness quality.
    // The verifier format (base64url, 32 random bytes) is fine for state too.
    return oauth2.GenerateVerifier()[:16]
}
```

- [ ] Build: `go build ./internal/auth/...`
- [ ] Fix any compilation errors
- [ ] Commit:
```bash
git add internal/auth/
git commit -m "feat(auth): add internal/auth package with PKCE, device flow, discovery, and token storage"
```

---

## Chunk 2: Auth Tests

### Task 3: Write discovery tests

**Files:**
- Create: `internal/auth/discovery_test.go`

- [ ] Create `internal/auth/discovery_test.go`:

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

func TestDiscoverMetadata_Success(t *testing.T) {
    srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        assert.Equal(t, "/.well-known/oauth-authorization-server", r.URL.Path)
        json.NewEncoder(w).Encode(authServerMetadata{
            AuthorizationEndpoint: "https://auth.example.com/authorize",
            TokenEndpoint:         "https://auth.example.com/token",
        })
    }))
    defer srv.Close()

    // Override hostToBaseURL by directly calling discoverMetadataFromURL (test helper)
    meta, err := discoverMetadataFromURL(context.Background(), srv.URL)
    require.NoError(t, err)
    assert.Equal(t, "https://auth.example.com/authorize", meta.AuthorizationEndpoint)
    assert.Equal(t, "https://auth.example.com/token", meta.TokenEndpoint)
}

func TestDiscoverMetadata_404(t *testing.T) {
    srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusNotFound)
    }))
    defer srv.Close()

    _, err := discoverMetadataFromURL(context.Background(), srv.URL)
    require.Error(t, err)
    assert.Contains(t, err.Error(), "does not support OAuth")
}

func TestDiscoverMetadata_Incomplete(t *testing.T) {
    srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(map[string]string{
            "authorization_endpoint": "https://auth.example.com/authorize",
            // missing token_endpoint
        })
    }))
    defer srv.Close()

    _, err := discoverMetadataFromURL(context.Background(), srv.URL)
    require.Error(t, err)
    assert.Contains(t, err.Error(), "incomplete OAuth metadata")
}

func TestHostToBaseURL(t *testing.T) {
    assert.Equal(t, "https://linkedin-be.glean.com", hostToBaseURL("linkedin"))
    assert.Equal(t, "https://mycompany-be.glean.com", hostToBaseURL("mycompany"))
    assert.Equal(t, "https://foo.bar.com", hostToBaseURL("foo.bar.com"))
    assert.Equal(t, "https://linkedin-be.glean.com", hostToBaseURL("linkedin-be.glean.com"))
}
```

Note: the tests call `discoverMetadataFromURL` which you need to add to `discovery.go` as an internal helper that takes a base URL instead of a host, so tests can use `httptest.NewServer`. Add this to `discovery.go`:

```go
// discoverMetadataFromURL fetches metadata from the given base URL directly.
// Used in tests to inject a mock server.
func discoverMetadataFromURL(ctx context.Context, baseURL string) (*authServerMetadata, error) {
    url := baseURL + "/.well-known/oauth-authorization-server"
    req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
    if err != nil {
        return nil, fmt.Errorf("building discovery request: %w", err)
    }
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("fetching OAuth metadata: %w", err)
    }
    defer resp.Body.Close()
    if resp.StatusCode == http.StatusNotFound {
        return nil, fmt.Errorf("this Glean instance does not support OAuth; use 'glean config --token' instead")
    }
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("OAuth discovery returned HTTP %d", resp.StatusCode)
    }
    var meta authServerMetadata
    if err := json.NewDecoder(resp.Body).Decode(&meta); err != nil {
        return nil, fmt.Errorf("parsing OAuth metadata: %w", err)
    }
    if meta.AuthorizationEndpoint == "" || meta.TokenEndpoint == "" {
        return nil, fmt.Errorf("incomplete OAuth metadata from %s", url)
    }
    return &meta, nil
}

// DiscoverMetadata is the public entry point that constructs the URL from the host.
func DiscoverMetadata(ctx context.Context, host string) (*authServerMetadata, error) {
    return discoverMetadataFromURL(ctx, hostToBaseURL(host))
}
```

- [ ] Run: `go test ./internal/auth/... -run TestDiscover -v`
- [ ] Expected: all pass

---

### Task 4: Write storage tests

**Files:**
- Create: `internal/auth/storage_test.go`

- [ ] Create `internal/auth/storage_test.go`:

```go
package auth

import (
    "os"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestSaveAndLoadOAuthToken_FileStorage(t *testing.T) {
    // Use a temp dir to avoid touching real config
    tmp := t.TempDir()
    origHome := os.Getenv("HOME")
    t.Setenv("HOME", tmp)
    defer func() { os.Setenv("HOME", origHome) }()

    host := "test.glean.com"
    tok := &OAuthToken{
        AccessToken:  "access-abc",
        RefreshToken: "refresh-xyz",
        Expiry:       time.Now().Add(time.Hour),
        TokenType:    "Bearer",
        Host:         host,
    }

    err := SaveOAuthToken(host, tok)
    require.NoError(t, err)

    loaded, err := LoadOAuthToken(host)
    require.NoError(t, err)
    require.NotNil(t, loaded)
    assert.Equal(t, "access-abc", loaded.AccessToken)
    assert.Equal(t, "refresh-xyz", loaded.RefreshToken)
    assert.Equal(t, "Bearer", loaded.TokenType)
}

func TestLoadOAuthToken_NotFound(t *testing.T) {
    tmp := t.TempDir()
    t.Setenv("HOME", tmp)

    tok, err := LoadOAuthToken("nonexistent.glean.com")
    require.NoError(t, err)
    assert.Nil(t, tok)
}

func TestDeleteOAuthToken(t *testing.T) {
    tmp := t.TempDir()
    t.Setenv("HOME", tmp)

    host := "delete-test.glean.com"
    tok := &OAuthToken{AccessToken: "to-delete", Host: host}
    require.NoError(t, SaveOAuthToken(host, tok))

    require.NoError(t, DeleteOAuthToken(host))

    loaded, err := LoadOAuthToken(host)
    require.NoError(t, err)
    assert.Nil(t, loaded)
}

func TestOAuthToken_IsExpired(t *testing.T) {
    expired := &OAuthToken{Expiry: time.Now().Add(-time.Hour)}
    assert.True(t, expired.IsExpired())

    valid := &OAuthToken{Expiry: time.Now().Add(time.Hour)}
    assert.False(t, valid.IsExpired())

    noExpiry := &OAuthToken{}
    assert.False(t, noExpiry.IsExpired())
}
```

Note: `SaveOAuthToken` tries the keyring first. In tests the keyring may fail (no system keyring in CI). To ensure the file fallback is tested, the test sets HOME to a temp dir and lets the keyring fail silently. Verify this works by checking `SaveOAuthToken`'s fallback logic — it calls `saveTokenToFile` if keyring fails.

- [ ] Run: `go test ./internal/auth/... -run TestSave -v`
- [ ] Fix if keyring attempts cause issues in CI-like environment (the keyring failure should silently fall back to file)
- [ ] Run: `go test ./internal/auth/... -v`
- [ ] Expected: all pass
- [ ] Commit:
```bash
git add internal/auth/
git commit -m "test(auth): add discovery and token storage tests"
```

---

### Task 5: Write callback server test

**Files:**
- Create: `internal/auth/callback_test.go`

- [ ] Create `internal/auth/callback_test.go`:

```go
package auth

import (
    "context"
    "fmt"
    "net/http"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestStartCallbackServer_SuccessCode(t *testing.T) {
    ctx := context.Background()
    port, ch, err := StartCallbackServer(ctx)
    require.NoError(t, err)
    assert.Greater(t, port, 0)

    // Simulate the browser redirect
    go func() {
        resp, err := http.Get(fmt.Sprintf("http://localhost:%d/callback?code=test-code&state=test-state", port))
        if err == nil {
            resp.Body.Close()
        }
    }()

    result := <-ch
    assert.NoError(t, result.Err)
    assert.Equal(t, "test-code", result.Code)
    assert.Equal(t, "test-state", result.State)
}

func TestStartCallbackServer_OAuthError(t *testing.T) {
    ctx := context.Background()
    port, ch, err := StartCallbackServer(ctx)
    require.NoError(t, err)

    go func() {
        resp, err := http.Get(fmt.Sprintf("http://localhost:%d/callback?error=access_denied&error_description=User+denied", port))
        if err == nil {
            resp.Body.Close()
        }
    }()

    result := <-ch
    require.Error(t, result.Err)
    assert.Contains(t, result.Err.Error(), "access_denied")
}
```

- [ ] Run: `go test ./internal/auth/... -run TestStartCallback -v`
- [ ] Expected: both pass
- [ ] Commit:
```bash
git add internal/auth/callback_test.go
git commit -m "test(auth): add callback server tests"
```

---

## Chunk 3: Update Client Auth Priority

### Task 6: Update internal/client/client.go for OAuth token priority

**Files:**
- Modify: `internal/client/client.go`

The new auth priority: env var → OAuth token → legacy static token.

- [ ] Read `internal/client/client.go`
- [ ] Add an import for `"github.com/scalvert/glean-cli/internal/auth"`
- [ ] Modify `New()` to check for an OAuth token if no static token is found:

```go
func New(cfg *config.Config) (*glean.Glean, error) {
    if cfg.GleanHost == "" {
        return nil, fmt.Errorf("Glean host not configured. Run 'glean config --host <host>' or set GLEAN_HOST")
    }

    token := cfg.GleanToken

    // If no static token, check for OAuth token
    if token == "" {
        oauthTok, err := auth.LoadOAuthToken(cfg.GleanHost)
        if err == nil && oauthTok != nil && !oauthTok.IsExpired() {
            token = oauthTok.AccessToken
        }
        // If expired and has refresh token, attempt refresh
        // (token refresh will be added in a follow-up; for now, prompt re-login)
        if err == nil && oauthTok != nil && oauthTok.IsExpired() {
            return nil, fmt.Errorf("OAuth token expired — run 'glean auth login' to re-authenticate")
        }
    }

    if token == "" {
        return nil, fmt.Errorf("not authenticated. Run 'glean auth login' or set GLEAN_API_TOKEN")
    }

    instance := extractInstance(cfg.GleanHost)
    // ... rest of existing code unchanged ...
```

- [ ] Note: update the error message on empty token to suggest `glean auth login` as the primary path
- [ ] Build: `go build -o glean-dev .`
- [ ] Run: `go test ./... -race`
- [ ] Expected: all existing tests still pass (they use `SetupTestWithResponse` which sets env vars)
- [ ] Commit:
```bash
git add internal/client/client.go
git commit -m "feat(client): add OAuth token priority in auth chain (env → oauth → legacy token)"
```

---

## Chunk 4: cmd/auth.go Command Group

### Task 7: Create cmd/auth.go

**Files:**
- Create: `cmd/auth.go`
- Modify: `cmd/root.go`

- [ ] Create `cmd/auth.go`:

```go
package cmd

import (
    "github.com/MakeNowJust/heredoc"
    "github.com/scalvert/glean-cli/internal/auth"
    "github.com/scalvert/glean-cli/internal/config"
    "github.com/spf13/cobra"
)

// NewCmdAuth creates and returns the auth command group.
func NewCmdAuth() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "auth",
        Short: "Authenticate with Glean",
        Long: heredoc.Doc(`
            Manage authentication with your Glean instance.

            Use 'glean auth login' to authenticate via your browser (recommended).
            For CI/CD environments, set the GLEAN_API_TOKEN environment variable instead.
        `),
    }

    cmd.AddCommand(newAuthLoginCmd())
    cmd.AddCommand(newAuthLogoutCmd())
    cmd.AddCommand(newAuthStatusCmd())
    return cmd
}

func newAuthLoginCmd() *cobra.Command {
    var noBrowser bool

    cmd := &cobra.Command{
        Use:   "login",
        Short: "Authenticate with Glean via your browser",
        Long: heredoc.Doc(`
            Opens your browser to authenticate with Glean using OAuth 2.0.

            After authentication, credentials are stored securely in your system keyring.
            Subsequent commands will use these credentials automatically.

            For headless environments (SSH, CI), use --no-browser to get a device code
            to enter on another device, or set GLEAN_API_TOKEN instead.

            Examples:
              # Interactive browser-based login (recommended)
              glean auth login

              # Headless/SSH environment
              glean auth login --no-browser
        `),
        RunE: func(cmd *cobra.Command, args []string) error {
            cfg, err := config.LoadConfig()
            if err != nil || cfg.GleanHost == "" {
                return fmt.Errorf("Glean host not configured; run 'glean config --host <host>' first")
            }
            return auth.Login(cmd.Context(), cfg.GleanHost, noBrowser)
        },
    }

    cmd.Flags().BoolVar(&noBrowser, "no-browser", false, "Use device flow instead of opening a browser")
    return cmd
}

func newAuthLogoutCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "logout",
        Short: "Remove stored Glean credentials",
        Long: heredoc.Doc(`
            Removes the OAuth token stored by 'glean auth login'.

            This does not revoke the token on the server side (Glean does not
            expose a revocation endpoint). The token will expire naturally.

            For tokens set via GLEAN_API_TOKEN or 'glean config --token', use
            'glean config --clear' instead.
        `),
        RunE: func(cmd *cobra.Command, args []string) error {
            cfg, err := config.LoadConfig()
            if err != nil || cfg.GleanHost == "" {
                return fmt.Errorf("Glean host not configured")
            }
            return auth.Logout(cfg.GleanHost)
        },
    }
}

func newAuthStatusCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "status",
        Short: "Show current authentication status",
        Long: heredoc.Doc(`
            Displays the current authentication state for your configured Glean instance.

            Shows whether you are authenticated, the token expiry time, and the host.
        `),
        RunE: func(cmd *cobra.Command, args []string) error {
            cfg, err := config.LoadConfig()
            if err != nil || cfg.GleanHost == "" {
                return fmt.Errorf("Glean host not configured; run 'glean config --host <host>' first")
            }
            return auth.Status(cfg.GleanHost)
        },
    }
}
```

Note: add `"fmt"` to the imports in this file.

- [ ] Add `NewCmdAuth()` to `cmd/root.go` `AddCommand` block, with `GroupID = "core"` if you implemented command groups

- [ ] Build: `go build -o glean-dev .`
- [ ] Test: `./glean-dev auth --help` — should show login/logout/status subcommands
- [ ] Test: `./glean-dev auth status` — should show "Not logged in" (since no token yet)
- [ ] Commit:
```bash
git add cmd/auth.go cmd/root.go
git commit -m "feat(auth): add glean auth login|logout|status command group"
```

---

## Chunk 5: End-to-End Verification

### Task 8: Full OAuth flow test

These are manual tests since the PKCE flow requires a real browser and a real Glean instance.

- [ ] Build final binary:
  ```bash
  go build -o glean-dev .
  ```
- [ ] Run all automated tests:
  ```bash
  go test ./... -race
  ```
  Expected: all pass

- [ ] **Test: auth status before login**
  ```bash
  ./glean-dev auth status
  ```
  Expected: "Not logged in to <host>"

- [ ] **Test: auth login (PKCE flow)**
  ```bash
  ./glean-dev auth login
  ```
  Expected:
  - "Opening your browser to authenticate with Glean..." is printed
  - Browser opens to Glean's OAuth page
  - After completing browser auth, terminal shows "✓ Authenticated with Glean (<host>)"

- [ ] **Test: auth status after login**
  ```bash
  ./glean-dev auth status
  ```
  Expected: "✓ Authenticated with Glean (<host>) — expires <timestamp>"

- [ ] **Test: search using OAuth token**
  ```bash
  ./glean-dev search "test"
  ```
  Expected: results returned without error (OAuth token used automatically)

- [ ] **Test: TUI using OAuth token**
  ```bash
  ./glean-dev
  ```
  Expected: TUI opens and chat works

- [ ] **Test: auth logout**
  ```bash
  ./glean-dev auth logout
  ```
  Expected: "✓ Logged out from Glean (<host>)"

- [ ] **Test: status after logout**
  ```bash
  ./glean-dev auth status
  ```
  Expected: "Not logged in to <host>"

- [ ] **Test: device flow**
  ```bash
  ./glean-dev auth login --no-browser
  ```
  Expected: URL and code printed; completing on another device authenticates

- [ ] **Test: CI escape hatch still works**
  ```bash
  GLEAN_API_TOKEN=<your-token> ./glean-dev search "test"
  ```
  Expected: works (env var takes precedence over OAuth)

- [ ] Update EVAL-CHECKLIST.md for OAuth items
- [ ] Commit:
  ```bash
  git add EVAL-CHECKLIST.md
  git commit -m "chore: mark OAuth auth items as closed in eval checklist"
  ```

---

## Token Refresh (Follow-Up)

Token refresh is intentionally deferred. The current implementation:
- Returns a clear error message when an OAuth token is expired
- Instructs user to run `glean auth login`

Full automatic refresh (using `oauth2.TokenSource`) can be added in a follow-up once the basic flow is confirmed working, since it requires storing the `oauth2.Config` alongside tokens and is non-trivial to test.

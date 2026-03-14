# Design: OAuth 2.0 Auth for glean-cli

**Date:** 2026-03-13
**Status:** Approved

---

## Overview

Add `glean auth` subcommand group implementing standards-compliant OAuth 2.0 with PKCE and OIDC for the glean-cli. Replaces static-token-as-primary-UX with a smooth interactive login flow.

Three auth paths are supported through a single discovery chain:
1. **API token** — static token, for CI/CD and instances without OAuth
2. **OAuth + DCR** — dynamic client registration, fully auto-discovered (most common)
3. **OAuth + static client** — pre-registered client_id/secret, also auto-discovered

All OAuth paths begin with RFC 9728 protected resource metadata discovery, which transparently routes to the right path. The user never needs to know which one their instance uses.

---

## Libraries

No hand-rolled protocol code. Three libraries cover all OAuth/OIDC work:

| Library | Purpose |
|---------|---------|
| `golang.org/x/oauth2` | PKCE primitives (`GenerateVerifier`, `S256ChallengeOption`, `VerifierOption`), token exchange, refresh |
| `github.com/int128/oauth2cli` | Local callback server on random port + browser opening — `GetToken()` does the full PKCE browser flow |
| `github.com/coreos/go-oidc/v3/oidc` | OIDC discovery via `NewProvider()`, ID token verification, UserInfo (email extraction) |

Minimal stdlib HTTP (~60 lines total, no protocol logic):
- **Protected resource metadata** — `GET /.well-known/oauth-protected-resource` → parse `authorization_servers[]`
- **DCR** — `POST registration_endpoint` → parse `client_id` + optional `client_secret`
- **Domain lookup** — `POST https://app.glean.com/config/search` → get backend URL from email domain

---

## Commands

```
glean auth login              # interactive PKCE flow (default)
glean auth login --no-browser # print URL + wait (headless/SSH environments)
glean auth logout             # remove stored OAuth tokens
glean auth status             # show auth state, expiry, email
```

For static client instances, user pre-configures once:
```
glean config --oauth-client-id <id>
glean config --oauth-client-secret <secret>   # optional, for confidential clients
```

---

## Auth Priority (updated client.go)

```
1. GLEAN_API_TOKEN env var      → use directly as Bearer token (CI/CD)
2. OAuth access token           → from keyring/file (written by glean auth login)
3. Legacy static token          → from keyring/file (glean config --token, migration path)
```

---

## Login Flow

### Entry

```
glean auth login
```

**If `GLEAN_HOST` already configured** → skip to discovery.

**If no host configured** → prompt interactively:
```
Enter your work email: steve@company.com
```
→ `POST https://app.glean.com/config/search {"emailDomain": "company.com", ...}`
→ Extract `queryURL` as backend base URL, store as `GLEAN_HOST`.

### Discovery Chain

```
GET <backend>/.well-known/oauth-protected-resource
→ { "authorization_servers": ["https://..."] }

GET <auth-server>/.well-known/oauth-authorization-server
→ { authorization_endpoint, token_endpoint, registration_endpoint?, ... }
```

`oidc.NewProvider()` drives the second step and gives us endpoint URLs + OIDC
ID token verifier in one call.

### Client ID Resolution

**DCR path** (registration_endpoint present):
1. Check `~/.local/state/glean-cli/<host>/client.json` for a previously registered client_id.
2. If none: `POST registration_endpoint` with `client_name: "glean-cli"`, `redirect_uris`, `grant_types: ["authorization_code", "refresh_token"]`, `token_endpoint_auth_method: "none"`.
3. Store returned `client_id` (and `client_secret` if provided) in `client.json` (0600).

**Static client path** (no registration_endpoint):
- Use `client_id` from `glean config --oauth-client-id`.
- If none configured → fall back to token prompt (see below).

### PKCE Browser Flow

`oauth2cli.GetToken()` handles this entirely:
- Generates `state` + `code_verifier` + `code_challenge` (S256).
- Starts local HTTP server on `127.0.0.1:<random-port>`.
- Opens browser to `authorization_endpoint?...&code_challenge=...&scope=openid email profile`.
- Waits for redirect to `http://127.0.0.1:<port>/callback?code=...`.
- Exchanges code for tokens via `token_endpoint`.

**`--no-browser` mode**: print the authorization URL, wait for user to visit it
manually and paste back the authorization code. `oauth2cli` supports this via
`LocalServerReadyChan` — we print the URL when the server is ready and the
callback still handles the redirect automatically once the user completes login.

### OIDC: Email Extraction

After receiving tokens, if an `id_token` is present:
```go
idToken, err := verifier.Verify(ctx, rawIDToken)
var claims struct{ Email string `json:"email"` }
idToken.Claims(&claims)
```

If no `id_token` (non-OIDC instance), call `provider.UserInfo()` as fallback.

### Token Storage

Per-host directory: `~/.local/state/glean-cli/<host-hash>/`

```
server.json   — { protectedResourceMetadata, authServerMetadata } (0600)
client.json   — { client_id, client_secret? }                     (0600)
tokens.json   — { access_token, refresh_token, expiry, email }    (0600)
```

Directory permissions: `0700`. Atomic writes (write to `.tmp`, rename).

### Success

```
✓ Authenticated as steve@company.com (company-be.glean.com)
```

---

## OAuth Not Supported (no authorization_endpoint)

Stay in `glean auth login`, prompt inline:

```
This Glean instance doesn't support OAuth.
Contact your Glean administrator to generate an API token.
  (Glean Admin → Settings → API Tokens)

Token: ████
```

Saves the token via existing `config.SaveConfig()`. Same as `glean config --token`
but integrated into the auth flow so the user doesn't need a second command.

---

## Token Refresh

On every command execution, `internal/client/client.go` checks the stored token
expiry (with 60s buffer). If expired and a `refresh_token` exists:

```go
ts := oauthCfg.TokenSource(ctx, &oauth2.Token{RefreshToken: storedRefreshToken})
newToken, err := ts.Token()
```

`golang.org/x/oauth2` handles the refresh transparently. Updated tokens are
written back to `tokens.json`. If refresh fails (revoked), return error:
`"Session expired — run 'glean auth login' to re-authenticate"`.

---

## glean auth status

```
✓ Authenticated as steve@company.com (company-be.glean.com)
  Token expires: 2026-03-14 09:23:11 UTC (in 1h 4m)
  Auth method: OAuth (DCR)
```

Or if not logged in:
```
Not authenticated.
Run 'glean auth login' to authenticate.
```

---

## New Package: internal/auth/

```
internal/auth/
  auth.go         — Login(), Logout(), Status(), EnsureAuth(), loadStoredToken()
  discovery.go    — fetchProtectedResource(), DCR() (~60 lines, 2 HTTP calls)
  domainlookup.go — LookupBackendURL(email) (~25 lines, 1 HTTP call)
  storage.go      — loadTokens(), saveTokens(), loadClient(), saveClient(), loadServer(), saveServer()
```

`oauth2cli` and `coreos/go-oidc` are used directly in `auth.go`. No abstraction
layers over them — they're the implementation.

---

## Files Changed

**New:**
- `cmd/auth.go` — `glean auth` command group
- `internal/auth/auth.go`
- `internal/auth/discovery.go`
- `internal/auth/domainlookup.go`
- `internal/auth/storage.go`
- `internal/auth/*_test.go`

**Modified:**
- `internal/client/client.go` — OAuth token in priority chain, auto-refresh
- `cmd/root.go` — register `NewCmdAuth()` as core command
- `go.mod` — add `oauth2cli`, `coreos/go-oidc/v3`
- `README.md` — document `glean auth login` as the recommended setup step

---

## Removed (done)

- `cmd/mcp.go` — deleted
- `internal/mcp/` — deleted
- `github.com/mark3labs/mcp-go` — removed from go.mod

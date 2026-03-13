# Design: OAuth 2.0 Auth + Pre-Release Bug Fix Sprint

**Date:** 2026-03-13
**Status:** Approved

---

## Overview

Two parallel tracks:

1. **OAuth 2.0 auth** — Add `glean auth` subcommand group implementing RFC 8252 (OAuth for Native Apps) + RFC 7636 (PKCE). This replaces static-token-as-primary-UX with a proper interactive login flow, while keeping `GLEAN_API_TOKEN` env var as a CI/CD escape hatch.

2. **CHK bug fixes** — Address all 31 items from the pre-release eval checklist, spanning release infrastructure, correctness bugs, documentation, and test coverage.

---

## Part 1: OAuth 2.0 Auth

### Design Principles

- Use standard libraries, not hand-rolled crypto. `golang.org/x/oauth2` provides `GenerateVerifier()`, `S256ChallengeOption()`, `DeviceAuth()`, token refresh — all RFC-compliant.
- Follow the gh CLI / Stripe CLI / Vercel CLI model: `auth login` is the interactive path, env var is the CI path.
- Support both PKCE Authorization Code Flow (default, browser-based) and Device Authorization Flow (`--no-browser`, for SSH/headless).

### New Commands

```
glean auth login           # PKCE flow (opens browser)
glean auth login --no-browser  # Device flow (user_code + polling)
glean auth logout          # Revokes and clears stored tokens
glean auth status          # Shows auth state, token expiry, scopes
```

### Auth Priority (updated client.go)

```
1. GLEAN_API_TOKEN env var  → use as Bearer token directly (no keyring, for CI)
2. OAuth access token       → from keyring (written by `glean auth login`)
3. Legacy static token      → from keyring/config file (migration path for existing users)
```

`glean config --token` remains for power users and migration, but `glean auth login` becomes the documented recommended path.

### PKCE Flow (default)

Libraries:
- `golang.org/x/oauth2` — PKCE primitives (`GenerateVerifier`, `S256ChallengeOption`), token exchange, refresh
- `github.com/pkg/browser` — cross-platform `browser.OpenURL()`
- `net/http` (stdlib) — local callback server (~15 lines)

Flow:
1. `glean auth login` discovers OAuth endpoints from `https://<host>-be.glean.com/.well-known/oauth-authorization-server`
2. Generates `code_verifier` (via `oauth2.GenerateVerifier()`) and `code_challenge` (via `S256ChallengeOption`)
3. Starts `net/http` server on a random available port (`localhost:0` → get actual port)
4. Opens `https://<auth-endpoint>?response_type=code&client_id=...&code_challenge=...&redirect_uri=http://localhost:<port>/callback` in the browser via `browser.OpenURL()`
5. Prints "Opening your browser to authenticate with Glean..." with the URL as fallback
6. Waits for the redirect to `localhost:<port>/callback?code=...`
7. Exchanges code + verifier via `oauth2.Config.Exchange()` with `oauth2.VerifierOption(verifier)`
8. Stores `access_token`, `refresh_token`, `expiry` in OS keyring (primary) → `~/.config/glean-cli/oauth-tokens.json` 0600 (fallback)
9. Prints "✓ Authenticated as <email> (<host>)"

### Device Flow (`--no-browser`)

```go
deviceAuth, err := oauthConfig.DeviceAuth(ctx, oauth2.S256ChallengeOption(verifier))
// Print: "Visit <verification_uri> and enter code: <user_code>"
token, err := oauthConfig.DeviceAccessToken(ctx, deviceAuth, oauth2.VerifierOption(verifier))
```

### Token Storage

```
Primary:  OS keyring (service: "glean-cli-oauth", user: "<host>")
Fallback: ~/.config/glean-cli/oauth-tokens.json  (0600, dir 0700)
```

Stored as JSON: `{"access_token": "...", "refresh_token": "...", "expiry": "2026-04-01T12:00:00Z", "token_type": "Bearer"}`

Token refresh happens automatically in `client.go` when the access token is within 60s of expiry.

### New Package: `internal/auth/`

```
internal/auth/
  auth.go          — Login() / Logout() / Status() / EnsureAuth()
  discovery.go     — fetchOAuthMetadata(host) → (authURL, tokenURL, clientID)
  storage.go       — loadOAuthToken(host) / saveOAuthToken(host, token)
  callback.go      — startCallbackServer() → (port, codeChan)
```

### OAuth Discovery

Discovery uses the full configured host (not the extracted instance name):
```
GET https://<full-host>/.well-known/oauth-authorization-server
→ { "authorization_endpoint": "...", "token_endpoint": "...", "device_authorization_endpoint": "..." }
```

Where `<full-host>` is:
- `linkedin-be.glean.com` if user configured `linkedin-be.glean.com` or `linkedin`
- `foo.bar.com` if user configured a custom hostname

`discovery.go` must construct the base URL from `cfg.GleanHost` directly (applying the same `-be.glean.com` expansion as `extractInstance` in reverse), **not** from the extracted instance name. Concretely: if `cfg.GleanHost` contains a dot, use it as-is; otherwise expand `<host>` → `<host>-be.glean.com`.

If discovery fails (pre-OAuth Glean instance), `glean auth login` returns a clear error: "This Glean instance does not support OAuth. Use `glean config --token` instead."

---

## Part 2: CHK Bug Fixes

### P0 — Release Blockers

| ID | Fix |
|----|-----|
| CHK-001 | Fix `$${ secrets.GITHUB_TOKEN }}` → `${{ secrets.GITHUB_TOKEN }}` in `release.yml:25` |
| CHK-002 | Fix `install.sh` URL OS case: remove `tr '[:upper:]' '[:lower:]'` so `Darwin`/`Linux` matches GoReleaser archive naming |
| CHK-003 | Remove `--template` and `--person` from README |

### P1 — Correctness Bugs

| ID | Fix |
|----|-----|
| CHK-004 | `cmd/api.go:98` — add `term.IsTerminal(int(os.Stdin.Fd()))` guard before `io.ReadAll(os.Stdin)` |
| CHK-005/006 | `internal/config/config.go` — rewrite `LoadConfig` to merge each field independently across all three sources (env → keyring → file), not gate entire sources on token presence |
| CHK-007 | `cmd/chat.go:259,261` — replace `fmt.Println()` with `fmt.Fprintln(s.cmd.OutOrStdout())` |
| CHK-008 | `internal/tui/model.go:113` — only handle `?` key when textarea is not focused; check `!m.textarea.Focused()` |
| CHK-009 | `internal/client/client.go` — pass `GleanPort` to SDK via custom server URL if set, or document port only applies to `glean api`; remove from `glean config --port` if not SDK-supported |
| CHK-010 | `cmd/api.go:189` — replace `http.DefaultClient` with a client that has a 30s timeout matching SDK |
| CHK-011 | `cmd/api.go:apiBaseURL` — apply same host normalization as `client.extractInstance` |
| CHK-012 | `cmd/api.go:previewRequest` — replace `fmt.Printf` with `fmt.Fprintf(cmd.OutOrStdout(), ...)` |
| CHK-022 | `internal/tui/model.go:234` — pass a `context.Context` into `Model` during `New()` (via `tea.WithContext` on the program, then store the ctx on the model); `callAPI()` uses that ctx so ctrl+c cancels in-flight requests |
| CHK-auth-gate | `cmd/root.go` TUI path — call `auth.EnsureAuth(ctx)` before `tui.New()` so users without any credentials get a clear "run `glean auth login`" error instead of a cryptic SDK failure |

### P1 — Infrastructure / Docs

| ID | Fix |
|----|-----|
| CHK-013 | Rewrite `README.md` — document TUI, all commands, core flags, auth setup |
| CHK-014 | `CONTRIBUTING.md:9` — Go 1.24 or higher |
| CHK-015 | `CONTRIBUTING.md:118` — remove CHANGELOG.md reference |
| CHK-016 | `install.sh:58` — scope chown to the glean binary only |
| CHK-017 | `install.sh:38` — remove `tar -tvf` debug line |
| CHK-018 | `.github/workflows/ci.yml:16` — update `go-version` to `["1.24"]` |
| CHK-023 | Re-record `demo/glean.tape` and regenerate `demo/readme.gif` after all fixes land |

### P1 — Features / UX

| ID | Fix |
|----|-----|
| CHK-019 | Add `cmd/search_test.go` with MockTransport — basic query, `--dry-run`, `--fields`, `--output ndjson`, missing-query error |
| CHK-020 | Register `cmd.GenBashCompletionCmd()` or use `root.InitDefaultCompletionCmd()` in `root.go` |
| CHK-021 | Add Cobra command groups to separate core from namespace commands in `--help` |

### P2 — Nice to Have (include if time allows)

CHK-024 through CHK-031 — session cleanup, dead code, dry-run completeness, namespace command help text, GLEAN_HOST docs, skills/ accuracy.

---

## Part 3: Testing Strategy

### Unit Tests (new files)

- `cmd/search_test.go` — primary user flow (CHK-019)
- `internal/auth/auth_test.go` — login/logout/status, token load/save, refresh logic
- `internal/auth/discovery_test.go` — OAuth metadata discovery with mock HTTP server
- `internal/config/config_test.go` — extend existing tests with cross-source merge cases (CHK-005/006)
- `cmd/api_test.go` — extend to test TTY guard (CHK-004), previewRequest output (CHK-012)

### Integration / E2E Verification

Each fix verified by building (`go build -o glean .`) and running the binary locally against the actual Glean instance. Specific checks:

```bash
# CHK-004: api stdin hang
./glean api users/me          # should error immediately, not hang

# CHK-005/006: config priority
GLEAN_API_TOKEN=xxx ./glean search "test"    # should use keyring host

# CHK-008: TUI ? key
./glean                       # type ? in message → should insert ?, not toggle help

# CHK-020: completion
./glean completion bash       # should output bash completion script

# Auth flow
./glean auth login            # should open browser and return "✓ Authenticated"
./glean auth status           # should show token info
./glean auth logout           # should clear tokens
./glean search "test"         # should work using OAuth token
```

---

## Libraries

| Library | Purpose | Already in go.mod? |
|---------|---------|-------------------|
| `golang.org/x/oauth2` | PKCE, token exchange, device flow, refresh | No — add |
| `github.com/pkg/browser` | Cross-platform browser opening | No — add |
| `net/http` (stdlib) | Local callback server | Yes (stdlib) |
| `golang.org/x/term` | TTY detection (CHK-004) | Already transitive dep |

---

## File Change Summary

**New files:**
- `cmd/auth.go` — `glean auth` command group
- `internal/auth/auth.go`
- `internal/auth/discovery.go`
- `internal/auth/storage.go`
- `internal/auth/callback.go`
- `internal/auth/auth_test.go`
- `cmd/search_test.go`

**Modified files:**
- `internal/client/client.go` — updated auth priority chain
- `internal/config/config.go` — fix cross-source merge (CHK-005/006)
- `cmd/api.go` — TTY guard, timeout, host normalization, previewRequest fix
- `cmd/chat.go` — fmt.Println fix
- `cmd/root.go` — command groups, completion registration
- `internal/tui/model.go` — ? key fix, context cancellation
- `README.md` — full rewrite
- `CONTRIBUTING.md` — Go version, CHANGELOG reference
- `install.sh` — OS case fix, chown fix, debug line removal
- `.github/workflows/ci.yml` — Go version
- `.github/workflows/release.yml` — `$$` typo fix

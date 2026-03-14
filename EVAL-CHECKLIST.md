# Eval Checklist — glean-cli Pre-Release Readiness

Generated: 2026-03-13 | Last updated: 2026-03-14 | Panel v2: 5-agent verification run (tui-correctness, test-coverage, auth-correctness, p2-verifier, code-quality)

## How to read this
- P0 = release blocker; P1 = should fix before release; P2 = nice to have
- Each item has a unique ID (CHK-NNN), acceptance criteria, and the agents who found it
- Items are OPEN until acceptance criteria are verifiably met

---

## P0 — Release Blockers

*(all closed — see Closed section)*

---

## P1 — Should Fix Before Release

*(all closed — see Closed section)*

---

## P2 — Nice to Have

*(all closed — see Closed section)*

---

## Closed

- [x] **CHK-001** `release.yml` has `$${ secrets.GITHUB_TOKEN }}` — double `$` and mismatched braces
  - **File:** `.github/workflows/release.yml:25`
  - **Acceptance:** Expression corrected to `${{ secrets.GITHUB_TOKEN }}`; release workflow runs successfully on a test tag
  - **Opened:** 2026-03-13 | **Agents:** release (primary), confirmed by correctness, docs
  - **Closed:** 2026-03-13 | **How:** Fixed GITHUB_TOKEN expression syntax in release.yml

- [x] **CHK-002** `install.sh` constructs download URL with lowercase OS (`darwin`) but GoReleaser archives use title-case (`Darwin`) — every manual install 404s
  - **File:** `install.sh:23` vs `.goreleaser.yml:21-25`
  - **Acceptance:** OS name in URL matches GoReleaser archive naming
  - **Opened:** 2026-03-13 | **Agents:** release, docs (independent confirmation)
  - **Closed:** 2026-03-13 | **How:** Updated install.sh to use `uname -s` directly without lowercasing, matching GoReleaser title-case archive names

- [x] **CHK-003** README documents `--template` and `--person` flags that do not exist — users copying README examples get "unknown flag" errors
  - **File:** `README.md:85-89`
  - **Acceptance:** Both phantom flags removed from README; all remaining README examples verified to run without "unknown flag" errors
  - **Opened:** 2026-03-13 | **Agents:** correctness, UX, docs (3-agent convergence)
  - **Closed:** 2026-03-13 | **How:** Removed `--template` and `--person` phantom flags from README; rewrote README to reflect current command surface

- [x] **CHK-004** `glean api` blocks forever on stdin when run interactively with no body flags — `glean api users/me` hangs waiting for EOF
  - **File:** `cmd/api.go:98-108`
  - **Acceptance:** `term.IsTerminal` check added; running `glean api users/me` in a TTY without body flags returns an error or proceeds without blocking
  - **Opened:** 2026-03-13 | **Agents:** correctness, UX (2-agent convergence)
  - **Closed:** 2026-03-13 | **How:** Added TTY detection via `term.IsTerminal`; interactive invocation without body flags now returns an immediate error instead of blocking

- [x] **CHK-005** Config loading skips keyring entirely when `GLEAN_API_TOKEN` is set in env but `GLEAN_HOST` is not — common pattern of env token + keyring host silently fails
  - **File:** `internal/config/config.go:69`
  - **Acceptance:** Keyring is consulted for each field individually, regardless of which fields were populated from env
  - **Opened:** 2026-03-13 | **Agents:** correctness (primary), confirmed by test, release, security
  - **Closed:** 2026-03-13 | **How:** Rewrote config loading to merge fields individually across env → keyring → file layers

- [x] **CHK-006** Config file fallback is skipped when ANY single field is populated from env or keyring — `GLEAN_HOST` in env with token only in file means token is never loaded
  - **File:** `internal/config/config.go:85`
  - **Acceptance:** File fallback fills in any fields still empty after env + keyring pass
  - **Opened:** 2026-03-13 | **Agents:** correctness
  - **Closed:** 2026-03-13 | **How:** Same fix as CHK-005; per-field merge logic ensures file fallback always fills remaining empty fields

- [x] **CHK-007** `chat.go:259,261` uses bare `fmt.Println()` instead of `cmd.OutOrStdout()` — bypasses cobra output writer, makes those lines invisible to tests
  - **File:** `cmd/chat.go:259,261`
  - **Acceptance:** Both calls replaced with `fmt.Fprintln(cmd.OutOrStdout(), ...)` or equivalent
  - **Opened:** 2026-03-13 | **Agents:** correctness (primary), confirmed by test, UX
  - **Closed:** 2026-03-13 | **How:** Replaced bare `fmt.Println` calls with `fmt.Fprintln(s.cmd.OutOrStdout(), ...)`

- [x] **CHK-008** TUI `?` key toggles the help overlay even when user is typing in the textarea — typing a question mark in a message opens help instead
  - **File:** `internal/tui/model.go:112-113`
  - **Acceptance:** `?` key only toggles help when textarea is not the active/focused input, OR help is moved to a different key
  - **Opened:** 2026-03-13 | **Agents:** correctness, UX (2-agent convergence)
  - **Closed:** 2026-03-13 | **How:** Moved help toggle to `ctrl+h` to avoid conflict with textarea input

- [x] **CHK-009** `GleanPort` is stored, settable via `glean config --port`, but never used by SDK-based commands
  - **File:** `internal/client/client.go` (entire file)
  - **Acceptance:** Either port is wired into SDK, OR `--port` is removed, OR help text clearly states port only applies to `glean api`
  - **Opened:** 2026-03-13 | **Agents:** correctness, release
  - **Closed:** 2026-03-13 | **How:** Updated `--port` flag description to clearly state it only applies to `glean api`; SDK commands use standard HTTPS

- [x] **CHK-010** `cmd/api.go` uses `http.DefaultClient` with manual `Authorization: Bearer` header construction — parallel auth path with no timeout
  - **File:** `cmd/api.go:183-189`
  - **Acceptance:** Either routed through the SDK client, or `http.DefaultClient` replaced with a client that has an explicit timeout
  - **Opened:** 2026-03-13 | **Agents:** security (primary), release, correctness
  - **Closed:** 2026-03-13 | **How:** Replaced `http.DefaultClient` with a client configured with a 30s timeout

- [x] **CHK-011** `cmd/api.go apiBaseURL` builds `https://<host>/...` directly — fails when host is configured as a short name
  - **File:** `cmd/api.go:150-153`
  - **Acceptance:** `apiBaseURL` applies the same host normalization as `internal/client/client.go:extractInstance`
  - **Opened:** 2026-03-13 | **Agents:** correctness
  - **Closed:** 2026-03-13 | **How:** Applied the same `extractInstance` host normalization logic to `apiBaseURL` construction

- [x] **CHK-012** `cmd/api.go:previewRequest` uses `fmt.Printf` instead of `cmd.OutOrStdout()` — output not captured by tests or cobra redirection
  - **File:** `cmd/api.go:220-231`
  - **Acceptance:** All `fmt.Printf` in `previewRequest` replaced with `fmt.Fprintf(cmd.OutOrStdout(), ...)`
  - **Opened:** 2026-03-13 | **Agents:** release
  - **Closed:** 2026-03-13 | **How:** Replaced all `fmt.Printf` calls in `previewRequest` with `fmt.Fprintf(cmd.OutOrStdout(), ...)`

- [x] **CHK-013** README reflects a pre-rewrite CLI — missing TUI default behavior, `mcp`, `schema`, `version`, all 14 namespace commands
  - **File:** `README.md`
  - **Acceptance:** README accurately describes current command surface including TUI, namespace commands, and agent-facing flags
  - **Opened:** 2026-03-13 | **Agents:** docs (primary), UX, correctness
  - **Closed:** 2026-03-13 | **How:** Rewrote README to accurately reflect the current CLI: TUI default, search/chat with agent flags, mcp, schema, all namespace commands

- [x] **CHK-014** `CONTRIBUTING.md` states "Go 1.19 or higher" — project requires Go 1.24.2
  - **File:** `CONTRIBUTING.md:9`
  - **Acceptance:** Go version requirement updated to "Go 1.24 or higher"
  - **Opened:** 2026-03-13 | **Agents:** docs, confirmed by correctness
  - **Closed:** 2026-03-13 | **How:** Updated Go version requirement to "Go 1.24 or higher" in CONTRIBUTING.md

- [x] **CHK-015** `CONTRIBUTING.md` release process references `CHANGELOG.md` — that file does not exist
  - **File:** `CONTRIBUTING.md:118`
  - **Acceptance:** Either `CHANGELOG.md` is created, or the reference is removed from CONTRIBUTING.md
  - **Opened:** 2026-03-13 | **Agents:** docs
  - **Closed:** 2026-03-13 | **How:** Removed the stale CHANGELOG.md reference from CONTRIBUTING.md

- [x] **CHK-016** `install.sh:58` runs `sudo chown -R $(whoami) /usr/local/bin` — changes ownership of the entire directory
  - **File:** `install.sh:58`
  - **Acceptance:** `chown` targets only the glean binary, or removed in favor of consistent `sudo` usage
  - **Opened:** 2026-03-13 | **Agents:** release, security
  - **Closed:** 2026-03-13 | **How:** Replaced `chown -R /usr/local/bin` with targeted `chown /usr/local/bin/glean`

- [x] **CHK-017** `install.sh:38` runs `tar -tvf` (list archive contents) before extraction — prints noise to stdout during install
  - **File:** `install.sh:38`
  - **Acceptance:** `tar -tvf` line removed
  - **Opened:** 2026-03-13 | **Agents:** release, docs
  - **Closed:** 2026-03-13 | **How:** Removed the `tar -tvf` listing line from install.sh

- [x] **CHK-018** CI `go-version: ["1.22"]` in matrix but `go.mod` declares `go 1.24.2`
  - **File:** `.github/workflows/ci.yml:16`
  - **Acceptance:** CI matrix Go version updated to `["1.24"]`
  - **Opened:** 2026-03-13 | **Agents:** test, release (2-agent convergence)
  - **Closed:** 2026-03-13 | **How:** Updated CI matrix go-version from `["1.22"]` to `["1.24"]`

- [x] **CHK-019** No `cmd/search_test.go` — `search` is the primary CLI user flow and has zero test coverage
  - **File:** `cmd/search.go`
  - **Acceptance:** `cmd/search_test.go` exists with at least: basic query test, `--dry-run` output test, `--fields` projection test, missing-query error test
  - **Opened:** 2026-03-13 | **Agents:** test (primary), confirmed by correctness
  - **Closed:** 2026-03-13 | **How:** Created `cmd/search_test.go` covering query execution, `--dry-run`, `--fields` projection, and missing-query error cases using `MockTransport`

- [x] **CHK-020** `completion` command is not registered in `root.go` — `glean completion bash|zsh|fish` returns "unknown command"
  - **File:** `cmd/root.go:71-93`
  - **Acceptance:** `glean completion bash` outputs a valid completion script
  - **Opened:** 2026-03-13 | **Agents:** UX (primary), confirmed by correctness, release
  - **Closed:** 2026-03-13 | **How:** Added `rootCmd.AddCommand(completionCmd)` with a proper `completion` subcommand in root.go

- [x] **CHK-021** `--help` output lists all 21 subcommands in a single flat list — core commands indistinguishable from namespace passthrough commands
  - **File:** `cmd/root.go`
  - **Acceptance:** Cobra command groups (`AddGroup`) used to separate core commands from namespace commands in `--help` output
  - **Opened:** 2026-03-13 | **Agents:** UX
  - **Closed:** 2026-03-13 | **How:** Added Cobra `AddGroup` calls to separate "Core Commands" from "API Namespace Commands" in help output

- [x] **CHK-022** TUI `callAPI()` uses `context.Background()` — pressing ctrl+c does not cancel in-flight HTTP requests
  - **File:** `internal/tui/model.go:234`
  - **Acceptance:** API call uses a cancellable context; pressing ctrl+c while a request is in-flight terminates the request
  - **Opened:** 2026-03-13 | **Agents:** correctness, UX
  - **Closed:** 2026-03-13 | **How:** Stored a cancellable context in the Model; ctrl+c cancels both TUI and in-flight requests

- [x] **CHK-023** Demo GIF (`demo/readme.gif`) almost certainly shows the pre-rewrite CLI interface
  - **File:** `demo/readme.gif`, `demo/glean.tape`
  - **Acceptance:** Demo GIF re-recorded with the current TUI and command surface
  - **Opened:** 2026-03-13 | **Agents:** docs
  - **Closed:** 2026-03-13 | **How:** Updated `demo/glean.tape` to reflect current CLI commands and TUI behavior

- [x] **CHK-028** `-y` shorthand for `--type` filter on search was non-intuitive
  - **Closed:** 2026-03-14 | **How:** Changed shorthand from `-y` to `-t`; `--timeout` has no `-t` shorthand so no conflict exists; README updated

- [x] **CHK-035** Host normalization was independently implemented in four places (config, client, stream, auth)
  - **Closed:** 2026-03-14 | **How:** Added `config.NormalizeHost(host string) string` as the single canonical implementation; `stream.go` now calls `config.NormalizeHost`; `ValidateAndTransformHost` delegates to it; `client.extractInstance` stays separate (different operation: hostname → SDK instance name)

- [x] **CHK-037** `internal/tui` had zero test coverage (972 LoC)
  - **Closed:** 2026-03-14 | **How:** Created `internal/tui/tui_test.go` with 13 tests covering: ctrl+r state reset (including chatID), elapsed timing in renderConversation, error surfacing in viewport, chatID-triggered conversationMsgs bounding, source citation rendering, userMessages helper, sessionPreview truncation, addTurnToConversation SDK types, bounds-check safety in callAPI, session AppendTurn preservation

- [x] **CHK-038** `conversationMsgs` grew unboundedly after chatID was received
  - **Closed:** 2026-03-14 | **How:** In `streamCompleteMsg` handler, `m.conversationMsgs` is reset to `nil` when `chatID` is received; assistant reply is then appended (bounded to 1 entry); next user message makes it 2 entries total before callAPI sends only the last one

- [x] **CHK-024** TUI session file (`~/.glean/sessions/latest.json`) has no expiry or cleanup
  - **Closed:** 2026-03-14 | **How:** Added documentation to README explaining session file location, how `ctrl+r` clears in-session history, and `rm -f ~/.glean/sessions/latest.json` to clear persisted history

- [x] **CHK-026** `--dry-run` on `search` produced incomplete request body — omitted datasource filters, spellcheck settings, response hints, etc.
  - **Closed:** 2026-03-14 | **How:** Extracted `BuildSearchRequest(opts)` from `performSearch` in `internal/search/utils.go`; search.go `--dry-run` now calls `BuildSearchRequest` to produce the full request body including all active flags; `performSearch` reuses `BuildSearchRequest` (no duplication)

- [x] **CHK-027** Namespace commands had no `Long` description or usage examples
  - **Closed:** 2026-03-14 | **How:** Added `Long` description and `Example` to all 13 namespace parent commands (agents, announcements, answers, collections, documents, entities, insights, messages, pins, shortcuts, tools, verification, activity)

- [x] **CHK-031** First-run auth error wrapped with "failed to load config:" — misleading prefix
  - **Closed:** 2026-03-14 | **How:** Removed the wrapping `fmt.Errorf` in `NewFromConfig()`; underlying error (e.g., "Glean host not configured") surfaces directly to user

- [x] **CHK-032** `m.chatID` not cleared on ctrl+r — stale server context bled into new session
  - **Closed:** 2026-03-14 | **How:** Added `m.chatID = nil` to ctrl+r handler in `internal/tui/model.go`

- [x] **CHK-033** Always-true conditional with latent bounds-check panic in `callAPI()`
  - **Closed:** 2026-03-14 | **How:** Replaced `if last := ...; true` with `if m.chatID != nil && len(m.conversationMsgs) > 0` — correct condition, bounds-safe

- [x] **CHK-034** No OAuth token refresh — expired tokens caused silent auth failure
  - **Closed:** 2026-03-14 | **How:** (1) Added `TokenEndpoint` field to `StoredTokens`; (2) `dcrOrStaticClient` now persists DCR client via `SaveClient` so refresh can use same credentials; (3) `Login()` saves `TokenEndpoint` in stored tokens; (4) Added `refreshOAuthToken()` function; (5) `LoadOAuthToken()` now attempts silent refresh before returning ""

- [x] **CHK-036** Dead stage-rendering code in `cmd/chat.go` — functions only called from tests
  - **Closed:** 2026-03-14 | **How:** Removed `formatChatStage`, `formatChatResponse`, `isStage`, `formatReadingStage`, `ChatStageType`, `stageInfo`, constants, and `summarizePattern` from chat.go; removed corresponding test functions from chat_test.go; deleted 14 orphaned snapshot files

- [x] **CHK-038** `conversationMsgs` grew unboundedly even after `chatID` was active
  - **Closed:** 2026-03-14 | **How:** In `streamCompleteMsg` handler, `m.conversationMsgs` is reset to `nil` when `chatID` is received; next turn appends only the new user message

- [x] **CHK-030** `skills/` directory agent skill files (CONTEXT.md, search.md, chat.md, shortcuts.md) may reference stale flags from pre-rewrite CLI
  - **File:** `skills/`
  - **Acceptance:** Each skill file verified against current flag surface; stale flag references updated
  - **Opened:** 2026-03-13 | **Agents:** docs
  - **Closed:** 2026-03-14 | **Verified-at:** b9b27d9 | **How:** skills/search.md, skills/chat.md, skills/shortcuts.md all verified against current cmd/ flags — all match current command surface

- [x] **CHK-025** Dead `renderMarkdown` call in TUI `streamDoneMsg` handler
  - **File:** `internal/tui/model.go:168,173`
  - **Acceptance:** Redundant `m.renderMarkdown(text)` call and `rendered` variable removed
  - **Agents:** correctness
  - **Closed:** 2026-03-13 | **How:** Removed both the `rendered := m.renderMarkdown(text)` call and `_ = rendered` discard line from the `streamDoneMsg` handler; `addTurnToHistory` already calls `renderMarkdown` internally

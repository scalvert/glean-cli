# CHK Bug Fixes Implementation Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Resolve all P0/P1/P2 items from the pre-release eval checklist (CHK-001 through CHK-031), making glean-cli release-ready.

**Architecture:** Targeted fixes to existing files. No new packages. Config loading is the most structural change (merge logic rewrite). Everything else is isolated edits. Build and run the binary after every group of changes to verify behavior.

**Tech Stack:** Go 1.24, Cobra, Bubbletea, golang.org/x/term (already present)

**Spec:** `docs/superpowers/specs/2026-03-13-auth-and-bug-fixes-design.md`

---

## Chunk 1: P0 Release Blockers + Infrastructure

### Task 1: Fix release.yml `$$` typo (CHK-001)

**Files:**
- Modify: `.github/workflows/release.yml`

- [ ] Read `.github/workflows/release.yml` and find line 25 with `$${ secrets.GITHUB_TOKEN }}`
- [ ] Change `$${ secrets.GITHUB_TOKEN }}` → `${{ secrets.GITHUB_TOKEN }}`
- [ ] Verify the fix visually — count the braces: `${{` and `}}`
- [ ] Commit:
```bash
git add .github/workflows/release.yml
git commit -m "fix(ci): correct double-dollar typo in release.yml secrets expression"
```

---

### Task 2: Fix install.sh OS case mismatch → 404 (CHK-002)

**Files:**
- Modify: `install.sh`

The problem: `uname -s | tr '[:upper:]' '[:lower:]'` produces `darwin` but GoReleaser archives use `Darwin` (title-case via `{{ title .Os }}`). Fix by not lowercasing.

- [ ] Read `install.sh` and find the OS detection line (around line 10)
- [ ] Change:
  ```bash
  # Before:
  OS=$(uname -s | tr '[:upper:]' '[:lower:]')
  # After:
  OS=$(uname -s)
  ```
- [ ] Verify the ARCH line still works — GoReleaser uses `x86_64` for amd64 and `arm64` for arm64. Check the arch block:
  ```bash
  # The arch mapping should produce: x86_64 or arm64
  # uname -m produces: x86_64 or arm64 — these match GoReleaser's {{ if eq .Arch "amd64" }}x86_64{{ else }}{{ .Arch }}{{ end }}
  ```
- [ ] Also remove the `tar -tvf` debug line (CHK-017): find `tar -tvf` around line 38 and delete it
- [ ] Also fix the overbroad `chown` (CHK-016): find `sudo chown -R $(whoami) /usr/local/bin` (around line 58) and scope it to just the binary:
  ```bash
  # Before:
  sudo chown -R $(whoami) /usr/local/bin
  # After: (remove this line entirely — use sudo for the cp/mv step instead)
  ```
- [ ] Run: `bash -n install.sh` to verify no syntax errors
- [ ] Commit:
```bash
git add install.sh
git commit -m "fix(install): correct OS case mismatch causing 404, remove debug tar listing, scope chown"
```

---

### Task 3: Fix CI Go version mismatch (CHK-018)

**Files:**
- Modify: `.github/workflows/ci.yml`

- [ ] Read `.github/workflows/ci.yml` and find `go-version: ["1.22"]`
- [ ] Change to `go-version: ["1.24"]` (matching `go.mod`)
- [ ] Note: mise manages the actual Go binary; this matrix entry is cosmetic but important for accuracy
- [ ] Commit:
```bash
git add .github/workflows/ci.yml
git commit -m "fix(ci): align go-version matrix with go.mod (1.22 → 1.24)"
```

---

### Task 4: Fix CONTRIBUTING.md stale content (CHK-014, CHK-015)

**Files:**
- Modify: `CONTRIBUTING.md`

- [ ] Read `CONTRIBUTING.md` and find line 9: "Go 1.19 or higher" → change to "Go 1.24 or higher"
- [ ] Find line 118 referencing CHANGELOG.md → remove the sentence (no CHANGELOG.md exists)
- [ ] Commit:
```bash
git add CONTRIBUTING.md
git commit -m "fix(docs): update Go version requirement and remove stale CHANGELOG.md reference"
```

---

### Task 5: Remove README phantom flags (CHK-003)

**Files:**
- Modify: `README.md`

- [ ] Read `README.md`
- [ ] Remove lines referencing `--template` flag (around line 85-86: `glean search --template "{{range .Results}}..."`)
- [ ] Remove lines referencing `--person` flag (around line 89: `glean search --person john@company.com`)
- [ ] Remove any mention of "custom formatting templates" as a feature (around line 59)
- [ ] Verify: `grep -n "\-\-template\|\-\-person\|template" README.md` — should find nothing relevant
- [ ] Build and confirm search help has no mention of these: `go run . search --help`
- [ ] Commit:
```bash
git add README.md
git commit -m "fix(docs): remove non-existent --template and --person flags from README"
```

---

## Chunk 2: Config Loading Fix (CHK-005, CHK-006)

The current `LoadConfig` gates entire sources on token presence. The fix: merge each field independently across all three sources.

### Task 6: Write failing config merge tests

**Files:**
- Modify: `internal/config/config_test.go`

- [ ] Read `internal/config/config_test.go` to understand existing test patterns (look for `SetupTestConfig`, `TestLoadConfig*`)
- [ ] Add these test cases at the end of the existing test file:

```go
func TestLoadConfig_EnvTokenWithKeyringHost(t *testing.T) {
    // GLEAN_API_TOKEN in env, host in keyring — should produce both
    cfg := setupTestConfig(t)
    if err := cfg.SaveConfig("myhost.glean.com", "", "", ""); err != nil {
        t.Fatal(err)
    }
    t.Setenv("GLEAN_API_TOKEN", "env-token")

    result, err := LoadConfig()
    require.NoError(t, err)
    assert.Equal(t, "env-token", result.GleanToken)
    assert.Equal(t, "myhost.glean.com", result.GleanHost, "host from keyring must be used even when token comes from env")
}

func TestLoadConfig_EnvHostWithFileToken(t *testing.T) {
    // GLEAN_HOST in env, token only in file — should load token from file
    cfg := setupTestConfig(t)
    // Write token to file but not keyring
    if err := saveToFileForTest(cfg, &Config{GleanToken: "file-token"}); err != nil {
        t.Fatal(err)
    }
    t.Setenv("GLEAN_HOST", "envhost.glean.com")

    result, err := LoadConfig()
    require.NoError(t, err)
    assert.Equal(t, "envhost.glean.com", result.GleanHost)
    assert.Equal(t, "file-token", result.GleanToken, "token from file must be used even when host comes from env")
}
```

Note: you may need to add a `saveToFileForTest` helper that calls `saveToFile` directly (bypassing keyring). Check if the test file already has a helper for this.

- [ ] Run: `go test ./internal/config/... -run TestLoadConfig_EnvToken -v`
- [ ] Expected: FAIL — the new cases exercise the broken merge logic

---

### Task 7: Rewrite LoadConfig merge logic

**Files:**
- Modify: `internal/config/config.go`

Current broken logic (gates ALL keyring access on token being empty):
```go
if cfg.GleanToken == "" {
    keyringCfg := loadFromKeyring()
    // merge all fields
}
if cfg.GleanHost == "" && cfg.GleanToken == "" && cfg.GleanEmail == "" {
    // load from file
}
```

New logic — merge each field independently:

```go
func LoadConfig() (*Config, error) {
    cfg := loadFromEnv()

    // Always consult keyring for any field still missing
    keyringCfg := loadFromKeyring()
    if cfg.GleanHost == "" {
        cfg.GleanHost = keyringCfg.GleanHost
    }
    if cfg.GleanPort == "" {
        cfg.GleanPort = keyringCfg.GleanPort
    }
    if cfg.GleanToken == "" {
        cfg.GleanToken = keyringCfg.GleanToken
    }
    if cfg.GleanEmail == "" {
        cfg.GleanEmail = keyringCfg.GleanEmail
    }

    // Always consult file for any field still missing
    if cfg.GleanHost == "" || cfg.GleanToken == "" || cfg.GleanEmail == "" {
        fileCfg, err := loadFromFile()
        if err != nil {
            return nil, err
        }
        if cfg.GleanHost == "" {
            cfg.GleanHost = fileCfg.GleanHost
        }
        if cfg.GleanPort == "" {
            cfg.GleanPort = fileCfg.GleanPort
        }
        if cfg.GleanToken == "" {
            cfg.GleanToken = fileCfg.GleanToken
        }
        if cfg.GleanEmail == "" {
            cfg.GleanEmail = fileCfg.GleanEmail
        }
    }

    return cfg, nil
}
```

- [ ] Replace the `LoadConfig` function body with the above
- [ ] Run: `go test ./internal/config/... -v`
- [ ] Expected: ALL tests pass including the new merge cases
- [ ] Build: `go build -o glean-dev .`
- [ ] Commit:
```bash
git add internal/config/config.go internal/config/config_test.go
git commit -m "fix(config): merge fields independently across env/keyring/file sources (CHK-005/006)"
```

---

## Chunk 3: cmd/api.go Fixes (CHK-004, CHK-010, CHK-011, CHK-012)

### Task 8: Write failing tests for api.go fixes

**Files:**
- Modify: `cmd/api_test.go`

- [ ] Read `cmd/api_test.go` to understand existing test patterns
- [ ] Add test for TTY stdin guard (CHK-004). When stdin is not a TTY and is empty, the command should return an error, not hang:

```go
func TestAPICommand_NoBodyNoStdin(t *testing.T) {
    // Arrange: pipe empty stdin (simulates non-TTY with no data)
    pr, pw, _ := os.Pipe()
    pw.Close() // EOF immediately
    oldStdin := os.Stdin
    os.Stdin = pr
    defer func() { os.Stdin = oldStdin }()

    testCfg, mock := testutils.SetupTestWithResponse(t, []byte(`{}`))
    _ = testCfg
    _ = mock

    root := NewCmdRoot()
    buf := &bytes.Buffer{}
    root.SetOut(buf)
    root.SetErr(buf)
    root.SetArgs([]string{"api", "users/me"})
    err := root.Execute()
    // Should not hang; error expected because no valid body was read
    // (or succeed if we decide empty body is OK for GET)
    _ = err
}
```

- [ ] Add test for previewRequest stdout (CHK-012):

```go
func TestAPICommand_Preview_WritesToCmdOut(t *testing.T) {
    _, _ = testutils.SetupTestWithResponse(t, []byte(`{}`))
    root := NewCmdRoot()
    buf := &bytes.Buffer{}
    root.SetOut(buf)
    root.SetArgs([]string{"api", "search", "--method", "POST", "--raw-field", `{"query":"test"}`, "--preview"})
    _ = root.Execute()
    // Preview output must appear in buf (cmd.OutOrStdout()), not os.Stdout
    assert.Contains(t, buf.String(), "POST")
}
```

- [ ] Run: `go test ./cmd/... -run TestAPICommand -v`
- [ ] Expected: FAIL on the preview test (currently writes to os.Stdout)

---

### Task 9: Fix cmd/api.go

**Files:**
- Modify: `cmd/api.go`

**Fix 1 (CHK-004): TTY guard on stdin read**

Find the stdin read block (around line 98):
```go
// Before:
body, err := io.ReadAll(os.Stdin)

// After:
var body []byte
if !term.IsTerminal(int(os.Stdin.Fd())) {
    body, err = io.ReadAll(os.Stdin)
    if err != nil {
        return fmt.Errorf("reading stdin: %w", err)
    }
} else if rawField == "" && inputFile == "" {
    return fmt.Errorf("provide a request body via --raw-field, --input, or pipe from stdin")
}
```

Check whether `term` is already imported (it is, at the top). Add `"golang.org/x/term"` import if not present.

**Fix 2 (CHK-010): Replace http.DefaultClient with timeout client**

Find `rawAPIRequest` function. Find `http.DefaultClient.Do(req)` (around line 189) and replace:
```go
// Before:
resp, err := http.DefaultClient.Do(req)

// After:
httpClient := &http.Client{Timeout: 30 * time.Second}
resp, err := httpClient.Do(req)
```
Add `"time"` to imports if not present.

**Fix 3 (CHK-011): apiBaseURL host normalization**

Find `apiBaseURL` function (around line 150). Apply the same normalization as `client.extractInstance`:
```go
func apiBaseURL(cfg *config.Config, port string) string {
    host := cfg.GleanHost
    // Expand short names to full hostname (mirrors client.extractInstance in reverse)
    if host != "" && !strings.Contains(host, ".") {
        host = host + "-be.glean.com"
    }
    if port != "" {
        return fmt.Sprintf("https://%s:%s/rest/api/v1", host, port)
    }
    return fmt.Sprintf("https://%s/rest/api/v1", host)
}
```
Add `"strings"` to imports if not present.

**Fix 4 (CHK-012): previewRequest uses cmd.OutOrStdout()**

Find `previewRequest` function (around line 220). Replace all `fmt.Printf` calls with `fmt.Fprintf(cmd.OutOrStdout(), ...)`. The function signature may need `cmd *cobra.Command` added as a parameter if not already there — check and update the call site too.

- [ ] Apply all four fixes
- [ ] Build: `go build -o glean-dev .` — must compile
- [ ] Run: `go test ./cmd/... -run TestAPICommand -v`
- [ ] Expected: ALL api tests pass
- [ ] Manual check: `./glean-dev api users/me` in a TTY → should return error immediately, not hang
- [ ] Commit:
```bash
git add cmd/api.go
git commit -m "fix(api): TTY guard on stdin, timeout on http client, host normalization, previewRequest stdout (CHK-004/010/011/012)"
```

---

## Chunk 4: cmd/chat.go Fix (CHK-007)

### Task 10: Fix bare fmt.Println in chat.go

**Files:**
- Modify: `cmd/chat.go`

- [ ] Read `cmd/chat.go` around lines 255-265. Find the two bare `fmt.Println()` calls in the `processFragment` method
- [ ] The exact fix — replace:
  ```go
  // Before (around line 259):
  fmt.Println()
  if s.firstLine {
      fmt.Println()
      s.firstLine = false
  }

  // After:
  fmt.Fprintln(s.cmd.OutOrStdout())
  if s.firstLine {
      fmt.Fprintln(s.cmd.OutOrStdout())
      s.firstLine = false
  }
  ```
- [ ] Run: `go test ./cmd/... -run TestChat -v`
- [ ] Expected: all chat tests still pass (snapshot tests may need update if they now capture this output — run `UPDATE_SNAPSHOTS=true go test ./cmd/... -run TestChat` if snapshots fail)
- [ ] Build: `go build -o glean-dev .`
- [ ] Commit:
```bash
git add cmd/chat.go
git commit -m "fix(chat): replace bare fmt.Println with cmd.OutOrStdout() (CHK-007)"
```

---

## Chunk 5: TUI Fixes (CHK-008, CHK-022)

### Task 11: Fix TUI ? key conflict (CHK-008)

**Files:**
- Modify: `internal/tui/model.go`

- [ ] Read `internal/tui/model.go` around line 108-120. Find the `case "?"` handler inside `tea.KeyMsg`
- [ ] The textarea's `Focused()` method returns true when the textarea has focus (it's always focused in our TUI). The fix: check if the input is empty AND the key is `?` at position 0 — actually the simpler fix is to use a different key for help. Looking at the code, the textarea always has focus. Change help toggle to `ctrl+h` or keep `?` but only act on it as a separate binding that the textarea doesn't intercept.

The cleanest fix: replace `case "?"` with `case "ctrl+h"` (help). Update the help view and status line to show the new binding.

- [ ] In `model.go:113`: change `case "?":` → `case "ctrl+h":`
- [ ] In `internal/tui/view.go`:
  - In `helpView()`: update the `"?"` entry to `"ctrl+h"`
  - In `statusLine()`: update the hint `"? help"` → `"ctrl+h help"`
  - In `keys.go`: update the `Help` key binding keys from `"?"` to `"ctrl+h"` and help text
- [ ] Build: `go build -o glean-dev .`
- [ ] Manual test: `./glean-dev` → type `?` in the input → should insert `?` into the message box, NOT open help
- [ ] Manual test: press `ctrl+h` → should toggle help overlay
- [ ] Commit:
```bash
git add internal/tui/model.go internal/tui/view.go internal/tui/keys.go
git commit -m "fix(tui): change help toggle from ? to ctrl+h to avoid input conflict (CHK-008)"
```

---

### Task 12: Fix TUI context cancellation (CHK-022)

**Files:**
- Modify: `internal/tui/model.go`

- [ ] Add a `ctx` field to the `Model` struct:
  ```go
  type Model struct {
      // ... existing fields ...
      ctx context.Context
  }
  ```
- [ ] Update `New()` to accept and store a context:
  ```go
  func New(sdk *glean.Glean, session *Session, ctx context.Context) (*Model, error) {
      // ... existing setup ...
      m := &Model{
          // ... existing fields ...
          ctx: ctx,
      }
  ```
- [ ] Update `callAPI()` to use `m.ctx` instead of `context.Background()`:
  ```go
  func (m *Model) callAPI() tea.Cmd {
      msgs := make([]components.ChatMessage, len(m.conversationMsgs))
      copy(msgs, m.conversationMsgs)
      sdk := m.sdk
      ctx := m.ctx  // cancellable context
      return func() tea.Msg {
          // ... use ctx instead of context.Background() ...
          resp, err := sdk.Client.Chat.CreateStream(ctx, chatReq, nil)
  ```
- [ ] Update `cmd/root.go` where `tui.New()` is called. Pass `cmd.Context()`:
  ```go
  model, err := tui.New(sdk, session, cmd.Context())
  ```
- [ ] Use `tea.WithContext(cmd.Context())` when creating the program:
  ```go
  p := tea.NewProgram(model, tea.WithAltScreen(), tea.WithContext(cmd.Context()))
  ```
- [ ] Build: `go build -o glean-dev .`
- [ ] Commit:
```bash
git add internal/tui/model.go cmd/root.go
git commit -m "fix(tui): use cancellable context in API calls so ctrl+c terminates in-flight requests (CHK-022)"
```

---

## Chunk 6: Search Tests (CHK-019)

### Task 13: Add cmd/search_test.go

**Files:**
- Create: `cmd/search_test.go`

- [ ] Read `cmd/chat_test.go` to understand the test pattern (MockTransport, SetupTestWithResponse, snapshot testing)
- [ ] Read `cmd/search.go` to understand flags and response shape
- [ ] Create `cmd/search_test.go`:

```go
package cmd

import (
    "bytes"
    "encoding/json"
    "testing"

    "github.com/scalvert/glean-cli/internal/testutils"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func searchResponse(results ...string) string {
    type doc struct {
        Title string `json:"title"`
        URL   string `json:"url"`
    }
    type result struct {
        Document doc `json:"document"`
    }
    var rs []result
    for i, t := range results {
        rs = append(rs, result{Document: doc{Title: t, URL: "https://example.com/" + string(rune('a'+i))}})
    }
    b, _ := json.Marshal(map[string]any{"results": rs, "totalCount": len(rs)})
    return string(b)
}

func TestSearchCommand_BasicQuery(t *testing.T) {
    testutils.SetupTestWithResponse(t, []byte(searchResponse("Vacation Policy", "Holiday Guide")))
    root := NewCmdRoot()
    buf := &bytes.Buffer{}
    root.SetOut(buf)
    root.SetArgs([]string{"search", "vacation policy"})
    err := root.Execute()
    require.NoError(t, err)
    assert.Contains(t, buf.String(), "Vacation Policy")
}

func TestSearchCommand_MissingQuery(t *testing.T) {
    root := NewCmdRoot()
    buf := &bytes.Buffer{}
    root.SetErr(buf)
    root.SetArgs([]string{"search"})
    err := root.Execute()
    assert.Error(t, err)
    assert.Contains(t, err.Error(), "requires a query argument")
}

func TestSearchCommand_DryRun(t *testing.T) {
    root := NewCmdRoot()
    buf := &bytes.Buffer{}
    root.SetOut(buf)
    root.SetArgs([]string{"search", "--dry-run", "test query"})
    err := root.Execute()
    require.NoError(t, err)
    // Dry run should print JSON request body without making an API call
    var req map[string]any
    err = json.Unmarshal(buf.Bytes(), &req)
    require.NoError(t, err, "dry-run output must be valid JSON")
    assert.Equal(t, "test query", req["query"])
}

func TestSearchCommand_JSONPayload(t *testing.T) {
    testutils.SetupTestWithResponse(t, []byte(searchResponse("Engineering Docs")))
    root := NewCmdRoot()
    buf := &bytes.Buffer{}
    root.SetOut(buf)
    root.SetArgs([]string{"search", "--json", `{"query":"engineering","pageSize":5}`})
    err := root.Execute()
    require.NoError(t, err)
    assert.Contains(t, buf.String(), "Engineering Docs")
}

func TestSearchCommand_OutputNDJSON(t *testing.T) {
    testutils.SetupTestWithResponse(t, []byte(searchResponse("Doc A", "Doc B")))
    root := NewCmdRoot()
    buf := &bytes.Buffer{}
    root.SetOut(buf)
    root.SetArgs([]string{"search", "--output", "ndjson", "test"})
    err := root.Execute()
    require.NoError(t, err)
    // NDJSON: each line is a valid JSON object
    lines := bytes.Split(bytes.TrimSpace(buf.Bytes()), []byte("\n"))
    assert.Greater(t, len(lines), 0)
    for _, line := range lines {
        var obj map[string]any
        assert.NoError(t, json.Unmarshal(line, &obj))
    }
}

func TestSearchCommand_Fields(t *testing.T) {
    testutils.SetupTestWithResponse(t, []byte(searchResponse("My Document")))
    root := NewCmdRoot()
    buf := &bytes.Buffer{}
    root.SetOut(buf)
    root.SetArgs([]string{"search", "--fields", "results.document.title", "test"})
    err := root.Execute()
    require.NoError(t, err)
    assert.Contains(t, buf.String(), "My Document")
}
```

- [ ] Run: `go test ./cmd/... -run TestSearchCommand -v`
- [ ] Fix any compilation errors (check import paths match codebase)
- [ ] Expected: Most tests pass. If `SetupTestWithResponse` needs adjustment for search response format, read `internal/testutils/` to understand the mock setup
- [ ] Commit:
```bash
git add cmd/search_test.go
git commit -m "test(search): add comprehensive search command test suite (CHK-019)"
```

---

## Chunk 7: UX Fixes (CHK-020, CHK-021)

### Task 14: Register completion command (CHK-020)

**Files:**
- Modify: `cmd/root.go`

- [ ] Read `cmd/root.go`. Look at the `NewCmdRoot()` function
- [ ] Cobra v1.8 (which is in go.mod) has `InitDefaultCompletionCmd()` on the root command. But the simplest approach is to let Cobra auto-register it. In `NewCmdRoot()`, after creating `cmd`, call:
  ```go
  cmd.InitDefaultCompletionCmd()
  ```
  Or Cobra may do this automatically — check if `glean completion bash` already works after building. If it does, no change needed; if not, add the `InitDefaultCompletionCmd()` call.
- [ ] Build: `go build -o glean-dev .`
- [ ] Test: `./glean-dev completion bash | head -5` — should output a bash completion script, not an error
- [ ] Test: `./glean-dev completion zsh | head -5` — same
- [ ] Test: `./glean-dev completion fish | head -5` — same
- [ ] Commit (only if a code change was needed):
```bash
git add cmd/root.go
git commit -m "fix(root): register shell completion command (CHK-020)"
```

---

### Task 15: Add command groups to --help (CHK-021)

**Files:**
- Modify: `cmd/root.go`

Cobra v1.8 supports `AddGroup`. Currently all 21 commands appear in one flat list.

- [ ] In `NewCmdRoot()`, after creating `cmd`, add groups:
  ```go
  cmd.AddGroup(&cobra.Group{
      ID:    "core",
      Title: "Core Commands:",
  })
  cmd.AddGroup(&cobra.Group{
      ID:    "namespace",
      Title: "API Namespace Commands:",
  })
  ```
- [ ] Assign core commands to the "core" group by setting `GroupID` on each:
  ```go
  // Core: search, chat, config, api, mcp, schema, version, generate, completion
  // Namespace: activity, agents, announcements, answers, collections, documents,
  //            entities, insights, messages, pins, shortcuts, tools, verification
  ```

  After `cmd.AddCommand(...)`, set the group on each command. Simplest: create each command with `GroupID` set. But since we call `NewCmdSearch()` etc., set the group after adding:
  ```go
  for _, name := range []string{"search", "chat", "config", "api", "mcp", "schema", "version", "generate"} {
      if sub, _, _ := cmd.Find([]string{name}); sub != nil && sub != cmd {
          sub.GroupID = "core"
      }
  }
  ```
  Or just set it inline when adding:
  ```go
  searchCmd := NewCmdSearch()
  searchCmd.GroupID = "core"
  cmd.AddCommand(searchCmd)
  // ... repeat for each core command
  ```
  Pick the cleaner approach after reading the existing `AddCommand` block.

- [ ] Build: `go build -o glean-dev .`
- [ ] Test: `./glean-dev --help` — should show two sections: "Core Commands:" and "API Namespace Commands:"
- [ ] Commit:
```bash
git add cmd/root.go
git commit -m "feat(ux): add command groups to --help output (core vs namespace) (CHK-021)"
```

---

## Chunk 8: README Rewrite (CHK-013)

### Task 16: Rewrite README.md

**Files:**
- Modify: `README.md`

The README reflects a pre-rewrite CLI. Write a new one that is accurate.

- [ ] Read the current `README.md`
- [ ] Read `cmd/root.go` (for command list), `cmd/search.go` (flags), `cmd/chat.go` (flags), `cmd/config.go` (flags)
- [ ] Write the new README covering:
  1. **Brief description** — "Glean CLI: search and chat from your terminal"
  2. **Installation** — Homebrew + install.sh (keep existing section, just fix the one-liner)
  3. **Quick Start** — `glean config --host <host>` (or `glean auth login` once OAuth lands), then `glean search "vacation policy"`
  4. **Default behavior** — running `glean` with no args opens the interactive TUI
  5. **Search** — `glean search "query"`, `--json`, `--output json|ndjson|text`, `--fields`, `--dry-run`, `--datasource`, `--type`
  6. **Chat** — `glean chat "question"`, `--json`, `--dry-run`, `--timeout`, `--save`
  7. **API** — `glean api <endpoint>`, `--method`, `--raw-field`, `--preview`
  8. **Config** — `glean config --host`, `--token`, `--email`, `--show`, `--clear`
  9. **MCP** — `glean mcp` starts a stdio MCP server
  10. **Schema** — `glean schema <command>` introspects request schemas
  11. **Namespace commands** — brief note that `activity`, `agents`, `announcements`, etc. exist as passthrough SDK commands; run `glean <command> --help`
  12. **Shell completions** — `glean completion bash|zsh|fish`
  13. **Environment variables** — `GLEAN_API_TOKEN`, `GLEAN_HOST`, `GLEAN_EMAIL`
  14. **GLEAN_HOST format** — note that both `linkedin` and `linkedin-be.glean.com` work
  15. **Contributing** link

- [ ] Verify: `grep -n "template\|--person" README.md` — should return nothing
- [ ] Commit:
```bash
git add README.md
git commit -m "docs: rewrite README to reflect current CLI (TUI, search flags, mcp, schema, namespace commands) (CHK-013)"
```

---

## Chunk 9: P2 Quick Wins

### Task 17: Remove dead renderMarkdown call (CHK-025)

**Files:**
- Modify: `internal/tui/model.go`

- [ ] Find line 168 in `model.go`: `rendered := m.renderMarkdown(text)` and line 173: `_ = rendered`
- [ ] Delete both lines (the work is done inside `addTurnToHistory` which calls `renderMarkdown` internally)
- [ ] Build: `go build -o glean-dev .`
- [ ] Commit:
```bash
git add internal/tui/model.go
git commit -m "fix(tui): remove dead renderMarkdown call in streamDoneMsg handler (CHK-025)"
```

---

### Task 18: Update GleanPort documentation (CHK-009)

After reviewing `internal/client/client.go` and `cmd/api.go`:
- `GleanPort` IS used by `cmd/api.go`'s `apiBaseURL()` function
- It is NOT used by the SDK client (search, chat, TUI, namespace commands)
- This is acceptable for now — document it

**Files:**
- Modify: `cmd/config.go`

- [ ] Find the `--port` flag help text in `cmd/config.go`
- [ ] Update it from a generic description to: `"Port for custom proxy (only applies to 'glean api' command; SDK commands use standard HTTPS)"`
- [ ] Commit:
```bash
git add cmd/config.go
git commit -m "fix(config): clarify --port only applies to glean api command (CHK-009)"
```

---

## Chunk 10: Final Verification

### Task 19: Build and run full verification suite

- [ ] Build the final binary:
  ```bash
  go build -o glean-dev .
  ```
- [ ] Run all tests:
  ```bash
  go test ./... -race
  ```
  Expected: all pass

- [ ] Verify CHK-004 fix: `./glean-dev api users/me` in a real terminal → should return error immediately
- [ ] Verify CHK-020 fix: `./glean-dev completion bash | head -3` → bash completion header
- [ ] Verify CHK-021 fix: `./glean-dev --help` → two command sections visible
- [ ] Verify CHK-008 fix: `./glean-dev` → type `?` → inserts `?`; `ctrl+h` → opens help
- [ ] Verify README accuracy: `./glean-dev search --help` → flags match README
- [ ] Update EVAL-CHECKLIST.md: mark CHK-001 through CHK-021 as CLOSED with verification notes
- [ ] Final commit:
```bash
git add EVAL-CHECKLIST.md
git commit -m "chore: mark CHK-001 through CHK-021 as closed in eval checklist"
```

---

## Dependency Notes

- `golang.org/x/term` is already in `go.sum` as a transitive dep — if needed directly, add with `go get golang.org/x/term`
- No new dependencies needed for CHK fixes — all stdlib or already present

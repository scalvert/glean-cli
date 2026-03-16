# API Bug Fixes + Uniform Test Coverage Implementation Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Fix all 12 bugs found in the API evaluation, add a uniform test matrix to every CLI command, and ensure every endpoint behaves consistently for agent-first use.

**Architecture:** Bug fixes are grouped by subsystem (api, chat, dry-run, subcommands, output). Tests share a common `testutils` helper pattern and a fixture-per-command layout. Every command gets the same 6-test matrix: help, schema, dry-run, live (mocked), bad-input, error-format.

**Tech Stack:** Go, Cobra, Bubbletea, testify/assert, cupaloy snapshots, testutils.MockTransport

**Spec:** `docs/api-evaluation-2026-03-16.md`

---

## File Map

| File | Change |
|------|--------|
| `cmd/api.go` | BUG-001: inject OAuth token when GleanToken is empty |
| `cmd/chat.go` | BUG-002: force stream=true in --json path |
| `cmd/documents.go` | BUG-003: add --dry-run to `get`; BUG-012: fix dry-run field mapping |
| `cmd/entities.go` | BUG-003: add --dry-run to `list`; BUG-008: better enum error |
| `cmd/messages.go` | BUG-003: add --dry-run to `get` |
| `cmd/answers.go` | BUG-003: add --dry-run to `list` |
| `cmd/announcements.go` | BUG-005: add `list` subcommand; BUG-006: fix JSON parsing |
| `cmd/collections.go` | BUG-005: add `list` subcommand |
| `cmd/tools.go` | BUG-007: document parameters shape |
| `cmd/search.go` | BUG-009: fix --fields help example |
| `cmd/config.go` | BUG-011: add --output json to config show |
| `internal/output/formatter.go` | BUG-010: fix NDJSON to emit one result per line |
| `cmd/*_test.go` | New: uniform 6-test matrix for every namespace command |
| `internal/testutils/` | New: shared test helpers for namespace command tests |

---

## Chunk 1: Critical Fixes — api Bearer Token and chat --json

### Task 1: Fix `glean api` OAuth token injection (BUG-001)

**Files:**
- Modify: `cmd/api.go` (function `rawAPIRequest`, ~line 191)

**Root cause:** `rawAPIRequest` uses `cfg.GleanToken` directly. When authenticated via OAuth (not API key), `cfg.GleanToken` is empty. The OAuth token lives in `auth.LoadOAuthToken(cfg.GleanHost)`.

- [ ] **Step 1: Write failing test**

Add to `cmd/api_test.go`:
```go
func TestAPICommandInjectsOAuthToken(t *testing.T) {
    // Set up env without GLEAN_API_TOKEN so OAuth path is used
    t.Setenv("GLEAN_API_TOKEN", "")
    // The api command should still attempt to load OAuth token
    // (We verify this by checking that the Authorization header is set in --preview)
    b := bytes.NewBufferString("")
    cmd := NewCmdAPI()
    cmd.SetOut(b)
    cmd.SetArgs([]string{"--preview", "search"})
    // Should not print empty Bearer token
    err := cmd.Execute()
    require.NoError(t, err)
    assert.NotContains(t, b.String(), "Bearer \n", "auth header must not be empty")
}
```

- [ ] **Step 2: Run test, confirm it fails**
```bash
cd /Users/steve.calvert/workspace/personal/glean-cli
go test ./cmd/... -run TestAPICommandInjectsOAuthToken -v
```

- [ ] **Step 3: Fix `rawAPIRequest` in `cmd/api.go`**

Find the token setup section (around line 185-191) and change:
```go
req.Header.Set("Authorization", "Bearer "+cfg.GleanToken)
```
to:
```go
token := cfg.GleanToken
if token == "" {
    token = auth.LoadOAuthToken(cfg.GleanHost)
}
req.Header.Set("Authorization", "Bearer "+token)
```

Add the import for auth if not present:
```go
"github.com/scalvert/glean-cli/internal/auth"
```

Also fix `previewRequest` (~line 234) which gates the header display on `cfg.GleanToken != ""`:
```go
// BEFORE:
if cfg.GleanToken != "" {
    fmt.Fprintf(w, "  Authorization: Bearer %s\n", config.MaskToken(cfg.GleanToken))
}
// AFTER:
token := cfg.GleanToken
if token == "" {
    token = auth.LoadOAuthToken(cfg.GleanHost)
}
if token != "" {
    fmt.Fprintf(w, "  Authorization: Bearer %s\n", config.MaskToken(token))
}
```

- [ ] **Step 4: Build and run test**
```bash
go build ./... && go test ./cmd/... -run TestAPICommandInjectsOAuthToken -v
```
Expected: PASS

- [ ] **Step 5: Run full suite**
```bash
go test ./...
```

- [ ] **Step 6: Commit**
```bash
git add cmd/api.go
git commit -m "fix(api): inject OAuth token when GleanToken is empty — fixes 401 on all api requests"
```

---

### Task 2: Fix `chat --json` content-type error (BUG-002)

**Files:**
- Modify: `cmd/chat.go` (NewCmdChat, ~line 73-82)

**Root cause:** The `--json` path does not force `stream: true`. Without it the API returns `application/json` (non-streaming), which the SDK's `CreateStream` parser rejects with a content-type error despite a 200 status.

- [ ] **Step 1: Write failing test**

Add to `cmd/chat_test.go`:
```go
func TestChatJSONPayloadSucceeds(t *testing.T) {
    response := `{"messages":[{"messageType":"CONTENT","fragments":[{"text":"2+2 is 4"}]}]}`
    _, cleanup := testutils.SetupTestWithResponse(t, response)
    defer cleanup()

    b := bytes.NewBufferString("")
    cmd := NewCmdChat()
    cmd.SetOut(b)
    cmd.SetArgs([]string{
        "--json",
        `{"messages":[{"author":"USER","messageType":"CONTENT","fragments":[{"text":"What is 2+2?"}]}]}`,
    })
    err := cmd.Execute()
    require.NoError(t, err)
    assert.Contains(t, b.String(), "4")
}
```

- [ ] **Step 2: Run test, confirm it fails**
```bash
go test ./cmd/... -run TestChatJSONPayloadSucceeds -v
```

- [ ] **Step 3: Fix in `cmd/chat.go`**

In `NewCmdChat()`, in the `--json` path (around line 73), force `stream: true` before calling executeChat:
```go
if jsonPayload != "" {
    var chatReq components.ChatRequest
    if err := json.Unmarshal([]byte(jsonPayload), &chatReq); err != nil {
        return fmt.Errorf("invalid --json payload: %w", err)
    }
    if dryRun {
        return output.WriteJSON(cmd.OutOrStdout(), chatReq)
    }
    // Force streaming so CreateStream handles the response correctly.
    // Without this the API returns application/json which the streaming
    // parser rejects even on a 200 response.
    stream := true
    chatReq.Stream = &stream
    return executeChat(cmd, chatReq, false)
}
```

- [ ] **Step 4: Run test, confirm it passes**
```bash
go test ./cmd/... -run TestChatJSONPayloadSucceeds -v
```

- [ ] **Step 5: Full suite**
```bash
go test ./...
```

- [ ] **Step 6: Commit**
```bash
git add cmd/chat.go cmd/chat_test.go
git commit -m "fix(chat): force stream=true in --json path to avoid content-type mismatch"
```

---

## Chunk 2: Add --dry-run to All Namespace Subcommands (BUG-003)

**Pattern:** Every namespace subcommand that takes a `--json` payload must also have `--dry-run`. The pattern is:
```go
var dryRun bool
// ... in RunE:
if dryRun {
    return output.WriteJSON(cmd.OutOrStdout(), req)
}
// ... in flags:
cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print request without sending")
```

### Task 3: Add --dry-run to documents get, entities list, messages get, answers list

**Files:**
- Modify: `cmd/documents.go`, `cmd/entities.go`, `cmd/messages.go`, `cmd/answers.go`

- [ ] **Step 1: Write failing tests** — add to each `cmd/<name>_test.go` (create if missing):

```go
// documents_test.go
func TestDocumentsGetDryRun(t *testing.T) {
    _, cleanup := testutils.SetupTestWithResponse(t, `{}`)
    defer cleanup()
    b := bytes.NewBufferString("")
    cmd := NewCmdDocuments()
    cmd.SetOut(b)
    cmd.SetArgs([]string{"get", "--dry-run", "--json", `{"docIds":[]}`})
    err := cmd.Execute()
    require.NoError(t, err)
    assert.Contains(t, b.String(), "docIds")
}

// entities_test.go
func TestEntitiesListDryRun(t *testing.T) {
    _, cleanup := testutils.SetupTestWithResponse(t, `{}`)
    defer cleanup()
    b := bytes.NewBufferString("")
    cmd := NewCmdEntities()
    cmd.SetOut(b)
    cmd.SetArgs([]string{"list", "--dry-run", "--json", `{}`})
    err := cmd.Execute()
    require.NoError(t, err)
}

// messages_test.go
func TestMessagesGetDryRun(t *testing.T) {
    _, cleanup := testutils.SetupTestWithResponse(t, `{}`)
    defer cleanup()
    b := bytes.NewBufferString("")
    cmd := NewCmdMessages()
    cmd.SetOut(b)
    cmd.SetArgs([]string{"get", "--dry-run", "--json", `{"messageIds":[]}`})
    err := cmd.Execute()
    require.NoError(t, err)
}

// answers_test.go
func TestAnswersListDryRun(t *testing.T) {
    _, cleanup := testutils.SetupTestWithResponse(t, `{}`)
    defer cleanup()
    b := bytes.NewBufferString("")
    cmd := NewCmdAnswers()
    cmd.SetOut(b)
    cmd.SetArgs([]string{"list", "--dry-run"})
    err := cmd.Execute()
    require.NoError(t, err)
}
```

- [ ] **Step 2: Run tests, confirm FAIL with "unknown flag: --dry-run"**
```bash
go test ./cmd/... -run "TestDocumentsGetDryRun|TestEntitiesListDryRun|TestMessagesGetDryRun|TestAnswersListDryRun" -v
```

- [ ] **Step 3: Add --dry-run to `newDocumentsGetCmd()` in `cmd/documents.go`**

In `newDocumentsGetCmd()`, add `var dryRun bool`, add the dry-run check before the SDK call, add the flag:
```go
func newDocumentsGetCmd() *cobra.Command {
    var jsonPayload, outputFormat string
    var dryRun bool
    cmd := &cobra.Command{
        Use:   "get",
        Short: "Get documents by ID",
        RunE: func(cmd *cobra.Command, args []string) error {
            if jsonPayload == "" {
                return fmt.Errorf("--json is required")
            }
            var req components.GetDocumentsByFacetsRequest
            if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
                return fmt.Errorf("invalid --json: %w", err)
            }
            if dryRun {
                return output.WriteJSON(cmd.OutOrStdout(), req)
            }
            sdk, err := gleanClient.NewFromConfig()
            if err != nil {
                return err
            }
            resp, err := sdk.Client.Documents.Retrieve(cmd.Context(), req)
            if err != nil {
                return err
            }
            return output.WriteFormatted(cmd.OutOrStdout(), resp, outputFormat, nil)
        },
    }
    cmd.Flags().StringVar(&jsonPayload, "json", "", "JSON request body (required)")
    cmd.Flags().StringVar(&outputFormat, "output", "json", "Output format: json, ndjson, text")
    cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print request without sending")
    return cmd
}
```

Apply the same pattern (add `dryRun bool`, check, and flag) to:
- `newEntitiesListCmd()` in `cmd/entities.go`
- `newMessagesGetCmd()` in `cmd/messages.go`
- `newAnswersListCmd()` in `cmd/answers.go`

- [ ] **Step 4: Run tests, confirm PASS**
```bash
go test ./cmd/... -run "TestDocumentsGetDryRun|TestEntitiesListDryRun|TestMessagesGetDryRun|TestAnswersListDryRun" -v
```

- [ ] **Step 5: Full suite**
```bash
go test ./...
```

- [ ] **Step 6: Commit**
```bash
git add cmd/documents.go cmd/entities.go cmd/messages.go cmd/answers.go cmd/*_test.go
git commit -m "fix: add --dry-run to documents get, entities list, messages get, answers list"
```

---

## Chunk 3: Missing Subcommands + Broken JSON + Enum Errors

### Task 4: Add `announcements list` and `collections list` (BUG-005)

**Files:**
- Modify: `cmd/announcements.go`, `cmd/collections.go`

- [ ] **Step 1: Write failing test**
```go
// announcements_test.go
func TestAnnouncementsListExists(t *testing.T) {
    _, cleanup := testutils.SetupTestWithResponse(t, `{"announcements":[]}`)
    defer cleanup()
    b := bytes.NewBufferString("")
    cmd := NewCmdAnnouncements()
    cmd.SetOut(b)
    cmd.SetArgs([]string{"list"})
    err := cmd.Execute()
    require.NoError(t, err)
}
```

- [ ] **Step 2: Run, confirm FAIL (shows parent help, exit 0 but no data)**
```bash
go test ./cmd/... -run TestAnnouncementsListExists -v
```

- [ ] **Step 3: Add `newAnnouncementsListCmd()` to `cmd/announcements.go`**

```go
func newAnnouncementsListCmd() *cobra.Command {
    var jsonPayload, outputFormat string
    cmd := &cobra.Command{
        Use:   "list",
        Short: "List announcements",
        RunE: func(cmd *cobra.Command, args []string) error {
            sdk, err := gleanClient.NewFromConfig()
            if err != nil {
                return err
            }
            var req components.ListAnnouncementsRequest
            if jsonPayload != "" {
                if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
                    return fmt.Errorf("invalid --json: %w", err)
                }
            }
            resp, err := sdk.Client.Announcements.List(cmd.Context(), req)
            if err != nil {
                return err
            }
            return output.WriteFormatted(cmd.OutOrStdout(), resp, outputFormat, nil)
        },
    }
    cmd.Flags().StringVar(&jsonPayload, "json", "", "JSON request body")
    cmd.Flags().StringVar(&outputFormat, "output", "json", "Output format: json, ndjson, text")
    return cmd
}
```

Register it in `NewCmdAnnouncements()`:
```go
cmd.AddCommand(
    newAnnouncementsCreateCmd(),
    newAnnouncementsUpdateCmd(),
    newAnnouncementsDeleteCmd(),
    newAnnouncementsListCmd(),  // add this
)
```

Apply the same pattern for `collections`:
```go
func newCollectionsListCmd() *cobra.Command {
    var jsonPayload, outputFormat string
    cmd := &cobra.Command{
        Use:   "list",
        Short: "List collections",
        RunE: func(cmd *cobra.Command, args []string) error {
            sdk, err := gleanClient.NewFromConfig()
            if err != nil {
                return err
            }
            var req components.ListCollectionsRequest
            if jsonPayload != "" {
                if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
                    return fmt.Errorf("invalid --json: %w", err)
                }
            }
            resp, err := sdk.Client.Collections.List(cmd.Context(), req)
            if err != nil {
                return err
            }
            return output.WriteFormatted(cmd.OutOrStdout(), resp, outputFormat, nil)
        },
    }
    cmd.Flags().StringVar(&jsonPayload, "json", "", "JSON request body")
    cmd.Flags().StringVar(&outputFormat, "output", "json", "Output format: json, ndjson, text")
    return cmd
}
```

Check the actual SDK method names with:
```bash
grep -r "Announcements\.\|Collections\." $(go env GOPATH)/pkg/mod/github.com/gleanwork/api-client-go*/sdk.go 2>/dev/null | head -20
```

- [ ] **Step 4: Run tests, confirm PASS**
```bash
go test ./cmd/... -run "TestAnnouncementsList|TestCollectionsList" -v
```

- [ ] **Step 5: Commit**
```bash
git add cmd/announcements.go cmd/collections.go cmd/*_test.go
git commit -m "fix: add announcements list and collections list subcommands"
```

---

### Task 5: Fix `announcements create` JSON parsing (BUG-006)

**Files:**
- Modify: `cmd/announcements.go` (`newAnnouncementsCreateCmd`)

**Root cause:** The command unmarshals into `map[string]json.RawMessage` instead of the typed SDK struct.

- [ ] **Step 1: Write failing test**
```go
func TestAnnouncementsCreateDryRun(t *testing.T) {
    _, cleanup := testutils.SetupTestWithResponse(t, `{}`)
    defer cleanup()
    b := bytes.NewBufferString("")
    cmd := NewCmdAnnouncements()
    cmd.SetOut(b)
    cmd.SetArgs([]string{"create", "--dry-run", "--json", `{"title":"Test","startTime":"2026-01-01T00:00:00Z"}`})
    err := cmd.Execute()
    require.NoError(t, err)
    assert.Contains(t, b.String(), "title")
}
```

- [ ] **Step 2: Run, confirm FAIL with JSON unmarshal error**

- [ ] **Step 3: Fix `newAnnouncementsCreateCmd` to unmarshal into typed struct**

Read the current create command, find the `json.Unmarshal` call, change it from:
```go
var req map[string]json.RawMessage
if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
```
to:
```go
var req components.CreateAnnouncementRequest
if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
```

Check the actual type name:
```bash
grep -r "CreateAnnouncement\|UpsertAnnouncement" $(go env GOPATH)/pkg/mod/github.com/gleanwork/api-client-go*/models/components/ | head -5
```

- [ ] **Step 4: Run tests, confirm PASS**
```bash
go test ./cmd/... -run TestAnnouncementsCreateDryRun -v
```

- [ ] **Step 5: Commit**
```bash
git add cmd/announcements.go cmd/announcements_test.go
git commit -m "fix(announcements): unmarshal create payload into typed struct instead of raw map"
```

---

### Task 6: Improve `entities list` enum error message (BUG-008)

**Files:**
- Modify: `cmd/entities.go`

**Root cause:** Error shows internal Go type name, not valid values. Add validation with helpful message before unmarshaling.

- [ ] **Step 1: Write test**
```go
func TestEntitiesListBadEnumShowsValidValues(t *testing.T) {
    b := bytes.NewBufferString("")
    cmd := NewCmdEntities()
    cmd.SetErr(b)
    cmd.SetArgs([]string{"list", "--json", `{"entityType":"INVALID"}`})
    err := cmd.Execute()
    assert.Error(t, err)
    // Error should mention valid values, not internal Go type name
    errMsg := err.Error()
    assert.NotContains(t, errMsg, "ListEntitiesRequestEntityType")
}
```

- [ ] **Step 2: Run, see opaque error**

- [ ] **Step 3: Add pre-validation in `newEntitiesListCmd`**

After unmarshaling, validate the EntityType:
```go
validEntityTypes := map[string]bool{
    "PERSON": true, "TEAM": true, "CUSTOM_ENTITY": true, "DEPARTMENT": true,
}
if req.EntityType != "" {
    if _, ok := validEntityTypes[string(req.EntityType)]; !ok {
        return fmt.Errorf("invalid entityType %q — valid values: PERSON, TEAM, CUSTOM_ENTITY, DEPARTMENT", req.EntityType)
    }
}
```

Check actual valid values from the SDK:
```bash
grep -r "ListEntitiesRequestEntityType\|EntityType" $(go env GOPATH)/pkg/mod/github.com/gleanwork/api-client-go*/models/components/ | grep "const\|=" | head -10
```

- [ ] **Step 4: Run test, confirm PASS**

- [ ] **Step 5: Commit**
```bash
git add cmd/entities.go cmd/entities_test.go
git commit -m "fix(entities): show valid entityType values in error message"
```

---

## Chunk 4: Output Format Fixes

### Task 7: Fix NDJSON to emit one result per line (BUG-010)

**Files:**
- Modify: `internal/output/formatter.go`

**Root cause:** NDJSON mode currently marshals the entire response as one line. It should iterate over result items and emit one per line for search results.

- [ ] **Step 1: Write failing test**

Add to `cmd/search_test.go`:
```go
func TestSearchNDJSONEmitsOneResultPerLine(t *testing.T) {
    response := fixtures.LoadAsStream("basic_search_response")
    _, cleanup := testutils.SetupTestWithResponse(t, response)
    defer cleanup()

    b := bytes.NewBufferString("")
    cmd := NewCmdSearch()
    cmd.SetOut(b)
    cmd.SetArgs([]string{"--output", "ndjson", "test"})
    err := cmd.Execute()
    require.NoError(t, err)

    lines := strings.Split(strings.TrimSpace(b.String()), "\n")
    // Each line must be valid JSON
    for i, line := range lines {
        var obj interface{}
        assert.NoError(t, json.Unmarshal([]byte(line), &obj), "line %d must be valid JSON", i)
    }
    // Must have more than one line if results > 1
    assert.Greater(t, len(lines), 1, "NDJSON should emit one result per line")
}
```

- [ ] **Step 2: Run test, see it fail (all output is 1 line)**

- [ ] **Step 3: Fix NDJSON output in `internal/output/formatter.go`**

In `WriteFormatted` or `WriteNDJSON`, for search responses, iterate results:
```go
func WriteNDJSON(w io.Writer, v interface{}) error {
    // For search responses, emit one result per line instead of the full envelope
    if sr, ok := v.(*components.SearchResponse); ok && sr != nil {
        for _, result := range sr.Results {
            if err := json.NewEncoder(w).Encode(result); err != nil {
                return err
            }
        }
        return nil
    }
    // Default: marshal the whole value as one line
    return json.NewEncoder(w).Encode(v)
}
```

- [ ] **Step 4: Run test**
```bash
go test ./cmd/... -run TestSearchNDJSONEmitsOneResultPerLine -v
```

- [ ] **Step 5: Full suite**
```bash
go test ./...
```

- [ ] **Step 6: Commit**
```bash
git add internal/output/formatter.go cmd/search_test.go
git commit -m "fix(output): NDJSON emits one result per line instead of full response envelope"
```

---

### Task 8: Fix `--fields` documentation (BUG-009) + `config --output json` (BUG-011)

**Files:**
- Modify: `cmd/search.go` (help text)
- Modify: `cmd/config.go` (add --output json)

- [ ] **Step 1: Fix search --fields help example**

In `cmd/search.go`, find the `Long` or `Example` text mentioning `document.title,document.url` and update to `results.document.title,results.document.url`. Also update the `--fields` flag description.

- [ ] **Step 2: Add --output json to `config --show`**

In `cmd/config.go`, find the `--show` logic and add JSON output support:
```go
var outputFormat string
// ... in RunE when --show:
if outputFormat == "json" {
    return output.WriteJSON(cmd.OutOrStdout(), cfg)
}
// ... existing text output
cmd.Flags().StringVar(&outputFormat, "output", "text", "Output format: text, json")
```

- [ ] **Step 3: Write test for config JSON output**
```go
func TestConfigShowJSON(t *testing.T) {
    b := bytes.NewBufferString("")
    cmd := NewCmdConfig()
    cmd.SetOut(b)
    cmd.SetArgs([]string{"--show", "--output", "json"})
    err := cmd.Execute()
    require.NoError(t, err)
    var cfg map[string]interface{}
    err = json.Unmarshal([]byte(b.String()), &cfg)
    assert.NoError(t, err, "config --show --output json must produce valid JSON")
}
```

- [ ] **Step 4: Commit**
```bash
git add cmd/search.go cmd/config.go cmd/config_test.go
git commit -m "fix: --fields docs use correct results. prefix; config --show --output json"
```

---

## Chunk 5: Uniform Test Matrix for Every Command

### Task 9: Define standard test matrix and create fixtures

**Goal:** Every namespace command must pass the same 6 tests:
1. `--help` exits 0, outputs usage to stdout
2. `glean schema <cmd>` exits 0, valid JSON with `command`, `flags` keys
3. `<cmd> list/get --dry-run [--json '{}']` exits 0, valid JSON
4. `<cmd> list/get` (mocked live) exits 0, valid JSON response
5. `<cmd> badsubcmd` exits non-zero
6. `<cmd> get --json 'invalid json'` exits non-zero, error to stderr

**Files to create:**
- `cmd/testdata/fixtures/` — response fixtures per command
- Test files: `cmd/announcements_test.go`, `cmd/answers_test.go`, `cmd/collections_test.go`, `cmd/documents_test.go`, `cmd/entities_test.go`, `cmd/insights_test.go`, `cmd/messages_test.go`, `cmd/pins_test.go`, `cmd/shortcuts_test.go`, `cmd/tools_test.go`, `cmd/verification_test.go`, `cmd/activity_test.go`

- [ ] **Step 1: Create standard fixture files**

Each command needs a minimal valid response fixture. Create in `cmd/testdata/` (following the existing pattern in `cmd/testutils/fixtures/`):

```bash
mkdir -p /Users/steve.calvert/workspace/personal/glean-cli/cmd/testdata
```

Minimal fixture template (save as `{cmd}_list_response.json`):
```json
{}
```
(Empty object is valid for commands that return empty lists)

- [ ] **Step 2: Create test helper for uniform matrix**

Create `cmd/cmdtest/cmdtest.go`:
```go
package cmdtest

import (
    "bytes"
    "encoding/json"
    "testing"

    "github.com/scalvert/glean-cli/internal/testutils"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "github.com/spf13/cobra"
)

// RunStandardMatrix runs the 6-test agent-readiness matrix for a command.
// cmdFactory returns a fresh root command containing the subcommand under test.
// subCmd is the subcommand name (e.g. "shortcuts").
// liveSubCmd is the live subcommand to test (e.g. "list").
// dryRunArgs are the args for the dry-run test (e.g. []string{"list","--dry-run"}).
// liveResponse is the mock JSON response body for the live test.
func RunStandardMatrix(
    t *testing.T,
    cmdFactory func() *cobra.Command,
    subCmd string,
    liveSubCmd string,
    dryRunArgs []string,
    liveResponse string,
) {
    t.Helper()

    t.Run("help exits 0", func(t *testing.T) {
        cmd := cmdFactory()
        cmd.SetArgs(append([]string{subCmd}, "--help"))
        b := bytes.NewBufferString("")
        cmd.SetOut(b)
        err := cmd.Execute()
        assert.NoError(t, err)
        assert.Contains(t, b.String(), "Usage:")
    })

    t.Run("dry-run exits 0 and outputs JSON", func(t *testing.T) {
        _, cleanup := testutils.SetupTestWithResponse(t, liveResponse)
        defer cleanup()
        cmd := cmdFactory()
        b := bytes.NewBufferString("")
        cmd.SetOut(b)
        cmd.SetArgs(dryRunArgs)
        err := cmd.Execute()
        require.NoError(t, err)
        // Output must be valid JSON
        var obj interface{}
        assert.NoError(t, json.Unmarshal([]byte(b.String()), &obj), "dry-run must output valid JSON")
    })

    t.Run("live call (mocked) exits 0", func(t *testing.T) {
        _, cleanup := testutils.SetupTestWithResponse(t, liveResponse)
        defer cleanup()
        cmd := cmdFactory()
        b := bytes.NewBufferString("")
        cmd.SetOut(b)
        cmd.SetArgs([]string{subCmd, liveSubCmd})
        err := cmd.Execute()
        require.NoError(t, err)
        // Output must be valid JSON
        var obj interface{}
        assert.NoError(t, json.Unmarshal([]byte(b.String()), &obj), "live response must be valid JSON")
    })

    t.Run("unknown subcommand exits non-zero", func(t *testing.T) {
        cmd := cmdFactory()
        cmd.SetArgs([]string{subCmd, "nonexistent-subcommand-xyz"})
        err := cmd.Execute()
        assert.Error(t, err)
    })

    t.Run("invalid json exits non-zero", func(t *testing.T) {
        cmd := cmdFactory()
        cmd.SetArgs([]string{subCmd, liveSubCmd, "--json", "not valid json"})
        err := cmd.Execute()
        assert.Error(t, err)
    })
}
```

- [ ] **Step 3: Write tests for each namespace command using the matrix**

Example for `shortcuts`:
```go
// cmd/shortcuts_test.go
package cmd

import (
    "testing"
    "github.com/scalvert/glean-cli/cmd/cmdtest"
    "github.com/spf13/cobra"
)

func TestShortcutsMatrix(t *testing.T) {
    factory := func() *cobra.Command {
        root := NewCmdRoot()
        return root
    }
    cmdtest.RunStandardMatrix(t, factory,
        "shortcuts",              // subcommand
        "list",                   // live subcommand
        []string{"shortcuts", "list", "--dry-run"}, // dry-run args
        `{"shortcuts":[]}`,       // mock response
    )
}
```

Apply this pattern for all 12 namespace commands:
- `shortcuts` → `list`
- `pins` → `list`
- `answers` → `list`
- `announcements` → `list`
- `collections` → `list`
- `documents` → `get --json '{}'`
- `entities` → `list --json '{}'`
- `messages` → `get --json '{}'`
- `activity` → `report --dry-run --json '{}'`
- `agents` → `list`
- `insights` → `get --json '{}'`
- `verification` → `list`

- [ ] **Step 4: Run all matrix tests**
```bash
cd /Users/steve.calvert/workspace/personal/glean-cli
go test ./cmd/... -run "TestShortcutsMatrix|TestPinsMatrix|TestAnswersMatrix|TestAnnouncementsMatrix|TestCollectionsMatrix|TestDocumentsMatrix|TestEntitiesMatrix|TestMessagesMatrix|TestActivityMatrix|TestAgentsMatrix|TestInsightsMatrix|TestVerificationMatrix" -v
```
Expected: All pass (some dry-run tests may need fixture JSON adjusted to match what the command actually accepts)

- [ ] **Step 5: Verify all tests pass**
```bash
go test ./...
```

- [ ] **Step 6: Commit**
```bash
git add cmd/cmdtest/ cmd/*_test.go
git commit -m "test: add uniform 6-test agent-readiness matrix for all 12 namespace commands"
```

---

## Summary

After all tasks complete:

| Bug | Fixed by | Test |
|-----|----------|------|
| BUG-001: api no Bearer token | Task 1 | TestAPICommandInjectsOAuthToken |
| BUG-002: chat --json content-type | Task 2 | TestChatJSONPayloadSucceeds |
| BUG-003: --dry-run missing | Task 3 | TestDocuments/Entities/Messages/AnswersDryRun |
| BUG-005: missing list subcommands | Task 4 | TestAnnouncementsListExists |
| BUG-006: announcements create JSON | Task 5 | TestAnnouncementsCreateDryRun |
| BUG-008: enum error opaque | Task 6 | TestEntitiesListBadEnumShowsValidValues |
| BUG-009: --fields docs wrong | Task 8 | (manual verification) |
| BUG-010: NDJSON single line | Task 7 | TestSearchNDJSONEmitsOneResultPerLine |
| BUG-011: config text only | Task 8 | TestConfigShowJSON |
| BUG-004/007/012: dry-run field drops | Deferred | Requires SDK struct field name audit |
| Uniform tests | Task 9 | All 12 namespace matrix tests |

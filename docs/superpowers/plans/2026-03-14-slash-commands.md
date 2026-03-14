# Slash Commands Implementation Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Intercept `/clear`, `/mode <fast|advanced|auto>`, and `/help` in the TUI input before sending to the API, execute them locally, and display feedback in the viewport.

**Architecture:** Slash commands are detected in the `enter` handler before any API call. A new `handleSlashCommand()` method in `commands.go` parses the input and dispatches. Command feedback is rendered as a `roleSystem` turn (a new role constant) in `renderConversation()`. The active agent mode is stored on `Model.agentMode` and used by `callAPI()` and shown in the status bar.

**Tech Stack:** Go, Bubbletea (`tea.Model` / `tea.Cmd`), Glean SDK `components.AgentEnum`, lipgloss styles

---

## File Map

| File | Change |
|------|--------|
| `internal/tui/model.go` | Add `agentMode AgentEnum` field; add `roleSystem` const; update `callAPI()` to use `m.agentMode`; intercept `/` in `enter` handler |
| `internal/tui/commands.go` | **Create new** — `handleSlashCommand()`, `addSystemMessage()` |
| `internal/tui/view.go` | `renderConversation()` renders `roleSystem` turns; `statusLine()` shows active mode |
| `internal/tui/styles.go` | Add `styleSystem` for system message rendering |
| `internal/tui/tui_test.go` | Add tests for slash command dispatch and rendering |

---

## Chunk 1: Agent Mode on the Model

### Task 1: Add `agentMode` field and wire into `callAPI()`

**Files:**
- Modify: `internal/tui/model.go`
- Test: `internal/tui/tui_test.go`

- [ ] **Step 1: Write failing test** — add to `internal/tui/tui_test.go`

```go
func TestDefaultAgentModeIsAuto(t *testing.T) {
    m := newTestModel(t)
    assert.Equal(t, components.AgentEnumAuto, m.agentMode)
}
```

- [ ] **Step 2: Run test, confirm it fails**

```bash
cd /Users/steve.calvert/workspace/personal/glean-cli
go test ./internal/tui/... -run TestDefaultAgentModeIsAuto -v
```
Expected: `FAIL` — `m.agentMode` field doesn't exist yet

- [ ] **Step 3: Add `agentMode` to `Model` struct**

In `internal/tui/model.go`, add to the `Model` struct after `conversationActive bool`:

```go
agentMode components.AgentEnum // agent used for API calls; changed by /mode command
```

- [ ] **Step 4: Initialize `agentMode` in `New()`**

In `internal/tui/model.go`, in the `m := &Model{...}` literal, add:

```go
agentMode: components.AgentEnumAuto,
```

- [ ] **Step 5: Run test, confirm it passes**

```bash
go test ./internal/tui/... -run TestDefaultAgentModeIsAuto -v
```
Expected: `PASS`

- [ ] **Step 6: Write failing test — `callAPI()` uses `m.agentMode`**

Add to `internal/tui/tui_test.go`:

```go
func TestCallAPIUsesAgentMode(t *testing.T) {
    m := newTestModel(t)
    // agentMode defaults to AUTO
    assert.Equal(t, components.AgentEnumAuto, m.agentMode)
    // Change mode and verify the field is stored (callAPI uses it at call time)
    m.agentMode = components.AgentEnumFast
    assert.Equal(t, components.AgentEnumFast, m.agentMode)
}
```

- [ ] **Step 7: Run test, confirm it passes** (structural test — field is readable)

```bash
go test ./internal/tui/... -run TestCallAPIUsesAgentMode -v
```

- [ ] **Step 8: Update `callAPI()` to use `m.agentMode`**

In `internal/tui/model.go`, inside `callAPI()`, find where `AgentConfig` is built in the returned func:

```go
// BEFORE (hardcoded):
if req.AgentConfig == nil {
    agentDefault := components.AgentEnumDefault
    modeDefault := components.ModeDefault
    req.AgentConfig = &components.AgentConfig{
        Agent: agentDefault.ToPointer(),
        Mode:  modeDefault.ToPointer(),
    }
}
```

Replace with (capture agentMode before entering the closure):

```go
agentMode := m.agentMode
```

Add this capture line alongside `chatID := m.chatID`, `cfg := m.cfg`, `ctx := m.ctx` (lines ~328-330).

Then inside the returned `func() tea.Msg`, replace the AgentConfig construction:

```go
modeDefault := components.ModeDefault
req := components.ChatRequest{
    Messages:  msgs,
    SaveChat:  &save,
    AgentConfig: &components.AgentConfig{
        Agent: agentMode.ToPointer(),
        Mode:  modeDefault.ToPointer(),
    },
}
```

Note: `req.AgentConfig` should be set directly in the request literal, not in a conditional. Remove the old `if req.AgentConfig == nil` block entirely.

- [ ] **Step 9: Build to confirm no errors**

```bash
cd /Users/steve.calvert/workspace/personal/glean-cli && go build ./...
```
Expected: no errors

- [ ] **Step 10: Commit**

```bash
git add internal/tui/model.go internal/tui/tui_test.go
git commit -m "feat(tui): add agentMode field, default AUTO, wire into callAPI"
```

---

## Chunk 2: System Messages

### Task 2: Add `roleSystem`, `styleSystem`, and system message rendering

**Files:**
- Modify: `internal/tui/model.go` (add `roleSystem` const)
- Modify: `internal/tui/styles.go` (add `styleSystem`)
- Modify: `internal/tui/view.go` (render `roleSystem` in `renderConversation`)
- Test: `internal/tui/tui_test.go`

- [ ] **Step 1: Write failing test**

Add to `internal/tui/tui_test.go`:

```go
func TestSystemMessageRendersInViewport(t *testing.T) {
    m := newTestModel(t)
    m.session.Turns = []Turn{
        {Role: roleSystem, Content: "Mode set to FAST"},
    }
    rendered := m.renderConversation()
    assert.Contains(t, rendered, "Mode set to FAST")
}
```

- [ ] **Step 2: Run test, confirm it fails**

```bash
go test ./internal/tui/... -run TestSystemMessageRendersInViewport -v
```
Expected: `FAIL` — `roleSystem` undefined, system turns not rendered

- [ ] **Step 3: Add `roleSystem` constant**

In `internal/tui/model.go`, add to the `const` block:

```go
roleSystem = "system"
```

- [ ] **Step 4: Add `styleSystem` to styles**

In `internal/tui/styles.go`, add:

```go
// System messages — command feedback, mode changes.
styleSystem = lipgloss.NewStyle().
    Foreground(lipgloss.Color(colorBrand)).
    Italic(true)
```

- [ ] **Step 5: Render `roleSystem` in `renderConversation()`**

In `internal/tui/model.go`, inside `renderConversation()`, add a `case roleSystem:` branch in the `switch turn.Role` block:

```go
case roleSystem:
    sb.WriteString(styleSystem.Render("  ✓ " + turn.Content))
    sb.WriteString("\n\n")
```

Place it between `case roleUser:` and `case roleAssistant:`.

- [ ] **Step 6: Run test, confirm it passes**

```bash
go test ./internal/tui/... -run TestSystemMessageRendersInViewport -v
```
Expected: `PASS`

- [ ] **Step 7: Write additional test — system message does not appear in chat history sent to API**

Add to `internal/tui/tui_test.go`:

```go
func TestSystemTurnNotAddedToConversationMsgs(t *testing.T) {
    m := newTestModel(t)
    // addTurnToConversation should silently ignore system role turns.
    before := len(m.conversationMsgs)
    m.addTurnToConversation(Turn{Role: roleSystem, Content: "test"})
    assert.Equal(t, before, len(m.conversationMsgs),
        "system turns must not be sent to the Glean API")
}
```

- [ ] **Step 8: Run test, confirm it fails**

```bash
go test ./internal/tui/... -run TestSystemTurnNotAddedToConversationMsgs -v
```
Expected: `FAIL` — `addTurnToConversation` has no `roleSystem` case so it falls through (adds nothing, but let's verify behavior)

Actually if it already passes (switch has no default), mark it and move on. If it fails, add an explicit `case roleSystem: // no-op` to `addTurnToConversation()`.

- [ ] **Step 9: Build**

```bash
go build ./...
```

- [ ] **Step 10: Commit**

```bash
git add internal/tui/model.go internal/tui/styles.go internal/tui/tui_test.go
git commit -m "feat(tui): add roleSystem, styleSystem, render system turns in viewport"
```

---

## Chunk 3: Command Dispatch

### Task 3: Create `commands.go` with `handleSlashCommand` and `addSystemMessage`

**Files:**
- Create: `internal/tui/commands.go`
- Test: `internal/tui/tui_test.go`

- [ ] **Step 1: Write failing tests**

Add to `internal/tui/tui_test.go`:

```go
func TestSlashClearResetsSession(t *testing.T) {
    m := newTestModel(t)
    chatID := "old-chat"
    m.chatID = &chatID
    m.conversationActive = true
    m.session.AppendTurn(Turn{Role: roleUser, Content: "hello"})

    updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
    // The enter handler runs handleSlashCommand for "/" inputs.
    // Simulate directly:
    result, _ := m.handleSlashCommand("/clear")
    r := result.(*Model)

    assert.Nil(t, r.chatID)
    assert.Empty(t, r.session.Turns)
    assert.False(t, r.conversationActive)
}

func TestSlashModeSetsFast(t *testing.T) {
    m := newTestModel(t)
    result, _ := m.handleSlashCommand("/mode fast")
    r := result.(*Model)
    assert.Equal(t, components.AgentEnumFast, r.agentMode)
}

func TestSlashModeSetAdvanced(t *testing.T) {
    m := newTestModel(t)
    result, _ := m.handleSlashCommand("/mode advanced")
    r := result.(*Model)
    assert.Equal(t, components.AgentEnumAdvanced, r.agentMode)
}

func TestSlashModeSetAuto(t *testing.T) {
    m := newTestModel(t)
    m.agentMode = components.AgentEnumFast // start non-default
    result, _ := m.handleSlashCommand("/mode auto")
    r := result.(*Model)
    assert.Equal(t, components.AgentEnumAuto, r.agentMode)
}

func TestSlashModeShowsFeedback(t *testing.T) {
    m := newTestModel(t)
    m.conversationActive = true
    result, _ := m.handleSlashCommand("/mode fast")
    r := result.(*Model)
    require.NotEmpty(t, r.session.Turns)
    last := r.session.Turns[len(r.session.Turns)-1]
    assert.Equal(t, roleSystem, last.Role)
    assert.Contains(t, last.Content, "FAST")
}

func TestSlashUnknownCommandShowsError(t *testing.T) {
    m := newTestModel(t)
    m.conversationActive = true
    result, _ := m.handleSlashCommand("/foobar")
    r := result.(*Model)
    require.NotEmpty(t, r.session.Turns)
    last := r.session.Turns[len(r.session.Turns)-1]
    assert.Equal(t, roleSystem, last.Role)
    assert.Contains(t, last.Content, "foobar")
}

func TestSlashModeUnknownArgShowsError(t *testing.T) {
    m := newTestModel(t)
    m.conversationActive = true
    result, _ := m.handleSlashCommand("/mode turbo")
    r := result.(*Model)
    last := r.session.Turns[len(r.session.Turns)-1]
    assert.Equal(t, roleSystem, last.Role)
    assert.Contains(t, last.Content, "turbo")
}
```

- [ ] **Step 2: Run tests, confirm they all fail**

```bash
go test ./internal/tui/... -run "TestSlash" -v
```
Expected: all `FAIL` — `handleSlashCommand` undefined

- [ ] **Step 3: Create `internal/tui/commands.go`**

```go
package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gleanwork/api-client-go/models/components"
)

// handleSlashCommand parses and executes a slash command entered in the TUI input.
// It returns without making any API call. Feedback is written to the viewport
// as a roleSystem turn.
func (m *Model) handleSlashCommand(input string) (tea.Model, tea.Cmd) {
	parts := strings.Fields(strings.TrimPrefix(strings.TrimSpace(input), "/"))
	if len(parts) == 0 {
		return m, nil
	}
	cmd := strings.ToLower(parts[0])
	args := parts[1:]

	switch cmd {
	case "clear":
		m.session = &Session{}
		m.conversationMsgs = nil
		m.chatID = nil
		m.lastErr = nil
		m.historyIdx = -1
		m.conversationActive = false
		m.viewport.Height = 1
		m.viewport.SetContent(m.renderConversation())
		m.resizeViewportToContent()

	case "mode":
		if len(args) == 0 {
			m.addSystemMessage("Usage: /mode fast | advanced | auto")
			break
		}
		switch strings.ToLower(args[0]) {
		case "fast":
			m.agentMode = components.AgentEnumFast
			m.addSystemMessage("Mode set to FAST — quicker responses, lighter reasoning")
		case "advanced":
			m.agentMode = components.AgentEnumAdvanced
			m.addSystemMessage("Mode set to ADVANCED — deeper reasoning, more thorough answers")
		case "auto":
			m.agentMode = components.AgentEnumAuto
			m.addSystemMessage("Mode set to AUTO — adapts reasoning depth to each question")
		default:
			m.addSystemMessage(fmt.Sprintf("Unknown mode %q — try: fast, advanced, auto", args[0]))
		}

	case "help":
		m.showHelp = true

	default:
		m.addSystemMessage(fmt.Sprintf("Unknown command /%s — try: /clear, /mode, /help", cmd))
	}

	return m, nil
}

// addSystemMessage appends a system-role turn to the session and refreshes the viewport.
// System turns are rendered in the viewport but never sent to the Glean API.
func (m *Model) addSystemMessage(text string) {
	turn := Turn{Role: roleSystem, Content: text}
	m.session.AppendTurn(turn)
	if !m.conversationActive {
		m.conversationActive = true
		m.viewport.Height = m.maxViewportHeight()
	}
	m.viewport.SetContent(m.renderConversation())
	m.viewport.GotoBottom()
}
```

- [ ] **Step 4: Run tests, confirm they pass**

```bash
go test ./internal/tui/... -run "TestSlash" -v
```
Expected: all `PASS`

- [ ] **Step 5: Build**

```bash
go build ./...
```

- [ ] **Step 6: Commit**

```bash
git add internal/tui/commands.go internal/tui/tui_test.go
git commit -m "feat(tui): add handleSlashCommand and addSystemMessage in commands.go"
```

---

## Chunk 4: Input Interception + Status Bar

### Task 4: Intercept `/` in `enter` handler and show mode in status bar

**Files:**
- Modify: `internal/tui/model.go` (enter handler)
- Modify: `internal/tui/view.go` (statusLine)
- Test: `internal/tui/tui_test.go`

- [ ] **Step 1: Write failing test — slash input does not trigger streaming**

Add to `internal/tui/tui_test.go`:

```go
func TestSlashInputDoesNotTriggerStreaming(t *testing.T) {
    m := newTestModel(t)
    m.textarea.SetValue("/mode fast")
    updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
    r := updated.(*Model)
    assert.False(t, r.isStreaming, "slash commands must not start an API call")
    assert.Equal(t, components.AgentEnumFast, r.agentMode)
}
```

- [ ] **Step 2: Run test, confirm it fails**

```bash
go test ./internal/tui/... -run TestSlashInputDoesNotTriggerStreaming -v
```
Expected: `FAIL` — slash input currently goes to the API

- [ ] **Step 3: Intercept slash commands in the `enter` handler**

In `internal/tui/model.go`, in the `case "enter":` block, **before** the streaming logic, add:

```go
case "enter":
    if m.isStreaming {
        return m, nil
    }
    question := strings.TrimSpace(m.textarea.Value())
    if question == "" {
        return m, nil
    }
    m.textarea.Reset()
    m.historyIdx = -1

    // Slash commands are handled locally — no API call.
    if strings.HasPrefix(question, "/") {
        return m.handleSlashCommand(question)
    }

    // ... rest of existing enter logic unchanged ...
    m.isStreaming = true
    m.requestStartTime = time.Now()
    // etc.
```

- [ ] **Step 4: Run test, confirm it passes**

```bash
go test ./internal/tui/... -run TestSlashInputDoesNotTriggerStreaming -v
```
Expected: `PASS`

- [ ] **Step 5: Write failing test — status bar shows active mode**

Add to `internal/tui/tui_test.go`:

```go
func TestStatusBarShowsAgentMode(t *testing.T) {
    m := newTestModel(t)
    m.agentMode = components.AgentEnumAdvanced

    status := m.statusLine()
    assert.Contains(t, status, "ADVANCED")
}
```

- [ ] **Step 6: Run test, confirm it fails**

```bash
go test ./internal/tui/... -run TestStatusBarShowsAgentMode -v
```
Expected: `FAIL` — status bar doesn't show mode yet

- [ ] **Step 7: Update `statusLine()` in `view.go` to show active mode**

In `internal/tui/view.go`, replace the `statusLine` `left` block:

```go
// Left side: mode badge + turn count.
modeLabel := styleStatusAccent.Render(string(m.agentMode))
var left string
turns := len(m.session.Turns)
if turns > 0 {
    left = modeLabel +
        styleStatusBar.Render("  ·  ") +
        styleStatusAccent.Render(fmt.Sprintf("%d", turns)) +
        styleStatusBar.Render(" turn"+pluralS(turns))
} else {
    left = modeLabel
}
```

- [ ] **Step 8: Run test, confirm it passes**

```bash
go test ./internal/tui/... -run TestStatusBarShowsAgentMode -v
```
Expected: `PASS`

- [ ] **Step 9: Run full test suite**

```bash
go test ./... && echo "ALL PASS"
```
Expected: all green

- [ ] **Step 10: Commit**

```bash
git add internal/tui/model.go internal/tui/view.go internal/tui/tui_test.go
git commit -m "feat(tui): intercept slash commands in enter handler, show agent mode in status bar"
```

---

## Chunk 5: Help View Update + Integration Smoke Test

### Task 5: Add slash commands to help view and update ctrl+h shortcuts list

**Files:**
- Modify: `internal/tui/view.go` (helpView)

- [ ] **Step 1: Update `helpView()` to list slash commands**

In `internal/tui/view.go`, in the `helpView()` function's `shortcuts` slice, add entries after the existing keyboard shortcuts:

```go
{"", ""}, // divider
{"/clear", "Start a new session"},
{"/mode fast|advanced|auto", "Switch agent reasoning depth"},
{"/help", "Show this help"},
```

The empty `{"", ""}` entry renders as a blank line divider — add a special case in the render loop to skip key rendering for blank keys:

```go
for _, s := range shortcuts {
    if s.key == "" {
        sb.WriteString("\n")
        continue
    }
    line := "  " +
        styleHelpKey.Render(fmt.Sprintf("%-30s", s.key)) +
        "  " +
        styleHelpDesc.Render(s.desc)
    sb.WriteString(line + "\n")
}
```

Update the format width from `%-26s` to `%-30s` to accommodate the longer slash command strings.

- [ ] **Step 2: Build and verify help view compiles**

```bash
go build ./... && echo "OK"
```

- [ ] **Step 3: Run full test suite one final time**

```bash
go test ./... && echo "ALL PASS"
```
Expected: all green

- [ ] **Step 4: Final commit**

```bash
git add internal/tui/view.go
git commit -m "feat(tui): add slash commands to help overlay"
```

---

## Summary

After all tasks complete:

- `/clear` resets session (same state as ctrl+r)
- `/mode fast|advanced|auto` changes the `AgentEnum` used in subsequent API calls
- `/help` opens the help overlay
- Unknown commands show a friendly error in the viewport
- Mode badge (`AUTO`, `FAST`, `ADVANCED`) visible in the status bar at all times
- All slash commands are intercepted before any network call
- System messages appear in the viewport but are never sent to Glean

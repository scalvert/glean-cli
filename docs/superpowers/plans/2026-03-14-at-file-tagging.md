# @ File Context Tagging Implementation Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** When the user types `@` followed by a partial filename in the TUI input, show a file picker overlay; selecting a file reads it and prepends its content to the message sent to Glean.

**Architecture:** A new `filepicker.go` file holds all pure file-system functions (`scanFiles`, `readAttachedFile`, `buildFileContext`, `parseFileQuery`). The Model grows picker state fields and three methods (`updateFilePicker`, `selectPickerItem`, `closePicker`). Key interception in `Update()` routes up/down/enter/tab/esc to the picker when it is open. The enter handler prepends file content to the API message while keeping the viewport display clean (original question only). The picker renders between the bottom delimiter and the input box; `maxViewportHeight` shrinks to accommodate it.

**Tech Stack:** Go, Bubbletea, `os.ReadDir`, `filepath`, lipgloss styles

---

## File Map

| File | Change |
|------|--------|
| `internal/tui/filepicker.go` | **Create** — `attachedFile`, `parseFileQuery`, `scanFiles`, `readAttachedFile`, `buildFileContext` |
| `internal/tui/model.go` | Add 5 picker fields to `Model`; add `updateFilePicker`, `selectPickerItem`, `closePicker`; key interception; enter handler file injection |
| `internal/tui/view.go` | Render picker overlay + attached-files chip; adjust layout |
| `internal/tui/styles.go` | Add `stylePickerHeader`, `stylePickerItem`, `stylePickerSelected`, `styleAttached` |
| `internal/tui/tui_test.go` | Tests for pure functions and picker state transitions |

---

## Chunk 1: Pure File Functions

### Task 1: Create `internal/tui/filepicker.go`

**Files:**
- Create: `internal/tui/filepicker.go`
- Test: `internal/tui/tui_test.go`

- [ ] **Step 1: Write failing tests**

Add to `internal/tui/tui_test.go`:

```go
func TestParseFileQueryDetectsAt(t *testing.T) {
    query, ok := parseFileQuery("hello @src/")
    assert.True(t, ok)
    assert.Equal(t, "src/", query)
}

func TestParseFileQueryNoAt(t *testing.T) {
    _, ok := parseFileQuery("hello world")
    assert.False(t, ok)
}

func TestParseFileQueryAtWithSpaceAfter(t *testing.T) {
    // "@foo bar" — space after means the @token is done, not a live query
    _, ok := parseFileQuery("@foo bar")
    assert.False(t, ok)
}

func TestParseFileQueryAtEnd(t *testing.T) {
    query, ok := parseFileQuery("look at this @")
    assert.True(t, ok)
    assert.Equal(t, "", query)
}

func TestBuildFileContextPrependsFiles(t *testing.T) {
    files := []attachedFile{
        {Path: "go.mod", Content: "module foo"},
    }
    result := buildFileContext(files, "what does this do?")
    assert.Contains(t, result, "[File: go.mod]")
    assert.Contains(t, result, "module foo")
    assert.Contains(t, result, "what does this do?")
    // File content must come before user message
    assert.Less(t, strings.Index(result, "[File:"), strings.Index(result, "what does this do?"))
}

func TestBuildFileContextNoFiles(t *testing.T) {
    result := buildFileContext(nil, "hello")
    assert.Equal(t, "hello", result)
}

func TestReadAttachedFileRejectsBinary(t *testing.T) {
    f, err := os.CreateTemp(t.TempDir(), "binary*.bin")
    require.NoError(t, err)
    _, _ = f.Write([]byte{0x00, 0x01, 0x02})
    f.Close()

    _, err = readAttachedFile(f.Name())
    assert.Error(t, err)
    assert.Contains(t, err.Error(), "binary")
}

func TestReadAttachedFileTruncatesLargeFiles(t *testing.T) {
    f, err := os.CreateTemp(t.TempDir(), "large*.txt")
    require.NoError(t, err)
    _, _ = f.WriteString(strings.Repeat("x", 15_000))
    f.Close()

    af, err := readAttachedFile(f.Name())
    require.NoError(t, err)
    assert.LessOrEqual(t, len(af.Content), 10_200) // 10k + truncation notice
    assert.Contains(t, af.Content, "truncated")
}

func TestReadAttachedFileReadsNormalFile(t *testing.T) {
    f, err := os.CreateTemp(t.TempDir(), "normal*.go")
    require.NoError(t, err)
    _, _ = f.WriteString("package main\n")
    f.Close()

    af, err := readAttachedFile(f.Name())
    require.NoError(t, err)
    assert.Equal(t, "package main\n", af.Content)
}
```

- [ ] **Step 2: Run tests, confirm they fail**

```bash
cd /Users/steve.calvert/workspace/personal/glean-cli
go test ./internal/tui/... -run "TestParseFile|TestBuildFile|TestReadAttached" -v 2>&1 | head -20
```
Expected: `FAIL` — functions not defined yet

- [ ] **Step 3: Create `internal/tui/filepicker.go`**

```go
package tui

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const maxFileBytes = 10_000

// attachedFile holds a file that the user has tagged with @ for inclusion
// in the next chat message sent to Glean.
type attachedFile struct {
	Path    string
	Content string
}

// parseFileQuery extracts the partial path typed after the last `@` in s.
// Returns ("", false) if no live @ query is present.
// A space after @ means the token is complete — not a live query.
func parseFileQuery(s string) (string, bool) {
	idx := strings.LastIndex(s, "@")
	if idx < 0 {
		return "", false
	}
	rest := s[idx+1:]
	// Space or tab after @ means the @ was already processed or is part of an email
	if strings.ContainsAny(rest, " \t\n") {
		return "", false
	}
	return rest, true
}

// scanFiles returns up to 20 file/directory paths that match the partial query.
// Hidden files are excluded unless the query starts with ".".
func scanFiles(query string) []string {
	dir := "."
	prefix := query

	if i := strings.LastIndex(query, "/"); i >= 0 {
		dir = query[:i]
		if dir == "" {
			dir = "/"
		}
		prefix = query[i+1:]
	}

	if strings.HasPrefix(dir, "~/") {
		home, err := os.UserHomeDir()
		if err == nil {
			dir = filepath.Join(home, dir[2:])
		}
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}

	var results []string
	for _, e := range entries {
		name := e.Name()
		if strings.HasPrefix(name, ".") && !strings.HasPrefix(prefix, ".") {
			continue
		}
		if prefix != "" && !strings.HasPrefix(strings.ToLower(name), strings.ToLower(prefix)) {
			continue
		}
		path := filepath.Join(dir, name)
		if dir == "." {
			path = name
		}
		if e.IsDir() {
			path += "/"
		}
		results = append(results, path)
		if len(results) >= 20 {
			break
		}
	}
	return results
}

// readAttachedFile reads a file and returns an attachedFile.
// Returns an error for binary files (null bytes detected) or unreadable files.
// Content is truncated to maxFileBytes with a notice if needed.
func readAttachedFile(path string) (attachedFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return attachedFile{}, fmt.Errorf("reading %s: %w", path, err)
	}

	// Binary detection: scan first 512 bytes for null bytes.
	sample := data
	if len(sample) > 512 {
		sample = sample[:512]
	}
	if bytes.ContainsRune(sample, 0) {
		return attachedFile{}, fmt.Errorf("%s appears to be a binary file — cannot attach", path)
	}

	content := string(data)
	if len(content) > maxFileBytes {
		content = content[:maxFileBytes] + "\n[... truncated at 10,000 characters ...]"
	}
	return attachedFile{Path: path, Content: content}, nil
}

// buildFileContext prepends file contents to the user message in a format
// Glean's AI can parse. Returns the original message unchanged if no files.
func buildFileContext(files []attachedFile, userMessage string) string {
	if len(files) == 0 {
		return userMessage
	}
	var sb strings.Builder
	for _, f := range files {
		sb.WriteString(fmt.Sprintf("[File: %s]\n```\n%s\n```\n\n", f.Path, f.Content))
	}
	sb.WriteString(userMessage)
	return sb.String()
}
```

- [ ] **Step 4: Add missing imports to tui_test.go**

The new tests need `"os"` — check the imports in `internal/tui/tui_test.go` and add it if missing.

- [ ] **Step 5: Run tests, confirm they pass**

```bash
go test ./internal/tui/... -run "TestParseFile|TestBuildFile|TestReadAttached" -v
```
Expected: all PASS

- [ ] **Step 6: Run full suite**

```bash
cd /Users/steve.calvert/workspace/personal/glean-cli && go test ./...
```
All green.

- [ ] **Step 7: Commit**

```bash
git add internal/tui/filepicker.go internal/tui/tui_test.go
git commit -m "feat(tui): add filepicker.go — parseFileQuery, scanFiles, readAttachedFile, buildFileContext"
```

---

## Chunk 2: Model State + Picker Methods

### Task 2: Add picker fields to Model and picker methods

**Files:**
- Modify: `internal/tui/model.go`
- Test: `internal/tui/tui_test.go`

- [ ] **Step 1: Write failing tests**

Add to `internal/tui/tui_test.go`:

```go
func TestUpdateFilePickerOpensOnAt(t *testing.T) {
    m := newTestModel(t)
    m.textarea.SetValue("@")
    m.updateFilePicker()
    // Current directory has files so picker should open
    // (if no files in CWD this test would fail — run from repo root)
    assert.True(t, m.showFilePicker)
}

func TestUpdateFilePickerClosesWhenNoAt(t *testing.T) {
    m := newTestModel(t)
    m.showFilePicker = true
    m.filePickerItems = []string{"foo.go"}
    m.textarea.SetValue("hello world")
    m.updateFilePicker()
    assert.False(t, m.showFilePicker)
}

func TestClosePickerResetsState(t *testing.T) {
    m := newTestModel(t)
    m.showFilePicker = true
    m.filePickerItems = []string{"a.go", "b.go"}
    m.filePickerIdx = 1
    m.closePicker()
    assert.False(t, m.showFilePicker)
    assert.Nil(t, m.filePickerItems)
    assert.Equal(t, 0, m.filePickerIdx)
}
```

- [ ] **Step 2: Run tests, confirm they fail**

```bash
go test ./internal/tui/... -run "TestUpdateFilePicker|TestClosePicker" -v 2>&1 | head -10
```
Expected: `FAIL` — fields not defined yet

- [ ] **Step 3: Add picker fields to `Model` struct in `internal/tui/model.go`**

After `conversationActive bool`, add:

```go
// File picker state — active when user types @ in the input.
showFilePicker  bool
filePickerQuery string
filePickerItems []string
filePickerIdx   int
attachedFiles   []attachedFile // files queued for next message
```

- [ ] **Step 4: Add picker methods to `internal/tui/model.go`** (append before `newGlamourRenderer`):

```go
// updateFilePicker inspects the textarea value after each keystroke and
// opens/refreshes/closes the file picker based on whether a live @query exists.
func (m *Model) updateFilePicker() {
	query, ok := parseFileQuery(m.textarea.Value())
	if !ok {
		if m.showFilePicker {
			m.closePicker()
		}
		return
	}
	items := scanFiles(query)
	if len(items) == 0 {
		if m.showFilePicker {
			m.closePicker()
		}
		return
	}
	m.filePickerQuery = query
	m.filePickerItems = items
	if !m.showFilePicker {
		m.showFilePicker = true
		m.filePickerIdx = 0
		m.recalculateLayout()
	} else if m.filePickerIdx >= len(m.filePickerItems) {
		m.filePickerIdx = len(m.filePickerItems) - 1
	}
}

// closePicker hides the file picker and resets its state.
func (m *Model) closePicker() {
	m.showFilePicker = false
	m.filePickerItems = nil
	m.filePickerIdx = 0
	m.recalculateLayout()
}

// selectPickerItem handles Enter/Tab on the highlighted picker item.
// Directories drill in; files are read, attached, and the @query removed.
func (m *Model) selectPickerItem() {
	if m.filePickerIdx >= len(m.filePickerItems) {
		return
	}
	selected := m.filePickerItems[m.filePickerIdx]

	// Directory — drill in by updating the textarea @path.
	if strings.HasSuffix(selected, "/") {
		current := m.textarea.Value()
		idx := strings.LastIndex(current, "@")
		if idx >= 0 {
			m.textarea.SetValue(current[:idx+1] + selected)
			m.textarea.CursorEnd()
		}
		m.updateFilePicker()
		return
	}

	// File — read it, attach, remove @query from textarea.
	if len(m.attachedFiles) >= 3 {
		m.addSystemMessage("Maximum 3 files per message")
		m.closePicker()
		return
	}

	af, err := readAttachedFile(selected)
	if err != nil {
		m.addSystemMessage(fmt.Sprintf("Cannot attach %s: %s", selected, err.Error()))
		m.closePicker()
		return
	}

	m.attachedFiles = append(m.attachedFiles, af)

	// Strip @query from textarea.
	current := m.textarea.Value()
	atIdx := strings.LastIndex(current, "@")
	if atIdx >= 0 {
		m.textarea.SetValue(current[:atIdx])
		m.textarea.CursorEnd()
	}

	m.closePicker()
	m.addSystemMessage(fmt.Sprintf("📎 %s attached (%d/%d)", filepath.Base(selected), len(m.attachedFiles), 3))
}
```

- [ ] **Step 5: Add `"path/filepath"` to imports in model.go** if not already present (check current imports).

- [ ] **Step 6: Run tests, confirm they pass**

```bash
go test ./internal/tui/... -run "TestUpdateFilePicker|TestClosePicker" -v
```
Expected: PASS (TestUpdateFilePickerOpensOnAt may be environment-dependent; if no files in CWD just verify it doesn't panic)

- [ ] **Step 7: Build**

```bash
cd /Users/steve.calvert/workspace/personal/glean-cli && go build ./...
```

- [ ] **Step 8: Commit**

```bash
git add internal/tui/model.go internal/tui/tui_test.go
git commit -m "feat(tui): add file picker state fields and updateFilePicker/selectPickerItem/closePicker methods"
```

---

## Chunk 3: Key Interception + Enter Handler

### Task 3: Route keys through picker, inject file context on send

**Files:**
- Modify: `internal/tui/model.go`
- Test: `internal/tui/tui_test.go`

- [ ] **Step 1: Write failing tests**

Add to `internal/tui/tui_test.go`:

```go
func TestPickerUpDownNavigation(t *testing.T) {
    m := newTestModel(t)
    m.showFilePicker = true
    m.filePickerItems = []string{"a.go", "b.go", "c.go"}
    m.filePickerIdx = 1

    // Down moves idx forward
    updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
    r := updated.(*Model)
    assert.Equal(t, 2, r.filePickerIdx)

    // Up moves idx back
    updated2, _ := r.Update(tea.KeyMsg{Type: tea.KeyUp})
    r2 := updated2.(*Model)
    assert.Equal(t, 1, r2.filePickerIdx)
}

func TestPickerEscClosesPicker(t *testing.T) {
    m := newTestModel(t)
    m.showFilePicker = true
    m.filePickerItems = []string{"a.go"}

    updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEsc})
    r := updated.(*Model)
    assert.False(t, r.showFilePicker)
}

func TestEnterWithAttachedFilesInjectsContext(t *testing.T) {
    m := newTestModel(t)
    m.conversationActive = true
    m.attachedFiles = []attachedFile{
        {Path: "go.mod", Content: "module test"},
    }
    m.textarea.SetValue("explain this")

    updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
    r := updated.(*Model)

    // Attached files cleared after send
    assert.Empty(t, r.attachedFiles)
    // conversationMsgs last entry has file context
    require.NotEmpty(t, r.conversationMsgs)
    lastMsg := r.conversationMsgs[len(r.conversationMsgs)-1]
    require.NotEmpty(t, lastMsg.Fragments)
    assert.Contains(t, *lastMsg.Fragments[0].Text, "[File: go.mod]")
    assert.Contains(t, *lastMsg.Fragments[0].Text, "explain this")
}

func TestEnterWithNoAttachedFilesSendsNormalMessage(t *testing.T) {
    m := newTestModel(t)
    m.conversationActive = true
    m.textarea.SetValue("hello")

    updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
    r := updated.(*Model)

    require.NotEmpty(t, r.conversationMsgs)
    lastMsg := r.conversationMsgs[len(r.conversationMsgs)-1]
    require.NotEmpty(t, lastMsg.Fragments)
    assert.Equal(t, "hello", *lastMsg.Fragments[0].Text)
}
```

- [ ] **Step 2: Run tests, confirm they fail**

```bash
go test ./internal/tui/... -run "TestPickerUp|TestPickerEsc|TestEnterWith" -v 2>&1 | head -20
```
Expected: `FAIL`

- [ ] **Step 3: Add picker key interception in `Update()` in `model.go`**

At the very top of the `case tea.KeyMsg:` block, **before** the existing `switch msg.String()`, add:

```go
case tea.KeyMsg:
    // Route navigation keys to the file picker when it is open.
    if m.showFilePicker {
        switch msg.String() {
        case "up", "ctrl+p":
            if m.filePickerIdx > 0 {
                m.filePickerIdx--
            }
            return m, nil
        case "down", "ctrl+n":
            if m.filePickerIdx < len(m.filePickerItems)-1 {
                m.filePickerIdx++
            }
            return m, nil
        case "enter", "tab":
            m.selectPickerItem()
            return m, nil
        case "esc":
            m.closePicker()
            return m, nil
        }
    }

    switch msg.String() {
    // ... existing cases unchanged ...
```

- [ ] **Step 4: Call `updateFilePicker()` after textarea update**

At the bottom of `Update()`, after `m.textarea, taCmd = m.textarea.Update(msg)`:

```go
m.textarea, taCmd = m.textarea.Update(msg)
if _, isKey := msg.(tea.KeyMsg); !isKey {
    m.viewport, vpCmd = m.viewport.Update(msg)
}
// Refresh file picker based on latest textarea value.
if _, isKey := msg.(tea.KeyMsg); isKey {
    m.updateFilePicker()
}
return m, tea.Batch(taCmd, vpCmd)
```

- [ ] **Step 5: Modify the `enter` handler to inject file context**

In the `case "enter":` block, replace the section that builds `turn` and appends to `conversationMsgs` with:

```go
// Build the message sent to the Glean API.
// If files are attached, prepend their content; display only the original question.
apiContent := question
if len(m.attachedFiles) > 0 {
    apiContent = buildFileContext(m.attachedFiles, question)
    m.attachedFiles = nil
}

// Display turn shows the original question only.
m.session.AddTurn(roleUser, question, nil)

// API message carries full file context when present.
apiText := apiContent
m.conversationMsgs = append(m.conversationMsgs, components.ChatMessage{
    Author:      components.AuthorUser.ToPointer(),
    MessageType: components.MessageTypeContent.ToPointer(),
    Fragments:   []components.ChatMessageFragment{{Text: &apiText}},
})

m.viewport.SetContent(m.renderConversation())
m.viewport.GotoBottom()

apiCmd := m.callAPI()
swCmd := m.stopwatch.Reset()
return m, tea.Batch(m.spinner.Tick, apiCmd, swCmd, m.stopwatch.Start())
```

Note: Remove the existing `turn := Turn{...}` / `m.addTurnToConversation(turn)` / `m.session.AddTurn(...)` lines and replace with the block above. The `addTurnToConversation` call is removed because we're directly appending to `conversationMsgs` above to inject full file context. The `session.AddTurn` still happens for viewport display (original question).

Also remove `m.viewport.SetContent(m.renderConversation())` and `m.viewport.GotoBottom()` that follow immediately after the old `addTurnToConversation` if they're duplicated — keep only one call at the end.

- [ ] **Step 6: Run tests, confirm they pass**

```bash
go test ./internal/tui/... -run "TestPickerUp|TestPickerEsc|TestEnterWith" -v
```
Expected: PASS

- [ ] **Step 7: Run full suite**

```bash
go test ./...
```
All green.

- [ ] **Step 8: Commit**

```bash
git add internal/tui/model.go internal/tui/tui_test.go
git commit -m "feat(tui): picker key interception, enter handler injects file context into API message"
```

---

## Chunk 4: View + Styles + Layout

### Task 4: Render picker overlay, attached file chip, adjust layout

**Files:**
- Modify: `internal/tui/view.go`
- Modify: `internal/tui/styles.go`
- Modify: `internal/tui/model.go` (layout constants)

- [ ] **Step 1: Add picker styles to `internal/tui/styles.go`**

```go
// File picker overlay styles.
stylePickerHeader = lipgloss.NewStyle().
    Foreground(lipgloss.Color(colorMuted)).
    Italic(true)

stylePickerItem = lipgloss.NewStyle().
    Foreground(lipgloss.Color(colorMuted))

stylePickerSelected = lipgloss.NewStyle().
    Foreground(lipgloss.Color(colorBrand)).
    Bold(true)

styleAttached = lipgloss.NewStyle().
    Foreground(lipgloss.Color(colorBrand))
```

- [ ] **Step 2: Add `filePickerView()` method to `view.go`**

```go
// filePickerView renders the file picker overlay shown when the user types @.
// Shows at most 5 items; the selected item is highlighted.
func (m *Model) filePickerView() string {
    if !m.showFilePicker || len(m.filePickerItems) == 0 {
        return ""
    }
    maxItems := 5
    items := m.filePickerItems
    if len(items) > maxItems {
        items = items[:maxItems]
    }
    var sb strings.Builder
    sb.WriteString(stylePickerHeader.Render("  @ file") + "\n")
    for i, item := range items {
        if i == m.filePickerIdx {
            sb.WriteString(stylePickerSelected.Render("  ▸ " + item))
        } else {
            sb.WriteString(stylePickerItem.Render("    " + item))
        }
        sb.WriteString("\n")
    }
    return sb.String()
}
```

- [ ] **Step 3: Add `attachedFilesView()` to `view.go`**

```go
// attachedFilesView renders a one-line chip for each file staged for the next message.
func (m *Model) attachedFilesView() string {
    if len(m.attachedFiles) == 0 {
        return ""
    }
    var parts []string
    for _, f := range m.attachedFiles {
        parts = append(parts, styleAttached.Render("📎 "+filepath.Base(f.Path)))
    }
    return "  " + strings.Join(parts, "   ")
}
```

Add `"path/filepath"` to `view.go` imports if not already present.

- [ ] **Step 4: Update `View()` in `view.go` to render picker and attached chip**

Replace the active-state `return` in `View()`:

```go
rule := styleDelimiter.Render(strings.Repeat("─", m.width))

parts := []string{header, rule, m.viewport.View(), rule}

if picker := m.filePickerView(); picker != "" {
    parts = append(parts, picker)
}
if chip := m.attachedFilesView(); chip != "" {
    parts = append(parts, chip)
}
parts = append(parts, inputBox, bottom)

return lipgloss.JoinVertical(lipgloss.Left, parts...)
```

- [ ] **Step 5: Adjust `maxViewportHeight()` in `model.go` to shrink when picker is open**

In `maxViewportHeight()`, add picker height to the constants:

```go
func (m *Model) maxViewportHeight() int {
    if m.width == 0 || m.height == 0 {
        return 4
    }
    const (
        inputH  = 3
        statusH = 1
        spacerH = 2 // top + bottom delimiters
    )
    pickerH := 0
    if m.showFilePicker {
        n := len(m.filePickerItems)
        if n > 5 {
            n = 5
        }
        pickerH = n + 1 // items + header line
    }
    chipH := 0
    if len(m.attachedFiles) > 0 {
        chipH = 1
    }
    h := m.height - logoHeaderLines - spacerH - inputH - statusH - pickerH - chipH
    if h < 4 {
        return 4
    }
    return h
}
```

Apply the same adjustment to `resizeViewportToContent()` — add `pickerH` and `chipH` to the maxVpH calculation there too.

- [ ] **Step 6: Build**

```bash
cd /Users/steve.calvert/workspace/personal/glean-cli && go build ./...
```

- [ ] **Step 7: Run full test suite**

```bash
go test ./...
```
All green.

- [ ] **Step 8: Commit**

```bash
git add internal/tui/view.go internal/tui/styles.go internal/tui/model.go
git commit -m "feat(tui): render file picker overlay and attached-file chip; adjust layout for picker height"
```

---

## Summary

After all tasks complete:
- Typing `@` in the input opens a file picker showing matching files in the current directory
- ↑/↓ navigate the picker; Enter/Tab selects; Esc dismisses
- Selecting a directory drills in; selecting a file reads it and shows `📎 filename.go` above the input
- Up to 3 files can be attached per message
- Binary files and unreadable files are rejected with a system message
- Files >10,000 chars are truncated with a notice
- On send: file contents prepended to the API message; viewport shows original question only
- The picker closes automatically when the `@` is cleared or no matches remain

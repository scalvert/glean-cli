# TUI Enhancements — Slash Commands, @ File Context, Images

## Overview

Three independent features that make the Glean TUI feel like a first-class AI interface.
Each is specced and implemented as its own unit.

---

## Feature 1: Slash Commands (`/clear`, `/mode`)

### What
When the user types `/` as the first character and presses enter, the input is
treated as a command rather than a chat message. No network call is made.

### Commands

| Command | Behavior |
|---------|----------|
| `/clear` | Resets the session (same as ctrl+r) |
| `/mode fast` | Sets agent to `AgentEnumFast` |
| `/mode advanced` | Sets agent to `AgentEnumAdvanced` |
| `/mode auto` | Sets agent to `AgentEnumAuto` (default) |
| `/help` | Shows the help overlay |

### Design

**Model additions:**
- `agentMode components.AgentEnum` — current agent (default: `AgentEnumAuto`)

**Input interception** (in `enter` handler, before API call):
```
if strings.HasPrefix(question, "/") {
    return m.handleSlashCommand(question)
}
```

**`handleSlashCommand(input string) (tea.Model, tea.Cmd)`**
- Parses the command token and optional arg
- Dispatches to the appropriate handler
- Returns a system message rendered in the viewport:
  `✓ Mode set to ADVANCED` (styled with muted/accent colors)
- Unknown commands show: `Unknown command: /foo — try /help`

**`callAPI()` update:**
- Uses `m.agentMode` instead of hardcoded `AgentEnumDefault`

**Status bar update:**
- Shows current mode when not streaming: `AUTO  ·  3 turns  ·  ctrl+r new …`

**Autocomplete (nice-to-have, not required for v1):**
- When input is exactly `/`, show inline hint: `/clear  /mode  /help`

### Acceptance
- `/mode auto|fast|advanced` changes subsequent API calls
- `/clear` resets session
- Unknown `/` commands show error in viewport, don't send to API
- Mode indicator visible in status bar

---

## Feature 2: @ File Context Tagging

### What
When the user types `@` in the input, a file picker overlay appears. Selecting a
file reads it and injects its content into the chat message sent to Glean.

### Design

**Model additions:**
- `showFilePicker bool`
- `filePickerQuery string` — text after `@`
- `filePickerItems []string` — filtered file candidates
- `filePickerIdx int` — selection cursor
- `attachedFiles []attachedFile` — files staged for this message

**`attachedFile` struct:**
```go
type attachedFile struct {
    Path    string
    Content string // truncated to max ~10k chars
}
```

**Trigger logic:**
- On every keystroke, scan textarea value for `@` followed by a partial path
- If found: compute `filePickerQuery`, populate `filePickerItems` via
  `filepath.Glob` and `os.ReadDir` in the current directory and common
  prefixes (`./`, `~/`, absolute)
- Render a small popup above the input (max 8 items, scrollable)

**Popup rendering** (in `View()`, between input and status bar):
```
  @ src/main.go
  @ src/config.go          ← highlighted
  @ src/handlers/
```

**Selection:**
- `↑`/`↓` navigate
- `enter` or `tab` selects: reads file, appends to `attachedFiles`, removes `@query` from textarea
- `esc` dismisses picker

**Message construction:**
When the user sends a message with attached files, prepend to the message text:
```
[File: src/config.go]
```go
package config
...
```
]

<user's original message>
```
This goes into the `Fragments[0].Text` field sent to the Glean chat API.

**Limits:**
- Max 3 files per message
- Max 10,000 characters per file (truncated with notice)
- Binary files: reject with inline error

### Acceptance
- `@` triggers picker within 100ms
- Selected file content prepended to chat message
- Glean AI responds with awareness of file content
- Binary/oversized files handled gracefully

---

## Feature 3: Drag & Drop / Paste Image Support

### What
User can drop an image onto the terminal or paste image data (clipboard). The
image is encoded and sent to the Glean chat API as a multimodal attachment.

### Constraints
- Terminal image paste uses **OSC 52** (clipboard) or **iTerm2 inline image protocol**
- Bubbletea does not natively surface image paste events — requires raw terminal
  event interception or a custom `tea.ExecCommand` approach
- Glean's chat API supports image attachments via `ChatFile` in `ChatMessageFragment`

### Design

**Detection approach:**
Bubbletea surfaces paste as `tea.KeyMsg` with bracketed paste sequences or as
`tea.PasteMsg` (v0.25+). Check if pasted bytes are a valid image (PNG/JPEG header).

**Model additions:**
- `pendingImage *pendingImageAttachment`

**`pendingImageAttachment` struct:**
```go
type pendingImageAttachment struct {
    Data     []byte
    MimeType string // "image/png" or "image/jpeg"
    Preview  string // braille/chafa thumbnail, shown in viewport
}
```

**UX flow:**
1. User drags/pastes image → detected as paste event
2. Image rendered as a small braille preview in the viewport (using chafa or
   a pure-Go braille renderer for thumbnails)
3. Status shows: `📎 image attached — press enter to send`
4. On enter: `ChatMessageFragment.File` populated with base64-encoded image

**API construction:**
```go
fragment := components.ChatMessageFragment{
    File: &components.ChatFile{
        MimeType: pendingImage.MimeType,
        Data:     base64.StdEncoding.EncodeToString(pendingImage.Data),
    },
}
```

**Fallback:**
If terminal doesn't support image paste detection (non-iTerm2/kitty), the feature
silently does nothing. No error shown.

### Acceptance
- Pasted PNG/JPEG detected and previewed in viewport
- Image sent with chat message to Glean
- Non-image paste continues to work normally (text paste into textarea)
- Graceful no-op on unsupported terminals

---

## Implementation Order

1. **Slash commands** — 1 session, low risk, immediately useful (`/mode` unlocks AUTO/FAST/ADVANCED)
2. **@ file context** — 1–2 sessions, high value for agent workflows
3. **Drag & drop images** — 1–2 sessions, highest complexity, terminal-dependent

Each feature gets its own implementation plan via `writing-plans`.

# Glean CLI — Development Rules

## Non-Negotiable: Verify Before Claiming Done

**Never say a feature is complete without running the actual binary.**

```bash
cd /Users/steve.calvert/workspace/glean/glean-cli
mise run build          # build the binary
./glean                 # run the TUI and test the feature manually
```

For TUI features specifically:
- Build the binary, run it, interact with the feature being built
- If it involves keystrokes (e.g. `/mode`, `@` file picker), actually type them
- If it involves streaming (stages, spinner, elapsed), ask a real question
- Screenshots from the user are ground truth — believe them over unit tests

Unit tests pass ≠ feature works. The TUI has UI state that tests cannot cover.

## Mandatory Pre-Push Requirement

**Before any `git push`, run ALL of the following locally. They MUST pass.**

```bash
go test ./...           # all tests must pass
golangci-lint run       # linter must be clean (matches CI config in .golangci.yml)
mise run build          # binary must build
```

Equivalently: `mise run test:all` (runs lint + test + build).

CI runs `golangci-lint` v2 with `gocritic`, `gosec`, `errcheck`, and other strict
linters configured in `.golangci.yml`. If your local `golangci-lint` version is older,
it may miss issues that CI catches. Always verify lint passes before pushing.

Pushing code that breaks CI is unacceptable. No exceptions.

## Project

- **Repo**: `github.com/gleanwork/glean-cli` — pre-1.0, breaking changes OK
- **Module**: `github.com/gleanwork/glean-cli`
- **Task runner**: `mise` (not go-task)
- **Go version**: 1.24

## Key Commands

```bash
mise run build          # build glean binary
mise run test           # run tests
mise run test:all       # lint + test + build (CI equivalent)
./glean                 # run TUI (requires valid GLEAN_HOST + token)
```

## Architecture

- `internal/tui/` — Bubbletea TUI (model.go, view.go, commands.go, filepicker.go, session.go, styles.go)
- `internal/client/` — Glean SDK wrapper + streaming HTTP client
- `internal/auth/` — OAuth PKCE + DCR flow
- `cmd/` — Cobra commands (search, chat, api, config, etc.)

## TUI Architecture Rules

- **Worktree LSP errors are noise** — the worktree (`/.claude/worktrees/`) is stale. Always build in the main repo at `/Users/steve.calvert/workspace/glean/glean-cli`
- **Viewport key isolation** — never pass `tea.KeyMsg` to viewport in the catch-all; scroll keys are handled explicitly above
- **Collect-then-display for content** — streaming stages show live via channel; content waits until complete
- **No viewport jumping** — `conversationActive` pins viewport height; `recalculateLayout()` only called on terminal resize or deliberate state changes

## Auth Test Isolation

**All tests that touch auth, config, or keyring state MUST call `authtest.IsolateAuthState(t)`.**

```go
import "github.com/gleanwork/glean-cli/internal/auth/authtest"

func TestSomethingWithAuth(t *testing.T) {
    authtest.IsolateAuthState(t)
    // ... test code
}
```

This redirects HOME, config path, and keyring to a temp directory so tests never read or delete real credentials. Without it, tests can silently wipe `~/.glean/config.json` and keyring entries.

## Common Mistakes to Avoid

1. **Rendering picker only in active state** — UI elements visible before a conversation starts need to work in the welcome state too (`!m.conversationActive` branch in View())
2. **Skipping `mise run build && ./glean` verification** — unit tests don't catch TUI rendering bugs
3. **Spinner below input** — spinner goes in the viewport content area, never in statusLine()
4. **Large decimal elapsed** — show integer seconds (`12s` not `11.6s`)

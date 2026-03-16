# <img src="demo/glean-logo.png" width="28" height="28" style="vertical-align: middle; margin-right: 4px"> Glean CLI

Your company's knowledge, search, and AI — from the command line.

Search across your company's knowledge, chat with Glean Assistant, manage the full Glean API surface, and integrate Glean into automated workflows — all without leaving the terminal.

<p>
  <a href="https://github.com/gleanwork/glean-cli/releases"><img src="https://img.shields.io/github/v/release/gleanwork/glean-cli" alt="latest release"></a>
  <a href="https://github.com/gleanwork/glean-cli/blob/main/LICENSE"><img src="https://img.shields.io/github/license/gleanwork/glean-cli" alt="license"></a>
  <a href="https://github.com/gleanwork/glean-cli/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/gleanwork/glean-cli/ci.yml?branch=main&label=CI" alt="CI status"></a>
</p>

## Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Why Glean CLI?](#why-glean-cli)
- [Authentication](#authentication)
- [Interactive TUI](#interactive-tui)
- [Commands](#commands)
- [Agent Workflow](#agent-workflow)
- [Environment Variables](#environment-variables)
- [Exit Codes](#exit-codes)
- [Contributing](#contributing)

## Installation

```bash
# Homebrew (recommended)
brew install gleanwork/tap/glean-cli

# Manual
curl -fsSL https://raw.githubusercontent.com/gleanwork/glean-cli/main/install.sh | sh
```

Pre-built binaries for macOS, Linux, and Windows are available on the [Releases](https://github.com/gleanwork/glean-cli/releases) page.

## Quick Start

```bash
# 1. Authenticate
glean auth login                          # OAuth via browser (recommended)
# — OR set env vars for CI/CD:
# export GLEAN_HOST=your-company-be.glean.com GLEAN_API_TOKEN=your-token

# 2. Search
glean search "vacation policy"

# 3. Chat
glean chat "Summarize our Q1 engineering goals"

# 4. Open the interactive TUI
glean
```

## Why Glean CLI?

If your workflow involves searching Glean, asking Glean AI questions, or wiring Glean data into scripts or agent pipelines — the CLI is faster than a browser tab and composable with everything else in your terminal.

Every command returns structured JSON. Use `--dry-run` to preview requests before they're sent. Use `glean schema <command>` to get machine-readable flag documentation. Results pipe cleanly to `jq`, scripts, or any tool that reads stdin.

```bash
# Discover what commands are available
glean schema | jq '.commands'

# Preview a request before sending
glean search --dry-run --datasource confluence "Q1 planning"

# Parse results with jq — each result has .title and nested .document fields
glean search "onboarding" | jq '.results[].title'

# Stream results as NDJSON — one SearchResult object per line
glean search "engineering docs" --output ndjson | jq .title
```

## Authentication

### OAuth (recommended)

```bash
glean auth login    # opens browser, completes PKCE flow
glean auth status   # verify credentials, host, and token expiry
glean auth logout   # remove all stored credentials
```

OAuth uses PKCE with Dynamic Client Registration — no client ID required. Tokens are stored securely in the system keyring and refreshed automatically.

For instances that don't support OAuth, `auth login` falls back to prompting for an API token.

### API Token (CI/CD)

Set credentials via environment variables — no interactive login needed:

```bash
export GLEAN_API_TOKEN=your-token
export GLEAN_HOST=your-company-be.glean.com
glean search "test"
```

Credentials are resolved in this order: environment variables → system keyring → `~/.glean/config.json`.

## Interactive TUI

Running `glean` with no arguments opens a full-screen chat powered by Glean Assistant.

```bash
glean            # open TUI
glean --continue # resume the most recent session
```

### Keyboard Shortcuts

| Key | Action |
|-----|--------|
| `enter` | Send message |
| `↑` / `↓` | Scroll history / navigate command picker |
| `ctrl+r` | New session |
| `ctrl+l` | Clear screen |
| `ctrl+y` | Copy last response |
| `ctrl+h` | Toggle help |
| `ctrl+c` / `esc` | Quit |

### Slash Commands

Type `/` in the input to open the command picker:

| Command | Description |
|---------|-------------|
| `/mode auto\|fast\|advanced` | Switch agent reasoning depth |
| `/clear` | Start a new session |
| `/help` | Show keyboard shortcuts |

### File Attachments

Type `@` followed by a filename to attach a local file to your message. The file's contents are sent as context to Glean AI.

```
@go.mod          # attach go.mod from current directory
@src/config.go   # attach a specific file
```

Use ↑/↓ to navigate matches, Enter to attach, Esc to dismiss.

## Commands

### Core

| Command | Description |
|---------|-------------|
| `glean search <query>` | Search across your company's knowledge |
| `glean chat <message>` | Chat with Glean Assistant (non-interactive) |
| `glean api <endpoint>` | Make a raw authenticated HTTP request to the Glean REST API |
| `glean schema [command]` | Show machine-readable JSON schema for any command |
| `glean mcp` | Start a stdio MCP server for AI agent integration |
| `glean auth` | Authenticate with Glean |
| `glean version` | Print the CLI version |

### `glean search`

```bash
glean search "vacation policy"
glean search "Q1 planning" --datasource confluence --page-size 5
glean search "docs" --fields "results.document.title,results.document.url"
glean search "docs" --output ndjson | jq .title
glean search --json '{"query":"onboarding","pageSize":3}'
glean search --dry-run "test"
```

| Flag | Description |
|------|-------------|
| `--output` / `--format` | `json` (default), `ndjson` (one result per line), `text` |
| `--fields` | Dot-path field projection — prefix paths with `results.` |
| `--datasource` / `-d` | Filter by datasource (repeatable) |
| `--type` / `-t` | Filter by document type (repeatable) |
| `--page-size` | Results per page (default 10) |
| `--json` | Raw SDK request body (overrides all flags) |
| `--dry-run` | Print request body without sending |

### `glean chat`

```bash
glean chat "What are our company holidays?"
glean chat --timeout 120000 "Summarize all Q1 OKRs across teams"
glean chat --json '{"messages":[{"author":"USER","messageType":"CONTENT","fragments":[{"text":"What is Glean?"}]}]}'
glean chat --dry-run "test"
```

| Flag | Description |
|------|-------------|
| `--timeout` | Request timeout in milliseconds (default 60000) |
| `--json` | Raw SDK request body (overrides all flags) |
| `--dry-run` | Print request body without sending |
| `--save` | Persist chat for continuation (default true) |

### `glean api`

Raw authenticated HTTP access to any Glean REST API endpoint (relative to `/rest/api/v1/`).

```bash
glean api search --method POST --raw-field '{"query":"rust","pageSize":3}'
glean api --preview search --method POST --raw-field '{"query":"test"}'
```

### API Namespace Commands

All namespace commands accept `--json`, `--output`, and `--dry-run`. Run `glean <command> --help` for full usage.

| Namespace | Subcommands | Description |
|-----------|-------------|-------------|
| `glean agents` | `list`, `get`, `schemas`, `run` | Manage and invoke Glean AI agents |
| `glean answers` | `list`, `get`, `create`, `update`, `delete` | Curated Q&A pairs |
| `glean announcements` | `create`, `update`, `delete` | Time-bounded company announcements |
| `glean collections` | `list`, `get`, `create`, `update`, `delete`, `add-items`, `delete-item` | Curated document collections |
| `glean documents` | `get`, `summarize`, `get-by-facets`, `get-permissions` | Document retrieval and summarization |
| `glean entities` | `list`, `read-people` | People, teams, and custom entities |
| `glean insights` | `get` | Search and usage analytics |
| `glean messages` | `get` | Retrieve indexed messages (Slack, Teams, etc.) |
| `glean pins` | `list`, `get`, `create`, `update`, `remove` | Promoted search results |
| `glean shortcuts` | `list`, `get`, `create`, `update`, `delete` | Go-links / memorable short URLs |
| `glean tools` | `list`, `run` | Glean platform tools |
| `glean verification` | `list`, `verify`, `remind` | Document verification and review |
| `glean activity` | `report`, `feedback` | User activity reporting |

#### Example payloads

```bash
# Retrieve a document by URL
glean documents get --json '{"documentSpecs":[{"url":"https://..."}]}'

# Summarize a document
glean documents summarize --json '{"documentSpecs":[{"url":"https://..."}]}'

# Look up people
glean entities list --json '{"entityType":"PEOPLE","query":"engineering"}'

# Create a go-link
glean shortcuts create --json '{"data":{"inputAlias":"onboarding","destinationUrl":"https://..."}}'

# Create a shortcut with a variable template
glean shortcuts create --json '{"data":{"inputAlias":"jira","urlTemplate":"https://jira.example.com/browse/{arg}"}}'

# Pin a result for a query
glean pins create --json '{"queries":["onboarding"],"documentId":"https://..."}'

# List available AI agents
glean agents list | jq '.SearchAgentsResponse.agents[] | {id: .agent_id, name: .name}'

# Run an agent
glean agents run --json '{"agent_id":"<id>"}'
```

## Agent Workflow

The CLI is designed as a first-class tool for AI coding agents. Every command returns JSON on stdout and errors on stderr with non-zero exit codes.

```bash
# 1. Discover all available commands
glean schema | jq '.commands'

# 2. Understand a command's flags
glean schema search | jq '.flags | keys'
glean schema search | jq '.flags["--output"]'

# 3. Preview the exact request before sending
glean shortcuts create --dry-run \
  --json '{"data":{"inputAlias":"test","destinationUrl":"https://example.com"}}'

# 4. Execute and parse results
glean search "engineering values" | jq '.results[].title'

# 5. Stream NDJSON for large result sets — one SearchResult object per line
glean search "all docs" --output ndjson --page-size 50 | jq .title
```

### MCP Server

For AI agents that support MCP, run `glean mcp` to expose Glean as an MCP tool server:

```bash
glean mcp
```

Add to `.claude/settings.json` for Claude Code:

```json
{
  "mcpServers": {
    "glean": { "command": "glean", "args": ["mcp"] }
  }
}
```

Available MCP tools: `glean_search`, `glean_chat`, `glean_schema`, `glean_people`.

## Environment Variables

| Variable | Description |
|----------|-------------|
| `GLEAN_API_TOKEN` | API token — overrides keyring and config file |
| `GLEAN_HOST` | Glean backend hostname (e.g. `your-company-be.glean.com`) |

Environment variables take precedence over stored configuration.

## Exit Codes

| Code | Meaning |
|------|---------|
| `0` | Success |
| `1` | General error (authentication failure, API error, invalid input) |

All error details are written to stderr. Stdout contains only structured output (JSON/NDJSON/text), making the CLI safe for piping.

## Shell Completions

```bash
glean completion bash   # Bash
glean completion zsh    # Zsh
glean completion fish   # Fish
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for development setup, coding conventions, and how to submit pull requests.

## License

MIT — see [LICENSE](LICENSE).

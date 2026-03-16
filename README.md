# <img src="demo/glean-logo.png" width="28" height="28" style="vertical-align: middle; margin-right: 4px"> Glean CLI

> Work seamlessly with Glean from your command line.

The Glean CLI (`glean`) brings Glean's search and AI capabilities directly to your terminal. Search across your company's knowledge, chat with Glean Assistant, invoke the full REST API surface, and wire Glean into AI agents via MCP — all from the command line.

## Installation

```bash
# Homebrew
brew install gleanwork/tap/glean-cli

# Manual
curl -fsSL https://raw.githubusercontent.com/gleanwork/glean-cli/main/install.sh | sh
```

## Quick Start

```bash
# 0. Authenticate (recommended: OAuth via browser)
glean auth login

# — OR — configure with an API token:
glean config --host your-company-be.glean.com --token YOUR_API_TOKEN

# 1. Search
glean search "vacation policy"

# 2. Chat
glean chat "Summarize our Q1 engineering goals"
```

## Interactive TUI

Running `glean` with no arguments opens a full-screen chat powered by Glean Assistant.

```bash
glean            # open TUI
glean --continue # resume the most recent session
```

### TUI Keyboard Shortcuts

| Key | Action |
|-----|--------|
| `enter` | Send message |
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

Type `@` followed by a filename to attach a file to your message. The file's contents are injected into the request as context for Glean AI.

```
@go.mod          # attach go.mod from current directory
@src/config.go   # attach a specific file
```

Up/Down to navigate matches, Enter to attach, Esc to dismiss.

Session history is persisted to `~/.glean/sessions/latest.json`. To clear it:
```bash
rm -f ~/.glean/sessions/latest.json
```

## Commands

### `glean search`

Search across your company's knowledge.

```bash
glean search "vacation policy"

# Pipe to jq
glean search "Q1 planning" | jq '.results[].document.title'

# Filter by datasource or type
glean search --datasource confluence "project planning"
glean search --type document "Q1 reports"

# Project specific fields (results.* prefix required)
glean search --fields "results.document.title,results.document.url" "onboarding"

# NDJSON — one result object per line, good for streaming pipelines
glean search --output ndjson "engineering docs" | jq .document.title

# Raw SDK request body
glean search --json '{"query":"Q1 reports","pageSize":5}'

# Preview request without sending
glean search --dry-run --datasource confluence "test"
```

Key flags:

| Flag | Description |
|------|-------------|
| `--output` / `--format` | `json` (default), `ndjson`, or `text` |
| `--fields` | Comma-separated dot-path projection (must include `results.` prefix) |
| `--json` | Complete JSON request body; overrides all other flags |
| `--dry-run` | Print request body without sending |
| `--datasource` / `-d` | Filter by datasource (repeatable) |
| `--type` / `-t` | Filter by document type (repeatable) |
| `--page-size` | Results per page (default 10) |

### `glean chat`

Ask Glean AI a question from the command line (non-interactive).

```bash
glean chat "What are our company holidays?"
glean chat "Summarize our engineering onboarding docs"

# Structured multi-message request
glean chat --json '{"messages":[{"author":"USER","messageType":"CONTENT","fragments":[{"text":"What is Glean?"}]}]}'

# Preview request without sending
glean chat --dry-run "test question"
```

Key flags:

| Flag | Description |
|------|-------------|
| `--json` | Complete JSON chat request body; overrides all other flags |
| `--dry-run` | Print request body without sending |
| `--timeout` | Request timeout in milliseconds (default 60000 — 60 seconds) |
| `--save` | Save the chat for later continuation (default true) |

### `glean auth`

Authenticate with Glean.

```bash
glean auth login    # OAuth via browser (recommended)
glean auth status   # check current auth state
glean auth logout   # remove stored credentials
```

OAuth uses PKCE with Dynamic Client Registration — no client ID configuration required for supported instances. For instances without OAuth, `auth login` falls back to an API token prompt.

### `glean api`

Make an authenticated HTTP request to any Glean REST API endpoint.

```bash
# GET
glean api search --method POST --raw-field '{"query":"rust programming","pageSize":3}'

# From a file
glean api search --method POST --input search-params.json

# Preview request (shows URL, headers, body)
glean api --preview search --method POST --raw-field '{"query":"test"}'
```

Endpoints are relative to `/rest/api/v1/`. Valid endpoints include: `search`, `chat`, `announcements`, `answers`, `collections`, `documents`, `entities`, `pins`, `shortcuts`, `tools`, and more.

### `glean schema`

Machine-readable JSON schema for any command — useful for agents discovering how to call the CLI.

```bash
glean schema              # list all 18 registered commands
glean schema search       # flags, types, defaults, and examples for search
glean schema shortcuts    # schema for shortcuts subcommands
```

### `glean config`

Manage credentials and connection settings.

```bash
glean config --host your-company-be.glean.com
glean config --token your-token
glean config --show               # show current config (token masked)
glean config --show --output json # machine-readable output
glean config --clear              # remove all stored credentials
```

Configuration is stored in the system keyring with fallback to `~/.glean/config.json`.

### `glean mcp`

Start a stdio MCP server that exposes Glean tools to AI agents (Claude Code, Cursor, etc.).

```bash
glean mcp
```

Available tools: `glean_search`, `glean_chat`, `glean_schema`, `glean_people`.

To wire into Claude Code, add to `.claude/settings.json`:

```json
{
  "mcpServers": {
    "glean": { "command": "glean", "args": ["mcp"] }
  }
}
```

## API Namespace Commands

Thin passthroughs to the full Glean SDK surface. Every subcommand accepts `--json`, `--output`, and `--dry-run`.

```bash
glean <command> --help          # full usage and examples
glean <command> list --dry-run  # preview request before sending
```

| Command | Subcommands | Key example |
|---------|-------------|-------------|
| `glean agents` | list, get, schemas, run | `glean agents list` |
| `glean answers` | list, get, create, update, delete | `glean answers list` |
| `glean collections` | list, get, create, update, delete | `glean collections list` |
| `glean documents` | get, summarize, get-by-facets, get-permissions | `glean documents get --json '{"documentSpecs":[{"url":"https://..."}]}'` |
| `glean entities` | list, read-people | `glean entities list --json '{"entityType":"PEOPLE","query":"engineering"}'` |
| `glean shortcuts` | list, get, create, update, delete | `glean shortcuts create --json '{"data":{"inputAlias":"onboarding","destinationUrl":"https://..."}}'` |
| `glean pins` | list, get, create, update, remove | `glean pins list` |
| `glean tools` | list, run | `glean tools list` |
| `glean announcements` | create, update, delete | `glean announcements create --json '{"title":"...","startTime":"...","endTime":"..."}'` |
| `glean activity` | report, feedback | `glean activity report --json '{"events":[...]}'` |
| `glean insights` | get | `glean insights get --json '{}'` |
| `glean messages` | get | `glean messages get --json '{"idType":"THREAD_ID","id":"...","datasource":"SLACK"}'` |
| `glean verification` | list, verify, remind | `glean verification list` |

## Agent Workflow

The CLI is designed as a first-class tool for AI coding agents. The recommended workflow:

```bash
# 1. Discover what's available
glean schema | jq '.commands'

# 2. Understand a command's flags and request shape
glean schema search | jq '.flags[] | {name, type, default}'

# 3. Preview the exact request before sending
glean search --dry-run --datasource confluence "onboarding" | jq .

# 4. Execute and parse results
glean search "engineering best practices" \
  --fields "results.document.title,results.document.url" \
  | jq '.results[].document'

# 5. Enumerate available AI agents
glean agents list | jq '.SearchAgentsResponse.agents[] | {id: .agent_id, name: .displayName}'
```

Every command returns JSON on stdout and errors on stderr with non-zero exit codes — predictable for scripting.

## Shell Completions

```bash
glean completion bash   # Bash
glean completion zsh    # Zsh
glean completion fish   # Fish
```

## Environment Variables

| Variable | Description |
|----------|-------------|
| `GLEAN_API_TOKEN` | API token (overrides keyring/config file) |
| `GLEAN_HOST` | Glean backend hostname (e.g. `your-company-be.glean.com`) |
| `GLEAN_EMAIL` | Email for impersonated API requests |

Environment variables take precedence over the system keyring and `~/.glean/config.json`.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for development setup and how to submit pull requests.

## License

MIT — see [LICENSE](LICENSE).

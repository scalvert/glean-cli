# <img src="demo/glean-logo.png" width="28" height="28" style="vertical-align: middle; margin-right: 4px"> Glean CLI (Unofficial)

> Work seamlessly with Glean from your command line.

![Glean CLI Demo](demo/readme.gif)

The Glean CLI (`glean`) brings Glean's search and AI capabilities directly to your terminal. Search across your company's knowledge, chat with Glean Assistant, invoke the full REST API surface, and wire Glean into AI agents via MCP — all from the command line.

## Installation

```bash
# Homebrew
brew install scalvert/tap/glean-cli

# Manual
curl -fsSL https://raw.githubusercontent.com/scalvert/glean-cli/main/install.sh | sh
```

## Quick Start

```bash
# 1. Configure credentials
glean config --host your-company --token your-token

# 2. Search
glean search "vacation policy"
```

## Default Behavior

Running `glean` with no arguments opens a full-screen interactive chat TUI powered by Glean Assistant. Session history is persisted across invocations. Use `--new` to start a fresh session.

```bash
glean           # open interactive TUI chat
glean --new     # open TUI with a blank session
```

## Commands

### `glean search`

Search across your company's knowledge.

```bash
glean search "vacation policy"
glean search "vacation policy" | jq '.results[].document.title'

# JSON output (default) — pipe-friendly
glean search --output json "meeting notes"

# NDJSON — one result per line
glean search --output ndjson "engineering docs" | head -3 | jq .document.title

# Project specific fields
glean search --fields "document.title,document.url" "onboarding"

# Filter by datasource or document type
glean search --datasource confluence "project planning"
glean search --type document "Q1 reports"

# Send a raw SDK request body
glean search --json '{"query":"Q1 reports","pageSize":5}' | jq .

# Preview request without sending
glean search --dry-run "test"
```

Key flags:

| Flag | Description |
|------|-------------|
| `--output` | `json` (default), `ndjson`, or `text` |
| `--fields` | Comma-separated dot-path projection (e.g. `document.title,document.url`) |
| `--json` | Complete JSON request body; overrides all other flags |
| `--dry-run` | Print request body without sending |
| `--datasource` | Filter by datasource (repeatable, `-d`) |
| `--type` | Filter by document type (repeatable, `-y`) |
| `--page-size` | Results per page (default 10) |
| `--timeout` | Request timeout in milliseconds (default 30000) |

### `glean chat`

Ask Glean AI a question from the command line (non-interactive, streaming output).

```bash
glean chat "What are our company holidays?"
glean chat --timeout 60000 "Tell me about our engineering team"
glean chat --save=false "What's our tech stack?"

# Send a raw SDK request body
glean chat --json '{"messages":[{"author":"USER","messageType":"CONTENT","fragments":[{"text":"What is Glean?"}]}]}'

# Preview request without sending
glean chat --dry-run "test question"
```

Key flags:

| Flag | Description |
|------|-------------|
| `--json` | Complete JSON chat request body; overrides all other flags |
| `--dry-run` | Print request body without sending |
| `--timeout` | Request timeout in milliseconds (default 30000) |
| `--save` | Save the chat for later continuation (default true) |

### `glean api`

Make an authenticated HTTP request to any Glean API endpoint.

```bash
# GET request
glean api users/me

# POST with a JSON body via --raw-field
glean api search --method POST --raw-field '{"query": "rust programming"}'

# POST with a body from a file
glean api search --method POST --input search-params.json

# POST with a body piped from stdin
echo '{"query": "rust"}' | glean api --method POST search

# Preview request without sending
glean api search --method POST --raw-field '{"query": "test"}' --preview

# Pipe to jq
glean api search --no-color | jq .results
```

Note: POST/PUT requests require a body. Provide it via `--raw-field`, `--input`, or stdin pipe.

Key flags:

| Flag | Description |
|------|-------------|
| `--method` / `-X` | HTTP method (default `GET`) |
| `--raw-field` | JSON string body |
| `--input` / `-F` | File to use as request body |
| `--preview` | Print request details without sending |
| `--no-color` | Disable colorized output |

### `glean config`

Manage credentials and connection settings.

```bash
# Set host — either short name or full hostname works
glean config --host linkedin
glean config --host linkedin-be.glean.com

# Set API token
glean config --token your-token

# Set user email
glean config --email you@company.com

# Set multiple values at once
glean config --host your-company --token your-token

# Show current configuration (token is masked)
glean config --show

# Clear all stored credentials
glean config --clear
```

Configuration is stored in the system keyring with fallback to `~/.glean/config.json`.

### `glean mcp`

Start a stdio MCP server that exposes Glean tools to AI agents (Claude Code, Cursor, etc.).

```bash
glean mcp
```

Available MCP tools: `glean_search`, `glean_chat`, `glean_schema`, `glean_people`.

To wire it into Claude Code, add to `.claude/settings.json`:

```json
{
  "mcpServers": {
    "glean": {
      "command": "glean",
      "args": ["mcp"]
    }
  }
}
```

### `glean schema`

Show machine-readable JSON schema for any command's flags and request format. Useful for agents.

```bash
glean schema            # list all commands with registered schemas
glean schema search     # full schema for the search command
glean schema chat       # full schema for the chat command
```

### API Namespace Commands

The following commands are thin passthroughs to the Glean SDK's API surface. Each accepts `--json`, `--output`, and `--dry-run`. Run `glean <command> --help` for details.

| Command | Description |
|---------|-------------|
| `glean activity` | Report user activity and feedback |
| `glean agents` | Manage and run Glean agents |
| `glean announcements` | Manage announcements |
| `glean answers` | Manage answers |
| `glean collections` | Manage collections |
| `glean documents` | Retrieve and summarize documents |
| `glean entities` | List and read entities and people |
| `glean insights` | Retrieve usage insights |
| `glean messages` | Retrieve messages |
| `glean pins` | Manage pins |
| `glean shortcuts` | Manage shortcuts (go-links) |
| `glean tools` | List and run Glean tools |
| `glean verification` | Manage document verification |

## Shell Completions

```bash
glean completion bash   # Bash
glean completion zsh    # Zsh
glean completion fish   # Fish
```

Follow the instructions printed by each command to install the completion script for your shell.

## Environment Variables

| Variable | Description |
|----------|-------------|
| `GLEAN_API_TOKEN` | API token (overrides keyring/config file) |
| `GLEAN_HOST` | Glean instance name or hostname |
| `GLEAN_EMAIL` | Email address for API requests |

Environment variables take precedence over the system keyring and `~/.glean/config.json`.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for how to submit pull requests and the project's code of conduct.

## License

MIT — see [LICENSE](LICENSE).

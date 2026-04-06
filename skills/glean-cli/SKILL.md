---
name: glean-cli
description: "Glean CLI: access company knowledge, search documents, chat with Glean Assistant, look up people, and manage enterprise content. Use when the user asks about internal docs, company information, people, policies, or enterprise data."
compatibility: Requires the glean binary on $PATH. Install via brew install gleanwork/tap/glean-cli
---

# Glean CLI

The `glean` command-line tool provides authenticated access to your company's Glean instance. It can search documents, chat with Glean Assistant, look up people and teams, manage collections, and more.

## When to Use

Use the Glean CLI when the user:

- Asks about internal documents, policies, wikis, or company knowledge
- Wants to find people, teams, or org structure
- Needs to search across enterprise data sources (Confluence, Jira, Google Drive, Slack, etc.)
- Asks questions that require company-specific context
- Wants to manage Glean resources (collections, shortcuts, pins, announcements)

## Installation

```bash
brew install gleanwork/tap/glean-cli
```

## Authentication

```bash
# Browser-based OAuth (interactive — recommended)
glean auth login

# Verify credentials
glean auth status

# CI/scripting (no interactive setup needed)
export GLEAN_API_TOKEN=your-token
export GLEAN_HOST=your-company-be.glean.com
```

Credentials resolve in this order: environment variables → system keyring → ~/.glean/config.json.

## CLI Syntax

```bash
glean <command> [subcommand] [flags]
```

### Global Flags

| Flag | Description |
|------|-------------|
| --output <FORMAT> | json (default), ndjson (one result per line), text |
| --fields <PATHS> | Dot-path field projection (e.g. results.document.title,results.document.url) |
| --json <PAYLOAD> | Complete JSON request body (overrides all other flags) |
| --dry-run | Print request body without sending |

## Schema Introspection

Always call `glean schema <command>` before invoking a command you haven't used before.

```bash
glean schema | jq '.commands'          # list all commands
glean schema search | jq '.flags'      # flags for search
```

## Security Rules

- **Never** output API tokens or secrets directly
- **Always** use --dry-run before write/delete operations in automated pipelines
- Prefer environment variables over config files for CI/CD

## Error Handling

All errors go to stderr; stdout contains only structured output.
Exit code 0 = success, non-zero = error.

## Available Commands

| Command | Description |
|---------|-------------|
| [glean activity](../glean-cli-activity/SKILL.md) | Report user activity and feedback. Subcommands: report, feedback. |
| [glean agents](../glean-cli-agents/SKILL.md) | Manage and run Glean agents. Subcommands: list, get, schemas, run. |
| [glean announcements](../glean-cli-announcements/SKILL.md) | Manage Glean announcements. Subcommands: create, update, delete. |
| [glean answers](../glean-cli-answers/SKILL.md) | Manage Glean answers. Subcommands: list, get, create, update, delete. |
| [glean api](../glean-cli-api/SKILL.md) | Make a raw authenticated HTTP request to any Glean REST API endpoint. |
| [glean chat](../glean-cli-chat/SKILL.md) | Have a conversation with Glean AI. Streams response to stdout. |
| [glean collections](../glean-cli-collections/SKILL.md) | Manage Glean collections. Subcommands: create, delete, update, add-items, delete-item. |
| [glean documents](../glean-cli-documents/SKILL.md) | Retrieve and summarize Glean documents. Subcommands: get, get-by-facets, get-permissions, summarize. |
| [glean entities](../glean-cli-entities/SKILL.md) | List and read Glean entities and people. Subcommands: list, read-people. |
| [glean insights](../glean-cli-insights/SKILL.md) | Retrieve Glean usage insights. Subcommands: get. |
| [glean messages](../glean-cli-messages/SKILL.md) | Retrieve Glean messages. Subcommands: get. |
| [glean pins](../glean-cli-pins/SKILL.md) | Manage Glean pins. Subcommands: list, get, create, update, remove. |
| [glean search](../glean-cli-search/SKILL.md) | Search for content in your Glean instance. Results are JSON. |
| [glean shortcuts](../glean-cli-shortcuts/SKILL.md) | Manage Glean shortcuts (go-links). Subcommands: list, get, create, update, delete. |
| [glean tools](../glean-cli-tools/SKILL.md) | List and run Glean tools. Subcommands: list, run. |
| [glean verification](../glean-cli-verification/SKILL.md) | Manage document verification. Subcommands: list, verify, remind. |


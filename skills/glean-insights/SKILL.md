---
name: glean-insights
description: "Retrieve search and usage analytics from Glean. Use when analyzing search patterns, popular queries, or platform adoption metrics."
---

# glean insights

> **PREREQUISITE:** Read `../glean-shared/SKILL.md` for auth, global flags, and security rules.

Retrieve Glean usage insights. Subcommands: get.

```bash
glean insights <subcommand> [flags]
```

## Subcommands

| Subcommand | Description |
|------------|-------------|
| `get` | Get analytics data |

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--dry-run` | boolean | false |  |
| `--json` | string |  | JSON request body (required) **(required)** |
| `--output` | json \| ndjson \| text | json |  |

## Examples

```bash
glean insights get --json '{"insightTypes":["SEARCH"]}' | jq .
```

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema insights

# List all available commands
glean schema | jq '.commands'
```

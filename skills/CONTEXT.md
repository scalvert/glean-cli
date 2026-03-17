---
name: glean-cli-context
description: CLI-wide guidance for using the glean command-line tool as an AI agent
---

# Glean CLI — Agent Context

`glean` is a CLI for searching and interacting with Glean's company knowledge platform.

## Key Principles

1. **Schema-first**: Always call `glean schema <command>` before invoking a command to understand
   current parameter types, required fields, and output format. Never guess at flag names.

2. **JSON-in, JSON-out**: All commands accept `--json <payload>` to pass a complete request body.
   All commands output JSON by default — pipe to `jq` for filtering.

3. **Dry-run before write**: Use `--dry-run` on create/update/delete operations to inspect the
   request body before sending.

4. **Ctrl+C-safe**: All commands respect context cancellation. Long-running operations can be
   interrupted.

## Authentication

Three credential sources (priority order):
1. `GLEAN_API_TOKEN` + `GLEAN_HOST` env vars (preferred for CI/scripting)
2. OAuth token via `glean auth login`
3. `~/.glean/config.json`

```bash
# CI/scripting (no interactive setup needed)
GLEAN_API_TOKEN=mytoken GLEAN_HOST=mycompany-be.glean.com glean search "test"
```

## Quick Start

```bash
# Discover available commands and their schemas
glean schema
glean schema search

# Search company knowledge
glean search "vacation policy" | jq '.results[0].title'
glean search --json '{"query":"Q1 reports","pageSize":5}' | jq '.results[].title'

# Chat with Glean AI
glean chat "What are our engineering principles?"
```

## Output Formats

- `--output json` (default): single JSON object, pipe-safe
- `--output ndjson`: one JSON line per result, good for streaming pipelines
- `--output text`: human-readable (avoid in agent pipelines)
- `--fields a,b.c`: project only specific dot-path fields

## Error Handling

All errors are written to stderr; stdout only contains structured output.
Exit code 0 = success, non-zero = error.

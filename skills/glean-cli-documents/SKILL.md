---
name: glean-cli-documents
description: "Retrieve, summarize, and inspect documents indexed by Glean. Use when getting document content, summaries, permissions, or metadata by URL."
---

# glean documents

> **PREREQUISITE:** Read `../glean-cli/SKILL.md` for auth, global flags, and security rules.

Retrieve and summarize Glean documents. Subcommands: get, get-by-facets, get-permissions, summarize.

```bash
glean documents <subcommand> [flags]
```

## Subcommands

| Subcommand | Description |
|------------|-------------|
| `get` | Retrieve document metadata by URL or ID |
| `get-by-facets` | Retrieve documents matching facet filters |
| `get-permissions` | Inspect who has access to a document |
| `summarize` | Generate an AI summary of a document |

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--dry-run` | boolean | false |  |
| `--json` | string |  | JSON request body |
| `--output` | json \| ndjson \| text | json |  |

## Examples

```bash
glean documents summarize --json '{"documentId":"DOC_ID"}' | jq .summary
```

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema documents

# List all available commands
glean schema | jq '.commands'
```

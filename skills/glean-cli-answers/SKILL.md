---
name: glean-cli-answers
description: "Manage curated Q&A pairs in Glean. Use when creating, updating, or listing company-approved answers to common questions."
---

# glean answers

> **PREREQUISITE:** Read `../glean-cli-shared/SKILL.md` for auth, global flags, and security rules.

Manage Glean answers. Subcommands: list, get, create, update, delete.

```bash
glean answers <subcommand> [flags]
```

## Subcommands

| Subcommand | Description |
|------------|-------------|
| `create` | Create a new answer |
| `delete` | Delete an answer |
| `get` | Get a specific answer |
| `list` | List all curated answers |
| `update` | Update an existing answer |

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--dry-run` | boolean | false |  |
| `--json` | string |  | JSON request body |
| `--output` | json \| ndjson \| text | json |  |

## Examples

```bash
glean answers list | jq '.[].id'
```

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema answers

# List all available commands
glean schema | jq '.commands'
```

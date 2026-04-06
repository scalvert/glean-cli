---
name: glean-cli-shortcuts
description: "Manage Glean go-links (shortcuts). Use when creating, listing, updating, or deleting short URL aliases like go/wiki."
---

# glean shortcuts

> **PREREQUISITE:** Read `../glean-cli-shared/SKILL.md` for auth, global flags, and security rules.

Manage Glean shortcuts (go-links). Subcommands: list, get, create, update, delete.

```bash
glean shortcuts <subcommand> [flags]
```

## Subcommands

| Subcommand | Description |
|------------|-------------|
| `create` | Create a new shortcut |
| `delete` | Delete a shortcut |
| `get` | Get a specific shortcut |
| `list` | List all shortcuts |
| `update` | Update an existing shortcut |

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--dry-run` | boolean | false | Print request without sending |
| `--json` | string |  | JSON request body (see Glean API docs for shape) |
| `--output` | json \| ndjson \| text | json |  |

## Examples

```bash
glean shortcuts list | jq '.results[].inputAlias'
glean shortcuts create --json '{"data":{"inputAlias":"test/link","destinationUrl":"https://example.com"}}'
```

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema shortcuts

# List all available commands
glean schema | jq '.commands'
```

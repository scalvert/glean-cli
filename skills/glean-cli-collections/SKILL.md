---
name: glean-cli-collections
description: "Manage curated document collections in Glean. Use when creating, updating, or organizing themed sets of documents."
---

# glean collections

> **PREREQUISITE:** Read `../glean-cli/SKILL.md` for auth, global flags, and security rules.

Manage Glean collections. Subcommands: create, delete, update, add-items, delete-item.

```bash
glean collections <subcommand> [flags]
```

## Subcommands

| Subcommand | Description |
|------------|-------------|
| `add-items` | Add documents to a collection |
| `create` | Create a new collection |
| `delete` | Delete a collection |
| `delete-item` | Remove a document from a collection |
| `get` | Get a specific collection |
| `list` | List all collections |
| `update` | Update a collection |

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--dry-run` | boolean | false |  |
| `--json` | string |  | JSON request body |
| `--output` | json \| ndjson \| text | json |  |

## Examples

```bash
glean collections create --json '{"name":"My Collection"}'
```

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema collections

# List all available commands
glean schema | jq '.commands'
```

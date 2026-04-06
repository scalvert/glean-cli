---
name: glean-cli-announcements
description: "Manage time-bounded company announcements in Glean. Use when creating, updating, or deleting announcements that surface across the Glean UI."
---

# glean announcements

> **PREREQUISITE:** Read `../glean-cli/SKILL.md` for auth, global flags, and security rules.

Manage Glean announcements. Subcommands: create, update, delete.

```bash
glean announcements <subcommand> [flags]
```

## Subcommands

| Subcommand | Description |
|------------|-------------|
| `create` | Create a new announcement |
| `delete` | Delete an announcement |
| `update` | Update an existing announcement |

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--dry-run` | boolean | false |  |
| `--json` | string |  | JSON request body (required) **(required)** |
| `--output` | json \| ndjson \| text | json |  |

## Examples

```bash
glean announcements create --json '{"title":"Company Update","body":"..."}'
```

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema announcements

# List all available commands
glean schema | jq '.commands'
```

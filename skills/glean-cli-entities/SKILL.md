---
name: glean-cli-entities
description: "Look up people, teams, and custom entities in Glean. Use when finding employees, org structure, team members, or expertise."
---

# glean entities

> **PREREQUISITE:** Read `../glean-cli/SKILL.md` for auth, global flags, and security rules.

List and read Glean entities and people. Subcommands: list, read-people.

```bash
glean entities <subcommand> [flags]
```

## Subcommands

| Subcommand | Description |
|------------|-------------|
| `list` | List entities by type and query |
| `read-people` | Get detailed people profiles |

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--json` | string |  | JSON request body **(required)** |
| `--output` | json \| ndjson \| text | json |  |

## Examples

```bash
glean entities read-people --json '{"query":"smith"}' | jq '.[].name'
```

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema entities

# List all available commands
glean schema | jq '.commands'
```

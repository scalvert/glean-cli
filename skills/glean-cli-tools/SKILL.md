---
name: glean-cli-tools
description: "List and run Glean platform tools. Use when discovering available platform tools or executing them programmatically."
---

# glean tools

> **PREREQUISITE:** Read `../glean-cli-shared/SKILL.md` for auth, global flags, and security rules.

List and run Glean tools. Subcommands: list, run.

```bash
glean tools <subcommand> [flags]
```

## Subcommands

| Subcommand | Description |
|------------|-------------|
| `list` | List available platform tools |
| `run` | Execute a platform tool |

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--dry-run` | boolean | false |  |
| `--json` | string |  | JSON request body |
| `--output` | json \| ndjson \| text | json |  |

## Examples

```bash
glean tools list | jq '.[].name'
```

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema tools

# List all available commands
glean schema | jq '.commands'
```

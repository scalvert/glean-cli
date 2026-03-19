---
name: glean-activity
description: "Report user activity and submit feedback to Glean. Use when logging user interactions or providing relevance feedback on search results."
---

# glean activity

> **PREREQUISITE:** Read `../glean-shared/SKILL.md` for auth, global flags, and security rules.

Report user activity and feedback. Subcommands: report, feedback.

```bash
glean activity <subcommand> [flags]
```

## Subcommands

| Subcommand | Description |
|------------|-------------|
| `feedback` | Submit feedback on search results |
| `report` | Report a user activity event |

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--dry-run` | boolean | false |  |
| `--json` | string |  | JSON request body (required) **(required)** |

## Examples

```bash
glean activity report --json '{"events":[{"action":"VIEW","url":"https://example.com"}]}'
```

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema activity

# List all available commands
glean schema | jq '.commands'
```

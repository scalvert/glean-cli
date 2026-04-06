---
name: glean-cli-verification
description: "Manage document verification and review workflows in Glean. Use when verifying document accuracy, listing pending verifications, or sending review reminders."
---

# glean verification

> **PREREQUISITE:** Read `../glean-cli-shared/SKILL.md` for auth, global flags, and security rules.

Manage document verification. Subcommands: list, verify, remind.

```bash
glean verification <subcommand> [flags]
```

## Subcommands

| Subcommand | Description |
|------------|-------------|
| `list` | List documents pending verification |
| `remind` | Send a verification reminder |
| `verify` | Mark a document as verified |

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--dry-run` | boolean | false |  |
| `--json` | string |  | JSON request body |
| `--output` | json \| ndjson \| text | json |  |

## Examples

```bash
glean verification list | jq '.[].document.title'
```

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema verification

# List all available commands
glean schema | jq '.commands'
```

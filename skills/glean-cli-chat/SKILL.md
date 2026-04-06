---
name: glean-cli-chat
description: "Chat with Glean Assistant from the command line. Use when asking questions, summarizing documents, or getting AI-powered answers about company knowledge."
---

# glean chat

> **PREREQUISITE:** Read `../glean-cli/SKILL.md` for auth, global flags, and security rules.

Have a conversation with Glean AI. Streams response to stdout.

```bash
glean chat [flags]
```

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--dry-run` | boolean | false | Print request body without sending |
| `--json` | string |  | Complete JSON chat request body (overrides individual flags) |
| `--message` | string |  | Chat message (positional arg) **(required)** |
| `--save` | boolean | true | Save the chat session |
| `--timeout` | integer | 30000 | Request timeout in milliseconds |

## Examples

```bash
glean chat "What are the company holidays?"
glean chat --json '{"messages":[{"author":"USER","messageType":"CONTENT","fragments":[{"text":"What is Glean?"}]}]}'
```

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema chat

# List all available commands
glean schema | jq '.commands'
```

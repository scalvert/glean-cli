---
name: glean-chat
description: "Chat with Glean Assistant from the command line. Use when asking questions, summarizing documents, or getting AI-powered answers about company knowledge."
---

# glean chat

> **PREREQUISITE:** Read `../glean-shared/SKILL.md` for auth, global flags, and security rules.

Have a conversation with Glean AI. Streams response to stdout.

```bash
glean chat [flags]
```

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--dry-run` | boolean | false | Print request body without sending |
| `--json` | string |  | Complete JSON chat request body (overrides individual flags) |
| `--message` | string |  | Chat message (positional arg; reads from stdin if omitted) |
| `--save` | boolean | true | Save the chat session |
| `--timeout` | integer | 30000 | Request timeout in milliseconds |

## Examples

```bash
glean chat "What are the company holidays?"
glean chat --json '{"messages":[{"author":"USER","messageType":"CONTENT","fragments":[{"text":"What is Glean?"}]}]}'
echo "What is Glean?" | glean chat
glean chat                                # interactive multiline input, Ctrl+D to send
```

When called without a message argument, reads from stdin until EOF (Ctrl+D).
This enables multiline messages and piping input from other commands.

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema chat

# List all available commands
glean schema | jq '.commands'
```

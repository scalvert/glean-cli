---
command: chat
description: Have a conversation with Glean AI
required_flags:
  - message (positional arg) OR --json
output_format: text (streaming response)
---

# glean chat

Ask Glean AI a question. The response is streamed with stage indicators for searching and reading.

## Basic Usage

```bash
# Simple question
glean chat "What are the engineering principles?"

# Full JSON payload (for multi-turn or custom agent config)
glean chat --json '{
  "messages": [
    {"author": "USER", "messageType": "CONTENT", "fragments": [{"text": "What is Glean?"}]}
  ],
  "agentConfig": {"agent": "DEFAULT", "mode": "DEFAULT"}
}'

# See what would be sent without sending
glean chat --dry-run "test question"
```

## Non-Interactive Usage

For scripting, `glean chat` streams text to stdout. Capture it with:

```bash
# Capture response (blocks until complete)
ANSWER=$(glean chat "What is the vacation policy?" 2>/dev/null)
```

## Request Schema (components.ChatRequest)

```json
{
  "messages": [
    {
      "author": "USER",
      "messageType": "CONTENT",
      "fragments": [{"text": "your question here"}]
    }
  ],
  "agentConfig": {
    "agent": "DEFAULT",
    "mode": "DEFAULT"
  },
  "saveChat": true,
  "timeoutMillis": 30000,
  "stream": true
}
```

## Pitfalls

- Output includes stage markers (✓ Searching:, ✓ Reading:) in text mode
- For machine-parseable output, use the MCP server (`glean mcp`) instead
- `--timeout` is the API timeout, not a local timeout

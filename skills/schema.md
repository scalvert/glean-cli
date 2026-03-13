---
command: schema
description: Runtime JSON schema introspection for any glean command
required_flags: none
output_format: json
---

# glean schema

Machine-readable schema for any glean command. Call this before invoking a command to understand
parameter types, required fields, defaults, and example invocations.

## Usage

```bash
# List all commands with registered schemas
glean schema | jq .commands

# Get full schema for a command
glean schema search | jq .flags
glean schema chat | jq .example
glean schema shortcuts | jq .description
```

## Output Shape

```json
{
  "command": "search",
  "description": "...",
  "flags": {
    "--query": {"type": "string", "required": true, "description": "..."},
    "--page-size": {"type": "integer", "default": 10, "description": "..."},
    "--json": {"type": "string", "description": "Complete JSON request body"},
    "--output": {"type": "enum", "values": ["json", "ndjson", "text"], "default": "json"},
    "--dry-run": {"type": "boolean", "default": false}
  },
  "example": "glean search \"vacation policy\" | jq ..."
}
```

## Best Practice for Agents

Always call `glean schema <command>` at the start of a session to understand current parameter
types. This eliminates guessing and works without documentation in context.

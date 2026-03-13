---
command: mcp
description: Start a stdio MCP server exposing Glean tools to AI agents
required_flags: none
output_format: MCP protocol (JSON-RPC over stdio)
---

# glean mcp

Start a Model Context Protocol (MCP) stdio server. Agents connect to this server to invoke
Glean operations as structured tools.

## Setup with Claude Code

Add to your project's `.claude/settings.json`:

```json
{
  "mcpServers": {
    "glean": {
      "command": "glean",
      "args": ["mcp"]
    }
  }
}
```

Or with explicit credentials:

```json
{
  "mcpServers": {
    "glean": {
      "command": "glean",
      "args": ["mcp"],
      "env": {
        "GLEAN_API_TOKEN": "your-token",
        "GLEAN_HOST": "your-instance-be.glean.com"
      }
    }
  }
}
```

## Available MCP Tools

| Tool | Description | Required params |
|------|-------------|-----------------|
| `glean_search` | Search company knowledge | `query` |
| `glean_chat` | Ask Glean AI a question | `message` |
| `glean_schema` | Get CLI command schema | `command` (optional) |
| `glean_people` | Search for employees | `query` |

## Usage Pattern

Once connected, the agent can call:

```
glean_schema({"command": "search"})  → JSON schema for search
glean_search({"query": "vacation policy", "pageSize": 5})  → search results
glean_chat({"message": "What is our engineering culture?"})  → AI response
glean_people({"query": "smith"})  → list of matching employees
```

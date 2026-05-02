# glean agents

Manage and run Glean agents. Subcommands: list, get, schemas, run.

```bash
glean agents <subcommand> [flags]
```

## Subcommands

| Subcommand | Description |
|------------|-------------|
| `get` | Get details of a specific agent |
| `list` | List all available agents |
| `run` | Run an agent |
| `schemas` | Get input/output schemas for an agent |

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--dry-run` | boolean | false |  |
| `--json` | string |  | JSON request body |
| `--output` | json \| ndjson \| text | json |  |

## Examples

```bash
glean agents list | jq '.[].id'
glean agents run --json '{"agentId":"my-agent","input":{"query":"test"}}'
```

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema agents

# List all available commands
glean schema | jq '.commands'
```

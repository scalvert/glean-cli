# glean messages

Retrieve Glean messages. Subcommands: get.

```bash
glean messages <subcommand> [flags]
```

## Subcommands

| Subcommand | Description |
|------------|-------------|
| `get` | Retrieve a specific message |

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--json` | string |  | JSON request body (required) **(required)** |
| `--output` | json \| ndjson \| text | json |  |

## Examples

```bash
glean messages get --json '{"messageId":"MSG_ID"}' | jq .
```

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema messages

# List all available commands
glean schema | jq '.commands'
```

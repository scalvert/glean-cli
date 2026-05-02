# glean insights

Retrieve Glean usage insights. Subcommands: get.

```bash
glean insights <subcommand> [flags]
```

## Subcommands

| Subcommand | Description |
|------------|-------------|
| `get` | Get analytics data |

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--dry-run` | boolean | false |  |
| `--json` | string |  | JSON request body (required) **(required)** |
| `--output` | json \| ndjson \| text | json |  |

## Examples

```bash
glean insights get --json '{"insightTypes":["SEARCH"]}' | jq .
```

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema insights

# List all available commands
glean schema | jq '.commands'
```

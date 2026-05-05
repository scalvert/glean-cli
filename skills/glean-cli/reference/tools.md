# glean tools

List and run Glean tools. Subcommands: list, run.

```bash
glean tools <subcommand> [flags]
```

## Subcommands

| Subcommand | Description |
|------------|-------------|
| `list` | List available platform tools |
| `run` | Execute a platform tool |

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--dry-run` | boolean | false |  |
| `--json` | string |  | JSON request body |
| `--output` | json \| ndjson \| text | json |  |

## Examples

```bash
glean tools list | jq '.[].name'
```

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema tools

# List all available commands
glean schema | jq '.commands'
```

# glean pins

Manage Glean pins. Subcommands: list, get, create, update, remove.

```bash
glean pins <subcommand> [flags]
```

## Subcommands

| Subcommand | Description |
|------------|-------------|
| `create` | Create a new pin |
| `get` | Get a specific pin |
| `list` | List all pins |
| `remove` | Remove a pin |
| `update` | Update a pin |

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--dry-run` | boolean | false |  |
| `--json` | string |  | JSON request body |
| `--output` | json \| ndjson \| text | json |  |

## Examples

```bash
glean pins list | jq '.[].id'
```

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema pins

# List all available commands
glean schema | jq '.commands'
```

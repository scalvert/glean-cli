# glean announcements

Manage Glean announcements. Subcommands: create, update, delete.

```bash
glean announcements <subcommand> [flags]
```

## Subcommands

| Subcommand | Description |
|------------|-------------|
| `create` | Create a new announcement |
| `delete` | Delete an announcement |
| `update` | Update an existing announcement |

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--dry-run` | boolean | false |  |
| `--json` | string |  | JSON request body (required) **(required)** |
| `--output` | json \| ndjson \| text | json |  |

## Examples

```bash
glean announcements create --json '{"title":"Company Update","body":"..."}'
```

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema announcements

# List all available commands
glean schema | jq '.commands'
```

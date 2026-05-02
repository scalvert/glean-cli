# glean verification

Manage document verification. Subcommands: list, verify, remind.

```bash
glean verification <subcommand> [flags]
```

## Subcommands

| Subcommand | Description |
|------------|-------------|
| `list` | List documents pending verification |
| `remind` | Send a verification reminder |
| `verify` | Mark a document as verified |

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--dry-run` | boolean | false |  |
| `--json` | string |  | JSON request body |
| `--output` | json \| ndjson \| text | json |  |

## Examples

```bash
glean verification list | jq '.[].document.title'
```

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema verification

# List all available commands
glean schema | jq '.commands'
```

---
command: shortcuts
description: Manage Glean shortcuts (go-links)
required_flags: --json for create/update/delete
output_format: json (default) | ndjson | text
---

# glean shortcuts

Manage Glean go-links (shortcuts). These are short aliases like `go/wiki` that redirect to URLs.

## Usage

```bash
# List all shortcuts
glean shortcuts list | jq '.results[].inputAlias'

# Create a shortcut
glean shortcuts create --json '{
  "data": {
    "inputAlias": "team/eng",
    "destinationUrl": "https://wiki.company.com/engineering"
  }
}'

# Dry-run a create
glean shortcuts create --dry-run --json '{
  "data": {"inputAlias": "test/link", "destinationUrl": "https://example.com"}
}'

# Delete a shortcut
glean shortcuts delete --json '{"id": 12345}'
```

## Request Shapes

### Create: components.CreateShortcutRequest
```json
{
  "data": {
    "inputAlias": "team/link",
    "destinationUrl": "https://...",
    "description": "Optional description",
    "unlisted": false
  }
}
```

### Delete: components.DeleteShortcutRequest
```json
{"id": 12345}
```

## Pitfalls

- The `create` payload uses a `data` wrapper (`{"data": {...}}`), NOT a flat object
- `inputAlias` is the `go/` part after the prefix (e.g. `team/wiki` for `go/team/wiki`)

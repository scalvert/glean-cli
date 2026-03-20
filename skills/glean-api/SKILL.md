---
name: glean-api
description: "Make raw authenticated HTTP requests to any Glean REST API endpoint. Use when no dedicated command exists or for advanced API access."
---

# glean api

> **PREREQUISITE:** Read `../glean-shared/SKILL.md` for auth, global flags, and security rules.

Make a raw authenticated HTTP request to any Glean REST API endpoint.

```bash
glean api [flags]
```

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--dry-run` | boolean | false | Same as --preview |
| `--input` | string |  | Path to a JSON file to use as request body |
| `--method` | GET \| POST \| PUT \| DELETE \| PATCH | GET | HTTP method |
| `--no-color` | boolean | false | Disable colorized output |
| `--preview` | boolean | false | Print request details without sending |
| `--raw` | boolean | false | Print raw response without syntax highlighting |
| `--raw-field` | string |  | JSON request body as a string |

## Examples

```bash
glean api search --method POST --raw-field '{"query":"test"}' --no-color | jq .results
```

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema api

# List all available commands
glean schema | jq '.commands'
```

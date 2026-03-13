---
command: search
description: Search Glean company knowledge
required_flags:
  - query (positional arg) OR --json
output_format: json (default) | ndjson | text
---

# glean search

Search company knowledge indexed by Glean. Results include documents, snippets, and metadata.

## Basic Usage

```bash
# Simple query (JSON output by default)
glean search "vacation policy" | jq '.results[0].document.title'

# Extract specific fields
glean search "Q1 reports" --fields "results.document.title,results.document.url"

# NDJSON — one result per line
glean search "engineering docs" --output ndjson | head -3 | jq .document.title

# Filter by datasource
glean search "deploy guide" --datasource confluence

# Full JSON payload (overrides all flags)
glean search --json '{"query":"Q1 reports","pageSize":5,"datasources":["confluence","gdrive"]}'
```

## Dry Run

```bash
# See what request would be sent
glean search --dry-run "test query"
```

## Request Schema (components.SearchRequest)

```json
{
  "query": "string (required)",
  "pageSize": 10,
  "maxSnippetSize": 0,
  "timeoutMillis": 30000,
  "disableSpellcheck": false,
  "requestOptions": {
    "facetFilters": [{"fieldName": "datasource", "values": [{"value": "confluence", "relationType": "EQUALS"}]}],
    "responseHints": ["RESULTS", "QUERY_METADATA"],
    "facetBucketSize": 10,
    "timezoneOffset": 0
  }
}
```

## Common Patterns

```bash
# Search and extract document titles
glean search "vacation policy" | jq '[.results[].document.title]'

# Search with datasource filter via flag
glean search "deploy guide" -d confluence | jq '.results | length'

# Search and get URLs only
glean search "engineering handbook" | jq '[.results[].document.url]'
```

## Pitfalls

- Results may be empty if the token doesn't have access to the datasource
- `pageSize` is a hint — the server may return more or fewer results
- Use `--fields` to limit output size when results are large

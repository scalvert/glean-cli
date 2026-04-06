---
name: glean-cli-search
description: "Search across company knowledge with the Glean CLI. Use when finding documents, policies, engineering docs, or any information across enterprise data sources."
---

# glean search

> **PREREQUISITE:** Read `../glean-cli/SKILL.md` for auth, global flags, and security rules.

Search for content in your Glean instance. Results are JSON.

```bash
glean search [flags]
```

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--datasource` | []string |  | Filter by datasource (repeatable) |
| `--disable-query-autocorrect` | boolean | false | Disable automatic query corrections |
| `--disable-spellcheck` | boolean | false | Disable spellcheck |
| `--dry-run` | boolean | false | Print request body without sending |
| `--facet-bucket-size` | integer | 10 | Maximum facet buckets per result |
| `--fetch-all-datasource-counts` | boolean | false | Return counts for all datasources |
| `--json` | string |  | Complete JSON request body (overrides individual flags) |
| `--max-snippet-size` | integer | 0 | Maximum snippet size in characters |
| `--output` | json \| ndjson \| text | json | Output format |
| `--page-size` | integer | 10 | Number of results per page |
| `--query` | string |  | Search query (positional arg) **(required)** |
| `--query-overrides-facet-filters` | boolean | false | Allow query operators to override facet filters |
| `--response-hints` | []string | [RESULTS QUERY_METADATA] | Response hints |
| `--return-llm-content` | boolean | false | Return expanded LLM-friendly content |
| `--tab` | []string |  | Filter by result tab IDs (repeatable) |
| `--timeout` | integer | 30000 | Request timeout in milliseconds |
| `--type` | []string |  | Filter by document type (repeatable) |

## Examples

```bash
glean search "vacation policy" | jq '.results[].document.title'
glean search --json '{"query":"Q1 reports","pageSize":5,"datasources":["confluence"]}' | jq .
```

## Discovering Commands

```bash
# Show machine-readable schema for this command
glean schema search

# List all available commands
glean schema | jq '.commands'
```

# Discover what commands are available
glean schema | jq '.commands'

# Preview a request before sending
glean search --dry-run --datasource confluence "Q1 planning"

# Parse results with jq — each result has .title and nested .document fields
glean search "onboarding" | jq '.results[].title'

# Stream results as NDJSON — one SearchResult object per line
glean search "engineering docs" --output ndjson | jq .title

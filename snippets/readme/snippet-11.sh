# 1. Discover all available commands
glean schema | jq '.commands'

# 2. Understand a command's flags
glean schema search | jq '.flags | keys'
glean schema search | jq '.flags["--output"]'

# 3. Preview the exact request before sending
glean shortcuts create --dry-run \
  --json '{"data":{"inputAlias":"test","destinationUrl":"https://example.com"}}'

# 4. Execute and parse results
glean search "engineering values" | jq '.results[].title'

# 5. Stream NDJSON for large result sets — one SearchResult object per line
glean search "all docs" --output ndjson --page-size 50 | jq .title

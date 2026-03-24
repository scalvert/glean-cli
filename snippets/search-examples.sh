glean search "vacation policy"
glean search "Q1 planning" --datasource confluence --page-size 5
glean search "docs" --fields "results.document.title,results.document.url"
glean search "docs" --output ndjson | jq .title
glean search --json '{"query":"onboarding","pageSize":3}'
glean search --dry-run "test"

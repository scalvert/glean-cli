# Retrieve a document by URL
glean documents get --json '{"documentSpecs":[{"url":"https://..."}]}'

# Summarize a document
glean documents summarize --json '{"documentSpecs":[{"url":"https://..."}]}'

# Look up people
glean entities list --json '{"entityType":"PEOPLE","query":"engineering"}'

# Create a go-link
glean shortcuts create --json '{"data":{"inputAlias":"onboarding","destinationUrl":"https://..."}}'

# Create a shortcut with a variable template
glean shortcuts create --json '{"data":{"inputAlias":"jira","urlTemplate":"https://jira.example.com/browse/{arg}"}}'

# Pin a result for a query
glean pins create --json '{"queries":["onboarding"],"documentId":"https://..."}'

# List available AI agents
glean agents list | jq '.agents[] | {id: .agent_id, name: .name}'

# Get a specific agent
glean agents get --json '{"agentId":"<id>"}'

# Get schemas for an agent
glean agents schemas --json '{"agentId":"<id>"}'

# Run an agent
glean agents run --json '{"agentId":"<id>","messages":[{"author":"USER","fragments":[{"text":"summarize Q1 results"}]}]}'

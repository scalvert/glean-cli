glean chat "What are our company holidays?"
glean chat --timeout 120000 "Summarize all Q1 OKRs across teams"
glean chat --json '{"messages":[{"author":"USER","messageType":"CONTENT","fragments":[{"text":"What is Glean?"}]}]}'
glean chat --dry-run "test"
echo "What is Glean?" | glean chat
glean chat                                # interactive multiline input, Ctrl+D to send

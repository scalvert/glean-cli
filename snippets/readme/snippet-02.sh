# 1. Authenticate
glean auth login                          # OAuth via browser (recommended)
# — OR set env vars for CI/CD:
# export GLEAN_SERVER_URL=https://your-server-url GLEAN_API_TOKEN=your-token

# 2. Search
glean search "vacation policy"

# 3. Chat
glean chat "Summarize our Q1 engineering goals"

# 4. Open the interactive TUI
glean

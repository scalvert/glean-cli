# Glean CLI (Unofficial)

> Work seamlessly with Glean from your command line.

![Glean CLI Demo](demo/readme.gif)

The Glean CLI (`glean`) brings Glean's powerful search and AI capabilities directly to your terminal. Search across your company's knowledge, chat with Glean Assistant, and manage your configuration all from the comfort of your command line.

## Features

- 🔍 **Powerful Search**: Search across all your company's content with rich filtering options
- 💬 **Interactive Chat**: Have natural conversations with Glean's AI about your company's knowledge
- 🔐 **Secure Authentication**: Credentials are stored securely in your system's keyring
- 🎨 **Beautiful Output**: Rich, colorized output with support for custom formatting
- 🛠️ **API Access**: Direct access to Glean's REST API for power users

## Installation

```bash
# Using homebrew
brew install scalvert/tap/glean-cli

# Manual installation
curl -fsSL https://raw.githubusercontent.com/scalvert/glean-cli/main/install.sh | sh
```

## Quick Start

1. Configure your Glean credentials:
```bash
glean config --host your-company --token your-token
```

2. Search for content:
```bash
# Basic search
glean search "vacation policy"

# Search with filters
glean search --datasource confluence "engineering docs"

# Custom output format
glean search --output json "meeting notes" | jq
```

3. Chat with Glean Assistant:
```bash
# Start a conversation
glean chat "What are our company holidays?"

# Extended chat session
glean chat --timeout 60000 "Tell me about our engineering team"
```

## Commands

- `glean search`: Search across your company's content
  - Supports filtering by datasource, type, and people
  - Custom output formatting with templates
  - JSON output for scripting

- `glean chat`: Have conversations with Glean Assistant
  - Natural language interactions with your company's knowledge base
  - Streaming responses with markdown rendering
  - Configurable timeouts and response saving

- `glean api`: Direct access to Glean's REST API
  - Support for all HTTP methods
  - Request preview
  - Custom headers and authentication

- `glean config`: Manage your configuration
  - Secure credential storage
  - Multiple configuration options
  - Easy setup and updates

## Examples

### Advanced Search

```bash
# Search with multiple filters
glean search --datasource confluence,drive --type document "project planning"

# Custom output template
glean search --template "{{range .Results}}{{.Title}} - {{.URL}}\n{{end}}" "meeting notes"

# Search with person filter
glean search --person john@company.com "team updates"
```

### Interactive Chat

```bash
# Basic chat with Glean Assistant
glean chat "What's our remote work policy?"

# Extended timeout for longer responses
glean chat --timeout 60000 "Tell me about our engineering team"

# Disable saving chat history (enabled by default)
glean chat --save=false "Tell me about our tech stack"
```

### API Access

```bash
# Get user information
glean api users/me

# Custom search request
glean api search --method POST --raw-field '{"query": "engineering", "pageSize": 5}'

# Preview API request
glean api search --preview --method POST --raw-field '{"query": "docs"}'
```

## Configuration

The CLI stores configuration securely using your system's keyring with fallback to file-based storage:

```bash
# Set Glean instance
glean config --host your-company

# Set API token
glean config --token your-token

# Set user email
glean config --email you@company.com

# View current config
glean config --show

# Clear all settings
glean config --clear
```

## Contributing

Contributions are welcome! Please read our [Contributing Guide](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

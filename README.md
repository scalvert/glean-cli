# Glean CLI

> ⚠️ **Note:** This project is currently under active development and may not be ready for production use. APIs and features may change without notice.

A command-line interface for interacting with Glean's API and services. This CLI tool provides a seamless way to work with Glean from your terminal.

## Installation

```bash
go install github.com/scalvert/glean-cli@latest
```

## Usage

```bash
glean [command] [flags]
```

## Available Commands

### `glean config`

Configure your Glean CLI credentials and settings.

```bash
# Set Glean host
glean config --host your-domain
# or
glean config --host your-domain-be.glean.com

# Set Glean API token
glean config --token your-token

# Set Glean user email
glean config --email user@company.com

# Show current configuration
glean config --show

# Clear all stored credentials
glean config --clear
```

### `glean api`

Make direct calls to the Glean API endpoints.

```bash
# Make a GET request
glean api /search

# Make a POST request
glean api --method POST /users

# Make a custom request
glean api -X PUT /update
```

### `glean generate`

Generate various resources and code for Glean integration.

#### Subcommands:

##### `glean generate openapi-spec`

Generate OpenAPI specifications from API definitions or curl commands.

```bash
# Generate from a file
glean generate openapi-spec -f input.txt -o spec.yaml

# Generate from stdin
echo "curl example.com/api" | glean generate openapi-spec

# Add custom instructions
glean generate openapi-spec -f input.txt --prompt "Include rate limiting details"
```

Options:
- `-f, --file`: Input file containing the API/curl command
- `-o, --output`: Output file for the OpenAPI spec (defaults to stdout)
- `-p, --prompt`: Additional instructions for the LLM
- `--model`: LLM model to use (default: "gpt-4")

## Development

### Requirements

- Go 1.19 or higher
- Make (optional)

### Building from Source

```bash
git clone https://github.com/scalvert/glean-cli.git
cd glean-cli
go build
```

### Running Tests

```bash
go test ./... -v
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
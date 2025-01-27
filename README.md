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
glean api <endpoint>

# Make a POST request with a request body
glean api <endpoint> --method POST --raw-field '{"key": "value"}'

# Make a request with a different HTTP method
glean api <endpoint> --method PUT --raw-field '{"key": "value"}'
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

## API Command

The `api` command allows you to make authenticated HTTP requests to the Glean API. It handles authentication and provides a convenient way to interact with any of Glean's REST API endpoints documented at [developers.glean.com](https://developers.glean.com).

### Request Options

- `--method, -X`: Specify the HTTP method (default: GET)
- `--raw-field`: Add a JSON string as the request body
- `--input, -F`: Read request body from a file
- `--preview`: Preview the request without sending it
- `--no-color`: Disable colorized output (useful when piping to jq)
- `--raw`: Print raw API response without formatting

### Examples

```bash
# Preview a request
glean api <endpoint> --method POST --raw-field '{"key": "value"}' --preview

# Pipe results to jq
glean api <endpoint> --no-color | jq '.'

# Read request body from a file
glean api <endpoint> --method POST --input params.json

# Read request body from stdin
echo '{"key": "value"}' | glean api <endpoint> --method POST
```

For available endpoints and their parameters, please refer to the [Glean API Documentation](https://developers.glean.com).

All requests automatically include the necessary authentication headers and follow Glean's REST API conventions.

## Development

### Requirements

- Go 1.19 or higher
- golangci-lint ([install](https://golangci-lint.run/welcome/install/#local-installation))

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

### Code Quality

```bash
# Run linters
golangci-lint run

# Fix auto-fixable issues
golangci-lint run --fix
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

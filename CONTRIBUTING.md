# Contributing to Glean CLI

> Thank you for your interest in contributing to the Glean CLI! This document provides guidelines and instructions for contributing.

## Development Setup

### Prerequisites

- Go 1.19 or higher
- [Task](https://taskfile.dev/#/installation) (task runner)
- Git

### Getting Started

1. Fork and clone the repository:
```bash
git clone https://github.com/YOUR-USERNAME/glean-cli.git
cd glean-cli
```

2. Install development dependencies:
```bash
task setup
```

3. Build the project:
```bash
task build
```

## Development Workflow

1. Create a new branch for your changes:
```bash
git checkout -b feature/your-feature-name
```

2. Make your changes, following our coding standards and practices.

3. Run tests:
```bash
# Run tests with verbose output
task test

# Run tests with colorized summary
task test:summary

# Run all checks (lint + test + build)
task test:all
```

4. Run linters:
```bash
# Run linter
task lint

# Run linter with auto-fix
task lint:fix
```

5. Install locally to test your changes:
```bash
task install
```

6. Commit your changes using conventional commit messages:
```bash
git commit -m "feat: add new feature"
git commit -m "fix: resolve issue with X"
```

## Available Tasks

Run `task --list-all` to see all available tasks. Common tasks include:

- `task setup`: Install required development dependencies
- `task build`: Build the CLI
- `task test:all`: Run all checks (used in CI)
- `task lint`: Run linters
- `task lint:fix`: Run linters with auto-fix
- `task install`: Install the CLI locally
- `task clean`: Clean build artifacts

## Pull Request Process

1. Update documentation (README.md, code comments) if you're changing functionality.
2. Add tests for any new features.
3. Ensure all tests pass and linters are clean.
4. Push your changes and create a pull request.
5. Fill out the pull request template with all required information.

## Code Style

- Follow standard Go conventions and idioms
- Use meaningful variable and function names
- Add comments for public APIs and complex logic
- Keep functions focused and concise
- Write tests for new functionality

## Testing

- Write unit tests for new features
- Use table-driven tests where appropriate
- Mock external dependencies
- Aim for high test coverage of critical paths
- Test both success and error cases

## Documentation

- Update README.md for user-facing changes
- Add godoc comments for exported functions and types
- Include examples in documentation where helpful
- Keep documentation up to date with code changes

## Release Process

1. Update version numbers in relevant files
2. Update CHANGELOG.md following Keep a Changelog format
3. Create a new release tag
4. Build and publish new artifacts

## Getting Help

- Open an issue for bugs or feature requests
- Ask questions in pull requests
- Be respectful and constructive in discussions

## License

By contributing to Glean CLI, you agree that your contributions will be licensed under the MIT License.
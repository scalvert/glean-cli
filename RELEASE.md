# Release Process

## Setup

After cloning the repository, run the following command to install required dependencies:

```bash
task setup
```

This will install:
- `svu` - for semantic versioning and tag management
- `gotestsum` - for improved test output and summaries
- `golangci-lint` - for code linting

## Creating a Release

To create a new release:

1. Ensure your working directory is clean and you're on the `main` branch
2. Run the release task:
   ```bash
   task release
   ```

   This will:
   - Run all checks (lint, test, build)
   - Show the current version and commits since last release
   - Calculate and propose the next version
   - Create and push a git tag
   - Trigger the release workflow

### Optional Parameters

- Specify a version:
  ```bash
  task release VERSION=v1.2.3
  ```

- Force update an existing tag:
  ```bash
  task release VERSION=v1.2.3 FORCE=true
  ```

## Monitoring

After creating a release:
1. The GitHub Actions workflow will be triggered automatically
2. Monitor the release at: https://github.com/scalvert/glean-cli/actions
3. Once complete, the release will be available at: https://github.com/scalvert/glean-cli/releases

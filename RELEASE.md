# Release Process

This document outlines the process for creating new releases of the Glean CLI.

## Versioning

We follow [Semantic Versioning](https://semver.org/) (SemVer):
- MAJOR version (X.0.0) - incompatible API changes
- MINOR version (0.X.0) - add functionality in a backward compatible manner
- PATCH version (0.0.X) - backward compatible bug fixes

## Prerequisites

Install the required tools:
```bash
# Install svu for automated versioning
go install github.com/caarlos0/svu@latest
```

## Release Steps

You can create a new release using either the automated task or the manual steps below.

### Automated Release

```bash
# This will run all checks, create and push the tag
task release
```

### Manual Steps

If you need more control over the release process, you can follow these steps manually:

1. **Ensure Main Branch is Ready**
   ```bash
   git checkout main
   git pull origin main
   ```

2. **Run Tests and Checks**
   ```bash
   task test
   task lint
   ```

3. **Create Release**
   ```bash
   # Preview the next version based on conventional commits
   svu next

   # Create and push the tag with the next version
   git tag -a $(svu next) -m "Release $(svu next)"
   git push origin $(svu next)
   ```

   Note: `svu` automatically determines the next version based on your commit messages:
   - `feat:` -> MINOR version bump
   - `fix:` -> PATCH version bump
   - `BREAKING CHANGE:` in commit body -> MAJOR version bump

4. **Monitor Release Process**
   - The GitHub Action will automatically:
     - Create a GitHub release
     - Build binaries for all supported platforms
     - Generate release notes from commits
     - Upload artifacts to the release
     - Update the Homebrew formula

5. **Verify Release**
   - Check the [GitHub releases page](https://github.com/scalvert/glean-cli/releases)
   - Verify Homebrew formula was updated
   - Test installation methods:
     ```bash
     # Homebrew
     brew update
     brew install scalvert/tap/glean-cli

     # Go Install
     go install github.com/scalvert/glean-cli@latest

     # Shell Script
     curl -fsSL https://raw.githubusercontent.com/scalvert/glean-cli/main/install.sh | sh
     ```

## Release Artifacts

Each release includes:
- Binary distributions for:
  - macOS (x86_64, arm64)
  - Linux (x86_64, arm64)
  - Windows (x86_64, arm64)
- Source code (zip, tar.gz)
- Checksums file
- Changelog

## Commit Convention

To make the most of automated versioning, follow these commit message conventions:

```
<type>: <description>

[optional body]

[optional footer(s)]
```

Types:
- `feat:` - New feature (MINOR version bump)
- `fix:` - Bug fix (PATCH version bump)
- `docs:` - Documentation only changes
- `style:` - Changes that do not affect the meaning of the code
- `refactor:` - Code change that neither fixes a bug nor adds a feature
- `perf:` - Code change that improves performance
- `test:` - Adding missing tests or correcting existing tests
- `chore:` - Changes to the build process or auxiliary tools

Breaking Changes:
- Add `BREAKING CHANGE:` to the commit body to trigger a MAJOR version bump
- Example:
  ```
  feat: change API endpoint structure

  BREAKING CHANGE: The API endpoint structure has changed from /v1/* to /api/v1/*
  ```

## Troubleshooting

1. **GitHub Action Failed**
   - Check the Action logs for errors
   - Ensure GITHUB_TOKEN has required permissions
   - Verify the tag follows the format `v*` (e.g., v1.0.0)

2. **Homebrew Update Failed**
   - Check if homebrew-tap repository exists
   - Verify repository permissions
   - Check the "Update Homebrew tap" step in the Action logs

3. **Bad Release**
   If a release has issues:
   ```bash
   # Get current version
   current_version=$(svu current)

   # Delete tag locally
   git tag -d $current_version

   # Delete tag remotely
   git push --delete origin $current_version
   ```
   Then fix the issues and retry the release process.

## Notes

- Release notes are automatically generated from commit messages
- Commits starting with `docs:`, `test:`, `ci:` are excluded from release notes
- The release process is automated via GitHub Actions
- Binary distributions are built using GoReleaser
- All releases are automatically published to GitHub Releases

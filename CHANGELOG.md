# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.16.1](https://github.com/gleanwork/glean-cli/compare/v0.16.0...v0.16.1) (2026-05-02)


### Bug Fixes

* **ci:** mint App token for release-please so tags trigger release.yml ([#108](https://github.com/gleanwork/glean-cli/issues/108)) ([d9184f8](https://github.com/gleanwork/glean-cli/commit/d9184f8fe2f9ef19b055843da077e6b78645497c))

## [0.16.0](https://github.com/gleanwork/glean-cli/compare/v0.15.0...v0.16.0) (2026-05-02)


### Features

* **ci:** automate releases via release-please ([#106](https://github.com/gleanwork/glean-cli/issues/106)) ([fd4b3c0](https://github.com/gleanwork/glean-cli/commit/fd4b3c05344f51da195534e5d2868fef85e56ef4))


### Continuous Integration

* mint homebrew-tap token via GitHub App instead of PAT ([#105](https://github.com/gleanwork/glean-cli/issues/105)) ([adbcb9c](https://github.com/gleanwork/glean-cli/commit/adbcb9c1a02475da3a4022bbddb5dd581c220c8b))

## [0.5.5] - 2026-03-17

### Fixed
- CI: skip `TestStateDir_FilePermissions` on Windows — Unix permission bits are not enforced on Windows
- Remove `glean version` subcommand in favour of the conventional `--version` flag

## [0.5.4] - 2026-03-17

### Added
- Update notifications: after each command, a background goroutine checks the GitHub releases API (cached for 24h in `~/.glean/update-check.json`) and prints a notice to stderr when a newer version is available

## [0.5.3] - 2026-03-17

### Added
- `--version` flag on the root command (via Cobra built-in)

## [0.5.2] - 2026-03-17

### Fixed
- Add `--dry-run` flag to `documents get-permissions`, `answers get`, and `shortcuts get` — these were inconsistently missing the flag vs their sibling subcommands

## [0.5.1] - 2026-03-17

### Added
- `User-Agent: glean-cli/<version>` header on all outbound HTTP requests (both SDK-routed and streaming chat), allowing Glean's backend to identify and attribute CLI traffic by version

## [0.5.0] - 2026-03-17

### Added
- Full release pipeline: GoReleaser with cosign signing, CycloneDX SBOM, Homebrew tap publishing
- Checksum verification in `install.sh`
- `SECURITY.md` with vulnerability disclosure process
- `--version` / `--help` flag improvements

### Fixed
- Release workflow now gates GoReleaser on tests and lint passing
- `glean chat --json --dry-run` now correctly includes `stream: true`
- All delete/remove subcommands now support `--dry-run`
- `documents get-by-facets`, `entities read-people`, `messages get` now support `--dry-run`
- README Quick Start uses correct full hostname format and includes auth as step 0
- `glean chat --timeout` help text corrected to reflect 60s default
- Error messages across namespace commands now include `--help` guidance

## [0.4.0] - 2026-03-14

### Added
- Full-screen interactive TUI as the default `glean` invocation (no args)
  - Streaming chat with live stage indicators (Searching / Reading / Writing)
  - Slash commands: `/mode auto|fast|advanced`, `/clear`, `/help`
  - `@filename` file attachment support
  - Session persistence with `--continue` flag
  - `ctrl+y` to copy last response
- `glean mcp` stdio MCP server exposing `glean_search`, `glean_chat`, `glean_schema`, `glean_people`
- `--fields` dot-path projection for search and namespace commands
- Agent skill files in `skills/`

## [0.3.0] - 2026-03-13

### Added
- 18 SDK namespace command groups: `activity`, `agents`, `announcements`, `answers`, `collections`, `documents`, `entities`, `insights`, `messages`, `pins`, `shortcuts`, `tools`, `verification`, plus core `search`, `chat`, `api`, `auth`, `schema`
- `--json` raw payload flag on all namespace commands
- `--output json|ndjson|text` on all commands
- `--dry-run` on all mutating commands
- `glean schema [command]` for machine-readable flag documentation

## [0.2.x] - 2025-2026

### Added
- OAuth PKCE + Dynamic Client Registration (`glean auth login`)
- Official Glean Go SDK (`github.com/gleanwork/api-client-go`) replacing hand-rolled HTTP client
- Shell completions: `glean completion bash|zsh|fish`
- Cross-platform builds (macOS, Linux, Windows — amd64 and arm64)

## [0.1.0] - 2025

### Added
- Initial release: `glean search` and `glean chat` commands
- API token authentication via environment variables

# Release Process

Releases are automated by [release-please](https://github.com/googleapis/release-please). Conventional commits on `main` accumulate into a release PR; merging that PR cuts the tag and triggers the binary-build pipeline.

## Day-to-day flow

1. **Write conventional commits** when merging to `main`. Prefixes drive the CHANGELOG sections:

   | Prefix      | Section in CHANGELOG         | Triggers version bump? |
   | ----------- | ---------------------------- | ---------------------- |
   | `feat:`     | Features                     | minor (pre-1.0)        |
   | `fix:`      | Bug Fixes                    | patch                  |
   | `perf:`     | Performance Improvements     | patch                  |
   | `revert:`   | Reverts                      | patch                  |
   | `deps:`     | Dependencies                 | —                      |
   | `docs:`     | Documentation                | —                      |
   | `ci:`       | Continuous Integration       | —                      |
   | `refactor:` | Code Refactoring             | —                      |
   | `test:`     | (hidden)                     | —                      |
   | `build:`    | (hidden)                     | —                      |
   | `chore:`    | (hidden)                     | —                      |

   Add `!` after the type (e.g. `feat!:`) or include `BREAKING CHANGE:` in the commit body to bump minor instead of patch while pre-1.0.

2. **Wait for the release PR**. release-please watches every push to `main` and opens or updates a PR titled `chore(main): release vX.Y.Z` containing the CHANGELOG.md diff and a version bump.

3. **Merge the release PR** when you're ready to ship. release-please then:
   - Creates the `vX.Y.Z` git tag
   - Creates the GitHub Release with the CHANGELOG entry as the body
   - The tag push triggers `.github/workflows/release.yml`, which runs GoReleaser to build binaries, sign them, upload SBOMs/checksums, and bump the `gleanwork/homebrew-tap` formula

Nothing to run locally.

## Monitoring

- Release PRs: https://github.com/gleanwork/glean-cli/pulls?q=is%3Apr+author%3Aapp%2Frelease-please
- Workflow runs: https://github.com/gleanwork/glean-cli/actions
- Published releases: https://github.com/gleanwork/glean-cli/releases

## Breaking glass

`mise run release` still exists as a manual fallback for cutting tags directly when the release-please flow is unavailable (tool outage, config bug, etc.). Prefer the automated flow.

```bash
mise run release                  # svu-computed next version
VERSION=v1.2.3 mise run release   # explicit version
```

## Setup (one-time, per clone)

```bash
mise run setup
```

Installs `svu`, `gotestsum`, `golangci-lint` — still needed locally for the pre-push checks required by `CLAUDE.md`.

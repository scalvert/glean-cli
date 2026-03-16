# Eval Checklist — glean-cli Beta Readiness

Generated: 2026-03-16 | Panel: beta-gate (4-agent — release-infra, security, api-correctness, ux-onboarding)

## How to read this
- P0 = beta blocker — do not ship without fixing
- P1 = should fix before beta tag; acceptable risk if documented
- P2 = nice to have for beta; defer to GA
- Items are OPEN until acceptance criteria are verifiably met

---

## P0 — Beta Blockers

- [ ] **CHK-B001** Release workflow runs GoReleaser without first running tests or lint
  - **File:** `.github/workflows/release.yml:12-34`
  - **Acceptance:** `goreleaser` job has `needs: [build, lint]` (or equivalent) so broken code cannot ship
  - **Agents:** release-infra

- [ ] **CHK-B002** `chat --json --dry-run` does not include `stream: true` in output
  - **File:** `cmd/chat.go:74` (stream assignment happens after the dryRun check at line 73)
  - **Acceptance:** `glean chat --dry-run --json '{"messages":[...]}'` returns JSON with `"stream": true`
  - **Agents:** api-correctness

- [ ] **CHK-B003** Delete/remove subcommands missing `--dry-run` flag — agents cannot preview destructive operations
  - **Files:** `cmd/shortcuts.go` (delete), `cmd/pins.go` (remove), `cmd/announcements.go` (delete), `cmd/answers.go` (delete), `cmd/collections.go` (delete, delete-item)
  - **Acceptance:** All delete/remove subcommands accept `--dry-run` and print the request body without executing
  - **Agents:** api-correctness

- [ ] **CHK-B004** `documents get-by-facets`, `entities read-people`, `messages get` missing `--dry-run`
  - **Files:** `cmd/documents.go:68`, `cmd/entities.go:70`, `cmd/messages.go:30`
  - **Acceptance:** All three accept `--dry-run` and print request body without executing
  - **Agents:** api-correctness

- [ ] **CHK-B005** README Quick Start host config mismatch — docs imply short-name auto-append (`your-company`), code does not auto-append
  - **File:** `README.md:22-26`
  - **Acceptance:** Quick Start uses only full hostname format (`your-company-be.glean.com`) OR code implements auto-append; both docs and code match
  - **Agents:** ux-onboarding

- [ ] **CHK-B006** `glean version` outputs `"glean version dev"` — released binaries must show semver
  - **File:** `cmd/version.go`, `.goreleaser.yml:17`
  - **Acceptance:** A binary built by GoReleaser shows `"glean version v0.x.y-beta.N"`, not "dev"
  - **Agents:** release-infra, api-correctness, ux-onboarding (3-agent convergence)

---

## P1 — Should Fix Before Beta Tag

- [ ] **CHK-B007** install.sh has no checksum/signature verification — downloads binary and extracts without integrity check
  - **File:** `install.sh:35-38`
  - **Acceptance:** install.sh downloads checksum file alongside binary and verifies before extraction; OR README removes curl-pipe-to-sh as recommended install path in favor of Homebrew
  - **Agents:** release-infra, security

- [ ] **CHK-B008** Namespace command error messages lack context — `"--json is required"` with no example or schema hint
  - **Files:** `cmd/shortcuts.go:45`, `cmd/entities.go:40`, similar in namespace commands
  - **Acceptance:** Error messages include one-line example payload or direct user to `--help`
  - **Agents:** ux-onboarding

- [ ] **CHK-B009** Quick Start does not include authentication as step 0
  - **File:** `README.md:19-30`
  - **Acceptance:** Quick Start begins with `glean auth login` or `glean config --token ... --host ...` before attempting any search/chat command
  - **Agents:** ux-onboarding

- [ ] **CHK-B010** `entities list` error check couples to Go internal type name (`ListEntitiesRequestEntityType`)
  - **File:** `cmd/entities.go:45`
  - **Acceptance:** Error message does not contain SDK type names; shows only user-visible valid values
  - **Agents:** api-correctness

- [ ] **CHK-B011** `chat --timeout` help text says "default 30s" but default is 60000ms (60s)
  - **File:** `cmd/chat.go:108`
  - **Acceptance:** Flag description reads `"Request timeout in milliseconds (default 60000 — 60 seconds)"`
  - **Agents:** ux-onboarding

---

## P2 — Nice to Have

- [ ] **CHK-B012** No binary GPG signing or cosign provenance in GoReleaser
  - **File:** `.goreleaser.yml`
  - **Acceptance:** checksums.txt is signed with GPG or cosign; release notes document how to verify
  - **Agents:** release-infra, security

- [ ] **CHK-B013** No SECURITY.md with vulnerability disclosure process
  - **Acceptance:** `SECURITY.md` exists with responsible disclosure contact and timeline
  - **Agents:** security

- [ ] **CHK-B014** SBOM not included in release artifacts
  - **Acceptance:** GoReleaser generates CycloneDX SBOM and includes in release
  - **Agents:** release-infra, security

- [ ] **CHK-B015** `glean config --show` doesn't clarify env var precedence
  - **Acceptance:** Output notes that env vars (`GLEAN_API_TOKEN`, `GLEAN_HOST`) override stored config
  - **Agents:** ux-onboarding

- [ ] **CHK-B016** TUI welcome screen slash command hint incomplete (says "Type / for commands" but only 3 exist)
  - **File:** `internal/tui/view.go:109`
  - **Acceptance:** Welcome hint either lists the 3 commands explicitly or is reworded to match actual scope
  - **Agents:** ux-onboarding

---

## Closed

*(items from prior panels — all closed before this beta gate panel)*

See git log for closure evidence on CHK-001 through CHK-038 from the March 13-14 panel.

---

## Panel Scores

| Agent | Lens | Score | Key Finding |
|-------|------|-------|-------------|
| release-infra | Release infra | 7/10 | Release workflow bypasses CI before GoReleaser |
| security | Security | 8/10 | No binary signing; otherwise solid OAuth + storage |
| ux-onboarding | UX / Onboarding | 6.5/10 | README host mismatch breaks first user step |
| api-correctness | API correctness | 5/10 | Missing --dry-run on delete ops; chat --json --dry-run broken |
| test-coverage | Test coverage | (no report) | — |
| devils-advocate | Devil's advocate | (no report) | — |

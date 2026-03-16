# API Fix + Test Coverage Tracker
**Started:** 2026-03-16 | **Status:** IN PROGRESS

This is the single source of truth for all bug fixes and test coverage work.
An item is only DONE when: fix committed + tests passing + verified.

---

## Bug Fixes

| ID | Description | File(s) | Status | Commit |
|----|-------------|---------|--------|--------|
| BUG-001 | `glean api` doesn't inject OAuth token → 401 on all calls | cmd/api.go | ⬜ TODO | — |
| BUG-002 | `chat --json` content-type error on 200 response | cmd/chat.go | ⬜ TODO | — |
| BUG-003a | `documents get` missing --dry-run | cmd/documents.go | ⬜ TODO | — |
| BUG-003b | `entities list` missing --dry-run | cmd/entities.go | ⬜ TODO | — |
| BUG-003c | `messages get` missing --dry-run | cmd/messages.go | ⬜ TODO | — |
| BUG-003d | `answers list` missing --dry-run | cmd/answers.go | ⬜ TODO | — |
| BUG-004 | Dry-run silently drops fields (activity, agents, insights, verification, pins, shortcuts) | multiple | ⬜ TODO | — |
| BUG-005a | `announcements list` subcommand missing (silent exit 0) | cmd/announcements.go | ⬜ TODO | — |
| BUG-005b | `collections list` subcommand missing (silent exit 0) | cmd/collections.go | ⬜ TODO | — |
| BUG-006 | `announcements create --dry-run` crashes — JSON parse error | cmd/announcements.go | ⬜ TODO | — |
| BUG-007 | `tools run` parameters field type error — no schema guidance | cmd/tools.go | ⬜ TODO | — |
| BUG-008 | `entities list` opaque enum error — doesn't show valid values | cmd/entities.go | ⬜ TODO | — |
| BUG-009 | `--fields` help example uses wrong path (`document.title` → `results.document.title`) | cmd/search.go | ⬜ TODO | — |
| BUG-010 | NDJSON emits whole response as one line instead of per-result | internal/output/formatter.go | ⬜ TODO | — |
| BUG-011 | `config --show` text-only, not JSON parseable | cmd/config.go | ⬜ TODO | — |
| BUG-012 | `documents summarize --dry-run` shows `documentSpecs:null` not input | cmd/documents.go | ⬜ TODO | — |

## Uniform Test Coverage

Every command must pass the same 6-test matrix:
1. `--help` exits 0, outputs usage
2. `glean schema <cmd>` exits 0, valid JSON
3. `--dry-run` exits 0, valid JSON
4. Live call (mocked) exits 0, valid JSON response
5. Unknown subcommand exits non-zero
6. Invalid `--json` exits non-zero

| Command | Tests Exist | Tests Pass | Status |
|---------|-------------|------------|--------|
| search | ✅ partial | — | ⬜ NEEDS EXPANSION |
| chat | ✅ partial | — | ⬜ NEEDS EXPANSION |
| api | ✅ partial | — | ⬜ NEEDS EXPANSION |
| config | ✅ partial | — | ⬜ NEEDS EXPANSION |
| activity | ❌ none | — | ⬜ TODO |
| agents | ❌ none | — | ⬜ TODO |
| announcements | ❌ none | — | ⬜ TODO |
| answers | ❌ none | — | ⬜ TODO |
| collections | ❌ none | — | ⬜ TODO |
| documents | ❌ none | — | ⬜ TODO |
| entities | ❌ none | — | ⬜ TODO |
| insights | ❌ none | — | ⬜ TODO |
| messages | ❌ none | — | ⬜ TODO |
| pins | ❌ none | — | ⬜ TODO |
| shortcuts | ❌ none | — | ⬜ TODO |
| tools | ❌ none | — | ⬜ TODO |
| verification | ❌ none | — | ⬜ TODO |

## Final Verification

- [ ] `go build ./...` — clean
- [ ] `go test ./...` — all pass
- [ ] `glean api users/me` — returns valid JSON (not 401)
- [ ] `glean chat --json '...'` — returns response (not content-type error)
- [ ] `glean schema` — lists all 18+ commands
- [ ] All namespace commands have passing test suite

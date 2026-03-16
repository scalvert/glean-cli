# Glean CLI — Agent Readiness Evaluation
**Date:** 2026-03-16
**Instance:** scio-prod-be.glean.com
**Auth:** OAuth token (expires 2026-03-21)
**Method:** 5 parallel test agents, 83 total test cases

---

## Executive Summary

The CLI has an excellent discovery layer (schema, help, dry-run on core commands) but has several bugs that make it unreliable for agent-first use. Two critical bugs block core functionality. Eight high-severity issues degrade agent usability across the namespace commands.

| Severity | Count | Description |
|----------|-------|-------------|
| CRITICAL  | 2 | Blocks core agent workflows entirely |
| HIGH      | 7 | Significantly impairs reliable agent use |
| MEDIUM    | 4 | Degrades UX, requires workarounds |
| INFO      | 3 | Auth/permissions — likely not code bugs |

---

## Overall Status by Command

| Command | Schema | --help | --dry-run | Live API | --json | NDJSON | Agent-Ready |
|---------|--------|--------|-----------|----------|--------|--------|-------------|
| **search** | ✅ | ✅ | ✅ | ✅ | ✅ | ⚠️ | **YES** |
| **chat** | ✅ | ✅ | ✅ | ✅ | ❌ | N/A | **PARTIAL** |
| **api** | ✅ | ✅ | N/A | ❌ | ❌ | N/A | **NO** |
| **schema** | N/A | ✅ | N/A | ✅ | N/A | N/A | **YES** |
| **auth** | N/A | ✅ | N/A | ✅ | N/A | N/A | **YES** |
| **config** | ✅ | ✅ | N/A | ⚠️ | N/A | N/A | **PARTIAL** |
| **version** | ✅ | ✅ | N/A | ✅ | N/A | N/A | **YES** |
| **documents** | ✅ | ✅ | ❌ | ❌ 401 | ⚠️ | N/A | **NO** |
| **entities** | ✅ | ✅ | ❌ | ❌ | N/A | N/A | **NO** |
| **messages** | ✅ | ✅ | ❌ | ❌ 401 | N/A | N/A | **NO** |
| **answers** | ✅ | ✅ | ❌ | ❌ 401 | N/A | N/A | **NO** |
| **announcements** | ✅ | ✅ | ❌ | ❌ | ❌ | N/A | **NO** |
| **collections** | ✅ | ✅ | ✅ | ❌ 401 | N/A | N/A | **PARTIAL** |
| **pins** | ✅ | ✅ | ⚠️ | ❌ 401 | N/A | N/A | **NO** |
| **shortcuts** | ✅ | ✅ | ⚠️ | ❌ 401 | N/A | N/A | **NO** |
| **activity** | ✅ | ✅ | ⚠️ | N/A | N/A | N/A | **PARTIAL** |
| **agents** | ✅ | ✅ | ⚠️ | ❌ 401 | N/A | N/A | **NO** |
| **insights** | ✅ | ✅ | ⚠️ | ❌ 401 | N/A | N/A | **NO** |
| **tools** | ✅ | ✅ | ❌ | ❌ 401 | N/A | N/A | **NO** |
| **verification** | ✅ | ✅ | ⚠️ | ❌ 401 | N/A | N/A | **NO** |

Legend: ✅ Works | ⚠️ Partial/Degraded | ❌ Fails | N/A Not applicable

---

## Critical Bugs

### BUG-001 — `glean api` does not inject Bearer token [CRITICAL]

**Command:** `glean api`
**Symptom:** Every request fails with `Error: API error (401): Bearer authorization not provided.`
**Evidence:**
```
$ ./glean api search --method POST --raw-field '{"query":"test","pageSize":1}'
Error: API error (401): Bearer authorization not provided. Not allowed
```
The `search` command works fine with the same credentials, confirming the token is available — the `api` command's HTTP client is simply not setting the `Authorization: Bearer <token>` header.

**Impact:** The `api` command is the agent's escape hatch for any endpoint not covered by dedicated commands. It is completely non-functional.

**Also:** `glean api users/me` returns 404, but this may be an invalid path (masked by the auth bug).

---

### BUG-002 — `chat --json` fails with content-type error on successful response [CRITICAL]

**Command:** `glean chat --json '{"messages":[...]}'`
**Symptom:** Exit code 1 despite a 200 HTTP response containing the correct answer.
```
Error: chat request failed: unknown content-type received: application/json: Status 200
```
The raw JSON response is dumped to stderr but the command reports failure.
**Root cause:** When using `--json` with an explicit message body, the streaming parser expects `text/event-stream` but the API returns `application/json` (non-streaming). The `--json` path likely forces `stream:true` in the request while setting an incompatible `Accept` header.

**Impact:** Agents cannot use structured multi-turn chat requests or custom agent configs via the CLI. Only the simple positional-arg form works.

---

## High Severity Issues

### BUG-003 — `--dry-run` absent on namespace commands despite schema advertising it [HIGH]

Schema reports `--dry-run` as a valid flag for these commands, but all reject it:
- `glean documents get --dry-run` → `Error: unknown flag: --dry-run`
- `glean entities list --dry-run` → `Error: unknown flag: --dry-run`
- `glean messages get --dry-run` → `Error: unknown flag: --dry-run`
- `glean answers list --dry-run` → `Error: unknown flag: --dry-run`

`glean tools run --dry-run` → `Error: invalid --json` (different failure mode)

**Impact:** Agents relying on schema introspection to discover capabilities will believe dry-run works everywhere, then get hard errors. Schema/implementation mismatch is the worst kind of documentation lie.

---

### BUG-004 — Dry-run silently drops/mangles input fields across ops commands [HIGH]

For these commands, `--dry-run` exits 0 but the output shows empty or incorrect data — the input JSON fields were silently discarded because the SDK struct field names differ from the intuitive JSON names:

| Command | Input | Dry-run Output | Lost Fields |
|---------|-------|----------------|-------------|
| `activity report` | `{"events":[{"action":"VIEW","docId":{...}}]}` | `{"events":[{"action":"VIEW","timestamp":"0001-01-01T00:00:00Z","url":""}]}` | `docId` → wrong fields |
| `agents run` | `{"agentId":"default","query":"test"}` | `{"agent_id":""}` | `agentId`, `query` both dropped |
| `insights get` | `{"insightType":"SEARCH"}` | `{}` | All fields dropped |
| `verification remind` | `{"docId":{"datasource":"...","objectId":"..."}}` | `{"documentId":""}` | `docId` → wrong name |
| `pins create` | `{"query":"...","url":"..."}` | `{}` | All fields dropped |
| `shortcuts create` | `{"urlTemplate":"...","shortcutId":"..."}` | `{"data":{}}` | All fields dropped |

**Impact:** Agent sends a request, gets exit 0, believes it worked. The actual API call sends garbage or an empty body. No warning, no error.

---

### BUG-005 — `announcements list` and `collections list` don't exist, exit 0 [HIGH]

Both commands are referenced in the `--help` examples but are not registered subcommands. Running them silently prints the parent help and exits 0:

```
$ ./glean announcements list
# prints parent announcements help, exit 0
$ ./glean collections list
# prints parent collections help, exit 0
```

**Impact:** Silent success with no output is the worst failure mode for agents. An agent would believe the call succeeded and move on with no data.

---

### BUG-006 — `announcements create --dry-run` crashes [HIGH]

```
$ ./glean announcements create --dry-run --json '{"title":"Test","body":"Test body",...}'
Error: invalid --json: json: cannot unmarshal string into Go value of type map[string]json.RawMessage
```
The JSON parser for this command expects `map[string]json.RawMessage` instead of a flat object.

---

### BUG-007 — `tools run` parameters field rejects valid JSON [HIGH]

```
$ ./glean tools run --json '{"toolName":"glean_search","parameters":{"query":"test"}}'
Error: invalid --json: json: cannot unmarshal string into Go struct field
ToolsCallRequest.parameters of type components.ToolsCallParameter
```
The `parameters` field expects a specific Go struct type, not a JSON map. The schema and help provide no guidance on the expected structure. An agent cannot construct a valid call without knowing the `ToolsCallParameter` schema.

---

### BUG-008 — `entities list` rejects valid enum value with opaque error [HIGH]

```
$ ./glean entities list --json '{"entityType":"PERSON","query":"steve"}'
Error: invalid --json: invalid value for ListEntitiesRequestEntityType: PERSON
```
The error names an internal Go type (`ListEntitiesRequestEntityType`) but does not list valid values. An agent cannot self-correct without knowing what values are accepted.

---

## Medium Severity Issues

### BUG-009 — `--fields` projection path format wrong in documentation [MEDIUM]

Help text example says `--fields "document.title,document.url"` but this produces `{}`. The correct path is `--fields "results.document.title,results.document.url"` (results must be prefixed).

---

### BUG-010 — NDJSON emits entire response as one line, not per-result [MEDIUM]

`--output ndjson` on search outputs the full response envelope as a single JSON line. The expected behavior is one line per result, enabling streaming/incremental processing. Instead, agents get one massive line containing the full response.

---

### BUG-011 — `config --show` output is text-only, not parseable [MEDIUM]

```
Current configuration:
  Host:  scio-prod-be.glean.com
  Token: [not set]
```
No `--output json` option. Agents cannot reliably parse this.

---

### BUG-012 — `documents summarize --dry-run` output doesn't reflect input [MEDIUM]

Input: `{"docId":{"datasource":"confluence","objectId":"123"}}`
Output: `{"documentSpecs": null}`
The dry-run maps `docId` to a different internal field name without warning.

---

## Auth/Permissions Issues (Not Code Bugs)

The following commands return HTTP 401 on live calls. This appears to be an OAuth token scope issue — the token is valid for search/chat but lacks permissions for these APIs:

- `glean agents list` — 401
- `glean tools list` — 401
- `glean insights get` — 401
- `glean verification list` — 401
- `glean pins list` — 401
- `glean shortcuts list` — 401
- `glean answers list` — 401
- `glean documents get` — 401

Error format is clean: `Error: API error occurred: Status 401 / Not allowed` (exit 1) — well-behaved for agent error detection.

---

## What Works Well

### Discovery Layer — EXCELLENT

The schema system is genuinely impressive for agent use:

```bash
$ glean schema              # lists all 18 commands as JSON
$ glean schema search       # full JSON schema with flag types, defaults, enums, examples
$ glean schema shortcuts    # subcommand listing, required flags, output format enum
```

All 18 commands return valid, machine-parseable JSON schema. This is exactly what an LLM needs for tool-use decisions.

### Search — FULLY AGENT-READY

- Live JSON output: valid, structured, rich metadata
- `--dry-run`: accurate, shows full request including facet filters
- `--datasource` / `--type` filtering: works
- `--json` raw payload: works
- `--output json`: valid, parseable
- `--output ndjson`: technically valid (but single line — see BUG-010)
- `--fields` projection: works with correct path prefix

### Chat (positional arg form) — WORKS

```bash
$ glean chat "What is Glean?"
```
Returns clean text to stdout. Easily captured by agents. No flags needed.

### Auth Preflight — WORKS

```bash
$ glean auth status    # exit 0 = authenticated, includes expiry
```
Agents can run this before any other command to verify credentials.

### Help Everywhere — CONSISTENT

Every command returns clean, readable help with flag types, defaults, and examples. Subcommand structure is logical and discoverable.

---

## Fix Priority

| Priority | Bug | Fix |
|----------|-----|-----|
| P0 | BUG-001: `api` missing Bearer token | Inject auth header in api command HTTP client |
| P0 | BUG-002: `chat --json` content-type error | Set correct `Accept` header or handle `application/json` response |
| P1 | BUG-003: `--dry-run` missing on namespace commands | Add `--dry-run` flag to all namespace subcommands |
| P1 | BUG-004: Dry-run drops input fields silently | Map SDK struct field names to expected JSON names; warn on unknown input fields |
| P1 | BUG-005: Missing `list` subcommands exit 0 | Register `announcements list` and `collections list`, or remove from help examples |
| P1 | BUG-006: `announcements create` JSON parse crash | Fix JSON parsing to accept flat object |
| P1 | BUG-007: `tools run` parameters type error | Document `ToolsCallParameter` shape in schema/help; accept map input |
| P1 | BUG-008: `entities list` opaque enum error | Add valid enum values to error message and schema |
| P2 | BUG-009: `--fields` wrong path in docs | Fix help example to use correct `results.` prefix |
| P2 | BUG-010: NDJSON single line | Emit one result per line in NDJSON mode |
| P2 | BUG-011: `config --show` text-only | Add `--output json` to config show |
| P2 | BUG-012: `summarize` dry-run mismatch | Fix field name mapping in dry-run output |

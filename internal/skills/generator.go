// Package skills generates SKILL.md files from the CLI's own schema registry.
//
// Usage: glean generate-skills [--output-dir skills/]
package skills

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/gleanwork/glean-cli/internal/schema"
)

// subcommandMap provides human-friendly descriptions for subcommands that
// the schema registry doesn't capture (schemas are per-namespace, not per-sub).
// Keys are "namespace.subcommand".
var subcommandMap = map[string]string{
	"agents.list":               "List all available agents",
	"agents.get":                "Get details of a specific agent",
	"agents.schemas":            "Get input/output schemas for an agent",
	"agents.run":                "Run an agent",
	"answers.list":              "List all curated answers",
	"answers.get":               "Get a specific answer",
	"answers.create":            "Create a new answer",
	"answers.update":            "Update an existing answer",
	"answers.delete":            "Delete an answer",
	"announcements.create":      "Create a new announcement",
	"announcements.update":      "Update an existing announcement",
	"announcements.delete":      "Delete an announcement",
	"collections.list":          "List all collections",
	"collections.get":           "Get a specific collection",
	"collections.create":        "Create a new collection",
	"collections.update":        "Update a collection",
	"collections.delete":        "Delete a collection",
	"collections.add-items":     "Add documents to a collection",
	"collections.delete-item":   "Remove a document from a collection",
	"documents.get":             "Retrieve document metadata by URL or ID",
	"documents.summarize":       "Generate an AI summary of a document",
	"documents.get-by-facets":   "Retrieve documents matching facet filters",
	"documents.get-permissions": "Inspect who has access to a document",
	"entities.list":             "List entities by type and query",
	"entities.read-people":      "Get detailed people profiles",
	"insights.get":              "Get analytics data",
	"messages.get":              "Retrieve a specific message",
	"pins.list":                 "List all pins",
	"pins.get":                  "Get a specific pin",
	"pins.create":               "Create a new pin",
	"pins.update":               "Update a pin",
	"pins.remove":               "Remove a pin",
	"shortcuts.list":            "List all shortcuts",
	"shortcuts.get":             "Get a specific shortcut",
	"shortcuts.create":          "Create a new shortcut",
	"shortcuts.update":          "Update an existing shortcut",
	"shortcuts.delete":          "Delete a shortcut",
	"tools.list":                "List available platform tools",
	"tools.run":                 "Execute a platform tool",
	"verification.list":         "List documents pending verification",
	"verification.verify":       "Mark a document as verified",
	"verification.remind":       "Send a verification reminder",
	"activity.report":           "Report a user activity event",
	"activity.feedback":         "Submit feedback on search results",
}

// skillDescription maps command names to richer descriptions for the SKILL.md
// frontmatter. These are tuned for agent skill triggering.
var skillDescription = map[string]string{
	"search":        "Search across company knowledge with the Glean CLI. Use when finding documents, policies, engineering docs, or any information across enterprise data sources.",
	"chat":          "Chat with Glean Assistant from the command line. Use when asking questions, summarizing documents, or getting AI-powered answers about company knowledge.",
	"api":           "Make raw authenticated HTTP requests to any Glean REST API endpoint. Use when no dedicated command exists or for advanced API access.",
	"agents":        "List, inspect, and run Glean AI agents. Use when discovering available agents, viewing agent schemas, or invoking agents programmatically.",
	"documents":     "Retrieve, summarize, and inspect documents indexed by Glean. Use when getting document content, summaries, permissions, or metadata by URL.",
	"collections":   "Manage curated document collections in Glean. Use when creating, updating, or organizing themed sets of documents.",
	"entities":      "Look up people, teams, and custom entities in Glean. Use when finding employees, org structure, team members, or expertise.",
	"answers":       "Manage curated Q&A pairs in Glean. Use when creating, updating, or listing company-approved answers to common questions.",
	"shortcuts":     "Manage Glean go-links (shortcuts). Use when creating, listing, updating, or deleting short URL aliases like go/wiki.",
	"pins":          "Manage promoted search results (pins) in Glean. Use when pinning specific documents to appear first for certain queries.",
	"announcements": "Manage time-bounded company announcements in Glean. Use when creating, updating, or deleting announcements that surface across the Glean UI.",
	"activity":      "Report user activity and submit feedback to Glean. Use when logging user interactions or providing relevance feedback on search results.",
	"verification":  "Manage document verification and review workflows in Glean. Use when verifying document accuracy, listing pending verifications, or sending review reminders.",
	"tools":         "List and run Glean platform tools. Use when discovering available platform tools or executing them programmatically.",
	"messages":      "Retrieve indexed messages from Slack, Teams, and other messaging platforms via Glean. Use when searching for or reading specific messages.",
	"insights":      "Retrieve search and usage analytics from Glean. Use when analyzing search patterns, popular queries, or platform adoption metrics.",
}

// FlagInfo holds rendered flag data for templates.
type FlagInfo struct {
	Name        string
	Type        string
	Default     string
	Description string
	Required    bool
}

// SkillData holds all data needed to render a SKILL.md template.
type SkillData struct {
	Name        string
	Description string
	Command     string
	SchemaDesc  string
	Flags       []FlagInfo
	Subcommands []SubcommandInfo
	Example     string
	HasDryRun   bool
}

// SubcommandInfo describes a subcommand.
type SubcommandInfo struct {
	Name        string
	Description string
}

var skillTmpl = template.Must(template.New("skill").Parse(`---
name: glean-{{ .Command }}
description: "{{ .Description }}"
---

# glean {{ .Command }}

> **PREREQUISITE:** Read ` + "`../glean-shared/SKILL.md`" + ` for auth, global flags, and security rules.

{{ .SchemaDesc }}

` + "```bash" + `
glean {{ .Command }}{{ if .Subcommands }} <subcommand>{{ end }} [flags]
` + "```" + `
{{ if .Subcommands }}
## Subcommands

| Subcommand | Description |
|------------|-------------|
{{ range .Subcommands -}}
| ` + "`{{ .Name }}`" + ` | {{ .Description }} |
{{ end }}{{ end }}
## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
{{ range .Flags -}}
| ` + "`{{ .Name }}`" + ` | {{ .Type }} | {{ .Default }} | {{ .Description }}{{ if .Required }} **(required)**{{ end }} |
{{ end }}
## Examples

` + "```bash" + `
{{ .Example }}
` + "```" + `

## Discovering Commands

` + "```bash" + `
# Show machine-readable schema for this command
glean schema {{ .Command }}

# List all available commands
glean schema | jq '.commands'
` + "```" + `
`))

var sharedTmpl = template.Must(template.New("shared").Parse(`---
name: glean-shared
description: "Glean CLI: Shared patterns for authentication, global flags, output formatting, and security rules."
compatibility: Requires the glean binary on $PATH. Install via brew install gleanwork/tap/glean-cli
---

# glean — Shared Reference

> **Read this first.** All other glean skills assume familiarity with auth, flags, and output formats described here.

## Installation

` + "```bash" + `
brew install gleanwork/tap/glean-cli
` + "```" + `

## Authentication

` + "```bash" + `
# Browser-based OAuth (interactive — recommended)
glean auth login

# Verify credentials
glean auth status

# CI/scripting (no interactive setup needed)
export GLEAN_API_TOKEN=your-token
export GLEAN_HOST=your-company-be.glean.com
` + "```" + `

Credentials resolve in this order: environment variables → system keyring → ~/.glean/config.json.

## CLI Syntax

` + "```bash" + `
glean <command> [subcommand] [flags]
` + "```" + `

### Global Flags

| Flag | Description |
|------|-------------|
| --output <FORMAT> | json (default), ndjson (one result per line), text |
| --fields <PATHS> | Dot-path field projection (e.g. results.document.title,results.document.url) |
| --json <PAYLOAD> | Complete JSON request body (overrides all other flags) |
| --dry-run | Print request body without sending |

## Schema Introspection

Always call glean schema <command> before invoking a command you haven't used before.

` + "```bash" + `
glean schema | jq '.commands'          # list all commands
glean schema search | jq '.flags'      # flags for search
` + "```" + `

## Security Rules

- **Never** output API tokens or secrets directly
- **Always** use --dry-run before write/delete operations in automated pipelines
- Prefer environment variables over config files for CI/CD

## Error Handling

All errors go to stderr; stdout contains only structured output.
Exit code 0 = success, non-zero = error.

## Available Commands

| Command | Description |
|---------|-------------|
{{ range .Commands -}}
| [glean {{ .Name }}](../glean-{{ .Name }}/SKILL.md) | {{ .Description }} |
{{ end }}
`))

// CommandEntry is used by the shared skill template.
type CommandEntry struct {
	Name        string
	Description string
}

// skipCommands lists commands that should not get their own skill.
var skipCommands = map[string]bool{
	"version": true,
}

// Generate writes SKILL.md files to outputDir for all registered schemas.
func Generate(outputDir string) error {
	// Generate shared skill
	commands := schema.List()
	var entries []CommandEntry
	for _, name := range commands {
		if skipCommands[name] {
			continue
		}
		s, err := schema.Get(name)
		if err != nil {
			continue
		}
		entries = append(entries, CommandEntry{Name: name, Description: s.Description})
	}

	if err := writeSharedSkill(outputDir, entries); err != nil {
		return fmt.Errorf("writing shared skill: %w", err)
	}
	fmt.Fprintf(os.Stderr, "  wrote glean-shared/SKILL.md\n")

	// Generate per-command skills
	for _, name := range commands {
		if skipCommands[name] {
			continue
		}
		s, err := schema.Get(name)
		if err != nil {
			continue
		}
		if err := writeCommandSkill(outputDir, name, s); err != nil {
			return fmt.Errorf("writing skill for %s: %w", name, err)
		}
		fmt.Fprintf(os.Stderr, "  wrote glean-%s/SKILL.md\n", name)
	}

	fmt.Fprintf(os.Stderr, "\nDone. Skills written to %s/\n", outputDir)
	return nil
}

func writeSharedSkill(outputDir string, commands []CommandEntry) error {
	dir := filepath.Join(outputDir, "glean-shared")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(dir, "SKILL.md"))
	if err != nil {
		return err
	}
	defer f.Close()

	data := struct {
		Commands []CommandEntry
	}{Commands: commands}
	return sharedTmpl.Execute(f, data)
}

func writeCommandSkill(outputDir, name string, s schema.CommandSchema) error {
	dir := filepath.Join(outputDir, "glean-"+name)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(dir, "SKILL.md"))
	if err != nil {
		return err
	}
	defer f.Close()

	data := buildSkillData(name, s)
	return skillTmpl.Execute(f, data)
}

func buildSkillData(name string, s schema.CommandSchema) SkillData {
	desc := skillDescription[name]
	if desc == "" {
		desc = s.Description
	}

	// Build flags
	var flags []FlagInfo
	for flagName, fs := range s.Flags {
		def := ""
		if fs.Default != nil {
			def = fmt.Sprintf("%v", fs.Default)
		}
		typ := fs.Type
		if len(fs.Enum) > 0 {
			typ = strings.Join(fs.Enum, " \\| ")
		}
		flags = append(flags, FlagInfo{
			Name:        flagName,
			Type:        typ,
			Default:     def,
			Description: fs.Description,
			Required:    fs.Required,
		})
	}
	sort.Slice(flags, func(i, j int) bool {
		return flags[i].Name < flags[j].Name
	})

	// Build subcommands from the subcommandMap
	var subs []SubcommandInfo
	for key, sdesc := range subcommandMap {
		parts := strings.SplitN(key, ".", 2)
		if parts[0] == name {
			subs = append(subs, SubcommandInfo{Name: parts[1], Description: sdesc})
		}
	}
	sort.Slice(subs, func(i, j int) bool {
		return subs[i].Name < subs[j].Name
	})

	hasDryRun := false
	for _, fl := range flags {
		if fl.Name == "--dry-run" {
			hasDryRun = true
			break
		}
	}

	return SkillData{
		Name:        "glean-" + name,
		Description: desc,
		Command:     name,
		SchemaDesc:  s.Description,
		Flags:       flags,
		Subcommands: subs,
		Example:     s.Example,
		HasDryRun:   hasDryRun,
	}
}

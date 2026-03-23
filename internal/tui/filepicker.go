package tui

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const maxFileBytes = 10_000

// attachedFile holds a file the user has tagged with @ for inclusion
// in the next chat message sent to Glean.
type attachedFile struct {
	Path    string
	Content string
}

// parseFileQuery extracts the partial path typed after the last `@` in s.
// Returns ("", false) if no live @ query is present.
// A space or tab after @ means the token is complete — not a live query.
func parseFileQuery(s string) (string, bool) {
	idx := strings.LastIndex(s, "@")
	if idx < 0 {
		return "", false
	}
	rest := s[idx+1:]
	if strings.ContainsAny(rest, " \t\n") {
		return "", false
	}
	return rest, true
}

// scanFiles returns up to 20 file/directory paths that match the partial query.
// Hidden files (dotfiles) are excluded unless the query starts with ".".
func scanFiles(query string) []string {
	dir := "."
	prefix := query

	if i := strings.LastIndex(query, "/"); i >= 0 {
		dir = query[:i]
		if dir == "" {
			dir = "/"
		}
		prefix = query[i+1:]
	}

	if strings.HasPrefix(dir, "~/") {
		home, err := os.UserHomeDir()
		if err == nil {
			dir = filepath.Join(home, dir[2:])
		}
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}

	var results []string
	for _, e := range entries {
		name := e.Name()
		if strings.HasPrefix(name, ".") && !strings.HasPrefix(prefix, ".") {
			continue
		}
		if prefix != "" && !strings.HasPrefix(strings.ToLower(name), strings.ToLower(prefix)) {
			continue
		}
		path := filepath.Join(dir, name)
		if dir == "." {
			path = name
		}
		if e.IsDir() {
			path += "/"
		}
		results = append(results, path)
		if len(results) >= 20 {
			break
		}
	}
	return results
}

// readAttachedFile reads a file and returns an attachedFile.
// Returns an error for binary files (null bytes detected) or unreadable files.
// Content is truncated to maxFileBytes with a notice if needed.
func readAttachedFile(path string) (attachedFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return attachedFile{}, fmt.Errorf("reading %s: %w", path, err)
	}

	// Binary detection: scan first 512 bytes for null bytes.
	sample := data
	if len(sample) > 512 {
		sample = sample[:512]
	}
	if bytes.ContainsRune(sample, 0) {
		return attachedFile{}, fmt.Errorf("%s appears to be a binary file — cannot attach", path)
	}

	content := string(data)
	if len(content) > maxFileBytes {
		content = content[:maxFileBytes] + "\n[... truncated at 10,000 characters ...]"
	}
	return attachedFile{Path: path, Content: content}, nil
}

// buildFileContext prepends file contents to the user message in a format
// Glean's AI can understand. Returns the original message unchanged if no files.
func buildFileContext(files []attachedFile, userMessage string) string {
	if len(files) == 0 {
		return userMessage
	}
	var sb strings.Builder
	for _, f := range files {
		fmt.Fprintf(&sb, "[File: %s]\n```\n%s\n```\n\n", f.Path, f.Content)
	}
	sb.WriteString(userMessage)
	return sb.String()
}

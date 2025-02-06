package testutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Fixtures provides a way to load test fixtures from files
type Fixtures struct {
	t        *testing.T
	basePath string
	files    map[string][]byte
}

// NewFixtures creates a new Fixtures instance with the given fixture files.
// The paths should be relative to the test file's location.
func NewFixtures(t *testing.T, paths ...string) *Fixtures {
	t.Helper()

	f := &Fixtures{
		t:        t,
		basePath: filepath.Join("testdata", "fixtures"),
		files:    make(map[string][]byte),
	}

	for _, path := range paths {
		name := filepath.Base(path)
		name = name[:len(name)-len(filepath.Ext(name))] // Remove extension
		f.files[name] = f.readFile(path)
	}

	return f
}

// Load returns the contents of a fixture file by name (without extension)
func (f *Fixtures) Load(name string) []byte {
	f.t.Helper()

	data, ok := f.files[name]
	if !ok {
		f.t.Fatalf("fixture not found: %s", name)
	}

	return data
}

// LoadAsStream returns the contents of a fixture file as a stream of JSON objects.
// This is useful for API responses that are streamed as newline-delimited JSON.
// If the file contains a single JSON object, it will be converted to a stream format.
// If the file contains multiple JSON objects (one per line), they will be preserved.
func (f *Fixtures) LoadAsStream(name string) []byte {
	f.t.Helper()

	data := f.Load(name)
	content := string(data)

	// Check if the content is already in NDJSON format (one JSON object per line)
	lines := strings.Split(content, "\n")
	var validLines []string
	hasValidJSON := false

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Try to parse as JSON to determine if it's a complete object
		if json.Valid([]byte(line)) {
			hasValidJSON = true
		}
		validLines = append(validLines, line)
	}

	// If no valid JSON lines found, treat as a single JSON object
	if !hasValidJSON {
		// Compact the JSON to remove formatting
		if json.Valid(data) {
			var buf bytes.Buffer
			if err := json.Compact(&buf, data); err == nil {
				validLines = []string{buf.String()}
			} else {
				// If we can't compact it, it might be intentionally invalid JSON
				validLines = []string{strings.TrimSpace(content)}
			}
		} else {
			// Invalid JSON, just use as is
			validLines = []string{strings.TrimSpace(content)}
		}
	}

	// Join with newlines and ensure trailing newline
	result := strings.Join(validLines, "\n")
	if !strings.HasSuffix(result, "\n") {
		result += "\n"
	}

	return []byte(result)
}

// LoadAsText returns the contents of a fixture file as plain text
func (f *Fixtures) LoadAsText(name string) string {
	f.t.Helper()
	return string(f.Load(name))
}

// readFile reads a fixture file from disk
func (f *Fixtures) readFile(path string) []byte {
	f.t.Helper()

	fullPath := filepath.Join(f.basePath, path)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		f.t.Fatalf("failed to read fixture %s: %v", path, err)
	}

	return data
}

// WithBasePath sets a custom base path for fixture files
func (f *Fixtures) WithBasePath(path string) *Fixtures {
	f.basePath = path
	return f
}

// MustExist verifies that all specified fixtures exist
func (f *Fixtures) MustExist(paths ...string) error {
	for _, path := range paths {
		fullPath := filepath.Join(f.basePath, path)
		if _, err := os.Stat(fullPath); err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("fixture file does not exist: %s", path)
			}
			return fmt.Errorf("error checking fixture file %s: %v", path, err)
		}
	}
	return nil
}

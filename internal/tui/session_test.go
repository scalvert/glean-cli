package tui

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAppendTurn_JSONL(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.jsonl")

	s := &Session{path: path}
	require.NoError(t, s.AppendTurn(Turn{Role: "user", Content: "hello"}))
	require.NoError(t, s.AppendTurn(Turn{Role: "assistant", Content: "hi there"}))
	require.NoError(t, s.AppendTurn(Turn{Role: "user", Content: "bye"}))

	assert.Len(t, s.Turns, 3)

	loaded, ok := loadJSONL(path)
	require.True(t, ok)
	assert.Len(t, loaded.Turns, 3)
	assert.Equal(t, "hello", loaded.Turns[0].Content)
	assert.Equal(t, "hi there", loaded.Turns[1].Content)
	assert.Equal(t, "bye", loaded.Turns[2].Content)
}

func TestLoadJSONL_SkipsCorruptLines(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.jsonl")

	lines := []string{
		`{"role":"user","content":"first"}`,
		`{not valid json`,
		`{"role":"assistant","content":"second"}`,
	}
	require.NoError(t, os.WriteFile(path, []byte(strings.Join(lines, "\n")+"\n"), 0600))

	s, ok := loadJSONL(path)
	require.True(t, ok)
	assert.Len(t, s.Turns, 2)
	assert.Equal(t, "first", s.Turns[0].Content)
	assert.Equal(t, "second", s.Turns[1].Content)
}

func TestMigrateFromJSON(t *testing.T) {
	dir := t.TempDir()
	jsonPath := filepath.Join(dir, "latest.json")
	jsonlPath := filepath.Join(dir, "latest.jsonl")

	legacy := Session{
		Turns: []Turn{
			{Role: "user", Content: "old question"},
			{Role: "assistant", Content: "old answer"},
		},
	}
	data, err := json.MarshalIndent(legacy, "", "  ")
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(jsonPath, data, 0600))

	s, ok := migrateFromJSON(jsonPath, jsonlPath)
	require.True(t, ok)
	assert.Len(t, s.Turns, 2)

	// JSONL file should exist
	_, err = os.Stat(jsonlPath)
	assert.NoError(t, err)

	// Legacy JSON file should be removed
	_, err = os.Stat(jsonPath)
	assert.True(t, os.IsNotExist(err))

	// Verify JSONL is readable
	loaded, ok := loadJSONL(jsonlPath)
	require.True(t, ok)
	assert.Len(t, loaded.Turns, 2)
	assert.Equal(t, "old question", loaded.Turns[0].Content)
}

func TestLoadJSONL_NonExistent(t *testing.T) {
	_, ok := loadJSONL("/nonexistent/path.jsonl")
	assert.False(t, ok)
}

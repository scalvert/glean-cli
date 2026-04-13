package tui

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

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

	_, err = os.Stat(jsonlPath)
	assert.NoError(t, err)

	_, err = os.Stat(jsonPath)
	assert.True(t, os.IsNotExist(err))

	loaded, ok := loadJSONL(jsonlPath)
	require.True(t, ok)
	assert.Len(t, loaded.Turns, 2)
	assert.Equal(t, "old question", loaded.Turns[0].Content)
}

func TestLoadJSONL_NonExistent(t *testing.T) {
	_, ok := loadJSONL("/nonexistent/path.jsonl")
	assert.False(t, ok)
}

func TestNewSessionID(t *testing.T) {
	id := newSessionID()
	assert.Len(t, id, 36) // UUID format: 8-4-4-4-12
	assert.Equal(t, byte('-'), id[8])
	assert.Equal(t, byte('-'), id[13])
	assert.Equal(t, byte('-'), id[18])
	assert.Equal(t, byte('-'), id[23])

	id2 := newSessionID()
	assert.NotEqual(t, id, id2)
}

func TestNewSession_GetsUUID(t *testing.T) {
	dir := t.TempDir()
	s := &Session{}
	s.id = ""
	s.path = ""

	// Override sessionsDir by setting path directly
	s.id = newSessionID()
	s.path = filepath.Join(dir, s.id+".jsonl")

	require.NoError(t, s.AppendTurn(Turn{Role: "user", Content: "hello"}))

	// File should be named with UUID
	assert.True(t, strings.HasSuffix(s.path, ".jsonl"))
	assert.NotContains(t, s.path, "latest")
	assert.Len(t, s.id, 36)

	_, err := os.Stat(s.path)
	assert.NoError(t, err)
}

func TestFindLatestSession_PicksMostRecent(t *testing.T) {
	dir := t.TempDir()

	// Create two session files with different mtimes
	older := filepath.Join(dir, "older-id.jsonl")
	newer := filepath.Join(dir, "newer-id.jsonl")

	require.NoError(t, os.WriteFile(older, []byte(`{"role":"user","content":"old"}`+"\n"), 0600))
	// Ensure different mtime
	past := time.Now().Add(-10 * time.Second)
	require.NoError(t, os.Chtimes(older, past, past))

	require.NoError(t, os.WriteFile(newer, []byte(`{"role":"user","content":"new"}`+"\n"), 0600))

	path, id := findLatestSession(dir)
	assert.Equal(t, newer, path)
	assert.Equal(t, "newer-id", id)
}

func TestFindLatestSession_IgnoresNonJSONL(t *testing.T) {
	dir := t.TempDir()

	require.NoError(t, os.WriteFile(filepath.Join(dir, "notes.txt"), []byte("not a session"), 0600))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "session.jsonl"), []byte(`{"role":"user","content":"hi"}`+"\n"), 0600))

	path, id := findLatestSession(dir)
	assert.Equal(t, filepath.Join(dir, "session.jsonl"), path)
	assert.Equal(t, "session", id)
}

func TestFindLatestSession_EmptyDir(t *testing.T) {
	dir := t.TempDir()
	path, id := findLatestSession(dir)
	assert.Empty(t, path)
	assert.Empty(t, id)
}

package update

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsNewer(t *testing.T) {
	tests := []struct {
		name           string
		latestTag      string
		currentVersion string
		want           bool
	}{
		{"newer patch", "v0.2.1", "v0.2.0", true},
		{"newer minor", "v0.3.0", "v0.2.0", true},
		{"newer major", "v1.0.0", "v0.9.0", true},
		{"same version", "v0.2.0", "v0.2.0", false},
		{"older version", "v0.1.0", "v0.2.0", false},
		{"without v prefix", "0.2.1", "0.2.0", true},
		{"mixed v prefix", "v0.2.1", "0.2.0", true},
		{"double digit minor newer", "v0.10.0", "v0.9.2", true},
		{"double digit minor older", "v0.9.2", "v0.10.0", false},
		{"double digit patch newer", "v0.9.12", "v0.9.9", true},
		{"double digit major newer", "v10.0.0", "v9.0.0", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isNewer(tt.latestTag, tt.currentVersion)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCheck_DevVersion(t *testing.T) {
	assert.Equal(t, "", check("dev"))
	assert.Equal(t, "", check(""))
}

func TestCheckAsync_DevVersion(t *testing.T) {
	ch := CheckAsync("dev")
	notice, ok := <-ch
	assert.False(t, ok, "channel should be closed without sending")
	assert.Empty(t, notice)
}

func TestReadCache_EmptyPath(t *testing.T) {
	_, err := readCache("")
	assert.Error(t, err)
}

func TestReadCache_MissingFile(t *testing.T) {
	_, err := readCache(filepath.Join(t.TempDir(), "nonexistent.json"))
	assert.Error(t, err)
}

func TestWriteAndReadCache(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, ".glean", "update-check.json")

	now := time.Now().Truncate(time.Second)
	entry := cacheEntry{CheckedAt: now, LatestTag: "v1.2.3"}

	err := writeCache(path, entry)
	require.NoError(t, err)

	got, err := readCache(path)
	require.NoError(t, err)
	assert.Equal(t, "v1.2.3", got.LatestTag)
	assert.True(t, got.CheckedAt.Equal(now))
}

func TestWriteCache_EmptyPath(t *testing.T) {
	err := writeCache("", cacheEntry{})
	assert.NoError(t, err, "empty path is a no-op, not an error")
}

func TestReadCache_InvalidJSON(t *testing.T) {
	path := filepath.Join(t.TempDir(), "bad.json")
	require.NoError(t, os.WriteFile(path, []byte("not json"), 0o600))

	_, err := readCache(path)
	assert.Error(t, err)
}

func TestWriteCache_CreatesDirectories(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "a", "b", "c", "cache.json")

	err := writeCache(path, cacheEntry{LatestTag: "v0.1.0"})
	require.NoError(t, err)

	data, err := os.ReadFile(path)
	require.NoError(t, err)

	var entry cacheEntry
	require.NoError(t, json.Unmarshal(data, &entry))
	assert.Equal(t, "v0.1.0", entry.LatestTag)
}

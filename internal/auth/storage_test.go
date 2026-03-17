package auth

import (
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func withTempHome(t *testing.T) {
	t.Helper()
	tmp := t.TempDir()
	t.Setenv("HOME", tmp)
}

func TestSaveAndLoadTokens(t *testing.T) {
	withTempHome(t)
	tok := &StoredTokens{
		AccessToken:  "at-abc",
		RefreshToken: "rt-xyz",
		Expiry:       time.Now().Add(time.Hour).Truncate(time.Second),
		Email:        "steve@glean.com",
	}
	require.NoError(t, SaveTokens("myhost.glean.com", tok))

	got, err := LoadTokens("myhost.glean.com")
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, "at-abc", got.AccessToken)
	assert.Equal(t, "rt-xyz", got.RefreshToken)
	assert.Equal(t, "steve@glean.com", got.Email)
}

func TestLoadTokens_Missing(t *testing.T) {
	withTempHome(t)
	got, err := LoadTokens("nobody.glean.com")
	require.NoError(t, err)
	assert.Nil(t, got)
}

func TestDeleteTokens(t *testing.T) {
	withTempHome(t)
	tok := &StoredTokens{AccessToken: "tok"}
	require.NoError(t, SaveTokens("host.glean.com", tok))
	require.NoError(t, DeleteTokens("host.glean.com"))
	got, err := LoadTokens("host.glean.com")
	require.NoError(t, err)
	assert.Nil(t, got)
}

func TestSaveAndLoadClient(t *testing.T) {
	withTempHome(t)
	cl := &StoredClient{ClientID: "cid-123", ClientSecret: "cs-abc"}
	require.NoError(t, SaveClient("myhost.glean.com", cl))

	got, err := LoadClient("myhost.glean.com")
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, "cid-123", got.ClientID)
}

func TestStoredTokens_IsExpired(t *testing.T) {
	assert.True(t, (&StoredTokens{Expiry: time.Now().Add(-time.Minute)}).IsExpired())
	assert.False(t, (&StoredTokens{Expiry: time.Now().Add(time.Hour)}).IsExpired())
	assert.False(t, (&StoredTokens{}).IsExpired())
}

func TestStateDir_FilePermissions(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Windows does not enforce Unix file permission bits")
	}
	withTempHome(t)
	tok := &StoredTokens{AccessToken: "tok"}
	require.NoError(t, SaveTokens("host.glean.com", tok))

	dir := stateDir("host.glean.com")
	info, err := os.Stat(dir)
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0700), info.Mode().Perm())

	tokFile := tokensPath("host.glean.com")
	fi, err := os.Stat(tokFile)
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0600), fi.Mode().Perm())
}

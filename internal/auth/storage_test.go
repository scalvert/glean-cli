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

func TestDeleteClient(t *testing.T) {
	withTempHome(t)
	cl := &StoredClient{ClientID: "cid-123", ClientSecret: "cs-abc"}
	require.NoError(t, SaveClient("host.glean.com", cl))
	require.NoError(t, DeleteClient("host.glean.com"))
	got, err := LoadClient("host.glean.com")
	require.NoError(t, err)
	assert.Nil(t, got)
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

func TestCanonicalHostKey(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"bare hostname", "acme-be.glean.com", "acme-be.glean.com"},
		{"https scheme", "https://acme-be.glean.com", "acme-be.glean.com"},
		{"http scheme", "http://acme-be.glean.com", "acme-be.glean.com"},
		{"trailing slash", "https://acme-be.glean.com/", "acme-be.glean.com"},
		{"multiple trailing slashes", "https://acme-be.glean.com///", "acme-be.glean.com"},
		{"mixed case", "HTTPS://Acme-Be.Glean.Com", "acme-be.glean.com"},
		{"surrounding whitespace", "  https://acme-be.glean.com  ", "acme-be.glean.com"},
		{"localhost with port", "http://localhost:8080", "localhost:8080"},
		{"vanity URL", "https://acmecorp-pl.glean.com", "acmecorp-pl.glean.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, canonicalHostKey(tt.input))
		})
	}
}

// TestStateDir_SameForLegacyAndNew ensures the GLEAN_HOST → GLEAN_SERVER_URL
// migration does not invalidate existing OAuth tokens: a bare hostname, the
// same hostname with an https scheme, and the same URL with a trailing slash
// must all hash to the same state directory.
func TestStateDir_SameForLegacyAndNew(t *testing.T) {
	withTempHome(t)

	bare := stateDir("acme-be.glean.com")
	scheme := stateDir("https://acme-be.glean.com")
	trailing := stateDir("https://acme-be.glean.com/")
	mixedCase := stateDir("HTTPS://ACME-BE.GLEAN.COM")

	assert.Equal(t, bare, scheme, "bare hostname and https URL must share a state dir")
	assert.Equal(t, bare, trailing, "trailing slash must not change state dir")
	assert.Equal(t, bare, mixedCase, "case must not change state dir")
}

// TestStateDir_DifferentHostsDiffer ensures distinct hosts still resolve to
// distinct directories after canonicalization.
func TestStateDir_DifferentHostsDiffer(t *testing.T) {
	withTempHome(t)

	a := stateDir("acme-be.glean.com")
	b := stateDir("other-be.glean.com")
	assert.NotEqual(t, a, b)
}

// TestTokensSurviveMigrationInputChange is the end-to-end guarantee: a token
// saved under a legacy GLEAN_HOST value is recoverable when the caller later
// passes the same tenant as a full GLEAN_SERVER_URL.
func TestTokensSurviveMigrationInputChange(t *testing.T) {
	withTempHome(t)
	tok := &StoredTokens{AccessToken: "legacy-token"}
	require.NoError(t, SaveTokens("acme-be.glean.com", tok))

	got, err := LoadTokens("https://acme-be.glean.com")
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, "legacy-token", got.AccessToken)
}

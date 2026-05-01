package auth_test

import (
	"context"
	"testing"
	"time"

	"github.com/gleanwork/glean-cli/internal/auth"
	"github.com/gleanwork/glean-cli/internal/auth/authtest"
	gleanClient "github.com/gleanwork/glean-cli/internal/client"
	"github.com/gleanwork/glean-cli/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func oauthToken() *auth.StoredTokens {
	return &auth.StoredTokens{
		AccessToken:   "oauth-access-token",
		RefreshToken:  "oauth-refresh-token",
		Expiry:        time.Now().Add(time.Hour),
		Email:         "user@example.com",
		TokenType:     "Bearer",
		TokenEndpoint: "https://example.com/oauth/token",
	}
}

func TestOAuthLoginStateRequiresPersistedServerURLAfterEnvIsRemoved(t *testing.T) {
	authtest.IsolateAuthState(t)

	const serverURL = "https://acme-be.glean.com"
	require.NoError(t, auth.SaveTokens(serverURL, oauthToken()))

	t.Setenv("GLEAN_SERVER_URL", serverURL)
	cfg, err := config.LoadConfig()
	require.NoError(t, err)

	token, authType := gleanClient.ResolveToken(cfg)
	assert.Equal(t, "oauth-access-token", token)
	assert.Equal(t, "OAUTH", authType)

	// Simulate a fresh shell/session after login where GLEAN_SERVER_URL is no longer set.
	t.Setenv("GLEAN_SERVER_URL", "")
	cfg, err = config.LoadConfig()
	require.NoError(t, err)
	assert.Empty(t, cfg.GleanServerURL, "server URL was never persisted by login")

	token, authType = gleanClient.ResolveToken(cfg)
	assert.Empty(t, token)
	assert.Empty(t, authType)
}

func TestOAuthTokenResolvesWhenServerURLIsPersisted(t *testing.T) {
	authtest.IsolateAuthState(t)

	const serverURL = "https://acme-be.glean.com"
	require.NoError(t, config.SaveServerURLToFile(serverURL))
	require.NoError(t, auth.SaveTokens(serverURL, oauthToken()))

	cfg, err := config.LoadConfig()
	require.NoError(t, err)
	require.Equal(t, serverURL, cfg.GleanServerURL)

	token, authType := gleanClient.ResolveToken(cfg)
	assert.Equal(t, "oauth-access-token", token)
	assert.Equal(t, "OAUTH", authType)
}

// TestTokensResolveAcrossSchemeForms confirms the canonical-host-key
// guarantee end-to-end at the auth layer: a token saved under a bare hostname
// is loadable when the persisted config later hands us a full URL. This is
// the behaviour that lets existing OAuth sessions survive the migration to
// GLEAN_SERVER_URL without forcing a re-login.
func TestTokensResolveAcrossSchemeForms(t *testing.T) {
	authtest.IsolateAuthState(t)

	// Tokens saved under the legacy bare-hostname shape.
	require.NoError(t, auth.SaveTokens("acme-be.glean.com", oauthToken()))

	// Config has moved to the new full-URL shape.
	require.NoError(t, config.SaveServerURLToFile("https://acme-be.glean.com"))

	cfg, err := config.LoadConfig()
	require.NoError(t, err)

	token, authType := gleanClient.ResolveToken(cfg)
	assert.Equal(t, "oauth-access-token", token, "tokens must be discoverable under the new server URL shape")
	assert.Equal(t, "OAUTH", authType)
}

func TestStaleAPITokenClearedOnOAuthLogin(t *testing.T) {
	authtest.IsolateAuthState(t)

	const serverURL = "https://acme-be.glean.com"

	// Simulate a stale API token in config.
	require.NoError(t, config.SaveConfig(serverURL, "stale-api-token"))

	cfg, err := config.LoadConfig()
	require.NoError(t, err)
	assert.Equal(t, "stale-api-token", cfg.GleanToken, "precondition: stale token exists")

	// Simulate what persistLoginState does (called during OAuth login).
	// We can't call Login() directly since it requires a browser, but
	// persistLoginState is the function that should clear stale tokens.
	require.NoError(t, config.ClearTokenFromStorage())
	require.NoError(t, config.SaveServerURLToFile(serverURL))
	require.NoError(t, auth.SaveTokens(serverURL, oauthToken()))

	// After OAuth login, the stale API token should be gone.
	cfg, err = config.LoadConfig()
	require.NoError(t, err)
	assert.Empty(t, cfg.GleanToken, "stale API token should be cleared after OAuth login")
	assert.Equal(t, serverURL, cfg.GleanServerURL, "server URL should remain")

	// OAuth token should now be resolvable.
	token, authType := gleanClient.ResolveToken(cfg)
	assert.Equal(t, "oauth-access-token", token)
	assert.Equal(t, "OAUTH", authType)
}

func TestLogoutClearsPersistedServerURLAndOAuthTokens(t *testing.T) {
	authtest.IsolateAuthState(t)

	const serverURL = "https://acme-be.glean.com"
	require.NoError(t, config.SaveServerURLToFile(serverURL))
	require.NoError(t, auth.SaveTokens(serverURL, oauthToken()))
	require.NoError(t, auth.SaveClient(serverURL, &auth.StoredClient{ClientID: "cid-123"}))

	require.NoError(t, auth.Logout(context.Background()))

	cfg, err := config.LoadConfig()
	require.NoError(t, err)
	assert.Empty(t, cfg.GleanServerURL)
	assert.Empty(t, cfg.GleanToken)

	tok, err := auth.LoadTokens(serverURL)
	require.NoError(t, err)
	assert.Nil(t, tok)

	cl, err := auth.LoadClient(serverURL)
	require.NoError(t, err)
	assert.Nil(t, cl)
}

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

func TestOAuthLoginStateRequiresPersistedHostAfterEnvHostIsRemoved(t *testing.T) {
	authtest.IsolateAuthState(t)

	const host = "acme-be.glean.com"
	require.NoError(t, auth.SaveTokens(host, oauthToken()))

	t.Setenv("GLEAN_HOST", host)
	cfg, err := config.LoadConfig()
	require.NoError(t, err)

	token, authType := gleanClient.ResolveToken(cfg)
	assert.Equal(t, "oauth-access-token", token)
	assert.Equal(t, "OAUTH", authType)

	// Simulate a fresh shell/session after login where GLEAN_HOST is no longer set.
	t.Setenv("GLEAN_HOST", "")
	cfg, err = config.LoadConfig()
	require.NoError(t, err)
	assert.Empty(t, cfg.GleanHost, "host was never persisted by login")

	token, authType = gleanClient.ResolveToken(cfg)
	assert.Empty(t, token)
	assert.Empty(t, authType)
}

func TestOAuthTokenResolvesWhenHostIsPersisted(t *testing.T) {
	authtest.IsolateAuthState(t)

	const host = "acme-be.glean.com"
	require.NoError(t, config.SaveHostToFile(host))
	require.NoError(t, auth.SaveTokens(host, oauthToken()))

	cfg, err := config.LoadConfig()
	require.NoError(t, err)
	require.Equal(t, host, cfg.GleanHost)

	token, authType := gleanClient.ResolveToken(cfg)
	assert.Equal(t, "oauth-access-token", token)
	assert.Equal(t, "OAUTH", authType)
}

func TestShortFormHostNormalizesConsistently(t *testing.T) {
	authtest.IsolateAuthState(t)

	const shortHost = "acme"
	const normalizedHost = "acme-be.glean.com"

	// Simulate: GLEAN_HOST=acme (short form) was set during login.
	// persistLoginState normalizes the host in the config file,
	// and SaveTokens must use the same normalized value.
	require.NoError(t, config.SaveHostToFile(shortHost))
	require.NoError(t, auth.SaveTokens(config.NormalizeHost(shortHost), oauthToken()))

	// Simulate next session: no env var, host loaded from config file.
	cfg, err := config.LoadConfig()
	require.NoError(t, err)
	require.Equal(t, normalizedHost, cfg.GleanHost, "config file should contain normalized host")

	// Token lookup must use the same normalized host.
	token, authType := gleanClient.ResolveToken(cfg)
	assert.Equal(t, "oauth-access-token", token, "tokens should be found via normalized host")
	assert.Equal(t, "OAUTH", authType)
}

func TestLogoutClearsPersistedHostAndOAuthTokens(t *testing.T) {
	authtest.IsolateAuthState(t)

	const host = "acme-be.glean.com"
	require.NoError(t, config.SaveHostToFile(host))
	require.NoError(t, auth.SaveTokens(host, oauthToken()))
	require.NoError(t, auth.SaveClient(host, &auth.StoredClient{ClientID: "cid-123"}))

	require.NoError(t, auth.Logout(context.Background()))

	cfg, err := config.LoadConfig()
	require.NoError(t, err)
	assert.Empty(t, cfg.GleanHost)
	assert.Empty(t, cfg.GleanToken)

	tok, err := auth.LoadTokens(host)
	require.NoError(t, err)
	assert.Nil(t, tok)

	cl, err := auth.LoadClient(host)
	require.NoError(t, err)
	assert.Nil(t, cl)
}

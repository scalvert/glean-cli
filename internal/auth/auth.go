package auth

import (
	"bufio"
	"context"
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gleanwork/glean-cli/internal/config"
	"github.com/int128/oauth2cli"
	"github.com/pkg/browser"
	"golang.org/x/oauth2"
)

//go:embed success.html
var successHTML string

// Login performs the full OAuth 2.0 PKCE login flow for the configured Glean host.
// If the host is not configured, prompts for a work email and auto-discovers it.
// If the instance doesn't support OAuth, falls back to an inline API token prompt.
func Login(ctx context.Context) error {
	host, err := resolveHost(ctx)
	if err != nil {
		return err
	}

	provider, endpoint, registrationEndpoint, err := discover(ctx, host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nOAuth discovery failed: %v\n", err)
		return promptForAPIToken(host)
	}

	// Find a free port for the local callback server.
	// This must happen before DCR so we register the exact redirect URI
	// that oauth2cli will use — a mismatch causes a silent hang.
	port, err := findFreePort()
	if err != nil {
		return fmt.Errorf("finding callback port: %w", err)
	}
	redirectURI := fmt.Sprintf("http://127.0.0.1:%d/callback", port)

	// Always do fresh DCR per login — the redirect URI (port) changes each time.
	clientID, clientSecret, err := dcrOrStaticClient(ctx, host, registrationEndpoint, redirectURI)
	if err != nil {
		return fmt.Errorf("resolving OAuth client: %w", err)
	}

	verifier := oauth2.GenerateVerifier()
	scopes := resolveScopes(provider)

	oauthCfg := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     endpoint,
		Scopes:       scopes,
		RedirectURL:  redirectURI,
	}

	// oauth2cli v1.15.1 does not open the browser itself — the caller must do it.
	// LocalServerReadyChan receives the local server URL once the callback server
	// is ready. We open the browser to that URL (which the local server redirects
	// to the real OAuth page), and also print the direct auth URL as a fallback.
	state := oauth2.GenerateVerifier()[:20]
	authURL := oauthCfg.AuthCodeURL(state, oauth2.S256ChallengeOption(verifier))

	readyChan := make(chan string, 1)
	go func() {
		select {
		case localURL := <-readyChan:
			fmt.Printf("Opening your browser to authenticate with Glean…\n")
			fmt.Printf("If your browser doesn't open, visit:\n  %s\n\n", authURL)
			fmt.Printf("Waiting for you to complete login in the browser…\n")
			if err := browser.OpenURL(localURL); err != nil {
				// Browser failed to open — the printed URL is the fallback.
				fmt.Printf("(Could not open browser automatically: %v)\n", err)
			}
		case <-ctx.Done():
		}
	}()

	token, err := oauth2cli.GetToken(ctx, oauth2cli.Config{
		OAuth2Config: oauthCfg,
		State:        state,
		// LocalServerBindAddress and LocalServerCallbackPath must match the
		// redirect_uri registered via DCR exactly. oauth2cli constructs the
		// redirect URL from LocalServerBindAddress (127.0.0.1:{port}) + path.
		LocalServerCallbackPath: "/callback",
		LocalServerBindAddress:  []string{fmt.Sprintf("127.0.0.1:%d", port)},
		LocalServerReadyChan:    readyChan,
		AuthCodeOptions:         []oauth2.AuthCodeOption{oauth2.S256ChallengeOption(verifier)},
		TokenRequestOptions:     []oauth2.AuthCodeOption{oauth2.VerifierOption(verifier)},
		LocalServerSuccessHTML:  successHTML,
	})
	if err != nil {
		return fmt.Errorf("OAuth login failed: %w", err)
	}

	email := extractEmailFromToken(ctx, provider, clientID, token)

	stored := &StoredTokens{
		AccessToken:   token.AccessToken,
		RefreshToken:  token.RefreshToken,
		Expiry:        token.Expiry,
		Email:         email,
		TokenType:     token.TokenType,
		TokenEndpoint: oauthCfg.Endpoint.TokenURL, // enables future token refresh
	}
	if err := SaveTokens(host, stored); err != nil {
		return fmt.Errorf("saving tokens: %w", err)
	}

	if email != "" {
		fmt.Printf("✓ Authenticated as %s (%s)\n", email, host)
	} else {
		fmt.Printf("✓ Authenticated with Glean (%s)\n", host)
	}
	return nil
}

// Logout removes stored OAuth tokens and clears the config file.
func Logout(ctx context.Context) error {
	cfg, err := config.LoadConfig()
	if err != nil || cfg.GleanHost == "" {
		return fmt.Errorf("no Glean host configured")
	}
	if err := DeleteTokens(cfg.GleanHost); err != nil {
		return fmt.Errorf("removing tokens: %w", err)
	}
	// Also wipe the config file so any stored API token / host is cleared.
	if config.ConfigPath != "" {
		_ = os.Remove(config.ConfigPath)
	}
	fmt.Printf("✓ Logged out from Glean (%s)\n", cfg.GleanHost)
	return nil
}

// Status prints the current authentication state.
func Status(ctx context.Context) error {
	cfg, _ := config.LoadConfig()
	if cfg == nil || cfg.GleanHost == "" {
		fmt.Println("Not configured.")
		fmt.Println("Run 'glean auth login' to authenticate.")
		return nil
	}

	if cfg.GleanToken != "" {
		masked := config.MaskToken(cfg.GleanToken)
		fmt.Printf("✓ Authenticated via API token\n  Host:  %s\n  Token: %s\n", cfg.GleanHost, masked)
		return nil
	}

	tok, err := LoadTokens(cfg.GleanHost)
	if err != nil {
		return fmt.Errorf("reading stored tokens: %w", err)
	}
	if tok == nil {
		fmt.Println("Not authenticated.")
		fmt.Println("Run 'glean auth login' to authenticate.")
		return nil
	}
	if tok.IsExpired() {
		fmt.Println("Token expired.")
		fmt.Println("Run 'glean auth login' to re-authenticate.")
		return nil
	}

	expStr := "no expiry"
	if !tok.Expiry.IsZero() {
		remaining := time.Until(tok.Expiry).Round(time.Minute)
		expStr = fmt.Sprintf("expires %s (in %v)", tok.Expiry.UTC().Format(time.RFC3339), remaining)
	}
	if tok.Email != "" {
		fmt.Printf("✓ Authenticated as %s (%s)\n  Token %s\n", tok.Email, cfg.GleanHost, expStr)
	} else {
		fmt.Printf("✓ Authenticated with Glean (%s)\n  Token %s\n", cfg.GleanHost, expStr)
	}
	return nil
}

// EnsureAuth returns nil if the client has usable credentials.
func EnsureAuth(ctx context.Context) error {
	cfg, err := config.LoadConfig()
	if err == nil && cfg.GleanToken != "" {
		return nil
	}
	if err == nil && cfg.GleanHost != "" {
		tok, _ := LoadTokens(cfg.GleanHost)
		if tok != nil && !tok.IsExpired() {
			return nil
		}
	}
	return fmt.Errorf("not authenticated — run 'glean auth login' to authenticate")
}

// LoadOAuthToken returns a valid, non-expired OAuth access token for host, or "".
// If the stored token is expired and a refresh token is available, it attempts
// a silent refresh and persists the new tokens before returning.
func LoadOAuthToken(host string) string {
	tok, err := LoadTokens(host)
	if err != nil || tok == nil {
		return ""
	}
	if !tok.IsExpired() {
		return tok.AccessToken
	}
	// Token expired — attempt silent refresh.
	if tok.RefreshToken != "" && tok.TokenEndpoint != "" {
		if refreshed, err := refreshOAuthToken(host, tok); err == nil {
			return refreshed.AccessToken
		}
	}
	return ""
}

// refreshOAuthToken exchanges a stored refresh_token for a new access token.
// The refreshed tokens are persisted to storage. Returns the updated StoredTokens.
func refreshOAuthToken(host string, tok *StoredTokens) (*StoredTokens, error) {
	cl, err := LoadClient(host)
	if err != nil || cl == nil {
		return nil, fmt.Errorf("no stored OAuth client for %s — re-run 'glean auth login'", host)
	}

	oauthCfg := oauth2.Config{
		ClientID:     cl.ClientID,
		ClientSecret: cl.ClientSecret,
		Endpoint:     oauth2.Endpoint{TokenURL: tok.TokenEndpoint},
	}

	existing := &oauth2.Token{
		AccessToken:  tok.AccessToken,
		RefreshToken: tok.RefreshToken,
		Expiry:       tok.Expiry,
		TokenType:    tok.TokenType,
	}

	src := oauthCfg.TokenSource(context.Background(), existing)
	newTok, err := src.Token()
	if err != nil {
		return nil, fmt.Errorf("token refresh failed: %w", err)
	}

	// Some providers rotate the refresh token; retain the old one if not.
	refreshToken := newTok.RefreshToken
	if refreshToken == "" {
		refreshToken = tok.RefreshToken
	}

	stored := &StoredTokens{
		AccessToken:   newTok.AccessToken,
		RefreshToken:  refreshToken,
		Expiry:        newTok.Expiry,
		Email:         tok.Email,
		TokenType:     newTok.TokenType,
		TokenEndpoint: tok.TokenEndpoint,
	}
	_ = SaveTokens(host, stored)
	return stored, nil
}

// resolveHost returns the configured Glean host, prompting for email if needed.
func resolveHost(ctx context.Context) (string, error) {
	cfg, _ := config.LoadConfig()
	if cfg != nil && cfg.GleanHost != "" {
		return cfg.GleanHost, nil
	}

	fmt.Print("Enter your work email: ")
	reader := bufio.NewReader(os.Stdin)
	email, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("reading email: %w", err)
	}
	email = strings.TrimSpace(email)

	fmt.Print("Looking up your Glean instance…")
	backendURL, err := LookupBackendURL(ctx, email)
	if err != nil {
		fmt.Println()
		return "", fmt.Errorf("could not find a Glean instance for %q: %w", email, err)
	}
	fmt.Println(" found.")

	host := strings.TrimPrefix(backendURL, "https://")
	host = strings.TrimPrefix(host, "http://")
	host = strings.SplitN(host, "/", 2)[0]

	_ = config.SaveConfig(host, "")
	return host, nil
}

// discover resolves the OAuth2 endpoint and registration endpoint for the Glean backend.
//
// Strategy:
//  1. Fetch RFC 9728 protected resource metadata → get authorization server URL
//  2. Try OIDC discovery (oidc.NewProvider) for full OIDC support
//  3. Fall back to RFC 8414 auth server metadata when OIDC is unavailable
//     (Glean uses RFC 8414 but does not serve /.well-known/openid-configuration)
//
// Returns (provider, oauth2Endpoint, registrationEndpoint, error).
// provider is nil when only RFC 8414 discovery succeeded.
func discover(ctx context.Context, host string) (*oidc.Provider, oauth2.Endpoint, string, error) {
	baseURL := "https://" + host
	meta, err := fetchProtectedResource(ctx, baseURL)
	if err != nil {
		return nil, oauth2.Endpoint{}, "", err
	}

	issuer := meta.AuthorizationServers[0]

	// Try full OIDC discovery first (supports ID token, UserInfo).
	provider, err := oidc.NewProvider(ctx, issuer)
	if err == nil {
		// Still need registration_endpoint, which oidc.Provider doesn't expose.
		authMeta, _ := fetchAuthServerMetadata(ctx, issuer)
		regEndpoint := ""
		if authMeta != nil {
			regEndpoint = authMeta.RegistrationEndpoint
		}
		return provider, provider.Endpoint(), regEndpoint, nil
	}

	// Fall back to RFC 8414 auth server metadata.
	authMeta, err := fetchAuthServerMetadata(ctx, issuer)
	if err != nil {
		return nil, oauth2.Endpoint{}, "", fmt.Errorf("OAuth discovery failed for %s: %w", issuer, err)
	}
	if authMeta.AuthorizationEndpoint == "" || authMeta.TokenEndpoint == "" {
		return nil, oauth2.Endpoint{}, "", fmt.Errorf("OAuth metadata missing required endpoints for %s", issuer)
	}

	return nil, oauth2.Endpoint{
		AuthURL:  authMeta.AuthorizationEndpoint,
		TokenURL: authMeta.TokenEndpoint,
	}, authMeta.RegistrationEndpoint, nil
}

// dcrOrStaticClient resolves the OAuth client_id/secret for a login session.
// It performs fresh DCR on each call (redirect URI includes port, so it changes).
// A successful DCR registration is persisted to storage so the same client
// credentials can be reused for token refresh later.
// Falls back to a static client configured via glean config --oauth-client-id.
func dcrOrStaticClient(ctx context.Context, host, registrationEndpoint, redirectURI string) (string, string, error) {
	if registrationEndpoint != "" {
		cl, err := registerClient(ctx, registrationEndpoint, redirectURI)
		if err == nil {
			// Persist so future token refresh can use the same client credentials.
			_ = SaveClient(host, cl)
			return cl.ClientID, cl.ClientSecret, nil
		}
		// DCR failed — log and fall through to static client
		fmt.Printf("Note: dynamic client registration failed (%v), trying static client\n", err)
	}

	cfg, _ := config.LoadConfig()
	if cfg != nil && cfg.OAuthClientID != "" {
		return cfg.OAuthClientID, cfg.OAuthClientSecret, nil
	}

	return "", "", fmt.Errorf("no OAuth client available — dynamic client registration failed and no static client is configured")
}

// resolveScopes returns the appropriate OAuth scopes for the given provider.
// For Glean native OAuth, we request all scopes supported by the CLI commands.
// The full list of supported scopes is available at:
//
//	GET <host>/.well-known/oauth-authorization-server/oauth → scopes_supported
func resolveScopes(provider *oidc.Provider) []string {
	if provider != nil {
		// Full OIDC: standard scopes for ID token + email
		return []string{oidc.ScopeOpenID, "email", "profile"}
	}
	// Glean native scopes — request all scopes required by CLI commands.
	// Previously only "chat", "search", "email" were requested, which caused
	// 401 errors on agents, tools, insights, verification, pins, shortcuts, etc.
	return []string{
		"activity",
		"agents",
		"announcements",
		"answers",
		"chat",
		"collections",
		"documents",
		"email",
		"entities",
		"insights",
		"offline_access", // enables token refresh
		"pins",
		"search",
		"shortcuts",
		"summarize",
		"tools",
		"verification",
	}
}

// findFreePort finds an available TCP port on localhost.
func findFreePort() (int, error) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

// fetchAuthServerMetadata fetches RFC 8414 Authorization Server Metadata.
//
// Per RFC 8414 §3, for an issuer with a path component (e.g. https://host/oauth),
// the discovery URL is: https://host/.well-known/oauth-authorization-server/oauth
// (origin + well-known + issuer-path), not https://host/oauth/.well-known/...
func fetchAuthServerMetadata(ctx context.Context, issuer string) (*authServerMeta, error) {
	parsed, err := url.Parse(strings.TrimRight(issuer, "/"))
	if err != nil {
		return nil, fmt.Errorf("invalid issuer URL %q: %w", issuer, err)
	}
	// RFC 8414 path-aware: origin + /.well-known/oauth-authorization-server + path
	u := parsed.Scheme + "://" + parsed.Host + "/.well-known/oauth-authorization-server" + parsed.Path
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := discoveryHTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("auth server metadata returned HTTP %d", resp.StatusCode)
	}
	var meta authServerMeta
	return &meta, json.NewDecoder(resp.Body).Decode(&meta)
}

type authServerMeta struct {
	Issuer                string `json:"issuer"`
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	RegistrationEndpoint  string `json:"registration_endpoint,omitempty"`
}

// extractEmailFromToken pulls the user email from the token.
// Tries OIDC ID token verification first (when provider available), then
// decodes the access token as a JWT (without verification) to read the email
// claim directly — works for Glean's RFC 8414 OAuth which issues JWT access tokens.
func extractEmailFromToken(ctx context.Context, provider *oidc.Provider, clientID string, token *oauth2.Token) string {
	// 1. OIDC ID token (full verification).
	if provider != nil {
		if rawIDToken, ok := token.Extra("id_token").(string); ok && rawIDToken != "" {
			verifier := provider.Verifier(&oidc.Config{ClientID: clientID})
			if idToken, err := verifier.Verify(ctx, rawIDToken); err == nil {
				var claims struct {
					Email string `json:"email"`
				}
				if err := idToken.Claims(&claims); err == nil && claims.Email != "" {
					return claims.Email
				}
			}
		}
		// UserInfo endpoint fallback.
		if ui, err := provider.UserInfo(ctx, oauth2.StaticTokenSource(token)); err == nil {
			var claims struct {
				Email string `json:"email"`
			}
			if err := ui.Claims(&claims); err == nil && claims.Email != "" {
				return claims.Email
			}
		}
	}

	// 2. Decode the access token as a JWT (no signature verification).
	// Glean issues JWT access tokens that contain the user's email claim.
	return EmailFromJWT(token.AccessToken)
}

// EmailFromJWT decodes a JWT payload (without verification) and returns the
// email claim, or "" if unavailable or the token is not a valid JWT.
func EmailFromJWT(raw string) string {
	parts := strings.Split(raw, ".")
	if len(parts) != 3 {
		return ""
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return ""
	}
	var claims struct {
		Email string `json:"email"`
	}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return ""
	}
	return claims.Email
}

// promptForAPIToken handles instances that don't support OAuth.
func promptForAPIToken(host string) error {
	fmt.Printf("There was an issue with OAuth. You can try using an API token instead.\n")
	fmt.Printf("Contact your Glean administrator to generate an API token.\n")
	fmt.Printf("  (Glean Admin → Settings → API Tokens)\n\n")
	fmt.Print("Token: ")
	reader := bufio.NewReader(os.Stdin)
	token, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("reading token: %w", err)
	}
	token = strings.TrimSpace(token)
	if token == "" {
		return fmt.Errorf("no token provided")
	}
	if err := config.SaveConfig(host, token); err != nil {
		return fmt.Errorf("saving token: %w", err)
	}
	fmt.Printf("✓ API token saved for %s\n", host)
	return nil
}

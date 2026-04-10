package auth

import (
	"bufio"
	"context"
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gleanwork/glean-cli/internal/config"
	"github.com/gleanwork/glean-cli/internal/debug"
	"github.com/gleanwork/glean-cli/internal/httputil"
	"github.com/int128/oauth2cli"
	"github.com/pkg/browser"
	"golang.org/x/oauth2"
)

var (
	loginLog     = debug.New("auth:login")
	hostLog      = debug.New("auth:resolve-host")
	discoveryLog = debug.New("auth:discovery")
	dcrLog       = debug.New("auth:dcr")
	tokenLog     = debug.New("auth:token")
	emailLog     = debug.New("auth:email")
	deviceLog    = debug.New("auth:device")
)

//go:embed success.html
var successHTML string

// errNoOAuthClient is returned by dcrOrStaticClient when neither DCR nor a
// static client is available. Login uses this to decide whether device flow
// is an appropriate fallback (as opposed to transient failures like network
// timeouts or the user closing their browser).
var errNoOAuthClient = errors.New("no OAuth client available")

// Login performs the full OAuth 2.0 login flow for the configured Glean host.
//
// Strategy (in order):
//  1. Authorization Code + PKCE via DCR or static client
//  2. Device Authorization Grant (RFC 8628) using the Glean-advertised client ID
//  3. Inline API token prompt when OAuth is not available at all
func Login(ctx context.Context) error {
	loginLog.Log("starting login flow")

	host, err := resolveHost(ctx)
	if err != nil {
		return err
	}
	loginLog.Log("host resolved: %s", host)

	disc, err := discover(ctx, host)
	if err != nil {
		loginLog.Log("OAuth discovery failed, falling back to API token: %v", err)
		fmt.Fprintf(os.Stderr, "\nOAuth discovery failed: %v\n", err)
		return promptForAPIToken(host)
	}
	loginLog.Log("OAuth discovery succeeded: auth=%s token=%s registration=%s", disc.Endpoint.AuthURL, disc.Endpoint.TokenURL, disc.RegistrationEndpoint)

	// Try DCR / static client first (standard authorization code flow).
	loginLog.Log("attempting authorization code + PKCE flow")
	authCodeErr := tryAuthCodeLogin(ctx, host, disc)
	if authCodeErr == nil {
		return nil
	}
	loginLog.Log("auth code flow failed: %v", authCodeErr)

	// Only fall back to device flow when the auth code flow failed because no
	// OAuth client could be obtained (DCR unsupported + no static client).
	// Transient failures (network, user closing browser, port conflicts) should
	// not silently switch to a different grant type.
	canDeviceFlow := disc.DeviceFlowClientID != "" && disc.DeviceAuthEndpoint != ""
	if errors.Is(authCodeErr, errNoOAuthClient) && canDeviceFlow {
		loginLog.Log("falling back to device authorization grant (client_id=%s)", disc.DeviceFlowClientID)
		fmt.Fprintf(os.Stderr, "\nYour SSO provider requires device-based login.\n")
		return deviceFlowLogin(ctx, host, disc)
	}

	return fmt.Errorf("authentication failed: %w", authCodeErr)
}

// tryAuthCodeLogin attempts the Authorization Code + PKCE flow via DCR or static client.
func tryAuthCodeLogin(ctx context.Context, host string, disc *discoveryResult) error {
	port, err := findFreePort()
	if err != nil {
		return fmt.Errorf("finding callback port: %w", err)
	}
	redirectURI := fmt.Sprintf("http://127.0.0.1:%d/glean-cli-callback", port)

	clientID, clientSecret, err := dcrOrStaticClient(ctx, host, disc.RegistrationEndpoint, redirectURI)
	if err != nil {
		return err
	}

	verifier := oauth2.GenerateVerifier()
	scopes := resolveScopes(disc.Provider)
	loginLog.Log("requesting scopes: %v", scopes)

	oauthCfg := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     disc.Endpoint,
		Scopes:       scopes,
		RedirectURL:  redirectURI,
	}

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
		LocalServerCallbackPath: "/glean-cli-callback",
		LocalServerBindAddress:  []string{fmt.Sprintf("127.0.0.1:%d", port)},
		LocalServerReadyChan:    readyChan,
		AuthCodeOptions:         []oauth2.AuthCodeOption{oauth2.S256ChallengeOption(verifier)},
		TokenRequestOptions:     []oauth2.AuthCodeOption{oauth2.VerifierOption(verifier)},
		LocalServerSuccessHTML:  successHTML,
	})
	if err != nil {
		return fmt.Errorf("OAuth login failed: %w", err)
	}

	return saveAndPrintToken(ctx, host, disc, oauthCfg.ClientID, token)
}

// saveAndPrintToken persists the OAuth token and client, then prints a success message.
func saveAndPrintToken(ctx context.Context, host string, disc *discoveryResult, clientID string, token *oauth2.Token) error {
	_ = SaveClient(host, &StoredClient{ClientID: clientID})

	email := extractEmailFromToken(ctx, disc.Provider, clientID, token)

	stored := &StoredTokens{
		AccessToken:   token.AccessToken,
		RefreshToken:  token.RefreshToken,
		Expiry:        token.Expiry,
		Email:         email,
		TokenType:     token.TokenType,
		TokenEndpoint: disc.Endpoint.TokenURL,
	}
	if err := persistLoginState(host, stored); err != nil {
		return err
	}

	if email != "" {
		fmt.Printf("✓ Authenticated as %s (%s)\n", email, host)
	} else {
		fmt.Printf("✓ Authenticated with Glean (%s)\n", host)
	}
	return nil
}

// persistLoginState stores the resolved host in config and persists OAuth tokens.
// Saving the host here ensures a successful `glean auth login` remains usable
// even when the host originally came from an environment variable.
//
// It also clears any existing API token from storage so that the new OAuth
// credentials take effect immediately — a stale API token in config would
// otherwise shadow the fresh OAuth token (ResolveToken prefers API tokens).
func persistLoginState(host string, tok *StoredTokens) error {
	if err := config.ClearTokenFromStorage(); err != nil {
		return fmt.Errorf("clearing stale API token: %w", err)
	}
	if err := config.SaveHostToFile(host); err != nil {
		return fmt.Errorf("saving host: %w", err)
	}
	if err := SaveTokens(host, tok); err != nil {
		return fmt.Errorf("saving tokens: %w", err)
	}
	return nil
}

// Logout removes stored OAuth tokens, OAuth client registration, and any saved
// config/keyring credentials for the current host.
func Logout(ctx context.Context) error {
	cfg, err := config.LoadConfig()
	if err != nil || cfg.GleanHost == "" {
		return fmt.Errorf("no Glean host configured")
	}
	if err := DeleteTokens(cfg.GleanHost); err != nil {
		return fmt.Errorf("removing tokens: %w", err)
	}
	if err := DeleteClient(cfg.GleanHost); err != nil {
		return fmt.Errorf("removing oauth client: %w", err)
	}
	if err := config.ClearConfig(); err != nil {
		return fmt.Errorf("clearing config: %w", err)
	}
	fmt.Printf("✓ Logged out from Glean (%s)\n", cfg.GleanHost)
	return nil
}

// TokenValidator validates credentials in a config against the Glean backend.
// It returns nil when the token is accepted, or an error describing the failure.
type TokenValidator func(ctx context.Context, cfg *config.Config) error

// Status prints the current authentication state.
// validateToken is used to verify API tokens against the backend (typically client.ValidateToken).
func Status(ctx context.Context, validateToken TokenValidator) error {
	cfg, _ := config.LoadConfig()
	if cfg == nil || cfg.GleanHost == "" {
		fmt.Println("Not configured.")
		fmt.Println("Run 'glean auth login' to authenticate.")
		return nil
	}

	if cfg.GleanToken != "" {
		masked := config.MaskToken(cfg.GleanToken)
		if err := validateToken(ctx, cfg); err != nil {
			fmt.Printf("✗ API token is invalid or expired\n  Host:  %s\n  Token: %s\n  Error: %v\n", cfg.GleanHost, masked, err)
			fmt.Println("Run 'glean auth login' to re-authenticate.")
			return nil
		}
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
	if err != nil {
		tokenLog.Log("load failed for %s: %v", host, err)
		return ""
	}
	if tok == nil {
		tokenLog.Log("no stored tokens for %s", host)
		return ""
	}
	if !tok.IsExpired() {
		tokenLog.Log("token valid (expires %s)", tok.Expiry.Format("15:04:05"))
		return tok.AccessToken
	}

	// Token expired — attempt silent refresh.
	if tok.RefreshToken == "" || tok.TokenEndpoint == "" {
		tokenLog.Log("token expired, cannot refresh (refresh_token=%t endpoint=%t)", tok.RefreshToken != "", tok.TokenEndpoint != "")
		return ""
	}
	tokenLog.Log("token expired, refreshing via %s", tok.TokenEndpoint)
	refreshed, err := refreshOAuthToken(host, tok)
	if err != nil {
		tokenLog.Log("refresh failed: %v", err)
		return ""
	}
	tokenLog.Log("refreshed (new expiry=%s)", refreshed.Expiry.Format("15:04:05"))
	return refreshed.AccessToken
}

// refreshOAuthToken exchanges a stored refresh_token for a new access token.
// The refreshed tokens are persisted to storage. Returns the updated StoredTokens.
func refreshOAuthToken(host string, tok *StoredTokens) (*StoredTokens, error) {
	cl, err := LoadClient(host)
	if err != nil || cl == nil {
		tokenLog.Log("no stored OAuth client for %s (err=%v)", host, err)
		return nil, fmt.Errorf("no stored OAuth client for %s — re-run 'glean auth login'", host)
	}
	tokenLog.Log("using stored client_id=%s for refresh", cl.ClientID)

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
	if err := SaveTokens(host, stored); err != nil {
		tokenLog.Log("persisting refreshed tokens failed: %v", err)
	}
	return stored, nil
}

// resolveHost returns the configured Glean host, prompting for email if needed.
// The returned host is always normalized (e.g. "linkedin" → "linkedin-be.glean.com")
// so that all downstream callers (token storage, config persistence) use a consistent value.
func resolveHost(ctx context.Context) (string, error) {
	cfg, _ := config.LoadConfig()
	if cfg != nil && cfg.GleanHost != "" {
		host := config.NormalizeHost(cfg.GleanHost)
		hostLog.Log("using configured host: %s", host)
		return host, nil
	}
	hostLog.Log("no host configured, prompting for email")

	fmt.Print("Enter your work email: ")
	reader := bufio.NewReader(os.Stdin)
	email, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("reading email: %w", err)
	}
	email = strings.TrimSpace(email)

	fmt.Print("Looking up your Glean instance…")
	hostLog.Log("looking up backend for email domain")
	backendURL, err := LookupBackendURL(ctx, email)
	if err != nil {
		fmt.Println()
		return "", fmt.Errorf("could not find a Glean instance for %q: %w", email, err)
	}
	fmt.Println(" found.")

	host := strings.TrimPrefix(backendURL, "https://")
	host = strings.TrimPrefix(host, "http://")
	host = strings.SplitN(host, "/", 2)[0]
	hostLog.Log("discovered host: %s", host)

	if err := config.SaveHostToFile(host); err != nil {
		hostLog.Log("best-effort host save failed: %v", err)
	}
	return host, nil
}

// discoveryResult holds all OAuth metadata discovered for a Glean backend.
type discoveryResult struct {
	Provider             *oidc.Provider
	Endpoint             oauth2.Endpoint
	RegistrationEndpoint string
	DeviceFlowClientID   string
	DeviceAuthEndpoint   string
}

// discover resolves the OAuth2 endpoint and registration endpoint for the Glean backend.
//
// Strategy:
//  1. Fetch RFC 9728 protected resource metadata → get authorization server URL
//  2. Try OIDC discovery (oidc.NewProvider) for full OIDC support
//  3. Fall back to RFC 8414 auth server metadata when OIDC is unavailable
//     (Glean uses RFC 8414 but does not serve /.well-known/openid-configuration)
func discover(ctx context.Context, host string) (*discoveryResult, error) {
	baseURL := "https://" + host
	discoveryLog.Log("fetching protected resource metadata: %s", baseURL)
	meta, err := fetchProtectedResource(ctx, baseURL)
	if err != nil {
		discoveryLog.Log("protected resource metadata failed: %v", err)
		return nil, err
	}

	issuer := meta.AuthorizationServers[0]
	discoveryLog.Log("authorization server: %s", issuer)

	// Try full OIDC discovery first (supports ID token, UserInfo).
	discoveryLog.Log("trying OIDC discovery at %s", issuer)
	provider, err := oidc.NewProvider(ctx, issuer)
	if err == nil {
		discoveryLog.Log("OIDC discovery succeeded")
		res := &discoveryResult{Provider: provider, Endpoint: provider.Endpoint()}
		res.DeviceFlowClientID = meta.GleanDeviceFlowClientID

		// Extract device_authorization_endpoint from OIDC provider claims
		// (RFC 8414 metadata may omit it even when OIDC metadata includes it).
		var providerClaims struct {
			RegistrationEndpoint        string `json:"registration_endpoint"`
			DeviceAuthorizationEndpoint string `json:"device_authorization_endpoint"`
		}
		if err := provider.Claims(&providerClaims); err == nil {
			res.RegistrationEndpoint = providerClaims.RegistrationEndpoint
			res.DeviceAuthEndpoint = providerClaims.DeviceAuthorizationEndpoint
		}

		// Supplement from RFC 8414 if OIDC claims were incomplete.
		if res.RegistrationEndpoint == "" || res.DeviceAuthEndpoint == "" {
			if authMeta, err := fetchAuthServerMetadata(ctx, issuer); err == nil {
				if res.RegistrationEndpoint == "" {
					res.RegistrationEndpoint = authMeta.RegistrationEndpoint
				}
				if res.DeviceAuthEndpoint == "" {
					res.DeviceAuthEndpoint = authMeta.DeviceAuthorizationEndpoint
				}
			}
		}
		return res, nil
	}
	discoveryLog.Log("OIDC discovery failed: %v, falling back to RFC 8414", err)

	// Fall back to RFC 8414 auth server metadata.
	authMeta, err := fetchAuthServerMetadata(ctx, issuer)
	if err != nil {
		return nil, fmt.Errorf("OAuth discovery failed for %s: %w", issuer, err)
	}
	if authMeta.AuthorizationEndpoint == "" || authMeta.TokenEndpoint == "" {
		discoveryLog.Log("RFC 8414 metadata incomplete: auth=%q token=%q", authMeta.AuthorizationEndpoint, authMeta.TokenEndpoint)
		return nil, fmt.Errorf("OAuth metadata missing required endpoints for %s", issuer)
	}
	discoveryLog.Log("RFC 8414 discovery succeeded: auth=%s token=%s", authMeta.AuthorizationEndpoint, authMeta.TokenEndpoint)

	return &discoveryResult{
		Endpoint: oauth2.Endpoint{
			AuthURL:  authMeta.AuthorizationEndpoint,
			TokenURL: authMeta.TokenEndpoint,
		},
		RegistrationEndpoint: authMeta.RegistrationEndpoint,
		DeviceFlowClientID:   meta.GleanDeviceFlowClientID,
		DeviceAuthEndpoint:   authMeta.DeviceAuthorizationEndpoint,
	}, nil
}

// dcrOrStaticClient resolves the OAuth client_id/secret for a login session.
// It performs fresh DCR on each call (redirect URI includes port, so it changes).
// A successful DCR registration is persisted to storage so the same client
// credentials can be reused for token refresh later.
// Falls back to a static client configured via glean config --oauth-client-id.
func dcrOrStaticClient(ctx context.Context, host, registrationEndpoint, redirectURI string) (string, string, error) {
	var dcrErr error
	if registrationEndpoint != "" {
		dcrLog.Log("registering client at %s with redirect %s", registrationEndpoint, redirectURI)
		cl, err := registerClient(ctx, registrationEndpoint, redirectURI)
		if err == nil {
			dcrLog.Log("registered client_id=%s", cl.ClientID)
			if err := SaveClient(host, cl); err != nil {
				dcrLog.Log("persisting client registration failed: %v", err)
			}
			return cl.ClientID, cl.ClientSecret, nil
		}
		dcrErr = err
		dcrLog.Log("DCR failed: %v, trying static client", err)
		fmt.Printf("Note: dynamic client registration failed (%v), trying static client\n", err)
	} else {
		dcrLog.Log("no registration endpoint, trying static client")
	}

	cfg, _ := config.LoadConfig()
	if cfg != nil && cfg.OAuthClientID != "" {
		dcrLog.Log("using static client_id=%s", cfg.OAuthClientID)
		return cfg.OAuthClientID, cfg.OAuthClientSecret, nil
	}

	if dcrErr != nil {
		return "", "", fmt.Errorf("%w: dynamic client registration failed (%v) and no static client is configured", errNoOAuthClient, dcrErr)
	}
	return "", "", fmt.Errorf("%w: no registration endpoint and no static client configured", errNoOAuthClient)
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
	discoveryLog.Log("fetching RFC 8414 metadata: %s", u)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := httputil.NewHTTPClient(10 * time.Second).Do(req)
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
	Issuer                      string `json:"issuer"`
	AuthorizationEndpoint       string `json:"authorization_endpoint"`
	TokenEndpoint               string `json:"token_endpoint"`
	RegistrationEndpoint        string `json:"registration_endpoint,omitempty"`
	DeviceAuthorizationEndpoint string `json:"device_authorization_endpoint,omitempty"`
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
					emailLog.Log("email from OIDC ID token: %s", claims.Email)
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
				emailLog.Log("email from UserInfo endpoint: %s", claims.Email)
				return claims.Email
			}
		}
	}

	// 2. Decode the access token as a JWT (no signature verification).
	// Glean issues JWT access tokens that contain the user's email claim.
	email := EmailFromJWT(token.AccessToken)
	if email != "" {
		emailLog.Log("email from JWT access token: %s", email)
	} else {
		emailLog.Log("could not extract email from token")
	}
	return email
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

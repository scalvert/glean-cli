package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gleanwork/glean-cli/internal/httputil"
	"github.com/pkg/browser"
	"golang.org/x/oauth2"
)

const (
	defaultPollInterval = 5 * time.Second
	maxPollInterval     = 60 * time.Second
	defaultExpiresIn    = 900 // 15 minutes
	maxExpiresIn        = 1800
)

// deviceAuthResponse is the response from the device authorization endpoint (RFC 8628 §3.2).
type deviceAuthResponse struct {
	DeviceCode              string `json:"device_code"`
	UserCode                string `json:"user_code"`
	VerificationURI         string `json:"verification_uri"`
	VerificationURIComplete string `json:"verification_uri_complete"`
	ExpiresIn               int    `json:"expires_in"`
	Interval                int    `json:"interval"`
}

// deviceTokenError is the error response from the token endpoint during polling.
type deviceTokenError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description,omitempty"`
}

// deviceFlowLogin performs the OAuth 2.0 Device Authorization Grant (RFC 8628).
func deviceFlowLogin(ctx context.Context, host string, disc *discoveryResult) error {
	scopes := resolveScopes(disc.Provider)
	deviceLog.Log("requesting device code from %s (client_id=%s)", disc.DeviceAuthEndpoint, disc.DeviceFlowClientID)

	authResp, err := requestDeviceCode(ctx, disc.DeviceAuthEndpoint, disc.DeviceFlowClientID, scopes)
	if err != nil {
		return fmt.Errorf("device authorization request failed: %w", err)
	}
	deviceLog.Log("device code received: user_code=%s verification_uri=%s expires_in=%d", authResp.UserCode, authResp.VerificationURI, authResp.ExpiresIn)

	verificationURL := authResp.VerificationURIComplete
	if verificationURL == "" {
		verificationURL = authResp.VerificationURI
	}

	parsed, err := url.Parse(verificationURL)
	if err != nil || parsed.Host == "" {
		return fmt.Errorf("device authorization returned invalid verification URL: %q", verificationURL)
	}
	if parsed.Scheme != "https" {
		return fmt.Errorf("device authorization returned non-HTTPS verification URL: %q", verificationURL)
	}

	fmt.Printf("\nTo authenticate, open this URL in your browser:\n\n  %s\n\n", verificationURL)
	if authResp.VerificationURIComplete == "" {
		fmt.Printf("Then enter code: %s\n\n", authResp.UserCode)
	} else {
		fmt.Printf("Your code: %s\n\n", authResp.UserCode)
	}
	fmt.Printf("Waiting for you to complete login in the browser…\n")

	_ = browser.OpenURL(verificationURL)

	deviceLog.Log("polling token endpoint %s (interval=%ds)", disc.Endpoint.TokenURL, authResp.Interval)
	token, err := pollForToken(ctx, disc.Endpoint.TokenURL, disc.DeviceFlowClientID, authResp)
	if err != nil {
		deviceLog.Log("device flow failed: %v", err)
		return fmt.Errorf("device flow login failed: %w", err)
	}
	deviceLog.Log("device flow token received")

	return saveAndPrintToken(ctx, host, disc, disc.DeviceFlowClientID, token)
}

// requestDeviceCode sends the initial device authorization request (RFC 8628 §3.1).
func requestDeviceCode(ctx context.Context, endpoint, clientID string, scopes []string) (*deviceAuthResponse, error) {
	data := url.Values{
		"client_id": {clientID},
		"scope":     {strings.Join(scopes, " ")},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("building device authorization request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	resp, err := httputil.NewHTTPClient(10 * time.Second).Do(req)
	if err != nil {
		return nil, fmt.Errorf("device authorization HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp deviceTokenError
		_ = json.NewDecoder(resp.Body).Decode(&errResp)
		desc := errResp.ErrorDescription
		if desc == "" {
			desc = errResp.Error
		}
		if errResp.Error == "unauthorized_client" {
			deviceLog.Log("IdP rejected device code grant for client %s: %s", clientID, desc)
			return nil, fmt.Errorf("%s\n\nAsk your IdP administrator to add the device_code grant type\nto OAuth app %s", desc, clientID)
		}
		if desc != "" {
			return nil, fmt.Errorf("device authorization failed: %s", desc)
		}
		return nil, fmt.Errorf("device authorization endpoint returned HTTP %d", resp.StatusCode)
	}

	var authResp deviceAuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return nil, fmt.Errorf("parsing device authorization response: %w", err)
	}
	if authResp.DeviceCode == "" {
		return nil, fmt.Errorf("device authorization response missing device_code")
	}
	if authResp.VerificationURI == "" && authResp.VerificationURIComplete == "" {
		return nil, fmt.Errorf("device authorization response missing verification_uri")
	}
	authResp.Interval = clampInt(authResp.Interval, int(defaultPollInterval/time.Second), int(maxPollInterval/time.Second))
	if authResp.ExpiresIn <= 0 {
		authResp.ExpiresIn = defaultExpiresIn
	} else if authResp.ExpiresIn > maxExpiresIn {
		authResp.ExpiresIn = maxExpiresIn
	}
	return &authResp, nil
}

func clampInt(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// pollForToken polls the token endpoint until the user completes authorization (RFC 8628 §3.4–3.5).
func pollForToken(ctx context.Context, tokenURL, clientID string, authResp *deviceAuthResponse) (*oauth2.Token, error) {
	interval := time.Duration(authResp.Interval) * time.Second
	deadline := time.Now().Add(time.Duration(authResp.ExpiresIn) * time.Second)

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(interval):
		}

		if time.Now().After(deadline) {
			return nil, fmt.Errorf("device code expired — run 'glean auth login' to try again")
		}

		token, status, err := exchangeDeviceCode(ctx, tokenURL, clientID, authResp.DeviceCode)
		if err != nil {
			return nil, err
		}
		if status == pollSlowDown {
			interval += 5 * time.Second
			continue
		}
		if status == pollPending {
			continue
		}
		return token, nil
	}
}

type pollStatus int

const (
	pollDone pollStatus = iota
	pollPending
	pollSlowDown
)

// exchangeDeviceCode attempts a single token exchange for a device code.
func exchangeDeviceCode(ctx context.Context, tokenURL, clientID, deviceCode string) (*oauth2.Token, pollStatus, error) {
	data := url.Values{
		"grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
		"client_id":   {clientID},
		"device_code": {deviceCode},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, pollDone, fmt.Errorf("building token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	resp, err := httputil.NewHTTPClient(10 * time.Second).Do(req)
	if err != nil {
		return nil, pollDone, fmt.Errorf("token exchange HTTP request: %w", err)
	}
	defer resp.Body.Close()

	body := json.NewDecoder(resp.Body)

	if resp.StatusCode != http.StatusOK {
		var tokenErr deviceTokenError
		_ = body.Decode(&tokenErr)
		switch tokenErr.Error {
		case "authorization_pending":
			return nil, pollPending, nil
		case "slow_down":
			return nil, pollSlowDown, nil
		case "expired_token":
			return nil, pollDone, fmt.Errorf("device code expired — run 'glean auth login' to try again")
		case "access_denied":
			return nil, pollDone, fmt.Errorf("authorization denied by user")
		default:
			desc := tokenErr.ErrorDescription
			if desc == "" {
				desc = tokenErr.Error
			}
			return nil, pollDone, fmt.Errorf("token request failed: %s", desc)
		}
	}

	var tokenResp struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token,omitempty"`
		Scope        string `json:"scope,omitempty"`
	}
	if err := body.Decode(&tokenResp); err != nil {
		return nil, pollDone, fmt.Errorf("parsing token response: %w", err)
	}
	if tokenResp.AccessToken == "" {
		return nil, pollDone, fmt.Errorf("token response missing access_token")
	}

	token := &oauth2.Token{
		AccessToken:  tokenResp.AccessToken,
		TokenType:    tokenResp.TokenType,
		RefreshToken: tokenResp.RefreshToken,
	}
	if tokenResp.ExpiresIn > 0 {
		token.Expiry = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)
	}
	return token, pollDone, nil
}

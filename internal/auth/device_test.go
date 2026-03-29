package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequestDeviceCode_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("Content-Type"))
		_ = r.ParseForm()
		assert.Equal(t, "client-id", r.FormValue("client_id"))
		assert.Contains(t, r.FormValue("scope"), "openid")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"device_code":      "dev-code",
			"user_code":        "USER-1",
			"verification_uri": "https://idp.example/verify",
		})
	}))
	defer srv.Close()

	restore := overrideDiscoveryHTTPClient(srv.Client())
	defer restore()

	resp, err := requestDeviceCode(context.Background(), srv.URL, "client-id", []string{"openid", "profile"})
	require.NoError(t, err)
	assert.Equal(t, "dev-code", resp.DeviceCode)
	assert.Equal(t, "USER-1", resp.UserCode)
	assert.Equal(t, "https://idp.example/verify", resp.VerificationURI)
}

func TestRequestDeviceCode_UnauthorizedClient(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error":             "unauthorized_client",
			"error_description": "client cannot use this grant",
		})
	}))
	defer srv.Close()

	restore := overrideDiscoveryHTTPClient(srv.Client())
	defer restore()

	_, err := requestDeviceCode(context.Background(), srv.URL, "my-client", []string{"openid"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "client cannot use this grant")
	assert.Contains(t, err.Error(), "device_code grant type")
	assert.Contains(t, err.Error(), "my-client")
}

func TestRequestDeviceCode_MissingDeviceCode(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"user_code":        "U",
			"verification_uri": "https://x",
		})
	}))
	defer srv.Close()

	restore := overrideDiscoveryHTTPClient(srv.Client())
	defer restore()

	_, err := requestDeviceCode(context.Background(), srv.URL, "cid", []string{"s"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "missing device_code")
}

func TestRequestDeviceCode_IntervalAndExpiresIn(t *testing.T) {
	cases := []struct {
		name          string
		rawInterval   int
		rawExpiresIn  int
		wantInterval  int
		wantExpiresIn int
	}{
		{"interval_below_min_clamped_to_5", 1, 100, 5, 100},
		{"interval_above_max_clamped_to_60", 100, 100, 60, 100},
		{"interval_in_range_unchanged", 30, 100, 30, 100},
		{"expires_in_zero_defaults_to_900", 5, 0, 5, defaultExpiresIn},
		{"expires_in_capped_at_1800", 5, 99999, 5, maxExpiresIn},
		{"expires_in_in_range_unchanged", 5, 600, 5, 600},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(w).Encode(map[string]any{
					"device_code":      "dc",
					"verification_uri": "https://v",
					"interval":         tc.rawInterval,
					"expires_in":       tc.rawExpiresIn,
				})
			}))
			defer srv.Close()

			restore := overrideDiscoveryHTTPClient(srv.Client())
			defer restore()

			resp, err := requestDeviceCode(context.Background(), srv.URL, "cid", []string{"s"})
			require.NoError(t, err)
			assert.Equal(t, tc.wantInterval, resp.Interval)
			assert.Equal(t, tc.wantExpiresIn, resp.ExpiresIn)
		})
	}
}

func TestRequestDeviceCode_MissingVerificationURI(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"device_code": "dc",
		})
	}))
	defer srv.Close()

	restore := overrideDiscoveryHTTPClient(srv.Client())
	defer restore()

	_, err := requestDeviceCode(context.Background(), srv.URL, "cid", []string{"s"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "missing verification_uri")
}

func TestExchangeDeviceCode_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		_ = r.ParseForm()
		assert.Equal(t, "urn:ietf:params:oauth:grant-type:device_code", r.FormValue("grant_type"))
		assert.Equal(t, "cid", r.FormValue("client_id"))
		assert.Equal(t, "dev", r.FormValue("device_code"))
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"access_token": "tok-123",
			"token_type":   "Bearer",
			"expires_in":   3600,
		})
	}))
	defer srv.Close()

	restore := overrideDiscoveryHTTPClient(srv.Client())
	defer restore()

	tok, status, err := exchangeDeviceCode(context.Background(), srv.URL, "cid", "dev")
	require.NoError(t, err)
	assert.Equal(t, pollDone, status)
	require.NotNil(t, tok)
	assert.Equal(t, "tok-123", tok.AccessToken)
	assert.Equal(t, "Bearer", tok.TokenType)
}

func TestExchangeDeviceCode_AuthorizationPending(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "authorization_pending"})
	}))
	defer srv.Close()

	restore := overrideDiscoveryHTTPClient(srv.Client())
	defer restore()

	tok, status, err := exchangeDeviceCode(context.Background(), srv.URL, "cid", "dev")
	require.NoError(t, err)
	assert.Nil(t, tok)
	assert.Equal(t, pollPending, status)
}

func TestExchangeDeviceCode_SlowDown(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "slow_down"})
	}))
	defer srv.Close()

	restore := overrideDiscoveryHTTPClient(srv.Client())
	defer restore()

	tok, status, err := exchangeDeviceCode(context.Background(), srv.URL, "cid", "dev")
	require.NoError(t, err)
	assert.Nil(t, tok)
	assert.Equal(t, pollSlowDown, status)
}

func TestExchangeDeviceCode_AccessDenied(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "access_denied"})
	}))
	defer srv.Close()

	restore := overrideDiscoveryHTTPClient(srv.Client())
	defer restore()

	tok, status, err := exchangeDeviceCode(context.Background(), srv.URL, "cid", "dev")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "authorization denied")
	assert.Nil(t, tok)
	assert.Equal(t, pollDone, status)
}

func TestExchangeDeviceCode_EmptyAccessToken(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"access_token": "",
			"token_type":   "Bearer",
		})
	}))
	defer srv.Close()

	restore := overrideDiscoveryHTTPClient(srv.Client())
	defer restore()

	_, status, err := exchangeDeviceCode(context.Background(), srv.URL, "cid", "dev")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "missing access_token")
	assert.Equal(t, pollDone, status)
}

func TestPollForToken_PendingThenSuccess(t *testing.T) {
	var calls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		n := calls.Add(1)
		if n == 1 {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "authorization_pending"})
			return
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"access_token": "final",
			"token_type":   "Bearer",
		})
	}))
	defer srv.Close()

	restore := overrideDiscoveryHTTPClient(srv.Client())
	defer restore()

	auth := &deviceAuthResponse{
		DeviceCode:      "dc",
		Interval:        0,
		ExpiresIn:       60,
		VerificationURI: "https://v",
	}

	tok, err := pollForToken(context.Background(), srv.URL, "cid", auth)
	require.NoError(t, err)
	require.NotNil(t, tok)
	assert.Equal(t, "final", tok.AccessToken)
	assert.Equal(t, int32(2), calls.Load())
}

func TestPollForToken_SlowDownIncreasesInterval(t *testing.T) {
	if testing.Short() {
		t.Skip("timing-based test waits ~5s after slow_down")
	}
	var calls atomic.Int32
	var firstAt, secondAt atomic.Int64

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		n := calls.Add(1)
		now := time.Now().UnixNano()
		if n == 1 {
			firstAt.Store(now)
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "slow_down"})
			return
		}
		secondAt.Store(now)
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"access_token": "ok",
			"token_type":   "Bearer",
		})
	}))
	defer srv.Close()

	restore := overrideDiscoveryHTTPClient(srv.Client())
	defer restore()

	auth := &deviceAuthResponse{
		DeviceCode:      "dc",
		Interval:        0,
		ExpiresIn:       120,
		VerificationURI: "https://v",
	}

	start := time.Now()
	tok, err := pollForToken(context.Background(), srv.URL, "cid", auth)
	elapsed := time.Since(start)

	require.NoError(t, err)
	require.NotNil(t, tok)
	assert.Equal(t, "ok", tok.AccessToken)
	assert.Equal(t, int32(2), calls.Load())

	gap := time.Duration(secondAt.Load() - firstAt.Load())
	assert.GreaterOrEqual(t, gap, 4*time.Second,
		"expected ~5s wait after slow_down increased interval; gap=%v", gap)
	assert.GreaterOrEqual(t, elapsed, 4*time.Second)
}

func TestPollForToken_ContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	auth := &deviceAuthResponse{
		DeviceCode:      "dc",
		Interval:        10,
		ExpiresIn:       3600,
		VerificationURI: "https://v",
	}

	_, err := pollForToken(ctx, "http://unused.example/token", "cid", auth)
	require.Error(t, err)
	assert.ErrorIs(t, err, context.Canceled)
}

// overrideDiscoveryHTTPClient swaps the package-level client used by device/discovery helpers.
// The returned function restores the previous client (call in defer).
func overrideDiscoveryHTTPClient(cl *http.Client) func() {
	prev := discoveryHTTPClient
	discoveryHTTPClient = cl
	return func() { discoveryHTTPClient = prev }
}

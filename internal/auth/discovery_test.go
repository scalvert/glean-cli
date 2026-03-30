package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchProtectedResource_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/.well-known/oauth-protected-resource", r.URL.Path)
		json.NewEncoder(w).Encode(map[string]any{
			"resource":              "https://example.glean.com",
			"authorization_servers": []string{"https://auth.example.com"},
		})
	}))
	defer srv.Close()

	result, err := fetchProtectedResource(context.Background(), srv.URL)
	require.NoError(t, err)
	assert.Equal(t, []string{"https://auth.example.com"}, result.AuthorizationServers)
}

func TestFetchProtectedResource_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	_, err := fetchProtectedResource(context.Background(), srv.URL)
	require.Error(t, err)
	var notSupported *ErrOAuthNotSupported
	assert.True(t, errors.As(err, &notSupported))
}

func TestFetchProtectedResource_RateLimited(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer srv.Close()

	_, err := fetchProtectedResource(context.Background(), srv.URL)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "429")
	var notSupported *ErrOAuthNotSupported
	assert.False(t, errors.As(err, &notSupported))
}

func TestFetchProtectedResource_ServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	_, err := fetchProtectedResource(context.Background(), srv.URL)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "500")
	var notSupported *ErrOAuthNotSupported
	assert.False(t, errors.As(err, &notSupported))
}

func TestFetchProtectedResource_EmptyAuthorizationServers(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{
			"resource":              "https://example.glean.com",
			"authorization_servers": []string{},
		})
	}))
	defer srv.Close()

	_, err := fetchProtectedResource(context.Background(), srv.URL)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "OAuth metadata is incomplete")
	var notSupported *ErrOAuthNotSupported
	assert.False(t, errors.As(err, &notSupported))
}

func TestFetchProtectedResource_NullAuthorizationServers(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{
			"resource": "https://example.glean.com",
		})
	}))
	defer srv.Close()

	_, err := fetchProtectedResource(context.Background(), srv.URL)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "OAuth metadata is incomplete")
	var notSupported *ErrOAuthNotSupported
	assert.False(t, errors.As(err, &notSupported))
}

func TestErrOAuthNotSupported_ErrorMessage(t *testing.T) {
	err := &ErrOAuthNotSupported{URL: "https://example.com/.well-known/oauth-protected-resource"}
	assert.Equal(t, "OAuth protected resource metadata not found at https://example.com/.well-known/oauth-protected-resource", err.Error())
}

func TestRegisterClient_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "glean-cli", body["client_name"])
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{
			"client_id": "dyn-client-id",
		})
	}))
	defer srv.Close()

	cl, err := registerClient(context.Background(), srv.URL, "http://127.0.0.1:9999/callback")
	require.NoError(t, err)
	assert.Equal(t, "dyn-client-id", cl.ClientID)
}

func TestRegisterClient_WithSecret(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{
			"client_id":     "cid",
			"client_secret": "cs",
		})
	}))
	defer srv.Close()

	cl, err := registerClient(context.Background(), srv.URL, "http://127.0.0.1:9999/callback")
	require.NoError(t, err)
	assert.Equal(t, "cs", cl.ClientSecret)
}

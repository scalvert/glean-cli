package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLookupBackendURL_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "glean.com", body["emailDomain"])
		json.NewEncoder(w).Encode(map[string]any{
			"search_config": map[string]any{
				"queryURL": "https://scio-prod-be.glean.com/",
			},
		})
	}))
	defer srv.Close()

	url, err := lookupBackendURL(context.Background(), "steve@glean.com", srv.URL)
	require.NoError(t, err)
	// Trailing slash must be stripped
	assert.Equal(t, "https://scio-prod-be.glean.com", url)
}

func TestLookupBackendURL_EmptyResult(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"search_config": map[string]any{}})
	}))
	defer srv.Close()

	_, err := lookupBackendURL(context.Background(), "user@unknown.com", srv.URL)
	require.Error(t, err)
}

func TestExtractDomain(t *testing.T) {
	assert.Equal(t, "glean.com", extractDomain("steve@glean.com"))
	assert.Equal(t, "", extractDomain("notanemail"))
}

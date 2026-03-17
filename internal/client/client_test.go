package client

import (
	"testing"

	"github.com/gleanwork/glean-cli/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractInstance(t *testing.T) {
	tests := []struct {
		host     string
		expected string
	}{
		{"linkedin-be.glean.com", "linkedin"},
		{"linkedin", "linkedin"},
		{"custom.example.com", "custom"},
		{"acme-corp-be.glean.com", "acme-corp"},
		{"single", "single"},
		{"deep.sub.domain.example.com", "deep"},
	}

	for _, tt := range tests {
		t.Run(tt.host, func(t *testing.T) {
			got := extractInstance(tt.host)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestNew_EmptyHost(t *testing.T) {
	cfg := &config.Config{GleanHost: "", GleanToken: "some-token"}
	_, err := New(cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "host not configured")
}

func TestNew_EmptyToken(t *testing.T) {
	// auth.LoadOAuthToken will return "" for a fake host, so token stays empty
	cfg := &config.Config{GleanHost: "test-be.glean.com", GleanToken: ""}
	_, err := New(cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not authenticated")
}

func TestNew_Success(t *testing.T) {
	cfg := &config.Config{GleanHost: "test-be.glean.com", GleanToken: "valid-token"}
	client, err := New(cfg)
	require.NoError(t, err)
	assert.NotNil(t, client)
}

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gleanwork/api-client-go/models/components"
	"github.com/scalvert/glean-cli/internal/auth"
	"github.com/scalvert/glean-cli/internal/config"
)

var streamHTTPClient = &http.Client{Timeout: 120 * time.Second}

// StreamChat makes a streaming chat request to the Glean API, bypassing the
// SDK's buffered CreateStream method which consumes the entire response body
// before returning. Returns the raw response body for progressive line-by-line
// reading. The caller is responsible for closing the returned body.
//
// The response body contains newline-delimited JSON (NDJSON): each line is a
// complete components.ChatResponse JSON object.
func StreamChat(ctx context.Context, cfg *config.Config, msgs []components.ChatMessage) (io.ReadCloser, error) {
	host := cfg.GleanHost
	if host == "" {
		return nil, fmt.Errorf("Glean host not configured")
	}

	token := cfg.GleanToken
	if token == "" {
		token = auth.LoadOAuthToken(host)
	}
	if token == "" {
		return nil, fmt.Errorf("not authenticated — run 'glean auth login'")
	}

	// Expand short hostname to full form (mirrors extractInstance in reverse).
	if !strings.Contains(host, ".") {
		host += "-be.glean.com"
	}

	agentDefault := components.AgentEnumDefault
	modeDefault := components.ModeDefault
	stream := true
	body := components.ChatRequest{
		Messages:    msgs,
		AgentConfig: &components.AgentConfig{Agent: agentDefault.ToPointer(), Mode: modeDefault.ToPointer()},
		Stream:      &stream,
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("marshaling chat request: %w", err)
	}

	url := fmt.Sprintf("https://%s/rest/api/v1/chat", host)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("building request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Authorization", "Bearer "+token)

	if cfg.GleanEmail != "" {
		req.Header.Set("X-Scio-Actas", cfg.GleanEmail)
	}

	resp, err := streamHTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("chat request failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("chat request returned HTTP %d", resp.StatusCode)
	}

	return resp.Body, nil
}

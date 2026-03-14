package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gleanwork/api-client-go/models/components"
	"github.com/scalvert/glean-cli/internal/auth"
	"github.com/scalvert/glean-cli/internal/config"
)

var streamHTTPClient = &http.Client{Timeout: 120 * time.Second}

// StreamChat makes a streaming chat request to the Glean API, bypassing the
// SDK's buffered CreateStream which reads the entire response before returning.
// The req.Stream field is forced to true. The caller is responsible for closing
// the returned io.ReadCloser.
//
// The response body is NDJSON: each line is a complete ChatResponse JSON object.
// Only messages with messageType == "CONTENT" carry user-facing text; callers
// should skip UPDATE, CONTROL, DEBUG, etc.
func StreamChat(ctx context.Context, cfg *config.Config, req components.ChatRequest) (io.ReadCloser, error) {
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

	host = config.NormalizeHost(host)

	stream := true
	req.Stream = &stream

	// Ensure AgentConfig defaults are set if not provided.
	if req.AgentConfig == nil {
		agentDefault := components.AgentEnumDefault
		modeDefault := components.ModeDefault
		req.AgentConfig = &components.AgentConfig{
			Agent: agentDefault.ToPointer(),
			Mode:  modeDefault.ToPointer(),
		}
	}

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshaling chat request: %w", err)
	}

	url := fmt.Sprintf("https://%s/rest/api/v1/chat", host)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("building request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "text/event-stream")
	httpReq.Header.Set("Authorization", "Bearer "+token)

	if cfg.GleanEmail != "" {
		httpReq.Header.Set("X-Scio-Actas", cfg.GleanEmail)
	}

	resp, err := streamHTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("chat request failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("chat request returned HTTP %d", resp.StatusCode)
	}

	return resp.Body, nil
}

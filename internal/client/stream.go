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
	"github.com/gleanwork/glean-cli/internal/config"
	"github.com/gleanwork/glean-cli/internal/debug"
	"github.com/gleanwork/glean-cli/internal/httputil"
)

var streamLog = debug.New("stream:connect")

// streamTimeout is a generous timeout for long-running AUTO/ADVANCED agent
// responses. Context cancellation (ctrl+c in the TUI) handles user-initiated
// cancellation; this timeout is only a backstop for genuine network hangs.
const streamTimeout = 10 * time.Minute

// StreamChat makes a streaming chat request to the Glean API, bypassing the
// SDK's buffered CreateStream which reads the entire response before returning.
// The req.Stream field is forced to true. The caller is responsible for closing
// the returned io.ReadCloser.
//
// The response body is NDJSON: each line is a complete ChatResponse JSON object.
// Only messages with messageType == "CONTENT" carry user-facing text; callers
// should skip UPDATE, CONTROL, DEBUG, etc.
func StreamChat(ctx context.Context, cfg *config.Config, req components.ChatRequest) (io.ReadCloser, error) {
	serverURL := cfg.GleanServerURL
	if serverURL == "" {
		return nil, fmt.Errorf("glean server URL not configured")
	}

	token, authType := ResolveToken(cfg)
	if token == "" {
		return nil, fmt.Errorf("not authenticated — run 'glean auth login'")
	}

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

	url := serverURL + "/rest/api/v1/chat"
	streamLog.Log("POST %s (%d bytes)", url, len(payload))
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("building request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "text/event-stream")
	httpReq.Header.Set("Authorization", "Bearer "+token)
	if authType != "" {
		httpReq.Header.Set("X-Glean-Auth-Type", authType)
	}

	resp, err := httputil.NewHTTPClient(streamTimeout).Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("chat request failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		streamLog.Log("chat request failed: HTTP %d (body discarded)", resp.StatusCode)
		resp.Body.Close()
		return nil, fmt.Errorf("chat request returned HTTP %d", resp.StatusCode)
	}
	streamLog.Log("stream connected: HTTP %d", resp.StatusCode)

	return resp.Body, nil
}

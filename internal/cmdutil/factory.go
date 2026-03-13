// Package cmdutil provides shared helpers for building SDK-backed CLI commands.
// Every namespace command follows the same pattern:
//   - Optional --json flag accepts a raw JSON payload (overrides typed flags)
//   - Optional --output flag (json | ndjson | text)
//   - Optional --dry-run flag prints the request and exits
//   - Results are written to an io.Writer for testability
package cmdutil

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/scalvert/glean-cli/internal/output"
)

// RunFunc is a function that calls the SDK and returns a result to be serialised.
type RunFunc func(ctx context.Context) (any, error)

// RunCommand executes a SDK-backed command with standard --json / --output / --dry-run handling.
// reqBody is the request that WOULD be sent (shown on --dry-run).
// runFn is the actual SDK call, only invoked when dryRun is false.
func RunCommand(ctx context.Context, w io.Writer, reqBody any, dryRun bool, outputFormat string, runFn RunFunc) error {
	if dryRun {
		return output.WriteJSON(w, reqBody)
	}

	result, err := runFn(ctx)
	if err != nil {
		return err
	}

	return output.WriteFormatted(w, result, outputFormat, nil)
}

// ParseJSON unmarshals a JSON string into dst.
func ParseJSON(s string, dst any) error {
	if err := json.Unmarshal([]byte(s), dst); err != nil {
		return fmt.Errorf("invalid --json payload: %w", err)
	}
	return nil
}

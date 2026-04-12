package errors_test

import (
	"errors"
	"fmt"
	"testing"

	clierrors "github.com/gleanwork/glean-cli/internal/errors"
)

func TestCLIError_ErrorWithoutSuggestion(t *testing.T) {
	err := &clierrors.CLIError{
		UserMessage: "something went wrong",
		ExitCode:    clierrors.ExitGeneralError,
	}
	want := "something went wrong"
	if got := err.Error(); got != want {
		t.Errorf("Error() = %q, want %q", got, want)
	}
}

func TestCLIError_ErrorWithSuggestion(t *testing.T) {
	err := &clierrors.CLIError{
		UserMessage: "not authenticated",
		Suggestion:  "Run: glean auth login",
		ExitCode:    clierrors.ExitAuthError,
	}
	want := "not authenticated\n\n  Run: glean auth login"
	if got := err.Error(); got != want {
		t.Errorf("Error() = %q, want %q", got, want)
	}
}

func TestCLIError_Unwrap(t *testing.T) {
	cause := fmt.Errorf("token expired")
	err := &clierrors.CLIError{
		UserMessage: "auth failed",
		ExitCode:    clierrors.ExitAuthError,
		Cause:       cause,
	}
	if got := err.Unwrap(); got != cause {
		t.Errorf("Unwrap() = %v, want %v", got, cause)
	}
}

func TestCLIError_UnwrapNilCause(t *testing.T) {
	err := &clierrors.CLIError{
		UserMessage: "auth failed",
		ExitCode:    clierrors.ExitAuthError,
	}
	if got := err.Unwrap(); got != nil {
		t.Errorf("Unwrap() = %v, want nil", got)
	}
}

func TestCLIError_ErrorsAs(t *testing.T) {
	cause := fmt.Errorf("token expired")
	inner := &clierrors.CLIError{
		UserMessage: "auth failed",
		ExitCode:    clierrors.ExitAuthError,
		Cause:       cause,
	}
	wrapped := fmt.Errorf("command failed: %w", inner)

	var cliErr *clierrors.CLIError
	if !errors.As(wrapped, &cliErr) {
		t.Fatal("errors.As did not find CLIError in chain")
	}
	if cliErr.ExitCode != clierrors.ExitAuthError {
		t.Errorf("ExitCode = %d, want %d", cliErr.ExitCode, clierrors.ExitAuthError)
	}
}

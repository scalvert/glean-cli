package errors

import "fmt"

const (
	ExitSuccess      = 0
	ExitGeneralError = 1
	ExitUsageError   = 2
	ExitAuthError    = 3
	ExitNetworkError = 4
	ExitAPIError     = 5
	ExitRateLimited  = 6
	ExitTimeout      = 7
)

type CLIError struct {
	UserMessage string
	Suggestion  string
	ExitCode    int
	Cause       error
}

func (e *CLIError) Error() string {
	if e.Suggestion != "" {
		return fmt.Sprintf("%s\n\n  %s", e.UserMessage, e.Suggestion)
	}
	return e.UserMessage
}

func (e *CLIError) Unwrap() error { return e.Cause }

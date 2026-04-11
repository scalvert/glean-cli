// main.go
package main

import (
	"errors"
	"os"

	"github.com/gleanwork/glean-cli/cmd"
	clierrors "github.com/gleanwork/glean-cli/internal/errors"
	"github.com/gleanwork/glean-cli/internal/httputil"
)

// version is set at build time via ldflags: -X main.version=<version>
var version = "dev"

func main() {
	cmd.SetVersion(version)
	httputil.SetVersion(version)
	if err := cmd.Execute(); err != nil {
		var cliErr *clierrors.CLIError
		if errors.As(err, &cliErr) {
			os.Exit(cliErr.ExitCode)
		}
		os.Exit(1)
	}
}

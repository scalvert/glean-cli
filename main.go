// main.go
package main

import (
	"os"

	"github.com/gleanwork/glean-cli/cmd"
	"github.com/gleanwork/glean-cli/internal/httputil"
)

// version is set at build time via ldflags: -X main.version=<version>
var version = "dev"

func main() {
	cmd.SetVersion(version)
	httputil.SetVersion(version)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

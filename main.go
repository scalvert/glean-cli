// main.go
package main

import (
	"os"

	"github.com/scalvert/glean-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

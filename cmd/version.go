package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cliVersion = "dev"

// SetVersion stores the build-time version for use in the version command.
func SetVersion(v string) {
	cliVersion = v
}

// NewCmdVersion creates and returns the version command.
func NewCmdVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the glean CLI version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "glean version %s\n", cliVersion)
		},
	}
}

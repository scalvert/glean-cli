package cmd

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

// rootCmd is the root command for the Glean CLI
var rootCmd *cobra.Command

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "glean",
		Short: "Glean CLI - A command-line interface for Glean operations.",
		Long: heredoc.Doc(`
			Work seamlessly with Glean from your command line.

			To get started, run 'glean --help'.
		`),
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				fmt.Fprintf(os.Stderr, "Error displaying help: %v\n", err)
			}
		},
	}

	// Add all subcommands
	cmd.AddCommand(
		NewCmdAPI(),
		NewCmdConfig(),
		NewCmdGenerate(),
		NewCmdOpenAPISpec(),
		NewCmdSearch(),
	)

	return cmd
}

func init() {
	rootCmd = NewCmdRoot()
}

func Execute() error {
	return rootCmd.Execute()
}

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
		// Silence usage output when an error occurs
		SilenceUsage: true,
		// Ensure consistent error formatting
		SilenceErrors: true,
	}

	// Add all subcommands
	cmd.AddCommand(
		NewCmdAPI(),
		NewCmdConfig(),
		NewCmdGenerate(),
		NewCmdSearch(),
		NewCmdChat(),
	)

	// Propagate settings to all subcommands
	for _, subCmd := range cmd.Commands() {
		subCmd.SilenceUsage = true
		subCmd.SilenceErrors = true
	}

	return cmd
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}
	return nil
}

func init() {
	rootCmd = NewCmdRoot()
}

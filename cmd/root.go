// Package cmd implements the command-line interface for the Glean CLI.
// It provides commands for interacting with Glean's API, managing configuration,
// and performing various operations like search and chat.
package cmd

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
// It provides help information and coordinates all subcommands.
var rootCmd *cobra.Command

// NewCmdRoot creates and returns the root command for the Glean CLI.
// It sets up the base command structure and adds all subcommands.
func NewCmdRoot() *cobra.Command {
	var verbosity int

	cmd := &cobra.Command{
		Use:   "glean",
		Short: "Glean CLI - A command-line interface for Glean operations.",
		Long: heredoc.Doc(`
			Work seamlessly with Glean from your command line.

			To get started, run 'glean --help'.
		`),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Set debug level based on verbosity flag count
			if verbosity > 0 {
				if verbosity > 3 {
					verbosity = 3 // Cap at level 3
				}

				os.Setenv("GLEAN_HTTP_DEBUG", fmt.Sprintf("%d", verbosity))
			}
		},
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

	cmd.PersistentFlags().CountVarP(&verbosity, "verbose", "v", "Increase verbosity level (can be used multiple times)")

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

// Execute runs the root command and handles any errors.
// It's the main entry point for the CLI application.
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

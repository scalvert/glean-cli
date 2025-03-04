// Package cmd implements the command-line interface for the Glean CLI.
package cmd

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

// NewCmdGenerate creates and returns the generate command.
// The generate command provides utilities for generating various resources
// like OpenAPI specifications and other Glean-related assets.
func NewCmdGenerate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate resources or code for Glean",
		Long: heredoc.Doc(`
			Use this command to generate various resources,
			such as OpenAPI specs, configurations, or other Glean-related assets.
		`),
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				fmt.Fprintf(os.Stderr, "Error displaying help: %v\n", err)
			}
		},
	}

	cmd.AddCommand(NewCmdOpenAPISpec())

	return cmd
}

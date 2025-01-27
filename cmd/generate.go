package cmd

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
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

func init() {
	rootCmd.AddCommand(generateCmd)
}

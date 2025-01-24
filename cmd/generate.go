package cmd

import (
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
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

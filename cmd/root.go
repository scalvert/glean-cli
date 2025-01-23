// cmd/root.go
package cmd

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "glean",
	Short: "Glean CLI - A command-line interface for Glean operations.",
	Long: heredoc.Doc(`
		Work seamlessly with Glean from your command line.

		To get started, run 'glean --help'.
	`),
	Run: func(cmd *cobra.Command, args []string) {
		// If no subcommand is provided, show help by default
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main().
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Here you can define any persistent flags or configuration for the root command.
	// For example:
	// rootCmd.PersistentFlags().BoolP("toggle", "t", false, "Help message for toggle")

	// Or define local flags for the root command only (non-persistent).
	// rootCmd.Flags().BoolP("version", "v", false, "Display version")
}

// cmd/api.go
package cmd

import (
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

// apiCmd represents the 'glean api' command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Make calls to the Glean API",
	Long: heredoc.Doc(`
		Send requests to the Glean API.
		This command is similar to 'gh api', allowing you to interact with Glean endpoints directly.

		Usage:
		  glean api --method POST /endpoint
		  glean api --method GET /search
	`),

	// Example: glean api --method POST /endpoint
	RunE: func(cmd *cobra.Command, args []string) error {
		// Implementation for the 'api' command
		method, _ := cmd.Flags().GetString("method")
		endpoint := ""
		if len(args) > 0 {
			endpoint = args[0]
		}

		fmt.Printf("Invoking Glean API with method=%s, endpoint=%s\n", strings.ToUpper(method), endpoint)
		// You can expand this to actually perform an HTTP request here.
		return nil
	},
}

func init() {
	// Add apiCmd as a subcommand of rootCmd
	rootCmd.AddCommand(apiCmd)

	// Here we define flags. e.g., glean api --method GET /some/path
	apiCmd.Flags().StringP("method", "X", "GET", "HTTP method to use (GET, POST, etc.)")
}

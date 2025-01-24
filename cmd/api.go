package cmd

import (
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

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

	RunE: func(cmd *cobra.Command, args []string) error {
		method, _ := cmd.Flags().GetString("method")
		endpoint := ""
		if len(args) > 0 {
			endpoint = args[0]
		}

		fmt.Printf("Invoking Glean API with method=%s, endpoint=%s\n", strings.ToUpper(method), endpoint)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	apiCmd.Flags().StringP("method", "X", "GET", "HTTP method to use (GET, POST, etc.)")
}

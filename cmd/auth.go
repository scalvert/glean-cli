package cmd

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/gleanwork/glean-cli/internal/auth"
	"github.com/spf13/cobra"
)

// NewCmdAuth creates the `glean auth` command group.
func NewCmdAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate with Glean",
		Long: heredoc.Doc(`
			Manage authentication with your Glean instance.

			Use 'glean auth login' to authenticate via your browser (recommended).
			For CI/CD environments, set GLEAN_API_TOKEN instead.
		`),
	}
	cmd.AddCommand(newAuthLoginCmd(), newAuthLogoutCmd(), newAuthStatusCmd())
	return cmd
}

func newAuthLoginCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Authenticate with Glean via your browser",
		Long: heredoc.Doc(`
			Opens your browser to authenticate with Glean using OAuth 2.0 + PKCE.

			If your Glean instance supports OAuth with Dynamic Client Registration
			(most instances), no additional configuration is needed.

			For instances with a pre-registered OAuth app, configure first:
			  glean config --oauth-client-id <id>

			For CI/CD environments, set GLEAN_API_TOKEN instead of using this command.

			Examples:
			  glean auth login
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return auth.Login(cmd.Context())
		},
		SilenceUsage: true,
	}
}

func newAuthLogoutCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "Remove stored Glean credentials",
		Long: heredoc.Doc(`
			Removes the OAuth token stored by 'glean auth login'.

			To clear a static API token, use 'glean config --clear' instead.
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return auth.Logout(cmd.Context())
		},
		SilenceUsage: true,
	}
}

func newAuthStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show current authentication status",
		RunE: func(cmd *cobra.Command, args []string) error {
			return auth.Status(cmd.Context())
		},
		SilenceUsage: true,
	}
}

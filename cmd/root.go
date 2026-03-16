// Package cmd implements the command-line interface for the Glean CLI.
// It provides commands for interacting with Glean's API, managing configuration,
// and performing various operations like search and chat.
package cmd

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gleanwork/glean-cli/internal/auth"
	gleanClient "github.com/gleanwork/glean-cli/internal/client"
	"github.com/gleanwork/glean-cli/internal/config"
	"github.com/gleanwork/glean-cli/internal/tui"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
// It provides help information and coordinates all subcommands.
var rootCmd *cobra.Command

// NewCmdRoot creates and returns the root command for the Glean CLI.
// It sets up the base command structure and adds all subcommands.
func NewCmdRoot() *cobra.Command {
	var verbosity int
	var continueSession bool

	cmd := &cobra.Command{
		Use:   "glean",
		Short: "Glean CLI - A command-line interface for Glean operations.",
		Long: heredoc.Doc(`
			Work seamlessly with Glean from your command line.

			Running 'glean' with no arguments opens the full-screen chat TUI.
			Run 'glean --help' for other available commands.
		`),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			_ = verbosity // reserved for future debug logging
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.LoadConfig()
			if err != nil {
				return err
			}

			// Validate credentials before opening the TUI.
			if _, err := gleanClient.NewFromConfig(); err != nil {
				return err
			}

			// Always start with a fresh session. Prior conversation context
			// flashing back on startup is jarring and unexpected. Users who
			// want to continue a prior conversation can use --continue.
			session := &tui.Session{}
			if continueSession {
				session = tui.LoadLatest()
			}

			// Resolve identity: "email · host" or just host.
			// Email comes from stored token or decoded from the JWT access token
			// (Glean's RFC 8414 access tokens are JWTs containing the email claim).
			identity := cfg.GleanHost
			if tok, err := auth.LoadTokens(cfg.GleanHost); err == nil && tok != nil {
				email := tok.Email
				if email == "" {
					email = auth.EmailFromJWT(tok.AccessToken)
				}
				if email != "" {
					identity = email + "  ·  " + cfg.GleanHost
				}
			}

			model, err := tui.New(cfg, session, identity, cmd.Context())
			if err != nil {
				return fmt.Errorf("failed to create TUI: %w", err)
			}

			p := tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseCellMotion(), tea.WithContext(cmd.Context()))
			finalModel, err := p.Run()
			// Print session stats to stdout after TUI exits (like Claude Code).
			if m, ok := finalModel.(*tui.Model); ok && len(m.Session().Turns) > 0 {
				fmt.Println(m.StatsLine())
			}
			return err
		},
		// Silence usage output when an error occurs
		SilenceUsage: true,
		// Ensure consistent error formatting
		SilenceErrors: true,
	}

	cmd.PersistentFlags().CountVarP(&verbosity, "verbose", "v", "Increase verbosity level (can be used multiple times)")
	cmd.Flags().BoolVar(&continueSession, "continue", false, "Resume the most recent saved chat session")

	cmd.AddGroup(&cobra.Group{
		ID:    "core",
		Title: "Core Commands:",
	})
	cmd.AddGroup(&cobra.Group{
		ID:    "namespace",
		Title: "API Namespace Commands:",
	})

	for _, sub := range []*cobra.Command{
		NewCmdAuth(),
		NewCmdSearch(),
		NewCmdChat(),
		NewCmdConfig(),
		NewCmdAPI(),
		NewCmdSchema(),
		NewCmdVersion(),
		NewCmdGenerate(),
	} {
		sub.GroupID = "core"
		cmd.AddCommand(sub)
	}

	for _, sub := range []*cobra.Command{
		NewCmdActivity(),
		NewCmdAgents(),
		NewCmdAnnouncements(),
		NewCmdAnswers(),
		NewCmdCollections(),
		NewCmdDocuments(),
		NewCmdEntities(),
		NewCmdInsights(),
		NewCmdMessages(),
		NewCmdPins(),
		NewCmdShortcuts(),
		NewCmdTools(),
		NewCmdVerification(),
	} {
		sub.GroupID = "namespace"
		cmd.AddCommand(sub)
	}

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

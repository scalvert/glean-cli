// Package cmd implements the command-line interface for the Glean CLI.
// It provides commands for interacting with Glean's API, managing configuration,
// and performing various operations like search and chat.
package cmd

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	tea "github.com/charmbracelet/bubbletea"
	gleanClient "github.com/scalvert/glean-cli/internal/client"
	"github.com/scalvert/glean-cli/internal/tui"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
// It provides help information and coordinates all subcommands.
var rootCmd *cobra.Command

// NewCmdRoot creates and returns the root command for the Glean CLI.
// It sets up the base command structure and adds all subcommands.
func NewCmdRoot() *cobra.Command {
	var verbosity int
	var newSession bool

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
			sdk, err := gleanClient.NewFromConfig()
			if err != nil {
				return err
			}

			var session *tui.Session
			if newSession {
				session = &tui.Session{}
			} else {
				session = tui.LoadLatest()
			}

			model, err := tui.New(sdk, session, cmd.Context())
			if err != nil {
				return fmt.Errorf("failed to create TUI: %w", err)
			}

			p := tea.NewProgram(model, tea.WithAltScreen(), tea.WithContext(cmd.Context()))
			_, err = p.Run()
			return err
		},
		// Silence usage output when an error occurs
		SilenceUsage: true,
		// Ensure consistent error formatting
		SilenceErrors: true,
	}

	cmd.PersistentFlags().CountVarP(&verbosity, "verbose", "v", "Increase verbosity level (can be used multiple times)")
	cmd.Flags().BoolVar(&newSession, "new", false, "Start a new chat session (discard saved history)")

	// Add all subcommands
	cmd.AddCommand(
		NewCmdActivity(),
		NewCmdAgents(),
		NewCmdAnnouncements(),
		NewCmdAnswers(),
		NewCmdAPI(),
		NewCmdChat(),
		NewCmdCollections(),
		NewCmdConfig(),
		NewCmdDocuments(),
		NewCmdEntities(),
		NewCmdGenerate(),
		NewCmdInsights(),
		NewCmdMCP(),
		NewCmdMessages(),
		NewCmdPins(),
		NewCmdSchema(),
		NewCmdSearch(),
		NewCmdShortcuts(),
		NewCmdTools(),
		NewCmdVerification(),
		NewCmdVersion(),
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

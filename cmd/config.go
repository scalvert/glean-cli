package cmd

import (
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/scalvert/glean-cli/pkg/config"
	"github.com/spf13/cobra"
)

var (
	host  string
	token string
	clear bool
	show  bool
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure Glean CLI credentials",
	Long: heredoc.Doc(`
		Configure credentials for the Glean CLI.

		Examples:
		  # Set Glean host
		  glean config --host instance.glean.com

		  # Set Glean API token
		  glean config --token your-token

		  # Show current configuration
		  glean config --show

		  # Clear all stored credentials
		  glean config --clear
	`),
	RunE: func(cmd *cobra.Command, args []string) error {
		if show {
			cfg, err := config.LoadConfig()
			if err != nil {
				return fmt.Errorf("failed to access keyring: %w", err)
			}

			fmt.Println("Current configuration:")
			fmt.Printf("  %-10s %s\n", "Host:", valueOrNotSet(cfg.GleanHost))

			// Mask token if present
			tokenDisplay := "[not set]"
			if cfg.GleanToken != "" {
				tokenDisplay = cfg.GleanToken[0:4] + strings.Repeat("*", len(cfg.GleanToken)-4)
			}
			fmt.Printf("  %-10s %s\n", "Token:", tokenDisplay)
			return nil
		}

		if clear {
			if err := config.ClearConfig(); err != nil {
				return fmt.Errorf("failed to clear configuration: %w", err)
			}
			fmt.Println("Configuration cleared successfully")
			return nil
		}

		if host == "" && token == "" {
			return fmt.Errorf("no configuration provided. Use --host or --token to set configuration")
		}

		if err := config.SaveConfig(host, token); err != nil {
			return fmt.Errorf("failed to save configuration: %w", err)
		}

		fmt.Println("Configuration saved successfully")
		return nil
	},
}

func valueOrNotSet(value string) string {
	if value == "" {
		return "[not set]"
	}
	return value
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().StringVar(&host, "host", "", "Glean instance hostname (e.g., <instance>.glean.com)")
	configCmd.Flags().StringVar(&token, "token", "", "Glean API token")
	configCmd.Flags().BoolVar(&clear, "clear", false, "Clear all stored credentials")
	configCmd.Flags().BoolVar(&show, "show", false, "Show current configuration")
}

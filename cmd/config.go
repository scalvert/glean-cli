package cmd

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/scalvert/glean-cli/pkg/config"
	"github.com/spf13/cobra"
)

var (
	host  string
	token string
	clear bool
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure Glean CLI credentials",
	Long: heredoc.Doc(`
		Configure credentials for the Glean CLI.

		Examples:
		  # Set Glean host
		  glean config --host your-instance.glean.com

		  # Set Glean API token
		  glean config --token your-token

		  # Clear all stored credentials
		  glean config --clear
	`),
	RunE: func(cmd *cobra.Command, args []string) error {
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

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().StringVar(&host, "host", "", "Glean instance hostname (e.g., your-instance.glean.com)")
	configCmd.Flags().StringVar(&token, "token", "", "Glean API token")
	configCmd.Flags().BoolVar(&clear, "clear", false, "Clear all stored credentials")
}

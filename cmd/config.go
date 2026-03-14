package cmd

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/scalvert/glean-cli/internal/config"
	"github.com/spf13/cobra"
)

// notSetValue is displayed when a configuration value is not set
const notSetValue = "[not set]"

// ConfigOptions holds the configuration options for the config command.
// It allows setting or clearing Glean credentials and connection settings.
type ConfigOptions struct {
	host              string // Glean instance hostname
	port              string // Glean instance port
	token             string // API token for authentication
	email             string // User's email address
	oauthClientID     string // OAuth client ID for pre-registered apps
	oauthClientSecret string // OAuth client secret for confidential apps
	clear             bool   // Whether to clear all configuration
	show              bool   // Whether to display current configuration
}

// NewCmdConfig creates and returns the config command.
// The config command manages CLI configuration, including credentials
// and connection settings, with support for secure storage.
func NewCmdConfig() *cobra.Command {
	opts := ConfigOptions{}

	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage Glean CLI configuration",
		Long: heredoc.Doc(`
			Configure credentials for the Glean CLI.

			Examples:
			  # Set Glean host (either format works)
			  glean config --host linkedin
			  glean config --host linkedin-be.glean.com

			  # Set Glean host and port (e.g. custom proxy)
			  glean config --host foo.bar.com --port 7960

			  # Set Glean API token
			  glean config --token your-token

			  # Set Glean user email
			  glean config --email user@company.com

			  # Show current configuration
			  glean config --show

			  # Clear all stored credentials
			  glean config --clear
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.show {
				cfg, err := config.LoadConfig()
				if err != nil {
					return fmt.Errorf("failed to access keyring: %w", err)
				}

				fmt.Println("Current configuration:")
				fmt.Printf("  %-10s %s\n", "Host:", valueOrNotSet(cfg.GleanHost))
				fmt.Printf("  %-10s %s\n", "Port:", valueOrNotSet(cfg.GleanPort))
				fmt.Printf("  %-10s %s\n", "Email:", valueOrNotSet(cfg.GleanEmail))
				// Mask token if present
				tokenDisplay := notSetValue
				if cfg.GleanToken != "" {
					tokenDisplay = config.MaskToken(cfg.GleanToken)
				}
				fmt.Printf("  %-10s %s\n", "Token:", tokenDisplay)
				return nil
			}

			if opts.clear {
				if err := config.ClearConfig(); err != nil {
					return fmt.Errorf("failed to clear configuration: %w", err)
				}
				fmt.Println("Configuration cleared successfully")
				return nil
			}

			if opts.host == "" && opts.port == "" && opts.token == "" && opts.email == "" && opts.oauthClientID == "" && opts.oauthClientSecret == "" {
				return fmt.Errorf("no configuration provided. Use --host, --port, --token, --email, or --oauth-client-id to set configuration")
			}

			if opts.host != "" || opts.port != "" || opts.token != "" || opts.email != "" {
				if err := config.SaveConfig(opts.host, opts.port, opts.token, opts.email); err != nil {
					return fmt.Errorf("failed to save configuration: %w", err)
				}
			}

			if opts.oauthClientID != "" || opts.oauthClientSecret != "" {
				if err := config.SaveOAuthClient(opts.oauthClientID, opts.oauthClientSecret); err != nil {
					return fmt.Errorf("failed to save OAuth configuration: %w", err)
				}
			}

			fmt.Println("Configuration saved successfully")
			return nil
		},
	}

	cmd.Flags().StringVar(&opts.host, "host", "", "Glean instance name or full hostname (e.g., 'linkedin' or 'linkedin-be.glean.com')")
	cmd.Flags().StringVar(&opts.port, "port", "", "Port for custom proxy (only applies to 'glean api' command; SDK commands use standard HTTPS)")
	cmd.Flags().StringVar(&opts.token, "token", "", "Glean API token")
	cmd.Flags().StringVar(&opts.email, "email", "", "Email address for API requests")
	cmd.Flags().StringVar(&opts.oauthClientID, "oauth-client-id", "", "OAuth client ID (for instances with a pre-registered OAuth app)")
	cmd.Flags().StringVar(&opts.oauthClientSecret, "oauth-client-secret", "", "OAuth client secret (for confidential OAuth apps)")
	cmd.Flags().BoolVar(&opts.clear, "clear", false, "Clear all stored credentials")
	cmd.Flags().BoolVar(&opts.show, "show", false, "Show current configuration")

	return cmd
}

func valueOrNotSet(value string) string {
	if value == "" {
		return "[not set]"
	}
	return value
}

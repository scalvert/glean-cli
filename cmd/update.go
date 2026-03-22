package cmd

import (
	"fmt"

	"github.com/gleanwork/glean-cli/internal/update"
	"github.com/spf13/cobra"
)

func NewCmdUpdate() *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Update the glean CLI to the latest version",
		Long: `Check for a newer release of the glean CLI and install it.

If glean was installed via Homebrew, this runs:
  brew upgrade gleanwork/tap/glean-cli

Otherwise the latest binary is downloaded from GitHub Releases,
its SHA-256 checksum is verified, and the running binary is replaced.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Read cliVersion at call time, not construction time, so the
			// ldflags-injected version (set via SetVersion in main) is used.
			if err := update.Upgrade(cliVersion); err != nil {
				return fmt.Errorf("update failed: %w", err)
			}
			return nil
		},
	}
}

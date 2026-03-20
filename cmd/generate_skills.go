package cmd

import (
	"github.com/gleanwork/glean-cli/internal/skills"
	"github.com/spf13/cobra"
)

// NewCmdGenerateSkills creates the generate-skills command.
func NewCmdGenerateSkills() *cobra.Command {
	var outputDir string

	cmd := &cobra.Command{
		Use:   "generate-skills",
		Short: "Generate SKILL.md files from the CLI's schema registry",
		Long: `Generates Agent Skills (https://agentskills.io) from the CLI's own
schema registry. Each command gets a SKILL.md with frontmatter, flags,
subcommands, and examples — derived deterministically from the registered
schemas.

Run this after adding or modifying commands to keep skills in sync.

  glean generate-skills                    # write to skills/
  glean generate-skills --output-dir out/  # write to out/`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return skills.Generate(outputDir)
		},
	}

	cmd.Flags().StringVar(&outputDir, "output-dir", "skills", "Output directory for generated skills")

	return cmd
}

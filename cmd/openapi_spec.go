package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/scalvert/glean-cli/pkg/llm"
	"github.com/scalvert/glean-cli/pkg/output"
	"github.com/spf13/cobra"
)

var (
	inputFile  string
	outputFile string
	prompt     string
	model      string
	noColor    bool
)

func newOpenapiSpecCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "openapi-spec",
		Short: "Generate an OpenAPI spec from an API definition or curl command",
		Long: heredoc.Doc(`
			Generate an OpenAPI specification from an API definition or curl command.

			The input can be provided via a file or through stdin:
			  glean generate openapi-spec -f api.txt
			  echo "curl example.com/api" | glean generate openapi-spec

			The generated spec will be in YAML format and can be saved to a file:
			  glean generate openapi-spec -f api.txt -o spec.yaml

			Add custom instructions to guide the generation:
			  glean generate openapi-spec -f api.txt --prompt "Include rate limiting details"

			Use --no-color to disable colorized output.
		`),
		Example: heredoc.Doc(`
			# Generate from a curl command
			$ echo 'curl -X POST https://api.example.com/users \
			  -H "Content-Type: application/json" \
			  -d '"'"'{"name": "John", "email": "john@example.com"}'"'"'' | \
			  glean generate openapi-spec

			# Generate from a file with custom instructions
			$ glean generate openapi-spec -f api.txt --prompt "Include authentication details"

			# Save to a file
			$ glean generate openapi-spec -f api.txt -o spec.yaml
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			var input []byte
			var err error

			if inputFile == "-" || inputFile == "" {
				input, err = io.ReadAll(os.Stdin)
				if err != nil {
					return fmt.Errorf("failed to read from stdin: %w", err)
				}
				if len(input) == 0 {
					return fmt.Errorf("no input provided")
				}
			} else {
				input, err = os.ReadFile(inputFile)
				if err != nil {
					return fmt.Errorf("failed to read input file: %w", err)
				}
			}

			spec, err := llm.GenerateOpenAPISpec(string(input), prompt, model)
			if err != nil {
				return fmt.Errorf("failed to generate OpenAPI spec: %w", err)
			}

			if outputFile != "" {
				if err := os.WriteFile(outputFile, []byte(spec), 0644); err != nil {
					return fmt.Errorf("failed to write output file: %w", err)
				}
				fmt.Fprintf(cmd.OutOrStdout(), "OpenAPI spec written to %s\n", outputFile)
				return nil
			}

			return output.WriteString(cmd.OutOrStdout(), spec, output.Options{
				NoColor: noColor,
				Format:  "yaml",
			})
		},
	}

	cmd.Flags().StringVarP(&inputFile, "file", "f", "", "Input file containing the API/curl command (use \"-\" for stdin)")
	cmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file for the OpenAPI spec")
	cmd.Flags().StringVarP(&prompt, "prompt", "p", "", "Additional instructions for the LLM")
	cmd.Flags().StringVar(&model, "model", "gpt-4", "LLM model to use")
	cmd.Flags().BoolVar(&noColor, "no-color", false, "Disable colorized output")

	return cmd
}

func init() {
	generateCmd.AddCommand(newOpenapiSpecCmd())
}

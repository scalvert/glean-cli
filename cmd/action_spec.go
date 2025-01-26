package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/scalvert/glean-cli/pkg/llm"
	"github.com/spf13/cobra"
)

var (
	inputFile  string
	outputFile string
	prompt     string
	model      string
)

var actionSpecCmd = &cobra.Command{
	Use:   "action-spec",
	Short: "Generate an OpenAPI spec from an API definition or curl command",
	Long: heredoc.Doc(`
		This command reads an input (either stdin or a file) describing an API endpoint or
		a curl command and generates an OpenAPI spec, simplifying the creation of Glean Actions.

		Usage:
		  # Generate from a file
		  glean generate action-spec -f input.txt -o spec.yaml

		  # Generate from stdin
		  echo "curl example.com/api" | glean generate action-spec

		  # Add custom instructions
		  glean generate action-spec -f input.txt --prompt "Include rate limiting details"

		Input Format:
		  - A curl command
		  - An API description in a simple format

		Output:
		  OpenAPI specification in YAML format
	`),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var input string

		if inputFile != "" {
			data, err := os.ReadFile(inputFile)

			if err != nil {
				return fmt.Errorf("could not read file: %w", err)
			}
			input = string(data)
		} else {
			stat, _ := os.Stdin.Stat()

			if (stat.Mode() & os.ModeCharDevice) != 0 {
				return fmt.Errorf("no input found; please provide a file or pipe input to stdin")
			}

			scanner := bufio.NewScanner(os.Stdin)

			var lines []string

			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				return fmt.Errorf("error reading input: %w", err)
			}

			input = strings.Join(lines, "\n")
		}

		spec, err := llm.GenerateOpenAPISpec(input, prompt, model)
		if err != nil {
			return err
		}

		if outputFile != "" {
			if err := os.WriteFile(outputFile, []byte(spec), 0644); err != nil {
				return fmt.Errorf("failed to write output file: %w", err)
			}
			fmt.Printf("OpenAPI spec written to %s\n", outputFile)
		} else {
			fmt.Println(spec)
		}

		return nil
	},
}

func init() {
	generateCmd.AddCommand(actionSpecCmd)

	actionSpecCmd.Flags().StringVarP(&inputFile, "file", "f", "", "Input file containing the API/curl command")
	actionSpecCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file for the OpenAPI spec (defaults to stdout)")
	actionSpecCmd.Flags().StringVarP(&prompt, "prompt", "p", "", "Additional instructions for the LLM")
	actionSpecCmd.Flags().StringVar(&model, "model", "gpt-4", "LLM model to use")
}

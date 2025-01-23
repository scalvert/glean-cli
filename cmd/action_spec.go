// cmd/action_spec.go
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

var inputFile string

var actionSpecCmd = &cobra.Command{
	Use:   "action-spec",
	Short: "Generate an OpenAPI spec from an API definition or curl command",
	Long: heredoc.Doc(`
		This command reads an input (either stdin or a file) describing an API endpoint or
		a curl command and generates an OpenAPI spec, simplifying the creation of Glean Actions.

		Usage:
		  glean generate action-spec -f input.txt
		  echo "curl example.com/api" | glean generate action-spec

		Input Format:
		  - A curl command
		  - An API description in a simple format

		Output:
		  OpenAPI specification suitable for creating Glean Actions
	`),
	RunE: func(cmd *cobra.Command, args []string) error {
		var scanner *bufio.Scanner
		if inputFile != "" {
			f, err := os.Open(inputFile)
			if err != nil {
				return fmt.Errorf("could not open file: %w", err)
			}
			defer f.Close()
			scanner = bufio.NewScanner(f)
		} else {
			// Read from stdin
			stat, _ := os.Stdin.Stat()
			if (stat.Mode() & os.ModeCharDevice) != 0 {
				// No data is being piped in
				return fmt.Errorf("no input found; please provide a file or pipe input to stdin")
			}
			scanner = bufio.NewScanner(os.Stdin)
		}

		var inputLines []string
		for scanner.Scan() {
			line := scanner.Text()
			inputLines = append(inputLines, line)
		}
		if err := scanner.Err(); err != nil {
			return fmt.Errorf("error reading input: %w", err)
		}

		// For now, just print what was read
		fmt.Println("Generating OpenAPI spec from the following lines:")
		for _, line := range inputLines {
			fmt.Println(line)
		}

		// TODO: Actually parse the curl or API spec and convert to OpenAPI specification
		// This is just a placeholder.
		fmt.Println("\n(Placeholder) OpenAPI spec generated successfully.")
		return nil
	},
}

func init() {
	generateCmd.AddCommand(actionSpecCmd)

	// For example, glean generate action-spec -f input.txt
	actionSpecCmd.Flags().StringVarP(&inputFile, "file", "f", "", "Input file containing the API/curl command")
}

package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/MakeNowJust/heredoc"
	"github.com/briandowns/spinner"
	"github.com/scalvert/glean-cli/pkg/config"
	"github.com/scalvert/glean-cli/pkg/http"
	"github.com/scalvert/glean-cli/pkg/output"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

type APIOptions struct {
	method      string
	requestBody string
	inputFile   string
	preview     bool
	raw         bool
	noColor     bool
}

func NewCmdAPI() *cobra.Command {
	opts := APIOptions{}

	cmd := &cobra.Command{
		Use:   "api",
		Short: "Make an authenticated HTTP request to the Glean API",
		Long: heredoc.Doc(`
			Makes an authenticated HTTP request to the Glean API and prints the response.

			The endpoint argument should be a path of a Glean API endpoint. For example:
			  glean api search
			  glean api users/me

			The default HTTP request method is "GET". To use a different method,
			use the --method flag:
			  glean api --method POST search

			Request body can be provided via --raw-field, --field, or --input:
			  echo '{"query": "rust programming"}' | glean api --method POST search
			  glean api --method POST search --raw-field '{"query": "rust programming"}'
			  glean api --method POST search --input request.json

			Pass --preview to print the request details without actually sending it.
			Pass --no-color to disable colorized output (useful when piping to jq).
		`),
		Example: heredoc.Doc(`
			# Get the current user
			$ glean api users/me

			# Search with parameters
			$ glean api search --method POST --raw-field '{"query": "rust programming"}'

			# Search with parameters from a file
			$ glean api search --method POST --input search-params.json

			# Preview the request
			$ glean api search --method POST --raw-field '{"query": "test"}' --preview

			# Pipe to jq
			$ glean api search --no-color | jq .results
		`),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.LoadConfig()
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}

			client, err := http.NewClient(cfg)
			if err != nil {
				return err
			}

			endpoint := args[0]

			var body map[string]interface{}
			if opts.requestBody != "" {
				if jsonErr := json.Unmarshal([]byte(opts.requestBody), &body); jsonErr != nil {
					return jsonErr
				}
			}

			if opts.inputFile != "" {
				data, readErr := os.ReadFile(opts.inputFile)
				if readErr != nil {
					return readErr
				}
				if jsonErr := json.Unmarshal(data, &body); jsonErr != nil {
					return jsonErr
				}
			}

			if opts.inputFile == "" && opts.requestBody == "" {
				stdinData, readErr := io.ReadAll(os.Stdin)
				if readErr != nil {
					return readErr
				}
				if len(stdinData) > 0 {
					if jsonErr := json.Unmarshal(stdinData, &body); jsonErr != nil {
						return fmt.Errorf("invalid JSON from stdin: %w", jsonErr)
					}
				}
			}

			req := &http.Request{
				Method: opts.method,
				Path:   endpoint,
				Body:   body,
			}

			if opts.preview {
				return previewRequest(req, opts.noColor)
			}

			// Only show spinner if we're in a terminal and not using --raw or --no-color
			useSpinner := term.IsTerminal(int(os.Stderr.Fd())) && !opts.raw && !opts.noColor

			var s *spinner.Spinner
			if useSpinner {
				s = spinner.New(spinner.CharSets[9], 100*time.Millisecond)
				s.Suffix = " Making request to Glean API..."
				s.Writer = os.Stderr
				s.FinalMSG = "Request complete!\n"
				s.Start()
				defer s.Stop()
			}

			resp, err := client.SendRequest(req)
			if err != nil {
				return err
			}

			return output.Write(os.Stdout, resp, output.Options{
				NoColor: opts.raw || opts.noColor,
				Format:  "json",
			})
		},
	}

	cmd.Flags().StringVarP(&opts.method, "method", "X", "GET", "The HTTP method for the request")
	cmd.Flags().StringVar(&opts.requestBody, "raw-field", "", "Add a JSON string as the request body")
	cmd.Flags().StringVarP(&opts.inputFile, "input", "F", "", "The file to use as body for the request (use \"-\" to read from standard input)")
	cmd.Flags().BoolVar(&opts.preview, "preview", false, "Preview the API request without sending it")
	cmd.Flags().BoolVar(&opts.raw, "raw", false, "Print raw API response")
	cmd.Flags().BoolVar(&opts.noColor, "no-color", false, "Disable colorized output")

	return cmd
}

func previewRequest(req *http.Request, noColor bool) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	client, err := http.NewClient(cfg)
	if err != nil {
		return err
	}

	fmt.Printf("Request Method: %s\n", req.Method)
	fmt.Printf("Request URL: %s\n", client.GetFullURL(req.Path))

	fmt.Println("\nRequest Headers:")
	fmt.Printf("  Content-Type: application/json\n")
	if cfg.GleanToken != "" {
		fmt.Printf("  Authorization: Bearer %s\n", config.MaskToken(cfg.GleanToken))
	}
	if cfg.GleanEmail != "" {
		fmt.Printf("  X-Glean-User-Email: %s\n", cfg.GleanEmail)
		fmt.Printf("  X-Scio-Actas: %s\n", cfg.GleanEmail)
	}
	fmt.Printf("  X-Glean-Auth-Type: string\n")

	if req.Body != nil {
		fmt.Println("\nRequest Body:")
		bodyBytes, err := json.Marshal(req.Body)
		if err != nil {
			return fmt.Errorf("failed to format request body: %w", err)
		}

		return output.Write(os.Stdout, bodyBytes, output.Options{
			NoColor: noColor,
			Format:  "json",
		})
	}

	return nil
}

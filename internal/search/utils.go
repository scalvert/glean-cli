package search

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os/exec"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/scalvert/glean-cli/internal/http"
)

const (
	keyLowerQ = "q"
	keyUpperQ = "Q"
	keyQuit   = "ctrl+c"
	keyEsc    = "esc"
	keySpace  = " "
)

// GetTimezoneOffset returns the current timezone offset in minutes
func GetTimezoneOffset() int {
	_, offset := time.Now().Zone()
	return offset / 60
}

// AddFacetFilter adds a facet filter to the search options
func AddFacetFilter(opts *Options, fieldName string, values []string) {
	filter := FacetFilter{
		FieldName: fieldName,
		Values:    make([]FilterValue, len(values)),
	}
	for i, value := range values {
		filter.Values[i] = FilterValue{
			Value:        value,
			RelationType: "EQUALS",
		}
	}
	opts.RequestOptions.FacetFilters = append(opts.RequestOptions.FacetFilters, filter)
}

// performSearch executes a search request with the given parameters
func performSearch(client http.Client, opts *Options, cursor, trackingToken string) (*Response, error) {
	requestBody := map[string]interface{}{
		"query":             opts.Query,
		"pageSize":          opts.PageSize,
		"disableSpellcheck": opts.DisableSpellcheck,
		"maxSnippetSize":    opts.MaxSnippetSize,
		"timeoutMillis":     opts.TimeoutMillis,
	}

	// Add optional parameters if they're set
	if opts.InputDetails != nil {
		requestBody["inputDetails"] = opts.InputDetails
	}
	if len(opts.People) > 0 {
		requestBody["people"] = opts.People
	}
	if opts.RequestOptions != nil && len(opts.RequestOptions.FacetFilters) > 0 {
		requestBody["requestOptions"] = opts.RequestOptions
	}
	if len(opts.ResultTabIds) > 0 {
		requestBody["resultTabIds"] = opts.ResultTabIds
	}
	if opts.SessionInfo != nil {
		requestBody["sessionInfo"] = opts.SessionInfo
	}
	if opts.SourceDocument != nil {
		requestBody["sourceDocument"] = opts.SourceDocument
	}
	if opts.Timestamp != "" {
		requestBody["timestamp"] = opts.Timestamp
	}
	if cursor != "" {
		requestBody["cursor"] = cursor
	}
	if trackingToken != "" {
		requestBody["trackingToken"] = trackingToken
	}

	req := &http.Request{
		Method: "POST",
		Path:   "search",
		Body:   requestBody,
	}

	resp, err := client.SendRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error making search request: %w", err)
	}

	var searchResp Response
	if err := json.Unmarshal(resp, &searchResp); err != nil {
		return nil, fmt.Errorf("error parsing search response: %w", err)
	}

	return &searchResp, nil
}

// openURL opens a URL in the default browser after validating it
func openURL(urlStr string) tea.Cmd {
	// Validate URL
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil
	}

	// Only allow http/https URLs
	if u.Scheme != "http" && u.Scheme != "https" {
		return nil
	}

	return tea.ExecProcess(exec.Command("open", urlStr), nil)
}

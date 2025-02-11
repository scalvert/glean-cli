package search

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/scalvert/glean-cli/pkg/http"
	"github.com/scalvert/glean-cli/pkg/theme"
	"github.com/scalvert/glean-cli/pkg/utils"
)

// resultItem represents a search result in the list
type resultItem struct {
	title  string
	url    string
	desc   string
	source string
	index  int
}

func (i resultItem) Title() string {
	return fmt.Sprintf("%s %s | %s",
		theme.Blue(fmt.Sprint(i.index+1)),
		theme.Blue(utils.FormatDatasource(i.source)),
		theme.Bold(i.title),
	)
}

func (i resultItem) Description() string {
	return theme.Yellow(i.url) + "\n" + i.desc
}

func (i resultItem) FilterValue() string {
	return i.title
}

// newSearchList creates a new list model with consistent styling
func newSearchList() list.Model {
	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(theme.GleanBlue.ToLipgloss()).
		Foreground(theme.GleanBlue.ToLipgloss()).
		Bold(true).
		Padding(0, 0, 0, 1)
	delegate.Styles.SelectedDesc = delegate.Styles.SelectedDesc.
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(theme.GleanBlue.ToLipgloss()).
		Foreground(theme.GleanYellow.ToLipgloss()).
		Padding(0, 0, 0, 1)
	delegate.Styles.NormalTitle = delegate.Styles.NormalTitle.
		Foreground(theme.GleanBlue.ToLipgloss()).
		Bold(true)
	delegate.Styles.NormalDesc = delegate.Styles.NormalDesc.
		Foreground(theme.GleanYellow.ToLipgloss())

	l := list.New([]list.Item{}, delegate, 0, 0)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.SetShowTitle(true)
	l.Title = "Search Results"
	l.Styles.Title = lipgloss.NewStyle().
		Foreground(theme.GleanBlue.ToLipgloss()).
		Bold(true).
		Padding(0, 1).
		MarginLeft(1)
	l.Styles.FilterCursor = lipgloss.NewStyle().
		Foreground(theme.GleanBlue.ToLipgloss())

	return l
}

// updateListWithResults updates a list model with search results
func updateListWithResults(l *list.Model, response *Response, existingItems []list.Item, query string) {
	if response == nil {
		return
	}

	baseIndex := len(existingItems)
	items := make([]list.Item, 0, len(response.Results))

	for i, result := range response.Results {
		var snippets []string
		for _, snippet := range result.Snippets {
			snippets = append(snippets, snippet.Text)
		}
		items = append(items, resultItem{
			index:  baseIndex + i,
			title:  result.Document.Title,
			url:    result.Document.URL,
			desc:   strings.Join(snippets, "\n"),
			source: result.Document.Datasource,
		})
	}

	if len(existingItems) == 0 {
		l.SetItems(items)
	} else {
		l.SetItems(append(existingItems, items...))
	}

	updateListTitle(l, response, query)
}

// updateListTitle updates the list title based on search response
func updateListTitle(l *list.Model, response *Response, query string) {
	if response.SuggestedSpellCorrectedQuery != "" {
		l.Title = fmt.Sprintf("Did you mean: %s?", response.SuggestedSpellCorrectedQuery)
	} else if response.RewrittenQuery != "" {
		l.Title = fmt.Sprintf("Showing results for: %s", response.RewrittenQuery)
	} else {
		l.Title = fmt.Sprintf("Search Results for: %s", query)
	}
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

package search

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os/exec"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/scalvert/glean-cli/pkg/http"
	"github.com/scalvert/glean-cli/pkg/theme"
	"github.com/scalvert/glean-cli/pkg/utils"
)

const (
	keyLowerQ = "q"
	keyUpperQ = "Q"
	keyQuit   = "ctrl+c"
	keyEsc    = "esc"
	keySpace  = " "
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
	if response == nil || response.Results == nil {
		return
	}

	baseIndex := len(existingItems)
	items := make([]list.Item, 0, len(response.Results))

	for i, result := range response.Results {
		if result.Document != nil {
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
		} else if len(result.StructuredResults) > 0 {
			for _, sr := range result.StructuredResults {
				if person, ok := sr.(map[string]interface{})["person"]; ok {
					if p, ok := person.(map[string]interface{}); ok {
						name := p["name"].(string)
						metadata := p["metadata"].(map[string]interface{})
						title := metadata["title"].(string)
						email := metadata["email"].(string)
						bio := ""
						if b, ok := metadata["bio"]; ok && b != nil {
							bio = b.(string)
						}
						profileUrl := fmt.Sprintf("https://app.glean.com/directory/people/profile?person=%s", email)

						items = append(items, resultItem{
							index:  baseIndex + i,
							title:  title,
							url:    profileUrl,
							desc:   bio,
							source: name,
						})
					}
				}
			}
		}
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
	switch {
	case response.SuggestedSpellCorrectedQuery != "":
		l.Title = fmt.Sprintf("Did you mean: %s?", response.SuggestedSpellCorrectedQuery)
	case response.RewrittenQuery != "":
		l.Title = fmt.Sprintf("Showing results for: %s", response.RewrittenQuery)
	default:
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

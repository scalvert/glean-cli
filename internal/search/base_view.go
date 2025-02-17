package search

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/scalvert/glean-cli/internal/http"
	"github.com/scalvert/glean-cli/internal/theme"
)

type searchResultMsg struct {
	response *Response
}

type errMsg struct {
	err error
}

type BaseSearchModel struct {
	list     list.Model
	err      error
	client   http.Client
	response *Response
	opts     *Options
	spinner  spinner.Model
	loading  bool
	showMore bool
}

func (m *BaseSearchModel) performSearch(cursor, trackingToken string) tea.Msg {
	m.loading = true
	resp, err := performSearch(m.client, m.opts, cursor, trackingToken)
	if err != nil {
		return errMsg{err}
	}
	return searchResultMsg{resp}
}

func (m *BaseSearchModel) loadMore() tea.Msg {
	if m.response == nil {
		return nil
	}
	resp, err := performSearch(m.client, m.opts, m.response.Cursor, m.response.TrackingToken)
	if err != nil {
		return errMsg{err}
	}
	return searchResultMsg{resp}
}

func (m *BaseSearchModel) errorView() string {
	if m.err != nil {
		return fmt.Sprintf("\nError: %v\nPress any key to exit", m.err)
	}
	return ""
}

func (m *BaseSearchModel) loadingView() string {
	if m.loading {
		return fmt.Sprintf("\n  Searching company knowledge... %s\n", m.spinner.View())
	}
	return ""
}

func (m *BaseSearchModel) loadMorePrompt() string {
	if m.showMore && m.response != nil && m.response.HasMoreResults {
		return "\n  Press SPACE to load more results, q to quit\n"
	}
	return ""
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

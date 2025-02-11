package search

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/scalvert/glean-cli/pkg/http"
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

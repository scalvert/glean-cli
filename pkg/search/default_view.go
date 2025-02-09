package search

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/scalvert/glean-cli/pkg/http"
	"github.com/scalvert/glean-cli/pkg/theme"
)

// searchDisplayModel represents a simple display model for non-interactive search
type searchDisplayModel struct {
	BaseSearchModel
}

func newSearchDisplayModel(opts *Options, client http.Client) *searchDisplayModel {
	s := spinner.New()
	s.Spinner = spinner.MiniDot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	l := newSearchList()
	return &searchDisplayModel{
		BaseSearchModel: BaseSearchModel{
			list:     l,
			spinner:  s,
			opts:     opts,
			client:   client,
			loading:  true,
			showMore: false,
		},
	}
}

func (m *searchDisplayModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		func() tea.Msg {
			return m.performSearch("", "")
		},
	)
}

func (m *searchDisplayModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "Q", "ctrl+c", "esc":
			return m, tea.Quit
		case " ": // Space to load more
			if !m.loading && m.response != nil && m.response.HasMoreResults {
				m.loading = true
				m.showMore = false
				return m, tea.Batch(
					m.spinner.Tick,
					m.loadMore,
				)
			}
		}

		// Handle enter key to open URL
		if msg.Type == tea.KeyEnter {
			if i, ok := m.list.SelectedItem().(resultItem); ok {
				return m, tea.ExecProcess(exec.Command("open", i.url), nil)
			}
		}

	case tea.WindowSizeMsg:
		h, v := lipgloss.NewStyle().Margin(1, 2).GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case searchResultMsg:
		m.loading = false
		m.showMore = true
		m.response = msg.response
		updateListWithResults(&m.list, m.response, m.list.Items(), m.opts.Query)
		return m, nil

	case errMsg:
		m.err = msg.err
		return m, tea.Quit

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *searchDisplayModel) View() string {
	if errView := m.errorView(); errView != "" {
		return errView
	}

	if loadingView := m.loadingView(); loadingView != "" {
		return loadingView
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\n  Search query: %s\n", theme.Bold(m.opts.Query)))
	sb.WriteString(m.list.View() + "\n")
	sb.WriteString(m.loadMorePrompt())

	return sb.String()
}

// RunSearch executes a non-interactive search with the given options
func RunSearch(opts *Options, client http.Client) error {
	p := tea.NewProgram(newSearchDisplayModel(opts, client))
	_, err := p.Run()
	return err
}

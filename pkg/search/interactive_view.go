package search

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/scalvert/glean-cli/pkg/http"
	"github.com/scalvert/glean-cli/pkg/theme"
)

// searchInteractiveModel represents the UI state for interactive search
type searchInteractiveModel struct {
	BaseSearchModel
	input     textinput.Model
	searching bool
}

func newSearchInteractiveModel(opts *Options, client http.Client) *searchInteractiveModel {
	s := spinner.New()
	s.Spinner = spinner.MiniDot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	l := newSearchList()

	ti := textinput.New()
	ti.Placeholder = "Enter search query..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50

	if opts.Query != "" {
		ti.SetValue(opts.Query)
	}

	return &searchInteractiveModel{
		BaseSearchModel: BaseSearchModel{
			list:    l,
			spinner: s,
			opts:    opts,
			client:  client,
			loading: false,
		},
		input:     ti,
		searching: opts.Query == "",
	}
}

func (m *searchInteractiveModel) Init() tea.Cmd {
	if m.opts.Query != "" {
		return tea.Batch(
			m.spinner.Tick,
			func() tea.Msg {
				return m.performSearch("", "")
			},
		)
	}
	return textinput.Blink
}

func (m *searchInteractiveModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.searching {
			switch msg.Type {
			case tea.KeyEnter:
				if m.input.Value() != "" {
					m.searching = false
					m.opts.Query = m.input.Value()
					return m, tea.Batch(
						m.spinner.Tick,
						func() tea.Msg {
							return m.performSearch("", "")
						},
					)
				}
			case tea.KeyCtrlC, tea.KeyEsc:
				return m, tea.Quit
			}

			var cmd tea.Cmd
			m.input, cmd = m.input.Update(msg)
			return m, cmd
		}

		// Not searching, handle list navigation
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

		if msg.String() == "/" {
			m.searching = true
			m.input.Focus()
			return m, textinput.Blink
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

	if m.searching {
		var cmd tea.Cmd
		m.input, cmd = m.input.Update(msg)
		return m, cmd
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *searchInteractiveModel) View() string {
	if errView := m.errorView(); errView != "" {
		return errView
	}

	if m.searching {
		return fmt.Sprintf("\n  Search: %s\n\n", m.input.View())
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

// RunInteractiveSearch executes an interactive search with the given options
func RunInteractiveSearch(opts *Options, client http.Client) error {
	p := tea.NewProgram(newSearchInteractiveModel(opts, client))
	_, err := p.Run()
	return err
}

package tui

import "github.com/charmbracelet/bubbles/key"

// keyMap defines all keyboard shortcuts for the chat TUI.
type keyMap struct {
	Send       key.Binding
	Newline    key.Binding
	Quit       key.Binding
	Clear      key.Binding
	NewSession key.Binding
	Help       key.Binding
	ScrollUp   key.Binding
	ScrollDown key.Binding
}

var keys = keyMap{
	Send: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "send message"),
	),
	Newline: key.NewBinding(
		key.WithKeys("shift+enter"),
		key.WithHelp("shift+enter", "new line"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "esc"),
		key.WithHelp("ctrl+c/esc", "quit"),
	),
	Clear: key.NewBinding(
		key.WithKeys("ctrl+l"),
		key.WithHelp("ctrl+l", "clear screen"),
	),
	NewSession: key.NewBinding(
		key.WithKeys("ctrl+r"),
		key.WithHelp("ctrl+r", "new session"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	ScrollUp: key.NewBinding(
		key.WithKeys("pgup"),
		key.WithHelp("pgup", "scroll up"),
	),
	ScrollDown: key.NewBinding(
		key.WithKeys("pgdown"),
		key.WithHelp("pgdown", "scroll down"),
	),
}

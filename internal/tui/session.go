// Package tui provides the full-screen chat TUI for the default glean invocation.
package tui

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Turn holds one exchange in the conversation history.
type Turn struct {
	Role    string   `json:"role"`              // "user" or "assistant"
	Content string   `json:"content"`           // full text of the message
	Sources []Source `json:"sources,omitempty"` // cited documents
	Elapsed string   `json:"elapsed,omitempty"` // response time, e.g. "4.2s"
}

// Source is a cited document reference shown below an AI response.
type Source struct {
	Title      string `json:"title"`
	URL        string `json:"url"`
	Datasource string `json:"datasource"`
}

// Session holds conversation history and can be persisted to disk.
type Session struct {
	Turns []Turn `json:"turns"`
}

// sessionsDir returns ~/.glean/sessions/.
func sessionsDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".glean", "sessions"), nil
}

// LoadLatest loads the last saved session, or returns an empty session if none exists.
func LoadLatest() *Session {
	dir, err := sessionsDir()
	if err != nil {
		return &Session{}
	}
	path := filepath.Join(dir, "latest.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return &Session{}
	}
	var s Session
	if err := json.Unmarshal(data, &s); err != nil {
		return &Session{}
	}
	return &s
}

// Save persists the session to ~/.glean/sessions/latest.json.
func (s *Session) Save() error {
	dir, err := sessionsDir()
	if err != nil {
		return fmt.Errorf("could not locate sessions dir: %w", err)
	}
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("could not create sessions dir: %w", err)
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "latest.json"), data, 0600)
}

// AddTurn appends a turn to the session and saves immediately.
func (s *Session) AddTurn(role, content string, sources []Source) {
	s.Turns = append(s.Turns, Turn{Role: role, Content: content, Sources: sources})
	_ = s.Save() // best-effort
}

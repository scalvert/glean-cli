// Package tui provides the full-screen chat TUI for the default glean invocation.
package tui

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gleanwork/glean-cli/internal/debug"
)

var sessionLog = debug.New("session:persist")

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
	path  string // resolved path to the session file
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
// It reads JSONL format first, falling back to legacy JSON if no JSONL file exists.
func LoadLatest() *Session {
	dir, err := sessionsDir()
	if err != nil {
		sessionLog.Log("load: sessions dir error: %v", err)
		return &Session{}
	}

	jsonlPath := filepath.Join(dir, "latest.jsonl")
	if s, ok := loadJSONL(jsonlPath); ok {
		return s
	}

	jsonPath := filepath.Join(dir, "latest.json")
	if s, ok := migrateFromJSON(jsonPath, jsonlPath); ok {
		return s
	}

	return &Session{}
}

func loadJSONL(path string) (*Session, bool) {
	f, err := os.Open(path)
	if err != nil {
		return nil, false
	}
	defer f.Close()

	var turns []Turn
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var turn Turn
		if err := json.Unmarshal(scanner.Bytes(), &turn); err != nil {
			sessionLog.Log("load: skipping malformed line: %v", err)
			continue
		}
		turns = append(turns, turn)
	}
	if err := scanner.Err(); err != nil {
		sessionLog.Log("load: scanner error: %v", err)
	}
	sessionLog.Log("loaded %d turns from %s", len(turns), path)
	return &Session{Turns: turns, path: path}, true
}

func migrateFromJSON(jsonPath, jsonlPath string) (*Session, bool) {
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, false
	}
	var legacy Session
	if err := json.Unmarshal(data, &legacy); err != nil {
		sessionLog.Log("load: legacy parse error: %v", err)
		return nil, false
	}
	sessionLog.Log("migrating %d turns from %s to %s", len(legacy.Turns), jsonPath, jsonlPath)

	legacy.path = jsonlPath
	for _, turn := range legacy.Turns {
		if err := appendTurnToFile(jsonlPath, turn); err != nil {
			sessionLog.Log("migration write error: %v", err)
			return &legacy, true
		}
	}
	os.Remove(jsonPath)
	return &legacy, true
}

func appendTurnToFile(path string, turn Turn) error {
	data, err := json.Marshal(turn)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(append(data, '\n'))
	return err
}

// ensurePath resolves and caches the session file path.
func (s *Session) ensurePath() (string, error) {
	if s.path != "" {
		return s.path, nil
	}
	dir, err := sessionsDir()
	if err != nil {
		return "", fmt.Errorf("could not locate sessions dir: %w", err)
	}
	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", fmt.Errorf("could not create sessions dir: %w", err)
	}
	s.path = filepath.Join(dir, "latest.jsonl")
	return s.path, nil
}

// AddTurn appends a turn to the session and persists it immediately.
func (s *Session) AddTurn(role, content string, sources []Source) error {
	return s.AppendTurn(Turn{Role: role, Content: content, Sources: sources})
}

// AppendTurn appends a complete Turn to the session and persists it immediately.
func (s *Session) AppendTurn(turn Turn) error {
	s.Turns = append(s.Turns, turn)
	path, err := s.ensurePath()
	if err != nil {
		sessionLog.Log("save failed: %v", err)
		return err
	}
	if err := appendTurnToFile(path, turn); err != nil {
		sessionLog.Log("save failed: %v", err)
		return err
	}
	return nil
}

// Save rewrites the entire session to disk. Used only for migration or
// exceptional cases — normal operation uses AppendTurn for O(1) writes.
func (s *Session) Save() error {
	path, err := s.ensurePath()
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, turn := range s.Turns {
		data, err := json.Marshal(turn)
		if err != nil {
			return err
		}
		if _, err := f.Write(append(data, '\n')); err != nil {
			return err
		}
	}
	return nil
}

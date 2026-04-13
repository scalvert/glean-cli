// Package tui provides the full-screen chat TUI for the default glean invocation.
package tui

import (
	"bufio"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
	id    string // session identifier (UUID)
}

// ID returns the session's unique identifier.
func (s *Session) ID() string { return s.id }

func newSessionID() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		panic("crypto/rand failed: " + err.Error())
	}
	b[6] = (b[6] & 0x0f) | 0x40 // version 4
	b[8] = (b[8] & 0x3f) | 0x80 // variant 1
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

// sessionsDir returns ~/.glean/sessions/.
func sessionsDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".glean", "sessions"), nil
}

// LoadLatest loads the most recently modified session, or returns an empty
// session if none exists. Session files are identified by mtime.
func LoadLatest() *Session {
	dir, err := sessionsDir()
	if err != nil {
		sessionLog.Log("load: sessions dir error: %v", err)
		return &Session{}
	}

	path, id := findLatestSession(dir)
	if path == "" {
		return &Session{}
	}

	s, ok := loadJSONL(path)
	if !ok {
		return &Session{}
	}
	s.id = id
	return s
}

func findLatestSession(dir string) (path, id string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		sessionLog.Log("load: %v", err)
		return "", ""
	}

	var latestPath string
	var latestID string
	var latestMtime int64

	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".jsonl") {
			continue
		}
		info, err := e.Info()
		if err != nil {
			continue
		}
		if mtime := info.ModTime().UnixNano(); mtime > latestMtime {
			latestMtime = mtime
			latestPath = filepath.Join(dir, e.Name())
			latestID = strings.TrimSuffix(e.Name(), ".jsonl")
		}
	}

	if latestPath == "" {
		// Fallback: check for legacy latest.json and migrate
		jsonPath := filepath.Join(dir, "latest.json")
		jsonlPath := filepath.Join(dir, "latest.jsonl")
		if s, ok := migrateFromJSON(jsonPath, jsonlPath); ok {
			// Re-save as a proper UUID-named file
			newID := newSessionID()
			newPath := filepath.Join(dir, newID+".jsonl")
			if err := os.Rename(jsonlPath, newPath); err == nil {
				s.path = newPath
				s.id = newID
				sessionLog.Log("migrated legacy session to %s", newPath)
				return newPath, newID
			}
			return jsonlPath, "latest"
		}
	}

	return latestPath, latestID
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
	if s.id == "" {
		s.id = newSessionID()
	}
	s.path = filepath.Join(dir, s.id+".jsonl")
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

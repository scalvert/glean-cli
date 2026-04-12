package auth

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gleanwork/glean-cli/internal/fileutil"
)

// StoredTokens holds persisted OAuth tokens for a Glean host.
type StoredTokens struct {
	AccessToken   string    `json:"access_token"`
	RefreshToken  string    `json:"refresh_token,omitempty"`
	Expiry        time.Time `json:"expiry,omitempty"`
	Email         string    `json:"email,omitempty"`
	TokenType     string    `json:"token_type,omitempty"`
	TokenEndpoint string    `json:"token_endpoint,omitempty"` // used for token refresh
}

// IsExpired returns true if the token expires within the next 60 seconds.
func (t *StoredTokens) IsExpired() bool {
	if t.Expiry.IsZero() {
		return false
	}
	return time.Now().Add(60 * time.Second).After(t.Expiry)
}

// StoredClient holds a registered or configured OAuth client for a Glean host.
type StoredClient struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret,omitempty"`
}

// stateDir returns ~/.local/state/glean-cli/<hash>/ for the given host.
func stateDir(host string) string {
	h := sha256.Sum256([]byte(host))
	key := fmt.Sprintf("%x", h[:8])
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".local", "state", "glean-cli", key)
}

func tokensPath(host string) string { return filepath.Join(stateDir(host), "tokens.json") }
func clientPath(host string) string { return filepath.Join(stateDir(host), "client.json") }

func ensureDir(host string) error {
	dir := stateDir(host)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	return os.Chmod(dir, 0700)
}

func writeJSON(path string, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return fileutil.WriteFileAtomic(path, data, 0600)
}

func readJSON(path string, v any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// SaveTokens persists OAuth tokens for the given host.
func SaveTokens(host string, tok *StoredTokens) error {
	if err := ensureDir(host); err != nil {
		return fmt.Errorf("creating state dir: %w", err)
	}
	return writeJSON(tokensPath(host), tok)
}

// LoadTokens returns stored tokens for the given host, or nil if none exist.
func LoadTokens(host string) (*StoredTokens, error) {
	var tok StoredTokens
	if err := readJSON(tokensPath(host), &tok); err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	return &tok, nil
}

// DeleteTokens removes stored tokens for the given host.
func DeleteTokens(host string) error {
	err := os.Remove(tokensPath(host))
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

// SaveClient persists an OAuth client registration for the given host.
func SaveClient(host string, cl *StoredClient) error {
	if err := ensureDir(host); err != nil {
		return fmt.Errorf("creating state dir: %w", err)
	}
	return writeJSON(clientPath(host), cl)
}

// LoadClient returns a stored client registration for the given host, or nil if none exist.
func LoadClient(host string) (*StoredClient, error) {
	var cl StoredClient
	if err := readJSON(clientPath(host), &cl); err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	return &cl, nil
}

// DeleteClient removes a stored client registration for the given host.
func DeleteClient(host string) error {
	err := os.Remove(clientPath(host))
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

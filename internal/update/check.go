// Package update provides background version checking against GitHub releases.
// The check runs asynchronously and caches its result for 24 hours so it never
// adds latency to any command. A notice is printed to stderr when a newer
// version is available.
package update

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gleanwork/glean-cli/internal/httputil"
	"golang.org/x/mod/semver"
)

const (
	cacheFile     = ".glean/update-check.json"
	checkInterval = 24 * time.Hour
	releaseAPIURL = "https://api.github.com/repos/gleanwork/glean-cli/releases/latest"
	devVersion    = "dev"
)

type cacheEntry struct {
	CheckedAt time.Time `json:"checked_at"`
	LatestTag string    `json:"latest_tag"`
}

// CheckAsync fires a background goroutine that checks for a newer release.
// If one is found and the check has not been performed in the last 24 hours,
// it sends a human-readable notice to noticeC. The caller should drain
// noticeC after the command completes and print any message to stderr.
func CheckAsync(currentVersion string) <-chan string {
	ch := make(chan string, 1)
	go func() {
		defer close(ch)
		notice := check(currentVersion)
		if notice != "" {
			ch <- notice
		}
	}()
	return ch
}

func check(currentVersion string) string {
	// Skip for dev builds.
	if currentVersion == devVersion || currentVersion == "" {
		return ""
	}

	cacheFilePath := cacheFilePath()

	// Read the cached result.
	entry, err := readCache(cacheFilePath)
	if err != nil || time.Since(entry.CheckedAt) > checkInterval {
		// Cache is stale or missing — fetch latest from GitHub.
		tag, err := fetchLatestTag()
		if err != nil {
			return "" // Network failure: silent no-op.
		}
		entry = cacheEntry{CheckedAt: time.Now(), LatestTag: tag}
		_ = writeCache(cacheFilePath, entry) // Best-effort.
	}

	if isNewer(entry.LatestTag, currentVersion) {
		return fmt.Sprintf(
			"A new release of glean is available: %s → %s\nRun: brew upgrade glean-cli\n      or: https://github.com/gleanwork/glean-cli/releases/latest",
			currentVersion, entry.LatestTag,
		)
	}
	return ""
}

func fetchLatestTag() (string, error) {
	client := httputil.NewHTTPClient(5 * time.Second)
	resp, err := client.Get(releaseAPIURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := json.Unmarshal(body, &release); err != nil {
		return "", err
	}
	return release.TagName, nil
}

// isNewer returns true if latestTag represents a version strictly greater
// than currentVersion. Both may be in "vMAJOR.MINOR.PATCH" or "MAJOR.MINOR.PATCH" form.
func isNewer(latestTag, currentVersion string) bool {
	return semver.Compare(canonical(latestTag), canonical(currentVersion)) > 0
}

// canonical ensures the version has a "v" prefix as required by golang.org/x/mod/semver.
func canonical(v string) string {
	if !strings.HasPrefix(v, "v") {
		v = "v" + v
	}
	return v
}

func cacheFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, cacheFile)
}

func readCache(path string) (cacheEntry, error) {
	var entry cacheEntry
	if path == "" {
		return entry, fmt.Errorf("no cache path")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return entry, err
	}
	err = json.Unmarshal(data, &entry)
	return entry, err
}

func writeCache(path string, entry cacheEntry) error {
	if path == "" {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o600)
}

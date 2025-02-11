package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"os"
)

// MaybeAnonymizeURL conditionally replaces URLs with randomly generated ones when GLEAN_CLI_ANONYMIZE is set.
// The same input URL will always generate the same random URL within a session.
func MaybeAnonymizeURL(urlStr string) string {
	if os.Getenv("GLEAN_CLI_ANONYMIZE") == "" {
		return urlStr
	}

	// Parse the URL to validate it
	_, err := url.Parse(urlStr)
	if err != nil {
		return urlStr
	}

	hash := sha256.Sum256([]byte(urlStr))
	shortHash := hex.EncodeToString(hash[:8])

	return "https://docs.example.com/" + shortHash
}

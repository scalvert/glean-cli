package utils

import "strings"

// FormatDatasource formats a datasource string for display
func FormatDatasource(s string) string {
	if s == "nonindexedshortcut" {
		return "GoLink"
	}

	words := []rune(s)
	if len(words) > 0 {
		words[0] = []rune(strings.ToUpper(string(words[0])))[0]
	}
	for i := 1; i < len(words); i++ {
		if words[i-1] == ' ' {
			words[i] = []rune(strings.ToUpper(string(words[i])))[0]
		}
	}
	return string(words)
}

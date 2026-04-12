package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ProjectConfig holds project-level CLI preferences that can be shared
// via a .glean/config.json file checked into a repository.
type ProjectConfig struct {
	DefaultOutput string `json:"default_output,omitempty"`
	DefaultMode   string `json:"default_mode,omitempty"`
	DefaultFields string `json:"default_fields,omitempty"`
}

// FindProjectConfig walks up from the current working directory looking for
// a .glean/config.json file. Returns nil, nil if no project config is found.
func FindProjectConfig() (*ProjectConfig, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, nil
	}
	return findProjectConfigFrom(dir)
}

func findProjectConfigFrom(dir string) (*ProjectConfig, error) {
	for {
		candidate := filepath.Join(dir, ".glean", "config.json")
		data, err := os.ReadFile(candidate)
		if err == nil {
			var cfg ProjectConfig
			if err := json.Unmarshal(data, &cfg); err != nil {
				return nil, fmt.Errorf("invalid %s: %w", candidate, err)
			}
			return &cfg, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return nil, nil
}

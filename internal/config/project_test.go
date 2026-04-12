package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindProjectConfigFrom_Found(t *testing.T) {
	root := t.TempDir()
	gleanDir := filepath.Join(root, ".glean")
	if err := os.MkdirAll(gleanDir, 0755); err != nil {
		t.Fatal(err)
	}
	content := `{"default_output":"json","default_mode":"chat","default_fields":"title,url"}`
	if err := os.WriteFile(filepath.Join(gleanDir, "config.json"), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	// Search from a nested subdirectory — should walk up and find it.
	nested := filepath.Join(root, "a", "b", "c")
	if err := os.MkdirAll(nested, 0755); err != nil {
		t.Fatal(err)
	}

	cfg, err := findProjectConfigFrom(nested)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg == nil {
		t.Fatal("expected config, got nil")
	}
	if cfg.DefaultOutput != "json" {
		t.Errorf("DefaultOutput = %q, want %q", cfg.DefaultOutput, "json")
	}
	if cfg.DefaultMode != "chat" {
		t.Errorf("DefaultMode = %q, want %q", cfg.DefaultMode, "chat")
	}
	if cfg.DefaultFields != "title,url" {
		t.Errorf("DefaultFields = %q, want %q", cfg.DefaultFields, "title,url")
	}
}

func TestFindProjectConfigFrom_NotFound(t *testing.T) {
	root := t.TempDir()
	cfg, err := findProjectConfigFrom(root)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg != nil {
		t.Fatalf("expected nil config, got %+v", cfg)
	}
}

func TestFindProjectConfigFrom_InvalidJSON(t *testing.T) {
	root := t.TempDir()
	gleanDir := filepath.Join(root, ".glean")
	if err := os.MkdirAll(gleanDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(gleanDir, "config.json"), []byte(`{bad json`), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := findProjectConfigFrom(root)
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
	if cfg != nil {
		t.Fatalf("expected nil config on error, got %+v", cfg)
	}
}

func TestFindProjectConfigFrom_EmptyConfig(t *testing.T) {
	root := t.TempDir()
	gleanDir := filepath.Join(root, ".glean")
	if err := os.MkdirAll(gleanDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(gleanDir, "config.json"), []byte(`{}`), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := findProjectConfigFrom(root)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg == nil {
		t.Fatal("expected config, got nil")
	}
	if cfg.DefaultOutput != "" || cfg.DefaultMode != "" || cfg.DefaultFields != "" {
		t.Errorf("expected all empty fields, got %+v", cfg)
	}
}

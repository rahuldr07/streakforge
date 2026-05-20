package app

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHelpCommand(t *testing.T) {
	var out bytes.Buffer
	err := RunWithConfig([]string{"help"}, Config{Stdout: &out, Home: t.TempDir()})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	if !strings.Contains(out.String(), "streakforge add") {
		t.Fatalf("help output missing add command: %s", out.String())
	}
}

func TestUnknownCommand(t *testing.T) {
	var out bytes.Buffer
	err := RunWithConfig([]string{"nope"}, Config{Stdout: &out, Home: t.TempDir()})
	if err == nil {
		t.Fatal("expected unknown command error")
	}
}

func TestDefaultConfigUsesUserHomeDirectory(t *testing.T) {
	cfg := DefaultConfig(&bytes.Buffer{}, &bytes.Buffer{})

	if cfg.Home == "" {
		t.Fatal("Home is empty")
	}
	if cfg.Stdout == nil {
		t.Fatal("Stdout is nil")
	}
	if cfg.Stderr == nil {
		t.Fatal("Stderr is nil")
	}
}

func TestStatsPrintsLongestCompletionCount(t *testing.T) {
	home := t.TempDir()
	data := `{
  "habits": [
    {"id": "a", "name": "Read", "completions": ["2026-05-17"]},
    {"id": "b", "name": "Run", "completions": ["2026-05-16", "2026-05-17"]}
  ]
}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"stats"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	if !strings.Contains(out.String(), "longest completions: 2") {
		t.Fatalf("stats output missing longest completions: %s", out.String())
	}
}

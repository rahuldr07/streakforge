package app

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
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

func TestNextPrintsPendingHabitsByStreakPriority(t *testing.T) {
	home := t.TempDir()
	now := time.Now()
	today := now.Format(time.DateOnly)
	yesterday := now.AddDate(0, 0, -1).Format(time.DateOnly)
	twoDaysAgo := now.AddDate(0, 0, -2).Format(time.DateOnly)
	data := `{
  "habits": [
    {"id": "read", "name": "Read", "completions": ["` + twoDaysAgo + `", "` + yesterday + `"]},
    {"id": "walk", "name": "Walk", "completions": ["` + yesterday + `"]},
    {"id": "water", "name": "Water", "completions": ["` + today + `"]},
    {"id": "old", "name": "Old", "archived": true}
  ]
}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"next"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(lines) != 2 {
		t.Fatalf("lines = %d, want 2; output=%q", len(lines), out.String())
	}
	if !strings.Contains(lines[0], "read  2 days  Read") {
		t.Fatalf("first next line = %q", lines[0])
	}
	if !strings.Contains(lines[1], "walk  1 day  Walk") {
		t.Fatalf("second next line = %q", lines[1])
	}
}

package habit

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestSaveReplacesExistingStore(t *testing.T) {
	path := filepath.Join(t.TempDir(), "store.json")
	original := Store{Habits: []Habit{{
		ID:          "old",
		Name:        "Old Habit",
		Description: strings.Repeat("stale-data", 20),
		CreatedAt:   time.Date(2026, 5, 17, 9, 0, 0, 0, time.UTC),
		Completions: []string{"2026-05-17"},
	}}}
	if err := Save(path, original); err != nil {
		t.Fatalf("save original: %v", err)
	}

	replacement := Store{Habits: []Habit{{
		ID:        "new",
		Name:      "New",
		CreatedAt: time.Date(2026, 5, 18, 9, 0, 0, 0, time.UTC),
	}}}
	if err := Save(path, replacement); err != nil {
		t.Fatalf("save replacement: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read store: %v", err)
	}
	if strings.Contains(string(data), "stale-data") {
		t.Fatalf("store still contains stale bytes: %s", string(data))
	}
}

func TestSaveCreatesParentDirectories(t *testing.T) {
	path := filepath.Join(t.TempDir(), "nested", "store.json")
	store := Store{Habits: []Habit{{
		ID:        "h1",
		Name:      "Read",
		CreatedAt: time.Date(2026, 5, 18, 9, 0, 0, 0, time.UTC),
	}}}

	if err := Save(path, store); err != nil {
		t.Fatalf("Save returned error: %v", err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("stat saved file: %v", err)
	}
}

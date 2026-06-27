package habit

import (
	"testing"
	"time"
)

func TestCompleteHabitOncePerDay(t *testing.T) {
	day := time.Date(2026, 5, 18, 9, 0, 0, 0, time.UTC)
	store := Store{Habits: []Habit{{ID: "h1", Name: "Run"}}}

	store, _, err := Complete(store, "h1", day)
	if err != nil {
		t.Fatalf("Complete returned error: %v", err)
	}
	store, _, err = Complete(store, "h1", day)
	if err != nil {
		t.Fatalf("Complete returned error: %v", err)
	}
	if got := len(store.Habits[0].Completions); got != 1 {
		t.Fatalf("completions = %d", got)
	}
}

func TestCompleteArchivedHabitReturnsError(t *testing.T) {
	day := time.Date(2026, 5, 18, 9, 0, 0, 0, time.UTC)
	store := Store{Habits: []Habit{{ID: "h1", Name: "Run", Archived: true}}}

	updated, _, err := Complete(store, "h1", day)

	if err != ErrArchived {
		t.Fatalf("err = %v, want %v", err, ErrArchived)
	}
	if got := len(updated.Habits[0].Completions); got != 0 {
		t.Fatalf("completions = %d", got)
	}
}

func TestUncompleteRemovesCompletionForDay(t *testing.T) {
	day := time.Date(2026, 5, 18, 9, 0, 0, 0, time.UTC)
	store := Store{Habits: []Habit{{
		ID:          "h1",
		Name:        "Run",
		Completions: []string{"2026-05-17", "2026-05-18"},
	}}}

	updated, _, removed, err := Uncomplete(store, "h1", day)
	if err != nil {
		t.Fatalf("Uncomplete returned error: %v", err)
	}
	if !removed {
		t.Fatal("removed = false, want true")
	}
	if got := updated.Habits[0].Completions; len(got) != 1 || got[0] != "2026-05-17" {
		t.Fatalf("completions = %+v", got)
	}
}

func TestUncompleteReportsMissingCompletion(t *testing.T) {
	day := time.Date(2026, 5, 18, 9, 0, 0, 0, time.UTC)
	store := Store{Habits: []Habit{{ID: "h1", Name: "Run"}}}

	_, _, removed, err := Uncomplete(store, "h1", day)
	if err != nil {
		t.Fatalf("Uncomplete returned error: %v", err)
	}
	if removed {
		t.Fatal("removed = true, want false")
	}
}

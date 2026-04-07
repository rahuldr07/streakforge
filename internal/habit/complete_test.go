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


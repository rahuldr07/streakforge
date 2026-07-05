package habit

import (
	"testing"
	"time"
)

func TestLaggingReturnsQuietActiveHabits(t *testing.T) {
	today := time.Date(2026, 7, 5, 12, 0, 0, 0, time.UTC)
	store := Store{Habits: []Habit{
		{ID: "read", Name: "Read", Completions: []string{"2026-07-01"}},
		{ID: "walk", Name: "Walk", Completions: []string{"2026-07-04"}},
		{ID: "water", Name: "Water"},
		{ID: "old", Name: "Old", Archived: true},
	}}

	lagging := Lagging(store, today, 3)

	if len(lagging) != 2 {
		t.Fatalf("len(lagging) = %d, want 2", len(lagging))
	}
	if lagging[0].ID != "read" || lagging[0].DaysQuiet != 4 {
		t.Fatalf("first lagging habit = %+v", lagging[0])
	}
	if lagging[1].ID != "water" || lagging[1].LastCompleted != "" {
		t.Fatalf("second lagging habit = %+v", lagging[1])
	}
}

func TestLastCompletedReturnsLatestCompletion(t *testing.T) {
	last := LastCompleted(Habit{Completions: []string{"2026-07-01", "2026-07-03", "2026-07-02"}})

	if last != "2026-07-03" {
		t.Fatalf("last = %q, want 2026-07-03", last)
	}
}

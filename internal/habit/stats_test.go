package habit

import (
	"testing"
	"time"
)

func TestSummarize(t *testing.T) {
	store := Store{Habits: []Habit{
		{ID: "a", Name: "A", Completions: []string{"2026-05-18"}},
		{ID: "b", Name: "B", Archived: true, Completions: []string{"2026-05-17", "2026-05-18"}},
	}}
	summary := Summarize(store)
	if summary.ActiveHabits != 1 || summary.ArchivedHabits != 1 {
		t.Fatalf("unexpected habit counts: %+v", summary)
	}
	if summary.TotalCompletions != 3 {
		t.Fatalf("total completions = %d", summary.TotalCompletions)
	}
	if summary.CompletionRate != 1.5 {
		t.Fatalf("completion rate = %f", summary.CompletionRate)
	}
}

func TestSummarizeAtIncludesBestCurrentStreak(t *testing.T) {
	today := time.Date(2026, 5, 18, 8, 0, 0, 0, time.UTC)
	store := Store{Habits: []Habit{
		{ID: "a", Name: "A", Completions: []string{"2026-05-17", "2026-05-18"}},
		{ID: "b", Name: "B", Completions: []string{"2026-05-18"}},
		{ID: "c", Name: "C", Archived: true, Completions: []string{"2026-05-16", "2026-05-17", "2026-05-18"}},
	}}

	summary := SummarizeAt(store, today)
	if summary.BestCurrentStreak != 2 {
		t.Fatalf("best current streak = %d", summary.BestCurrentStreak)
	}
}

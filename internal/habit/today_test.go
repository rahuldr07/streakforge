package habit

import (
	"testing"
	"time"
)

func TestTodayShowsCompletionStatus(t *testing.T) {
	day := time.Date(2026, 5, 27, 9, 0, 0, 0, time.UTC)
	store := Store{Habits: []Habit{
		{ID: "a", Name: "Read", Completions: []string{"2026-05-27"}},
		{ID: "b", Name: "Run"},
		{ID: "c", Name: "Old", Archived: true, Completions: []string{"2026-05-27"}},
	}}

	statuses := Today(store, day)
	if len(statuses) != 2 {
		t.Fatalf("statuses = %d", len(statuses))
	}
	if !statuses[0].Completed || statuses[1].Completed {
		t.Fatalf("unexpected statuses: %+v", statuses)
	}
}

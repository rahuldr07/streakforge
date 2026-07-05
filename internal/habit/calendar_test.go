package habit

import (
	"testing"
	"time"
)

func TestCompletionCalendarCountsActiveHabitsByDay(t *testing.T) {
	today := time.Date(2026, 7, 5, 12, 0, 0, 0, time.UTC)
	store := Store{Habits: []Habit{
		{ID: "read", Name: "Read", Completions: []string{"2026-07-04", "2026-07-05"}},
		{ID: "walk", Name: "Walk", Completions: []string{"2026-07-05", "2026-07-05"}},
		{ID: "old", Name: "Old", Archived: true, Completions: []string{"2026-07-05"}},
	}}

	calendar := CompletionCalendar(store, today, 3)

	if len(calendar) != 3 {
		t.Fatalf("len(calendar) = %d, want 3", len(calendar))
	}
	if calendar[0].Date != "2026-07-03" || calendar[0].Complete != 0 {
		t.Fatalf("first day = %+v", calendar[0])
	}
	if calendar[1].Date != "2026-07-04" || calendar[1].Complete != 1 || calendar[1].Total != 2 {
		t.Fatalf("second day = %+v", calendar[1])
	}
	if calendar[2].Date != "2026-07-05" || calendar[2].Complete != 2 || calendar[2].Rate != 1 {
		t.Fatalf("third day = %+v", calendar[2])
	}
}

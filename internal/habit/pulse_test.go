package habit

import (
	"testing"
	"time"
)

func TestWeeklyPulseCountsActiveHabitCompletions(t *testing.T) {
	today := time.Date(2026, 5, 26, 10, 0, 0, 0, time.UTC)
	store := Store{Habits: []Habit{
		{Name: "Read", Completions: []string{"2026-05-26"}},
		{Name: "Run", Completions: []string{"2026-05-26"}, Archived: true},
	}}

	pulse := WeeklyPulse(store, today)
	last := pulse[len(pulse)-1]
	if last.Date != "2026-05-26" || last.Completion != 1 {
		t.Fatalf("last pulse = %+v", last)
	}
}

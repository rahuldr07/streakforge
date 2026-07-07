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

func TestPulseUsesRequestedDayCount(t *testing.T) {
	today := time.Date(2026, 5, 26, 10, 0, 0, 0, time.UTC)

	pulse := Pulse(Store{}, today, 3)

	if len(pulse) != 3 {
		t.Fatalf("len(pulse) = %d, want 3", len(pulse))
	}
	if pulse[0].Date != "2026-05-24" || pulse[2].Date != "2026-05-26" {
		t.Fatalf("pulse range = %+v", pulse)
	}
}

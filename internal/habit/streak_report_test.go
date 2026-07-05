package habit

import (
	"testing"
	"time"
)

func TestStreakReportsSortByCurrentStreak(t *testing.T) {
	today := time.Date(2026, 7, 5, 12, 0, 0, 0, time.UTC)
	store := Store{Habits: []Habit{
		{ID: "walk", Name: "Walk", Completions: []string{"2026-07-05"}},
		{ID: "read", Name: "Read", Completions: []string{"2026-07-04", "2026-07-05"}},
		{ID: "water", Name: "Water", Completions: []string{"2026-07-03"}},
		{ID: "old", Name: "Old", Archived: true, Completions: []string{"2026-07-05"}},
	}}

	reports := StreakReports(store, today)

	if len(reports) != 3 {
		t.Fatalf("len(reports) = %d, want 3", len(reports))
	}
	if reports[0].ID != "read" || reports[0].CurrentStreak != 2 {
		t.Fatalf("first report = %+v", reports[0])
	}
	if reports[1].ID != "walk" || reports[1].LastCompleted != "2026-07-05" {
		t.Fatalf("second report = %+v", reports[1])
	}
	if reports[2].ID != "water" || reports[2].CurrentStreak != 0 {
		t.Fatalf("third report = %+v", reports[2])
	}
}

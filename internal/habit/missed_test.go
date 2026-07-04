package habit

import (
	"testing"
	"time"
)

func TestMissedReturnsHabitsSkippedYesterday(t *testing.T) {
	day := time.Date(2026, 5, 18, 12, 0, 0, 0, time.UTC)
	store := Store{Habits: []Habit{
		{ID: "read", Name: "Read", Completions: []string{"2026-05-16"}},
		{ID: "walk", Name: "Walk", Completions: []string{"2026-05-17"}},
		{ID: "water", Name: "Water"},
		{ID: "old", Name: "Old", Archived: true},
	}}

	missed := Missed(store, day)

	if len(missed) != 2 {
		t.Fatalf("len(missed) = %d, want 2", len(missed))
	}
	if missed[0].ID != "water" {
		t.Fatalf("first missed habit = %q, want water", missed[0].ID)
	}
	if missed[0].LastCompleted != "" {
		t.Fatalf("water last completed = %q, want empty", missed[0].LastCompleted)
	}
	if missed[1].ID != "read" {
		t.Fatalf("second missed habit = %q, want read", missed[1].ID)
	}
	if missed[1].LastCompleted != "2026-05-16" {
		t.Fatalf("read last completed = %q, want 2026-05-16", missed[1].LastCompleted)
	}
}

package habit

import (
	"testing"
	"time"
)

func TestNextSkipsCompletedAndArchivedHabits(t *testing.T) {
	day := time.Date(2026, 5, 27, 9, 0, 0, 0, time.UTC)
	store := Store{Habits: []Habit{
		{ID: "a", Name: "Read", Completions: []string{"2026-05-27"}},
		{ID: "b", Name: "Run"},
		{ID: "c", Name: "Old", Archived: true},
	}}

	actions := Next(store, day)
	if len(actions) != 1 {
		t.Fatalf("actions = %d, want 1", len(actions))
	}
	if actions[0].ID != "b" {
		t.Fatalf("actions[0].ID = %q, want b", actions[0].ID)
	}
}

func TestNextPrioritizesStreaksAtRisk(t *testing.T) {
	day := time.Date(2026, 5, 27, 9, 0, 0, 0, time.UTC)
	store := Store{Habits: []Habit{
		{ID: "a", Name: "Stretch", Completions: []string{"2026-05-26"}},
		{ID: "b", Name: "Read", Completions: []string{"2026-05-24", "2026-05-25", "2026-05-26"}},
		{ID: "c", Name: "Run"},
	}}

	actions := Next(store, day)
	want := []string{"b", "a", "c"}
	for i, id := range want {
		if actions[i].ID != id {
			t.Fatalf("actions[%d].ID = %q, want %q; actions=%+v", i, actions[i].ID, id, actions)
		}
	}
	if actions[0].CurrentStreak != 3 {
		t.Fatalf("top streak = %d, want 3", actions[0].CurrentStreak)
	}
}

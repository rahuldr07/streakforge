package habit

import "testing"

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
}


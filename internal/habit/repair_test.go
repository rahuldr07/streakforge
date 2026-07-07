package habit

import "testing"

func TestRepairSortsAndDeduplicatesCompletions(t *testing.T) {
	store := Store{Habits: []Habit{{
		ID:          "read",
		Name:        "Read",
		Completions: []string{"2026-07-05", "2026-07-03", "2026-07-05"},
	}}}

	repaired, result := Repair(store)

	if !result.Changed {
		t.Fatal("repair did not report changes")
	}
	if result.DuplicateRemoved != 1 {
		t.Fatalf("duplicates removed = %d, want 1", result.DuplicateRemoved)
	}
	if result.CompletionsSorted != 1 {
		t.Fatalf("sorted count = %d, want 1", result.CompletionsSorted)
	}
	completions := repaired.Habits[0].Completions
	if len(completions) != 2 || completions[0] != "2026-07-03" || completions[1] != "2026-07-05" {
		t.Fatalf("completions = %#v", completions)
	}
}

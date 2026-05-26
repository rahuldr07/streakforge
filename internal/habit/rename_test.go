package habit

import "testing"

func TestRenameHabit(t *testing.T) {
	store := Store{Habits: []Habit{{ID: "h1", Name: "Old"}}}

	store, renamed, err := Rename(store, "h1", "  New Name  ", "updated")
	if err != nil {
		t.Fatalf("Rename returned error: %v", err)
	}
	if renamed.Name != "New Name" || store.Habits[0].Description != "updated" {
		t.Fatalf("renamed habit = %+v", renamed)
	}
}

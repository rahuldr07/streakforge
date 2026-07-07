package habit

import "testing"

func TestUnarchiveRestoresArchivedHabit(t *testing.T) {
	store := Store{Habits: []Habit{{ID: "read", Name: "Read", Archived: true}}}

	updated, restored, err := Unarchive(store, "read")
	if err != nil {
		t.Fatalf("Unarchive returned error: %v", err)
	}
	if restored.Archived {
		t.Fatal("restored habit is still archived")
	}
	if updated.Habits[0].Archived {
		t.Fatal("stored habit is still archived")
	}
}

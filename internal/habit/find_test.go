package habit

import "testing"

func TestFindReturnsHabitByID(t *testing.T) {
	store := Store{Habits: []Habit{{ID: "h1", Name: "Read"}}}

	h, err := Find(store, "h1")
	if err != nil {
		t.Fatalf("Find returned error: %v", err)
	}
	if h.Name != "Read" {
		t.Fatalf("name = %q", h.Name)
	}
}

func TestSearchMatchesActiveNameAndDescription(t *testing.T) {
	store := Store{Habits: []Habit{
		{ID: "read", Name: "Read", Description: "Books before bed"},
		{ID: "run", Name: "Run"},
		{ID: "old", Name: "Old books", Archived: true},
	}}

	matches := Search(store, "book")
	if len(matches) != 1 {
		t.Fatalf("matches = %d, want 1", len(matches))
	}
	if matches[0].ID != "read" {
		t.Fatalf("match ID = %q", matches[0].ID)
	}
}

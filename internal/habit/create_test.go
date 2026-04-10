package habit

import (
	"testing"
	"time"
)

func TestAddHabit(t *testing.T) {
	now := time.Date(2026, 5, 18, 12, 0, 0, 0, time.UTC)
	store, created, err := Add(Store{}, "  Read  ", "daily pages", now)
	if err != nil {
		t.Fatalf("Add returned error: %v", err)
	}
	if created.Name != "Read" {
		t.Fatalf("name = %q", created.Name)
	}
	if len(store.Habits) != 1 {
		t.Fatalf("habits = %d", len(store.Habits))
	}
}

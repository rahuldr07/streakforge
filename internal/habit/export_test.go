package habit

import (
	"strings"
	"testing"
	"time"
)

func TestWriteCSV(t *testing.T) {
	var out strings.Builder
	err := WriteCSV(&out, Store{Habits: []Habit{{
		ID:          "h1",
		Name:        "Read",
		Description: "daily pages",
		CreatedAt:   time.Date(2026, 5, 18, 9, 0, 0, 0, time.UTC),
	}}})
	if err != nil {
		t.Fatalf("WriteCSV returned error: %v", err)
	}
	if !strings.Contains(out.String(), "id,name,description,created_at,archived,completions") {
		t.Fatalf("unexpected csv header: %s", out.String())
	}
	if !strings.Contains(out.String(), "h1,Read,daily pages,2026-05-18T09:00:00Z,false,0") {
		t.Fatalf("unexpected csv: %s", out.String())
	}
}

func TestWriteJSON(t *testing.T) {
	var out strings.Builder
	err := WriteJSON(&out, Store{Habits: []Habit{{
		ID:          "h1",
		Name:        "Read",
		Description: "daily pages",
		CreatedAt:   time.Date(2026, 5, 18, 9, 0, 0, 0, time.UTC),
	}}})
	if err != nil {
		t.Fatalf("WriteJSON returned error: %v", err)
	}
	if !strings.Contains(out.String(), `"id": "h1"`) {
		t.Fatalf("json missing id: %s", out.String())
	}
	if !strings.Contains(out.String(), `"name": "Read"`) {
		t.Fatalf("json missing name: %s", out.String())
	}
}

func TestWriteCompletionCSV(t *testing.T) {
	var out strings.Builder
	err := WriteCompletionCSV(&out, Store{Habits: []Habit{{
		ID:          "h1",
		Name:        "Read",
		Completions: []string{"2026-07-04", "2026-07-05"},
	}}})
	if err != nil {
		t.Fatalf("WriteCompletionCSV returned error: %v", err)
	}
	if !strings.Contains(out.String(), "habit_id,habit_name,date,archived") {
		t.Fatalf("unexpected csv header: %s", out.String())
	}
	if !strings.Contains(out.String(), "h1,Read,2026-07-05,false") {
		t.Fatalf("unexpected csv row: %s", out.String())
	}
}

func TestWriteActiveCSVSkipsArchivedHabits(t *testing.T) {
	var out strings.Builder
	err := WriteActiveCSV(&out, Store{Habits: []Habit{
		{ID: "read", Name: "Read"},
		{ID: "old", Name: "Old", Archived: true},
	}})
	if err != nil {
		t.Fatalf("WriteActiveCSV returned error: %v", err)
	}
	if !strings.Contains(out.String(), "read,Read") {
		t.Fatalf("active csv missing active habit: %s", out.String())
	}
	if strings.Contains(out.String(), "old,Old") {
		t.Fatalf("active csv included archived habit: %s", out.String())
	}
}

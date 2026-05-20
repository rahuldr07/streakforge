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

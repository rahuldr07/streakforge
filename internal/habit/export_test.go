package habit

import (
	"strings"
	"testing"
)

func TestWriteCSV(t *testing.T) {
	var out strings.Builder
	err := WriteCSV(&out, Store{Habits: []Habit{{ID: "h1", Name: "Read"}}})
	if err != nil {
		t.Fatalf("WriteCSV returned error: %v", err)
	}
	if !strings.Contains(out.String(), "h1,Read,false,0") {
		t.Fatalf("unexpected csv: %s", out.String())
	}
}


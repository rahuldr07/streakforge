package habit

import (
	"strings"
	"testing"
	"time"
)

func TestDoctorReportsStoreIssues(t *testing.T) {
	today := time.Date(2026, 7, 5, 12, 0, 0, 0, time.UTC)
	store := Store{Habits: []Habit{
		{ID: "read", Name: "Read", Completions: []string{"2026-07-05", "2026-07-05", "bad"}},
		{ID: "read", Name: "Duplicate"},
		{ID: "walk", Completions: []string{"2026-07-06"}},
	}}

	issues := Doctor(store, today)
	messages := issueMessages(issues)

	for _, want := range []string{
		"duplicate habit id: read",
		"habit walk has empty name",
		"habit read has duplicate completion: 2026-07-05",
		"habit read has invalid completion date: bad",
		"habit walk has future completion: 2026-07-06",
	} {
		if !strings.Contains(messages, want) {
			t.Fatalf("doctor output missing %q in %q", want, messages)
		}
	}
}

func TestDoctorWarnsForEmptyStore(t *testing.T) {
	issues := Doctor(Store{}, time.Date(2026, 7, 5, 0, 0, 0, 0, time.UTC))

	if len(issues) != 1 || issues[0].Severity != "warn" || issues[0].Message != "store has no habits" {
		t.Fatalf("issues = %+v", issues)
	}
}

func issueMessages(issues []StoreIssue) string {
	var messages []string
	for _, issue := range issues {
		messages = append(messages, issue.Message)
	}
	return strings.Join(messages, "\n")
}

package habit

import (
	"testing"
	"time"
)

func TestCurrentStreak(t *testing.T) {
	today := time.Date(2026, 5, 18, 8, 0, 0, 0, time.UTC)
	h := Habit{Completions: []string{"2026-05-16", "2026-05-17", "2026-05-18"}}
	if got := CurrentStreak(h, today); got != 3 {
		t.Fatalf("streak = %d", got)
	}
}


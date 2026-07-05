package habit

import "time"

type StoreIssue struct {
	Severity string
	Message  string
}

func Doctor(store Store, today time.Time) []StoreIssue {
	var issues []StoreIssue
	if len(store.Habits) == 0 {
		return []StoreIssue{{Severity: "warn", Message: "store has no habits"}}
	}

	ids := make(map[string]bool, len(store.Habits))
	for _, h := range store.Habits {
		if h.ID == "" {
			issues = append(issues, StoreIssue{Severity: "error", Message: "habit has empty id"})
		} else if ids[h.ID] {
			issues = append(issues, StoreIssue{Severity: "error", Message: "duplicate habit id: " + h.ID})
		}
		ids[h.ID] = true

		if h.Name == "" {
			issues = append(issues, StoreIssue{Severity: "error", Message: "habit " + habitLabel(h) + " has empty name"})
		}
		seenCompletions := make(map[string]bool, len(h.Completions))
		for _, completion := range h.Completions {
			day, err := time.Parse(time.DateOnly, completion)
			if err != nil {
				issues = append(issues, StoreIssue{Severity: "error", Message: "habit " + habitLabel(h) + " has invalid completion date: " + completion})
				continue
			}
			if seenCompletions[completion] {
				issues = append(issues, StoreIssue{Severity: "warn", Message: "habit " + habitLabel(h) + " has duplicate completion: " + completion})
			}
			seenCompletions[completion] = true
			if day.After(startOfDay(today)) {
				issues = append(issues, StoreIssue{Severity: "warn", Message: "habit " + habitLabel(h) + " has future completion: " + completion})
			}
		}
	}
	return issues
}

func habitLabel(h Habit) string {
	if h.ID != "" {
		return h.ID
	}
	if h.Name != "" {
		return h.Name
	}
	return "<unknown>"
}

func startOfDay(day time.Time) time.Time {
	return time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())
}

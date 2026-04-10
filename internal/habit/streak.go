package habit

import "time"

func CurrentStreak(h Habit, today time.Time) int {
	completed := make(map[string]bool, len(h.Completions))
	for _, day := range h.Completions {
		completed[day] = true
	}
	count := 0
	cursor := today
	for {
		key := cursor.Format(time.DateOnly)
		if !completed[key] {
			return count
		}
		count++
		cursor = cursor.AddDate(0, 0, -1)
	}
}

package habit

import "time"

type Summary struct {
	ActiveHabits      int
	ArchivedHabits    int
	TotalCompletions  int
	LongestCompletion int
	BestCurrentStreak int
	CompletionRate    float64
}

func Summarize(store Store) Summary {
	return SummarizeAt(store, time.Now())
}

func SummarizeAt(store Store, today time.Time) Summary {
	var summary Summary
	for _, h := range store.Habits {
		if h.Archived {
			summary.ArchivedHabits++
		} else {
			summary.ActiveHabits++
			streak := CurrentStreak(h, today)
			if streak > summary.BestCurrentStreak {
				summary.BestCurrentStreak = streak
			}
		}
		count := len(h.Completions)
		summary.TotalCompletions += count
		if count > summary.LongestCompletion {
			summary.LongestCompletion = count
		}
	}
	if len(store.Habits) > 0 {
		summary.CompletionRate = float64(summary.TotalCompletions) / float64(len(store.Habits))
	}
	return summary
}

package habit

type Summary struct {
	ActiveHabits      int
	ArchivedHabits    int
	TotalCompletions  int
	LongestCompletion int
}

func Summarize(store Store) Summary {
	var summary Summary
	for _, h := range store.Habits {
		if h.Archived {
			summary.ArchivedHabits++
		} else {
			summary.ActiveHabits++
		}
		count := len(h.Completions)
		summary.TotalCompletions += count
		if count > summary.LongestCompletion {
			summary.LongestCompletion = count
		}
	}
	return summary
}


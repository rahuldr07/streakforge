package habit

import (
	"sort"
	"time"
)

type MissedHabit struct {
	ID            string
	Name          string
	LastCompleted string
	CurrentStreak int
}

func Missed(store Store, day time.Time) []MissedHabit {
	yesterday := day.AddDate(0, 0, -1)
	key := yesterday.Format(time.DateOnly)
	missed := make([]MissedHabit, 0, len(store.Habits))

	for _, h := range store.Habits {
		if h.Archived || contains(h.Completions, key) {
			continue
		}
		missed = append(missed, MissedHabit{
			ID:            h.ID,
			Name:          h.Name,
			LastCompleted: lastCompletedBefore(h.Completions, yesterday),
			CurrentStreak: CurrentStreak(h, day),
		})
	}

	sort.SliceStable(missed, func(i, j int) bool {
		if missed[i].LastCompleted != missed[j].LastCompleted {
			return missed[i].LastCompleted < missed[j].LastCompleted
		}
		if missed[i].Name != missed[j].Name {
			return missed[i].Name < missed[j].Name
		}
		return missed[i].ID < missed[j].ID
	})

	return missed
}

func lastCompletedBefore(completions []string, day time.Time) string {
	cutoff := day.Format(time.DateOnly)
	last := ""
	for _, completion := range completions {
		if completion < cutoff && completion > last {
			last = completion
		}
	}
	return last
}

package habit

import (
	"sort"
	"time"
)

type NextAction struct {
	ID            string
	Name          string
	CurrentStreak int
}

func Next(store Store, day time.Time) []NextAction {
	today := day.Format(time.DateOnly)
	yesterday := day.AddDate(0, 0, -1)
	actions := make([]NextAction, 0, len(store.Habits))

	for _, h := range store.Habits {
		if h.Archived || contains(h.Completions, today) {
			continue
		}
		actions = append(actions, NextAction{
			ID:            h.ID,
			Name:          h.Name,
			CurrentStreak: CurrentStreak(h, yesterday),
		})
	}

	sort.SliceStable(actions, func(i, j int) bool {
		if actions[i].CurrentStreak != actions[j].CurrentStreak {
			return actions[i].CurrentStreak > actions[j].CurrentStreak
		}
		if actions[i].Name != actions[j].Name {
			return actions[i].Name < actions[j].Name
		}
		return actions[i].ID < actions[j].ID
	})

	return actions
}

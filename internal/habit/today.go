package habit

import "time"

type TodayStatus struct {
	ID        string
	Name      string
	Completed bool
}

func Today(store Store, day time.Time) []TodayStatus {
	key := day.Format(time.DateOnly)
	statuses := make([]TodayStatus, 0, len(store.Habits))
	for _, h := range store.Habits {
		if h.Archived {
			continue
		}
		statuses = append(statuses, TodayStatus{
			ID:        h.ID,
			Name:      h.Name,
			Completed: contains(h.Completions, key),
		})
	}
	return statuses
}

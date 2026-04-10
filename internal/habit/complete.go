package habit

import "time"

func Complete(store Store, id string, day time.Time) (Store, Habit, error) {
	key := day.Format(time.DateOnly)
	for i, h := range store.Habits {
		if h.ID != id {
			continue
		}
		if !contains(h.Completions, key) {
			h.Completions = append(h.Completions, key)
		}
		store.Habits[i] = h
		return store, h, nil
	}
	return store, Habit{}, ErrNotFound
}

func contains(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

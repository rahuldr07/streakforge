package habit

import "time"

func Complete(store Store, id string, day time.Time) (Store, Habit, error) {
	key := day.Format(time.DateOnly)
	for i, h := range store.Habits {
		if h.ID != id {
			continue
		}
		if h.Archived {
			return store, Habit{}, ErrArchived
		}
		if !contains(h.Completions, key) {
			h.Completions = append(h.Completions, key)
		}
		store.Habits[i] = h
		return store, h, nil
	}
	return store, Habit{}, ErrNotFound
}

func Uncomplete(store Store, id string, day time.Time) (Store, Habit, bool, error) {
	key := day.Format(time.DateOnly)
	for i, h := range store.Habits {
		if h.ID != id {
			continue
		}
		if h.Archived {
			return store, Habit{}, false, ErrArchived
		}
		completions := h.Completions[:0]
		removed := false
		for _, completion := range h.Completions {
			if completion == key {
				removed = true
				continue
			}
			completions = append(completions, completion)
		}
		h.Completions = completions
		store.Habits[i] = h
		return store, h, removed, nil
	}
	return store, Habit{}, false, ErrNotFound
}

func contains(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

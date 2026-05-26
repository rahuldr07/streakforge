package habit

import "strings"

func Rename(store Store, id string, name string, description string) (Store, Habit, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return store, Habit{}, ErrEmptyName
	}
	for i, h := range store.Habits {
		if h.ID != id {
			continue
		}
		h.Name = name
		h.Description = strings.TrimSpace(description)
		store.Habits[i] = h
		return store, h, nil
	}
	return store, Habit{}, ErrNotFound
}

package habit

import (
	"strings"
	"time"
)

func Add(store Store, name string, description string, now time.Time) (Store, Habit, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return store, Habit{}, ErrEmptyName
	}
	id, err := NewID()
	if err != nil {
		return store, Habit{}, err
	}
	h := Habit{
		ID:          id,
		Name:        name,
		Description: strings.TrimSpace(description),
		CreatedAt:   now.UTC(),
		Completions: []string{},
	}
	store.Habits = append(store.Habits, h)
	return store, h, nil
}

package habit

import "strings"

func Find(store Store, id string) (Habit, error) {
	for _, h := range store.Habits {
		if h.ID == id {
			return h, nil
		}
	}
	return Habit{}, ErrNotFound
}

func Search(store Store, query string) []Habit {
	query = strings.ToLower(strings.TrimSpace(query))
	if query == "" {
		return nil
	}
	matches := make([]Habit, 0)
	for _, h := range store.Habits {
		if h.Archived {
			continue
		}
		name := strings.ToLower(h.Name)
		description := strings.ToLower(h.Description)
		if strings.Contains(name, query) || strings.Contains(description, query) {
			matches = append(matches, h)
		}
	}
	return matches
}

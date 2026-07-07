package habit

func Archive(store Store, id string) (Store, Habit, error) {
	for i, h := range store.Habits {
		if h.ID != id {
			continue
		}
		h.Archived = true
		store.Habits[i] = h
		return store, h, nil
	}
	return store, Habit{}, ErrNotFound
}

func Unarchive(store Store, id string) (Store, Habit, error) {
	for i, h := range store.Habits {
		if h.ID != id {
			continue
		}
		h.Archived = false
		store.Habits[i] = h
		return store, h, nil
	}
	return store, Habit{}, ErrNotFound
}

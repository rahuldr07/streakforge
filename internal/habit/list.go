package habit

func Active(store Store) []Habit {
	habits := make([]Habit, 0, len(store.Habits))
	for _, h := range store.Habits {
		if !h.Archived {
			habits = append(habits, h)
		}
	}
	return habits
}

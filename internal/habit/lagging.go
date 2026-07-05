package habit

import (
	"sort"
	"time"
)

type LaggingHabit struct {
	ID            string
	Name          string
	LastCompleted string
	DaysQuiet     int
}

func Lagging(store Store, today time.Time, minDays int) []LaggingHabit {
	if minDays < 0 {
		minDays = 0
	}
	today = startOfDay(today)
	habits := make([]LaggingHabit, 0, len(store.Habits))
	for _, h := range Active(store) {
		last := LastCompleted(h)
		daysQuiet := minDays
		if last == "" {
			daysQuiet = minDays + 1
		} else {
			day, err := time.ParseInLocation(time.DateOnly, last, today.Location())
			if err != nil {
				continue
			}
			daysQuiet = int(today.Sub(day).Hours() / 24)
		}
		if daysQuiet < minDays {
			continue
		}
		habits = append(habits, LaggingHabit{
			ID:            h.ID,
			Name:          h.Name,
			LastCompleted: last,
			DaysQuiet:     daysQuiet,
		})
	}

	sort.SliceStable(habits, func(i, j int) bool {
		if habits[i].DaysQuiet != habits[j].DaysQuiet {
			return habits[i].DaysQuiet > habits[j].DaysQuiet
		}
		if habits[i].Name != habits[j].Name {
			return habits[i].Name < habits[j].Name
		}
		return habits[i].ID < habits[j].ID
	})
	return habits
}

func LastCompleted(h Habit) string {
	last := ""
	for _, completion := range h.Completions {
		if completion > last {
			last = completion
		}
	}
	return last
}

package habit

import "time"

type DayPulse struct {
	Date       string
	Completion int
}

func WeeklyPulse(store Store, today time.Time) []DayPulse {
	start := today.AddDate(0, 0, -6)
	pulse := make([]DayPulse, 0, 7)
	for i := 0; i < 7; i++ {
		day := start.AddDate(0, 0, i).Format(time.DateOnly)
		count := 0
		for _, h := range store.Habits {
			if h.Archived {
				continue
			}
			if contains(h.Completions, day) {
				count++
			}
		}
		pulse = append(pulse, DayPulse{Date: day, Completion: count})
	}
	return pulse
}

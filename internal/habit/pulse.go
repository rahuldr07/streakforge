package habit

import "time"

type DayPulse struct {
	Date       string
	Completion int
}

func WeeklyPulse(store Store, today time.Time) []DayPulse {
	return Pulse(store, today, 7)
}

func Pulse(store Store, today time.Time, days int) []DayPulse {
	if days <= 0 {
		return nil
	}
	start := today.AddDate(0, 0, -days+1)
	pulse := make([]DayPulse, 0, days)
	for i := 0; i < days; i++ {
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

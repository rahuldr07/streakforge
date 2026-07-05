package habit

import "time"

type CalendarDay struct {
	Date     string
	Complete int
	Total    int
	Rate     float64
}

func CompletionCalendar(store Store, today time.Time, days int) []CalendarDay {
	if days <= 0 {
		return nil
	}
	active := Active(store)
	completedByDay := make(map[string]int)
	for _, h := range active {
		seen := make(map[string]bool, len(h.Completions))
		for _, completion := range h.Completions {
			if seen[completion] {
				continue
			}
			seen[completion] = true
			completedByDay[completion]++
		}
	}

	calendar := make([]CalendarDay, 0, days)
	start := today.AddDate(0, 0, -days+1)
	for offset := 0; offset < days; offset++ {
		date := start.AddDate(0, 0, offset).Format(time.DateOnly)
		day := CalendarDay{
			Date:     date,
			Complete: completedByDay[date],
			Total:    len(active),
		}
		if day.Total > 0 {
			day.Rate = float64(day.Complete) / float64(day.Total)
		}
		calendar = append(calendar, day)
	}
	return calendar
}

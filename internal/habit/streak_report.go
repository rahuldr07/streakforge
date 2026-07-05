package habit

import (
	"sort"
	"time"
)

type StreakReport struct {
	ID            string
	Name          string
	CurrentStreak int
	LastCompleted string
}

func StreakReports(store Store, today time.Time) []StreakReport {
	reports := make([]StreakReport, 0, len(store.Habits))
	for _, h := range Active(store) {
		reports = append(reports, StreakReport{
			ID:            h.ID,
			Name:          h.Name,
			CurrentStreak: CurrentStreak(h, today),
			LastCompleted: LastCompleted(h),
		})
	}

	sort.SliceStable(reports, func(i, j int) bool {
		if reports[i].CurrentStreak != reports[j].CurrentStreak {
			return reports[i].CurrentStreak > reports[j].CurrentStreak
		}
		if reports[i].LastCompleted != reports[j].LastCompleted {
			return reports[i].LastCompleted > reports[j].LastCompleted
		}
		if reports[i].Name != reports[j].Name {
			return reports[i].Name < reports[j].Name
		}
		return reports[i].ID < reports[j].ID
	})
	return reports
}

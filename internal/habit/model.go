package habit

import "time"

type Habit struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	Archived    bool      `json:"archived"`
	Completions []string  `json:"completions"`
}

type Store struct {
	Habits []Habit `json:"habits"`
}


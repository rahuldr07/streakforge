package habit

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"strconv"
	"time"
)

func WriteCSV(w io.Writer, store Store) error {
	return writeHabitCSV(w, store.Habits)
}

func WriteActiveCSV(w io.Writer, store Store) error {
	return writeHabitCSV(w, Active(store))
}

func writeHabitCSV(w io.Writer, habits []Habit) error {
	writer := csv.NewWriter(w)
	if err := writer.Write([]string{"id", "name", "description", "created_at", "archived", "completions"}); err != nil {
		return err
	}
	for _, h := range habits {
		if err := writer.Write([]string{
			h.ID,
			h.Name,
			h.Description,
			h.CreatedAt.Format(time.RFC3339),
			strconv.FormatBool(h.Archived),
			strconv.Itoa(len(h.Completions)),
		}); err != nil {
			return err
		}
	}
	writer.Flush()
	return writer.Error()
}

func WriteCompletionCSV(w io.Writer, store Store) error {
	writer := csv.NewWriter(w)
	if err := writer.Write([]string{"habit_id", "habit_name", "date", "archived"}); err != nil {
		return err
	}
	for _, h := range store.Habits {
		for _, completion := range h.Completions {
			if err := writer.Write([]string{
				h.ID,
				h.Name,
				completion,
				strconv.FormatBool(h.Archived),
			}); err != nil {
				return err
			}
		}
	}
	writer.Flush()
	return writer.Error()
}

func WriteJSON(w io.Writer, store Store) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(store)
}

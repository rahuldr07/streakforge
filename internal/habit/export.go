package habit

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"strconv"
	"time"
)

func WriteCSV(w io.Writer, store Store) error {
	writer := csv.NewWriter(w)
	if err := writer.Write([]string{"id", "name", "description", "created_at", "archived", "completions"}); err != nil {
		return err
	}
	for _, h := range store.Habits {
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

func WriteJSON(w io.Writer, store Store) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(store)
}

package habit

import (
	"encoding/csv"
	"io"
	"strconv"
)

func WriteCSV(w io.Writer, store Store) error {
	writer := csv.NewWriter(w)
	if err := writer.Write([]string{"id", "name", "archived", "completions"}); err != nil {
		return err
	}
	for _, h := range store.Habits {
		if err := writer.Write([]string{
			h.ID,
			h.Name,
			strconv.FormatBool(h.Archived),
			strconv.Itoa(len(h.Completions)),
		}); err != nil {
			return err
		}
	}
	writer.Flush()
	return writer.Error()
}

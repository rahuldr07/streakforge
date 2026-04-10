package habit

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

func DefaultPath(home string) string {
	if home == "" {
		return ".streakforge.json"
	}
	return filepath.Join(home, ".streakforge.json")
}

func Load(path string) (Store, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return Store{}, nil
	}
	if err != nil {
		return Store{}, err
	}
	var store Store
	if err := json.Unmarshal(data, &store); err != nil {
		return Store{}, err
	}
	return store, nil
}

func Save(path string, store Store) error {
	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(data, '\n'), 0o600)
}

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
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(dir, ".streakforge-*.tmp")
	if err != nil {
		return err
	}
	tmpPath := tmp.Name()
	defer os.Remove(tmpPath)

	if _, err := tmp.Write(append(data, '\n')); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmpPath, path)
}

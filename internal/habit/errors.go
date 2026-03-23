package habit

import "errors"

var (
	ErrEmptyName = errors.New("habit name is required")
	ErrNotFound  = errors.New("habit not found")
)


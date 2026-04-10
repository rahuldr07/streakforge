package habit

import "errors"

var (
	ErrEmptyName = errors.New("habit name is required")
	ErrArchived  = errors.New("habit is archived")
	ErrNotFound  = errors.New("habit not found")
)

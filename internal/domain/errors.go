package domain

import "errors"

var (
	ErrNotFound   = errors.New("todo not found")
	ErrDuplicated = errors.New("todo already exists")
	ErrValidation = errors.New("invalid todo")
)

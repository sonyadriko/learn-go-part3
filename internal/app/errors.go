package app

import "errors"

var (
	// ErrNotFound indicates that the requested resource was not found.
	ErrNotFound = errors.New("not found")
	// Add other custom error definitions as needed.
)

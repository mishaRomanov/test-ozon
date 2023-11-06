package storage

import "errors"

var (
	ErrNotFound      = errors.New("link not found")
	ErrNotWritten    = errors.New("could not write the data")
	ErrEmptyInput    = errors.New("data wasn't written, input empty")
	ErrAlreadyExists = errors.New("short link already exists")
)

// Storager interface is for different storage solutions
// whether it's postgres or in-memory
type Storager interface {
	GetValue(string) (string, error)
	LookUp(string) (bool, error)
	WriteValue(string, string) error
}

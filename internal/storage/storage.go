package storage

import "errors"

var (
	ErrNotFound      = errors.New("link not found")
	ErrNotWritten    = errors.New("could not write the data")
	ErrEmptyInput    = errors.New("data wasn't written, input empty")
	ErrAlreadyExists = errors.New("short link already exists")
)

// создаем интерфейс как для постгреса
// так и для  in-memory решения
type Storager interface {
	GetValue() error
	WriteValue() error
}

package storage

import (
	"github.com/mishaRomanov/test-ozon/internal/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteValue(t *testing.T) {
	object := NewCache()
	t.Run("Successfull writing", func(t *testing.T) {
		assert.Equal(t, nil, object.WriteValue("short", "full"))
	})
	t.Run("Link already exists", func(t *testing.T) {
		assert.Equal(t, storage.ErrAlreadyExists, object.WriteValue("short", "full"))
	})

	//tests for empty input
	t.Run("Empty input: both", func(t *testing.T) {
		err := object.WriteValue("", "")
		assert.Equal(t, storage.ErrEmptyInput, err)
	})
	t.Run("Empty input: full link", func(t *testing.T) {
		err := object.WriteValue("", "world")
		assert.Equal(t, storage.ErrEmptyInput, err)
	})
	t.Run("Empty input: short link", func(t *testing.T) {
		err := object.WriteValue("hello", "")
		assert.Equal(t, storage.ErrEmptyInput, err)
	})
}

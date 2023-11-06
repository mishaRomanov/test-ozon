package storage

import (
	"github.com/mishaRomanov/test-ozon/internal/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteValue(t *testing.T) {
	testmap := NewCache()
	testmap.WriteValue("a", "b")
	assert.Equal(t, testmap.WriteValue("", ""), storage.ErrEmptyInput)
	assert.Equal(t, testmap.WriteValue("a", "b"), storage.ErrAlreadyExists)
}

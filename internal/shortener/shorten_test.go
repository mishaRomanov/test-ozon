package shorten

import (
	"github.com/mishaRomanov/test-ozon/internal/storage"
	cache "github.com/mishaRomanov/test-ozon/internal/storage/cache"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeAShortLink(t *testing.T) {
	//test for already shortened link
	t.Run("Link already shortened", func(t *testing.T) {
		object := cache.NewCache()
		object.WriteValue("short", "full")
		assert.Equal(t, storage.ErrAlreadyExists, object.WriteValue("short", "full"))
	})
	//test for empty input
	t.Run("Empty input", func(t *testing.T) {
		object := cache.NewCache()
		err := object.WriteValue("", "")
		assert.Equal(t, storage.ErrEmptyInput, err)
	})
}

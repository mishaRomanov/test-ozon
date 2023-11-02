package storage

import (
	"github.com/mishaRomanov/test-ozon/internal/storage"
	"sync"
)

type Cache struct {
	cache map[string]string
	mu    sync.Mutex
}

// GetValue возвращает значение
func (c *Cache) GetValue(val string) (string, error) {
	res, ok := c.cache[val]
	if !ok {
		return "", storage.ErrNotFound
	}
	return res, nil
}

// WriteValue записывает значения и проверяет
// не пустые ли значения передаваемые в метод
func (c *Cache) WriteValue(short, full string) error {
	c.mu.Lock()
	if short == "" || full == "" {
		return storage.ErrEmptyInput
	}
	defer c.mu.Unlock()
	c.cache[short] = full
	return nil
}

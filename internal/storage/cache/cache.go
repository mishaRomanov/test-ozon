package storage

import (
	"github.com/mishaRomanov/test-ozon/internal/storage"
	"sync"
)

type Cache struct {
	Cache map[string]string
	Mu    sync.Mutex
}

// WriteValue записывает значения и проверяет
// не пустые ли значения передаваемые в метод
func (c *Cache) WriteValue(short, full string) error {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	if short == "" || full == "" {
		return storage.ErrEmptyInput
	}
	c.Cache[short] = full
	return nil
}

// GetValue возвращает значение
func (c *Cache) GetValue(val string) (string, error) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	res, ok := c.Cache[val]
	if !ok {
		return "", storage.ErrNotFound
	}
	return res, nil
}
func NewCache() *Cache {
	var c Cache
	c.Cache = make(map[string]string)
	return &c
}

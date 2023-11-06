package storage

import (
	"github.com/mishaRomanov/test-ozon/internal/storage"
	"sync"
)

type Cache struct {
	Cache map[string]string
	Mu    sync.Mutex
}

// WriteValue.....writes value!
// and checks if there is one
func (c *Cache) WriteValue(short, full string) error {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	if short == "" || full == "" {
		return storage.ErrEmptyInput
	}
	//check if we even have a value
	if _, ok := c.Cache[short]; ok {
		return storage.ErrAlreadyExists
	}
	c.Cache[short] = full
	return nil
}

// Returns a value if one is found
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

func (c *Cache) LookUp(value string) (bool, error) {
	for _, link := range c.Cache {
		if link == value {
			return true, storage.ErrAlreadyExists
		}
	}
	return false, nil
}

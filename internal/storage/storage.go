package storage

import (
	"sync"
)

type DataCache struct {
	cache map[string]string
	mu sync.RWMutex
}
func (c *DataCache) New() {
	c.cache = make(map[string]string)
}

func (c *DataCache) Set(key string,  value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	c.cache[key] = value
}

func (c *DataCache) Get(key string) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	name := c.cache[key]
	return name
}

func (c *DataCache) Delete(key string) {
	c.mu.Lock()
	_, ok := c.cache[key]

	if ok {
		delete(c.cache, key)
		c.mu.Unlock()
	} else {
		c.mu.Unlock()
	}
}

func (c *DataCache) Update(key string,  value string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.cache[key]
	if ok {
		c.cache[key] = value
		return true
	} else {
		return false
	}
}

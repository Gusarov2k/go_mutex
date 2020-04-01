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
	c.cache[key] = value
	c.mu.Unlock()
}

func (c *DataCache) Get(key string) string {
	c.mu.Lock()
	name := c.cache[key]
	defer c.mu.Unlock()
	if name == "" {
		return ""
	}
	return name
}

func (c *DataCache) Delete(key string) string {
	c.mu.Lock()
	_, ok := c.cache[key]

	if ok {
		delete(c.cache, key)
		c.mu.Unlock()
		return "key deleted"
	} else {
		c.mu.Unlock()
		return "key not find"
	}
}

func (c *DataCache) Update(key string,  value string) string {
	c.mu.Lock()
	_, ok := c.cache[key]
	if ok {
		c.cache[key] = value
		c.mu.Unlock()
		return "key update"
	} else {
		c.mu.Unlock()
		return "key not find"
	}
}

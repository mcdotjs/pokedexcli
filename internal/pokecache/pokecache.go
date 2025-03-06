package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	value map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		value: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		c.mu.Lock()
		<-ticker.C
		for key, entry := range c.value {
			if time.Now().Sub(entry.createdAt) > interval {
				delete(c.value, key)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) Add(k string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
	c.value[k] = entry

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, exist := c.value[key]
	if exist {
		return v.val, true
	}
	return []byte{}, false
}

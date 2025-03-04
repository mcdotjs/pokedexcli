package pokecache

import (
	"fmt"
	"sync"
	"time"

	"github.com/charmbracelet/bubbles/key"
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
	//start feap loop
	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) reapLoop(interval time.Duration) {

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
	v, exist := c.value[key]
	if exist {
		return v.val, false
	}
	return []byte{}, false
}

package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entry map[string]cacheEntry
	mu    *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		Entry: make(map[string]cacheEntry),
		mu:    &sync.RWMutex{},
	}

	go newCache.reapLoop(interval)
	return newCache
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	var addEntry cacheEntry
	addEntry.createdAt = time.Now()
	addEntry.val = val
	c.Entry[key] = addEntry
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, exists := c.Entry[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c Cache) reapLoop(interval time.Duration) {
	for {
		for key := range c.Entry {
			if time.Since(c.Entry[key].createdAt) > interval {
				c.mu.Lock()
				delete(c.Entry, key)
				c.mu.Unlock()
			}
		}
		if len(c.Entry) <= 0 {
			return
		}
	}
}

func (c Cache) CheckExpiry() {

}

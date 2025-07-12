package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	Entries map[string]cacheEntry
	mu      *sync.Mutex
}

func NewCache(duration time.Duration) *Cache {
	cache := Cache{
		Entries: make(map[string]cacheEntry),
		mu:      &sync.Mutex{},
	}
	go cache.reapLoop(duration)
	return &cache
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
	c.Entries[key] = entry

	//fmt.Printf("%s was cached.\n", key)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.Entries[key]
	if exists {
		return entry.val, exists
	} else {
		return nil, exists
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.Entries {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.Entries, k)
		}
	}
}

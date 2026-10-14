package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	CacheMap map[string]cacheEntry
	mu       sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(ttl time.Duration) Cache {
	c := Cache{
		CacheMap: make(map[string]cacheEntry),
	}

	go c.readLoop(ttl)
	return c
}

func (c *Cache) Add(key string, v []byte) error {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.CacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       v,
	}
	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.CacheMap[key]
	if ok {
		v.createdAt = time.Now()
		c.CacheMap[key] = v
	}

	return v.val, ok
}

func (c *Cache) readLoop(ttl time.Duration) {

	ticker := time.NewTicker(ttl)
	for {

		<-ticker.C

		for key, val := range c.CacheMap {
			expectedttl := val.createdAt.Add(ttl)
			if expectedttl.Before(time.Now()) {
				c.mu.Lock()
				delete(c.CacheMap, key)
				c.mu.Unlock()
			}

		}

	}

}

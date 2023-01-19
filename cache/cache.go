package cache

import (
	"sync"
	"time"
)

type cacheItem struct {
	value      interface{}
	ttl        time.Duration
	expireTime time.Time
}

type Cache struct {
	data   map[string]*cacheItem
	mu     sync.RWMutex
	delete chan string
}

func New() *Cache {
	c := &Cache{
		data:   make(map[string]*cacheItem),
		delete: make(chan string),
	}
	go c.cleanup()
	return c
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	expireTime := time.Now().Add(ttl)
	c.data[key] = &cacheItem{value: value, ttl: ttl, expireTime: expireTime}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.data[key]
	if ok {
		if time.Now().After(item.expireTime) {
			c.delete <- key
			return nil, false
		}
		return item.value, true
	}
	return nil, false
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func (c *Cache) cleanup() {
	for {
		select {
		case key := <-c.delete:
			c.mu.Lock()
			delete(c.data, key)
			c.mu.Unlock()
		case <-time.After(time.Second * 30):
			c.mu.Lock()
			for key, item := range c.data {
				if time.Now().After(item.expireTime) {
					delete(c.data, key)
				}
			}
			c.mu.Unlock()
		}
	}
}

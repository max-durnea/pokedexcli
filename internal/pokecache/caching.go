package pokecache
import (
	"time"
	"sync"
)
type Cache struct {
	Entries map[string]cacheEntry
	mu sync.Mutex
	ttl time.Duration
}

type cacheEntry struct{
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		Entries: map[string]cacheEntry{},
		ttl: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte){
	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Entries[key]= entry
}

func (c *Cache) Get(key string) ([]byte, bool){
	c.mu.Lock()
	defer c.mu.Unlock()
	val,ok := c.Entries[key]
	if !ok{
		return []byte{}, false
	}
	return val.val, true
}

func (c *Cache) reapLoop(){
	ticker := time.NewTicker(c.ttl)
	for range ticker.C {
		c.mu.Lock()
		for k,v := range c.Entries{
			if time.Since(v.createdAt) > c.ttl{
				delete(c.Entries,k)
			}
		}
		c.mu.Unlock()
	}
}
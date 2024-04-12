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
	cacheEntryMap map[string]cacheEntry
	lock          *sync.Mutex
}

func NewCache(interval time.Duration) Cache {

	var ourCache Cache
	ourCache.cacheEntryMap = make(map[string]cacheEntry)
	ourCache.lock = new(sync.Mutex);

	go ourCache.readLoop(interval);
	return ourCache; 

}

func (c *Cache) Add(key string, param_val []byte) {

	c.lock.Lock()
 	defer c.lock.Unlock()
	var temp_entry cacheEntry
	temp_entry.val = param_val
	temp_entry.createdAt = time.Now().UTC()

	c.cacheEntryMap[key] = temp_entry

	

}

func (c *Cache) Get(key string) ([]byte, bool) {

	c.lock.Lock()
	defer c.lock.Unlock()
	value, ok := c.cacheEntryMap[key]

	return value.val, ok;
}

func (c *Cache) readLoop(valid_duration time.Duration) {
	interval := valid_duration
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.lock.Lock()

		for x, y := range c.cacheEntryMap{

			
			current_time := time.Now().UTC()
			if current_time.Sub(y.createdAt) > valid_duration {
				delete(c.cacheEntryMap, x);
			}

		}


		c.lock.Unlock()
	
	}
}



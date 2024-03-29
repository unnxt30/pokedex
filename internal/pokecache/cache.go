package main

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cahceEntry_map map[string]cacheEntry
	lock           *sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {

}

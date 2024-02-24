package mapcache

import (
	"fmt"
	"time"
)

var _ Cache = InMemoryCache{}

type CacheEntry struct {
	settledAt time.Time
	value     interface{}
}

type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}

type InMemoryCache struct {
	expireIn time.Duration
	Cache    map[string]CacheEntry

	setKey chan *chSet
}

func (c InMemoryCache) Get(key string) interface{} {
	// if time
	return c.Cache[key].value
}

type chSet struct {
	key       string
	settledAt time.Time
}

func (c InMemoryCache) Set(key string, value interface{}) {

	c.setKey <- &chSet{key: key, settledAt: time.Now()}
	c.Cache[key] = CacheEntry{settledAt: time.Now(), value: value}
}

func NewInMemoryCache(expireIn time.Duration) *InMemoryCache {

	cache := &InMemoryCache{
		expireIn: expireIn,
		Cache:    make(map[string]CacheEntry),
		setKey:   make(chan *chSet, 10),
	}

	go cache.StartCacheShedulerLoop()

	return cache
}

func (c *InMemoryCache) StartCacheShedulerLoop() {

	queue := []chSet{}

	for {

		select {
		case newkey := <-c.setKey:
			fmt.Println("select new key", newkey)
			queue = append(queue, *newkey)
		// почему ловим дедлок без дефолта?(
		default:
			if len(queue) == 0 {
				continue
			}
			if time.Since(queue[0].settledAt) > c.expireIn {
				key := queue[0].key
				delete(c.Cache, key)

				queue = queue[1:]
				fmt.Println("просрочка",key)
			}
		}

	}
}

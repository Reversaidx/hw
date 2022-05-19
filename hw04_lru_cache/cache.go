package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	if _, ok := c.items[key]; ok {
		List.MoveToFront(c.queue, c.items[key])
		c.items[key].Value = value
		return true
	}

	if c.queue.Len()+1 > c.capacity {
		delete(c.items, c.itemsKey[c.queue.Back()])
		List.Remove(c.queue, c.queue.Back())
	}
	c.items[key] = List.PushFront(c.queue, value)
	c.itemsKey[c.queue.Front()] = key
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	// Lock so only one goroutine at a time can access the map c.v.
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.items[key]
	if !ok {
		return nil, false
	}

	List.PushFront(c.queue, value)
	c.itemsKey[c.queue.Front()] = key
	return c.items[key].Value, true
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	itemsKey map[*ListItem]Key
	mu       sync.Mutex
}

/*type cacheItem struct {
	key   Key
	value interface{}
}*/

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		itemsKey: make(map[*ListItem]Key),
	}
}
func test() {
	l := NewCache(5)
	l.Clear()
}

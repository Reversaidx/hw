package hw04lrucache

import "fmt"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	curValue, ok := c.items[key]
	if ok && curValue.Value != nil {
		c.items[key].Value = value
		List.MoveToFront(c.queue, c.items[key])
		return true
	} else {
		if c.queue.Len() == c.capacity {
			List.Remove(c.queue, c.queue.Back())
		}

		c.items[key] = List.PushFront(c.queue, value)
		return false
	}
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	value, ok := c.items[key]
	fmt.Println(ok)
	if !ok || value == nil {
		return nil, false
	}
	List.PushFront(c.queue, value)
	return c.items[key].Value, true
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   string
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

func (c *lruCache) Set(key Key, value interface{}) bool {

	_, ok := c.items[key]
	if ok {
		List.MoveToFront(c.queue, c.items[key])
		c.items[key].Value = value
		c.itemsKey[c.items[key]] = key
		return true
	} else {
		if c.queue.Len()+1 > c.capacity {
			delete(c.items, c.itemsKey[c.queue.Back()])
			List.Remove(c.queue, c.queue.Back())

		}
		c.items[key] = List.PushFront(c.queue, value)
		c.itemsKey[c.queue.Front()] = key
		return false
	}
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	value, ok := c.items[key]
	if !ok {
		return nil, false
	}
	delete(c.itemsKey, value)
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

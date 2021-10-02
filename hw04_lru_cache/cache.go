package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

// type cacheItem struct {
// 	key   string
// 	value interface{}
// }

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if i, ok := c.items[key]; ok {
		i.Value = value
		c.queue.MoveToFront(i)
		return true
	}

	i := c.queue.PushFront(value)
	if c.queue.Len() > c.capacity {
		c.queue.Remove(c.queue.Back())
	}
	c.items[key] = i

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	val, ok := c.items[key]
	if !ok {
		return nil, false
	}

	c.queue.MoveToFront(val)
	return val.Value, true
}

func (c *lruCache) Clear() {
	c.items = nil
	c.queue = nil
}

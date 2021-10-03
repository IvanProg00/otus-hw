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

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if i, ok := c.items[key]; ok {
		i.Value = cacheItem{key: key, value: value}
		c.queue.MoveToFront(i)
		return true
	}

	i := c.queue.PushFront(cacheItem{key, value})
	if c.queue.Len() > c.capacity {
		if cache, ok := c.queue.Back().Value.(cacheItem); ok {
			c.items[cache.key] = nil
			delete(c.items, cache.key)
		}
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
	cache, ok := val.Value.(cacheItem)
	if !ok {
		return nil, false
	}
	return cache.value, true
}

func (c *lruCache) Clear() {
	for k := range c.items {
		delete(c.items, k)
	}
	for i := c.queue.Front(); i != nil; i = i.Next {
		i.Prev = nil
		i.Value = nil
		i.Next = nil
	}
}

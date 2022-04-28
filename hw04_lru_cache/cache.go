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

func (lru *lruCache) Set(key Key, value interface{}) bool {
	el, ok := lru.items[key]

	if ok {
		// Элемент уже есть в очереди - обновим и переместим в начало
		el.Value = value
		lru.queue.MoveToFront(el)
	} else {
		// Эленента еще нет в очереди
		el = lru.queue.PushFront(value)

		// Если размер очереди > объема кеша - удалим малоиспользуемый элемент (с конца)
		if lru.queue.Len() > lru.capacity {
			elToRemove := lru.queue.Back()

			// из мапы
			for mapKey, mapValue := range lru.items {
				if mapValue.Value == elToRemove.Value {
					delete(lru.items, mapKey)
					break
				}
			}
			// из очереди
			lru.queue.Remove(elToRemove)
		}
	}
	lru.items[key] = el

	return ok
}

func (lru lruCache) Get(key Key) (interface{}, bool) {
	el, ok := lru.items[key]
	var res interface{}

	if ok {
		lru.queue.MoveToFront(el)
		res = el.Value
	}
	return res, ok
}

func (lru *lruCache) Clear() {
	lru.queue = nil
	lru.items = nil
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

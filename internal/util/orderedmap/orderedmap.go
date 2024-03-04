package orderedmap

import (
	"container/list"
	"sync"
)

type item struct {
	key      string
	value    interface{}
	queueElt *list.Element
}

type OrderedMap struct {
	values map[string]*item
	queue  *list.List
	mutex  sync.RWMutex
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{
		values: make(map[string]*item),
		mutex:  sync.RWMutex{},
		queue:  list.New(),
	}
}

func (om *OrderedMap) Set(key string, value interface{}) {
	newItem := &item{
		key:   key,
		value: value,
	}
	om.mutex.Lock()
	defer om.mutex.Unlock()

	newItem.queueElt = om.queue.PushBack(newItem)
	om.values[key] = newItem
}

func (om *OrderedMap) Get(key string) (interface{}, bool) {
	om.mutex.RLock()
	defer om.mutex.RUnlock()

	if n, ok := om.values[key]; ok {
		return n.value, true
	}
	return nil, false
}

func (om *OrderedMap) DeleteItem(key string) {
	om.mutex.Lock()
	defer om.mutex.Unlock()

	if n, ok := om.values[key]; ok {
		om.queue.Remove(n.queueElt)
		delete(om.values, key)
	}
}

func (om *OrderedMap) GetAll() ([]string, []interface{}) {
	om.mutex.RLock()
	defer om.mutex.RUnlock()

	keys := make([]string, 0, len(om.values))
	values := make([]interface{}, 0, len(om.values))
	queueLen := om.queue.Len()
	i := 0

	for queueItem := om.queue.Front(); i < queueLen; queueItem = queueItem.Next() {
		item := queueItem.Value.(*item)

		keys = append(keys, item.key)
		values = append(values, item.value)
		i++
	}
	return keys, values
}

package list

import (
	"sync"
)

type Item struct {
	data       interface{}
	prev, next *Item
	list       *List // The list to which this element belongs.
}

// <--> A (first) <--> B <--> C <--> D (last) <-->
type List struct {
	first *Item
	last  *Item
	len   int
	lock  sync.RWMutex
}

func New() *List {
	return &List{}
}

func (list *List) Length() int {
	return list.len
}

func (list *List) IsEmpty() bool {
	return list.len == 0
}

func (list *List) First() *Item {
	return list.first
}

func (list *List) Last() *Item {
	return list.last
}

func (item *Item) Prev() *Item {
	if p := item.prev; item.list != nil && p != item.list.last {
		return p
	}
	return nil
}

func (item *Item) Next() *Item {
	if p := item.next; item.list != nil && p != item.list.first {
		return p
	}
	return nil
}

func (list *List) Has(data interface{}) bool {
	list.lock.RLock()
	defer list.lock.RUnlock()

	if (list.IsEmpty()) {
		return false
	}

	for cur, last := list.First(), list.Last(); cur != nil; cur = cur.next {
		if cur.data == data {
			return true
		} else if (cur == last) {
			return false
		}
	}
	return false
}

func (list *List) Remove(data interface{}) bool {
	list.lock.Lock()
	defer list.lock.Unlock()

	if (list.IsEmpty()) {
		return false
	}

	for cur, last := list.First(), list.Last(); cur != nil; cur = cur.next {
		if cur.data == data {
			cur.prev.next = cur.next
			if cur.next != nil {
				cur.next.prev = cur.prev
			}
			list.len--
			cur.data = nil
			cur.prev = nil
			cur.next = nil
			cur.list = nil
			return true
		} else if (cur == last) {
			return false
		}
	}
	return false
}

// Insert to the last
func (list *List) Insert(data interface{}) {
	list.lock.Lock()
	defer list.lock.Unlock()

	newItem := &Item{data: data, list: list}
	if list.IsEmpty() {
		newItem.prev = newItem
		newItem.next = newItem
		list.first = newItem
		list.last = newItem
	} else {
		newItem.prev = list.last
		newItem.next = list.first
		list.first.prev = newItem
		list.last.next = newItem
		list.last = newItem
	}
	list.len++
}
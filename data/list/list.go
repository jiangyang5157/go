package list

import (
	"sync"
)

type Element struct {
	value       interface{}
	prev, next *Element
	list       *List // The list to which this element belongs.
}

// <--> A (first) <--> B <--> C <--> D (last) <-->
type List struct {
	first *Element
	last  *Element
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

func (list *List) First() *Element {
	return list.first
}

func (list *List) Last() *Element {
	return list.last
}

func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != e.list.last {
		return p
	}
	return nil
}

func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != e.list.first {
		return p
	}
	return nil
}

func (list *List) Has(value interface{}) bool {
	list.lock.RLock()
	defer list.lock.RUnlock()

	if (list.IsEmpty()) {
		return false
	}

	for cur, last := list.First(), list.Last(); cur != nil; cur = cur.next {
		if cur.value == value {
			return true
		} else if (cur == last) {
			return false
		}
	}
	return false
}

func (list *List) Remove(value interface{}) bool {
	list.lock.Lock()
	defer list.lock.Unlock()

	if (list.IsEmpty()) {
		return false
	}

	for cur, last := list.First(), list.Last(); cur != nil; cur = cur.next {
		if cur.value == value {
			cur.prev.next = cur.next
			if cur.next != nil {
				cur.next.prev = cur.prev
			}
			list.len--
			cur.value = nil
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
func (list *List) Insert(value interface{}) {
	list.lock.Lock()
	defer list.lock.Unlock()

	newElement := &Element{value: value, list: list}
	if list.IsEmpty() {
		newElement.prev = newElement
		newElement.next = newElement
		list.first = newElement
		list.last = newElement
	} else {
		newElement.prev = list.last
		newElement.next = list.first
		list.first.prev = newElement
		list.last.next = newElement
		list.last = newElement
	}
	list.len++
}
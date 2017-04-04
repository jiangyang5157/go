package list

import (
	"errors"
)

//  nil <-- A (first) <--> B <--> C <--> D (last) --> nil
type List struct {
	first, last *Element
	size        int
}

type Element struct {
	list       *List
	prev, next *Element
	value      interface{}
}

func NewList() *List {
	return &List{size: 0}
}

func NewElement(list *List, value interface{}) *Element {
	return &Element{list: list, value: value}
}

func (list *List) Size() int {
	return list.size
}

func (list *List) IsEmpty() bool {
	return list.Size() == 0
}

func (list *List) First() *Element {
	return list.first
}

func (list *List) Last() *Element {
	return list.last
}

func (e *Element) Prev() *Element {
	return e.prev
}

func (e *Element) Next() *Element {
	return e.next
}

func (list *List) Concat(other *List) {
	list.last.next, other.first.prev = other.first, list.last
	list.last = other.last
	list.size += other.size
}

func (list *List) Each(f func(element *Element)) {
	for e := list.first; e != nil; e = e.next {
		f(e)
	}
}

func (list *List) Find(value interface{}) (*Element, error) {
	for e := list.first; e != nil; e = e.next {
		if e.value == value {
			return e, nil
		}
	}
	return nil, errors.New("Element not found")
}

func (list *List) Get(index int) (*Element, error) {
	if index > list.size-1 {
		return nil, errors.New("Index out of range")
	}
	node := list.first
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node, nil
}

func (list *List) Prepend(value interface{}) {
	e := NewElement(list, value)
	if list.size == 0 {
		list.first = e
		list.last = list.first
	} else {
		list.first.prev = e
		e.next = list.first
		list.first = e
	}
	list.size++
}

func (list *List) Append(value interface{}) {
	e := NewElement(list, value)
	if list.size == 0 {
		list.first = e
		list.last = list.first
	} else {
		list.last.next = e
		e.prev = list.last
		list.last = e
	}
	list.size++
}

func (list *List) Add(value interface{}, index int) error {
	if index > list.size {
		return errors.New("Index out of range")
	}

	if index == 0 {
		list.Prepend(value)
		return nil
	}

	if index == list.size {
		list.Append(value)
		return nil
	}

	next, _ := list.Get(index)
	prev := next.prev

	e := NewElement(list, value)
	prev.next = e
	e.prev = prev
	next.prev = e
	e.next = next
	list.size++
	return nil
}

func (list *List) Remove(value interface{}) error {
	if list.size == 1 && list.first.value == value {
		list.first = nil
		list.last = nil
		list.size = 0
		return nil
	}

	for e := list.first; e != nil; e = e.next {
		if e.value == value {
			if e.next != nil {
				e.next.prev = e.prev
			}
			if e.prev != nil {
				e.prev.next = e.next
			}
			e.value = nil
			list.size--
			return nil
		}
	}
	return errors.New("Element not found")
}

func (list *List) Clear() {
	list.Each(func(e *Element) {
		if e.next != nil {
			e.next.prev = e.prev
		}
		if e.prev != nil {
			e.prev.next = e.next
		}
		e.value = nil
		list.size--
	})
	list.first = nil
	list.last = nil
}

package heap

import (
	"errors"

	"github.com/jiangyang5157/go/data/queue"
)

type PriorityElement struct {
	value    interface{}
	priority int
}

func NewPriorityElement(value interface{}, priority int) *PriorityElement {
	return &PriorityElement{value: value, priority: priority}
}

func NewPriorityQueue(compare Comparable) *Heap {
	return New(compare)
}

func (h *Heap) ChangePriority(value interface{}, priority int) error {
	length := h.Length()
	if length == 0 {
		return errors.New("Empty priority queue")
	}

	tmp := queue.New()
	var popped PriorityElement
	for h.Length() > 0 {
		popped = h.Extract().(PriorityElement)
		if popped.value == value {
			popped.priority = priority
			h.Insert(popped)
			break
		} else {
			tmp.Push(popped)
		}
	}
	var err error
	if tmp.Length() == length {
		err = errors.New("Element not found")
	}
	for tmp.Length() > 0 {
		// recover pq
		h.Insert(tmp.Pop().(PriorityElement))
	}
	return err
}

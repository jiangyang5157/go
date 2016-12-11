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

func (pq *Heap) ChangePriority(value interface{}, priority int) error {
	length := pq.Length()
	if length == 0 {
		return errors.New("Empty priority queue")
	}

	tmp := queue.New()
	var popped PriorityElement
	for pq.Length() > 0 {
		popped = pq.Extract().(PriorityElement)
		if popped.value == value {
			popped.priority = priority
			pq.Insert(popped)
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
		pq.Insert(tmp.Pop().(PriorityElement))
	}
	return err
}

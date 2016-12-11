package heap

import (
	"errors"

	"github.com/jiangyang5157/go/data/queue"
)

type PriorityQueue struct {
	data Heap
}

type PriorityElement struct {
	value    interface{}
	priority int
}

func NewPriorityElement(value interface{}, priority int) *PriorityElement {
	return &PriorityElement{value: value, priority: priority}
}

func NewPriorityQueue(compare Comparable) *PriorityQueue {
	return &PriorityQueue{data: *New(compare)}
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.Length() == 0
}

func (pq *PriorityQueue) Length() int {
	return pq.data.Length()
}

func (pq *PriorityQueue) Insert(e PriorityElement) {
	pq.data.Insert(e)
}

func (pq *PriorityQueue) Extract() PriorityElement {
	return pq.data.Extract().(PriorityElement)
}

func (pq *PriorityQueue) Peek() PriorityElement {
	return pq.data.Peek().(PriorityElement)
}

func (pq *PriorityQueue) ChangePriority(value interface{}, priority int) error {
	length := pq.Length()
	if length == 0 {
		return errors.New("Empty priority queue")
	}

	tmp := queue.New()
	var popped PriorityElement
	for pq.Length() > 0 {
		popped = pq.Extract()
		if popped.value == value {
			popped.priority = priority
			pq.data.Insert(popped)
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
		pq.data.Insert(tmp.Pop().(PriorityElement))
	}
	return err
}

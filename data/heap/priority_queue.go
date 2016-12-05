package heap

import (
	"errors"

	"github.com/jiangyang5157/go/data/queue"
)

type PriorityQueue struct {
	data Heap
}

type Element struct {
	value    interface{}
	priority int
}

func NewElement(value interface{}, priority int) *Element {
	return &Element{value: value, priority: priority}
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

func (pq *PriorityQueue) Insert(e Element) {
	pq.data.Insert(e)
}

func (pq *PriorityQueue) Extract() Element {
	return pq.data.Extract().(Element)
}

func (pq *PriorityQueue) Peek() Element {
	return pq.data.Peek().(Element)
}

func (pq *PriorityQueue) ChangePriority(value interface{}, priority int) error {
	length := pq.Length()
	if length == 0 {
		return errors.New("Empty priority queue")
	}

	tmp := queue.New()
	var popped Element
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
		pq.data.Insert(tmp.Pop().(Element))
	}
	return err
}

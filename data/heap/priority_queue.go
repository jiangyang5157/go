package heap

import (
	"github.com/jiangyang5157/go/data/queue"
	"errors"
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
	if (pq.Length() == 0) {
		return errors.New("Empty priority queue")
	}

	var err error
	tmp := queue.New()
	popped := pq.Extract()
	for value != popped.value {
		if pq.Length() == 0 {
			err = errors.New("Element not found")
			break
		}
		tmp.Push(popped)
		popped = pq.Extract()
	}
	popped.priority = priority
	pq.data.Insert(popped)

	// re-insert tmp into pq
	for tmp.Length() > 0 {
		pq.data.Insert(tmp.Pop().(Element))
	}
	return err
}
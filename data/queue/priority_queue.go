package queue

import "github.com/jiangyang5157/go/data/heap"

type PriorityQueue struct {
	data heap.Heap
}

type Element struct {
	value    interface{}
	priority int
}

func NewElement(value interface{}, priority int) *Element {
	return &Element{value: value, priority: priority}
}

func NewPriorityQueue(compare heap.Comparable) *PriorityQueue {
	return &PriorityQueue{data: *heap.New(compare)}
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
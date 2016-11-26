package heap

import (
	"testing"
	"fmt"
)

func maxHeapCompare(a interface{}, b interface{}) int {
	switch {
	case a.(int) < b.(int):
		return -1
	case a.(int) > b.(int):
		return 1
	default:
		return 0
	}
}

func minHeapCompare(a interface{}, b interface{}) int {
	switch {
	case a.(int) < b.(int):
		return 1
	case a.(int) > b.(int):
		return -1
	default:
		return 0
	}
}

func Test_Heap(t *testing.T) {
	heap := New(maxHeapCompare)
	if heap.IsEmpty() != true {
		t.Error("IsEmpty doesn't work as expected")
	}
	if heap.Length() != 0 {
		t.Error("Length doesn't work as expected")
	}

	heap.Insert(8)
	heap.Insert(7)
	heap.Insert(6)
	heap.Insert(3)
	heap.Insert(1)
	heap.Insert(0)
	heap.Insert(2)
	heap.Insert(4)
	heap.Insert(5)
	heap.Insert(5)

	if heap.Length() != 10 {
		t.Error("Length doesn't work as expected")
	}

	sorted := make([]int, 0)
	for heap.Length() > 0 {
		sorted = append(sorted, heap.Extract().(int))
	}

	for i := 0; i < len(sorted); i++ {
		fmt.Printf("%v, ", sorted[i])
	}
	fmt.Println("")
}
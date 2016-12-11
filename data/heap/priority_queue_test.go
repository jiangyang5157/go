package heap

import (
	"fmt"
	"testing"
)

func maxPqCompare(a interface{}, b interface{}) int {
	switch {
	case a.(PriorityElement).priority < b.(PriorityElement).priority:
		return -1
	case a.(PriorityElement).priority > b.(PriorityElement).priority:
		return 1
	default:
		return 0
	}
}

func minPqCompare(a interface{}, b interface{}) int {
	switch {
	case a.(PriorityElement).priority < b.(PriorityElement).priority:
		return 1
	case a.(PriorityElement).priority > b.(PriorityElement).priority:
		return -1
	default:
		return 0
	}
}

func Test_MaxPq(t *testing.T) {
	pq := NewPriorityQueue(maxPqCompare)

	if pq.IsEmpty() != true {
		t.Error("IsEmpty doesn't work as expected")
	}
	if pq.Length() != 0 {
		t.Error("Length doesn't work as expected")
	}

	pq.Insert(*NewPriorityElement(8, 10))
	pq.Insert(*NewPriorityElement(7, 11))
	pq.Insert(*NewPriorityElement(6, 12))
	pq.Insert(*NewPriorityElement(3, 13))
	pq.Insert(*NewPriorityElement(1, 14))
	pq.Insert(*NewPriorityElement(0, 19))
	pq.Insert(*NewPriorityElement(2, 18))
	pq.Insert(*NewPriorityElement(4, 17))
	pq.Insert(*NewPriorityElement(9, 16))
	pq.Insert(*NewPriorityElement(5, 15))

	if pq.Length() != 10 {
		t.Error("Length doesn't work as expected")
	}

	sorted := make([]PriorityElement, 0)
	for pq.Length() > 0 {
		sorted = append(sorted, pq.Extract().(PriorityElement))
	}
	for i := 0; i < len(sorted); i++ {
		fmt.Printf("%v, ", sorted[i])
	}
	fmt.Println("")
}

func Test_MinPq(t *testing.T) {
	pq := NewPriorityQueue(minPqCompare)

	if pq.IsEmpty() != true {
		t.Error("IsEmpty doesn't work as expected")
	}
	if pq.Length() != 0 {
		t.Error("Length doesn't work as expected")
	}

	pq.Insert(*NewPriorityElement(8, 10))
	pq.Insert(*NewPriorityElement(7, 11))
	pq.Insert(*NewPriorityElement(6, 12))
	pq.Insert(*NewPriorityElement(3, 13))
	pq.Insert(*NewPriorityElement(1, 14))
	pq.Insert(*NewPriorityElement(0, 19))
	pq.Insert(*NewPriorityElement(2, 18))
	pq.Insert(*NewPriorityElement(4, 17))
	pq.Insert(*NewPriorityElement(9, 16))
	pq.Insert(*NewPriorityElement(5, 15))

	if pq.Length() != 10 {
		t.Error("Length doesn't work as expected")
	}

	sorted := make([]PriorityElement, 0)
	for pq.Length() > 0 {
		sorted = append(sorted, pq.Extract().(PriorityElement))
	}
	for i := 0; i < len(sorted); i++ {
		fmt.Printf("%v, ", sorted[i])
	}
	fmt.Println("")
}

func TestChangePriority(t *testing.T) {
	pq := NewPriorityQueue(maxPqCompare)
	pq.Insert(*NewPriorityElement(8, 10))
	pq.Insert(*NewPriorityElement(7, 11))
	pq.Insert(*NewPriorityElement(6, 12))
	pq.Insert(*NewPriorityElement(3, 13))
	pq.Insert(*NewPriorityElement(1, 14))
	pq.Insert(*NewPriorityElement(0, 19))
	pq.Insert(*NewPriorityElement(2, 18))
	pq.Insert(*NewPriorityElement(4, 17))
	pq.Insert(*NewPriorityElement(9, 16))
	pq.Insert(*NewPriorityElement(5, 15))

	if pq.ChangePriority(22, 222) == nil {
		t.Error("ChangePriority doesn't work as expected")
	}
	if pq.ChangePriority(2, 222) == nil {
		popped := pq.Extract().(PriorityElement)
		if popped.value != 2 {
			t.Error("ChangePriority doesn't work as expected")
		}
	} else {
		t.Error("ChangePriority doesn't work as expected")
	}
}

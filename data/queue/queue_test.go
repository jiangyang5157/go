package queue

import (
	"testing"
)

func Test_Queue(t *testing.T) {
	queue := New()

	if queue.IsEmpty() != true {
		t.Error("IsEmpty doesn't work as expected")
	}

	queue.Insert(1)
	queue.Insert(2)
	queue.Insert(3)
	queue.Insert(4)

	if queue.Length() != 4 {
		t.Error("length doesn't work as expected")
	}

	if queue.Remove() != 1 {
		t.Error("remove doesn't work as expected")
	}

	if queue.Length() != 3 {
		t.Error("Length doesn't work as expected")
	}

	if queue.Peek() != 2 {
		t.Error("Peek doesn't work as expected")
	}
}
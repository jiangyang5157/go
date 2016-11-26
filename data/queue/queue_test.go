package queue

import (
	"testing"
)

func Test_Queue(t *testing.T) {
	queue := New()

	if queue.IsEmpty() != true {
		t.Error("IsEmpty doesn't work as expected")
	}

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)

	if queue.Size() != 4 {
		t.Error("Size doesn't work as expected")
	}

	if queue.Pop() != 1 {
		t.Error("remove doesn't work as expected")
	}

	if queue.Size() != 3 {
		t.Error("Size doesn't work as expected")
	}

	if queue.Peek() != 2 {
		t.Error("Peek doesn't work as expected")
	}
}
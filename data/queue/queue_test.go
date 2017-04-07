package queue

import "testing"

func Test_Queue(t *testing.T) {
	queue := NewQueue()

	if queue.IsEmpty() != true {
		t.Error("IsEmpty doesn't work as expected")
	}

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)

	if queue.Length() != 4 {
		t.Error("Length doesn't work as expected")
	}

	if queue.Pop() != 1 {
		t.Error("Remove doesn't work as expected")
	}

	if queue.Length() != 3 {
		t.Error("Length doesn't work as expected")
	}

	if queue.Peek() != 2 {
		t.Error("Peek doesn't work as expected")
	}

	if queue.Length() != 3 {
		t.Error("Length doesn't work as expected")
	}

	if queue.Pop() != 2 {
		t.Error("Remove doesn't work as expected")
	}

	if queue.Pop() != 3 {
		t.Error("Remove doesn't work as expected")
	}

	if queue.Pop() != 4 {
		t.Error("Remove doesn't work as expected")
	}

	if queue.Pop() != nil {
		t.Error("Pop on an empty queue doesn't work as expected")
	}
}

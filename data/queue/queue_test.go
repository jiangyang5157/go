package queue

import (
	"testing"
)

func Test_Queue(t *testing.T) {
	queue := New()

	if !queue.IsEmpty() {
		t.Error("IsEmpty is wrong")
	}

	queue.Add(1)
	queue.Add(2)
	queue.Add(3)
	queue.Add(4)

	if queue.Length() != 4 {
		t.Error("length is wrong")
	}

	if queue.Remove() != 1 {
		t.Error("remove is wrong")
	}

	if queue.Length() != 3 {
		t.Error("length is wrong after pop")
	}

	if queue.Peek() != 2 {
		t.Error("Peek is wrong")
	}
}
package stack

import (
	"testing"
)

func Test_Stack(t *testing.T) {
	stack := New()

	if stack.IsEmpty() != true {
		t.Error("IsEmpty doesn't work as expected")
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	if stack.Size() != 4 {
		t.Error("Size doesn't work as expected")
	}

	if stack.Pop() != 4 {
		t.Error("Pop doesn't work as expected")
	}

	if stack.Size() != 3 {
		t.Error("Size doesn't work as expected")
	}

	if stack.Peek() != 3 {
		t.Error("Peek doesn't work as expected")
	}
}
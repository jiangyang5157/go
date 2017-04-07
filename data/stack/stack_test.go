package stack

import (
	"testing"
)

func Test_Stack(t *testing.T) {
	stack := NewStack()

	if stack.IsEmpty() != true {
		t.Error("IsEmpty doesn't work as expected")
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	if stack.Length() != 4 {
		t.Error("Length doesn't work as expected")
	}

	if stack.Pop() != 4 {
		t.Error("Pop doesn't work as expected")
	}

	if stack.Length() != 3 {
		t.Error("Length doesn't work as expected")
	}

	if stack.Peek() != 3 {
		t.Error("Peek doesn't work as expected")
	}
}

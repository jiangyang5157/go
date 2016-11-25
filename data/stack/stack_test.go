package stack

import (
	"testing"
)

func Test_Stack(t *testing.T) {
	stack := New()

	if !stack.IsEmpty() {
		t.Error("IsEmpty is wrong")
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	if stack.Length() != 4 {
		t.Error("Length is wrong")
	}

	if stack.Pop() != 4 {
		t.Error("Pop is wrong")
	}

	if stack.Length() != 3 {
		t.Error("Length is wrong after pop")
	}

	if stack.Peek() != 3 {
		t.Error("Peek is wrong")
	}
}
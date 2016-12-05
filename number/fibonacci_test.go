package number

import (
	"testing"
)

func Test_TopDownFibonacci(t *testing.T) {
	if TopDownFibonacci(0) != 0 {
		t.Error("TopDownFibonacci(0) is wrong")
	}

	if TopDownFibonacci(1) != 1 {
		t.Error("TopDownFibonacci(1) is wrong")
	}

	if TopDownFibonacci(2) != 1 {
		t.Error("TopDownFibonacci(2) is wrong")
	}

	if TopDownFibonacci(3) != 2 {
		t.Error("TopDownFibonacci(3) is wrong")
	}

	if TopDownFibonacci(4) != 3 {
		t.Error("TopDownFibonacci(4) is wrong")
	}

	if TopDownFibonacci(5) != 5 {
		t.Error("TopDownFibonacci(5) is wrong")
	}

	if TopDownFibonacci(6) != 8 {
		t.Error("TopDownFibonacci(6) is wrong")
	}

	if TopDownFibonacci(7) != 13 {
		t.Error("TopDownFibonacci(7) is wrong")
	}
}

func Test_BottomUpFibonacci(t *testing.T) {
	if BottomUpFibonacci(0) != 0 {
		t.Error("BottomUpFibonacci(0) is wrong")
	}

	if BottomUpFibonacci(1) != 1 {
		t.Error("BottomUpFibonacci(1) is wrong")
	}

	if BottomUpFibonacci(2) != 1 {
		t.Error("BottomUpFibonacci(2) is wrong")
	}

	if BottomUpFibonacci(3) != 2 {
		t.Error("BottomUpFibonacci(3) is wrong")
	}

	if BottomUpFibonacci(4) != 3 {
		t.Error("BottomUpFibonacci(4) is wrong")
	}

	if BottomUpFibonacci(5) != 5 {
		t.Error("BottomUpFibonacci(5) is wrong")
	}

	if BottomUpFibonacci(6) != 8 {
		t.Error("BottomUpFibonacci(6) is wrong")
	}

	if BottomUpFibonacci(7) != 13 {
		t.Error("BottomUpFibonacci(7) is wrong")
	}
}

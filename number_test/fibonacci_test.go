package number

import (
	"testing"
	"github.com/jiangyang5157/go/number"
)

func Test_TopDoenFibonacci(t *testing.T) {
	if number.TopDoenFibonacci(0) != 0 {
		t.Error("TopDoenFibonacci(0) is wrong")
	}

	if number.TopDoenFibonacci(1) != 1 {
		t.Error("TopDoenFibonacci(1) is wrong")
	}

	if number.TopDoenFibonacci(2) != 1 {
		t.Error("TopDoenFibonacci(2) is wrong")
	}

	if number.TopDoenFibonacci(3) != 2 {
		t.Error("TopDoenFibonacci(3) is wrong")
	}

	if number.TopDoenFibonacci(4) != 3 {
		t.Error("TopDoenFibonacci(4) is wrong")
	}

	if number.TopDoenFibonacci(5) != 5 {
		t.Error("TopDoenFibonacci(5) is wrong")
	}

	if number.TopDoenFibonacci(6) != 8 {
		t.Error("TopDoenFibonacci(6) is wrong")
	}

	if number.TopDoenFibonacci(7) != 13 {
		t.Error("TopDoenFibonacci(7) is wrong")
	}
}

func Test_BottomUpFibonacci(t *testing.T) {
	if number.BottomUpFibonacci(0) != 0 {
		t.Error("BottomUpFibonacci(0) is wrong")
	}

	if number.BottomUpFibonacci(1) != 1 {
		t.Error("BottomUpFibonacci(1) is wrong")
	}

	if number.BottomUpFibonacci(2) != 1 {
		t.Error("BottomUpFibonacci(2) is wrong")
	}

	if number.BottomUpFibonacci(3) != 2 {
		t.Error("BottomUpFibonacci(3) is wrong")
	}

	if number.BottomUpFibonacci(4) != 3 {
		t.Error("BottomUpFibonacci(4) is wrong")
	}

	if number.BottomUpFibonacci(5) != 5 {
		t.Error("BottomUpFibonacci(5) is wrong")
	}

	if number.BottomUpFibonacci(6) != 8 {
		t.Error("BottomUpFibonacci(6) is wrong")
	}

	if number.BottomUpFibonacci(7) != 13 {
		t.Error("BottomUpFibonacci(7) is wrong")
	}
}
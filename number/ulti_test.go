package number

import (
	"testing"
)

func Test_RandomArray(t *testing.T) {
	arr := RandomArray(10)
	if arr == nil {
		t.Error("RandomArray(10) returns nil")
	}
	if len(arr) != 10 {
		t.Error("Lenth doesn't work as expected")
	}

	arr = RandomArray(0)
	if arr == nil {
		t.Error("RandomArray(0) returns nil")
	}
	if len(arr) != 0 {
		t.Error("Lenth doesn't work as expected")
	}

	arr = RandomArray(1)
	if arr[0] != 0 {
		t.Error("RandomArray(1)[0] is wrong")
	}
}

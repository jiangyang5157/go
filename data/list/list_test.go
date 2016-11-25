package list

import (
	"testing"
)

func Test_List(t *testing.T) {
	list := New()

	if list.IsEmpty() != true {
		t.Error("IsEmpty doesn't work as expected")
	}
	if list.Has(222) != false {
		t.Error("Has doesn't work as expected")
	}
	if list.Remove(222) != false {
		t.Error("Remove doesn't work as expected")
	}
	list.Insert(1)
	if list.IsEmpty() != false {
		t.Error("IsEmpty doesn't work as expected")
	}
	if list.Has(1) != true {
		t.Error("Has doesn't work as expected")
	}
	if list.Remove(1) != true {
		t.Error("Remove doesn't work as expected")
	}
	list.Insert(1)
	list.Insert(4)
	list.Insert(3)
	list.Insert(10)
	list.Insert(103)
	list.Insert(56)

	if list.Length() != 6 {
		t.Error("Length doesn't work as expected")
	}

	if list.Remove(4) != true {
		t.Error("Remove doesn't work as expected")
	}
	if list.Length() != 5 {
		t.Error("Length doesn't work as expected")
	}
}
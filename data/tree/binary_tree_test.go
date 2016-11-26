package tree

import (
	"testing"
)

func compare(a interface{}, b interface{}) int {
	switch {
	case a.(int) < b.(int):
		return -1
	case a.(int) > b.(int):
		return 1
	default:
		return 0
	}
}

func Test_BinaryTree(t *testing.T) {
	root := New(compare)

	root.Insert(1)
	root.Insert(2)
	root.Insert(3)
	root.Insert(4)

	findTree := root.Search(3)
	if findTree.value != 3 {
		t.Error("Search doesn't work as expected")
	}

	findNilTree := root.Search(222)
	if findNilTree != nil {
		t.Error("Search doesn't work as expected")
	}
}
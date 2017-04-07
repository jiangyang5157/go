package binarysearchtree

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
	root := NewNode(compare)

	root.Insert(1)
	root.Insert(2)
	root.Insert(3)
	root.Insert(4)

	if root.Search(3).value != 3 {
		t.Error("Search doesn't work as expected")
	}

	if root.Search(222) != nil {
		t.Error("Search doesn't work as expected")
	}
}

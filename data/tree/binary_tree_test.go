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
	tree := New(compare)

	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)

	findTree := tree.Search(3)
	if findTree.value != 3 {
		t.Error("Search doesn't work as expected")
	}

	findNilTree := tree.Search(222)
	if findNilTree != nil {
		t.Error("Search doesn't work as expected")
	}
}
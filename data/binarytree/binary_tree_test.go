package binarytree

import (
	"testing"
)

func less(a interface{}, b interface{}) bool {
	if a.(int) < b.(int) {
		return true
	} else {
		return false
	}
}

func Test_BinaryTree(t *testing.T) {
	tree := New(less)

	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)

	findTree := tree.Search(3)
	if findTree.data != 3 {
		t.Error("Search doesn't work as expected")
	}

	findNilTree := tree.Search(222)
	if findNilTree != nil {
		t.Error("Search doesn't work as expected")
	}
}
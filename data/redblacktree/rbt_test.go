package redblacktree

import "testing"

type Int int

func (a Int) LessThan(b KeyType) bool {
	return a < b.(Int)
}

func Test_Print(t *testing.T) {
	tree := NewTree()

	tree.Insert(Int(1), "123")
	tree.Insert(Int(3), "234")
	tree.Insert(Int(4), "dfa3")
	tree.Insert(Int(6), "sd4")
	tree.Insert(Int(5), "jcd4")
	tree.Insert(Int(2), "bcd4")
	if tree.Size() != 6 {
		t.Error("Error size")
		return
	}
	tree.Print()
}

func Test_Search(t *testing.T) {
	tree := NewTree()

	tree.Insert(Int(1), "123")
	tree.Insert(Int(3), "234")
	tree.Insert(Int(4), "dfa3")
	tree.Insert(Int(6), "sd4")
	tree.Insert(Int(5), "jcd4")
	tree.Insert(Int(2), "bcd4")

	n := tree.Search(Int(4))
	if n.Value != "dfa3" {
		t.Error("Error value")
		return
	}
	n.Value = "bdsf"
	if n.Value != "bdsf" {
		t.Error("Error value modify")
		return
	}
	value := tree.FindValue(Int(5)).(string)
	if value != "jcd4" {
		t.Error("Error value after modifyed other node")
		return
	}
}

func Test_Delete(t *testing.T) {
	tree := NewTree()

	tree.Insert(Int(1), "123")
	tree.Insert(Int(3), "234")
	tree.Insert(Int(4), "dfa3")
	tree.Insert(Int(6), "sd4")
	tree.Insert(Int(5), "jcd4")
	tree.Insert(Int(2), "bcd4")

	for i := 1; i <= 6; i++ {
		tree.Delete(Int(i))
		if tree.Size() != 6-i {
			t.Error("Delete Error")
		}
	}
	tree.Insert(Int(1), "bcd4")
	tree.Clear()
	tree.Print()
	if tree.FindValue(Int(1)) != nil {
		t.Error("Can't clear")
		return
	}
}

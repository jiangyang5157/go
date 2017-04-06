package redblacktree

import "testing"

type MyElement struct {
	key   int
	value string
}

func MyCompareFunc(a, b Element) int {
	return a.(MyElement).key - b.(MyElement).key
}

func Test_Insert(t *testing.T) {
	tree := NewRbTree(MyCompareFunc)
	tree.insert(MyElement{10, "value_10"})
	tree.insert(MyElement{20, "value_20"})
	tree.insert(MyElement{11, "value_11"})
	tree.insert(MyElement{22, "value_22"})
	tree.insert(MyElement{1, "value_1"})
	tree.insert(MyElement{2, "value_2"})
	if tree.size != 6 {
		t.Error("Error size")
		return
	}
	// tree.Print()
}

func Test_Find(t *testing.T) {
	tree := NewRbTree(MyCompareFunc)
	tree.insert(MyElement{10, "value_10"})
	tree.insert(MyElement{20, "value_20"})
	tree.insert(MyElement{11, "value_11"})
	tree.insert(MyElement{22, "value_22"})
	tree.insert(MyElement{1, "value_1"})
	tree.insert(MyElement{2, "value_2"})
	if tree.find(MyElement{key: 2}) == nil {
		t.Error("Should find")
		return
	}

	if tree.find(MyElement{key: 2222}) != nil {
		t.Error("Should not find")
		return
	}
}

func Test_MinAndMax(t *testing.T) {
	tree := NewRbTree(MyCompareFunc)
	tree.insert(MyElement{10, "value_10"})
	tree.insert(MyElement{20, "value_20"})
	tree.insert(MyElement{11, "value_11"})
	tree.insert(MyElement{22, "value_22"})
	tree.insert(MyElement{1, "value_1"})
	tree.insert(MyElement{2, "value_2"})

	if minimum(tree.root).element.(MyElement).key != 1 {
		t.Error("Should find minimum element: key = 1")
		return
	}

	if maximum(tree.root).element.(MyElement).key != 22 {
		t.Error("Should find maximum element: key = 22")
		return
	}
}

func Test_Next(t *testing.T) {
	tree := NewRbTree(MyCompareFunc)
	tree.insert(MyElement{10, "value_10"})
	tree.insert(MyElement{20, "value_20"})
	tree.insert(MyElement{11, "value_11"})
	tree.insert(MyElement{22, "value_22"})
	tree.insert(MyElement{1, "value_1"})
	tree.insert(MyElement{2, "value_2"})

	// tree.Print()
	n := tree.find(MyElement{key: 1})
	// n.Print() // 1
	// n.next().Print() // 2
	// n.next().next().Print() // 10
	// n.next().next().next().Print() // 11
	// n.next().next().next().next().Print() // 20
	// n.next().next().next().next().next().Print() // 22
	if n.next().next().next().next().next().next() != nil {
		t.Error("Should not have next value")
		return
	}
}

func Test_Previous(t *testing.T) {
	tree := NewRbTree(MyCompareFunc)
	tree.insert(MyElement{10, "value_10"})
	tree.insert(MyElement{20, "value_20"})
	tree.insert(MyElement{11, "value_11"})
	tree.insert(MyElement{22, "value_22"})
	tree.insert(MyElement{1, "value_1"})
	tree.insert(MyElement{2, "value_2"})

	// tree.Print()
	n := tree.find(MyElement{key: 22})
	// n.Print() // 22
	// n.previous().Print() // 20
	// n.previous().previous().Print() // 11
	// n.previous().previous().previous().Print() // 10
	// n.previous().previous().previous().previous().Print() // 2
	// n.previous().previous().previous().previous().previous().Print() // 1
	if n.previous().previous().previous().previous().previous().previous() != nil {
		t.Error("Should not have previous value")
		return
	}
}

func Test_Delete(t *testing.T) {
	tree := NewRbTree(MyCompareFunc)
	tree.insert(MyElement{10, "value_10"})
	tree.insert(MyElement{20, "value_20"})
	tree.insert(MyElement{11, "value_11"})
	tree.insert(MyElement{22, "value_22"})
	tree.insert(MyElement{1, "value_1"})
	tree.insert(MyElement{2, "value_2"})

	tree.delete(tree.find(MyElement{key: 11}))
	// tree.print()
}

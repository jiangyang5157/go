package binarysearchtree

/*
BST (ordered) is good at search: balanced BST O(log(n))
Worst case, becomes "linked list"
*/

type Comparable func(a interface{}, b interface{}) int

type Node struct {
	value   interface{}
	compare Comparable
	left    *Node
	right   *Node
}

func New(compare Comparable) *Node {
	return &Node{compare: compare}
}

// O(log(n))
func (n *Node) Insert(value interface{}) {
	if n.value == nil {
		n.value = value
		n.right = New(n.compare)
		n.left = New(n.compare)
	} else {
		if n.compare(value, n.value) < 0 {
			n.left.Insert(value)
		} else {
			n.right.Insert(value)
		}
	}
}

// O(log(n))
func (n *Node) Search(value interface{}) *Node {
	if n.value == nil {
		return nil
	} else if n.value == value {
		return n
	} else {
		if n.compare(value, n.value) < 0 {
			return n.left.Search(value)
		} else {
			return n.right.Search(value)
		}
	}
}

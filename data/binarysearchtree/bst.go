package binarysearchtree

/*
BST (ordered) is good at search: balanced BST O(log(n))
Worst case, becomes "linked list"
*/

type CompareFunc func(a interface{}, b interface{}) int

type Node struct {
	left, right *Node
	value       interface{}
	compare     CompareFunc
}

func NewNode(compare CompareFunc) *Node {
	return &Node{compare: compare}
}

// O(log(n))
func (n *Node) Insert(value interface{}) {
	if value == nil {
		return
	}

	if n.value == nil {
		n.value = value
		n.right = NewNode(n.compare)
		n.left = NewNode(n.compare)
	}

	x := n.compare(value, n.value)
	if x < 0 {
		n.left.Insert(value)
	} else if x > 0 {
		n.right.Insert(value)
	} else {
		// exist, new node is not required
		return
	}
}

// O(log(n))
func (n *Node) Search(value interface{}) *Node {
	if value == nil {
		return nil
	}

	if n.value == nil {
		return nil
	}

	x := n.compare(value, n.value)
	if x < 0 {
		return n.left.Search(value)
	} else if x > 0 {
		return n.right.Search(value)
	} else {
		return n
	}
}

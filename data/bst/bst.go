package bst

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

func (node *Node) Insert(value interface{}) {
	if node.value == nil {
		node.value = value
		node.right = New(node.compare)
		node.left = New(node.compare)
	} else {
		if node.compare(value, node.value) < 0 {
			node.left.Insert(value)
		} else {
			node.right.Insert(value)
		}
	}
}

func (node *Node) Search(value interface{}) *Node {
	if node.value == nil {
		return nil
	} else if node.value == value {
		return node
	} else {
		if node.compare(value, node.value) < 0 {
			return node.left.Search(value)
		} else {
			return node.right.Search(value)
		}
	}
}

package binarytree

type Comparable func(a interface{}, b interface{}) int

type Node struct {
	data    interface{}
	compare Comparable
	left    *Node
	right   *Node
}

func New(compare Comparable) *Node {
	return &Node{compare: compare}
}

func (node *Node) IsEmpty() bool {
	return node.data == nil
}

func (node *Node) Insert(data interface{}) {
	if node.IsEmpty() {
		node.data = data
		node.right = New(node.compare)
		node.left = New(node.compare)
	} else {
		if node.compare(data, node.data) < 0 {
			node.left.Insert(data)
		} else {
			node.right.Insert(data)
		}
	}
}

func (node *Node) Search(data interface{}) *Node {
	if node.IsEmpty() {
		return nil
	}
	if node.data == data {
		return node
	} else {
		if node.compare(data, node.data) < 0 {
			return node.left.Search(data)
		} else {
			return node.right.Search(data)
		}
	}
}
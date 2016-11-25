package binarytree

type Comparable func(a interface{}, b interface{}) bool

type BinaryTree struct {
	node  interface{}
	left  *BinaryTree
	right *BinaryTree
	less  Comparable
}

func New(less Comparable) *BinaryTree {
	tree := &BinaryTree{}
	tree.less = less
	return tree
}

func (tree *BinaryTree) IsEmpty() bool {
	return tree.node == nil
}

func (tree *BinaryTree) Insert(node interface{}) {
	if tree.IsEmpty() {
		tree.node = node
		tree.right = New(tree.less)
		tree.left = New(tree.less)
	} else {
		if tree.less(node, tree.node) {
			tree.left.Insert(node)
		} else {
			tree.right.Insert(node)
		}
	}
}

func (tree *BinaryTree) Search(node interface{}) *BinaryTree {
	if tree.IsEmpty() {
		return nil
	}
	if tree.node == node {
		return tree
	} else {
		if tree.less(node, tree.node) {
			return tree.left.Search(node)
		} else {
			return tree.right.Search(node)
		}
	}
}
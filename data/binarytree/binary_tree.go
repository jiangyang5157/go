package binarytree

type Comparable func(a interface{}, b interface{}) bool

type BinaryTree struct {
	data  interface{}
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
	return tree.data == nil
}

func (tree *BinaryTree) Insert(data interface{}) {
	if tree.IsEmpty() {
		tree.data = data
		tree.right = New(tree.less)
		tree.left = New(tree.less)
	} else {
		if tree.less(data, tree.data) {
			tree.left.Insert(data)
		} else {
			tree.right.Insert(data)
		}
	}
}

func (tree *BinaryTree) Search(data interface{}) *BinaryTree {
	if tree.IsEmpty() {
		return nil
	}
	if tree.data == data {
		return tree
	} else {
		if tree.less(data, tree.data) {
			return tree.left.Search(data)
		} else {
			return tree.right.Search(data)
		}
	}
}
package redblacktree

import "fmt"

const (
	RED = 0
	BLACK = 1
)

type ValueType interface{}

type KeyType interface {
	LessThan(interface{}) bool
}

type node struct {
	left, right, parent *node
	color               int
	Key                 KeyType
	Value               ValueType
}

type Tree struct {
	root *node
	size int
}

// Return a new rbtree
func NewTree() *Tree {
	return &Tree{}
}

// Return the size of the rbtree
func (t *Tree) Size() int {
	return t.size
}

// Clear the rbtree
func (t *Tree) Clear() {
	t.root = nil
	t.size = 0
}

// Whether the rbtree is empty
func (t *Tree) IsEmpty() bool {
	return t.root == nil
}

// Get color of the node
func getColor(n *node) int {
	if n == nil {
		return BLACK
	}
	return n.color
}

// Find the minimum node of subtree n.
func minimum(n *node) *node {
	for n.left != nil {
		n = n.left
	}
	return n
}

// Find the maximum node of subtree n.
func maximum(n *node) *node {
	for n.right != nil {
		n = n.right
	}
	return n
}

// Create the rbtree's iterator that points to the minmum node
func (t *Tree) Iterator() *node {
	return minimum(t.root)
}

// Find the node by key and return it,if not exists return nil
func (t *Tree) findNode(key KeyType) *node {
	for x := t.root; x != nil; {
		if key.LessThan(x.Key) {
			x = x.left
		} else {
			if key == x.Key {
				return x
			} else {
				x = x.right
			}
		}
	}
	return nil
}

// Find the node and return it as a iterator
func (t *Tree) FindIt(key KeyType) *node {
	return t.findNode(key)
}

// Find the node and return its value
func (t *Tree) FindValue(key KeyType) ValueType {
	n := t.findNode(key)
	if n != nil {
		return n.Value
	}
	return nil
}

// Insert the key-value pair into the rbtree
func (t *Tree) Insert(key KeyType, value ValueType) {
	var y *node
	for x := t.root; x != nil; {
		y = x
		if key.LessThan(x.Key) {
			x = x.left
		} else {
			x = x.right
		}
	}

	z := &node{parent: y, color: RED, Key: key, Value: value}
	t.size += 1

	if y == nil {
		z.color = BLACK
		t.root = z
		return
	} else if z.Key.LessThan(y.Key) {
		y.left = z
	} else {
		y.right = z
	}
	t.rb_insert_fixup(z)

}

// Delete the node by key
func (t *Tree) Delete(key KeyType) {
	z := t.findNode(key)
	if z == nil {
		return
	}

	var x, y, parent *node
	y = z
	y_original_color := y.color
	parent = z.parent
	if z.left == nil {
		x = z.right
		t.transplant(z, z.right)
	} else if z.right == nil {
		x = z.left
		t.transplant(z, z.left)
	} else {
		y = minimum(z.right)
		y_original_color = y.color
		x = y.right

		if y.parent == z {
			if x == nil {
				parent = y
			} else {
				x.parent = y
			}
		} else {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}
	if y_original_color == BLACK {
		t.rb_delete_fixup(x, parent)
	}
	t.size -= 1
}


// Return the node's successor as an iterator
func (n *node) Next() *node {
	return successor(n)
}

// Return the successor of the node
func successor(x *node) *node {
	if x.right != nil {
		return minimum(x.right)
	}
	y := x.parent
	for y != nil && x == y.right {
		x = y
		y = x.parent
	}
	return y
}


// Transplant the subtree u and v
func (t *Tree) transplant(u, v *node) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v == nil {
		return
	}
	v.parent = u.parent
}

func (t *Tree) rb_insert_fixup(z *node) {
	var y *node
	for z.parent != nil && z.parent.color == RED {
		if z.parent == z.parent.parent.left {
			y = z.parent.parent.right
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = BLACK
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.left_rotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.right_rotate(z.parent.parent)
			}
		} else {
			y = z.parent.parent.left
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.right_rotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.left_rotate(z.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}

func (t *Tree) rb_delete_fixup(x, parent *node) {
	var w *node

	for x != t.root && getColor(x) == BLACK {
		if x != nil {
			parent = x.parent
		}
		if x == parent.left {
			w = parent.right
			if w.color == RED {
				w.color = BLACK
				parent.color = RED
				t.left_rotate(x.parent)
				w = parent.right
			}
			if getColor(w.left) == BLACK && getColor(w.right) == BLACK {
				w.color = RED
				x = parent
			} else {
				if getColor(w.right) == BLACK {
					if w.left != nil {
						w.left.color = BLACK
					}
					w.color = RED
					t.right_rotate(w)
					w = parent.right
				}
				w.color = parent.color
				parent.color = BLACK
				if w.right != nil {
					w.right.color = BLACK
				}
				t.left_rotate(parent)
				x = t.root
			}
		} else {
			w = parent.left
			if w.color == RED {
				w.color = BLACK
				parent.color = RED
				t.right_rotate(parent)
				w = parent.left
			}
			if getColor(w.left) == BLACK && getColor(w.right) == BLACK {
				w.color = RED
				x = parent
			} else {
				if getColor(w.left) == BLACK {
					if w.right != nil {
						w.right.color = BLACK
					}
					w.color = RED
					t.left_rotate(w)
					w = parent.left
				}
				w.color = parent.color
				parent.color = BLACK
				if w.left != nil {
					w.left.color = BLACK
				}
				t.right_rotate(parent)
				x = t.root
			}
		}
	}
	if x != nil {
		x.color = BLACK
	}
}

func (t *Tree) left_rotate(x *node) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *Tree) right_rotate(x *node) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = x
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.right = x
	x.parent = y
}

func (n *node) preorder() {
	fmt.Printf("(%v %v)", n.Key, n.Value)
	if n.parent == nil {
		fmt.Printf("nil")
	} else {
		fmt.Printf("whose parent is %v", n.parent.Key)
	}
	if n.color == RED {
		fmt.Println(" and color RED")
	} else {
		fmt.Println(" and color BLACK")
	}
	if n.left != nil {
		fmt.Printf("%v's left child is ", n.Key)
		n.left.preorder()
	}
	if n.right != nil {
		fmt.Printf("%v's right child is ", n.Key)
		n.right.preorder()
	}
}

func (t *Tree) Preorder() {
	fmt.Println("preorder begin!")
	if t.root != nil {
		t.root.preorder()
	}
	fmt.Println("preorder end!")
}
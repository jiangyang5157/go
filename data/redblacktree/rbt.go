package redblacktree

/*
It is a special BST, where:
+ Each node is either BLACK or RED
+ Root is BLACK
+ All leaves (nil children) are BLACK
+ Every RED node has two BLACK children
+ Every path from root to leaf has the same number of BLACK nodes, the number also called "black-height"

It is not a completed balance BST, but it guarantees a almost-balanced BST:
+ It sacrifices balanced-hight for a O(log(n)) Insert/Delete/Search
+ Any un-almost-balanced situation caused by Insert/Delete can be fix in maximum 3 times rotation

eg: JDK.TreeMap, JDK.TreeSet
*/

import "fmt"

const (
	RED   = 0
	BLACK = 1
)

type KeyType interface {
	LessThan(interface{}) bool
}

type ValueType interface{}

type node struct {
	color               int
	left, right, parent *node
	Key                 KeyType
	Value               ValueType
}

type Tree struct {
	root *node
	size int
}

func (n *node) Print() {
	fmt.Printf("[%v]{value=%v, ", n.Key, n.Value)

	if n.color == RED {
		fmt.Printf("color=RED, ")
	} else {
		fmt.Printf("color=BLACK, ")
	}

	if n.parent == nil {
		fmt.Printf("parent=nil\n")
	} else {
		fmt.Printf("parent.key=%v\n", n.parent.Key)
	}

	if n.left != nil {
		fmt.Printf("left=\n")
		n.left.Print()
	}
	if n.right != nil {
		fmt.Printf("right=\n")
		n.right.Print()
	}
	fmt.Printf("}[%v]\n", n.Key)
}

func (t *Tree) Print() {
	if t.root != nil {
		t.root.Print()
	}
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

func (t *Tree) Find(key KeyType) *node {
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

// Find the minimum node of the subtree-n: n, n.left, n.right.
func minimum(n *node) *node {
	for n.left != nil {
		n = n.left
	}
	return n
}

// Find the maximum node of the subtree-n: n, n.left, n.right.
func maximum(n *node) *node {
	for n.right != nil {
		n = n.right
	}
	return n
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

func (t *Tree) Iterator() *node {
	return minimum(t.root)
}

// Return the node's successor
func (n *node) Next() *node {
	return successor(n)
}

// Return the successor of the node
func successor(x *node) *node {
	if x.right != nil {
		// successor is the minimum node of the subtree-x.right:
		return minimum(x.right)
	}
	y := x.parent
	for y != nil && x == y.right {
		// Skip parent that less than x
		x = y
		y = y.parent
	}
	return y
}

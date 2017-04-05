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
	LessThan(KeyType) bool
}

func lessThan(a, b KeyType) bool {
	return a.LessThan(b)
}

type ValueType interface{}

type node struct {
	color               int
	left, right, parent *node
	Key                 KeyType
	Value               ValueType
}

type RbTree struct {
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

func (t *RbTree) Print() {
	if t.root != nil {
		t.root.Print()
	}
}

// Return a new rbtree
func NewTree() *RbTree {
	return &RbTree{}
}

// Return the size of the rbtree
func (t *RbTree) Size() int {
	return t.size
}

// Clear the rbtree
func (t *RbTree) Clear() {
	t.root = nil
	t.size = 0
}

// Whether the rbtree is empty
func (t *RbTree) IsEmpty() bool {
	return t.root == nil
}

// Get color of the node
func getColor(n *node) int {
	if n == nil {
		return BLACK
	}
	return n.color
}

func (t *RbTree) Search(key KeyType) *node {
	x := t.root
	for x != nil {
		if key.LessThan(x.Key) {
			x = x.left
		} else if x.Key.LessThan(key) {
			x = x.right
		} else {
			break
		}

	}
	return x
}

// Find the node and return its value
func (t *RbTree) FindValue(key KeyType) ValueType {
	n := t.Search(key)
	if n != nil {
		return n.Value
	}
	return nil
}

func minimum(n *node) *node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func maximum(n *node) *node {
	for n.right != nil {
		n = n.right
	}
	return n
}

// Return the successor of the node
func successor(x *node) *node {
	if x.right != nil {
		// Get the minimum from the right sub-tree if it existed.
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

// Insert the key-value pair into the rbtree
func (t *RbTree) Insert(key KeyType, value ValueType) *node {
	x := t.root
	var y *node
	for x != nil {
		y = x
		if key.LessThan(x.Key) {
			x = x.left
		} else if x.Key.LessThan(key) {
			x = x.right
		} else {
			// not allow to override existing node
			return x
		}
	}

	z := &node{parent: y, color: RED, Key: key, Value: value}
	if y == nil {
		z.color = BLACK
		t.root = z
	} else if z.Key.LessThan(y.Key) {
		y.left = z
	} else {
		y.right = z
	}
	t.size++
	t.insertFixup(z)
	return z
}

func (t *RbTree) insertFixup(z *node) {
	var y *node
	for z.parent != nil && z.parent.color == RED {
		if z.parent == z.parent.parent.left {
			y = z.parent.parent.right
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.leftRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.rightRotate(z.parent.parent)
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
					t.rightRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.leftRotate(z.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}

// Delete the node by key
func (t *RbTree) Delete(key KeyType) {
	z := t.Search(key)
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
		t.deleteFixup(x, parent)
	}
	t.size -= 1
}

// Transplant the subtree u and v
func (t *RbTree) transplant(u, v *node) {
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

func (t *RbTree) deleteFixup(x, parent *node) {
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
				t.leftRotate(x.parent)
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
					t.rightRotate(w)
					w = parent.right
				}
				w.color = parent.color
				parent.color = BLACK
				if w.right != nil {
					w.right.color = BLACK
				}
				t.leftRotate(parent)
				x = t.root
			}
		} else {
			w = parent.left
			if w.color == RED {
				w.color = BLACK
				parent.color = RED
				t.rightRotate(parent)
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
					t.leftRotate(w)
					w = parent.left
				}
				w.color = parent.color
				parent.color = BLACK
				if w.left != nil {
					w.left.color = BLACK
				}
				t.rightRotate(parent)
				x = t.root
			}
		}
	}
	if x != nil {
		x.color = BLACK
	}
}

// func (t *Rbtree) delete(key *Node) *Node {
// 	z := t.search(key)
//
// 	if z == t.NIL {
// 		return t.NIL
// 	}
// 	ret := &Node{t.NIL, t.NIL, t.NIL, z.Color, z.Item}
//
// 	var y *Node
// 	var x *Node
//
// 	if z.Left == t.NIL || z.Right == t.NIL {
// 		y = z
// 	} else {
// 		y = t.successor(z)
// 	}
//
// 	if y.Left != t.NIL {
// 		x = y.Left
// 	} else {
// 		x = y.Right
// 	}
//
// 	// Even if x is NIL, we do the assign. In that case all the NIL nodes will
// 	// change from {nil, nil, nil, BLACK, nil} to {nil, nil, ADDR, BLACK, nil},
// 	// but do not worry about that because it will not affect the compare
// 	// between Node-X with Node-NIL
// 	x.Parent = y.Parent
//
// 	if y.Parent == t.NIL {
// 		t.root = x
// 	} else if y == y.Parent.Left {
// 		y.Parent.Left = x
// 	} else {
// 		y.Parent.Right = x
// 	}
//
// 	if y != z {
// 		z.Item = y.Item
// 	}
//
// 	if y.Color == BLACK {
// 		t.deleteFixup(x)
// 	}
//
// 	t.count--
//
// 	return ret
// }
//
// func (t *Rbtree) deleteFixup(x *Node) {
// 	for x != t.root && x.Color == BLACK {
// 		if x == x.Parent.Left {
// 			w := x.Parent.Right
// 			if w.Color == RED {
// 				w.Color = BLACK
// 				x.Parent.Color = RED
// 				t.leftRotate(x.Parent)
// 				w = x.Parent.Right
// 			}
// 			if w.Left.Color == BLACK && w.Right.Color == BLACK {
// 				w.Color = RED
// 				x = x.Parent
// 			} else {
// 				if w.Right.Color == BLACK {
// 					w.Left.Color = BLACK
// 					w.Color = RED
// 					t.rightRotate(w)
// 					w = x.Parent.Right
// 				}
// 				w.Color = x.Parent.Color
// 				x.Parent.Color = BLACK
// 				w.Right.Color = BLACK
// 				t.leftRotate(x.Parent)
// 				// this is to exit while loop
// 				x = t.root
// 			}
// 		} else { // the code below is has left and right switched from above
// 			w := x.Parent.Left
// 			if w.Color == RED {
// 				w.Color = BLACK
// 				x.Parent.Color = RED
// 				t.rightRotate(x.Parent)
// 				w = x.Parent.Left
// 			}
// 			if w.Left.Color == BLACK && w.Right.Color == BLACK {
// 				w.Color = RED
// 				x = x.Parent
// 			} else {
// 				if w.Left.Color == BLACK {
// 					w.Right.Color = BLACK
// 					w.Color = RED
// 					t.leftRotate(w)
// 					w = x.Parent.Left
// 				}
// 				w.Color = x.Parent.Color
// 				x.Parent.Color = BLACK
// 				w.Left.Color = BLACK
// 				t.rightRotate(x.Parent)
// 				x = t.root
// 			}
// 		}
// 	}
// 	x.Color = BLACK
// }

/*
      |                         |
      A                         C
     / \      left rotate      / \
  	B  C       -------->      A  E
      / \                    / \
     D  E                   B  D
*/
func (t *RbTree) leftRotate(x *node) {
	if x.right == nil {
		// left rotation, the right child should not be nil
		return
	}
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

/*
      |                         |
      A                         B
     / \      right rotate     / \
  	B  C       -------->      D  A
   / \                          / \
  D  E                         E  C
*/
func (t *RbTree) rightRotate(x *node) {
	if x.left == nil {
		// right rotation, the left child should not be nil
		return
	}
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

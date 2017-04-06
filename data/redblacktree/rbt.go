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

//https://github.com/yasushi-saito/rbtree/blob/master/rbtree.go
type KeyType interface {
	LessThan(KeyType) bool
}

type ValueType interface{}

func lessThan(a, b KeyType) bool {
	return a.LessThan(b)
}

type Node struct {
	color               int
	left, right, parent *Node
	Key                 KeyType
	Value               ValueType
}

type RbTree struct {
	root *Node
	size int
}

func (t *RbTree) Print() {
	if t.root != nil {
		t.root.Print()
	}
}

func (n *Node) Print() {
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

func NewTree() *RbTree {
	return &RbTree{root: nil, size: 0}
}

func (n *Node) isLeftChild() bool {
	if n == nil || n.parent == nil {
		return false
	} else {
		return n == n.parent.left
	}
}

func (n *Node) isRightChild() bool {
	if n == nil || n.parent == nil {
		return false
	} else {
		return n == n.parent.right
	}
}

func (t *RbTree) Size() int {
	return t.size
}

func (t *RbTree) IsEmpty() bool {
	return t.Size() == 0
}

func (t *RbTree) Clear() {
	t.root = nil
	t.size = 0
}

func minimum(n *Node) *Node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func maximum(n *Node) *Node {
	for n.right != nil {
		n = n.right
	}
	return n
}

/*
      |
      A
     / \
  	B  C
      / \
     D  E

		 A is D/E's grandparent
*/
func (n *Node) grandparent() *Node {
	if n != nil && n.parent != nil {
		return n.parent.parent
	} else {
		return nil
	}
}

/*
      |
      A
     / \
  	B  C
      / \
     D  E

		 B is D/E's uncle
*/
func (n *Node) uncle() *Node {
	g := n.grandparent()
	if g == nil {
		return nil
	}
	if n.parent.isLeftChild() {
		return g.right
	} else {
		return g.left
	}
}

/*
      |
      A
     / \
  	B  C

		C is B's sibling
*/
func (n *Node) sibling() *Node {
	if n == nil || n.parent == nil {
		return nil
	}
	if n.isLeftChild() {
		return n.parent.right
	} else {
		return n.parent.left
	}
}

/*
      |                         |
      A       right rotate      B
     / \       -------->       / \
  	B  C       <--------      D  A
   / \        left rotate       / \
  D  E                         E  C
*/
func (t *RbTree) rightRotate(x *Node) {
	y := x.left
	if y == nil {
		// right rotation, the left child should not be nil
		return
	}

	// Turn y's right sub-tree into x's left sub-tree
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}

	// y's new parent was x's parent
	y.parent = x.parent

	// Set x's old parent to point to y instead of x
	if x.parent == nil {
		// if we're at the root
		t.root = x
	} else {
		if x == x.parent.left {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
	}

	// x's new parent is y
	x.parent = y
	// put x on y's right
	y.right = x
}

/*
      |                         |
      y       right rotate      x
     / \       -------->       / \
  	x  C       <--------      A  y
   / \        left rotate       / \
  A  B                         B  C
*/
func (t *RbTree) leftRotate(x *Node) {
	y := x.right
	if y == nil {
		// left rotation, the right child should not be nil
		return
	}

	// Turn y's left sub-tree into x's right sub-tree
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}

	// y's new parent was x's parent
	y.parent = x.parent

	// Set x's old parent to point to y instead of x
	if x.parent == nil {
		// if we're at the root
		t.root = y
	} else {
		if x == x.parent.left {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
	}

	// x's new parent is y
	x.parent = y
	// put x on y's left
	y.left = x
}

// ################################################################

func (t *RbTree) search(key KeyType) *Node {
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
func (t *RbTree) SearchValue(key KeyType) ValueType {
	n := t.search(key)
	if n != nil {
		return n.Value
	}
	return nil
}

// insert the key-value pair into the rbtree
func (t *RbTree) insert(key KeyType, value ValueType) *Node {
	var x, y, z *Node = t.root, nil, nil
	// locate z's parent y
	for x != nil {
		y = x
		if key.LessThan(x.Key) {
			x = x.left
		} else if x.Key.LessThan(key) {
			x = x.right
		} else {
			// return existing node has same key
			return x
		}
	}

	z = &Node{parent: y, color: RED, Key: key, Value: value}
	if y == nil {
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

func (t *RbTree) insertFixup(n *Node) {
	for n.parent != nil && n.parent.color == RED {
		// Since the color of the parent is RED, so the parent is not root.
		// The grandparent must be exist.
		g := n.grandparent()
		if n.parent == g.left {
			uncle := g.right
			if uncle != nil && uncle.color == RED {
				// Parent and uncle are both RED, the grandparent must be BLACK
				n.parent.color = BLACK
				uncle.color = BLACK
				g.color = RED
				n = g
			} else {
				if n == n.parent.right {
					// Parent is RED and uncle is BLACK, the current node is right child
					n = n.parent
					t.leftRotate(n)
				}
				n.parent.color = BLACK
				g.color = RED
				t.rightRotate(g)
			}
		} else {
			uncle := g.left
			if uncle != nil && uncle.color == RED {
				n.parent.color = BLACK
				uncle.color = BLACK
				g.color = RED
				n = g
			} else {
				if n == n.parent.left {
					n = n.parent
					t.rightRotate(n)
				}
				n.parent.color = BLACK
				g.color = RED
				t.leftRotate(g)
			}
		}
	}
	t.root.color = BLACK
}

// Delete the node by key
func (t *RbTree) Delete(key KeyType) {
	z := t.search(key)
	if z == nil {
		return
	}

	var x, y, parent *Node
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
func (t *RbTree) transplant(u, v *Node) {
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

func (t *RbTree) deleteFixup(x, parent *Node) {
	var w *Node
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

// Create the rbtree's iterator that points to the minmum node
func (t *RbTree) Iterator() *Node {
	return minimum(t.root)
}

func (n *Node) Next() *Node {
	return successor(n)
}

// Return the successor of the node
func successor(x *Node) *Node {
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

// // Return the minimum node that's larger than N. Return nil if no such
// // node is found.
// func (n *node) doNext() *node {
// 	if n.right != nil {
// 		m := n.right
// 		for m.left != nil {
// 			m = m.left
// 		}
// 		return m
// 	}
//
// 	for n != nil {
// 		p := n.parent
// 		if p == nil {
// 			return nil
// 		}
// 		if n.isLeftChild() {
// 			return p
// 		}
// 		n = p
// 	}
// 	return nil
// }
// // Return the maximum node that's smaller than N. Return nil if no
// // such node is found.
// func (n *node) doPrev() *node {
// 	if n.left != nil {
// 		return maxPredecessor(n)
// 	}
//
// 	for n != nil {
// 		p := n.parent
// 		if p == nil {
// 			break
// 		}
// 		if n.isRightChild() {
// 			return p
// 		}
// 		n = p
// 	}
// 	return negativeLimitNode
// }
// // Return the predecessor of "n".
// func maxPredecessor(n *node) *node {
// 	doAssert(n.left != nil)
// 	m := n.left
// 	for m.right != nil {
// 		m = m.right
// 	}
// 	return m
// }

// Get color of the node
func getColor(n *Node) int {
	if n == nil {
		return BLACK
	}
	return n.color
}

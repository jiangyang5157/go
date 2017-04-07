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
	BLACK = 0
	RED   = 1
)

type Element interface{}

type CompareFunc func(a, b Element) int

type Node struct {
	left, right, parent *Node
	element             Element
	color               int
}

type RbTree struct {
	root    *Node
	size    int
	compare CompareFunc
}

func NewRbTree(compare CompareFunc) *RbTree {
	if compare == nil {
		return nil
	}
	return &RbTree{root: nil, size: 0, compare: compare}
}

func (n *Node) isLeftChild() bool {
	if n.parent == nil {
		return false
	}
	return n == n.parent.left
}

func (n *Node) isRightChild() bool {
	if n.parent == nil {
		return false
	}
	return n == n.parent.right
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
	if n.parent == nil {
		return nil
	}
	return n.parent.parent
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
	if n.parent == nil {
		return nil
	}
	if n.isLeftChild() {
		return n.parent.right
	} else {
		return n.parent.left
	}
}

// Traverse from the node to left recursively until left is nil.
func minimum(n *Node) *Node {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}
	return n
}

// Traverse from the node to right recursively until right is nil.
func maximum(n *Node) *Node {
	if n == nil {
		return nil
	}
	for n.right != nil {
		n = n.right
	}
	return n
}

func (n *Node) next() *Node {
	if n.right != nil {
		return minimum(n.right)
	}
	p := n.parent
	for n.isRightChild() {
		n = p
		p = p.parent
	}
	return p
}

func (n *Node) previous() *Node {
	if n.left != nil {
		return maximum(n.left)
	}
	p := n.parent
	for n.isLeftChild() {
		n = p
		p = p.parent
	}
	return p
}

func (t *RbTree) find(item Element) *Node {
	n := t.root
	for n != nil {
		x := t.compare(item, n.element)
		if x < 0 {
			n = n.left
		} else if x > 0 {
			n = n.right
		} else {
			break
		}
	}
	return n
}

/*
      |                         |
      y       right rotate      x
     / \       -------->       / \
  	x  C       <--------      A  y
   / \        left rotate       / \
  A  B                         B  C
*/
func (t *RbTree) rightRotate(x *Node) {
	if x == nil {
		return
	}
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
		if x.isLeftChild() {
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
	if x == nil {
		return
	}
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
		if x.isLeftChild() {
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

func (t *RbTree) insert(item Element) *Node {
	if item == nil {
		return nil
	}

	var node *Node
	if t.root == nil {
		// new node is root
		node = &Node{parent: nil, element: item}
		t.root = node
		t.size++
	} else {
		parent := t.root
		for {
			x := t.compare(item, parent.element)
			if x < 0 {
				if parent.left == nil {
					node = &Node{parent: parent, element: item}
					parent.left = node
					t.size++
					break
				} else {
					parent = parent.left
				}
			} else if x > 0 {
				if parent.right == nil {
					node = &Node{parent: parent, element: item}
					parent.right = node
					t.size++
					break
				} else {
					parent = parent.right
				}
			} else {
				// exist, new node is not required
				break
			}
		}
	}

	t.insertFixup(node)
	return node
}

func (t *RbTree) insertFixup(n *Node) {
	if n == nil {
		return
	}

	n.color = RED
	for {
		if n.parent == nil {
			// n is root
			n.color = BLACK
			break
		}
		if n.parent.color == BLACK {
			// parent is BLACK, nothing required
			break
		}
		// parent is RED

		grandparent := n.grandparent()
		uncle := n.uncle()
		if uncle != nil && uncle.color == RED {
			// If parent and uncle are both RED, paint both BLACK and make grandparent RED. Then repeat.
			n.parent.color = BLACK
			uncle.color = BLACK
			grandparent.color = RED
			n = grandparent
			continue
		}
		// parent is RED, uncle is BLACK - step 1

		if n.isRightChild() && n.parent.isLeftChild() {
			t.leftRotate(n.parent)
			n = n.left
			continue
		}
		if n.isLeftChild() && n.parent.isRightChild() {
			t.rightRotate(n.parent)
			n = n.right
			continue
		}
		// parent is RED, uncle is BLACK - step 2

		n.parent.color = BLACK
		grandparent.color = RED
		if n.isLeftChild() {
			t.rightRotate(grandparent)
		} else {
			t.leftRotate(grandparent)
		}

		break
	}
}

func (t *RbTree) delete(n *Node) {
	if n == nil {
		return
	}

	parent := n.parent
	var x *Node
	y := n
	yColor := y.color
	if n.left == nil {
		x = n.right
		t.transplant(n, x)
	} else if n.right == nil {
		x = n.left
		t.transplant(n, x)
	} else {
		y = minimum(n.right)
		yColor = y.color
		x = y.right

		if y.parent == n {
			if x == nil {
				parent = y
			} else {
				x.parent = y
			}
		} else {
			t.transplant(y, y.right)
			y.right = n.right
			y.right.parent = y
		}
		t.transplant(n, y)
		y.left = n.left
		y.left.parent = y
		y.color = n.color
	}
	if yColor == BLACK {
		t.deleteFixup(x, parent)
	}
	t.size--
}

// Transplant the subtree u and v
func (t *RbTree) transplant(u, v *Node) {
	if u.parent == nil {
		t.root = v
	} else if u.isLeftChild() {
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
		if x.isLeftChild() {
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

func getColor(n *Node) int {
	if n == nil {
		return BLACK
	}
	return n.color
}

func (t *RbTree) print() {
	if t.root != nil {
		t.root.print()
	}
}

func (n *Node) print() {
	fmt.Printf("%v{", n.element)

	fmt.Printf("color=%v, ", n.color)

	if n.parent == nil {
		fmt.Printf("parent=nil\n")
	} else {
		fmt.Printf("parent=%v\n", n.parent.element)
	}

	if n.left != nil {
		fmt.Printf("left=\n")
		n.left.print()
	}

	if n.right != nil {
		fmt.Printf("right=\n")
		n.right.print()
	}

	fmt.Printf("}%v\n", n.element)
}

package tries

type Node struct {
	parent    *Node
	children  map[byte]*Node
	completed bool
}

func NewTries() *Node {
	root := newNode()
	root.completed = true
	return root
}

func newNode() *Node {
	return &Node{children: make(map[byte]*Node)}
}

func (n *Node) Add(word string) *Node {
	return n.add([]byte(word))
}

func (n *Node) add(bytes []byte) *Node {
	var currNode *Node = n
	for i := 0; i < len(bytes); i++ {
		c := bytes[i]
		child, exist := currNode.children[c]
		if exist {
			currNode = child
			return currNode.add(bytes[i+1:])
		} else {
			child = newNode()
			child.parent = currNode
			currNode.children[c] = child
			currNode = child
		}
	}
	currNode.completed = true
	return currNode
}

func (n *Node) Search(word string) *Node {
	return n.search([]byte(word))
}

func (n *Node) search(bytes []byte) *Node {
	currNode := n.searchBeginWith(bytes)
	if currNode == nil {
		return nil
	}
	if currNode.completed != true {
		return nil
	}
	return currNode
}

func (n *Node) SearchBeginWith(beginWith string) *Node {
	return n.searchBeginWith([]byte(beginWith))
}

func (n *Node) searchBeginWith(bytes []byte) *Node {
	var currNode *Node = n
	for i := 0; i < len(bytes); i++ {
		c := bytes[i]
		child, exist := currNode.children[c]
		if exist {
			currNode = child
			return currNode.searchBeginWith(bytes[i+1:])
		} else {
			return nil
		}
	}
	return currNode
}

func (n *Node) Remove(word string) bool {
	return n.remove([]byte(word))
}

func (n *Node) remove(bytes []byte) bool {
	currNode := n.search(bytes)
	if currNode == nil {
		return false
	}
	currNode.completed = false

	var parentNode *Node
	for i := len(bytes) - 1; i >= 0; i-- {
		c := bytes[i]
		parentNode = currNode.parent
		if len(currNode.children) == 0 && currNode.completed == false {
			delete(parentNode.children, c)
		}
		currNode = parentNode
	}
	return true
}

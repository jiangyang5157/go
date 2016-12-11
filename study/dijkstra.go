package study

import (
	"math"
	"container/heap"
)

// O(|V|^2) Dijkstra's algorithm (Greedy algorithm) is an algorithm for finding the shortest paths between nodes in a graph

type vertex interface{}

type edge struct {
	v1, v2 interface{}
	dist   int
}

type node struct {
	v         interface{} // vertex name
	neighbors []neighbor  // edges from this vertex
	tent      int         // tentative distance
	prev      *node       // previous node in shortest path back to start
	done      bool        // true when tent and prev represent shortest path
	rx        int         // heap remove/fix index
}

type neighbor struct {
	theNode *node
	dist    int
}

type path struct {
	path   []vertex
	length int
}

type nodeHeap []*node

func (nh nodeHeap) Len() int {
	return len(nh)
}
func (nh nodeHeap) Less(i, j int) bool {
	return nh[i].tent < nh[j].tent
}
func (nh nodeHeap) Swap(i, j int) {
	nh[i], nh[j] = nh[j], nh[i]
	nh[i].rx = i
	nh[j].rx = j
}
func (nh *nodeHeap) Push(x interface{}) {
	n := x.(*node)
	n.rx = len(*nh)
	*nh = append(*nh, n)
}
func (nh *nodeHeap) Pop() interface{} {
	nhRef := *nh
	lastIndex := len(nhRef) - 1
	lastNode := nhRef[lastIndex]
	*nh = nhRef[:lastIndex]
	lastNode.rx = -1
	return lastNode
}

func linkGraph(edges []edge, directed bool, start vertex, end vertex) (nodes []*node, startNode *node, endNode *node) {
	all := make(map[vertex]*node)
	for _, e := range edges {
		n1 := all[e.v1]
		n2 := all[e.v2]
		if n1 == nil {
			n1 = &node{v: e.v1}
			all[e.v1] = n1
		}
		if n2 == nil {
			n2 = &node{v: e.v2}
			all[e.v2] = n2
		}
		n1.neighbors = append(n1.neighbors, neighbor{theNode: n2, dist: e.dist})
		if !directed {
			n2.neighbors = append(n2.neighbors, neighbor{theNode: n1, dist: e.dist})
		}
	}
	nodes = make([]*node, len(all))
	var n int
	for _, nd := range all {
		nodes[n] = nd
		n++
	}
	return nodes, all[start], all[end]
}

// dijkstra is a heap-enhanced version of Dijkstra's shortest path algorithm.
// If endNode is specified, only a single path is returned.
// If endNode is nil, paths to all nodes are returned.
// Note input all nodes is needed to efficiently accomplish WP steps 1 and 2.
// This initialization could be done in linkGraph, but is done here to more
// closely follow the WP algorithm.
func dijkstra(nodes []*node, startNode, endNode *node) (ret []path) {
	// WP steps 1 and 2.
	for _, n := range nodes {
		n.tent = math.MaxInt32
		n.done = false
		n.prev = nil
		n.rx = -1
	}
	current := startNode
	current.tent = 0
	var unvis nodeHeap

	for {
		// WP step 3: update tentative distances to neighbors
		for _, nb := range current.neighbors {
			if nd := nb.theNode; !nd.done {
				if d := current.tent + nb.dist; d < nd.tent {
					nd.tent = d
					nd.prev = current
					if nd.rx < 0 {
						heap.Push(&unvis, nd)
					} else {
						heap.Fix(&unvis, nd.rx)
					}
				}
			}
		}
		// WP step 4: mark current node visited, record path and distance
		current.done = true
		if endNode == nil || current == endNode {
			// record path and distance for return value
			distance := current.tent
			// recover path by tracing prev links,
			var p []vertex
			for ; current != nil; current = current.prev {
				p = append(p, current.v)
			}
			// then reverse list
			for i := (len(p) + 1) / 2; i > 0; i-- {
				p[i - 1], p[len(p) - i] = p[len(p) - i], p[i - 1]
			}
			ret = append(ret, path{p, distance}) // pl is return value
			// WP step 5 (case of end node reached)
			if endNode != nil {
				return
			}
		}
		if len(unvis) == 0 {
			break // WP step 5 (case of no more reachable nodes)
		}
		// WP step 6: new current is node with smallest tentative distance
		current = heap.Pop(&unvis).(*node)
	}
	return
}

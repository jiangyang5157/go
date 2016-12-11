package study

import (
	"math"
	"container/heap"
)

// O(|V|^2) Dijkstra's algorithm (Greedy algorithm) is an algorithm for finding the shortest paths between nodes in a graph


type vertex interface{}

type edge struct {
	v1, v2 vertex
	dist   int
}

type node struct {
	v         vertex     // vertex name
	neighbors []neighbor // edges from this vertex
	tent      int        // tentative distance
	prev      *node      // previous node in shortest path back to start
	done      bool       // true when tent and prev represent shortest path
	rx        int        // heap.Remove index
}

type neighbor struct {
	host *node
	dist int
}

type path struct {
	path   []vertex
	length int
}


// linkGraph constructs a linked representation of the input graph,
// with additional fields needed by the shortest path algorithm.
//
// Return value allNodes will contain all nodes found in the input graph,
// even ones not reachable from the start node.
// Return values startNode, endNode will be nil if the specified start or
// end node names are not found in the graph.
func linkGraph(graph []edge, directed bool,
start, end string) (allNodes []*node, startNode, endNode *node) {

	all := make(map[interface{}]*node)
	// one pass over graph to collect nodes and link neighbors
	for _, e := range graph {
		n1 := all[e.v1]
		n2 := all[e.v2]
		// add previously unseen nodes
		if n1 == nil {
			n1 = &node{v: e.v1}
			all[e.v1] = n1
		}
		if n2 == nil {
			n2 = &node{v: e.v2}
			all[e.v2] = n2
		}
		// link neighbors
		n1.neighbors = append(n1.neighbors, neighbor{n2, e.dist})
		if !directed {
			n2.neighbors = append(n2.neighbors, neighbor{n1, e.dist})
		}
	}
	allNodes = make([]*node, len(all))
	var n int
	for _, nd := range all {
		allNodes[n] = nd
		n++
	}
	return allNodes, all[start], all[end]
}

// dijkstra is a heap-enhanced version of Dijkstra's shortest path algorithm.
//
// If endNode is specified, only a single path is returned.
// If endNode is nil, paths to all nodes are returned.
//
// Note input allNodes is needed to efficiently accomplish WP steps 1 and 2.
// This initialization could be done in linkGraph, but is done here to more
// closely follow the WP algorithm.
func dijkstra(allNodes []*node, startNode, endNode *node) (pl []path) {
	// WP steps 1 and 2.
	for _, nd := range allNodes {
		nd.tent = math.MaxInt32
		nd.done = false
		nd.prev = nil
		nd.rx = -1
	}
	current := startNode
	current.tent = 0
	var unvis ndList

	for {
		// WP step 3: update tentative distances to neighbors
		for _, nb := range current.neighbors {
			if nd := nb.host; !nd.done {
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
			pl = append(pl, path{p, distance}) // pl is return value
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

// ndList implements container/heap
type ndList []*node

func (n ndList) Len() int {
	return len(n)
}
func (n ndList) Less(i, j int) bool {
	return n[i].tent < n[j].tent
}
func (n ndList) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
	n[i].rx = i
	n[j].rx = j
}
func (n *ndList) Push(x interface{}) {
	nd := x.(*node)
	nd.rx = len(*n)
	*n = append(*n, nd)
}
func (n *ndList) Pop() interface{} {
	s := *n
	last := len(s) - 1
	r := s[last]
	*n = s[:last]
	r.rx = -1
	return r
}
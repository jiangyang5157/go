package study

import (
	"testing"
	"fmt"
)

func Test_(t *testing.T) {
	edges := []edge{
		{"a", "b", 7},
		{"a", "c", 9},
		{"a", "f", 14},
		{"b", "c", 10},
		{"b", "d", 15},
		{"c", "d", 11},
		{"c", "f", 2},
		{"d", "e", 6},
		{"e", "f", 9},
	}
	start := "a"
	end := "e"
	directed := false
	nodes, startNode, endNode := linkGraph(edges, directed, start, end)

	if directed {
		fmt.Print("Directed")
	} else {
		fmt.Print("Undirected")
	}
	fmt.Printf(" graph with %d nodes, %d edges\n", len(nodes), len(edges))
	if startNode == nil {
		fmt.Printf("Start node %q not found in graph\n", start)
		return
	}
	if endNode == nil {
		fmt.Printf("End node %q not found in graph\n", end)
		return
	}

	paths := dijkstra(nodes, startNode, endNode)
	fmt.Println("Shortest path(s):")
	for _, p := range paths {
		fmt.Println(p.path, "length", p.length)
	}
}

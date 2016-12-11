package study

import (
	"testing"
	"fmt"
)

func Test_dijkstra(t *testing.T) {
	g := &graph{}
	g.edges = make(map[int][]*edge)
	g.nodes = make(map[int]struct{})

	g.load("data_dijkstra.txt")
	var shortestPath1 map[int]int = g.dijkstra(1)
	// 2599,2610,2947,2052,2367,2399,2029,2442,2505,3068
	fmt.Printf("%d,%d,%d,%d,%d,%d,%d,%d,%d,%d\n", shortestPath1[7], shortestPath1[37], shortestPath1[59], shortestPath1[82], shortestPath1[99], shortestPath1[115], shortestPath1[133], shortestPath1[165], shortestPath1[188], shortestPath1[197])

	g.load("data_dijkstra_test.txt")
	var shortestPath2 map[int]int = g.dijkstra(0)
	// 23,22,4,31,36
	fmt.Printf("%d,%d,%d,%d,%d\n", shortestPath2[7], shortestPath2[1], shortestPath2[2], shortestPath2[3], shortestPath2[4])
}

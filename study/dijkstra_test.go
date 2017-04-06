package study

import (
	"fmt"
	"testing"
)

func Test_dijkstra(t *testing.T) {
	g := &graph{}
	g.edges = make(map[int][]*edge)
	g.nodes = make(map[int]struct{})

	g.load("data_dijkstra.txt")
	var distToTheSource1 map[int]int = g.dijkstra(1)
	// 2599,2610,2947,2052,2367,2399,2029,2442,2505,3068
	fmt.Printf("%d,%d,%d,%d,%d,%d,%d,%d,%d,%d\n", distToTheSource1[7], distToTheSource1[37], distToTheSource1[59], distToTheSource1[82], distToTheSource1[99], distToTheSource1[115], distToTheSource1[133], distToTheSource1[165], distToTheSource1[188], distToTheSource1[197])

	g.load("data_dijkstra_test.txt")
	var distToTheSource2 map[int]int = g.dijkstra(0)
	// 23,22,4,31,36
	fmt.Printf("%d,%d,%d,%d,%d\n", distToTheSource2[7], distToTheSource2[1], distToTheSource2[2], distToTheSource2[3], distToTheSource2[4])
}

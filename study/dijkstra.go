package study
// O(|V|^2) Dijkstra's algorithm (Greedy algorithm) is an algorithm for finding the shortest paths between nodes in a graph

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type graph struct {
	edges map[int][]*edge
	nodes map[int]struct{}
}

type edge struct {
	head   int
	length int
}

func (g *graph) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		edges := strings.Split(strings.TrimSpace(scanner.Text()), "\t")

		// Convert tail vertex to number
		tail, err := strconv.Atoi(edges[0])
		if err != nil {
			log.Fatal(err)
		}
		g.nodes[tail] = struct{}{}

		for i := 1; i < len(edges); i++ {
			data := strings.Split(edges[i], ",")

			// Convert adjacent vertex to number
			head, err := strconv.Atoi(data[0])
			if err != nil {
				log.Fatal(err)
			}

			// Convert length to number
			length, err := strconv.Atoi(data[1])
			if err != nil {
				log.Fatal(err)
			}

			g.nodes[head] = struct{}{}
			e := edge{head: head, length: length}
			g.edges[tail] = append(g.edges[tail], &e)
		}
	}

	return scanner.Err()
}

// Return the map to track distances from source vertex
func (g *graph) dijkstra(source int) map[int]int {
	if _, ok := g.nodes[source]; !ok {
		return nil // source doesn't exist
	}

	const MAX_DIST int = int(1 << 15) - 1 // 32767

	dist := make(map[int]int)
	for index := range g.nodes {
		if index == source {
			dist[source] = 0 // Distance from source to source is zero
		} else {
			dist[index] = MAX_DIST // Initalize distances to maximum
		}
	}

	var tmpIndex int
	for {
		if len(g.nodes) == 0 {
			break
		}
		// Find vertex with minimum distance
		min := MAX_DIST
		for index := range g.nodes {
			if dist[index] < min {
				min = dist[index]
				tmpIndex = index
			}
		}

		// Remove minimum vertex
		delete(g.nodes, tmpIndex)

		// Calculate minimum edge distance
		for _, edge := range g.edges[tmpIndex] {
			if dist[tmpIndex] + edge.length < dist[edge.head] {
				dist[edge.head] = dist[tmpIndex] + edge.length
			}
		}
	}

	return dist
}
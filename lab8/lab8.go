package main

// Graph represented as adjacency matrix
type Graph struct {
	Capacity [][]int // capacity[i][j] = capacity from i to j
	N        int     // number of nodes
}

// Exercise 1: Ford-Fulkerson algorithm
func (g *Graph) maxFlow(source, sink int) int {
	// TODO: Implement Ford-Fulkerson
	// Return value of maximum flow
	return 0
}

// Exercise 2: Find min-cut
func (g *Graph) minCut(source int) []int {
	// TODO: Find nodes reachable from source in residual graph
	// Return list of nodes in min-cut (set A)
	return []int{}
}

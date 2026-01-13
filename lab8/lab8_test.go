package main

import (
	"testing"
)

func createTestGraph1() *Graph {
	// Graph from example
	g := &Graph{N: 4}
	g.Capacity = make([][]int, 4)
	for i := range g.Capacity {
		g.Capacity[i] = make([]int, 4)
	}

	// s=0, a=1, b=2, t=3
	g.Capacity[0][1] = 10 // s->a
	g.Capacity[0][2] = 5  // s->b
	g.Capacity[1][2] = 15 // a->b
	g.Capacity[1][3] = 5  // a->t
	g.Capacity[2][3] = 10 // b->t

	return g
}

func TestMaxFlow(t *testing.T) {
	g := createTestGraph1()
	flow := g.maxFlow(0, 3) // s=0, t=3

	expected := 15
	if flow != expected {
		t.Errorf("Max flow = %d, expected %d", flow, expected)
	}
}

func TestMinCut(t *testing.T) {
	g := createTestGraph1()
	g.maxFlow(0, 3)
	cut := g.minCut(0)

	// In min-cut, source and possibly some other nodes are reachable
	// For this graph, min-cut nodes should be {0, 1} or similar
	if len(cut) == 0 {
		t.Error("Min-cut should contain at least source node")
	}

	// Source should always be in cut
	foundSource := false
	for _, node := range cut {
		if node == 0 {
			foundSource = true
		}
	}
	if !foundSource {
		t.Error("Source node not in min-cut")
	}
}

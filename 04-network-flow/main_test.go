package topic4

import (
	"testing"
)

func TestMaxFlow(t *testing.T) {
	tests := []struct {
		name         string
		vertices     int
		edges        [][3]int // u, v, capacity
		source, sink int
		expectedFlow int
	}{
		{
			name:     "simple graph",
			vertices: 4,
			edges: [][3]int{
				{0, 1, 3},
				{0, 2, 2},
				{1, 2, 5},
				{1, 3, 2},
				{2, 3, 3},
			},
			source:       0,
			sink:         3,
			expectedFlow: 5,
		},
		{
			name:     "graph from slides",
			vertices: 6,
			edges: [][3]int{
				{0, 1, 2},
				{0, 2, 3},
				{0, 3, 3},
				{1, 2, 1},
				{1, 4, 1},
				{2, 3, 2},
				{2, 4, 1},
				{3, 4, 2},
				{3, 5, 3},
				{4, 5, 2},
			},
			source:       0,
			sink:         5,
			expectedFlow: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewFlowGraph(tt.vertices)
			for _, e := range tt.edges {
				g.AddEdge(e[0], e[1], e[2])
			}

			flow := g.MaxFlow(tt.source, tt.sink)
			if flow != tt.expectedFlow {
				t.Errorf("Expected flow %d, got %d", tt.expectedFlow, flow)
			}

			// Verify flow conservation (except source and sink)
			for v := 0; v < tt.vertices; v++ {
				if v == tt.source || v == tt.sink {
					continue
				}
				inFlow, outFlow := 0, 0
				for u := 0; u < tt.vertices; u++ {
					inFlow += g.Flow[u][v]
					outFlow += g.Flow[v][u]
				}
				if inFlow != outFlow {
					t.Errorf("Flow not conserved at vertex %d: in=%d, out=%d", v, inFlow, outFlow)
				}
			}
		})
	}
}

func TestBipartiteMatching(t *testing.T) {
	tests := []struct {
		name     string
		left     int
		right    int
		edges    [][2]int // left, right
		expected int      // maximum matching size
	}{
		{
			name:  "simple bipartite",
			left:  3,
			right: 3,
			edges: [][2]int{
				{0, 0},
				{0, 1},
				{1, 0},
				{1, 2},
				{2, 1},
			},
			expected: 3, // Perfect matching possible
		},
		{
			name:  "no perfect matching",
			left:  3,
			right: 3,
			edges: [][2]int{
				{0, 0},
				{1, 0},
				{1, 1},
				{2, 1},
			},
			expected: 2, // Only 2 matches possible
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create bipartite graph
			g := BipartiteGraph{
				Left:  tt.left,
				Right: tt.right,
				Edges: make([][]bool, tt.left),
			}
			for i := range g.Edges {
				g.Edges[i] = make([]bool, tt.right)
			}
			for _, e := range tt.edges {
				g.Edges[e[0]][e[1]] = true
			}

			matching := MaxBipartiteMatching(g)
			if len(matching) != tt.expected {
				t.Errorf("Expected %d matches, got %d", tt.expected, len(matching))
			}

			// Verify matching is valid
			leftUsed := make(map[int]bool)
			rightUsed := make(map[int]bool)
			for _, m := range matching {
				l, r := m[0], m[1]
				if leftUsed[l] {
					t.Errorf("Left node %d matched twice", l)
				}
				if rightUsed[r] {
					t.Errorf("Right node %d matched twice", r)
				}
				if !g.Edges[l][r] {
					t.Errorf("Edge (%d,%d) not in original graph", l, r)
				}
				leftUsed[l] = true
				rightUsed[r] = true
			}
		})
	}
}

package topic4

// ========== Ford-Fulkerson (Edmonds-Karp) ==========

// Graph represents a flow network
type FlowGraph struct {
	Vertices int
	Capacity [][]int // capacity[u][v]
	Flow     [][]int // flow[u][v]
}

// NewFlowGraph creates a new flow graph
func NewFlowGraph(vertices int) *FlowGraph {
	cap := make([][]int, vertices)
	flow := make([][]int, vertices)
	for i := range cap {
		cap[i] = make([]int, vertices)
		flow[i] = make([]int, vertices)
	}
	return &FlowGraph{
		Vertices: vertices,
		Capacity: cap,
		Flow:     flow,
	}
}

// AddEdge adds a directed edge with capacity
func (g *FlowGraph) AddEdge(u, v, capacity int) {
	g.Capacity[u][v] = capacity
}

// MaxFlow computes maximum flow from source to sink
func (g *FlowGraph) MaxFlow(source, sink int) int {
	// TODO: Implement Edmonds-Karp (BFS-based Ford-Fulkerson)
	// 1. Initialize flow to 0
	// 2. While augmenting path exists:
	//    - Find path using BFS in residual graph
	//    - Compute bottleneck
	//    - Augment flow
	// 3. Return total flow
	return 0
}

// bfs finds augmenting path in residual graph
func (g *FlowGraph) bfs(source, sink int) ([]int, int) {
	// TODO: Implement BFS to find path
	return nil, 0
}

// ========== Bipartite Matching ==========

// BipartiteGraph represents bipartite graph
type BipartiteGraph struct {
	Left, Right int      // number of nodes in left and right partitions
	Edges       [][]bool // adjacency matrix
}

// MaxBipartiteMatching finds maximum matching
func MaxBipartiteMatching(g BipartiteGraph) [][2]int {
	// TODO: Implement via max flow
	// 1. Create flow network
	// 2. Add source and sink
	// 3. Run max flow
	// 4. Extract matching from flow
	return [][2]int{}
}

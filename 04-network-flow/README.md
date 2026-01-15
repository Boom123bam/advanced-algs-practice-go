# Topic 4: Network Flow - Practice Instructions

## Overview
Network flow models flow through directed graphs with capacities. Max-flow min-cut theorem provides duality between flow value and cut capacity.

## Algorithms to Implement

### 1. Ford-Fulkerson (Edmonds-Karp Implementation)

**Problem:** Find maximum flow from source s to sink t in directed graph with edge capacities.

**Key Concepts:**
- **Flow:** f(u,v) satisfying:
  - Capacity: 0 ≤ f(u,v) ≤ c(u,v)
  - Conservation: Σᵥ f(u,v) = Σᵥ f(v,u) for all u ≠ s,t
- **Residual Graph:** G_f has edges with leftover capacity
  - Forward edge: capacity = c(u,v) - f(u,v)
  - Backward edge: capacity = f(u,v) (for undoing flow)
- **Augmenting Path:** Path from s to t in residual graph

**Detailed Steps (Edmonds-Karp - BFS version):**
1. **Initialize:**
   - flow[u][v] = 0 for all edges
   - max_flow = 0

2. **While augmenting path exists:**
   - Find shortest augmenting path using BFS in residual graph
   - Compute bottleneck = min residual capacity along path
   - Augment flow along path:
     - For forward edges: flow[u][v] += bottleneck
     - For backward edges: flow[v][u] -= bottleneck
   - max_flow += bottleneck

3. **Return** max_flow

**Residual Capacity Calculation:**
```
residual_capacity(u,v) = 
  c(u,v) - f(u,v)    if (u,v) is original edge and f(u,v) < c(u,v)
  f(v,u)             if (v,u) is original edge and f(v,u) > 0
  0                  otherwise
```

**Time Complexity:** O(VE²) for Edmonds-Karp

**Finding Min Cut (from Max-Flow):**
1. After max flow computed, run BFS from s in residual graph
2. Min cut = (A,B) where:
   - A = vertices reachable from s in residual graph
   - B = remaining vertices
3. Cut capacity = sum of capacities of edges from A to B

### 2. Bipartite Matching via Max Flow

**Problem:** Find maximum matching in bipartite graph G=(X∪Y, E).

**Reduction to Max Flow:**
1. **Construct Flow Network:**
   - Add source s with edges to all vertices in X (capacity 1)
   - Add sink t with edges from all vertices in Y (capacity 1)
   - Keep original edges X→Y with capacity 1
   
2. **Run Max Flow:** Maximum flow = maximum matching size
   
3. **Extract Matching:**
   - For each edge (x,y) where x∈X, y∈Y
   - If flow[x][y] = 1, then (x,y) is in matching

**Example from Slides:**
- Left: {x₁, x₂, x₃, x₄}
- Right: {y₁, y₂, y₃, y₄}
- Edges: various connections
- Max matching size = 4 (perfect matching)

**Time Complexity:** O(VE) or O(√V E) with specialized algorithms

### 3. Additional Flow Applications

**Baseball Elimination (from slides):**
1. **Construct Flow Network:**
   - Source → game nodes (capacity = games remaining)
   - Game nodes → team nodes (capacity = ∞)
   - Team nodes → sink (capacity = w_i + r_i - w_x, where x is team being tested)
   
2. **Check:** If max flow = total remaining games, team not eliminated

**Multiple Sources/Sinks:**
- Add super-source connected to all sources
- Add super-sink connected from all sinks

**Circulation with Demands:**
- Each vertex has supply/demand
- Convert to max flow with super source/sink

### Key Theorems and Proofs

**Max-Flow Min-Cut Theorem:**
- Value of max flow = capacity of min cut
- Three equivalent conditions:
  1. f is maximum flow
  2. No augmenting paths in residual graph
  3. |f| = capacity of some cut (A,B)

**Proof Technique (from slides):**
1. If no augmenting path, define cut (A,B) where A = vertices reachable from s
2. Show all edges from A to B are saturated
3. Show all edges from B to A have zero flow
4. Therefore |f| = capacity(A,B)

### Implementation Details

**Data Structures:**
1. **Adjacency Matrix:** Simple but O(V²) space
2. **Adjacency List:** More efficient for sparse graphs
3. **Residual Graph:** Can be computed on-the-fly

**Edge Representation:**
```
type Edge struct {
    to     int
    rev    int    // index of reverse edge in adjacency list
    cap    int    // capacity
    flow   int    // current flow
}
```

**BFS for Augmenting Path:**
- Need to store parent pointers for path reconstruction
- Store bottleneck value for each vertex

### Common Pitfalls

1. **Integer vs. Real Capacities:** Ford-Fulkerson may not terminate with irrational capacities
2. **Multiple Edges:** Need to handle parallel edges
3. **Backward Edges:** Crucial for correctness
4. **Stopping Condition:** When BFS can't reach t

### Testing Strategies

1. **Small Hand-calculated Examples:** Verify step-by-step
2. **Flow Conservation:** Check Σᵥ f(u,v) = 0 for all u ≠ s,t
3. **Capacity Constraints:** Check 0 ≤ f(u,v) ≤ c(u,v)
4. **Duality:** Verify |f| ≤ capacity of any cut
5. **Special Cases:** Single edge, bipartite graphs, etc.

### Performance Optimization

1. **Capacity Scaling:** Start with large Δ, reduce gradually
2. **Dinic's Algorithm:** O(V²E) or O(√V E) for bipartite matching
3. **Push-Relabel:** O(V³) or O(V²√E) with gap heuristic
4. **BFS vs. DFS:** BFS finds shortest augmenting paths (Edmonds-Karp)
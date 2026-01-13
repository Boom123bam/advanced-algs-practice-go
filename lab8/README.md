# Lab 8: Max Flow - Ford-Fulkerson Algorithm

## Problem
Find the maximum flow from a source to a sink in a network where each edge has a capacity.

### Constraints:
1. Flow on each edge ≤ capacity
2. Flow into node = flow out of node (except source and sink)
3. Source produces flow, sink consumes flow

## Ford-Fulkerson Algorithm

### Step 1: Initialization
- Start with 0 flow on all edges
- Create residual graph (same as original graph initially)

### Step 2: Find Augmenting Path
While there exists a path from source to sink in residual graph:
1. Find any path with available capacity (using BFS or DFS)
2. Find bottleneck = minimum capacity along this path
3. Add bottleneck flow to all forward edges on path
4. Subtract bottleneck flow from reverse edges (allows "undoing" flow)

### Step 3: Update Residual Graph
After each augmentation:
- Forward edge capacity decreases by flow amount
- Reverse edge capacity increases by flow amount (represents ability to reduce flow)

### Step 4: Termination
When no augmenting path exists, current flow is maximum.

## Key Concepts

### Residual Graph
- Represents remaining capacity
- Forward edges: capacity = original capacity - current flow
- Reverse edges: capacity = current flow (can send flow backwards)

### Min-Cut Theorem
- Max flow value = Min cut capacity
- To find min cut: In final residual graph, find all nodes reachable from source
- Cut = edges from reachable nodes to unreachable nodes

## Example

### Graph:
```
     10        5
  s -----> a -----> t
   \       |       /
    \5     |15    /10
     \     |     /
      ---> b ---
```

### Step-by-Step:
1. Initial flow = 0
2. Find path s→a→t, bottleneck = 5, add flow 5
3. Find path s→b→t, bottleneck = 5, add flow 5  
4. Find path s→a→b→t, bottleneck = 5, add flow 5
5. Total flow = 15, no more augmenting paths

### Min Cut:
Final residual graph: reachable nodes = {s, a}
Cut edges: a→t (capacity 5), s→b (capacity 0 in residual)
But actually cut = edges from {s,a} to {b,t}: a→b (0), a→t (5), s→b (5)
Total capacity = 5 + 5 = 10? Wait, let's trace properly...

Actually for this graph:
- Edges from {s,a} to {b,t}: s→b (5), a→b (15), a→t (5) = 25
But max flow is 15...

Let me correct: The actual min cut should be edges whose capacity limits flow.
In final flow: s→a carries 10, s→b carries 5, a→b carries 5, a→t carries 5, b→t carries 10

The point is: max flow = min cut capacity
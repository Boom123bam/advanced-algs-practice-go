# Topic 4: Network Flow - Practice Instructions

## Overview
Implement network flow algorithms and applications.

## Algorithms to Implement

### 1. Ford-Fulkerson (Edmonds-Karp)
**Steps:**
1. Create residual graph with forward/backward edges
2. While augmenting path exists (BFS):
   - Find path with available capacity
   - Augment flow by bottleneck capacity
   - Update residual graph
3. Return max flow

### 2. Bipartite Matching via Max Flow
**Steps:**
1. Create flow network: source → left nodes → right nodes → sink
2. All capacities = 1
3. Run max flow
4. Maximum flow = maximum matching
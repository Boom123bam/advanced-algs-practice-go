# Topic 2: Greedy Algorithms - Practice Instructions

## Overview
Greedy algorithms make locally optimal choices at each step with the hope of finding a global optimum. They're simple but require proof of correctness.

## Algorithms to Implement

### 1. Interval Scheduling

**Problem:** Given n jobs with start and finish times, select maximum number of compatible jobs.

**Detailed Steps:**
1. **Sorting:** Sort jobs by earliest finish time (ascending)
2. **Greedy Selection:**
   - Initialize: selected = [], last_finish = -∞
   - For each job in sorted order:
     - If job.start ≥ last_finish:
       - Add job to selected
       - last_finish = job.finish
3. **Return** selected jobs

**Proof of Optimality (from slides):**
1. Let S* be any optimal solution
2. Show greedy solution S is at least as good by exchange argument
3. If first job in S* differs from S, replace it with S's first job
   (which finishes no later)
4. Continue this exchange to transform S* into S without decreasing size

**Time Complexity:** O(n log n) for sorting

**Variations:**
- Weighted version (requires DP)
- With resource constraints

### 2. Interval Partitioning (Classroom Scheduling)

**Problem:** Schedule all lectures into minimum number of classrooms.

**Detailed Steps:**
1. **Sorting:** Sort lectures by start time (ascending)
2. **Greedy Assignment:**
   - Track when each classroom becomes free
   - For each lecture:
     - If any classroom is free at lecture.start, assign to it
     - Otherwise, allocate new classroom
3. **Implementation:** Use min-heap to track classroom finish times

**Key Insight:**
- Number of classrooms needed ≥ depth (maximum overlapping intervals)
- Greedy algorithm achieves exactly depth

**Algorithm from Slides:**
```
d = 0  // number of classrooms
for j = 1 to n:
  if lecture j compatible with some classroom k:
    schedule in classroom k
  else:
    allocate new classroom d+1
    d = d+1
```

**Time Complexity:** O(n log n)

### 3. Minimizing Lateness

**Problem:** Schedule jobs with durations and deadlines to minimize maximum lateness.

**Detailed Steps:**
1. **Sorting:** Sort jobs by earliest deadline (not shortest processing time!)
2. **Schedule in Order:**
   - current_time = 0
   - max_lateness = 0
   - For each job in sorted order:
     - finish_time = current_time + job.duration
     - lateness = max(0, finish_time - job.deadline)
     - max_lateness = max(max_lateness, lateness)
     - current_time = finish_time
3. **Return** schedule and max_lateness

**Why Earliest Deadline? Counterexamples:**
- Shortest processing time first: fails with (duration=10, deadline=10) vs (duration=1, deadline=100)
- Smallest slack first: also fails

**Proof Technique (from slides):**
1. Show there exists optimal schedule with no idle time
2. Show there exists optimal schedule with no inversions
   (job i scheduled before j but deadline_i > deadline_j)
3. Greedy schedule has no inversions
4. Swapping inverted jobs doesn't increase max lateness

**Time Complexity:** O(n log n)

### 4. Dijkstra's Algorithm

**Problem:** Find shortest paths from source to all vertices in weighted graph.

**Detailed Steps:**
1. **Initialization:**
   - dist[v] = ∞ for all v
   - dist[source] = 0
   - Priority queue Q contains all vertices
   
2. **Main Loop:**
   - While Q not empty:
     - u = extract-min(Q)  // vertex with smallest dist
     - For each neighbor v of u:
       - If dist[u] + weight(u,v) < dist[v]:
         - dist[v] = dist[u] + weight(u,v)
         - decrease-key(Q, v, dist[v])  // update priority

**Key Data Structure:** Min-heap priority queue

**Proof of Correctness:**
- Maintain invariant: for vertices in S (explored), dist[] is correct shortest path
- When adding new vertex v, its distance is correct because:
  - Any other path to v must leave S at some edge (x,y)
  - But dist[x] + weight(x,y) ≥ dist[v] by choice of v

**Complexity:**
- With array: O(V²)
- With binary heap: O((V+E) log V)
- With Fibonacci heap: O(E + V log V)

**Important:** Doesn't work with negative weights!

### General Greedy Algorithm Template
```
Sort items according to some criterion
Initialize solution = empty
For each item in sorted order:
  If item compatible with current solution:
    Add item to solution
Return solution
```

**When to use Greedy:**
- Problem has optimal substructure
- Greedy choice property (local optimal → global optimal)
- Need efficient O(n log n) solution
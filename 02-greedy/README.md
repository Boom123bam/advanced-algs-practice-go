# Topic 2: Greedy Algorithms - Practice Instructions

## Overview
Implement greedy algorithms that make locally optimal choices.

## Algorithms to Implement

### 1. Interval Scheduling
**Steps:**
1. Sort intervals by earliest finish time
2. Greedily select intervals that don't overlap with previously selected
3. Return maximum number of compatible intervals

### 2. Interval Partitioning
**Steps:**
1. Sort intervals by earliest start time
2. Assign each interval to a compatible classroom (track classroom finish times)
3. Minimize number of classrooms used

### 3. Minimizing Lateness
**Steps:**
1. Sort jobs by earliest deadline
2. Schedule jobs in that order
3. Track lateness (finish time - deadline, if positive)

### 4. Dijkstra's Algorithm
**Steps:**
1. Maintain set of explored nodes with known shortest distance
2. Use priority queue to repeatedly select node with minimum distance
3. Update distances to neighbors
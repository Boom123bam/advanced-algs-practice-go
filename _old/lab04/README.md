# Lab 4: Interval Scheduling

## Key Concepts:
- **Interval Scheduling**: Select maximum number of non-overlapping intervals
- **Interval Partitioning**: Minimum number of "rooms" needed for all intervals
- **Greedy Approach**: Make locally optimal choices at each step

## Exercises:

### Exercise 1: Interval Scheduling (Earliest Finish Time)
**Goal**: Find maximum number of non-overlapping intervals  
**Rough instructions**:
1. Sort intervals by finish time (earliest finish first)
2. Start with empty schedule
3. For each interval in sorted order:
   - If it doesn't overlap with last scheduled interval, add it
   - Otherwise, skip it

**Example**:
Intervals: [(1,3), (2,4), (3,5), (5,7)]
Sorted: [(1,3), (2,4), (3,5), (5,7)]
Schedule: (1,3) → (3,5) → (5,7) = 3 intervals

### Exercise 2: Interval Scheduling (Latest Start Time)
**Goal**: Same as Exercise 1, but build from end  
**Rough instructions**:
1. Sort intervals by start time (latest start first)
2. Start from end time
3. For each interval in sorted order:
   - If it doesn't overlap with previously scheduled interval, add it
   - Build schedule backwards

### Exercise 3: Interval Partitioning
**Goal**: Find minimum number of "rooms" needed  
**Rough instructions**:
1. Sort intervals by start time
2. Keep track of when each room becomes available
3. For each interval:
   - Check if any room is free (finish time ≤ interval start time)
   - If yes, assign to that room
   - If no, create new room
   
**Example**:
Intervals: [(1,4), (2,5), (3,6), (5,7)]
Rooms needed: 3
- Room 1: (1,4), (5,7)
- Room 2: (2,5)
- Room 3: (3,6)

### Exercise 4: Test Different Orderings
Test these sorting orders for Interval Partitioning:
1. Earliest finish time
2. Latest start time  
3. Latest finish time
4. Shortest interval first
5. Longest interval first

Which give optimal solutions?
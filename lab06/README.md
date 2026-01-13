# Lab 6: Dynamic Programming - Knapsack

## Key Concept: Dynamic Programming
Dynamic Programming solves problems by:
1. Breaking them into overlapping subproblems
2. Storing solutions to subproblems in a table
3. Building up the solution from smaller to larger problems

## Exercise: 0-1 Knapsack with Dynamic Programming

### Goal
Find the optimal solution for the 0-1 Knapsack problem (not just a greedy approximation).

### Rough Instructions

#### Step 1: Understand the Problem
You have:
- `n` items, each with a value and weight
- A knapsack with maximum capacity `W`
- Need to maximize total value without exceeding weight capacity

#### Step 2: Create the DP Table
Create a table `dp[i][w]` where:
- `i` = considering first `i` items (0 to n)
- `w` = current weight capacity (0 to W)
- Value = maximum value achievable with first `i` items and capacity `w`

#### Step 3: Base Cases
- `dp[0][w] = 0` (no items, so value is 0)
- `dp[i][0] = 0` (no capacity, so value is 0)

#### Step 4: Recurrence Relation
For each `i` from 1 to n, and each `w` from 1 to W:
```
if weight[i-1] > w:
    dp[i][w] = dp[i-1][w]  # Can't take item i
else:
    dp[i][w] = max(
        dp[i-1][w],  # Don't take item i
        value[i-1] + dp[i-1][w - weight[i-1]]  # Take item i
    )
```

#### Step 5: Fill the Table
Fill row by row, from top-left to bottom-right.

#### Step 6: Get the Answer
The answer is `dp[n][W]` (bottom-right corner of table).

#### Step 7: Traceback (Optional)
To find which items were selected:
1. Start at `dp[n][W]`
2. Move upwards through table:
   - If `dp[i][w] != dp[i-1][w]`, item `i` was taken
   - Subtract its weight and continue
   - Otherwise, move to `dp[i-1][w]`

### Example
Items: A(value=60, weight=10), B(value=100, weight=20), C(value=120, weight=30)
Capacity = 50

Table (simplified):
```
        Capacity: 0  10  20  30  40  50
No items:        0   0   0   0   0   0
Item A only:     0  60  60  60  60  60
Items A,B:       0  60 100 160 160 160  
Items A,B,C:     0  60 100 160 180 220
```
Optimal value = 220 (items B and C)
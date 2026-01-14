# Topic 3: Dynamic Programming - Practice Instructions

## Overview
Implement dynamic programming solutions with memoization/bottom-up approaches.

## Algorithms to Implement

### 1. Weighted Interval Scheduling
**Steps:**
1. Sort intervals by finish time
2. Compute p(j) = last compatible interval before j
3. DP recurrence: OPT(j) = max(v_j + OPT(p(j)), OPT(j-1))
4. Reconstruct solution

### 2. Knapsack Problem (0-1)
**Steps:**
1. DP[i][w] = max value using first i items with weight limit w
2. Recurrence: DP[i][w] = max(DP[i-1][w], v_i + DP[i-1][w-w_i])
3. Reconstruct items

### 3. Sequence Alignment (Edit Distance)
**Steps:**
1. DP[i][j] = min cost to align first i chars of X and j chars of Y
2. Recurrence with mismatch cost α and gap cost δ
3. Trace back for alignment
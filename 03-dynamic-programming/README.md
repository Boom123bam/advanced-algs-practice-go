# Topic 3: Dynamic Programming - Practice Instructions

## Overview
Dynamic programming solves problems by breaking them into overlapping subproblems, solving each once, and storing solutions (memoization). Useful for optimization problems.

## Algorithms to Implement

### 1. Weighted Interval Scheduling

**Problem:** Jobs have start, finish, and weight/value. Find maximum weight set of non-overlapping jobs.

**Detailed Steps:**
1. **Preprocessing:**
   - Sort jobs by finish time: f₁ ≤ f₂ ≤ ... ≤ fₙ
   - Compute p(j) for each job j:
     - p(j) = largest index i < j such that job i is compatible with j
     - i.e., fᵢ ≤ sⱼ
     - Compute using binary search: O(n log n)

2. **DP Recurrence (from slides):**
   ```
   OPT(j) = max {
     vⱼ + OPT(p(j)),  // include job j
     OPT(j-1)          // exclude job j
   }
   ```
   Base case: OPT(0) = 0

3. **Implementation Approaches:**
   - **Memoization (top-down):**
     ```
     M[0] = 0
     for j=1 to n: M[j] = empty
     
     function OPT(j):
       if M[j] empty:
         M[j] = max(vⱼ + OPT(p(j)), OPT(j-1))
       return M[j]
     ```
   - **Iterative (bottom-up):**
     ```
     M[0] = 0
     for j=1 to n:
       M[j] = max(vⱼ + M[p(j)], M[j-1])
     ```

4. **Reconstruct Solution:**
   - Start from j = n
   - If vⱼ + M[p(j)] > M[j-1], include job j, move to p(j)
   - Else, move to j-1

**Time Complexity:** O(n log n) with sorting and binary search

### 2. Knapsack Problem (0-1)

**Problem:** Items have weight and value. Select items with total weight ≤ capacity to maximize total value.

**Detailed Steps:**
1. **DP Definition:**
   - Let OPT(i, w) = max value using first i items with weight limit w
   
2. **Recurrence (from slides):**
   ```
   OPT(i, w) = 
     0                                    if i = 0
     OPT(i-1, w)                          if wᵢ > w
     max {
       OPT(i-1, w),                       // don't take i
       vᵢ + OPT(i-1, w - wᵢ)             // take i
     }                                    otherwise
   ```

3. **Implementation:**
   - Create 2D array M[n+1][W+1]
   - Initialize M[0][w] = 0 for all w
   - Fill row by row:
     ```
     for i = 1 to n:
       for w = 0 to W:
         if wᵢ > w:
           M[i][w] = M[i-1][w]
         else:
           M[i][w] = max(M[i-1][w], vᵢ + M[i-1][w-wᵢ])
     ```

4. **Reconstruct Solution:**
   - Start at i=n, w=W
   - While i > 0:
     - If M[i][w] ≠ M[i-1][w], item i is taken
       - Add i to solution
       - w = w - wᵢ
     - i = i-1

**Time Complexity:** O(nW) — pseudo-polynomial (exponential in input size)

**Space Optimization:** Can use 1D array since we only need previous row

### 3. Sequence Alignment (Edit Distance)

**Problem:** Align two strings with minimum cost (gaps and mismatches).

**Detailed Parameters (from slides):**
- δ = gap penalty (typically 1 or 2)
- α[a,b] = mismatch cost (0 if a=b, typically 1 if a≠b)

**Detailed Steps:**
1. **DP Definition:**
   - Let OPT(i,j) = min cost to align first i chars of X and j chars of Y
   
2. **Recurrence:**
   ```
   OPT(i,j) = min {
     α[Xᵢ, Yⱼ] + OPT(i-1,j-1),  // match/mismatch
     δ + OPT(i-1,j),            // gap in Y
     δ + OPT(i,j-1)             // gap in X
   }
   ```
   Base cases: OPT(i,0) = i·δ, OPT(0,j) = j·δ

3. **Implementation:**
   - Create (m+1)×(n+1) table
   - Fill base cases
   - Fill row by row, column by column
   
4. **Reconstruct Alignment:**
   - Start at (m,n)
   - Trace back choosing minimum cost option
   - Build aligned strings backwards

**Example from Slides:**
Align "ocurrance" with "occurrence"
Best alignment has cost 6 (with δ=2, α=1 for mismatch)

**Time Complexity:** O(mn)

**Variations:**
- Local alignment (Smith-Waterman)
- Affine gap penalties

### DP General Framework

1. **Identify Subproblems:**
   - What are the smaller instances?
   - Usually prefix/suffix or subset problems
   
2. **Define Recurrence:**
   - Express solution in terms of smaller subproblems
   - Consider all possible choices at current step
   
3. **Choose Implementation Strategy:**
   - Memoization (top-down) vs. tabulation (bottom-up)
   - Consider space optimization
   
4. **Reconstruct Solution:**
   - Store parent pointers or trace back from final state
   - May need additional pass through DP table

### Common DP Patterns

1. **Prefix/Suffix:** LCS, edit distance
2. **Interval:** matrix chain multiplication
3. **Tree:** independent set in trees
4. **Subset:** knapsack, partition
5. **State Machine:** sequence alignment

### Tips for Implementation
- Start with brute force recursion
- Add memoization
- Convert to iterative if needed
- Watch for integer overflow in counts
- Use appropriate data structures for reconstruction
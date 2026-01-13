# Lab 7: Sequence Alignment (Needleman-Wunsch Algorithm)

## Goal
Align two strings with minimum "cost" by inserting gaps.

### Costs:
- **Match**: 0 cost (characters are the same)
- **Mismatch**: penalty (e.g., 1)
- **Gap**: penalty (e.g., 1)

## Rough Algorithm

### Step 1: Create DP Table
Create table `dp[i][j]` where:
- `i` = first `i` characters of string X
- `j` = first `j` characters of string Y
- Value = minimum cost to align these prefixes

### Step 2: Base Cases
- `dp[0][j] = j * gapCost` (align empty string with Y using gaps)
- `dp[i][0] = i * gapCost` (align X with empty string using gaps)

### Step 3: Recurrence Relation
For each `i` from 1 to len(X), and `j` from 1 to len(Y):
```
matchCost = 0 if X[i-1] == Y[j-1] else mismatchCost
dp[i][j] = min(
    dp[i-1][j-1] + matchCost,     # align X[i] with Y[j]
    dp[i-1][j] + gapCost,         # align X[i] with gap
    dp[i][j-1] + gapCost          # align gap with Y[j]
)
```

### Step 4: Fill the Table
Fill from top-left to bottom-right.

### Step 5: Traceback to Get Alignment
Start from `dp[len(X)][len(Y)]` and work backwards:
- If came from diagonal: align X[i] with Y[j]
- If came from above: align X[i] with gap
- If came from left: align gap with Y[j]

### Example: Align "CAT" with "CAR"
Costs: mismatch=1, gap=1

Table:
```
    ""  C  A  R
""  0   1  2  3
C   1   0  1  2
A   2   1  0  1
T   3   2  1  1  ← Final cost = 1
```

Traceback gives alignment:
```
C A T
C A R
Cost: 0 0 1 = 1
```

### Visualizing the Process
```
String X: CAT
String Y: CAR

Three choices at each step:
1. Match/mismatch: C with C (match, cost 0)
2. Gap in Y: C with - (gap, cost 1)  
3. Gap in X: - with C (gap, cost 1)

Choose the minimum cost path!
```
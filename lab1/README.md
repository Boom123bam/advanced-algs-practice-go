# Lab 1: Divide and Conquer

## How Divide and Conquer Works:
1. **Divide**: Break the problem into smaller subproblems (usually halves)
2. **Conquer**: Solve the subproblems recursively
3. **Combine**: Merge the solutions to solve the original problem

## Exercises:

### Exercise 1: Sum Array
**Goal**: Calculate sum of all numbers in an array  
**Rough instructions**: 
- If array is empty, return 0
- If array has 1 element, return that element  
- Otherwise: split array into two halves, sum each half recursively, add results

### Exercise 2: Plus Minus  
**Goal**: Compute v1 - v2 + v3 - v4 + ...  
**Rough instructions**:
- Base case: empty array → 0, single element → that element
- For array [a,b,c,d]: compute (a-b) + (c-d) recursively
- Split array in half, compute each half, combine carefully with alternating signs

### Exercise 3: Find Maximum
**Goal**: Find largest number in array  
**Rough instructions**:
- Base case: single element → return it
- Split array in half, find max in each half, return larger of the two

### Exercise 4: Find Max and Min
**Goal**: Find both largest and smallest numbers  
**Rough instructions**:
- Base case: single element → return (element, element)
- Split array in half, get (max1, min1) and (max2, min2) from each half
- Combine: overall max = max(max1, max2), overall min = min(min1, min2)

### Exercise 5: Count Value
**Goal**: Count how many times a specific value appears  
**Rough instructions**:
- Base case: empty array → return 0
- Single element: if element equals value → return 1, else → 0  
- Split array in half, count in each half, add results

### Exercise 6: Maximum Subarray Sum (Harder)
**Goal**: Find maximum sum of any contiguous subarray  
**Rough instructions**:
- Split array into left and right halves
- Three possibilities for max sum:
  1. Entirely in left half → solve recursively
  2. Entirely in right half → solve recursively  
  3. Crosses middle → compute max sum ending at middle + max sum starting at middle
- Return maximum of these three

**Hint for crossing sum**:
- Left side: start from middle, go left, find maximum sum
- Right side: start from middle+1, go right, find maximum sum
- Crossing sum = left sum + right sum
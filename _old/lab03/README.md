# Lab 3: Further Divide and Conquer

## Exercises:

### Exercise 1: MergeSort
**Goal**: Sort array using MergeSort algorithm  
**Rough instructions**:
1. **Base case**: If array has 0 or 1 element, it's already sorted
2. **Divide**: Split array into two equal halves
3. **Conquer**: Recursively sort each half using MergeSort
4. **Combine**: Merge the two sorted halves into one sorted array

**Merge step (combining two sorted arrays)**:
- Use two pointers, one for each array
- Compare elements at pointers, take smaller one
- Move pointer of the array you took from
- Continue until all elements are merged

**Example merging [2,5,8] and [1,3,6]**:
- Compare 2 and 1 → take 1 (from right)
- Compare 2 and 3 → take 2 (from left)
- Compare 5 and 3 → take 3 (from right)
- Compare 5 and 6 → take 5 (from left)
- Compare 8 and 6 → take 6 (from right)
- Take remaining 8 → result: [1,2,3,5,6,8]

### Exercise 2: Search in Sorted 2D Array
**Goal**: Check if target exists in sorted 2D matrix  
**Matrix properties**: Each row sorted left→right, each column sorted top→bottom

**Rough instructions - Strategy 1 (Divide and Conquer)**:
1. Start at bottom-left corner (or top-right)
2. Compare with target:
   - If equal → found it!
   - If current > target: target can't be in this column (all elements below are larger) → move left
   - If current < target: target can't be in this row (all elements to left are smaller) → move down
3. Continue until found or out of bounds

**Rough instructions - Strategy 2 (True Divide and Conquer)**:
1. Pick middle element of matrix
2. Compare with target:
   - If equal → found it!
   - If current > target: target can only be in top-left quadrant
   - If current < target: target can only be in bottom-right quadrant
3. Recursively search appropriate quadrant(s)

**Example searching for 15 in matrix**:
Starting at bottom-left (31):
- 31 > 15 → move left to 29
- 29 > 15 → move left to 11
- 11 < 15 → move down to 26
- 26 > 15 → move left to 12
- 12 < 15 → move down to 15 → found!

**Complexity**: O(m+n) for first strategy, O(log(mn)) for divide and conquer approach
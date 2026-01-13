# Lab 2: Quicksort and k-th Smallest

## Key Concept: Pivot-based Division
These algorithms use a "pivot" element to divide arrays into smaller/larger parts.

## Exercises:

### Exercise 1: Quicksort
**Goal**: Sort array using Quicksort algorithm  
**Rough instructions**:
1. Choose a pivot element (simplest: use first or last element)
2. Partition array into:
   - Left: elements < pivot
   - Middle: pivot element(s)
   - Right: elements > pivot
3. Recursively sort left and right parts
4. Combine: left + middle + right

**Example for [5,2,8,1,4]**:
- Pivot = 5
- Left = [2,1,4] (elements < 5)
- Right = [8] (elements > 5)
- Sort left → [1,2,4], sort right → [8]
- Result: [1,2,4] + [5] + [8] = [1,2,4,5,8]

### Exercise 2: Find k-th Smallest Element
**Goal**: Find k-th smallest element without fully sorting  
**Rough instructions**:
1. Choose a pivot (like in Quicksort)
2. Partition array into left (< pivot), middle (= pivot), right (> pivot)
3. Count elements in left part (let's call it `leftCount`)
4. Decide where to search:
   - If k <= leftCount: k-th smallest is in left part → search there
   - If leftCount < k <= leftCount + len(middle): k-th smallest is the pivot
   - Otherwise: k-th smallest is in right part → search there with adjusted k
   
**Example**: Find 3rd smallest in [7,10,4,3,20,15]
- Pivot = 7
- Left = [4,3] (<7), Middle = [7], Right = [10,20,15] (>7)
- leftCount = 2, middleCount = 1
- k=3: Since 2 < 3 <= 3 (2+1), 3rd smallest is pivot = 7

### Exercise 3: Find Median
**Goal**: Find middle value(s) of array  
**Rough instructions**:
- If array length is odd: median = middle element (k = n/2 + 1)
- If array length is even: median = average of two middle elements
  - Need both k = n/2 and k = n/2 + 1 smallest elements
  - You can call kthSmallest twice, or modify it to return both

**Examples**:
- [1,3,5,7,9] (odd length): median = 5 (3rd smallest of 5 elements)
- [1,2,3,4] (even length): median = (2+3)/2 = 2.5
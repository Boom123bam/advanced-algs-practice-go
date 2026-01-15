# Topic 1: Divide and Conquer - Practice Instructions

## Overview
Divide and conquer algorithms break problems into smaller subproblems, solve them recursively, and combine solutions. This paradigm often yields O(n log n) solutions.

## Algorithms to Implement

### 1. Quicksort

**Detailed Steps:**
1. **Base Case:** If array has 0 or 1 elements, it's already sorted
2. **Divide:**
   - Choose pivot element (use last element as in slides)
   - Partition array into two subarrays:
     - Elements ≤ pivot
     - Elements > pivot
   - Pivot element ends up in its final sorted position
3. **Conquer:**
   - Recursively apply Quicksort to left subarray (elements ≤ pivot)
   - Recursively apply Quicksort to right subarray (elements > pivot)
4. **Combine:** No explicit combine step - array is sorted in place

**Implementation Details from Slides:**
- Use `rearrange(S, a, b)` function that:
  - Takes array S, indices a (start), b (end)
  - Sets P = S[b] (pivot)
  - Uses two pointers: l = a, r = b - 1
  - While l ≤ r:
    - Move l right while S[l] ≤ P
    - Move r left while S[r] > P
    - If l < r, swap S[l] and S[r]
  - Finally swap S[l] and S[b] (place pivot)
  - Return l (pivot's final position)

**Time Complexity:**
- Best/Average: O(n log n)
- Worst (already sorted): O(n²)

**Tips:**
- Implement both recursive and iterative versions
- Test with arrays containing duplicates

### 2. Mergesort

**Detailed Steps:**
1. **Base Case:** If array has 0 or 1 elements, return
2. **Divide:**
   - Find middle index: mid = (left + right) / 2
   - Divide array into two halves: left = arr[left..mid], right = arr[mid+1..right]
3. **Conquer:**
   - Recursively sort left half
   - Recursively sort right half
4. **Combine (Merge):**
   - Create temporary array
   - Use three pointers: i (left), j (right), k (result)
   - While both halves have elements:
     - Take smaller of arr[i] and arr[j], copy to result
     - Increment appropriate pointers
   - Copy remaining elements from either half

**Implementation Details from Slides:**
- Merge function is the key: O(n) time
- Stable sort (preserves order of equal elements)
- Requires O(n) extra space

**Time Complexity:**
- Always O(n log n)
- Space: O(n)

**Tips:**
- Implement bottom-up (iterative) version too
- Use for linked lists (no extra space needed)

### 3. Counting Inversions

**Detailed Steps:**
1. Modify Mergesort to count inversions during merge step
2. When merging two sorted halves:
   - If element from right half is smaller than element from left half
   - It forms inversions with ALL remaining elements in left half
   - Count = count + (mid - i + 1) where i is current left index
3. Return total inversion count

**Example from Slides:**
Input: [1, 5, 4, 8, 10, 2, 6, 9, 12, 11, 3, 7]
Total inversions: 22

**Time Complexity:** O(n log n)

**Real-world Application:** Measuring similarity between rankings

**Tips:**
- Return both sorted array and count from recursive calls
- Use long integers for count (can be up to n(n-1)/2)

### 4. Closest Pair of Points (2D)

**Detailed Steps:**
1. **Preprocessing:**
   - Sort points by x-coordinate → Px
   - Sort points by y-coordinate → Py (keep for strip checking)
   
2. **Base Case:** If ≤ 3 points, use brute force O(n²) comparison
   
3. **Divide:**
   - Find vertical line L that splits points into two halves
   - Left half: points with x ≤ L
   - Right half: points with x > L
   
4. **Conquer:**
   - Recursively find closest pair in left half → δ₁
   - Recursively find closest pair in right half → δ₂
   - δ = min(δ₁, δ₂)
   
5. **Combine (Check strip):**
   - Consider points within δ distance of line L
   - Sort these points by y-coordinate (use Py)
   - For each point, compare with next 11 points (or 7 in some implementations)
   - This is based on geometric packing argument from slides

**Key Insight from Slides:**
- Only need to compare points within δ strip
- Each point needs to compare with at most 11 neighbors in y-sorted order
- This ensures O(n log n) overall complexity

**Time Complexity:** O(n log n)

**Tips:**
- Precompute sorted arrays to avoid sorting at each recursion
- Use Master Theorem: T(n) = 2T(n/2) + O(n) → O(n log n)

### Master Theorem Practice

For each algorithm, identify parameters for Master Theorem:
- a = number of subproblems
- b = factor by which problem size reduces
- f(n) = cost of dividing and combining

**Example (Mergesort):**
- T(n) = 2T(n/2) + O(n)
- a = 2, b = 2, f(n) = O(n)
- Case 2: O(n log n)
# Topic 1: Divide and Conquer - Practice Instructions

## Overview
Implement the following divide-and-conquer algorithms as described in the lectures.

## Algorithms to Implement

### 1. Quicksort
**Steps:**
1. Choose a pivot (use the last element)
2. Partition the array into elements ≤ pivot and > pivot
3. Recursively sort the two partitions
4. Combine (pivot is already in correct position)

**Key Points:**
- Implement in-place partitioning
- Time: O(n log n) average, O(n²) worst-case
- Use the algorithm from slides with `l` and `r` pointers

### 2. Mergesort
**Steps:**
1. Divide array into two halves
2. Recursively sort each half
3. Merge the two sorted halves using temporary array

**Key Points:**
- Stable sort (preserves order of equal elements)
- Time: O(n log n) always
- Need O(n) extra space

### 3. Counting Inversions
**Steps:**
1. Modify mergesort to count inversions during merge
2. When element from right half is smaller, it forms inversions with all remaining elements in left half
3. Return total inversion count

**Key Points:**
- Use mergesort framework
- Count while merging

### 4. Closest Pair of Points (2D)
**Steps:**
1. Sort points by x-coordinate (preprocess)
2. Recursively find closest distance in left and right halves
3. Check for closer pairs within δ strip of midline
4. Only compare points within 11 positions in y-sorted order

**Key Points:**
- Sort by x and y coordinates initially
- Use Master Theorem to analyze complexity: O(n log n)
## Lab 1: Divide and Conquer

### Exercise 1: Sum of Array
```go
// Approach 1: Sum between indices
func sum(arr []int, left, right int) int {
    // Base case: when left >= right
    // Divide: split into two halves
    // Combine: add the two sums
}

// Approach 2: Split the array
func listSum(arr []int) int {
    // Base case: array of length 0 or 1
    // Divide: split array into two halves
    // Recursively solve each half
    // Combine: add the two results
}
```

### Exercise 2: Alternating Sum
```go
func plusMinus(arr []int) int {
    // Compute: v1 - v2 + v3 - v4 + ...
    // Base case: array of length 1 or 2
    // Divide: split array into two halves
    // Combine: consider how to combine alternating sums
}
```

### Exercise 3: Maximum Element
```go
func findMax(arr []int) int {
    // Base case: array of length 1
    // Divide: split array into two halves
    // Recursively find max in each half
    // Combine: return the larger of the two
}
```

### Exercise 4: Maximum and Minimum
```go
func findMinMax(arr []int) (int, int) {
    // Return both min and max
    // Base case: array of length 1 or 2
    // Divide: split array into two halves
    // Recursively find min/max in each half
    // Combine: return overall min and max
}
```

### Exercise 5: Count Occurrences
```go
func countValue(arr []int, target int) int {
    // Count how many times target appears
    // Base case: array of length 1
    // Divide: split array into two halves
    // Recursively count in each half
    // Combine: add the two counts
}
```

### Exercise 6: Maximum Subarray Sum
```go
func maxSubarraySum(arr []int) int {
    // Find the contiguous subarray with maximum sum
    // Base case: array of length 1
    // Divide: split array into two halves
    // Three possibilities:
    // 1. Max sum in left half
    // 2. Max sum in right half
    // 3. Max sum crossing the midpoint
    // Return the largest of these three
}
```

## Lab 2: Quicksort and Selection

### Exercise 1: Quicksort
```go
func quickSort(arr []int) []int {
    // Base case: array of length 0 or 1
    // Choose a pivot
    // Divide: partition into elements < pivot and > pivot
    // Recursively sort left and right partitions
    // Combine: left + pivot + right
}
```

### Exercise 2: k-th Smallest Element
```go
func kthSmallest(arr []int, k int) int {
    // Find the k-th smallest element (1-indexed)
    // Base case: array of length 1
    // Choose a pivot
    // Partition into left (< pivot) and right (> pivot)
    // Compare sizes with k:
    // - If left size >= k: search in left
    // - If left size == k-1: return pivot
    // - Otherwise: search in right with adjusted k
}
```

### Exercise 3: Median
```go
func findMedian(arr []int) float64 {
    // Find median using kthSmallest
    // If odd length: return middle element
    // If even length: return average of two middle elements
}
```

## Lab 3: Further Divide and Conquer

### Exercise 1: Merge Sort
```go
func mergeSort(arr []int) []int {
    // Base case: array of length 0 or 1
    // Divide: split array into two equal halves
    // Recursively sort each half
    // Combine: merge two sorted halves
}

func merge(left, right []int) []int {
    // Merge two sorted arrays into one sorted array
}
```

### Exercise 2: Search in Sorted 2D Array
```go
func searchInSortedTable(matrix [][]int, target int) bool {
    // Search in a 2D array where:
    // - Each row is sorted left to right
    // - Each column is sorted top to bottom
    // Base case: empty matrix or 1x1 matrix
    // Choose a strategic position (like top-right or bottom-left)
    // Compare target with that position
    // Eliminate impossible rows/columns
    // Recursively search remaining submatrix
}
```

## Simplified Implementation Hints:

1. **Base cases**: Usually when array length is 0, 1, or 2
2. **Divide**: Split at midpoint: `mid := len(arr)/2`
3. **Conquer**: Recursive calls: `solve(arr[:mid])` and `solve(arr[mid:])`
4. **Combine**: Simple operation like `max`, `min`, `+`, etc.

Each exercise focuses on one core concept while keeping the Go implementations clean and minimal. Start with the simplest cases and build up to more complex ones.
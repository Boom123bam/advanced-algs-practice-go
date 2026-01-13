# Lab 2: Quicksort and k-th Smallest

This lab focuses on two related Divide and Conquer algorithms that use a pivot element.

## Exercises:

### Exercise 1: Quicksort
Implement `quickSort(arr []int) []int` to sort an array using the Quicksort algorithm.

### Exercise 2: Find k-th Smallest Element
Implement `kthSmallest(arr []int, k int) int` to find the k-th smallest element in an unsorted array.
- k is 1-indexed (1 = smallest, 2 = second smallest, etc.)
- Use a pivot-based Divide and Conquer approach similar to Quicksort
- Average time complexity should be O(n)

### Exercise 3: Find Median
Implement `findMedian(arr []int) float64` to find the median value of an array.
- If array length is odd: return middle element
- If array length is even: return average of two middle elements
- Use your kthSmallest implementation as inspiration
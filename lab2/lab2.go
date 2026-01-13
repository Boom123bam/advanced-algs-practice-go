package main

// Exercise 1: Quicksort
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)/2]
	left := []int{}
	right := []int{}
	for i, num := range arr {
		if i == len(arr)/2 {
			continue
		}
		if num < pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}
	return append(append(quickSort(left), pivot), quickSort(right)...)
}

// Exercise 2: Find k-th Smallest Element
func kthSmallest(arr []int, k int) int {
	// k is 1-indexed (1 = smallest element)
	if len(arr) == 1 {
		return arr[0]
	}
	mid := len(arr) / 2
	pivot := arr[mid]
	left := []int{}
	right := []int{}
	for i, num := range arr {
		if i == mid {
			continue
		}
		if num < pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}
	if k <= len(left) {
		return kthSmallest(left, k)
	}
	if k == len(left)+1 {
		return pivot
	}
	return kthSmallest(right, k-len(left)-1)
}

// Exercise 3: Find Median
func findMedian(arr []int) float64 {
	// Hint: Use kthSmallest as inspiration

	if len(arr)%2 == 1 {
		return float64(kthSmallest(arr, len(arr)/2+1))
	}
	return float64(kthSmallest(arr, len(arr)/2)+kthSmallest(arr, len(arr)/2+1)) / 2

}

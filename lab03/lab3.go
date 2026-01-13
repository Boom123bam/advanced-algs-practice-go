package main

// Exercise 1: MergeSort
func mergeSort(arr []int) []int {
	// You'll need a merge helper function
	if len(arr) <= 1 {
		return arr
	}
	left := arr[:len(arr)/2]
	right := arr[len(arr)/2:]
	return merge(mergeSort(left), mergeSort(right))
}

// Helper function for mergeSort
func merge(left, right []int) []int {
	result := []int{}
	for i, j := 0, 0; i < len(left) || j < len(right); {
		if i == len(left) {
			result = append(result, right[j])
			j++
			continue
		}
		if j == len(right) {
			result = append(result, left[i])
			i++
			continue
		}
		l, r := left[i], right[j]
		if l < r {
			result = append(result, l)
			i++
		} else {
			result = append(result, r)
			j++
		}

	}
	return result
}

// Exercise 2: Search in Sorted 2D Array
func searchInSorted2D(matrix [][]int, target int) bool {
	// Matrix properties: rows sorted left->right, columns sorted top->bottom
	// btm left
	x := 0
	y := len(matrix) - 1
	for x < len(matrix[0]) && y >= 0 {
		current := matrix[y][x]
		if current == target {
			return true
		}
		if current > target {
			y -= 1
		} else {
			x += 1
		}
	}
	return false
}

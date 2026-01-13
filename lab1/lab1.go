package main

// Exercise 1: Sum Array
func sumArray(arr []int) int {
	switch len(arr) {
	case 0:
		return 0
	case 1:
		return arr[0]
	default:
		return sumArray(arr[:len(arr)/2]) + sumArray(arr[len(arr)/2:])
	}
}

// Exercise 2: Plus Minus
func plusMinus(arr []int) int {
	switch len(arr) {
	case 0:
		return 0
	case 1:
		return arr[0]
	default:
		mid := len(arr) / 2
		if mid%2 == 0 {
			return plusMinus(arr[:mid]) + plusMinus(arr[mid:])
		}
		return plusMinus(arr[:mid]) - plusMinus(arr[mid:])

	}
}

// Exercise 3: Find Maximum
func findMax(arr []int) int {
	switch len(arr) {
	case 0:
		return 0
	case 1:
		return arr[0]
	default:
		a := findMax(arr[:len(arr)/2])
		b := findMax(arr[len(arr)/2:])
		if a > b {
			return a
		}
		return b
	}
}

// Exercise 4: Find Max and Min
func findMaxMin(arr []int) (int, int) {
	switch len(arr) {
	case 0:
		return 0, 0
	case 1:
		return arr[0], arr[0]
	default:
		max1, min1 := findMaxMin(arr[:len(arr)/2])
		max2, min2 := findMaxMin(arr[len(arr)/2:])
		max := max1
		min := min2
		if max2 > max1 {
			max = max2
		}
		if min2 > min1 {
			min = min2
		}
		return max, min
	}
}

// Exercise 5: Count Value
func countValue(arr []int, value int) int {
	switch len(arr) {
	case 0:
		return 0
	case 1:
		if arr[0] == value {
			return 1
		}
		return 0
	default:
		return countValue(arr[:len(arr)/2], value) + countValue(arr[len(arr)/2:], value)
	}
}

// Exercise 6: Maximum Subarray Sum
func maxSubarraySum(arr []int) int {
	// TODO: Implement using Divide and Conquer
	// Hint: You might want a helper function that returns
	// max sum in left, max sum in right, max sum crossing
	switch len(arr) {
	case 0:
		return 0
	case 1:
		return arr[0]
	default:
		leftMax := maxSubarraySum(arr[:len(arr)/2])
		rightMax := maxSubarraySum(arr[len(arr)/2:])
		midLeftSum := arr[len(arr)/2-1]
		midRightSum := arr[len(arr)/2]
		midLeftMax := midLeftSum
		midRightMax := midRightSum

		for i := len(arr)/2 - 2; i >= 0; i-- {
			midLeftSum += arr[i]
			if midLeftSum > midLeftMax {
				midLeftMax = midLeftSum
			}
		}
		for i := len(arr)/2 + 1; i < len(arr); i++ {
			midRightSum += arr[i]
			if midRightSum > midRightMax {
				midRightMax = midRightSum
			}
		}
		midMax := midRightMax + midLeftMax
		max := leftMax
		if rightMax > max {
			max = rightMax
		}
		if midMax > max {
			max = midMax
		}
		return max
	}
}

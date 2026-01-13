package main

import (
	"testing"
)

func TestSumArray(t *testing.T) {
	tests := []struct {
		arr      []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{}, 0},
		{[]int{5}, 5},
		{[]int{-1, 2, -3, 4, -5}, -3},
	}

	for _, test := range tests {
		result := sumArray(test.arr)
		if result != test.expected {
			t.Errorf("sumArray(%v) = %d, expected %d", test.arr, result, test.expected)
		}
	}
}

func TestPlusMinus(t *testing.T) {
	tests := []struct {
		arr      []int
		expected int
	}{
		{[]int{6, 5, 4, 3, 2, 1}, 3},
		{[]int{1, 2, 3}, 2},
		{[]int{10}, 10},
		{[]int{5, 3}, 2},
	}

	for _, test := range tests {
		result := plusMinus(test.arr)
		if result != test.expected {
			t.Errorf("plusMinus(%v) = %d, expected %d", test.arr, result, test.expected)
		}
	}
}

func TestFindMax(t *testing.T) {
	tests := []struct {
		arr      []int
		expected int
	}{
		{[]int{3, 7, 2, 9, 1}, 9},
		{[]int{-5, -2, -10}, -2},
		{[]int{42}, 42},
		{[]int{5, 5, 5, 5}, 5},
	}

	for _, test := range tests {
		result := findMax(test.arr)
		if result != test.expected {
			t.Errorf("findMax(%v) = %d, expected %d", test.arr, result, test.expected)
		}
	}
}

func TestFindMaxMin(t *testing.T) {
	tests := []struct {
		arr         []int
		expectedMax int
		expectedMin int
	}{
		{[]int{3, 7, 2, 9, 1}, 9, 1},
		{[]int{-5, -2, -10}, -2, -10},
		{[]int{42}, 42, 42},
		{[]int{5, 5, 5, 5}, 5, 5},
	}

	for _, test := range tests {
		max, min := findMaxMin(test.arr)
		if max != test.expectedMax || min != test.expectedMin {
			t.Errorf("findMaxMin(%v) = (%d, %d), expected (%d, %d)",
				test.arr, max, min, test.expectedMax, test.expectedMin)
		}
	}
}

func TestCountValue(t *testing.T) {
	tests := []struct {
		arr      []int
		value    int
		expected int
	}{
		{[]int{1, 2, 3, 2, 4, 2, 5}, 2, 3},
		{[]int{1, 2, 3, 4, 5}, 6, 0},
		{[]int{}, 5, 0},
		{[]int{7, 7, 7}, 7, 3},
	}

	for _, test := range tests {
		result := countValue(test.arr, test.value)
		if result != test.expected {
			t.Errorf("countValue(%v, %d) = %d, expected %d",
				test.arr, test.value, result, test.expected)
		}
	}
}

func TestMaxSubarraySum(t *testing.T) {
	tests := []struct {
		arr      []int
		expected int
	}{
		{[]int{-1, 2, 3, -4, 5, -6}, 6},
		{[]int{-2, -3, 4, -1, -2, 1, 5, -3}, 7},
		{[]int{-5, -2, -1, -3}, -1},
		{[]int{1, 2, 3, 4}, 10},
	}

	for _, test := range tests {
		result := maxSubarraySum(test.arr)
		if result != test.expected {
			t.Errorf("maxSubarraySum(%v) = %d, expected %d",
				test.arr, result, test.expected)
		}
	}
}

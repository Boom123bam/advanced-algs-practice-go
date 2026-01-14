package main

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{[]int{5, 2, 8, 1, 9}, []int{1, 2, 5, 8, 9}},
		{[]int{}, []int{}},
		{[]int{3}, []int{3}},
		{[]int{9, 7, 5, 3, 1}, []int{1, 3, 5, 7, 9}},
		{[]int{3, 1, 4, 1, 5, 9, 2, 6}, []int{1, 1, 2, 3, 4, 5, 6, 9}},
	}

	for _, test := range tests {
		result := mergeSort(test.arr)
		if len(result) != len(test.expected) {
			t.Errorf("mergeSort(%v) length = %d, expected %d",
				test.arr, len(result), len(test.expected))
			continue
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("mergeSort(%v)[%d] = %d, expected %d",
					test.arr, i, result[i], test.expected[i])
				break
			}
		}
	}
}

func TestSearchInSorted2D(t *testing.T) {
	matrix := [][]int{
		{1, 4, 8, 15, 23},
		{3, 7, 9, 35, 68},
		{9, 12, 15, 36, 79},
		{11, 26, 28, 42, 88},
		{29, 32, 38, 66, 91},
		{31, 34, 43, 67, 94},
	}

	tests := []struct {
		target   int
		expected bool
	}{
		{66, true},
		{1, true},
		{94, true},
		{15, true},
		{100, false},
		{0, false},
		{2, false},
		{43, true},
		{99, false},
	}

	for _, test := range tests {
		result := searchInSorted2D(matrix, test.target)
		if result != test.expected {
			t.Errorf("searchInSorted2D(matrix, %d) = %v, expected %v",
				test.target, result, test.expected)
		}
	}
}

// Test with smaller matrix
func TestSearchInSorted2DSmall(t *testing.T) {
	smallMatrix := [][]int{
		{1, 3},
		{2, 4},
	}

	tests := []struct {
		target   int
		expected bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, true},
		{5, false},
		{0, false},
	}

	for _, test := range tests {
		result := searchInSorted2D(smallMatrix, test.target)
		if result != test.expected {
			t.Errorf("searchInSorted2D(smallMatrix, %d) = %v, expected %v",
				test.target, result, test.expected)
		}
	}
}

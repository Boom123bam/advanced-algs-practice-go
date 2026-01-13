package main

import (
	"math"
	"testing"
)

func TestQuickSort(t *testing.T) {
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
		result := quickSort(test.arr)
		if len(result) != len(test.expected) {
			t.Errorf("quickSort(%v) length = %d, expected %d",
				test.arr, len(result), len(test.expected))
			continue
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("quickSort(%v)[%d] = %d, expected %d",
					test.arr, i, result[i], test.expected[i])
				break
			}
		}
	}
}

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		arr      []int
		k        int
		expected int
	}{
		{[]int{7, 10, 4, 3, 20, 15}, 3, 7},
		{[]int{7, 10, 4, 3, 20, 15}, 1, 3},
		{[]int{7, 10, 4, 3, 20, 15}, 6, 20},
		{[]int{3}, 1, 3},
		{[]int{5, 5, 5, 5, 5}, 3, 5},
	}

	for _, test := range tests {
		result := kthSmallest(test.arr, test.k)
		if result != test.expected {
			t.Errorf("kthSmallest(%v, %d) = %d, expected %d",
				test.arr, test.k, result, test.expected)
		}
	}
}

func TestFindMedian(t *testing.T) {
	tests := []struct {
		arr      []int
		expected float64
	}{
		{[]int{7, 10, 4, 3, 20, 15}, 8.5}, // (7+10)/2
		{[]int{7, 10, 4, 3, 20}, 7.0},     // middle element
		{[]int{1, 2, 3, 4}, 2.5},          // (2+3)/2
		{[]int{5}, 5.0},
		{[]int{2, 4, 6, 8, 10, 12}, 7.0}, // (6+8)/2
	}

	for _, test := range tests {
		result := findMedian(test.arr)
		// Use a small epsilon for floating point comparison
		if math.Abs(result-test.expected) > 0.0001 {
			t.Errorf("findMedian(%v) = %f, expected %f",
				test.arr, result, test.expected)
		}
	}
}

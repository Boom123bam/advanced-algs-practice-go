package topic1

import (
	"math"
	"testing"
)

func TestQuicksort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse sorted",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "random array",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6},
			expected: []int{1, 1, 2, 3, 4, 5, 6, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)

			QSort(arr, 0, len(arr)-1)

			if len(arr) != len(tt.expected) {
				t.Errorf("Expected length %d, got %d", len(tt.expected), len(arr))
				return
			}

			for i := range arr {
				if arr[i] != tt.expected[i] {
					t.Errorf("At index %d: expected %d, got %d", i, tt.expected[i], arr[i])
				}
			}
		})
	}
}

func TestMergesort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "random array",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6},
			expected: []int{1, 1, 2, 3, 4, 5, 6, 9},
		},
		{
			name:     "array with duplicates",
			input:    []int{5, 2, 8, 2, 5, 8, 1},
			expected: []int{1, 2, 2, 5, 5, 8, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeSort(tt.input)

			if len(result) != len(tt.expected) {
				t.Errorf("Expected length %d, got %d", len(tt.expected), len(result))
				return
			}

			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("At index %d: expected %d, got %d", i, tt.expected[i], result[i])
				}
			}
		})
	}
}

func TestCountInversions(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "empty array",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "reverse sorted",
			input:    []int{5, 4, 3, 2, 1},
			expected: 10, // n(n-1)/2 for n=5
		},
		{
			name:     "example from slides",
			input:    []int{1, 5, 4, 8, 10, 2, 6, 9, 12, 11, 3, 7},
			expected: 22,
		},
		{
			name:     "simple case",
			input:    []int{2, 4, 1, 3, 5},
			expected: 3, // (2,1), (4,1), (4,3)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountInversions(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %d inversions, got %d", tt.expected, result)
			}
		})
	}
}

func TestClosestPair(t *testing.T) {
	// Helper function to compare floats with tolerance
	almostEqual := func(a, b float64) bool {
		return math.Abs(a-b) < 1e-9
	}

	tests := []struct {
		name           string
		points         []Point
		expectedDist   float64
		shouldHavePair bool
	}{
		{
			name: "two points",
			points: []Point{
				{0, 0},
				{3, 4},
			},
			expectedDist:   5.0, // √(3² + 4²) = 5
			shouldHavePair: true,
		},
		{
			name: "three points forming right triangle",
			points: []Point{
				{0, 0},
				{3, 0},
				{0, 4},
			},
			expectedDist:   3.0, // Distance between (0,0) and (3,0)
			shouldHavePair: true,
		},
		{
			name: "points from lecture example",
			points: []Point{
				{0, 0},
				{5, 5},
				{1, 2},
				{3, 1},
				{4, 3},
			},
			expectedDist:   math.Sqrt(2), // Distance between (1,2) and (2,3) = √2
			shouldHavePair: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p1, p2, dist := ClosestPair(tt.points)

			if !almostEqual(dist, tt.expectedDist) {
				t.Errorf("Expected distance %.6f, got %.6f", tt.expectedDist, dist)
			}

			if tt.shouldHavePair {
				// Verify the returned points are actually in the input
				found1, found2 := false, false
				for _, p := range tt.points {
					if almostEqual(p.X, p1.X) && almostEqual(p.Y, p1.Y) {
						found1 = true
					}
					if almostEqual(p.X, p2.X) && almostEqual(p.Y, p2.Y) {
						found2 = true
					}
				}
				if !found1 || !found2 {
					t.Error("Returned points not found in input")
				}

				// Verify distance calculation is correct
				calcDist := distance(p1, p2)
				if !almostEqual(calcDist, dist) {
					t.Errorf("Distance mismatch: calculated %.6f, returned %.6f", calcDist, dist)
				}
			}
		})
	}
}

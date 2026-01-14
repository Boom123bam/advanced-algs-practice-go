package topic3

import (
	"testing"
)

func TestWeightedSchedule(t *testing.T) {
	// Example from slides
	intervals := []WeightedInterval{
		{Start: 1, End: 4, Weight: 2},
		{Start: 3, End: 5, Weight: 4},
		{Start: 0, End: 6, Weight: 4},
		{Start: 4, End: 7, Weight: 7},
		{Start: 3, End: 8, Weight: 2},
		{Start: 5, End: 9, Weight: 1},
		{Start: 6, End: 10, Weight: 3},
		{Start: 8, End: 11, Weight: 5},
	}

	t.Run("weighted interval scheduling", func(t *testing.T) {
		selected, totalWeight := WeightedSchedule(intervals)

		// Verify no overlaps
		for i := 0; i < len(selected)-1; i++ {
			if selected[i].End > selected[i+1].Start {
				t.Errorf("Intervals overlap: %v and %v", selected[i], selected[i+1])
			}
		}

		// Check weight is reasonable (should be optimal)
		// Expected from slides: maximum weight is 12
		expectedMinWeight := 12
		if totalWeight < expectedMinWeight {
			t.Errorf("Expected weight at least %d, got %d", expectedMinWeight, totalWeight)
		}
	})
}

func TestKnapsack(t *testing.T) {
	// Example from slides
	items := []Item{
		{Weight: 1, Value: 1},
		{Weight: 2, Value: 6},
		{Weight: 5, Value: 18},
		{Weight: 6, Value: 22},
		{Weight: 7, Value: 28},
	}

	tests := []struct {
		name     string
		capacity int
		expected int
	}{
		{"capacity 11", 11, 40},
		{"capacity 10", 10, 35},
		{"capacity 5", 5, 18},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, totalValue := Knapsack(items, tt.capacity)
			if totalValue != tt.expected {
				t.Errorf("Expected value %d, got %d", tt.expected, totalValue)
			}
		})
	}
}

func TestAlignSequences(t *testing.T) {
	tests := []struct {
		name     string
		x        string
		y        string
		expected int // minimum edit distance
	}{
		{
			name:     "identical strings",
			x:        "ACG",
			y:        "ACG",
			expected: 0, // No cost for matches
		},
		{
			name:     "example from slides",
			x:        "ocurrance",
			y:        "occurrence",
			expected: 6, // With given costs in slides
		},
		{
			name:     "empty string",
			x:        "",
			y:        "ABC",
			expected: 3 * GapCost, // Three gaps
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AlignSequences(tt.x, tt.y)
			if result.Cost != tt.expected {
				t.Errorf("Expected cost %d, got %d", tt.expected, result.Cost)
			}

			// Verify aligned strings have same length
			if len(result.XAligned) != len(result.YAligned) {
				t.Error("Aligned strings have different lengths")
			}

			// Verify alignment is valid (no simultaneous gaps)
			for i := 0; i < len(result.XAligned); i++ {
				if result.XAligned[i] == '-' && result.YAligned[i] == '-' {
					t.Error("Both strings have gap at same position")
				}
			}
		})
	}
}

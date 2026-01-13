package main

import (
	"testing"
)

func TestKnapsackDP(t *testing.T) {
	tests := []struct {
		items    []Item
		capacity int
		expected int
	}{
		{
			[]Item{
				{"A", 60, 10},
				{"B", 100, 20},
				{"C", 120, 30},
			},
			50,
			220, // B + C = 100 + 120
		},
		{
			[]Item{
				{"Diamond", 10, 2},
				{"Gold", 20, 3},
				{"Silver", 15, 4},
			},
			6,
			30, // Diamond + Gold = 10 + 20
		},
	}

	for i, test := range tests {
		result := knapsackDP(test.items, test.capacity)
		if result != test.expected {
			t.Errorf("Test %d: Got value %d, expected %d", i, result, test.expected)
		}
	}
}

func TestKnapsackDPSolution(t *testing.T) {
	items := []Item{
		{"A", 60, 10},
		{"B", 100, 20},
		{"C", 120, 30},
	}
	capacity := 50

	value, selected := knapsackDPSolution(items, capacity)

	if value != 220 {
		t.Errorf("Expected value 220, got %d", value)
	}

	// Check that selected items don't exceed capacity
	totalWeight := 0
	for _, item := range selected {
		totalWeight += item.Weight
	}
	if totalWeight > capacity {
		t.Errorf("Selected items weight %d exceeds capacity %d", totalWeight, capacity)
	}
}

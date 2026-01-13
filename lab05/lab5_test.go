package main

import (
	"testing"
)

func TestKnapsackGreedy(t *testing.T) {
	items := []Item{
		{"A", 60, 10},
		{"B", 100, 20},
		{"C", 120, 30},
	}
	capacity := 50

	// Test different strategies
	strategies := []string{"weight", "value", "density"}

	for _, strategy := range strategies {
		selected, value := knapsackGreedy(items, capacity, strategy)
		totalWeight := 0
		for _, item := range selected {
			totalWeight += item.Weight
		}
		if totalWeight > capacity {
			t.Errorf("Strategy %s: Total weight %d exceeds capacity %d",
				strategy, totalWeight, capacity)
		}
		t.Logf("Strategy %s: Value = %d, Items = %v", strategy, value, selected)
	}
}

func TestFractionalKnapsack(t *testing.T) {
	items := []Item{
		{"A", 60, 10},  // ratio = 6
		{"B", 100, 20}, // ratio = 5
		{"C", 120, 30}, // ratio = 4
	}
	capacity := 50

	selected, value := fractionalKnapsack(items, capacity)
	totalWeight := 0.0
	for _, item := range selected {
		// For fractional, we'd need to track fraction taken
		totalWeight += float64(item.Weight)
	}

	// For fractional knapsack with these items and capacity=50:
	// Take all of A (10kg, value=60)
	// Take all of B (20kg, value=100)
	// Take 20/30 of C (20kg, value=80)
	// Total: 50kg, value=240
	expectedValue := 240.0

	if value != expectedValue {
		t.Errorf("Fractional knapsack: Got value %.2f, expected %.2f", value, expectedValue)
	}
}

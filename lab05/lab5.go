package main

type Item struct {
	Name   string
	Value  int
	Weight int
}

// Exercise 1-6: Different greedy strategies for 0-1 Knapsack
func knapsackGreedy(items []Item, capacity int, strategy string) ([]Item, int) {
	// TODO: Implement different greedy strategies
	// strategy can be: "weight", "value", "density" (value/weight), etc.
	// Return selected items and total value
	return []Item{}, 0
}

// Exercise 10: Fractional Knapsack
func fractionalKnapsack(items []Item, capacity int) ([]Item, float64) {
	// TODO: Implement fractional knapsack
	// Items can be split, take fractions if needed
	// Return items with fractions and total value
	return []Item{}, 0.0
}

// Helper to test if greedy is optimal
func testGreedyVsOptimal(strategy string, numItems int, maxTrials int) bool {
	// TODO: Generate random items, compare greedy solution with optimal (brute force)
	// Return true if greedy always optimal for small cases
	return false
}

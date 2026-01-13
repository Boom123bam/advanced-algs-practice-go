package main

type Item struct {
	Name   string
	Value  int
	Weight int
}

// Exercise 1: Knapsack DP - returns max value
func knapsackDP(items []Item, capacity int) int {
	// TODO: Implement dynamic programming solution
	// Create dp table and fill it using recurrence relation
	return 0
}

// Exercise 2: Knapsack DP with solution traceback
func knapsackDPSolution(items []Item, capacity int) (int, []Item) {
	// TODO: Return max value AND which items were selected
	// Need to trace back through dp table
	return 0, []Item{}
}

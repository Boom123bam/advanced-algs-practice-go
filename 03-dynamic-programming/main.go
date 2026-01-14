package topic3

// ========== Weighted Interval Scheduling ==========

type WeightedInterval struct {
	Start, End, Weight int
}

// WeightedSchedule finds maximum weight set of non-overlapping intervals
func WeightedSchedule(intervals []WeightedInterval) ([]WeightedInterval, int) {
	// TODO: Implement
	// 1. Sort by finish time
	// 2. Compute p(j) using binary search
	// 3. DP with memoization
	// 4. Reconstruct solution
	return []WeightedInterval{}, 0
}

// ========== Knapsack Problem ==========

type Item struct {
	Weight, Value int
}

// Knapsack solves 0-1 knapsack problem
func Knapsack(items []Item, capacity int) ([]Item, int) {
	// TODO: Implement
	// 1. Create DP table
	// 2. Fill using recurrence
	// 3. Reconstruct items
	return []Item{}, 0
}

// ========== Sequence Alignment ==========

// AlignmentResult contains the aligned sequences and cost
type AlignmentResult struct {
	XAligned string
	YAligned string
	Cost     int
}

// constants for alignment costs
const (
	GapCost      = 2
	MatchCost    = 0
	MismatchCost = 1
)

// AlignSequences computes optimal alignment of two strings
func AlignSequences(x, y string) AlignmentResult {
	// TODO: Implement Needleman-Wunsch algorithm
	// 1. Initialize DP table
	// 2. Fill using recurrence
	// 3. Trace back to get alignment
	return AlignmentResult{}
}

// mismatchCost returns cost for aligning chars a and b
func mismatchCost(a, b byte) int {
	if a == b {
		return MatchCost
	}
	return MismatchCost
}

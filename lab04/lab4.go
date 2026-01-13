package main

// Interval represents a time interval with start and end times
type Interval struct {
	Start int
	End   int
}

// Exercise 1: Interval Scheduling (Earliest Finish Time)
func intervalSchedulingEarliestFinish(intervals []Interval) []Interval {
	// TODO: Implement greedy algorithm with earliest finish time first
	// Sort by end time, then select non-overlapping intervals
	return []Interval{}
}

// Exercise 2: Interval Scheduling (Latest Start Time)
func intervalSchedulingLatestStart(intervals []Interval) []Interval {
	// TODO: Implement greedy algorithm with latest start time first
	// Sort by start time descending, build schedule from end
	return []Interval{}
}

// Exercise 3: Interval Partitioning
func intervalPartitioning(intervals []Interval) [][]Interval {
	// TODO: Implement algorithm to find minimum number of rooms
	// Return schedule showing which intervals go in each room
	return [][]Interval{}
}

// Helper: Check if two intervals overlap
func overlaps(a, b Interval) bool {
	// TODO: Return true if intervals a and b overlap
	return a.Start < b.End && b.Start < a.End
}

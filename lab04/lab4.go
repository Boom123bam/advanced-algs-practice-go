package main

import (
	"math"
	"sort"
)

// Interval represents a time interval with start and end times
type Interval struct {
	Start int
	End   int
}

// Exercise 1: Interval Scheduling (Earliest Finish Time)
func intervalSchedulingEarliestFinish(intervals []Interval) []Interval {
	// Sort by end time, then select non-overlapping intervals
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].End < intervals[j].End
	})
	lastEndTime := 0
	result := []Interval{}
	for _, interval := range intervals {
		if lastEndTime <= interval.Start {
			result = append(result, interval)
			lastEndTime = interval.End
		}
	}
	return result
}

// Exercise 2: Interval Scheduling (Latest Start Time)
func intervalSchedulingLatestStart(intervals []Interval) []Interval {
	// Sort by start time descending, build schedule from end
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start > intervals[j].Start
	})
	lastStartTime := math.MaxInt
	result := []Interval{}
	for _, interval := range intervals {
		if lastStartTime >= interval.End {
			result = append([]Interval{interval}, result...)
			lastStartTime = interval.Start
		}
	}
	return result
}

// Exercise 3: Interval Partitioning
func intervalPartitioning(intervals []Interval) [][]Interval {
	// Return schedule showing which intervals go in each room
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	rooms := [][]Interval{}
	for _, interval := range intervals {
		foundRoom := false
		for i, room := range rooms {
			if room[len(room)-1].End <= interval.Start {
				rooms[i] = append(room, interval)
				foundRoom = true
				break
			}
		}
		if !foundRoom {
			rooms = append(rooms, []Interval{interval})
		}
	}
	return rooms
}

// Helper: Check if two intervals overlap
func overlaps(a, b Interval) bool {
	// TODO: Return true if intervals a and b overlap
	return a.Start < b.End && b.Start < a.End
}

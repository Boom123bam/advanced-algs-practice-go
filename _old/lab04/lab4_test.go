package main

import (
	"testing"
)

func TestIntervalSchedulingEarliestFinish(t *testing.T) {
	tests := []struct {
		intervals []Interval
		expected  int // expected number of intervals in optimal schedule
	}{
		{
			[]Interval{{1, 4}, {2, 5}, {7, 9}, {6, 10}, {8, 11}},
			2, // (1,4), (7,9), or similar
		},
		{
			[]Interval{{1, 3}, {2, 4}, {3, 5}, {5, 7}},
			3, // (1,3), (3,5), (5,7)
		},
		{
			[]Interval{{1, 2}, {2, 3}, {3, 4}},
			3, // All can be scheduled
		},
	}

	for i, test := range tests {
		result := intervalSchedulingEarliestFinish(test.intervals)
		if len(result) != test.expected {
			t.Errorf("Test %d: Got %d intervals, expected %d", i, len(result), test.expected)
		}
	}
}

func TestIntervalSchedulingLatestStart(t *testing.T) {
	tests := []struct {
		intervals []Interval
		expected  int // should give same number as earliest finish
	}{
		{
			[]Interval{{1, 4}, {2, 5}, {7, 9}, {6, 10}, {8, 11}},
			2,
		},
		{
			[]Interval{{1, 3}, {2, 4}, {3, 5}, {5, 7}},
			3,
		},
	}

	for i, test := range tests {
		result := intervalSchedulingLatestStart(test.intervals)
		if len(result) != test.expected {
			t.Errorf("Test %d: Got %d intervals, expected %d", i, len(result), test.expected)
		}
	}
}

func TestIntervalPartitioning(t *testing.T) {
	tests := []struct {
		intervals []Interval
		expected  int // minimum number of rooms needed
	}{
		{
			[]Interval{{1, 4}, {2, 5}, {3, 6}, {5, 7}},
			3,
		},
		{
			[]Interval{{9, 10}, {9, 11}, {10, 12}, {11, 13}},
			2,
		},
		{
			[]Interval{{1, 2}, {3, 4}, {5, 6}},
			1,
		},
	}

	for i, test := range tests {
		result := intervalPartitioning(test.intervals)
		if len(result) != test.expected {
			t.Errorf("Test %d: Got %d rooms, expected %d", i, len(result), test.expected)
		}
	}
}

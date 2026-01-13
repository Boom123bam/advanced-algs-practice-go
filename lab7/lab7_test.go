package main

import (
	"testing"
)

func TestAlignSequences(t *testing.T) {
	tests := []struct {
		x, y     string
		mismatch int
		gap      int
		expected int
	}{
		{"CAT", "CAR", 1, 1, 1},         // One mismatch
		{"SUNDAY", "SATURDAY", 1, 1, 3}, // Need to align
		{"A", "A", 1, 1, 0},             // Perfect match
		{"A", "B", 1, 1, 1},             // Mismatch
		{"", "ABC", 1, 1, 3},            // All gaps
	}

	for i, test := range tests {
		cost, _ := alignSequences(test.x, test.y, test.mismatch, test.gap)
		if cost != test.expected {
			t.Errorf("Test %d: Align(%q, %q) cost = %d, expected %d",
				i, test.x, test.y, cost, test.expected)
		}
	}
}

func TestTraceAlignment(t *testing.T) {
	x, y := "CAT", "CAR"
	mismatch, gap := 1, 1

	cost, dp := alignSequences(x, y, mismatch, gap)
	alignX, alignY := traceAlignment(x, y, dp, mismatch, gap)

	if cost != 1 {
		t.Errorf("Expected cost 1, got %d", cost)
	}

	// Check alignment strings have same length
	if len(alignX) != len(alignY) {
		t.Errorf("Alignment strings different lengths: %d vs %d",
			len(alignX), len(alignY))
	}

	t.Logf("Alignment: %s\n         %s", alignX, alignY)
}

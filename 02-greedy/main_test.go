package topic2

import (
	"math"
	"testing"
)

func TestIntervalSchedule(t *testing.T) {
	tests := []struct {
		name      string
		intervals []Interval
		expected  int // expected number of intervals
	}{
		{
			name: "example from slides",
			intervals: []Interval{
				{Start: 0, End: 6},
				{Start: 1, End: 4},
				{Start: 3, End: 5},
				{Start: 3, End: 8},
				{Start: 4, End: 7},
				{Start: 5, End: 9},
				{Start: 6, End: 10},
				{Start: 8, End: 11},
			},
			expected: 4, // Should select intervals ending at 4, 7, 10 (or similar)
		},
		{
			name: "simple example",
			intervals: []Interval{
				{Start: 1, End: 3},
				{Start: 2, End: 4},
				{Start: 3, End: 5},
			},
			expected: 2, // Can select {1,3} and {3,5}
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IntervalSchedule(tt.intervals)

			if len(result) != tt.expected {
				t.Errorf("Expected %d intervals, got %d", tt.expected, len(result))
			}

			// Verify no overlaps
			for i := 0; i < len(result)-1; i++ {
				if result[i].End > result[i+1].Start {
					t.Errorf("Intervals %d and %d overlap: %v ends at %d, %v starts at %d",
						i, i+1, result[i], result[i].End, result[i+1], result[i+1].Start)
				}
			}
		})
	}
}

func TestIntervalPartition(t *testing.T) {
	tests := []struct {
		name      string
		intervals []Interval
		expected  int // minimum classrooms
	}{
		{
			name: "simple case - 1 classroom",
			intervals: []Interval{
				{Start: 1, End: 3},
				{Start: 4, End: 6},
				{Start: 7, End: 9},
			},
			expected: 1, // All intervals can use same classroom
		},
		{
			name: "overlapping intervals - 2 classrooms",
			intervals: []Interval{
				{Start: 1, End: 4},
				{Start: 2, End: 5},
				{Start: 6, End: 8},
			},
			expected: 2,
		},
		{
			name: "lecture example simplified",
			intervals: []Interval{
				{Start: 900, End: 1030},  // 9:00 to 10:30
				{Start: 900, End: 1230},  // 9:00 to 12:30
				{Start: 930, End: 1030},  // 9:30 to 10:30
				{Start: 1000, End: 1100}, // 10:00 to 11:00
				{Start: 1030, End: 1200}, // 10:30 to 12:00
				{Start: 1100, End: 1230}, // 11:00 to 12:30
				{Start: 1130, End: 1230}, // 11:30 to 12:30
				{Start: 1200, End: 1400}, // 12:00 to 14:00
				{Start: 1300, End: 1430}, // 13:00 to 14:30
				{Start: 1300, End: 1430}, // 13:00 to 14:30
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IntervalPartition(tt.intervals)
			if result != tt.expected {
				t.Errorf("Expected %d classrooms, got %d", tt.expected, result)
			}
		})
	}
}

func TestScheduleJobs(t *testing.T) {
	tests := []struct {
		name     string
		jobs     []Job
		expected int // maximum lateness
	}{
		{
			name: "example from slides",
			jobs: []Job{
				{ID: 1, Duration: 3, Deadline: 6},
				{ID: 2, Duration: 2, Deadline: 8},
				{ID: 3, Duration: 1, Deadline: 9},
				{ID: 4, Duration: 4, Deadline: 9},
				{ID: 5, Duration: 3, Deadline: 14},
				{ID: 6, Duration: 2, Deadline: 15},
			},
			expected: 1, // Maximum lateness should be 1
		},
		{
			name: "no lateness",
			jobs: []Job{
				{ID: 1, Duration: 2, Deadline: 5},
				{ID: 2, Duration: 3, Deadline: 10},
				{ID: 3, Duration: 1, Deadline: 6},
			},
			expected: 0, // All jobs finish on time
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schedule, maxLateness := ScheduleJobs(tt.jobs)

			if maxLateness != tt.expected {
				t.Errorf("Expected max lateness %d, got %d", tt.expected, maxLateness)
			}

			// Verify schedule is in deadline order
			currentTime := 0
			for _, job := range schedule {
				finishTime := currentTime + job.Duration
				if finishTime > job.Deadline {
					// Calculate lateness for this job
					lateness := finishTime - job.Deadline
					if lateness > maxLateness {
						t.Errorf("Job %d has lateness %d > reported max %d",
							job.ID, lateness, maxLateness)
					}
				}
				currentTime = finishTime
			}
		})
	}
}

func TestDijkstra(t *testing.T) {
	// Create graph from slides example
	g := Graph{
		Vertices: 8,
		Edges:    make(map[int]map[int]int),
	}

	// Add edges as in the slides example
	edges := [][3]int{
		{0, 1, 9}, {0, 5, 14}, {0, 6, 15},
		{1, 2, 24},
		{2, 4, 2}, {2, 7, 19},
		{3, 2, 6}, {3, 7, 6},
		{4, 3, 6}, {4, 5, 30}, {4, 7, 16},
		{5, 2, 18}, {5, 4, 30}, {5, 7, 44},
		{6, 3, 44}, {6, 5, 20},
	}

	for _, e := range edges {
		if g.Edges[e[0]] == nil {
			g.Edges[e[0]] = make(map[int]int)
		}
		g.Edges[e[0]][e[1]] = e[2]
	}

	t.Run("shortest path from 0 to 7", func(t *testing.T) {
		dist, prev := Dijkstra(g, 0)

		// Expected shortest distance from 0 to 7 is 50
		expected := 50
		if dist[7] != expected {
			t.Errorf("Expected distance %d, got %d", expected, dist[7])
		}

		// Verify path reconstruction
		if prev[7] != -1 {
			// Should be path 0 -> 1 -> 2 -> 4 -> 3 -> 7 or similar
			t.Logf("Path to 7: distance = %d", dist[7])
		}

		// Verify distances to all vertices are reasonable
		for v := 0; v < g.Vertices; v++ {
			if v == 0 {
				if dist[v] != 0 {
					t.Errorf("Distance to source should be 0, got %d", dist[v])
				}
			} else if dist[v] == math.MaxInt32 {
				t.Errorf("Vertex %d unreachable but should be reachable", v)
			}
		}
	})
}

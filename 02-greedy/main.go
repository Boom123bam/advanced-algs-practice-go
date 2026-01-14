package topic2

import (
	"container/heap"
	"math"
)

// ========== Interval Scheduling ==========

type Interval struct {
	Start, End int
}

// IntervalSchedule finds maximum number of non-overlapping intervals
func IntervalSchedule(intervals []Interval) []Interval {
	// TODO: Implement interval scheduling
	// 1. Sort intervals by end time
	// 2. Greedily select intervals
	return []Interval{}
}

// ========== Interval Partitioning ==========

// IntervalPartition finds minimum classrooms needed
func IntervalPartition(intervals []Interval) int {
	// TODO: Implement interval partitioning
	// 1. Sort intervals by start time
	// 2. Track when classrooms become free
	// 3. Assign to first available classroom or create new one
	return 0
}

// ========== Minimizing Lateness ==========

type Job struct {
	ID       int
	Duration int // processing time
	Deadline int
}

// ScheduleJobs schedules to minimize maximum lateness
func ScheduleJobs(jobs []Job) ([]Job, int) {
	// TODO: Implement
	// 1. Sort by deadline
	// 2. Process jobs in order
	// 3. Track current time and lateness
	return []Job{}, 0
}

// ========== Dijkstra's Algorithm ==========

type Graph struct {
	Vertices int
	Edges    map[int]map[int]int // adjacency list with weights
}

// Dijkstra finds shortest paths from source to all vertices
func Dijkstra(g Graph, source int) ([]int, []int) {
	// TODO: Implement Dijkstra's algorithm
	// 1. Initialize distances to infinity
	// 2. Use priority queue
	// 3. Relax edges
	dist := make([]int, g.Vertices)
	prev := make([]int, g.Vertices)
	for i := range dist {
		dist[i] = math.MaxInt32
		prev[i] = -1
	}
	dist[source] = 0

	return dist, prev
}

// Priority queue implementation for Dijkstra
type Item struct {
	vertex int
	dist   int
	index  int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// Update modifies the distance of an item in the queue
func (pq *PriorityQueue) Update(item *Item, dist int) {
	item.dist = dist
	heap.Fix(pq, item.index)
}

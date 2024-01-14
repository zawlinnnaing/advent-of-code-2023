package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	row          int
	col          int
	heatLoss     int
	directionRow int
	directionCol int
	n            int
	index        int
}

func (item *Item) String() string {
	return fmt.Sprintf("CurrentPount: (%d, %d); Direction: (%d, %d); HeatLoss: %d; Straight: %d",
		item.row, item.col, item.directionRow, item.directionCol, item.heatLoss, item.n)
}

type PriorityQueue []*Item

type Grid [][]int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].heatLoss < pq[j].heatLoss
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func NewPriorityQueue() *PriorityQueue {
	pq := PriorityQueue{}
	heap.Init(&pq)
	return &pq
}

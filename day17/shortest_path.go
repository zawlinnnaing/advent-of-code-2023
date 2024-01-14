package main

import (
	"container/heap"
)

var directions = [][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

type SeenItem struct {
	r  int
	c  int
	dr int
	dc int
	n  int
}

func FindShortPathWithPQ(grid Grid) int {
	seen := make(map[SeenItem]bool)
	pq := NewPriorityQueue()
	heap.Push(pq, &Item{
		row:          0,
		col:          0,
		heatLoss:     0,
		directionRow: 0,
		directionCol: 0,
		n:            0,
	})
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		if item.row == len(grid)-1 && item.col == len(grid[0])-1 {
			return item.heatLoss
		}
		seenItem := SeenItem{
			r:  item.row,
			c:  item.col,
			dr: item.directionRow,
			dc: item.directionCol,
			n:  item.n,
		}
		if _, ok := seen[seenItem]; ok {
			continue
		}
		seen[seenItem] = true
		// Going straight
		if item.n < 3 && (item.directionCol != 0 || item.directionRow != 0) {
			nextRow := item.row + item.directionRow
			nextCol := item.col + item.directionCol
			if nextRow >= 0 && nextRow < len(grid) && nextCol >= 0 && nextCol < len(grid[0]) {
				nextItem := Item{
					row:          nextRow,
					col:          nextCol,
					heatLoss:     item.heatLoss + grid[nextRow][nextCol],
					directionRow: item.directionRow,
					directionCol: item.directionCol,
					n:            item.n + 1,
				}
				heap.Push(pq, &nextItem)
			}
		}
		// Turning left or right
		for _, direction := range directions {
			//// going same direction (straight) skips
			//if direction[0] == item.directionRow && direction[1] == item.directionCol {
			//	continue
			//}
			//// Going reverse, not allowed
			//if direction[0] == -item.directionRow && direction[1] == -item.directionCol {
			//	//fmt.Println("direction", direction, item.directionRow, item.directionCol)
			//	continue
			//}
			if (direction[0] != item.directionRow || direction[1] != item.directionCol) && (direction[0] != -item.directionRow || direction[1] != -item.directionCol) {
				nextRow, nextCol := item.row+direction[0], item.col+direction[1]
				if nextRow >= 0 && nextRow < len(grid) && nextCol >= 0 && nextCol < len(grid[0]) {
					nextItem := Item{
						row:          nextRow,
						col:          nextCol,
						heatLoss:     item.heatLoss + grid[nextRow][nextCol],
						directionRow: direction[0],
						directionCol: direction[1],
						n:            1,
					}
					heap.Push(pq, &nextItem)
				}
			}
		}
	}
	return 0
}

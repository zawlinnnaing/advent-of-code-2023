package main

import (
	"container/heap"
)

func RunPart2(grid Grid) int {
	pq := NewPriorityQueue()
	seen := make(map[SeenItem]bool)
	heap.Push(pq, &Item{
		row:          0,
		col:          0,
		heatLoss:     0,
		directionRow: 0,
		directionCol: 0,
		n:            0,
		index:        0,
	})
	numRows := len(grid)
	numCols := len(grid[0])
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		if item.row == numRows-1 && item.col == numCols-1 {
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
		// going straight
		if item.n < 10 && (item.directionRow != 0 || item.directionCol != 0) {
			// Going down
			if item.directionRow == 1 && item.directionCol == 0 {
				supposedEnd := item.row + (10 - item.n)
				if supposedEnd >= numRows-4 && supposedEnd < numRows-1 && item.row >= numRows-4 {
					continue
				}
			}
			// Going right
			if item.directionRow == 0 && item.directionCol == 1 {
				supposedEnd := item.col + (10 - item.n)
				if supposedEnd >= numCols-4 && supposedEnd < numCols-1 && item.col >= numCols-4 {
					continue
				}
			}
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

		// For initial case
		if item.directionRow == 0 && item.directionCol == 0 && item.n == 0 {
			for _, direction := range directions {
				// going same direction (straight) skips
				if direction[0] == item.directionRow && direction[1] == item.directionCol {
					continue
				}
				// Going reverse, not allowed
				if direction[0] == -item.directionRow && direction[1] == -item.directionCol {
					continue
				}
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

		if item.n >= 4 {
			// Turning left or right
			for _, direction := range directions {
				// going same direction (straight) skips
				if direction[0] == item.directionRow && direction[1] == item.directionCol {
					continue
				}
				// Going reverse, not allowed
				if direction[0] == -item.directionRow && direction[1] == -item.directionCol {
					continue
				}
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

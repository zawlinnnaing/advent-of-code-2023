package main

import "sort"

func RunPart2(grid Grid) int {
	numRows := len(grid)
	numCols := len(grid[0])
	totals := []int{}
	for rowIdx := 0; rowIdx < numRows; rowIdx++ {
		for colIdx := 0; colIdx < numCols; colIdx++ {
			startPoint := [2]int{rowIdx, colIdx}
			if rowIdx == 0 {
				totals = append(totals, CalculateTotalEnergizedTiles(grid, startPoint, "bottom"))
			}
			if colIdx == 0 {
				totals = append(totals, CalculateTotalEnergizedTiles(grid, startPoint, "right"))
			}
			if rowIdx == numRows-1 {
				totals = append(totals, CalculateTotalEnergizedTiles(grid, startPoint, "top"))
			}
			if colIdx == numCols-1 {
				totals = append(totals, CalculateTotalEnergizedTiles(grid, startPoint, "left"))
			}
		}
	}
	sort.Ints(totals)
	return totals[len(totals)-1]
}

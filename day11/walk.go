package main

import "math"

func findSteps(source [2]int, target [2]int, emptyRows, emptyCols []int, stepsToAdd int) int {
	rowDiff := int(math.Abs(float64(source[0] - target[0])))
	colDiff := int(math.Abs(float64(source[1] - target[1])))
	additionalSteps := 0
	for _, rowIdx := range emptyRows {
		if rowIdx >= source[0] && rowIdx <= target[0] || rowIdx >= target[0] && rowIdx <= source[0] {
			additionalSteps += stepsToAdd
		}
	}
	for _, colIdx := range emptyCols {
		if colIdx >= source[1] && colIdx <= target[1] || colIdx >= target[1] && colIdx <= source[1] {
			additionalSteps += stepsToAdd
		}
	}
	return rowDiff + colDiff + additionalSteps
}

func findGalaxyCoords(universe Universe) [][2]int {
	galaxyCoords := [][2]int{}
	for rowIdx, row := range universe {
		for colIdx, cell := range row {
			if cell == GALAXY {
				galaxyCoords = append(galaxyCoords, [2]int{rowIdx, colIdx})
			}
		}
	}
	return galaxyCoords
}

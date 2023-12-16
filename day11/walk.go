package main

import "math"

func findSteps(source [2]int, target [2]int) int {
	rowDiff := int(math.Abs(float64(source[0] - target[0])))
	colDiff := int(math.Abs(float64(source[1] - target[1])))
	return rowDiff + colDiff
}

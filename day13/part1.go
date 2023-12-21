package main

import (
	"reflect"
)

func CheckReflection(pattern Pattern, firstPoint, secondPoint int) bool {
	currentPoint, nextPoint := firstPoint, secondPoint
	for currentPoint >= 0 && nextPoint < len(pattern) {
		if !reflect.DeepEqual(pattern[currentPoint], pattern[nextPoint]) {
			return false
		}
		currentPoint -= 1
		nextPoint += 1
	}
	return true
}

func FindReflectionPoint(pattern Pattern) (int, int) {
	rows := len(pattern)
	possiblePoints := [2]int{0, 0}
	for i := 0; i < rows-1; i++ {
		currentRow := pattern[i]
		nextRow := pattern[i+1]
		if reflect.DeepEqual(currentRow, nextRow) && CheckReflection(pattern, i-1, i+2) {
			if i == 0 {
				possiblePoints = [2]int{i, i + 1}
				continue
			}
			return i, i + 1
		}
	}
	return possiblePoints[0], possiblePoints[1]
}

func runPart1(patterns []Pattern) int {
	total := 0
	for _, pattern := range patterns {
		firstPoint, secondPoint := FindReflectionPoint(pattern)
		//Row reflection point
		if firstPoint != secondPoint {
			total += secondPoint * 100
			continue
		}
		//	Col reflection point
		firstPoint, secondPoint = FindReflectionPoint(pattern.Transpose())
		total += secondPoint
	}

	return total
}

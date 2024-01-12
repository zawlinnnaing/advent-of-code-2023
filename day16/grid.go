package main

import (
	"slices"
)

type Grid [][]string

type Path struct {
	point     [2]int
	direction string
}

var directionsMap = map[string][2]int{
	"left":   {0, -1},
	"right":  {0, 1},
	"top":    {-1, 0},
	"bottom": {1, 0},
}

func calculateReflectDirection(mirror string, currentDirection string) string {
	if mirror == "/" {
		switch currentDirection {
		case "left":
			return "bottom"
		case "right":
			return "top"
		case "top":
			return "right"
		case "bottom":
			return "left"
		}
	}
	if mirror == "\\" {
		switch currentDirection {
		case "left":
			return "top"
		case "right":
			return "bottom"
		case "top":
			return "left"
		case "bottom":
			return "right"
		}
	}
	return currentDirection
}

func calculateSplitDirections(mirror string, currentDirection string) []string {
	if mirror == "|" && slices.Contains([]string{"left", "right"}, currentDirection) {
		return []string{"top", "bottom"}
	}
	if mirror == "-" && slices.Contains([]string{"top", "bottom"}, currentDirection) {
		return []string{"left", "right"}
	}

	return []string{currentDirection}
}

func getNextPoint(currentPoint [2]int, direction string) [2]int {
	directionPoint := directionsMap[direction]
	return [2]int{currentPoint[0] + directionPoint[0], currentPoint[1] + directionPoint[1]}
}

func Walk(grid Grid, seen *[][]bool, currentPoint [2]int, direction string, path *[]Path) {
	numRows := len(grid)
	numCols := len(grid[0])
	if currentPoint[0] < 0 || currentPoint[0] >= numRows || currentPoint[1] < 0 || currentPoint[1] >= numCols {
		return
	}
	currentPath := Path{
		point:     currentPoint,
		direction: direction,
	}
	if slices.Contains(*path, currentPath) {
		return
	}
	(*seen)[currentPoint[0]][currentPoint[1]] = true
	currentValue := grid[currentPoint[0]][currentPoint[1]]
	*path = append(*path, currentPath)
	if currentValue == "\\" || currentValue == "/" {
		newDirection := calculateReflectDirection(currentValue, direction)
		nextPoint := getNextPoint(currentPoint, newDirection)
		Walk(grid, seen, nextPoint, newDirection, path)
		return
	}
	if currentValue == "|" || currentValue == "-" {
		newDirections := calculateSplitDirections(currentValue, direction)
		for _, newDirection := range newDirections {
			nextPoint := getNextPoint(currentPoint, newDirection)
			Walk(grid, seen, nextPoint, newDirection, path)
		}
		return
	}
	nextPoint := getNextPoint(currentPoint, direction)
	Walk(grid, seen, nextPoint, direction, path)
}

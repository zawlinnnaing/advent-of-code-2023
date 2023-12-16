package main

import (
	"math"
)

func findArea(path *Path) int {
	area := 0
	for i := 0; i < path.Len(); i++ {
		nextIdx := (i + 1) % path.Len()
		currentPoint := path.points[i]
		nextPoint := path.points[nextIdx]

		area += currentPoint.colIdx*nextPoint.rowIdx - nextPoint.colIdx*currentPoint.rowIdx
	}
	return int(math.Abs(float64(area / 2)))
}

func RunPart2(maze [][]string) int {
	path := Path{points: make([]Point, 0)}
	startPoint := findStartPoint(maze)
	seen := [][]bool{}
	for i := 0; i < len(maze[0]); i++ {
		innerSeen := make([]bool, len(maze[0]))
		seen = append(seen, innerSeen)
	}
	walk(maze, startPoint, seen, true, &path)
	area := findArea(&path)
	return int(math.Abs(float64(area + 1 - (path.Len() / 2))))
}

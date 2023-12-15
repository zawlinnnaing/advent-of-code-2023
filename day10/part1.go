package main

import (
	"strings"
)

var dirs = [4][2]int{
	{0, 1}, {0, -1}, {1, 0}, {-1, 0},
}

var pipeDirMap = map[string][2][2]int{
	"|": {
		{-1, 0}, {1, 0},
	},
	"-": {{0, -1}, {0, 1}},
	"F": {{1, 0}, {0, 1}},
	"7": {{1, 0}, {0, -1}},
	"L": {{0, 1}, {-1, 0}},
	"J": {{0, -1}, {-1, 0}},
}

func findStartPoint(maps [][]string) Point {
	var startPoint Point
	for rowIdx, row := range maps {
		for colIdx, cell := range row {
			if cell == "S" {
				startPoint.rowIdx = rowIdx
				startPoint.colIdx = colIdx
				return startPoint
			}
		}
	}
	return startPoint
}

func outOfBound(pipes [][]string, point Point) bool {
	maxRowIdx := len(pipes)
	maxColIdx := len(pipes[0])
	return point.rowIdx >= maxRowIdx || point.rowIdx < 0 || point.colIdx < 0 || point.colIdx >= maxColIdx
}

func canWalk(pipes [][]string, currentPoint Point, nextPoint Point) bool {
	if outOfBound(pipes, nextPoint) {
		return false
	}
	currentPipe := pipes[currentPoint.rowIdx][currentPoint.colIdx]
	if currentPipe == "S" {
		return true
	}
	nextPipe := pipes[nextPoint.rowIdx][nextPoint.colIdx]
	if nextPipe == "." {
		return false
	}
	walkableCoords := pipeDirMap[currentPipe]
	isWalkable := false
	relativeCoord := [2]int{0, 0}
	for _, coords := range walkableCoords {
		if currentPoint.rowIdx+coords[0] == nextPoint.rowIdx && currentPoint.colIdx+coords[1] == nextPoint.colIdx {
			isWalkable = true
			relativeCoord = coords
		}
	}
	if !isWalkable {
		return false
	}
	switch currentPipe {
	case "|":
		if nextPipe == "|" {
			return true
		}
		if relativeCoord == walkableCoords[0] {
			return strings.Contains("F7", nextPipe)
		}
		return strings.Contains("LJ", nextPipe)
	case "-":
		if nextPipe == "-" {
			return true
		}
		if relativeCoord == walkableCoords[0] {
			return strings.Contains("FL", nextPipe)
		}
		return strings.Contains("7J", nextPipe)
	case "F":
		if relativeCoord == walkableCoords[0] {
			return strings.Contains("|LJ", nextPipe)
		}
		return strings.Contains("-7J", nextPipe)
	case "7":
		if relativeCoord == walkableCoords[0] {
			return strings.Contains("|LJ", nextPipe)
		}
		return strings.Contains("-LF", nextPipe)
	case "L":
		if relativeCoord == walkableCoords[0] {
			return strings.Contains("7J-", nextPipe)
		}
		return strings.Contains("|7F", nextPipe)
	case "J":
		if relativeCoord == walkableCoords[0] {
			return strings.Contains("-FL", nextPipe)
		}
		return strings.Contains("|7F", nextPipe)
	}

	return false

}

func walk(maze [][]string, currentPoint Point, seen [][]bool, initialCall bool, path *Path) bool {
	//Base case: Off the map
	if outOfBound(maze, currentPoint) {
		return false
	}
	// Base case: hit the wall
	if maze[currentPoint.rowIdx][currentPoint.colIdx] == "." {
		return false
	}
	// Base case: arrive start
	if !initialCall && maze[currentPoint.rowIdx][currentPoint.colIdx] == "S" {
		return true
	}
	//	Base case: already visited
	if seen[currentPoint.rowIdx][currentPoint.colIdx] {
		return false
	}
	seen[currentPoint.rowIdx][currentPoint.colIdx] = true
	path.Push(currentPoint)
	for _, dir := range dirs {
		nextPoint := Point{
			rowIdx: currentPoint.rowIdx + dir[0],
			colIdx: currentPoint.colIdx + dir[1],
		}
		if canWalk(maze, currentPoint, nextPoint) && walk(maze, nextPoint, seen, false, path) {
			return true
		}
	}
	return false
}

func RunPart1(maze [][]string) int {
	startPoint := findStartPoint(maze)
	seen := make([][]bool, 0)
	for i := 0; i < len(maze[0]); i++ {
		innerSeen := make([]bool, len(maze[0]))
		seen = append(seen, innerSeen)
	}
	path := Path{points: make([]Point, 0)}
	walk(maze, startPoint, seen, true, &path)
	return path.Len() / 2
}

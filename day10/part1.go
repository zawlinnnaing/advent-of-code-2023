package main

import "fmt"

var dirs = [4][2]int{
	{0, 1}, {0, -1}, {1, 0}, {-1, 0},
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

func findPath(maze [][]string, currentPoint Point, seen [][]bool, initialCall bool, path *Path) bool {
	maxRowIdx := len(maze)
	maxColIdx := len(maze[0])

	fmt.Println("currentPoint", currentPoint, path)
	//Base case: Off the map
	if currentPoint.rowIdx >= maxRowIdx || currentPoint.rowIdx < 0 || currentPoint.colIdx < 0 || currentPoint.colIdx >= maxColIdx {
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
		if findPath(maze, Point{
			rowIdx: currentPoint.rowIdx + dir[0],
			colIdx: currentPoint.colIdx + dir[1],
		}, seen, false, path) {
			return true
		}
	}
	path.Pop()
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
	findPath(maze, startPoint, seen, true, &path)
	fmt.Println("path", path, path.Len())
	return path.Len() / 2
}

package main

type Point struct {
	rowIdx int
	colIdx int
}

type Path struct {
	points []Point
}

func (path *Path) Pop() (v Point) {
	pathLength := len(path.points)
	if len(path.points) == 0 {
		return
	}
	if len(path.points) == 1 {
		v = path.points[0]
		path.points = []Point{}
		return
	}
	newPoints, lastItem := path.points[:pathLength-2], path.points[pathLength-1]
	path.points = newPoints
	return lastItem
}

func (path *Path) Push(point Point) {
	path.points = append(path.points, point)
}

func (path *Path) Len() int {
	return len(path.points)
}

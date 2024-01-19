package main

import (
	"math"
)

var directionsMap = map[string][2]int{
	"R": {0, 1},
	"L": {0, -1},
	"U": {-1, 0},
	"D": {1, 0},
}

type Vertex struct {
	row   int
	col   int
	color string
}

type DigPlan struct {
	direction string
	step      int
	color     string
}

func getDigVertices(plans []DigPlan) []Vertex {
	var vertices []Vertex
	currentVertex := Vertex{
		row:   0,
		col:   0,
		color: "",
	}
	vertices = append(vertices, currentVertex)
	for _, plan := range plans {
		direction := directionsMap[plan.direction]
		//vertex := Vertex{
		//	row:   currentVertex.row + direction[0],
		//	col:   currentVertex.col + direction[1],
		//	color: plan.color,
		//}
		//currentVertex = vertex
		verticesInDirection := []Vertex{}
		for i := 0; i < plan.step; i++ {
			vertex := Vertex{
				row:   currentVertex.row + direction[0],
				col:   currentVertex.col + direction[1],
				color: plan.color,
			}
			verticesInDirection = append(verticesInDirection, vertex)
			currentVertex = vertex
		}
		vertices = append(vertices, verticesInDirection...)
	}
	return vertices
}

func getArea(vertices []Vertex) float64 {
	var area float64 = 0
	numVertices := len(vertices)
	for vertexIdx, vertex := range vertices {
		j := (vertexIdx + 1) % numVertices
		area += float64(vertex.row*vertices[j].col - vertex.col*vertices[j].row)
	}
	area = math.Abs(area) / 2
	return area
}

func getInterior(area float64, numOfPoint int) int {
	// Using pick's Theorem to get number of points interior, https://en.wikipedia.org/wiki/Pick%27s_theorem
	interior := area + 1 - (float64(numOfPoint) / 2)
	return int(interior) + numOfPoint
}

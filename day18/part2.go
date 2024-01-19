package main

import (
	"math"
	"strconv"
)

type TransformedPlan struct {
	direction string
	step      int
}

type AnalyzeResult struct {
	maxRows   int
	maxCols   int
	numPoints int
}

type Side struct {
	startPoint [2]int
	endPoint   [2]int
}

var colorToDirectionMap = map[string]string{
	"0": "R",
	"1": "D",
	"2": "L",
	"3": "U",
}

func getEndPoint(startPoint [2]int, plan TransformedPlan) [2]int {
	directionPoint := directionsMap[plan.direction]
	return [2]int{startPoint[0] + directionPoint[0]*plan.step, startPoint[1] + directionPoint[1]*plan.step}
}

func TransformPlans(plans *[]DigPlan) *[]TransformedPlan {
	var transformedPlans []TransformedPlan
	for _, plan := range *plans {
		hexadecimal := plan.color[1:6]
		step, err := strconv.ParseInt(hexadecimal, 16, 64)
		if err != nil {
			panic(err)
		}
		direction := colorToDirectionMap[string([]rune(plan.color)[6])]
		transformedPlan := TransformedPlan{
			step:      int(step),
			direction: direction,
		}
		transformedPlans = append(transformedPlans, transformedPlan)
	}
	return &transformedPlans
}

func analyze(plans *[]TransformedPlan) (AnalyzeResult, *[]Side) {
	var analyzeResult AnalyzeResult
	var sides []Side
	var currentPoint = [2]int{0, 0}
	for _, plan := range *plans {
		endpoint := getEndPoint(currentPoint, plan)
		side := Side{
			startPoint: currentPoint,
			endPoint:   endpoint,
		}
		analyzeResult.numPoints += int(math.Abs(float64(side.endPoint[0] - side.startPoint[0])))
		analyzeResult.numPoints += int(math.Abs(float64(side.endPoint[1] - side.startPoint[1])))
		if endpoint[0] > analyzeResult.maxRows {
			analyzeResult.maxRows = endpoint[0]
		}
		if endpoint[1] > analyzeResult.maxCols {
			analyzeResult.maxCols = endpoint[1]
		}
		currentPoint = endpoint
		sides = append(sides, side)
	}
	return analyzeResult, &sides
}

func RunPart2(plans []DigPlan) int {
	transformedPlans := TransformPlans(&plans)
	result, sides := analyze(transformedPlans)
	vertices := []Vertex{}
	for _, side := range *sides {
		vertices = append(vertices, Vertex{
			row:   side.startPoint[0],
			col:   side.startPoint[1],
			color: "",
		}, Vertex{
			row: side.endPoint[0],
			col: side.endPoint[1],
		})
	}
	area := getArea(vertices)
	return getInterior(area, result.numPoints)
}

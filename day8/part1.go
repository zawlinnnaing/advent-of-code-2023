package main

import "strings"

type NetworkMap map[string][2]string

func RunPart1(direction string, network NetworkMap) int {
	numSteps := 0
	directions := strings.Split(direction, "")
	currentPos := network["AAA"]
	for {
		for _, dir := range directions {
			numSteps += 1
			endNode := currentPos[0]
			if dir == "R" {
				endNode = currentPos[1]
			}
			if endNode == "ZZZ" {
				return numSteps
			}
			currentPos = network[endNode]
		}
	}
}

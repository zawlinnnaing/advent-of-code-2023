package main

import "strings"

type NetworkMap [][3]string

func RunPart1(direction string, mappings NetworkMap) int {
	numSteps := 0
	directions := strings.Split(direction, "")
	network := make(map[string][2]string)
	for _, mapping := range mappings {
		key := mapping[0]
		value := mapping[1:]
		network[key] = [2]string{
			value[0], value[1],
		}
	}
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

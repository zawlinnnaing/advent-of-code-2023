package main

import (
	"fmt"
	"strings"
)

func findLCM(numbers []int64) int64 {
	// Initialize the result with the first number
	result := numbers[0]

	// Iterate through the rest of the numbers
	for _, num := range numbers[1:] {
		// Calculate the least common multiple using the formula: LCM(a, b) = |a * b| / GCD(a, b)
		// where GCD is the greatest common divisor
		gcd := calculateGCD(result, num)
		lcm := result * num / gcd

		// Update the result with the new LCM
		result = lcm
	}

	return result
}

// Function to calculate the greatest common divisor using the Euclidean algorithm
func calculateGCD(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func RunPart2(direction string, mappings NetworkMap) int {
	startPoints := []string{}
	network := make(map[string][2]string)
	for _, mapping := range mappings {
		key := mapping[0]
		if strings.HasSuffix(key, "A") {
			startPoints = append(startPoints, key)
		}
		network[key] = [2]string{
			mapping[1], mapping[2],
		}
	}
	fmt.Println("startPoints", startPoints)
	directions := strings.Split(direction, "")
	stepsToReachEnd := make([]int64, len(startPoints))
	for pointIdx, startPoint := range startPoints {
		currentPoint := startPoint
		var numSteps int64 = 0
		finished := false
		for !finished {
			for _, dir := range directions {
				numSteps += 1
				idx := 0
				if dir == "R" {
					idx = 1
				}
				currentPoint = network[currentPoint][idx]
				fmt.Println(currentPoint, network[currentPoint], dir, idx)
				if strings.HasSuffix(currentPoint, "Z") {
					stepsToReachEnd[pointIdx] = numSteps
					finished = true
					break
				}
			}
		}
	}

	fmt.Println("steps to reach end", stepsToReachEnd)
	return int(findLCM(stepsToReachEnd))
}

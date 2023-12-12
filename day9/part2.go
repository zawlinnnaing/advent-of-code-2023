package main

func RunPart2(histories [][]int) int {
	total := 0
	for _, history := range histories {
		steps := getSteps(history)
		firstVal := 0
		for i := len(steps) - 1; i > 0; i-- {
			currentStep := steps[i]
			currentStepFirst := currentStep[0]
			prevStep := steps[i-1]
			prevStepFirstVal := prevStep[0]
			firstVal = prevStepFirstVal - currentStepFirst
			steps[i-1] = append([]int{firstVal}, steps[i-1]...)
		}
		total += firstVal
	}
	return total
}

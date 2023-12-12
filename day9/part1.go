package main

func checkIfAllZeros(history []int) bool {
	for _, item := range history {
		if item != 0 {
			return false
		}
	}
	return true
}

func RunPart1(histories [][]int) int {
	total := 0
	for _, history := range histories {
		steps := [][]int{
			history,
		}
		i := 0
		for {
			currentStep := steps[i]
			nextStep := []int{}
			for j := 0; j < len(currentStep)-1; j++ {
				diff := currentStep[j+1] - currentStep[j]
				nextStep = append(nextStep, diff)
			}
			steps = append(steps, nextStep)
			if checkIfAllZeros(nextStep) {
				break
			}
			i += 1
		}
		prevValue := 0
		for k := len(steps) - 1; k >= 0; k-- {
			currentStep := steps[k]
			lastValue := currentStep[len(currentStep)-1]
			newValue := prevValue + lastValue
			steps[k] = append(currentStep, newValue)
			prevValue = newValue
		}
		total += prevValue
	}
	return total
}

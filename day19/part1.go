package main

func RunPart1(workflows []Workflow, inputMaps []InputMap) int {
	workflowsMap := makeWorkflowsMap(&workflows)
	total := 0
	for _, inputMap := range inputMaps {
		result := runWorkflows(workflowsMap, inputMap)
		if result == "A" {
			total += getSum(inputMap)
		}
	}
	return total
}

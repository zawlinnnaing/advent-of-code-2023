package main

import "fmt"

type Rule struct {
	result    string
	condition string
	value     int
	inputKey  string
}

type Workflow struct {
	name  string
	rules []Rule
}

type InputMap map[string]int

func runWorkflow(workflow Workflow, inputMap InputMap) string {
	for _, rule := range workflow.rules {
		if rule.condition == "" {
			return rule.result
		}
		inputValue := inputMap[rule.inputKey]
		satisifies := false
		if rule.condition == ">" {
			satisifies = inputValue > rule.value
		} else {
			satisifies = inputValue < rule.value
		}
		if satisifies {
			return rule.result
		}
	}
	panic(fmt.Sprintf("Invalid workflow: %v, input: %v", workflow, inputMap))
}

func runWorkflows(workflowsMap map[string]Workflow, inputMap InputMap) string {
	result := "in"
	for result != "A" && result != "R" {
		workflow := workflowsMap[result]
		result = runWorkflow(workflow, inputMap)
		//fmt.Println("workflow", result, inputMap)
	}
	return result
}

func makeWorkflowsMap(workflows *[]Workflow) map[string]Workflow {
	var workflowsMap = make(map[string]Workflow)
	for _, workflow := range *workflows {
		workflowsMap[workflow.name] = workflow
	}
	return workflowsMap
}

func getSum(inputMap InputMap) int {
	sum := 0
	for _, value := range inputMap {
		sum += value
	}
	return sum
}

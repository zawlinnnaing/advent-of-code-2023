package main

import "fmt"

type Boundary struct {
	min int
	max int
}

type BoundMap map[string]Boundary

func getDistinctPossibility(input BoundMap) int {
	dp := 1
	for _, val := range input {
		dp *= val.max - val.min + 1
	}
	return dp
}

func getAccepted(workflowsMap *map[string]Workflow, workflow Workflow, ranges BoundMap) int {
	if workflow.name == "R" {
		return 0
	}
	if workflow.name == "A" {
		fmt.Println("accepted ranges", ranges)
		return getDistinctPossibility(ranges)
	}
	total := 0
	for _, rule := range workflow.rules {
		currentRange := ranges[rule.inputKey]
		if rule.condition != "" {
			passedHalf := Boundary{}
			rejectedHalf := Boundary{}
			if rule.condition == "<" {
				passedHalf = Boundary{
					min: currentRange.min,
					max: rule.value - 1,
				}
				rejectedHalf = Boundary{
					min: rule.value,
					max: currentRange.max,
				}
			} else {
				passedHalf = Boundary{
					min: rule.value + 1,
					max: currentRange.max,
				}
				rejectedHalf = Boundary{
					min: currentRange.min,
					max: rule.value,
				}
			}
			if passedHalf.max >= passedHalf.min {
				newRanges := make(BoundMap)
				for key, value := range ranges {
					newRanges[key] = value
				}
				newRanges[rule.inputKey] = passedHalf
				total += getAccepted(workflowsMap, (*workflowsMap)[rule.result], newRanges)
			}
			if rejectedHalf.max >= rejectedHalf.min {
				ranges[rule.inputKey] = rejectedHalf
			} else {
				//	If rejected half is empty
				break
			}
		} else {
			//	Fallback rule
			total += getAccepted(workflowsMap, (*workflowsMap)[rule.result], ranges)
		}
	}
	return total
}
func RunPart2(workflows []Workflow, _ []InputMap) int {
	workflows = append(workflows, []Workflow{
		{
			name: "A",
		},
		{
			name: "R",
		},
	}...)
	workflowsMap := makeWorkflowsMap(&workflows)
	defaultBoundary := Boundary{
		min: 1,
		max: 4000,
	}
	boundaryMap := BoundMap{
		"x": defaultBoundary,
		"m": defaultBoundary,
		"a": defaultBoundary,
		"s": defaultBoundary,
	}

	return getAccepted(&workflowsMap, workflowsMap["in"], boundaryMap)
}

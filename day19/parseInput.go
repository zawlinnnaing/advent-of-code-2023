package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseInput(inputPath string) ([]Workflow, []InputMap, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	mode := "workflow"
	var workflows []Workflow
	var inputMaps []InputMap
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			mode = "inputMap"
			continue
		}
		if mode == "workflow" {
			workflow, err := extractWorkflow(text)
			if err != nil {
				return nil, nil, err
			}
			workflows = append(workflows, workflow)
		} else {
			inputMap, err := extractInputMap(text)
			if err != nil {
				return nil, nil, err
			}
			inputMaps = append(inputMaps, inputMap)
		}
	}
	return workflows, inputMaps, nil
}

func extractWorkflow(input string) (Workflow, error) {
	var workflow Workflow
	idx := strings.Index(input, "{")
	workflow.name = input[:idx]
	rulesStr := input[idx+1 : len(input)-1]
	ruleStrs := strings.Split(rulesStr, ",")
	for _, ruleStr := range ruleStrs {
		var rule Rule
		conditionIdx := strings.IndexAny(ruleStr, "<>")
		if conditionIdx == -1 {
			rule.result = ruleStr
			workflow.rules = append(workflow.rules, rule)
			continue
		}
		rule.condition = string(ruleStr[conditionIdx])
		colonIdx := strings.Index(ruleStr, ":")
		rule.inputKey = ruleStr[:conditionIdx]
		rule.result = ruleStr[colonIdx+1:]
		valueStr := ruleStr[conditionIdx+1 : colonIdx]
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			return workflow, err
		}
		rule.value = value
		workflow.rules = append(workflow.rules, rule)
	}
	return workflow, nil
}

func extractInputMap(input string) (InputMap, error) {
	input = input[1 : len(input)-1]
	parts := strings.Split(input, ",")
	var inputMap = make(InputMap)
	for _, part := range parts {
		equalIdx := strings.Index(part, "=")
		key, valueStr := part[:equalIdx], part[equalIdx+1:]
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			return nil, err
		}
		inputMap[key] = value
	}
	return inputMap, nil
}

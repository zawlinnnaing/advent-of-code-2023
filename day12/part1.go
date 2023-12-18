package main

type Input struct {
	parts   []string
	records []int
}

const DAMAGED = "#"
const OPERATION = "."
const UNKNOWN = "?"

func RunPart1(inputs []Input) int {
	totalResult := 0
	for _, input := range inputs {
		//cache := map[string]map[string]int{}
		totalResult += count(input.parts, input.records)
	}
	return totalResult
}

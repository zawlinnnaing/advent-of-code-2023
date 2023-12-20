package main

import "strings"

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
		totalResult += count(strings.Join(input.parts, ""), &input.records)
	}
	return totalResult
}

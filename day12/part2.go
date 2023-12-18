package main

import "fmt"

func Unfold(input Input) Input {
	//originalParts := []string{}
	//copy(input.parts, originalParts)
	//originalRecords := []int{}
	//copy(input.records, originalRecords)
	partsLen := len(input.parts)
	recordsLen := len(input.records)
	for i := 0; i < 5; i++ {
		input.parts = append(input.parts, append([]string{UNKNOWN}, input.parts[:partsLen]...)...)
		input.records = append(input.records, input.records[:recordsLen]...)
	}
	return input
}

func RunPart2(inputs []Input) int {
	total := 0
	for _, input := range inputs {
		unfolded := Unfold(input)
		cacheMap := map[string]map[string]int{}
		result := count(unfolded.parts, unfolded.records)
		fmt.Println(unfolded, result)
		total += result
	}
	return total
}

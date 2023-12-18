package main

import (
	"fmt"
	"strings"
	"sync"
)

func Unfold(input Input) Input {
	partsLen := len(input.parts)
	recordsLen := len(input.records)
	for i := 0; i < 4; i++ {
		input.parts = append(input.parts, append([]string{UNKNOWN}, input.parts[:partsLen]...)...)
		input.records = append(input.records, input.records[:recordsLen]...)
	}
	return input
}

func RunPart2(inputs []Input) int {
	total := NewTotalCounter()
	wg := sync.WaitGroup{}
	for _, input := range inputs {
		wg.Add(1)
		go func(input2 Input) {
			unfolded := Unfold(input2)
			result := count(strings.Join(unfolded.parts, ""), unfolded.records)
			fmt.Println("finished", input2.parts, input2.records)
			total.Add(result)
			wg.Done()
		}(input)
	}
	wg.Wait()
	return total.Get()
}

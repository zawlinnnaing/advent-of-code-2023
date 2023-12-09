package main

import (
	"fmt"
	"sort"
	"strings"
)

type deal struct {
	hand string
	bid  int
	rank int
}

func determineRank(inp deal) int {
	dealMap := map[string]int{}
	cards := strings.Split(inp.hand, "")
	for _, card := range cards {
		dealMap[card] += 1
	}
	numOfThreePairs := 0
	numOfTwoPairs := 0
	for _, occurrence := range dealMap {
		if occurrence == 5 {
			return FIVE_OF_KIND
		}
		if occurrence == 4 {
			return FOUR_OF_KIND
		}
		if occurrence == 3 {
			numOfThreePairs += 1
		}
		if occurrence == 2 {
			numOfTwoPairs += 1
		}
	}

	if numOfThreePairs == 1 {
		if numOfTwoPairs == 1 {
			return FULL_HOUSE
		}
		return THREE_OF_KIND
	}
	if numOfTwoPairs == 2 {
		return TWO_PAIR
	}
	if numOfTwoPairs == 1 {
		return ONE_PAIR
	}
	return HIGH_CARD
}

func RunPart1(inputs []deal) int {
	outputs := &dealHeap{}
	for _, input := range inputs {
		(&input).rank = determineRank(input)
		outputs.Push(input)
	}

	sort.Sort(outputs)

	idx := outputs.Len()
	total := 0
	for outputs.Len() > 0 {
		output := outputs.Pop()
		total += idx * output.bid
		idx -= 1
	}
	return total
}

func SolvePart1() {
	inputs, err := parseInput("./inputs.txt")
	if err != nil {
		panic(err)
	}
	output := RunPart1(inputs)
	fmt.Println("Solution for part1 ==>", output)
}

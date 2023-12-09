package main

import (
	"fmt"
	"sort"
	"strings"
)

func determineRankWithJoker(inp deal) int {
	dealMap := map[string]int{}
	cards := strings.Split(inp.hand, "")
	for _, card := range cards {
		dealMap[card] += 1
	}
	numOfThreePairs := 0
	numOfTwoPairs := 0
	for card, occurrence := range dealMap {
		if card == "J" {
			continue
		}
		if occurrence == 5 {
			return FIVE_OF_KIND
		}
		if occurrence == 4 {
			if dealMap["J"] != 0 {
				return FIVE_OF_KIND
			}
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
		if dealMap["J"] != 0 {
			return THREE_OF_KIND + dealMap["J"] + 1
		}
		if numOfTwoPairs == 1 {
			return FULL_HOUSE
		}
		return THREE_OF_KIND
	}
	if numOfTwoPairs == 2 {
		if dealMap["J"] != 0 {
			return TWO_PAIR + dealMap["J"] + 1
		}

		return TWO_PAIR
	}
	if numOfTwoPairs == 1 {
		if dealMap["J"] != 0 {
			switch dealMap["J"] {
			case 1:
				return THREE_OF_KIND
			case 2:
				return FOUR_OF_KIND
			case 3:
				return FIVE_OF_KIND
			}
		}
		return ONE_PAIR
	}
	if strings.Contains(inp.hand, "J") {
		switch dealMap["J"] {
		case 1:
			return ONE_PAIR
		case 2:
			return THREE_OF_KIND
		case 3:
			return FOUR_OF_KIND
		default:
			return FIVE_OF_KIND
		}
	}

	return HIGH_CARD

}

func RunPart2(inputs []deal) int {
	outputs := &dealSorter{
		deals:    make([]deal, len(inputs)),
		cardsMap: cardsMapWithJoker,
	}
	for idx, input := range inputs {
		input.rank = determineRankWithJoker(input)
		outputs.deals[idx] = input
	}
	sort.Sort(outputs)
	return outputs.Total()
}

func SolvePart2() {
	inputs, err := parseInput("./inputs.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Solution for part2 ==>", RunPart2(inputs))
}

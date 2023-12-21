package main

import "fmt"

func main() {
	patterns, err := parseInput("./data/input.txt")
	if err != nil {
		panic(err)
	}
	part1 := runPart1(patterns)
	fmt.Println("Solution for part1 ===>", part1)
	part2 := runPart2(patterns)
	fmt.Println("Solution for part2 ==>", part2)
}

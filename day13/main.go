package main

import "fmt"

func main() {
	patterns, err := parseInput("./data/input.txt")
	if err != nil {
		panic(err)
	}
	part1 := runPart1(patterns)
	fmt.Println("Solution for part1 ===>", part1)
}

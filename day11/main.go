package main

import "fmt"

func main() {
	universe, err := parseInput("./data/input.txt")
	if err != nil {
		panic(err)
	}
	part1 := RunPart1(universe)
	fmt.Println("Solution for part1 ==>", part1)
}

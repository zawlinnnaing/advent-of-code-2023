package main

import "fmt"

func main() {
	grid, err := parseInput("./data/input.txt")
	if err != nil {
		panic(err)
	}
	part1 := RunPart1(grid)
	fmt.Println("Solution for part1 ==>", part1)
	part2 := RunPart2(grid)
	fmt.Println("Solution for part2 ==>", part2)
}

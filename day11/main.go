package main

import "fmt"

func main() {
	universe, err := parseInput("./data/input.txt")
	if err != nil {
		panic(err)
	}
	part1 := Solve(universe, 2)
	fmt.Println("Solution for part1 ==>", part1)
	part2 := Solve(universe, 1000000)
	fmt.Println("Solution for part2 ===>", part2)
}

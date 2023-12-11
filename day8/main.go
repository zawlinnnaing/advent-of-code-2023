package main

import "fmt"

func main() {
	direction, mappings, err := parseInput("./inputs.txt")
	if err != nil {
		panic(err)
	}
	part1output := RunPart1(direction, mappings)
	fmt.Println("Solution for part1 ===>", part1output)
	part2output := RunPart2(direction, mappings)
	fmt.Println("Solution for part2 ===>", part2output)
}

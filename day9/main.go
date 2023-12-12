package main

import "fmt"

func main() {
	inputs, err := parseInput("./inputs.txt")
	if err != nil {
		panic(err)
	}
	part1 := RunPart1(inputs)
	fmt.Println("Solution for part1 ===>", part1)
	part2 := RunPart2(inputs)
	fmt.Println("Solution for part2 ===>", part2)
}

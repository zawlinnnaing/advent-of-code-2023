package main

import "fmt"

func main() {
	inputs, err := parseInput("./inputs.txt")
	if err != nil {
		panic(err)
	}
	part1 := RunPart1(inputs)
	fmt.Println("Solution for part1 ===>", part1)
}

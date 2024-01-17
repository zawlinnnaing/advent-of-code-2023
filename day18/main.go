package main

import "fmt"

func main() {
	plans, err := parseInput("./data/input.txt")
	if err != nil {
		panic(err)
	}
	part1 := RunPart1(plans)
	fmt.Println("Solution for part1===>", part1)
}

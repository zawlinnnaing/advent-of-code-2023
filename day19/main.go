package main

import "fmt"

func main() {
	workflows, inputMaps, err := parseInput("./data/input.txt")
	if err != nil {
		panic(err)
	}
	part1 := RunPart1(workflows, inputMaps)
	fmt.Println("Solution for part1 ==>", partå1)
}

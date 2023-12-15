package main

import "fmt"

func main() {
	pipesMap, err := parseInput("./inputs.txt")
	if err != nil {
		panic(err)
	}
	part1 := RunPart1(pipesMap)
	fmt.Println("Solution for part1 ==>", part1)
}

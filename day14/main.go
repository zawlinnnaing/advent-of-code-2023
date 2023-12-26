package main

import "fmt"

func main() {
	plane, err := parseInputFile("./data/input.txt")
	if err != nil {
		panic(err)
	}
	part1 := RunPart1(plane)
	fmt.Println("Solution for part1 ===>", part1)
}

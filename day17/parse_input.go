package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseInput(inputFile string) (Grid, error) {
	var grid Grid
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		textStrs := strings.Split(scanner.Text(), "")
		numArr := []int{}
		for _, numStr := range textStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}
			numArr = append(numArr, num)
		}
		grid = append(grid, numArr)
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return grid, nil
}

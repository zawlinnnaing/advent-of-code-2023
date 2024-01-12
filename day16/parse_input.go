package main

import (
	"bufio"
	"os"
	"strings"
)

func parseInput(filePath string) (Grid, error) {
	var grid Grid
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		grid = append(grid, strings.Split(text, ""))
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

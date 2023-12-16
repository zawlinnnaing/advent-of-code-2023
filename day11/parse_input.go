package main

import (
	"bufio"
	"os"
	"strings"
)

func parseInput(filePath string) (Universe, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var galaxies Universe
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cells := strings.Split(line, "")
		galaxies = append(galaxies, cells)
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return galaxies, nil
}

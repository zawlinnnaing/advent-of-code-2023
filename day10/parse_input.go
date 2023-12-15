package main

import (
	"bufio"
	"os"
	"strings"
)

func parseInput(filePath string) ([][]string, error) {
	var pipesMap [][]string
	file, err := os.Open(filePath)
	if err != nil {
		return pipesMap, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		cells := strings.Split(row, "")
		pipesMap = append(pipesMap, cells)
	}
	if err := scanner.Err(); err != nil {
		return pipesMap, err
	}
	return pipesMap, nil
}

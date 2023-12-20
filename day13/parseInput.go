package main

import (
	"bufio"
	"os"
	"strings"
)

func parseInput(filePath string) ([]Pattern, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	patterns := []Pattern{}
	scanner := bufio.NewScanner(file)
	pattern := Pattern{}
	for scanner.Scan() {
		patternLine := scanner.Text()
		if patternLine == "" {
			patterns = append(patterns, pattern)
			pattern = Pattern{}
		} else {
			pattern = append(pattern, strings.Split(patternLine, ""))
		}
	}
	if len(pattern) > 0 {
		patterns = append(patterns, pattern)
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return patterns, nil
}

package main

import (
	"bufio"
	"os"
	"strings"
)

func parseInputFile(filePath string) (Plane, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var plane Plane
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		blocks := strings.Split(text, "")
		plane = append(plane, blocks)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return plane, err
}

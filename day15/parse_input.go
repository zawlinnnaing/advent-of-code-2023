package main

import (
	"os"
	"strings"
)

func parseInput(inputFile string) ([]string, error) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), ","), nil
}

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseInput(filePath string) ([]Input, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var inputs []Input
	for scanner.Scan() {
		input := Input{}
		text := scanner.Text()
		textSplits := strings.Split(text, " ")
		input.parts = strings.Split(textSplits[0], "")
		recordStrs := strings.Split(textSplits[1], ",")
		input.records = []int{}
		for _, recordStr := range recordStrs {
			record, err := strconv.Atoi(recordStr)
			if err != nil {
				return nil, err
			}
			input.records = append(input.records, record)
		}
		inputs = append(inputs, input)
	}
	return inputs, nil
}

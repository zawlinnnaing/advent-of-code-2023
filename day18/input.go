package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseInput(filePath string) ([]DigPlan, error) {
	var digPlans []DigPlan
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	parenRegex := regexp.MustCompile(`[()]`)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		step, err := strconv.Atoi(text[1])
		if err != nil {
			return nil, err
		}

		digPlan := DigPlan{
			direction: text[0],
			step:      step,
			color:     parenRegex.ReplaceAllString(text[2], ""),
		}
		digPlans = append(digPlans, digPlan)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return digPlans, nil
}

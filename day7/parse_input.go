package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseInput(inputFile string) ([]deal, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	outputs := []deal{}
	for scanner.Scan() {
		text := scanner.Text()
		strs := strings.Split(text, " ")
		if len(strs) != 2 {
			break
		}
		bid, err := strconv.Atoi(strs[1])
		if err != nil {
			return nil, err
		}
		outputs = append(outputs, deal{
			hand: strs[0],
			bid:  bid,
		})
	}
	return outputs, nil
}

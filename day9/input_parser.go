package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	// Create a slice of slices to store the data
	var data [][]int

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line into individual numbers
		numStr := strings.Fields(scanner.Text())
		var row []int

		// Convert each string number to an integer and append to the row slice
		for _, num := range numStr {
			val, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return nil, err
			}
			row = append(row, val)
		}

		// Append the row to the data slice
		data = append(data, row)
	}
	return data, nil
}

package main

import (
	"os"
	"strings"
)

func parseInput(file string) (string, NetworkMap, error) {
	mappings := make(map[string][2]string)
	content, err := os.ReadFile(file)
	if err != nil {
		return "", mappings, err
	}
	inputs := string(content)
	lines := strings.Split(inputs, "\n")
	direction := strings.TrimSpace(lines[0])
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

		// Split the line by "="
		parts := strings.Split(line, "=")

		// Trim spaces from both key and value
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove parentheses from the value
		value = strings.Trim(value, "()")

		// Split the value by comma
		values := strings.Split(value, ",")

		// Trim spaces from each element in values
		values[0] = strings.TrimSpace(values[0])
		values[1] = strings.TrimSpace(values[1])

		// Store the key-value pair in the map
		mappings[key] = [2]string{values[0], values[1]}
	}

	return direction, mappings, nil

}

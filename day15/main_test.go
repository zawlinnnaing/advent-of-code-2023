package main

import (
	"testing"
)

func TestProblem(t *testing.T) {
	testCases := []struct {
		name      string
		inputFile string
		expected  int
		run       func(input []string) int
	}{
		{
			name:      "Part1",
			inputFile: "./data/example.txt",
			expected:  1320,
			run:       RunPart1,
		},
		{
			name:      "Part2",
			inputFile: "./data/example.txt",
			expected:  145,
			run:       RunPart2,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			input, err := parseInput(testCase.inputFile)
			if err != nil {
				t.Error(err)
			}
			output := testCase.run(input)
			if output != testCase.expected {
				t.Errorf("Expected: %d, recevied: %d", testCase.expected, output)
			}
		})
	}
}

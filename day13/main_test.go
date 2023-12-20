package main

import (
	"testing"
)

func TestProblem(t *testing.T) {
	testCases := []struct {
		name      string
		run       func(patterns []Pattern) int
		inputFile string
		expected  int
	}{
		{
			name:      "Part1",
			run:       runPart1,
			inputFile: "./data/example.txt",
			expected:  405,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			patterns, err := parseInput(testCase.inputFile)
			if err != nil {
				t.Error(err)
			}
			output := testCase.run(patterns)
			if output != testCase.expected {
				t.Errorf("Expected: %d, received: %d", testCase.expected, output)
			}
		})
	}
}

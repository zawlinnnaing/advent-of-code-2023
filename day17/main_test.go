package main

import "testing"

func TestProblem(t *testing.T) {
	testCases := []struct {
		name      string
		run       func(grid Grid) int
		inputFile string
		expected  int
	}{
		{
			name:      "Part1",
			run:       RunPart1,
			inputFile: "./data/example.txt",
			expected:  102,
		},
		{
			name:      "Part2",
			run:       RunPart2,
			inputFile: "./data/example.txt",
			expected:  94,
		},
		{
			name:      "Part2 - Example2",
			run:       RunPart2,
			inputFile: "./data/example2.txt",
			expected:  71,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			grid, err := parseInput(testCase.inputFile)
			if err != nil {
				t.Error(err)
			}
			output := testCase.run(grid)
			if output != testCase.expected {
				t.Errorf("Expected %d, received %d", testCase.expected, output)
			}
		})
	}
}

package main

import "testing"

func TestProblem(t *testing.T) {
	testCases := []struct {
		name     string
		filePath string
		expected int
		run      func(grid Grid) int
	}{
		{
			name:     "Part1",
			filePath: "./data/example.txt",
			expected: 46,
			run:      RunPart1,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			grid, err := parseInput(testCase.filePath)
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

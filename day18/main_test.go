package main

import "testing"

func TestProblem(t *testing.T) {
	testCases := []struct {
		name      string
		inputFile string
		expected  int
		run       func(plans []DigPlan) int
	}{
		{
			name:      "Part1",
			inputFile: "./data/example.txt",
			expected:  62,
			run:       RunPart1,
		},
		{
			name:      "Part2",
			inputFile: "./data/example.txt",
			expected:  952408144115,
			run:       RunPart2,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			plans, err := parseInput(testCase.inputFile)
			if err != nil {
				t.Fatal(err)
			}
			output := testCase.run(plans)
			if output != testCase.expected {
				t.Errorf("Expected %d, received %d", testCase.expected, output)
			}
		})
	}
}

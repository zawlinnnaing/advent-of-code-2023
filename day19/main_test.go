package main

import "testing"

func TestProblem(t *testing.T) {
	testCases := []struct {
		name      string
		inputFile string
		expected  int
		run       func(workflows []Workflow, inputMaps []InputMap) int
	}{
		{
			name:      "Part1",
			inputFile: "./data/example.txt",
			expected:  19114,
			run:       RunPart1,
		},
		{
			name:      "Part2",
			inputFile: "./data/example.txt",
			expected:  167409079868000,
			run:       RunPart2,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			workflows, inputMaps, err := parseInput(testCase.inputFile)
			if err != nil {
				t.Fatal(err)
			}
			output := testCase.run(workflows, inputMaps)
			if output != testCase.expected {
				t.Errorf("Expected %d, received: %d, difference: %d", testCase.expected, output, testCase.expected-output)
			}
		})
	}
}

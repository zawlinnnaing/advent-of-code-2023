package main

import "testing"

func TestProblem(t *testing.T) {
	testCases := []struct {
		name      string
		expected  int
		run       func(inputs []Input) int
		inputFile string
	}{
		{
			name:      "Part1",
			expected:  21,
			run:       RunPart1,
			inputFile: "./data/example.txt",
		},
		{
			name:      "Part2",
			expected:  525152,
			run:       RunPart2,
			inputFile: "./data/example.txt",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			inputs, err := parseInput(testCase.inputFile)
			if err != nil {
				t.Error(err)
			}
			output := testCase.run(inputs)
			if output != testCase.expected {
				t.Errorf("Expeced %d, recevied %d", testCase.expected, output)
			}
		})
	}
}

package main

import "testing"

func TestProblem(t *testing.T) {
	testCases := []struct {
		filePath string
		expect   int
		run      func(plane Plane) int
		name     string
	}{
		{
			filePath: "./data/example.txt",
			expect:   136,
			run:      RunPart1,
			name:     "Part1",
		},
		{
			filePath: "./data/example.txt",
			expect:   64,
			run:      RunPart2,
			name:     "Part2",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			input, err := parseInputFile(testCase.filePath)
			if err != nil {
				t.Error(err)
			}
			output := testCase.run(input)
			if output != testCase.expect {
				t.Errorf("Expected %d, received %d", testCase.expect, output)
			}
		})
	}
}

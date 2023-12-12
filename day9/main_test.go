package main

import "testing"

func TestProblem(t *testing.T) {
	histories := [][]int{
		{0, 3, 6, 9, 12, 15},
		{1, 3, 6, 10, 15, 21},
		{10, 13, 16, 21, 30, 45},
	}
	testCases := []struct {
		name   string
		expect int
		run    func(inputs [][]int) int
	}{
		{
			expect: 114,
			name:   "Part1",
			run:    RunPart1,
		},
		{
			expect: 2,
			name:   "Part2",
			run:    RunPart2,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			output := testCase.run(histories)
			if output != testCase.expect {
				t.Errorf("Expected output to be: %d, received %d", testCase.expect, output)
			}
		})
	}
}

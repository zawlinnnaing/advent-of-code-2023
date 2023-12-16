package main

import (
	"testing"
)

func TestProblem(t *testing.T) {
	testCases := []struct {
		filePath      string
		name          string
		exp           int
		stepsToExpand int
		run           func(galaxies Universe, stepsToExpand int) int
	}{
		{
			filePath:      "./data/map1.txt",
			name:          "Part1 - Map1",
			exp:           374,
			stepsToExpand: 2,
			run:           Solve,
		},
		{
			filePath:      "./data/map1.txt",
			name:          "Part2 -Map1",
			exp:           1030,
			stepsToExpand: 10,
			run:           Solve,
		},
		{
			filePath:      "./data/map1.txt",
			name:          "Part2 -Map1",
			exp:           8410,
			stepsToExpand: 100,
			run:           Solve,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			galaxies, err := parseInput(testCase.filePath)
			if err != nil {
				t.Error(err)
			}
			output := testCase.run(galaxies, testCase.stepsToExpand)
			if output != testCase.exp {
				t.Errorf("Expected output to be %d, received: %d", testCase.exp, output)
			}
		})
	}
}

func TestSanity(t *testing.T) {
	expSteps := 96580
	steps := 0
	for i := 1; i < 440; i++ {
		steps += i
	}
	if expSteps != steps {
		t.Errorf("Sanity test failed. %d, %d", expSteps, steps)
	}
}

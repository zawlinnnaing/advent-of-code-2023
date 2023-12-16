package main

import (
	"testing"
)

func TestProblem(t *testing.T) {
	testCases := []struct {
		filePath string
		name     string
		exp      int
		run      func(galaxies Universe) int
	}{
		{
			filePath: "./data/map1.txt",
			name:     "Part1 - Map1",
			exp:      374,
			run:      RunPart1,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			galaxies, err := parseInput(testCase.filePath)
			if err != nil {
				t.Error(err)
			}
			output := testCase.run(galaxies)
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

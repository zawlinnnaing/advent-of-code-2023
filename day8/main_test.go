package main

import "testing"

func TestProblem(t *testing.T) {
	testCases := []struct {
		direction string
		network   NetworkMap
		expSteps  int
	}{
		{
			direction: "RL",
			network: map[string][2]string{
				"AAA": {"BBB", "CCC"},
				"BBB": {"DDD", "EEE"},
				"CCC": {"ZZZ", "GGG"},
			},
			expSteps: 2,
		},
		{
			direction: "LLR",
			network: map[string][2]string{
				"AAA": {"BBB", "BBB"},
				"BBB": {"AAA", "ZZZ"},
				"ZZZ": {"ZZZ", "ZZZ"},
			},
			expSteps: 6,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.direction, func(t *testing.T) {
			output := RunPart1(testCase.direction, testCase.network)
			if output != testCase.expSteps {
				t.Errorf("Expected steps: %d, received: %d", testCase.expSteps, output)
			}
		})
	}
}

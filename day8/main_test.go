package main

import "testing"

func TestPart1(t *testing.T) {
	testCases := []struct {
		direction string
		network   NetworkMap
		expSteps  int
	}{
		{
			direction: "RL",
			network: NetworkMap{
				{"AAA", "BBB", "CCC"},
				{"BBB", "DDD", "EEE"},
				{"CCC", "ZZZ", "GGG"},
			},
			expSteps: 2,
		},
		{
			direction: "LLR",
			network: NetworkMap{
				{"AAA", "BBB", "BBB"},
				{"BBB", "AAA", "ZZZ"},
				{"ZZZ", "ZZZ", "ZZZ"},
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

func TestPart2(t *testing.T) {
	direction := "LR"
	mappings := NetworkMap{
		{"11A", "11B", "XXX"},
		{"11B", "XXX", "11Z"},
		{"11Z", "11B", "XXX"},
		{"22A", "22B", "XXX"},
		{"22B", "22C", "22C"},
		{"22C", "22Z", "22Z"},
		{"22Z", "22B", "22B"},
		{"XXX", "XXX", "XXX"},
	}
	expSteps := 6
	output := RunPart2(direction, mappings)
	if expSteps != output {
		t.Errorf("Expected steps to be: %d, received: %d", expSteps, output)
	}

}

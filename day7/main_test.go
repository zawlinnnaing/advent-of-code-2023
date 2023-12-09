package main

import "testing"

func TestRun(t *testing.T) {
	inputs := []deal{
		{hand: "32T3K", bid: 765},
		{
			hand: "T55J5",
			bid:  684,
		},
		{
			hand: "KK677",
			bid:  28,
		},
		{
			hand: "KTJJT",
			bid:  220,
		},
		{
			hand: "QQQJA",
			bid:  483,
		},
	}
	testCases := []struct {
		run      func(inputs []deal) int
		expected int
		name     string
	}{
		{
			name:     "Part1",
			run:      RunPart1,
			expected: 6440,
		},
		{
			name:     "Part2",
			run:      RunPart2,
			expected: 5905,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			expected := testCase.expected
			output := testCase.run(inputs)
			if expected != output {
				t.Errorf("Expected total to be: %d, received: %d", expected, output)
			}
		})
	}

}

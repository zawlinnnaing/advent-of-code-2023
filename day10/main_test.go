package main

import "testing"

func TestProblem(t *testing.T) {
	map1 := [][]string{
		{".", ".", ".", ".", "."},
		{".", "S", "-", "7", "."},
		{".", "|", ".", "|", "."},
		{".", "L", "-", "J", "."},
		{".", ".", ".", ".", "."},
	}
	map2 := [][]string{
		{".", ".", "F", "7", "."},
		{".", "F", "J", "|", "."},
		{"S", "J", ".", "L", "7"},
		{"|", "F", "-", "-", "J"},
		{"L", "J", ".", ".", "."},
	}

	map3, err := parseInput("./map3.txt")
	map4, err := parseInput("./map4.txt")
	map5, err := parseInput("./map5.txt")
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name    string
		pattern [][]string
		exp     int
		run     func(m [][]string) int
	}{
		{
			name:    "Part1 - Map1",
			pattern: map1,
			exp:     4,
			run:     RunPart1,
		},
		{
			name:    "Part1 - Map2",
			pattern: map2,
			exp:     8,
			run:     RunPart1,
		},
		{
			name:    "Part2 - Map3",
			pattern: map3,
			exp:     4,
			run:     RunPart2,
		},
		{
			name:    "Part2 - Map4",
			pattern: map4,
			exp:     8,
			run:     RunPart2,
		},
		{
			name:    "Part2 - map5",
			pattern: map5,
			exp:     10,
			run:     RunPart2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			output := testCase.run(testCase.pattern)
			if output != testCase.exp {
				t.Errorf("Expected output: %d, received %d", testCase.exp, output)
			}
		})
	}
}

package main

import "testing"

func TestPart1(t *testing.T) {
	testCase := struct {
		races  []Race
		output int
	}{
		races: []Race{
			{
				Time:   7,
				Record: 9,
			},
			{
				Time:   15,
				Record: 40,
			},
			{
				Time:   30,
				Record: 200,
			},
		},
		output: 288,
	}
	output := RunPart1(testCase.races)
	if output != testCase.output {
		t.Errorf("Expected output to be %d,recevied %d", testCase.output, output)
	}
}

func TestPart2(t *testing.T) {
	testCase := struct {
		race   Race
		output int
	}{
		race: Race{
			Time:   71530,
			Record: 940200,
		},
		output: 71503,
	}
	output := FindWaysToBeat(testCase.race)
	if output != testCase.output {
		t.Errorf("Expect output to be %d, recevied: %d", testCase.output, output)
	}
}

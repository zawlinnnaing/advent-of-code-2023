package main

import "testing"

func TestExample(t *testing.T) {
	testCases := []struct {
		input             []int
		part1             int
		part2             int
		seedToSoil        []RangeMap
		soilToFertilizer  []RangeMap
		fertilizerToWater []RangeMap
		waterToLight      []RangeMap
		lightToTemp       []RangeMap
		tempToHumid       []RangeMap
		humidToLocation   []RangeMap
	}{
		{
			seedToSoil: []RangeMap{
				{
					sourceStart: 98,
					destStart:   50,
					rangeLength: 2,
				},
				{
					sourceStart: 50,
					destStart:   52,
					rangeLength: 48,
				},
			},
			soilToFertilizer: []RangeMap{
				{
					sourceStart: 15,
					destStart:   0,
					rangeLength: 37,
				},
				{
					sourceStart: 52,
					destStart:   37,
					rangeLength: 2,
				},
				{
					sourceStart: 0,
					destStart:   39,
					rangeLength: 15,
				},
			},
			fertilizerToWater: []RangeMap{
				{
					sourceStart: 53,
					destStart:   49,
					rangeLength: 8,
				},
				{
					sourceStart: 11,
					destStart:   0,
					rangeLength: 42,
				},
				{
					sourceStart: 0,
					destStart:   42,
					rangeLength: 7,
				},
				{
					sourceStart: 7,
					destStart:   57,
					rangeLength: 4,
				},
			},
			waterToLight: []RangeMap{
				{
					sourceStart: 18,
					destStart:   88,
					rangeLength: 7,
				},
				{
					sourceStart: 25,
					destStart:   18,
					rangeLength: 70,
				},
			},
			lightToTemp: []RangeMap{
				{
					sourceStart: 77,
					destStart:   45,
					rangeLength: 23,
				},
				{
					sourceStart: 45,
					destStart:   81,
					rangeLength: 19,
				},
				{
					sourceStart: 64,
					destStart:   68,
					rangeLength: 13,
				},
			},
			tempToHumid: []RangeMap{
				{
					sourceStart: 69,
					destStart:   0,
					rangeLength: 1,
				},
				{
					sourceStart: 0,
					destStart:   1,
					rangeLength: 69,
				},
			},
			humidToLocation: []RangeMap{
				{
					sourceStart: 56,
					destStart:   60,
					rangeLength: 37,
				},
				{
					sourceStart: 93,
					destStart:   56,
					rangeLength: 4,
				},
			},
			input: []int{79, 14, 55, 13},
			part1: 35,
			part2: 46,
		},
	}
	for _, testCase := range testCases {
		output := Run(testCase.input, testCase.seedToSoil, testCase.soilToFertilizer, testCase.fertilizerToWater, testCase.waterToLight, testCase.lightToTemp, testCase.tempToHumid, testCase.humidToLocation)
		if output != testCase.part1 {
			t.Errorf("Expected part1 to be: %d, recevied: %d", testCase.part1, output)
		}
		part2Output := RunPart2(testCase.input, testCase.seedToSoil, testCase.soilToFertilizer, testCase.fertilizerToWater, testCase.waterToLight, testCase.lightToTemp, testCase.tempToHumid, testCase.humidToLocation)
		if part2Output != testCase.part2 {
			t.Errorf("Expect part2 result to be: %d, recevied: %d", testCase.part2, part2Output)
		}
	}

}

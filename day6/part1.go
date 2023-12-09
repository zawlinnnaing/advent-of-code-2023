package main

import (
	"fmt"
	"math"
)

type Race struct {
	Time   int
	Record int
}

func FindWaysToBeat(race Race) int {
	holdOffset := int(math.Ceil(float64(race.Record) / float64(race.Time)))
	holdOffsetToBeat := 0
	for i := holdOffset; i < race.Time; i++ {
		raceTime := race.Time - i
		record := raceTime * i
		if record > race.Record {
			holdOffsetToBeat = i
			break
		}
	}
	if holdOffsetToBeat == 0 {
		return 0
	}
	return (race.Time - holdOffsetToBeat) - (holdOffsetToBeat) + 1
}

func RunPart1(races []Race) int {
	totalWays := 1
	for _, race := range races {
		totalWays *= FindWaysToBeat(race)
	}
	return totalWays
}

func SolvePart1() {
	solution := RunPart1(part1Inputs)
	fmt.Println("Solution for part 1 is ==>", solution)
}

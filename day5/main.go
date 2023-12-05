package main

import (
	"fmt"
	"os"
	"path"
	"slices"
)

type RangeMap struct {
	sourceStart int
	destStart   int
	rangeLength int
}

func findDest(source int, rangeMaps []RangeMap) int {
	var dest = source
	for _, rangeMap := range rangeMaps {
		if rangeMap.sourceStart > source {
			continue
		}
		sourceRangeEnd := rangeMap.sourceStart + rangeMap.rangeLength
		if rangeMap.sourceStart <= source && sourceRangeEnd >= source {
			destLength := source - rangeMap.sourceStart
			dest = rangeMap.destStart + destLength
		}
	}
	return dest
}

func Run(initialSeeds []int, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemp, tempToHumid, humidToLocation []RangeMap) int {
	dests := make([]int, len(initialSeeds))
	for idx, seed := range initialSeeds {
		soil := findDest(seed, seedToSoil)
		fertilizer := findDest(soil, soilToFertilizer)
		water := findDest(fertilizer, fertilizerToWater)
		light := findDest(water, waterToLight)
		temp := findDest(light, lightToTemp)
		humid := findDest(temp, tempToHumid)
		location := findDest(humid, humidToLocation)
		dests[idx] = location
	}
	return slices.Min(dests)
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	seeds, maps, err := ParseFile(path.Join(dir, "input.txt"))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	output := Run(seeds, maps[0], maps[1], maps[2], maps[3], maps[4], maps[5], maps[6])
	fmt.Println("Result is ==>:", output)
}

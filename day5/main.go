package main

import (
	"fmt"
	"math"
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
	usedMap := RangeMap{}
	for _, rangeMap := range rangeMaps {
		if rangeMap.sourceStart > source {
			continue
		}

		sourceRangeEnd := rangeMap.sourceStart + rangeMap.rangeLength
		if rangeMap.sourceStart <= source && sourceRangeEnd >= source {
			usedMap = rangeMap
			destLength := source - rangeMap.sourceStart
			dest = rangeMap.destStart + destLength
			break
		}
	}
	if dest <= 10 {
		fmt.Println("Dest:", dest, source, "used map:", usedMap)
	}
	return dest
}

func findSeedLocation(seed int, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemp, tempToHumid, humidToLocation []RangeMap) int {
	soil := findDest(seed, seedToSoil)
	fertilizer := findDest(soil, soilToFertilizer)
	water := findDest(fertilizer, fertilizerToWater)
	light := findDest(water, waterToLight)
	temp := findDest(light, lightToTemp)
	humid := findDest(temp, tempToHumid)
	location := findDest(humid, humidToLocation)
	return location
}

func divideRange(inputRange []int, seedMaps []RangeMap) [][]int {
	output := [][]int{}
	if len(inputRange) < 2 {
		panic("Input range less than 2")
		return [][]int{inputRange}
	}
	passThroughRange := []int{inputRange[0], inputRange[1]}

	for _, seedMap := range seedMaps {
		if passThroughRange == nil {
			break
		}

		mapStart := seedMap.sourceStart
		mapEnd := mapStart + seedMap.rangeLength
		ptStart := passThroughRange[0]
		ptEnd := passThroughRange[1]

		if ptStart == ptEnd {
			output = append(output, passThroughRange)
			passThroughRange = nil
		}

		if ptEnd == mapStart {
			output = append(output, []int{mapStart, mapStart})
			passThroughRange = []int{
				ptStart,
				int(math.Max(float64(ptStart), float64(ptEnd-1))),
			}
			continue
		}
		if ptStart == mapEnd {
			output = append(output, []int{mapEnd, mapEnd})
			passThroughRange = []int{
				int(math.Min(float64(ptEnd), float64(mapEnd+1))),
				ptEnd,
			}
			continue
		}

		// Right side overlapping
		if ptEnd >= mapStart && ptEnd <= mapEnd {
			newStart := int(math.Max(float64(ptStart), float64(mapStart)))
			newRange := []int{
				newStart,
				ptEnd,
			}
			output = append(output, newRange)
			// Meaning all values in pass through range falls into one map
			if newStart == ptStart {
				passThroughRange = nil
				break
			}
			passThroughRange = []int{
				ptStart,
				int(math.Max(float64(newStart-1), float64(ptStart))),
			}

		} else if ptStart >= mapStart && ptStart <= mapEnd {
			// Left side overlapping
			newEnd := int(math.Min(float64(ptEnd), float64(mapEnd)))
			output = append(output, []int{
				ptStart,
				newEnd,
			})
			if ptEnd == newEnd {
				passThroughRange = nil
				break
			}
			passThroughRange = []int{
				int(math.Min(float64(newEnd+1), float64(ptEnd))),
				ptEnd,
			}
		}
	}
	// Meaning there are still some values not fallen into maps
	if passThroughRange != nil {
		output = append(output, passThroughRange)
	}
	return output
}

func Run(initialSeeds []int, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemp, tempToHumid, humidToLocation []RangeMap) int {
	dests := make([]int, len(initialSeeds))
	for idx, seed := range initialSeeds {
		location := findSeedLocation(seed, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemp, tempToHumid, humidToLocation)
		dests[idx] = location
	}
	return slices.Min(dests)
}

func divideAndMap(inputRanges [][]int, seedMaps []RangeMap) [][]int {
	divided := [][]int{}
	for _, inputRange := range inputRanges {
		divided = append(divided, divideRange(inputRange, seedMaps)...)
	}
	outputs := [][]int{}
	for _, inputRange := range divided {
		outputRange := []int{}
		for _, input := range inputRange {
			outputRange = append(outputRange, findDest(input, seedMaps))
		}
		outputs = append(outputs, outputRange)
	}
	return outputs
}

func RunPart2(seeds []int, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemp, tempToHumid, humidToLocation []RangeMap) int {
	outputs := []int{}
	for i := 0; i < len(seeds); i += 2 {
		seedStart := seeds[i]
		seedRange := seeds[i+1]
		endSeed := seedStart + seedRange
		soilOutputs := divideAndMap([][]int{
			{seedStart, endSeed},
		}, seedToSoil)
		fertilizerOutputs := divideAndMap(soilOutputs, soilToFertilizer)
		waterOutputs := divideAndMap(fertilizerOutputs, fertilizerToWater)
		lightOutputs := divideAndMap(waterOutputs, waterToLight)
		tempOutputs := divideAndMap(lightOutputs, lightToTemp)
		humidOutputs := divideAndMap(tempOutputs, tempToHumid)
		locationOutputs := divideAndMap(humidOutputs, humidToLocation)
		for _, outputPairs := range locationOutputs {
			for _, output := range outputPairs {
				//if output < 10 {
				//	continue
				//}
				outputs = append(outputs, output)
			}
		}
	}
	return slices.Min(outputs)
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
	part1 := Run(seeds, maps[0], maps[1], maps[2], maps[3], maps[4], maps[5], maps[6])
	part2Output := RunPart2(seeds, maps[0], maps[1], maps[2], maps[3], maps[4], maps[5], maps[6])
	fmt.Println("Part 1 ==>:", part1, "Part 2 ==>", part2Output)
}

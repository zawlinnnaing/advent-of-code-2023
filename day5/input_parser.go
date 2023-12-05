package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ParseFile(input string) ([]int, [][]RangeMap, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	seeds := []int{}
	seedToSoil := []RangeMap{}
	soilToFertilizer := []RangeMap{}
	fertilizerToWater := []RangeMap{}
	waterToLight := []RangeMap{}
	lightToTemp := []RangeMap{}
	tempToHumid := []RangeMap{}
	humidToLocation := []RangeMap{}
	currentMap := ""
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "seeds:") {
			numText := strings.ReplaceAll(text, "seeds:", "")
			numStrs := strings.Split(numText, " ")
			for _, numStr := range numStrs {
				num, err := strconv.Atoi(numStr)
				if err == nil {
					//return nil, nil, fmt.Errorf("failed to parse seeds: %v, %w", numStrs, err)
					seeds = append(seeds, num)
				}
			}
		}
		if strings.Contains(text, "map:") {
			currentMap = text
			continue
		}
		if text == "\n" {
			continue
		}
		numStrs := strings.Split(text, " ")
		if len(numStrs) != 3 {
			continue
		}

		nums := make([]int, 3)
		for idx, numStr := range numStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, nil, err
			}
			nums[idx] = num
		}
		rangeMap := RangeMap{
			sourceStart: nums[1],
			destStart:   nums[0],
			rangeLength: nums[2],
		}
		if strings.Contains(currentMap, "seed-to-soil") {
			seedToSoil = append(seedToSoil, rangeMap)
		}
		if strings.Contains(currentMap, "soil-to-fertilizer") {
			soilToFertilizer = append(soilToFertilizer, rangeMap)
		}
		if strings.Contains(currentMap, "fertilizer-to-water") {
			fertilizerToWater = append(fertilizerToWater, rangeMap)
		}
		if strings.Contains(currentMap, "water-to-light") {
			waterToLight = append(waterToLight, rangeMap)
		}
		if strings.Contains(currentMap, "light-to-temp") {
			lightToTemp = append(lightToTemp, rangeMap)
		}
		if strings.Contains(currentMap, "temperature-to-humidity") {
			tempToHumid = append(tempToHumid, rangeMap)
		}
		if strings.Contains(currentMap, "humidity-to-location") {
			humidToLocation = append(humidToLocation, rangeMap)
		}
	}

	return seeds, [][]RangeMap{
		seedToSoil,
		soilToFertilizer,
		fertilizerToWater,
		waterToLight,
		lightToTemp,
		tempToHumid,
		humidToLocation,
	}, nil
}

package main

import (
	"strconv"
	"strings"
)

func populateHashMap(hashMap *HashMap, input string) {
	if strings.Contains(input, "=") {
		labelsAndFocal := strings.Split(input, "=")
		focalPower, err := strconv.Atoi(labelsAndFocal[1])
		if err != nil {
			panic(err)
		}
		lens := LabelLens{
			label: labelsAndFocal[0],
			focal: focalPower,
		}
		boxNum := Hash(lens.label)
		hashMap.Insert(boxNum, lens)
		return
	}
	label := strings.Split(input, "-")[0]
	boxNum := Hash(label)
	hashMap.Remove(boxNum, label)
}

func RunPart2(inputs []string) int {
	hashMap := NewHashMap()
	for _, input := range inputs {
		populateHashMap(hashMap, input)
	}
	total := 0
	for boxNum, lenses := range *hashMap {
		for lensIdx, lens := range lenses {
			total += (boxNum + 1) * (lensIdx + 1) * lens.focal
		}
	}
	return total
}

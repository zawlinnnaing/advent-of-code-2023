package main

import (
	"fmt"
)

type Universe [][]string

const EMPTY_SPACE = "."
const GALAXY = "#"

func findEmptyRowsAndCols(galaxies Universe) ([]int, []int) {
	colsWithGalaxy := make(map[int]bool)
	rowsWithGalaxy := make(map[int]bool)
	for rowIdx, galaxyRow := range galaxies {
		for colIdx, cell := range galaxyRow {
			if cell == GALAXY {
				colsWithGalaxy[colIdx] = true
				rowsWithGalaxy[rowIdx] = true
			}
		}
	}
	emptyRows := []int{}
	for rowIdx, _ := range galaxies {
		if !rowsWithGalaxy[rowIdx] {
			emptyRows = append(emptyRows, rowIdx)
		}
	}
	emptyCols := []int{}
	for colIdx, _ := range galaxies[0] {
		if !colsWithGalaxy[colIdx] {
			emptyCols = append(emptyCols, colIdx)
		}
	}
	return emptyRows, emptyCols
}

func expand(galaxies Universe) Universe {
	emptyRows, emptyCols := findEmptyRowsAndCols(galaxies)
	for i := len(emptyRows) - 1; i >= 0; i-- {
		rowIdx := emptyRows[i]
		emptyRow := make([]string, len(galaxies[rowIdx]))
		for j := 0; j < len(emptyRow); j++ {
			emptyRow[j] = EMPTY_SPACE
		}
		galaxies = append(galaxies[:rowIdx], append([][]string{emptyRow}, galaxies[rowIdx:]...)...)
	}

	for i := len(emptyCols) - 1; i >= 0; i-- {
		colIdx := emptyCols[i]
		for rowIdx, row := range galaxies {
			galaxies[rowIdx] = append(row[:colIdx], append([]string{"."}, row[colIdx:]...)...)
		}
	}
	return galaxies
}
func findPairs(galaxyCoords [][2]int) [][2][2]int {
	pairs := [][2][2]int{}
	for idx, galaxyCoord := range galaxyCoords {
		for i := idx + 1; i < len(galaxyCoords); i++ {
			pair := [2][2]int{
				galaxyCoord, galaxyCoords[i],
			}
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

func RunPart1(universe Universe) int {
	expanded := expand(universe)
	galaxyCoords := [][2]int{}
	for rowIdx, row := range expanded {
		for colIdx, cell := range row {
			if cell == GALAXY {
				galaxyCoords = append(galaxyCoords, [2]int{rowIdx, colIdx})
			}
		}
	}
	pairs := findPairs(galaxyCoords)
	fmt.Println("len pairs", len(pairs))
	totalSteps := 0
	for _, pair := range pairs {
		totalSteps += findSteps(pair[0], pair[1])
	}
	return totalSteps
}

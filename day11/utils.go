package main

type Universe [][]string

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

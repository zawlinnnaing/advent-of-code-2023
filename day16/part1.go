package main

func RunPart1(grid Grid) int {
	return CalculateTotalEnergizedTiles(grid, [2]int{0, 0}, "right")
}

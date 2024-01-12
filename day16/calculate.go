package main

func CalculateTotalEnergizedTiles(grid Grid, startPoint [2]int, direction string) int {
	seen := [][]bool{}
	for _, row := range grid {
		seen = append(seen, make([]bool, len(row)))
	}
	Walk(grid, &seen, startPoint, direction, &[]Path{})
	total := 0
	for _, row := range seen {
		for _, cell := range row {
			if cell {
				total += 1
			}
		}
	}
	return total
}

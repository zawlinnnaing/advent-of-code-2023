package main

func RunPart1(grid Grid) int {
	seen := [][]bool{}
	for _, row := range grid {
		seen = append(seen, make([]bool, len(row)))
	}
	Walk(grid, &seen, [2]int{0, 0}, "right", &[]Path{})
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

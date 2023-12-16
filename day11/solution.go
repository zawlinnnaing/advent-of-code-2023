package main

func Solve(universe Universe, stepsToExpand int) int {
	emptyRows, emptyCols := findEmptyRowsAndCols(universe)
	galaxyCoords := findGalaxyCoords(universe)
	pairs := findPairs(galaxyCoords)
	totalSteps := 0
	for _, pair := range pairs {
		totalSteps += findSteps(pair[0], pair[1], emptyRows, emptyCols, stepsToExpand-1)
	}
	return totalSteps
}

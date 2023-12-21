package main

func CheckReflectionWithSmudge(pattern Pattern, firstRow, secondRow int, alreadyHaveSmudged bool) bool {
	currentRow, nextRow := firstRow, secondRow
	totalErr := 0
	if alreadyHaveSmudged {
		totalErr += 1
	}
	for currentRow >= 0 && nextRow < len(pattern) {
		isEqual, err := CheckEqualWithSmudge(pattern[currentRow], pattern[nextRow])
		totalErr += err
		if !isEqual || totalErr > 1 {
			return false
		}
		currentRow -= 1
		nextRow += 1
	}
	return totalErr == 1
}

func CheckEqualWithSmudge(row1 []string, row2 []string) (bool, int) {
	totalErr := 0
	for idx, cell1 := range row1 {
		cell2 := row2[idx]
		if cell1 != cell2 {
			totalErr += 1
		}
		if totalErr > 1 {
			return false, totalErr
		}
	}
	return true, totalErr
}

func FindReflectPointWithSmudge(pattern Pattern) (int, int) {
	rows := len(pattern)
	possiblePoints := [2]int{0, 0}
	for i := 0; i < rows-1; i++ {
		currentRow := pattern[i]
		nextRow := pattern[i+1]
		isEqualWithSmudge, err := CheckEqualWithSmudge(currentRow, nextRow)
		if isEqualWithSmudge && CheckReflectionWithSmudge(pattern, i-1, i+2, err > 0) {
			if i == 0 {
				possiblePoints = [2]int{i, i + 1}
				continue
			}
			return i, i + 1
		}
	}
	return possiblePoints[0], possiblePoints[1]
}

func runPart2(patterns []Pattern) int {
	total := 0
	for _, pattern := range patterns {
		firstPoint, secondPoint := FindReflectPointWithSmudge(pattern)
		//Row reflection point
		if firstPoint != secondPoint {
			total += secondPoint * 100
			continue
		}
		//	Col reflection point
		firstPoint, secondPoint = FindReflectPointWithSmudge(pattern.Transpose())
		total += secondPoint
	}
	return total
}

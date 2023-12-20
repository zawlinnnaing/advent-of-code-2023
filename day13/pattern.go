package main

type Pattern [][]string

func (p Pattern) GetCol(idx int) []string {
	result := []string{}
	for _, row := range p {
		for i := 0; i < len(row); i++ {
			if idx == i {
				result = append(result, row[i])
			}
		}
	}
	return result
}

func (p Pattern) Transpose() Pattern {
	rows := len(p)
	cols := len(p[0])

	transposed := make(Pattern, cols)

	for i := 0; i < cols; i++ {
		transposed[i] = make([]string, rows)
		for j := 0; j < rows; j++ {
			transposed[i][j] = p[j][i]
		}
	}
	return transposed
}

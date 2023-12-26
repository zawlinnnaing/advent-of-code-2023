package main

import (
	"fmt"
	"strings"
)

type Plane [][]string

const (
	BLOCK = "#"
	EMPTY = "."
	ROCK  = "O"
)

var directions = [][2]int{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func (p *Plane) Tilt() {
	maxRows := len(*p)
	maxCols := len((*p)[0])
	for rowIdx, row := range *p {
		for colIdx, _ := range row {
			p.MoveByOffset(rowIdx, colIdx, directions[0], maxRows, maxCols)
		}
	}
}

func (p *Plane) Spin(direction [2]int) {
	maxRows := len(*p)
	maxCols := len((*p)[0])
	if direction[0] > 0 {
		for rowIdx := maxRows - 1; rowIdx >= 0; rowIdx-- {
			row := (*p)[rowIdx]
			for colIdx, _ := range row {
				p.MoveByOffset(rowIdx, colIdx, direction, maxRows, maxCols)
			}
		}
		return
	}
	if direction[1] > 0 {
		for rowIdx, _ := range *p {
			for colIdx := maxCols - 1; colIdx >= 0; colIdx-- {
				p.MoveByOffset(rowIdx, colIdx, direction, maxRows, maxCols)
			}
		}
		return
	}
	for rowIdx, row := range *p {
		for colIdx, _ := range row {
			p.MoveByOffset(rowIdx, colIdx, direction, maxRows, maxCols)
		}
	}
}

func (p *Plane) String() string {
	str := ""
	for _, row := range *p {
		str += fmt.Sprintf("%v\n", row)
	}
	return str
}

func (p *Plane) parse(plane string) {
	var newPlane Plane
	rows := strings.Split(strings.ReplaceAll(strings.ReplaceAll(plane, "[", ""), "]", ""), "\n")

	for _, row := range rows {
		if row == "" {
			continue
		}
		newPlane = append(newPlane, strings.Split(row, " "))
	}
	*p = newPlane
}

func (p *Plane) TotalLoad() int {
	totalLoad := 0
	maxRows := len(*p)
	for rowIdx, row := range *p {
		loadFactor := maxRows - rowIdx
		for _, cell := range row {
			if cell != ROCK {
				continue
			}
			totalLoad += loadFactor
		}
	}
	return totalLoad
}

func (p *Plane) MoveByOffset(rowIdx, colIdx int, offset [2]int, maxRows, maxCols int) {
	currentCell := (*p)[rowIdx][colIdx]
	if currentCell != ROCK {
		return
	}
	nextRowIdx, nextColIdx := rowIdx+offset[0], colIdx+offset[1]
	if nextRowIdx >= maxRows || nextColIdx >= maxCols || nextRowIdx < 0 || nextColIdx < 0 {
		return
	}
	nextCell := (*p)[nextRowIdx][nextColIdx]
	if nextCell != EMPTY {
		return
	}
	(*p)[nextRowIdx][nextColIdx] = (*p)[rowIdx][colIdx]
	(*p)[rowIdx][colIdx] = EMPTY
	p.MoveByOffset(nextRowIdx, nextColIdx, offset, maxRows, maxCols)
}

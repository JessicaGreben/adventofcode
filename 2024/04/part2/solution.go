package main

import (
	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int64, error) {
	m, err := fileinput.ConvertToMatrix(file)
	if err != nil {
		return -1, err
	}

	var total int64
	for r := range m {
		for c := range m[r] {
			if m[r][c] == "M" {
				for direction := range xDirections {
					aRow, aCol := r+direction.r, c+direction.c
					if !inBounds(m, aRow, aCol) {
						continue
					}
					if m[aRow][aCol] != "A" {
						continue
					}
					sRow, sCol := aRow+direction.r, aCol+direction.c
					if !inBounds(m, sRow, sCol) {
						continue
					}
					if m[sRow][sCol] != "S" {
						continue
					}
					if isXmas(m, r, c, direction) {
						total++
					}
				}
			}
		}
	}
	return total / 2, nil
}

type pos struct {
	r, c int
}

type xMasLocation struct {
	mLocation pos
	dir       pos
}

var xDirections = map[pos][]xMasLocation{
	pos{1, 1}: []xMasLocation{
		xMasLocation{pos{2, 0}, pos{-1, 1}},
		xMasLocation{pos{0, 2}, pos{1, -1}},
	},
	pos{1, -1}: []xMasLocation{
		xMasLocation{pos{0, -2}, pos{1, 1}},
		xMasLocation{pos{2, 0}, pos{-1, -1}},
	},
	pos{-1, 1}: []xMasLocation{
		xMasLocation{pos{-2, 0}, pos{1, 1}},
		xMasLocation{pos{0, 2}, pos{-1, -1}},
	},
	pos{-1, -1}: []xMasLocation{
		xMasLocation{pos{-2, 0}, pos{1, -1}},
		xMasLocation{pos{0, -2}, pos{-1, 1}},
	},
}

// isMas returns true if there is a second "MAS" that is in a valid position.
func isXmas(m [][]string, r, c int, direction pos) bool {
	for _, x := range xDirections[direction] {
		dir := x.mLocation
		xRow, xCol := r+dir.r, c+dir.c
		if !inBounds(m, xRow, xCol) {
			continue
		}
		if m[xRow][xCol] != "M" {
			continue
		}

		nextDir := x.dir

		aRow, aCol := xRow+nextDir.r, xCol+nextDir.c
		if !inBounds(m, aRow, aCol) {
			continue
		}
		if m[aRow][aCol] != "A" {
			continue
		}
		sRow, sCol := aRow+nextDir.r, aCol+nextDir.c
		if !inBounds(m, sRow, sCol) {
			continue
		}
		if m[sRow][sCol] != "S" {
			continue
		}
		return true
	}
	return false
}

func inBounds(m [][]string, r, c int) bool {
	return r >= 0 && r < len(m) && c >= 0 && c < len(m[0])
}

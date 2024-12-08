package main

import (
	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
)

type pos struct {
	r, c int
}

func solution(file string) (int64, error) {
	m, err := fileinput.ConvertToMatrix(file)
	if err != nil {
		return -1, err
	}

	antennaLocations := map[string][]pos{}
	for r := range m {
		for c := range m[r] {
			curr := m[r][c]
			if curr == "." {
				continue
			}
			antennaLocations[curr] = append(antennaLocations[curr], pos{r, c})
		}
	}

	antinodes := map[pos]int{}
	for _, locations := range antennaLocations {
		for i := range locations {
			for j := i + 1; j < len(locations); j++ {
				a, b := locations[i], locations[j]
				for _, n := range getAntinodes(m, a, b) {
					antinodes[n]++
				}
			}
		}
	}
	return int64(len(antinodes)), nil
}

func getAntinodes(m [][]string, a, b pos) []pos {
	result := []pos{}
	diffRow, diffCol := b.r-a.r, a.c-b.c
	if diffRow == 0 || diffCol < 0 {
		diffCol = b.c - a.c
	}

	var row1Dir, col1Dir, row2Dir, col2Dir int
	switch {
	case diffCol == 0 || diffRow == 0 || diffCol < 0:
		row1Dir, col1Dir = -1, -1
		row2Dir, col2Dir = 1, 1
	case diffCol > 0:
		diffCol = a.c - b.c
		row1Dir, col1Dir = -1, 1
		row2Dir, col2Dir = 1, -1
	}

	antinodeRow, antinodeCol := a.r+(row1Dir*diffRow), a.c+(col1Dir*diffCol)
	if inBounds(m, antinodeRow, antinodeCol) {
		result = append(result, pos{antinodeRow, antinodeCol})
	}
	antinodeRow, antinodeCol = b.r+(row2Dir*diffRow), b.c+(col2Dir*diffCol)
	if inBounds(m, antinodeRow, antinodeCol) {
		result = append(result, pos{antinodeRow, antinodeCol})
	}
	return result
}

func inBounds(m [][]string, r, c int) bool {
	return r >= 0 && r < len(m) && c >= 0 && c < len(m[r])
}

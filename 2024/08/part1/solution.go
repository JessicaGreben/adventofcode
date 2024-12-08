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
				getAntinodes(m, a, b, antinodes)
				getAntinodes(m, b, a, antinodes)
			}
		}
	}
	return int64(len(antinodes)), nil
}

func getAntinodes(m [][]string, a, b pos, antinodes map[pos]int) {
	diffRow, diffCol := a.r-b.r, a.c-b.c
	antinodeRow, antinodeCol := a.r+diffRow, a.c+diffCol
	if inBounds(m, antinodeRow, antinodeCol) {
		antinodes[pos{antinodeRow, antinodeCol}]++
		antinodeRow, antinodeCol = antinodeRow+diffRow, antinodeCol+diffCol
	}
}
func inBounds(m [][]string, r, c int) bool {
	return r >= 0 && r < len(m) && c >= 0 && c < len(m[r])
}

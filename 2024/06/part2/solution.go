package main

import (
	"fmt"

	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int64, error) {
	m, err := fileinput.ConvertToMatrix(file)
	if err != nil {
		return -1, err
	}

	var startDirection pos
	var startRow, startCol int
	for r := range m {
		for c := range m[r] {
			curr := m[r][c]
			if direction, ok := locationPointerToDirections[curr]; ok {
				startDirection = direction
				startRow, startCol = r, c
			}
		}
	}

	var result int64
	for r := range m {
		for c := range m[r] {
			cp := copyMatrix(m)
			curr := cp[r][c]
			if curr == "." {
				cp[r][c] = blocked
			} else {
				continue
			}
			cp[startRow][startCol] = "."
			if traverse(cp, startRow, startCol, startDirection, map[seenKey]bool{}) {
				result++
			}
		}
	}

	return result, nil
}

func copyMatrix(m [][]string) [][]string {
	cp := make([][]string, len(m))
	for i := range m {
		cp[i] = make([]string, len(m[i]))
		copy(cp[i], m[i])
	}
	return cp
}

type pos struct {
	r, c   int
	symbol string
}

var locationPointerToDirections = map[string]pos{
	"<": pos{0, -1, "<"},
	">": pos{0, 1, ">"},
	"^": pos{-1, 0, "^"},
	"v": pos{1, 0, "v"},
}

var nextDirection = map[string]string{
	"<": "^",
	">": "v",
	"^": ">",
	"v": "<",
}

const (
	blocked = "#"
	visited = "X"
)

type seenKey struct {
	currRow, currCol int
	dir              pos
}

func traverse(m [][]string, r, c int, direction pos, seen map[seenKey]bool) bool {
	if !inBounds(m, r, c) {
		return false
	}

	key := seenKey{r, c, direction}
	if _, ok := seen[key]; ok {
		return true
	}

	if isBlocked(m, r, c) {
		nextDirectionPointer := nextDirection[direction.symbol]
		nextDirection := locationPointerToDirections[nextDirectionPointer]
		prevRow, prevCol := r-direction.r, c-direction.c
		return traverse(m, prevRow, prevCol, nextDirection, seen)
	}

	if m[r][c] == "." {
		m[r][c] = visited
		seen[key] = true
	}
	nextRow, nextCol := r+direction.r, c+direction.c
	return traverse(m, nextRow, nextCol, direction, seen)
}

func isBlocked(m [][]string, r, c int) bool {
	return m[r][c] == blocked
}

func inBounds(m [][]string, r, c int) bool {
	return r >= 0 && r < len(m) && c >= 0 && c < len(m[r])
}

func print(m [][]string) {
	for r := range m {
		fmt.Println(m[r])
	}
	fmt.Println()
}

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

	for r := range m {
		for c := range m[r] {
			curr := m[r][c]
			direction, ok := locationPointerToDirections[curr]
			if ok {
				m[r][c] = "."
				return traverse(m, r, c, direction, 0)
			}
		}
	}

	return -1, nil
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

// can you cross a path you've already been
// can you turn more than once
// dont double count a path you have already crossed
func traverse(m [][]string, r, c int, direction pos, count int64) (int64, error) {
	if !inBounds(m, r, c) {
		return count, nil
	}

	if isBlocked(m, r, c) {
		nextDirectionPointer := nextDirection[direction.symbol]
		nextDirection := locationPointerToDirections[nextDirectionPointer]
		prevRow, prevCol := r-direction.r, c-direction.c
		return traverse(m, prevRow, prevCol, nextDirection, count)
	}

	if m[r][c] == "." {
		m[r][c] = visited
		count++
	}

	nextRow, nextCol := r+direction.r, c+direction.c
	return traverse(m, nextRow, nextCol, direction, count)
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
}

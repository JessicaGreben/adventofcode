package main

import (
	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
	"github.com/jessicagreben/adventofcode/pkg/matrix"
)

func solution(file string, moves string) (int64, error) {
	m, err := fileinput.ConvertToMatrix(file)
	if err != nil {
		return -1, err
	}
	var startRow, startCol int
	for r := range m {
		for c := range m[r] {
			if m[r][c] == "@" {
				startRow, startCol = r, c
				break
			}
		}
	}

	traverse(m, moves, startRow, startCol, 0)

	var result int64
	for r := range m {
		for c := range m[r] {
			if m[r][c] == "O" {
				result += int64(100*r + c)
			}
		}
	}
	return result, nil
}

var moveDirections = map[byte]matrix.Element{
	'<': {0, -1},
	'>': {0, 1},
	'v': {1, 0},
	'^': {-1, 0},
}

func traverse(m [][]string, moves string, r, c, moveIdx int) {
	if moveIdx >= len(moves) {
		return
	}

	currMove := moves[moveIdx]
	currDirection := moveDirections[currMove]
	nextRow, nextCol := r+currDirection.Row, c+currDirection.Col
	nr, nc := r, c

	switch next := m[nextRow][nextCol]; next {
	case "O":
		if ok := moveBox(m, nextRow, nextCol, currDirection); ok {
			swap(m, r, c, nextRow, nextCol)
			nr, nc = nextRow, nextCol
		}
	case ".":
		swap(m, r, c, nextRow, nextCol)
		nr, nc = nextRow, nextCol
	}

	traverse(m, moves, nr, nc, moveIdx+1)
}

func moveBox(m [][]string, r, c int, direction matrix.Element) bool {
	nextRow, nextCol := r+direction.Row, c+direction.Col
	switch next := m[nextRow][nextCol]; next {
	case ".":
		swap(m, r, c, nextRow, nextCol)
		return true
	case "O":
		if ok := moveBox(m, nextRow, nextCol, direction); ok {
			swap(m, r, c, nextRow, nextCol)
			return true
		}
	}
	return false
}

func swap(m [][]string, currRow, currCol, r, c int) {
	m[currRow][currCol], m[r][c] = m[r][c], m[currRow][currCol]
}

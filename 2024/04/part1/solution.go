package main

import (
	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int64, error) {
	m, err := fileinput.ConvertToMatrix(file)
	if err != nil {
		return -1, err
	}

	allDirections := [][]int{
		{1, 0}, {-1, 0}, // row
		{0, 1}, {0, -1}, // col
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1}, // diagonal
	}

	var total int64
	for r := range m {
		for c := range m[r] {
			if m[r][c] == "X" {
				for _, direction := range allDirections {
					nextRow, nextCol := r+direction[0], c+direction[1]
					if !inBounds(m, nextRow, nextCol) {
						continue
					}
					total += countXmas(m, nextRow, nextCol, 1, direction)
				}
			}
		}
	}
	return total, nil
}

const XMAS = "XMAS"

func countXmas(m [][]string, r, c, i int, direction []int) int64 {
	if i == len(XMAS)-1 {
		if m[r][c] == string(XMAS[i]) {
			return 1
		}
		return 0
	}

	if m[r][c] != string(XMAS[i]) {
		return 0
	}
	nextRow, nextCol := r+direction[0], c+direction[1]
	if !inBounds(m, nextRow, nextCol) {
		return 0
	}
	return countXmas(m, nextRow, nextCol, i+1, direction)
}

func inBounds(m [][]string, r, c int) bool {
	return r >= 0 && r < len(m) && c >= 0 && c < len(m[0])
}

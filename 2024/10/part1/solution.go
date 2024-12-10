package main

import (
	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
	"github.com/jessicagreben/adventofcode/pkg/matrix"
)

type pos struct {
	r, c int
}

func solution(file string) (int64, error) {
	m, err := fileinput.ConvertToMatrix(file)
	if err != nil {
		return -1, err
	}

	var result int64
	for r := range m {
		for c := range m[r] {
			if m[r][c] == "0" {
				scores := map[pos]bool{}
				getScore(m, r, c, scores)
				result += int64(len(scores))
			}
		}
	}
	return result, nil
}

var nextValue = map[string]string{
	"0": "1",
	"1": "2",
	"2": "3",
	"3": "4",
	"4": "5",
	"5": "6",
	"6": "7",
	"7": "8",
	"8": "9",
}

var directions = []pos{
	{1, 0}, {-1, 0}, {0, 1}, {0, -1},
}

func getScore(m [][]string, r, c int, scores map[pos]bool) {
	if m[r][c] == "9" {
		scores[pos{r, c}] = true
		return
	}
	next := nextValue[m[r][c]]
	for _, dir := range directions {
		nextRow, nextCol := r+dir.r, c+dir.c
		if !matrix.InBounds(m, nextRow, nextCol) {
			continue
		}
		if m[nextRow][nextCol] == next {
			getScore(m, nextRow, nextCol, scores)
		}
	}
}

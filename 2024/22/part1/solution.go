package main

import (
	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int64, error) {
	input, err := fileinput.New(file)
	if err != nil {
		return -1, err
	}

	var result int64
	for line := range input.All() {
		lineInts, err := parseLine(line)
		if err != nil {
			return -1, err
		}
		result += x(lineInts[0])

	}
	return result, nil
}

func x(initSecret int64) int64 {
	curr := initSecret
	for i := 0; i < 2000; i++ {
		// 1
		x := curr * 64
		curr ^= x
		curr %= 16777216
		// 2
		y := curr / 32
		curr ^= y
		curr %= 16777216
		// 3
		z := curr * 2048
		curr ^= z
		curr %= 16777216
	}
	return curr
}

func parseLine(line string) ([]int64, error) {
	return fileinput.ParseLineInt64(line, " ", -1)
}

package main

import (
	"github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int64, error) {
	return input.ForEachLineInt64(file, processLine)
}

func processLine(line string) (int64, error) {
	return -1, nil
}

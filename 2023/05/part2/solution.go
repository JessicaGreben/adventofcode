package main

import (
	"github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int, error) {
	return input.ForEachLine(file, processLine)
}

func processLine(line string) (int, error) {
	return -1, nil
}

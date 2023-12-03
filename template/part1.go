package main

import (
	"github.com/jessicagreben/adventofcode/pkg/input"
)

func part1(file string) (int, error) {
	return input.ForEachLine(file, part1ProcessLine)
}

func part1ProcessLine(line string) (int, error) {
	return -1, nil
}

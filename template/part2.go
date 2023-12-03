package main

import (
	"github.com/jessicagreben/adventofcode/pkg/input"
)

func part2(file string) (int, error) {
	return input.ForEachLine(file, part2ProcessLine)
}

func part2ProcessLine(line string) (int, error) {
	return -1, nil
}

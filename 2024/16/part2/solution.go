package main

import (
	"fmt"

	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int64, error) {
	input, err := fileinput.New(file)
	if err != nil {
		return -1, err
	}

	for line := range input.All() {
		lintInts, err := parseLine(line)
		if err != nil {
			return -1, err
		}
		part1, part2 := lintInts[0], lintInts[1]
		fmt.Println(part1, part2)
	}
	return -1, nil
}

func parseLine(line string) ([]int64, error) {
	return fileinput.ParseLineInt64(line, " ", 2)
}

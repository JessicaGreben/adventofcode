package main

import (
	"fmt"

	"github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int64, error) {
	in, err := input.NewInput(file)
	if err != nil {
		return -1, err
	}

	for line := range in.All() {
		linePartInts, err := processLine(line)
		if err != nil {
			return -1, err
		}
		part1, part2 := linePartInts[0], linePartInts[1]
		fmt.Println(part1, part2)
	}
	return -1, nil
}

func processLine(line string) ([]int64, error) {
	return input.ParseLineInt64(line, " ", 2)
}

package main

import (
	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int64, error) {
	input, err := fileinput.New(file)
	if err != nil {
		return -1, err
	}

	var safeCount int64
	for line := range input.All() {
		lineInts, err := parseLine(line)
		if err != nil {
			return -1, err
		}

		if isSafe(lineInts) {
			safeCount++
		}
	}
	return safeCount, nil
}

// isSafe returns true if both of the following are true:
// - the levels are either all increasing or all decreasing.
// - any two adjacent levels differ by at least one and at most three.
func isSafe(lineInts []int64) bool {
	return (isIncreasing(lineInts) || isDecreasing(lineInts)) &&
		adjacentLevelsOK(lineInts)
}

// Any two adjacent levels differ by at least one and at most three.
func adjacentLevelsOK(lineInts []int64) bool {
	for i := range len(lineInts) - 1 {
		c, n := lineInts[i], lineInts[i+1]
		if abs(c-n) < 1 || abs(c-n) > 3 {
			return false
		}
	}
	return true
}

func isIncreasing(lineInts []int64) bool {
	for i := range len(lineInts) - 1 {
		c, n := lineInts[i], lineInts[i+1]
		if c >= n {
			return false
		}
	}
	return true
}

func isDecreasing(lineInts []int64) bool {
	for i := range len(lineInts) - 1 {
		c, n := lineInts[i], lineInts[i+1]
		if c <= n {
			return false
		}
	}
	return true
}

func parseLine(line string) ([]int64, error) {
	return fileinput.ParseLineInt64(line, " ", -1)
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

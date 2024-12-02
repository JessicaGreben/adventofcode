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

		if isSafe(lineInts, -1) {
			safeCount++
		}
	}
	return safeCount, nil
}

func isSafe(lineInts []int64, skipIdx int) bool {
	if len(lineInts) <= 2 {
		return true
	}

	i, j := 0, 1
	if skipIdx == 0 {
		i, j = 1, 2
	}
	if skipIdx == 1 {
		j = 2
	}
	allIncreasing := lineInts[i] < lineInts[j]
	allDecreasing := lineInts[i] > lineInts[j]
	for j < len(lineInts) {
		if i == skipIdx {
			i++
			if i == j {
				j++
			}
			continue
		}
		if j == skipIdx {
			j++
			continue
		}
		c, n := lineInts[i], lineInts[j]
		if c < n && allIncreasing {
			// safe
		} else if c > n && allDecreasing {
			// safe
		} else {
			if skipIdx >= 0 {
				return false
			}
			if i == 0 {
				return isSafe(lineInts, i) ||
					isSafe(lineInts, j)
			}
			return isSafe(lineInts, i) ||
				isSafe(lineInts, i-1) ||
				isSafe(lineInts, j)
		}
		if abs(c-n) < 1 || abs(c-n) > 3 {
			if skipIdx >= 0 {
				return false
			}
			if i == 0 {
				return isSafe(lineInts, i) ||
					isSafe(lineInts, j)
			}
			return isSafe(lineInts, i) ||
				isSafe(lineInts, i-1) ||
				isSafe(lineInts, j)
		}
		i++
		j++
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

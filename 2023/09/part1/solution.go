package main

import (
	"strconv"
	"strings"

	"github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int64, error) {
	return input.ForEachLineInt64(file, processLine)
}

func processLine(line string) (int64, error) {
	ln := strings.Split(line, " ")
	prev := make([]int64, 0, len(ln))
	for _, n := range ln {
		x, err := strconv.Atoi(n)
		if err != nil {
			return -1, err
		}
		prev = append(prev, int64(x))
	}

	lastVals := []int64{prev[len(prev)-1]}
	var allZeros bool
	for !allZeros {
		allZeros = true
		var diff int64
		next := make([]int64, 0, len(prev)-1)
		for i, j := 0, 1; j < len(prev); i, j = i+1, j+1 {
			currDiff := prev[j] - prev[i]
			if currDiff != 0 {
				allZeros = false
			}
			diff = currDiff
			next = append(next, diff)
		}
		lastVals = append(lastVals, diff)
		prev = next
	}
	var sum int64 = lastVals[len(lastVals)-2]
	for i := len(lastVals) - 3; i >= 0; i-- {
		sum += lastVals[i]
	}
	return sum, nil
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

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

	firstVals := []int64{prev[0]}
	var allZeros bool
	for !allZeros {
		allZeros = true
		var diff int64
		next := make([]int64, 0, len(prev)-1)
		for i, j := 0, 1; j < len(prev); i, j = i+1, j+1 {
			diff = prev[i] - prev[j]
			if diff != 0 {
				allZeros = false
			}
			next = append(next, diff)
		}
		firstVals = append(firstVals, next[0])
		prev = next
	}
	var sum int64 = firstVals[len(firstVals)-2]
	for i := len(firstVals) - 3; i >= 0; i-- {
		sum += firstVals[i]
	}
	return sum, nil
}

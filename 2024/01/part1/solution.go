package main

import (
	"sort"

	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int64, error) {
	input, err := fileinput.New(file)
	if err != nil {
		return -1, err
	}

	leftIDs, rightIDs := []int{}, []int{}
	for line := range input.All() {
		lineParts, err := fileinput.ParseLineInt64(line, "   ", 2)
		if err != nil {
			return -1, err
		}
		leftIDs = append(leftIDs, int(lineParts[0]))
		rightIDs = append(rightIDs, int(lineParts[1]))
	}

	sort.Ints(leftIDs)
	sort.Ints(rightIDs)

	var sumDifference int
	for i := range leftIDs {
		sumDifference += abs(leftIDs[i] - rightIDs[i])
	}
	return int64(sumDifference), nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

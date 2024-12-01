package main

import (
	"github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int64, error) {
	in, err := input.NewInput(file)
	if err != nil {
		return -1, err
	}

	leftIDs := []int64{}
	rightIDsFrequency := map[int64]int64{}
	for line := range in.All() {
		lineParts, err := input.ParseLineInt64(line, "   ", 2)
		if err != nil {
			return -1, err
		}

		leftID, rightID := int(lineParts[0]), int(lineParts[1])
		leftIDs = append(leftIDs, int64(leftID))
		rightIDsFrequency[int64(rightID)]++
	}

	var similarityScore int64
	for _, id := range leftIDs {
		similarityScore += id * rightIDsFrequency[id]
	}
	return similarityScore, nil

}

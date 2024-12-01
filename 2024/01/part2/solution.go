package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution(file string) (int64, error) {
	leftIDs, rightIDsFrequency, err := processInput(file)
	if err != nil {
		return -1, err
	}
	var similarityScore int64
	for _, id := range leftIDs {
		similarityScore += id * rightIDsFrequency[id]
	}
	return similarityScore, nil
}

func processInput(file string) ([]int64, map[int64]int64, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}
	scanner := bufio.NewScanner(fd)

	leftIDs := []int64{}
	rightIDsFrequency := map[int64]int64{}
	for scanner.Scan() {
		line := scanner.Text()
		IDs := strings.Split(line, "   ")
		if len(IDs) != 2 {
			return nil, nil, fmt.Errorf("wrong number of ids, want=2, got=%d, IDs=%v", len(IDs), IDs)
		}
		leftID, rightID := IDs[0], IDs[1]
		leftIDInt, err := strconv.Atoi(leftID)
		if err != nil {
			return nil, nil, err
		}
		rightIDInt, err := strconv.Atoi(rightID)
		if err != nil {
			return nil, nil, err
		}
		leftIDs = append(leftIDs, int64(leftIDInt))
		rightIDsFrequency[int64(rightIDInt)]++
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return leftIDs, rightIDsFrequency, nil
}

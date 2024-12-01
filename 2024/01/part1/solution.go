package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(file string) (int64, error) {
	leftIDs, rightIDs, err := processInput(file)
	if err != nil {
		return -1, err
	}
	var sumDifference int
	for i := range leftIDs {
		sumDifference += abs(leftIDs[i] - rightIDs[i])
	}
	return int64(sumDifference), nil
}

func processInput(file string) ([]int, []int, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}
	scanner := bufio.NewScanner(fd)
	leftIDs, rightIDs := []int{}, []int{}
	for scanner.Scan() {
		line := scanner.Text()
		IDs := strings.Split(line, "   ")
		if len(IDs) != 2 {
			return nil, nil, fmt.Errorf("wrong number of ID, want=2, got=%d, IDs=%v", len(IDs), IDs)
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
		leftIDs = append(leftIDs, leftIDInt)
		rightIDs = append(rightIDs, rightIDInt)
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	if len(leftIDs) != len(rightIDs) {
		return nil, nil, fmt.Errorf("left and right input are not the same len, len(left)=%v, len(right)=%v", len(leftIDs), len(rightIDs))
	}
	sort.Ints(leftIDs)
	sort.Ints(rightIDs)
	return leftIDs, rightIDs, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

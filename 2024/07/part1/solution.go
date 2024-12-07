package main

import (
	"errors"
	"strconv"
	"strings"

	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int64, error) {
	input, err := fileinput.New(file)
	if err != nil {
		return -1, err
	}

	var result int64
	for line := range input.All() {
		testVal, nums, err := parseLine(line)
		if err != nil {
			return -1, err
		}
		if isTrue(nums, 0, "+", 0, testVal) {
			result += testVal
			continue
		}
		if isTrue(nums, 0, "*", 0, testVal) {
			result += testVal
			continue
		}
	}
	return result, nil
}

func isTrue(nums []int64, i int, operator string, curr, desired int64) bool {
	if i >= len(nums) {
		return curr == desired
	}
	if curr > desired {
		return false
	}
	switch {
	case operator == "+":
		curr += nums[i]
	case operator == "*":
		curr *= nums[i]
	}

	return isTrue(nums, i+1, "+", curr, desired) ||
		isTrue(nums, i+1, "*", curr, desired)
}

func parseLine(line string) (int64, []int64, error) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return -1, nil, errors.New("expected two parts")
	}
	testVal := parts[0]
	testValInt, err := strconv.Atoi(testVal)
	if err != nil {
		return -1, nil, err
	}
	nums, err := fileinput.ParseLineInt64(parts[1], " ", -1)
	if err != nil {
		return -1, nil, err
	}

	return int64(testValInt), nums, nil
}

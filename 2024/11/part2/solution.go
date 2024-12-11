package main

import (
	"strconv"

	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
)

type memoKey struct {
	stone      int64
	blinkCount int
}

func solution(file string) (int64, error) {
	input, err := fileinput.New(file)
	if err != nil {
		return -1, err
	}

	var result int64
	memo := map[memoKey]int64{}
	for line := range input.All() {
		stones, err := fileinput.ParseLineInt64(line, " ", -1)
		if err != nil {
			return -1, err
		}

		for _, stone := range stones {
			result += dfs(stone, 0, 75, 1, memo)
		}
	}

	return result, nil
}

func dfs(stone int64, blinkCount, desiredBlinkCount int, totalStoneCount int64, memo map[memoKey]int64) int64 {
	if blinkCount == desiredBlinkCount {
		return totalStoneCount
	}
	v, ok := memo[memoKey{stone, blinkCount}]
	if ok {
		return v
	}

	switch {
	case stone == 0:
		result := dfs(1, blinkCount+1, desiredBlinkCount, totalStoneCount, memo)
		memo[memoKey{stone, blinkCount}] = result
		return result
	case len(strconv.Itoa(int(stone)))%2 == 0:
		firstHalf, secondHalf := splitDigits(stone)
		resultA := dfs(firstHalf, blinkCount+1, desiredBlinkCount, totalStoneCount, memo)
		resultB := dfs(secondHalf, blinkCount+1, desiredBlinkCount, 1, memo)
		memo[memoKey{stone, blinkCount}] = resultA + resultB
		return resultA + resultB
	default:
		result := dfs(stone*2024, blinkCount+1, desiredBlinkCount, totalStoneCount, memo)
		memo[memoKey{stone, blinkCount}] = result
		return result
	}
}

func splitDigits(x int64) (int64, int64) {
	var firstHalf, secondHalf int64
	digits := []int{}
	for x > 0 {
		digits = append(digits, int(x)%10)
		x /= 10
	}
	i := len(digits) / 2
	factor := 1
	for i < len(digits) {
		firstHalf += int64(digits[i]) * int64(factor)
		factor *= 10
		i++
	}
	factor = 1
	i = 0
	for i < len(digits)/2 {
		secondHalf += int64(digits[i]) * int64(factor)
		factor *= 10
		i++
	}
	return firstHalf, secondHalf
}

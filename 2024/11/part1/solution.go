package main

import (
	"strconv"

	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int64, error) {
	input, err := fileinput.New(file)
	if err != nil {
		return -1, err
	}

	var stones []int64
	for line := range input.All() {
		stones, err = fileinput.ParseLineInt64(line, " ", -1)
		if err != nil {
			return -1, err
		}

		for range 25 {
			stones = blink(stones)
		}
	}

	return int64(len(stones)), nil
}

func blink(input []int64) []int64 {
	result := []int64{}
	for _, x := range input {
		switch {
		case x == 0:
			result = append(result, 1)
		case len(strconv.Itoa(int(x)))%2 == 0:
			firstHalf, secondHalf := splitDigits(x)
			result = append(result, firstHalf, secondHalf)
		default:
			result = append(result, x*2024)
		}
	}
	return result
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
